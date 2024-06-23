---
title: Change management policy
slug: change-management
policy: true
faq: false
weight: 9
---

To avoid potential security incidents, Tailscale requires change management controls to ensure only authorized changes are made to its environment and processes.

### Environment
[https://alsmola.github.io/graphgrc/soc2/cc52.html](https://alsmola.github.io/graphgrc/soc2/cc52.html)
[https://alsmola.github.io/graphgrc/soc2/cc81.html](https://alsmola.github.io/graphgrc/soc2/cc81.html)
[https://alsmola.github.io/graphgrc/soc2/cc68.html](https://alsmola.github.io/graphgrc/soc2/cc68.html)

#### Code changes

Changes to code in Tailscale’s environment made by an employee or contractor must be tested and approved by another employee prior to being merged and rolled out.

Tailscale uses branch protection rules on GitHub to require changes be made through a pull request with a second review prior to merging code.

Exceptionally, employees can push changes without a second review where they are required to mitigate an incident. Changes pushed without prior approval are tagged and audited after the fact, within 2 business days.

Changes to update dependencies, publish documentation, changes to the marketing website, or non-substantive code changes are exempt from this policy.

##### Dependencies
[https://alsmola.github.io/graphgrc/soc2/cc81.html](https://alsmola.github.io/graphgrc/soc2/cc81.html)
[https://alsmola.github.io/graphgrc/soc2/cc34.html](https://alsmola.github.io/graphgrc/soc2/cc34.html)
[https://alsmola.github.io/graphgrc/soc2/cc68.html](https://alsmola.github.io/graphgrc/soc2/cc68.html)

Dependencies can be updated without requiring a separate reviewer. Tailscale reviews the release notes for a dependency prior to merging the changes.

##### Documentation
[https://alsmola.github.io/graphgrc/soc2/cc81.html](https://alsmola.github.io/graphgrc/soc2/cc81.html)
[https://alsmola.github.io/graphgrc/soc2/cc72.html](https://alsmola.github.io/graphgrc/soc2/cc72.html)

Documentation can be updated without requiring a separate reviewer.

#### Infrastructure changes
[https://alsmola.github.io/graphgrc/soc2/cc81.html](https://alsmola.github.io/graphgrc/soc2/cc81.html)

Employees should notify others prior to making changes to Tailscale’s infrastructure, e.g., over Slack. Where infrastructure is codified and uses a deployment tool, infrastructure changes should be approved by another employee prior to being deployed.

#### Customer accounts
[https://alsmola.github.io/graphgrc/soc2/cc81.html](https://alsmola.github.io/graphgrc/soc2/cc81.html)

Tailscale may make changes to customers’ networks and accounts in Tailscale at their request. Changes are initiated by customer support tickets.

Tailscale may also make changes to customer environments without the customer initiating the request, such as when required by law or due to an urgent security issue.

### Security policies
[https://alsmola.github.io/graphgrc/soc2/cc81.html](https://alsmola.github.io/graphgrc/soc2/cc81.html)

Security policies must have a change log to allow auditing of past changes, including when and by whom these changes were made. Tailscale stores these security policies in GitHub and uses git to track changes.

Tailscale will review and evaluate its security policies, adapt them as needed due to changing risks, and validate if the implemented information security continuity controls are sufficient on a quarterly basis.

