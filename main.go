package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime/types"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

func parseMarkdownWithMeta(filename string) (map[string]interface{}, []string, []string, error) {
	metadata := map[string]interface{}{}
	headers := []string{}
	contents := []string{}

	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)
	source, err := os.ReadFile(filename)
	if err != nil {
		return metadata, headers, contents, err
	}
	var buf bytes.Buffer
	context := parser.NewContext()
	if err := markdown.Convert([]byte(source), &buf, parser.WithContext(context)); err != nil {
		return metadata, headers, contents, err
	}
	metadata = meta.Get(context)
	var currentContent strings.Builder
	reader := text.NewReader(source)
	node := goldmark.DefaultParser().Parse(reader)
	collecting := false

	ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if heading, ok := node.(*ast.Heading); ok && entering {
			if collecting {
				contents = append(contents, strings.TrimSpace(currentContent.String()))
				currentContent.Reset()
			}
			headerText := string(heading.Text(reader.Source()))
			if !strings.HasPrefix(headerText, "title:") {
				headers = append(headers, headerText)
			}
			collecting = false
		} else if heading, ok = node.(*ast.Heading); ok && !entering {
			collecting = true
		} else if collecting && entering {
			if textNode, ok := node.(*ast.Text); ok {
				currentContent.Write(textNode.Segment.Value(reader.Source()))
			} else if codeNode, ok := node.(*ast.CodeBlock); ok {
				currentContent.Write(codeNode.Text(reader.Source()))
			}
		}
		return ast.WalkContinue, nil
	})
	if collecting {
		contents = append(contents, strings.TrimSpace(currentContent.String()))
	}
	return metadata, headers, contents, nil
}

func findMarkdownFiles(dir string) ([]string, error) {
	var markdownFiles []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") && path != "_index.md" && path != "README.md" {
			markdownFiles = append(markdownFiles, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return markdownFiles, nil
}

func toSafeName(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, " ", "-"))
}

func getGRCMappingsFromBedrock(content string) ([]string, error) {
	KNOWLEDGE_BASE_ID := "FMO5PGNIVI"
	KNOWLEDGE_BASE_MODEL_ID := "anthropic.claude-3-sonnet-20240229-v1:0"
	KNOWLEDGE_BASE_NUMBER_OF_RESULT := int32(5)
	mappings := []string{}
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		return mappings, err
	}

	type ClaudeRequest struct {
		Prompt            string   `json:"prompt"`
		MaxTokensToSample int      `json:"max_tokens_to_sample"`
		Temperature       float64  `json:"temperature,omitempty"`
		StopSequences     []string `json:"stop_sequences,omitempty"`
	}

	type ClaudeResponse struct {
		Completion string `json:"completion"`
	}
	if len(content) > 1000 {
		content = content[:1000]
	}
	// Create a BedrockAgentRuntime client
	client := bedrockagentruntime.NewFromConfig(cfg)
	prompt := `You are providing mappings of security and privacy compliance programs to security policy documents. I will provide you with a set of search results relating to security and privacy compliance programs. The user will provide you with a set of security policy statements. Your job is to list the compliance program controls that map to the security policy, using the search results as additional context. You should only provide the policy ID (what is in the heading) and the control framework and ID (e.g. SOC 2 A1.3) - each on a separate line. Don't provide any text description - just return the results as JSON so they can be read by a machine if you are called over an API. Important - only return SOC 2 controls.
                            
	Here are the search results in numbered order:
	$search_results$
	
	$output_format_instructions$`
	// invoke bedrock agent runtime to retrieve and generate
	output, err := client.RetrieveAndGenerate(
		context.Background(),
		&bedrockagentruntime.RetrieveAndGenerateInput{
			Input: &types.RetrieveAndGenerateInput{
				Text: aws.String(content),
			},
			RetrieveAndGenerateConfiguration: &types.RetrieveAndGenerateConfiguration{
				Type: types.RetrieveAndGenerateTypeKnowledgeBase,
				KnowledgeBaseConfiguration: &types.KnowledgeBaseRetrieveAndGenerateConfiguration{
					KnowledgeBaseId: aws.String(KNOWLEDGE_BASE_ID),
					ModelArn:        aws.String(KNOWLEDGE_BASE_MODEL_ID),
					RetrievalConfiguration: &types.KnowledgeBaseRetrievalConfiguration{
						VectorSearchConfiguration: &types.KnowledgeBaseVectorSearchConfiguration{
							NumberOfResults: aws.Int32(KNOWLEDGE_BASE_NUMBER_OF_RESULT),
						},
					},
					GenerationConfiguration: &types.GenerationConfiguration{
						PromptTemplate: &types.PromptTemplate{
							TextPromptTemplate: &prompt,
						},
					},
				},
			},
		},
	)
	if err != nil {
		return mappings, err
	}

	allParts := []string{}
	parts := strings.Split(*output.Output.Text, "\n")
	for _, part := range parts {
		allParts = append(allParts, strings.Split(part, " ")...)
	}
	for _, p := range allParts {
		if strings.HasPrefix(p, "A") || strings.HasPrefix(p, "C") || strings.HasPrefix(p, "P") {
			mappings = append(mappings, p)
		}
	}
	return mappings, nil
}

