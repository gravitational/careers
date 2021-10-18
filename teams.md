# Summary

This is meant to be a general guide to learn about the different kinds of projects our engineering teams are working on at Teleport and to help decide which may be the best fit for those looking to join.

## Core

The Teleport Core team is focused on design and development of the core product: Teleport. It is split into multiple sub-teams that work on specific areas of the product.

### Server and Kubernetes Access

The Server and Kubernetes Access team works on [SSH](https://goteleport.com/ssh-server-access/) and [Kubernetes](https://goteleport.com/kubernetes-access/) servers and proxies. The team recently added support for [eBPF based restricted sessions](https://goteleport.com/docs/server-access/guides/restricted-session) and is working on extending eBPF based monitoring and access controls to Kubernetes clusters.

### Application and Database Access

The Application and Database Access team works on providing access to [web applications](https://goteleport.com/docs/application-access) and [databases](https://goteleport.com/docs/database-access/) behind NATs and firewalls with security and compliance needs. The team has recently worked on adding [certificate based authentication support to CLI applications](https://github.com/gravitational/teleport/pull/5918) and adding support for the [MongoDB protocol](https://github.com/gravitational/teleport/pull/7213) to Database Access.

#### Here are the examples of projects you may be working on:

* Adding support for more database access [protocols](https://github.com/gravitational/teleport/issues?q=is%3Aopen+is%3Aissue+label%3Adatabase-access+label%3Adb%2Frequested) and improving the existing ones.
* Working on [features](https://github.com/gravitational/teleport/issues?q=is%3Aissue+is%3Aopen+label%3Aapplication-access+label%3Afeature-request) for application access identity-aware proxy.
* Building advanced access controls such as [session recording](https://github.com/gravitational/teleport/issues/5799), [data masking](https://github.com/gravitational/teleport/issues/7150) and [per-session MFA](https://github.com/gravitational/teleport/issues/6172).
* Implementing cloud access solutions for AWS, GCP and Azure.
* Improving the UX of application and database access products.

### Product Security

The Product Security team is focused on application security for the entire product. This team implements security controls and works with security researchers to find and fix security vulnerabilities.

The team has recently added support for [Per-Session MFA](https://goteleport.com/docs/access-controls/guides/per-session-mfa/), [Access Requests](https://goteleport.com/docs/enterprise/workflow/), [user and session locking](https://github.com/gravitational/teleport/pull/7286), and mitigated issued found during [security audits](https://goteleport.com/resources/audits/).

### Performance and Scaling

The Performance and Scaling team is focused on scaling, performance, and robustness of Teleport in large deployments.

The team is currently working on scaling Teleport to [100k node clusters](https://github.com/gravitational/teleport/issues/4173) and [secure automatic upgrades](https://github.com/gravitational/teleport/pull/6691) of large fleets of servers.

### Teleport Terminal

The Teleport Terminal team is focused on creating desktop applications that make Teleport easier to use and more secure.

The team is working on pairing [hardware devices](https://github.com/gravitational/teleport/pull/7808) with certificate authentication to give users a seamless highly secure environment.

### Desktop Access

The Teleport Desktop Access team is focused on secure and user-friendly remote desktop access.

The team is currently working on browser-based [Windows Desktop Access](https://github.com/gravitational/teleport/pull/7725).

### Frontend Applications

The Teleport Frontend Applications team is focused on building user-friendly web applications.

* Adding session event player to the audit log.
* Building a new user interface for Teleport's terminal.
* Improving Node Search. [#2699](https://github.com/gravitational/teleport/issues/2699)
* Ability to comment / link to specific times in recorded sessions. [#996](https://github.com/gravitational/teleport/issues/996)
* Adding better support for WebAuthn (second factor).

### Internal Tools

The Internal Tools team is focused on building simple and secure foundational tools and processes to increase developer autonomy and productivity in a distributed environment.

* Workflow automation bots [#8116](https://github.com/gravitational/teleport/pull/8116)
* Designing and building next generation CI and CD platforms.

### Teleport Devops

Teleport DevOps is responsible for building, maintaining, and testing open-source modules for various configuration management platforms. This will be the open-source code that enables both our OSS and Enterprise customers to deploy and maintain Teleport on their chosen configuration management system. We’re looking to cover Kubernetes, Terraform on various clouds, Ansible, Salt, Puppet, etc.

## Cloud Platform team

The Teleport Cloud offering is new; we're on a journey of defining and building a culture around production, offering a secure and reliable hosted version of Teleport Enterprise as a service.

Our mandate is straightforward; we need to be prepared to secure, monitor, maintain, scale, investigate, and automate our production environment for Teleport Cloud. Our work includes a wide variety of projects drawing on various skill sets, from ensuring we're up to date on patches to redesigning parts of the Teleport Core project to be more appealing to the Cloud.

#### Here are some of the projects we're working on right now:

* Rewriting the Teleport core project networking so that agents run by customers require fewer connections and overhead when phoning home.
* We are expanding Teleport Cloud to multiple regions, reducing latency between users and their infrastructure.
* We are deploying auto scalers, so no one needs to be paged to scale our infrastructure when there is a burst in new accounts.
* We're improving our security posture, currently focusing on ensuring we have an audit trail for generated and published assets and ensuring we're using encryption at rest in the right places.
* Progressing beyond minimum viable monitoring and alerting and fully adopting Grafana Cloud as our centralized metrics, alerting and logging platform.

#### The toolset we currently use is:

* Golang
* Terraform and Packer
* Kubernetes
* Prometheus / Alertmanager / Loki internally and on Grafana Cloud
* Amazon Web Services
* Drone.io
* and More

## Cloud Application team

Teleport Cloud is the SaaS version of Teleport that we launched in early 2021. As the applications team for Teleport Cloud, our goal is to manage the deployments of Teleport for our cloud customers and provide a stable, easy-to-use product.

The team intersects with almost all other departments in the company and as such there is a large variety of projects that we're involved in, for example:

* Custom Kubernetes operators
* Subscription & billing logic
* Internal dashboards & admin panels
* UI & UX improvements for Teleport
* Analytics & Monitoring
* Asset distribution

#### We use the following stack:

* Go
* Typescript
* React
* Kubernetes
* AWS
* Terraform
* PostgreSQL
* gRPC
* Grafana & Prometheus
