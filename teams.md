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

We want to ensure Teleport is easy to install and maintain within our customers infrastructure, whether they're enterprise or opensource customers. This means supporting the plethora of configuration management tools and approaches different customers use. In this role you get to work with the most popular configuration management tooling to ensure we have modules and docs for using Teleport with the customers preferred deployment platform. This role allows a great opportunity to work with Kubernetes, Terraform, Packer, Ansible, Salt, Puppet, and more within various cloud providers.

Success in this role ensures that customers of Teleport can not only install but maintain their teleport deployments on their own infrastructure.

## Cloud
The Teleport Cloud offering is new; we're on a journey of defining and building a culture around production, offering a secure and reliable hosted version of Teleport Enterprise as a service. Our mandate is straightforward; we need to be prepared to secure, develop, monitor, maintain, scale, investigate, and automate the SAAS version of Teleport. Our work includes a wide variety of projects drawing on various skill sets, from ensuring we're up to date on patches to redesigning parts of the Teleport Core project to leverage the capabilities a cloud environment can deliver.


### Cloud Tools Team
TBD

### Cloud Reliability Team

The Cloud Reliability team combines practices from SRE, Production Engineering, and Devops to take ownership of our production environment, focussing on the security, monitoring, scalability, maintenance, and engineering of production. This team is relatively new, we're building practices appropriate to a new SAAS offering, and taking on appropriate technical debt for a team our size, as we expand and grow the Teleport Cloud product.

Here are some of the projects we're working on right now:

- We're scaling our cloud capabilities, moving from decisions made during prototyping that limited scale, to designs that can sustain thousands of tenants.
- We're expanding Teleport Cloud to multiple regions around the world with automatic routing the optimal path for users to reach their infrastructure.
- We're rewriting our monitoring stack which started quick and dirty, into a platform we can use to escalate the real emergencies, and leave our engineers with a good nights sleep for everything else.
- Disaster recovery planning and tests that prepare us for mistakes that may occur in production


#### The toolset we currently use is:

* Golang
* Terraform and Packer
* Kubernetes
* Prometheus / Alertmanager / Loki internally and on Grafana Cloud
* Amazon Web Services
* Drone.io
* and More

### Cloud Application Team

The Cloud Applications team focusses on the software heavy lifting of the cloud product, writing our controllers, networking improvements, and management plane of Teleport Cloud to offer a stable and easy-to-use product.

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

### Cloud Security Infrastructure Team

At Teleport, we carry a philosophy that all teams are responsible for security, and it isn't someone else's job to make our product secure. With that in mind, we want to put together a team that can fill in our gaps, and focus on driving our security capabilities forward. Members of the security infrastructure team will build common infra and tooling that enhances our security posture, such as secure access to our production environment. They'll also consult with other teams and coordinate cross-team initiatives to ensure we're continually making Teleport Cloud a more difficult product to breach. 

This is a new team, that can draw on development, sre, and devops skills with an interest in developing expertise in technical security.

Projects within the team:
- Secure admin access to the production environment - ensuring our attack chain requires multiple vulnerabilities to compromise customers secrets.
- Executing on our security roadmap, from ensuring we have proper scanning and response to documented vulnerabilities, to tackling complex software supply chain problems.
- Building common audit infrastructure, that other teams can simply plug into



