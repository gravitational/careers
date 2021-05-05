# Summary

Teleport Cloud provides secure access to servers, databases, and applications of our customers. The platform team is responsible for automating, scaling, monitoring, and securing the infrastructure that underpins Teleport Cloud. This team draws on a wide variety of skills and experience to build reliable infra, and as such we use different interviews based on the candidate's background.


# Rationale

These exercises have two goals.

- It helps us understand what to expect from you as an engineer, how you reason about production systems and investigate faults.
- It helps you get a feel for what it would be like to work at Teleport, as this exercise aims to simulate some of our day-as-usual and expose you to the type of work we're doing here.

We believe this technique is better and is more fun compared to whiteboard/quiz interviews so common in the industry. It's not without the downsides as it could take longer than traditional interviews. [Some of the best teams use coding challenges.](https://sockpuppet.org/blog/2015/03/06/the-hiring-post/)

We appreciate your time and are looking forward to hacking on this challenge together.


# Backgrounds

## Development Focus: Systems / Backend Developers
Systems developers on the platform team leverage a strong background in software development to build, maintain, deploy core production services.

### Challenge
1. [Systems developer challenge](../systems/worker.pdf)

### Levels
See [levels.pdf](../../levels.pdf)


## Operations Focus: Devops / DevSecOps / SREs / Production engineers / etc
Team members with experience in various operations focused disciplines are focussed on the day to day delivery of platform components, such as build/testing automation of infrastructure as code, deploying and configuring our metrics, alarming and logging systems, automating build pipeliness, and other day to day activities to enable our cloud infrastructure.


### Challenge
We use two interviews for team members with operational backgrounds, the first is a lightweight scripting / programming challenge, and the second is a troubleshooting production test. See the following links for outlines of each challenge.
1. [Devops development challenge](devops.md)
2. [Troubleshooting challenge](troubleshooting.md)

### Levels



| Focus Area | Level 1 | Level 2 | Level 3 | Level 4 |
|---|---|---|---|---|
| Communications & Writing | Updates internal documentation as required | Write new documentation, run-books, guides. Able to lead knowledge sharing sessions | Contribute to engineering blog, meetups, booth duty, etc. to advertise product and engineering culture at Teleport | |
| Security | Follows security guidelines and best practices in a production environment such as proper secret handling | Able to analyze and secure new tools, software deployments, or configuration as per best practices | Take charge of security incident coordination, create and execute improvements based on findings | Lead security design work for meeting compliance, engineering, or team goals |
| Oncall | Participates in on-call rotation with backup available when needed | Investigates and solves problems in the production environment | Provide leadership, guidance, and act as escalation point for junior team members | Helps set and guide the production culture for the team |
| Operations | Documents and executes maintenance activities in production such as testing and patching upstream applications | Develops and tests procedures for production in a key area (backup/restore as an example) | | Introduces new operational concepts such as chaos engineering. |
| Networking | Able to capture traffic, check routing tables, test network connectivity, etc. | | Deep understanding of linux and container networking internals and configuration, AWS routing, TCP/IP tuning and optimization, internet routing, load balancing, etc. | Able to design solutions that cross team boundaries that make teleport cloud a best in class service | 
| Systems Engineering | | Design, write, and test automation for different environments, such as chat bots, upgrade scripts, deployments, etc. | Engineers key area of the cloud platform, such as design and deployment of autoscaling logic, patch automation, backup & restore, etc. | |
| OS Understanding | Working understanding of core OS concepts, filesystem, networking, cgroups, namespaces, security mechanisms, etc. | | Solid understanding of linux workload isolation, syscall filtering, internals, for high security and production workloads | |
| Tooling | Able to patch / update internal tools, test and apply configuration updates, etc. | Take ownership of key tools, like CI/CD platform, Terraform, AWS Configuration, etc. | Take design lead on key operational areas, such as introducing new tooling, new monitoring, or complex automation | Able to work with vendors or opensource communities to ensure external tooling meets our security, compliance, and operational needs |

<br />
Note: These levels are used as a guide, not every engineer specifically matches every skill in every focus area of the guide. Higher level engineers are expected to have a broader set of skills than junior team members. Engineers that focus on certain disciplines do not require as broad an area of skills, such as team members focussing on Tooling may not require as deep network understand. All team members are expected to meet Communications, Security, Oncall, and Operations requirements.


