# Summary

This is meant to be a general guide to learn about the different kinds of
projects our engineering teams are working on at Teleport and to help decide
which may be the best fit for those looking to join.

## Core

The Teleport Core team is focused on design and development of the core product:
Teleport. It is split into multiple sub-teams that work on specific areas of the
product.

### Server and Kubernetes Access

The Server and Kubernetes Access team works on
[SSH](https://goteleport.com/ssh-server-access/) and
[Kubernetes](https://goteleport.com/kubernetes-access/) servers and proxies. For
Server Access, the team is working on automatic user provisioning and
configuration, as well as IP-based restrictions. For Kubernetes access, the team
is building out features to support Teleport deployments that don't require
static tokens.

### Application and Database Access

The Application and Database Access team works on providing access to [web applications](https://goteleport.com/docs/application-access) and [databases](https://goteleport.com/docs/database-access/) behind NATs and firewalls with security and compliance needs. The team has recently worked on adding [certificate based authentication support to CLI applications](https://github.com/gravitational/teleport/pull/5918) and adding support for the [MongoDB protocol](https://github.com/gravitational/teleport/pull/7213) to Database Access.

#### Here are the examples of projects you may be working on:

- Adding support for more database access [protocols](https://github.com/gravitational/teleport/issues?q=is%3Aopen+is%3Aissue+label%3Adatabase-access+label%3Adb%2Frequested) and improving the existing ones.
- Working on [features](https://github.com/gravitational/teleport/issues?q=is%3Aissue+is%3Aopen+label%3Aapplication-access+label%3Afeature-request) for application access identity-aware proxy.
- Building advanced access controls such as [session recording](https://github.com/gravitational/teleport/issues/5799), [data masking](https://github.com/gravitational/teleport/issues/7150) and [per-session MFA](https://github.com/gravitational/teleport/issues/6172).
- Implementing cloud access solutions for AWS, GCP and Azure.
- Improving the UX of application and database access products.

### Product Security

The Product Security team is focused on application security for the Teleport
product. This team implements security controls and works with security
researchers to find and fix security vulnerabilities.

The team is focused on releasing a preview of passwordless support, which will
allow users to log in to Teleport using biometric proof of identity in
conjunction with a proof of presence.

Teleport Engineering also has a general [Security team](#security) that focuses
on holistic security across the organization.

### Performance and Scaling

The Performance and Scaling team is focused on scaling, performance, and
robustness of Teleport in large deployments.

The team is currently working on scaling Teleport to
[100k node clusters](https://github.com/gravitational/teleport/issues/4173) and
[secure automatic upgrades](https://github.com/gravitational/teleport/pull/6691)
of large fleets of servers.

### Machine ID

The [Machine ID](https://goteleport.com/blog/machine-to-machine-access/) team is
focused on bringing all the advantages and convenience that Teleport provides
for human users to machine use cases.

The team is currently working on:

- support for database access, kubernetes access, and application access
- simplifying setup and configuration
- support for issuing host certificates
- CA rotation

### Desktop Access

The Teleport Desktop Access team is focused on secure and user-friendly remote desktop access.

The team is currently working on browser-based Windows Desktop Access, including:

- MFA
- Shared clipboard
- Session recording
- Performance optimization and security hardening

### Teleport Connect

The [Teleport Connect](https://goteleport.com/blog/teleport-connect/) team is
focused on creating desktop applications that make Teleport easier to use and
more secure.

The team is working towards the first release of Teleport Connect, adding
support for SSO, Server Access, Database Access, and Kubernetes Access.

### Frontend Applications

The Teleport Frontend Applications team is focused on building user-friendly web
applications.

- An updated web UI for passwordless logins
- Advanced filtering and clickable labels
- Improved UI for access requests

### Internal Tools

The Core Tooling team contributes force multiplication efforts to help engineers
developing Teleport and across the org. This team is the backbone to ensuring
the rest of the development team remains incredibly productive, and that we
operate in the open, with an open source code base.

This team is responsible for the tooling necessary to build and release Teleport
artifacts, including container images, AMIs, Helm Charts, and
integration/distribution to package managers.

The team is currently working on:

- Redefining our CI tooling, to ensure high confidence that the CI tooling is
  not a security liability, while making the development teams more productive.
- Building robust and essential Release engineering tooling, so the code we
  produce can get into our customers hands in a secure manner.
- Internal automation, bots, and integrations that keep the team members
  autonomous.

### Teleport Devops

The Teleport codebase is used by DevOps teams around the world. These customers
and OSS users use a variety of tools in their day to day jobs. This role is all
about ensuring that we provide stable and reliable integrations with the
configuration management tools that are most important to our community. We want
to ensure that whether it's Kubernetes or Ansible, we have a fully realized way
to deploy and maintain Teleport. This is a great opportunity to work with
Kubernetes, Terraform (on various cloud providers), Ansible, Salt, Puppet, and
so much more!

## Cloud

The Teleport Cloud offering is new; we're on a journey of defining and building
a culture around production services, offering a secure and reliable hosted
version of Teleport. Our mandate is straightforward; we need to be prepared to
secure, develop, monitor, maintain, scale, investigate, and automate the cloud
hosted version of Teleport and the systems that support our work. Our work
varies from the mundane to the remarkable, ensuring that the cloud version of
Teleport is a compelling option for our customers.

### Technologies

We use the following technologies to build Teleport Cloud:

- Golang
- Typescript & React
- PostgreSQL
- Terraform and Packer
- Amazon Web Services
- Kubernetes
- Prometheus / Alertmanager / Loki internally and on Grafana Cloud
- Drone.io

### Cloud Reliability

Our Cloud Reliability Team defines and builds a culture around production,
ensuring scaling of our product offering against company/product growth.

Some projects we're working on right now:

- Expanding Teleport Cloud to multiple regions, reducing latency between users
  and their infrastructure.
- Defining SLOs, and developing centralized monitoring and alerting based on them
- Evaluating and implementing database and storage solutions that match or
  exceed our SLOs and growth demands
- Developing procedures and automation to ensure reliable and safe rollouts

### Cloud Applications

As the applications team for Teleport Cloud, we manage the Teleport deployments
for our cloud customers and provide a stable, easy-to-use product.

The team intersects with almost all other departments in the company and offers
a broad range of projects. For example:

- Building Kubernetes operators
- Developing subscription & billing logic
- Contributing UI & UX improvements for Teleport
- Rewriting Teleport core networking to work better for cloud use cases.
- Asset distribution

### Cloud Tooling

The Cloud Tooling team contributes force multiplication efforts to help
engineers developing Teleport Cloud and across the org. We're currently working
on:

- Building and maintaining Teleport clusters that engineers use to securely
  access our infrastructure backing Teleport Cloud.
- Improving static analysis and automated tests.
- Developing security conscious CI guidelines and tooling.
- While we build up the Cloud Security team, tooling is heavily invested in
  improving our security posture
  audit trail for generated and published assets and ensuring we're using
  encryption at rest in the right places.

## Security

At Teleport, each and every engineer is responsible for security of their work.
In addition to this individual mandate and our
[Product Security team](#product-security), we maintain a Security team
focused on organization-wide efforts. We're currently working on the following areas:

- Software supply chain security. We ensure infrastructure and code is protected
  and auditable from developer to production.
- Teleport's bug bounty program.
- Working with consultants and independent experts to perform blackbox, whitebox,
  and red team validation of our code and security controls.
- Updating compliance documentation, internal controls, and our corporate policies.

### What is the difference between the "Product Security" and "Security" teams?
The Product Security team works primarily in the Teleport codebase, developing
new security features and fixing bugs for the next Teleport release. The
Security team addresses all elements of information security, including cloud
security, app security, IT security, GRC, policies, training, and our bug bounty.

To illustrate, you'd find a Product Security team member hacking on issues like
[#10375](https://github.com/gravitational/teleport/issues/10375). You might find
a Security team member improving our internal Okta terraform or improving the
policies and infrastructure backing our promises at https://goteleport.com/security/.
