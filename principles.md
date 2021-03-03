## Product and Engineering Principles

**Simpler always wins in the end**

We choose simpler designs whenever possible. Having less configuration knobs
is better for security. Complex products are easier to hack and harder
to use. Security feature is only good, when it's easy to use.

That's why we stick to [client certificates](https://goteleport.com/teleport/how-it-works/certificate-based-authentication-ssh-kubernetes/)
and as the only way to authenticate and [reduce the number of ports and protocols](https://github.com/gravitational/teleport/issues/5777).

**Security through transparency**

The only way to build secure systems is through peer review and transparency.

That's why we have [open design process](https://github.com/gravitational/teleport/pull/5135/files),
and [publish security audits](https://blog.doyensec.com/2020/03/02/gravitational-audit.html) even
if the reports are not always shiny.

**Security and customers are not the product**

Everyone deserves strong security and privacy.

That is why all security features in Teleport are free and our software does not spy on our customers.
When we want your data, we ask [for permission first](https://github.com/gravitational/teleport/pull/550).
