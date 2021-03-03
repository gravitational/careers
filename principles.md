## Product and Engineering Principles

**We choose simpler designs whenever possible.**

Having less configuration knobs is better for security. Complex products are easier to hack and harder
to use. Security feature is only good if it make user experience better, otherwise users will find ways to work around it.

That's why we prefer [client certificates](https://goteleport.com/teleport/how-it-works/certificate-based-authentication-ssh-kubernetes/)
and as way to authenticate and [reduce the number of ports, configuration options and protocols](https://github.com/gravitational/teleport/issues/5777).

**Security through transparency**

The only way to build secure systems is through peer review and transparency.

That's why we have [open design process](https://github.com/gravitational/teleport/pull/5135/files),
and [publish security audits](https://blog.doyensec.com/2020/03/02/gravitational-audit.html) even
if reports are critical.

**We are not selling security and customers are not the product**

Our products improve developer efficiency and don't stay in the way. They have to be secure, but security
is not what we sell. Everyone deserves strong security and privacy.
That's why we built OSS version of Teleport to be as secure as possible. Our software does not spy on our customers.
When we want your data, we ask [for permission first](https://github.com/gravitational/teleport/pull/5505).
