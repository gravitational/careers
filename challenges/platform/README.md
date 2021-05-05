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


| Level  | Oncall  | Operations  | Security  | Communications / Writing  | Networking  | Systems Engineering  | OS Understanding  | Tooling |
|---|---|---|---|---|---|---|---|---|
| 1 | Participates in on-call and able to execute run books | Documents and Executes maintenance activities in production | Able to follow security guidelines and best practices in a production environment  | Contributor to internal documentation such as updating run books | Able to troubleshoot introductory networking issues, network policies, routing, etc. | Focus on and contribute to specific area of production system (such as the deployment of our monitoring stack)  | Working understanding of core OS concepts, filesystem, networking, cgroups, namespaces, security mechanisms, etc.  | Work with team members on deployment and maintenance of internal tooling |
| 2 | Can investigate and solve a majority of production incidents and issues with peer review  | Develops and tests maintenance of production environment (backup/restore as an example)  | Able to analyze and secure new tools or systems connectivity  | Write new documentation, run-books, guides. Able to lead knowledge sharing sessions  |   | Design, write, and test automation for different environments, such as chat bots, upgrade scripts, etc. |   | Able to take ownership and key contributor to key tooling.|
| 3 | Provide leadership, guidance, and act as escalation point for junior team members |   | Take charge of security incident coordination, create and execute improvements based on new findings | Contribute to engineering blog, meetups, booth duty, etc to advertise product and engineering culture at Teleport  | Solid understanding of advanced networking, routing, load balancing concepts and scaling production systems  | Design and implement significant automation for challenging problem spaces, such as implementing horizontal/vertical autoscaling and resource allocation | Solid understanding of workload isolation, syscall filtering, linux internals, for high security and production workloads  | Take design lead on key operational areas, such as introducing new tooling, new monitoring, or complex automation |
| 4 | Helps set and guide the production culture for the team | Finds the areas we are weak and takes ownership on solutions  | Lead security design work for meeting compliance, engineering, or team goals  |   | Able to design solutions that cross team boundaries that make teleport cloud a best in class service  |   |   | Able to work with vendors or opensource communities to ensure external tooling meets our security, compliance, and operational needs |

