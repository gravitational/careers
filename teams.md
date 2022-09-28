# Summary

This page is meant to be a general guide to learn about the different teams at
Teleport to help candidates decide which may be the best fit for them.

## Core

The Teleport Core team is focused on design and development of the core
product: Teleport. It is split into multiple sub-teams that work on specific
areas of the product.

### Technologies

We use the following technologies to build Teleport.

* Go
* Rust
* GitHub Actions
* Docker

### Performance and Scaling

The Performance and Scaling team is focused on scaling, performance, and
robustness of Teleport in large deployments.

Recent projects:

* Improve `tsh` initial connection latency
  [#16257](https://github.com/gravitational/teleport/pull/16257)
  [#16456](https://github.com/gravitational/teleport/pull/16456)
* Improve performance of Teleport on large clusters
  [#16636](https://github.com/gravitational/teleport/pull/16636)

### Platform Security

The Platform Security team is focused on application security for the Teleport
product. This team implements security controls and works with security
researchers to find and fix security vulnerabilities.

Teleport Engineering also has a general [Security
team](https://github.com/gravitational/careers/blob/main/teams.md#security)
that focuses on holistic security across the organization.

Recent projects:

* Add support for passwordless authentication
  [#9296](https://github.com/gravitational/teleport/pull/9296)
* Add support for searching for resource
  [#11261](https://github.com/gravitational/teleport/pull/11261)

### Server Access and Kubernetes Access

The Server and Kubernetes Access teams work on secure access using short-lived
credentials to SSH servers and Kubernetes clusters.

Recent projects:

* Add support for non-disk based certificate support to `tsh`
  [#14774](https://github.com/gravitational/teleport/pull/14774)
  [#15334](https://github.com/gravitational/teleport/pull/15334)
  [#15335](https://github.com/gravitational/teleport/pull/15335)
  [#15336](https://github.com/gravitational/teleport/pull/15336)
* Switch from scp to sftp
  [#13216](https://github.com/gravitational/teleport/pull/13216)
  [#13491](https://github.com/gravitational/teleport/pull/13491)
  [#16601](https://github.com/gravitational/teleport/pull/16601)
* Automatic discovery and enrollment of servers on AWS
  [#12410](https://github.com/gravitational/teleport/pull/12410)

### Application Access

The Application Access team works on providing access to web and console
applications behind NATs and firewalls with security and compliance
needs.

Recent projects:

* Add certificate-based authentication support to CLI applications
  [#5918](https://github.com/gravitational/teleport/pull/5918)

### Database Access

The Database Access team works on providing access to databases behind NATs and
firewalls with security and compliance needs.

Recent projects:

* Added support for mTLS over an AWS ALB
  [#14369](https://github.com/gravitational/teleport/issues/14369)
  [#15669](https://github.com/gravitational/teleport/pull/15669)
* Added support for Redis
  [#10053](https://github.com/gravitational/teleport/pull/10053)
* Added support for Snowflake
  [#12006](https://github.com/gravitational/teleport/pull/12006)
* Added support for MongoDB
  [#7213](https://github.com/gravitational/teleport/pull/7213)

### Desktop Access

The Teleport Desktop Access team is focused on secure and user-friendly remote
desktop access.

Recent projects:

* Add support for non-AD hosts
* Added support for file system sharing
  [#16022](https://github.com/gravitational/teleport/pull/16022)

### Machine ID

The Machine ID team is focused on bringing all the advantages and convenience
that Teleport provides for human users to machine use cases.

Recent projects:

* Added Kubernetes Access support
  [#14269](https://github.com/gravitational/teleport/pull/14269)
* Added Application Access support
  [#14551](https://github.com/gravitational/teleport/pull/14551)
* Added Database Access support
  [#11596](https://github.com/gravitational/teleport/pull/11596)
  [#11904](https://github.com/gravitational/teleport/pull/11904)
* Added host certificate support
  [#15592](https://github.com/gravitational/teleport/pull/15593)

### Access Provider

The Access Provider team is focused on web and desktop applications that make
Teleport easier to use and more secure.

Recent projects:

* Added support for passwordless
  [#724](https://github.com/gravitational/webapps/pull/724)
* Added support for server-side paginating and filtering
  [#9633](https://github.com/gravitational/teleport/pull/9633)

### Access Manager

The Access Manager team is focused on simplifying connecting and configuring
Teleport.

Recent projects:

* Teleport Discover
  [#13440](https://github.com/gravitational/teleport/pull/13440)

### Integrations

The Integration team is focused on building strong integrations with tools like
Kubernetes, Terraform, Ansible, and more.

Recent projects:

* Added support for Microsoft Teams to Access Requests
* Added Teleport Kubernetes Operator
  [#13331](https://github.com/gravitational/teleport/pull/13331)

### Tools

The Core Tooling team contributes force multiplication efforts to help
engineers developing Teleport and across the org. This team is the backbone to
ensuring the rest of the development team remains incredibly productive, and
that we operate in the open, with an open source code base.

This team is responsible for the tooling necessary to build and release
Teleport artifacts, including container images, AMIs, Helm Charts, and
integration/distribution to package managers.

Recent projects:

* Create multi-arch container images
  [#16688](https://github.com/gravitational/teleport/pull/16688)
* Build package distribution infrastructure for apt/deb
  [#10746](https://github.com/gravitational/teleport/pull/10746)
* Secure build infrastructure

## Cloud

The Teleport Cloud team develops and operates Teleport Cloud for customers as
service. Cloud exists to help alleviate the burden of running and maintaining
secure global low-latency access to infrastructure.

### Technologies

We use the following technologies to build Teleport Cloud:

* Go
* Typescript & React
* PostgreSQL
* Terraform and Packer
* Amazon Web Services (AWS)
* Kubernetes
* Prometheus/Alertmanager/Loki/Grafana

### Product User Experience

The team is focused on delivering a great user experience for Cloud users.

Recent projects:

* Add support for Upgrade Windows
* Simplify the Teleport Downloads page
* User journey tracking

### Product Infrastructure

The team is focused on scalability, security, and analytics of Teleport on our
Cloud platform.

Recent projects:

* Reduce tenant on-boarding time
* Build reporting and analytics platform

### Platform Metrics & CI/CD

The team is focused on release automation, observability, and developer
experience.

Recent projects:

* CI/CD with GitHub Actions
* Develop more sophisticated resource growth alerts
* Automate manual processes
* Improve platform monitoring and alerting

### Platform Infrastructure

The team is focused on building secure, reliable, and low latency
infrastructure for the Cloud platform.

Recent projects:

* Reduce connection latency for geo-distributed infrastructure
* Build next-generation Cloud platform infrastructure
* Update disaster recovery and backups infrastructure
* Multi-region support for Kubernetes operator

## Security

At Teleport, each and every engineer is responsible for security of their work.
In addition to this individual mandate and our [Product Security
team](#product-security), we maintain a Security team focused on
organization-wide efforts. We're currently working on the following areas:

* Software supply chain security. We ensure infrastructure and code is
  protected and auditable from developer to production.
* Teleport's bug bounty program.
* Working with consultants and independent experts to perform blackbox,
  whitebox, and red team validation of our code and security controls.
* Updating compliance documentation, internal controls, and our corporate
  policies.

### What is the difference between the "Product Security" and "Security" teams?

The Product Security team works primarily in the Teleport codebase, developing
new security features and fixing bugs for the next Teleport release. The
Security team addresses all elements of information security, including cloud
security, app security, IT security, GRC, policies, training, and our bug
bounty.

To illustrate, you'd find a Product Security team member hacking on issues like
[#10375](https://github.com/gravitational/teleport/issues/10375). You might
find a Security team member improving our internal Okta terraform or improving
the policies and infrastructure backing our promises at
https://goteleport.com/security/.
