---
title: Password policy
slug: password
policy: true
faq: false
weight: 8
---

To avoid potential security incidents, Tailscale requires employees to follow password requirements.

### Scope

This policy applies to passwords for any application or server accessed by Tailscale employees, contractors, or vendors. _It does not apply to the passwords customers of Tailscale use to access the Tailscale service._

### Password strength
[https://alsmola.github.io/graphgrc/soc2/cc62.html](https://alsmola.github.io/graphgrc/soc2/cc62.html)
[https://alsmola.github.io/graphgrc/soc2/cc11.html](https://alsmola.github.io/graphgrc/soc2/cc11.html)
[https://alsmola.github.io/graphgrc/soc2/cc51.html](https://alsmola.github.io/graphgrc/soc2/cc51.html)
[https://alsmola.github.io/graphgrc/soc2/cc58.html](https://alsmola.github.io/graphgrc/soc2/cc58.html)

Passwords must be unique for each use.

Passwords must be randomly generated.

Default passwords on all systems are changed after installation. Initial passwords generated for new users must be changed after login.

Passwords do not need to be regularly rotated. However, if a password is known or thought to be compromised, it must be rotated to a new password.

### Single sign-on
[https://alsmola.github.io/graphgrc/soc2/a13.html](https://alsmola.github.io/graphgrc/soc2/a13.html)
[https://alsmola.github.io/graphgrc/soc2/a14.html](https://alsmola.github.io/graphgrc/soc2/a14.html)

Where a third-party application supports single sign-on, it must be used.

### Multi-factor authentication
[https://alsmola.github.io/graphgrc/soc2/cc61.html](https://alsmola.github.io/graphgrc/soc2/cc61.html)

Where a third-party application supports multi-factor authentication, it must be used. Use of multi-factor is enforced where possible.

Acceptable forms of multi-factor authentication include authentication apps or a WebAuthn token. Embedded tokens (e.g., TouchID) are permitted. WebAuthn hardware or embedded hardware tokens are preferred to authentication apps.

### Password manager
[https://alsmola.github.io/graphgrc/soc2/cc61.html](https://alsmola.github.io/graphgrc/soc2/cc61.html)
[https://alsmola.github.io/graphgrc/soc2/cc63.html](https://alsmola.github.io/graphgrc/soc2/cc63.html)
[https://alsmola.github.io/graphgrc/soc2/cc64.html](https://alsmola.github.io/graphgrc/soc2/cc64.html)

Where SSO is not used, and where possible, passwords should be stored in a password manager.

### Encryption at rest
[https://alsmola.github.io/graphgrc/soc2/a12.html](https://alsmola.github.io/graphgrc/soc2/a12.html)

Passwords should be stored encrypted at rest.

### Logging
[https://alsmola.github.io/graphgrc/soc2/cc67.html](https://alsmola.github.io/graphgrc/soc2/cc67.html)

Passwords should not be logged.

### Requirements for specific use cases
[https://alsmola.github.io/graphgrc/soc2/pi14.html](https://alsmola.github.io/graphgrc/soc2/pi14.html)
[https://alsmola.github.io/graphgrc/soc2/pi15.html](https://alsmola.github.io/graphgrc/soc2/pi15.html)

#### Servers

Access to servers, for both production as well as development and testing infrastructure, must be with a password and MFA or with per-user public keys (e.g., SSH keys). Only Tailscale-based network authentication is permitted for services not exposed to the Internet.

#### Automated processes
[https://alsmola.github.io/graphgrc/soc2/cc61.html](https://alsmola.github.io/graphgrc/soc2/cc61.html)
[https://alsmola.github.io/graphgrc/soc2/cc62.html](https://alsmola.github.io/graphgrc/soc2/cc62.html)
[https://alsmola.github.io/graphgrc/soc2/cc63.html](https://alsmola.github.io/graphgrc/soc2/cc63.html)
[https://alsmola.github.io/graphgrc/soc2/cc64.html](https://alsmola.github.io/graphgrc/soc2/cc64.html)

Automated processes, including deployment or CI/CD tools, should use passwords or API keys to access and communicate with other systems. Passwords used in scripts must be encrypted at rest.

#### End user devices
[https://alsmola.github.io/graphgrc/soc2/cc61.html](https://alsmola.github.io/graphgrc/soc2/cc61.html)
[https://alsmola.github.io/graphgrc/soc2/cc63.html](https://alsmola.github.io/graphgrc/soc2/cc63.html)
[https://alsmola.github.io/graphgrc/soc2/cc66.html](https://alsmola.github.io/graphgrc/soc2/cc66.html)
[https://alsmola.github.io/graphgrc/soc2/cc67.html](https://alsmola.github.io/graphgrc/soc2/cc67.html)
[https://alsmola.github.io/graphgrc/soc2/cc68.html](https://alsmola.github.io/graphgrc/soc2/cc68.html)

End user devices must use passwords to encrypt their disks and unlock the device. These must be unique for each individual but may be reused across an individual’s devices. These do not need to be randomly generated.

#### SaaS applications or other software
[https://alsmola.github.io/graphgrc/soc2/cc67.html](https://alsmola.github.io/graphgrc/soc2/cc67.html)

Access to third party applications must use SSO where possible, MFA where possible, and enforce MFA where possible.

An individual’s password for their password management vault must be unique. These do not need to be randomly generated.