func rewritePolicyWithLinks(filename string, slug string, sections map[string][]string) error {
	source, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	content := string(source)
	lines := strings.Split(content, "\n")
	var buf bytes.Buffer
	var currentHeading string
	headerPattern := regexp.MustCompile(`^#{1,6}\s*(.*)$`)
	for _, line := range lines {
		buf.WriteString(line + "\n")
		match := headerPattern.FindStringSubmatch(line)
		if match != nil {
			currentHeading = strings.TrimSpace(match[1])
			id := slugAndHeaderToID(slug, currentHeading)
			if links, found := sections[id]; found {
				for _, link := range links {
					additionalContent := fmt.Sprintf("[%s](%s)\n", link, link)
					buf.WriteString(additionalContent)
				}

			}
		}
	}
	buf.WriteString("\n")
	modifiedContent := buf.Bytes()
	err = os.WriteFile(filename, modifiedContent, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}
func slugAndHeaderToID(slug interface{}, header string) string {
	return fmt.Sprintf("%s-%s", slug, toSafeName(header))
}

func alphanumericOnly(input string) string {
	var result []rune
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result = append(result, r)
		}
	}
	return string(result)
}

func graphGRCLink(s string) string {
	filename := alphanumericOnly(strings.ToLower(strings.ReplaceAll(s, "SOC 2 ", "")))
	return fmt.Sprintf("https://alsmola.github.io/graphgrc/soc2/%s.html", filename)
}

func main() {
	markdownFiles, err := findMarkdownFiles(".")
	if err != nil {
		panic(err)
	}
	policyStatements := map[string]string{}
	for _, f := range markdownFiles {
		log.Println(f)
		metadata, headers, contents, err := parseMarkdownWithMeta(f)
		if err != nil {
			panic(err)
		}
		slugI, ok := metadata["slug"]
		slug := fmt.Sprintf("%s", slugI)
		if !ok || slug == string("") {
			continue
		}
		for idx, header := range headers {
			if header != "Scope" && header != "Context" {
				id := slugAndHeaderToID(slug, header)
				if len(contents[idx]) > 5 {
					policyStatements[id] = contents[idx]
				}
			}
		}
		sections := map[string][]string{}
		for id, content := range policyStatements {
			mappings, err := getGRCMappingsFromBedrock(content)
			links := []string{}
			for _, m := range mappings {
				links = append(links, graphGRCLink(m))
			}
			sections[id] = links
			if err != nil {
				panic(err)
			}
		}
		err = rewritePolicyWithLinks(f, slug, sections)
		if err != nil {
			panic(err)
		}
	}
}
