# Summary

Monitor, deploy, and extend a prototype service that provides an API for Kubernetes resources

# Rationale

This exercise has two goals:

* It helps us to understand what to expect from you as a Site Reliability Engineer,
  how you reason about a production service, how you reason about system design, your
  opinions on automation & tooling, and how you communicate when trying to solve problems.
* It helps you get a feel for what it would be like to work at Teleport, as this
  exercise aims to simulate our day-as-usual and expose you to the type of work
  we're doing here.

We believe this technique is not only better, but also is more fun compared to
whiteboard/quiz interviews, so common in the industry. It's not without the
downsides - it could take longer than traditional interviews.

[Some of the best teams use coding challenges.](https://sockpuppet.org/blog/2015/03/06/the-hiring-post/)

We appreciate your time and are looking forward to hacking on this project together.

# Requirements

This challenge covers 5 engineering levels at Teleport. Level 6 is only for internal promotions. Check the
[engineering levels document](https://raw.githubusercontent.com/gravitational/careers/main/levels.pdf)
for more details.

* Your submission will need to meet the requirements of the level you are applying for.
* Split the submission into 2-3 pull requests for us to review. We will review
  every pull request and provide our feedback.
* We are going to inspect the repository, test it, and get back to you.

## Server

An incomplete implementation of a Go server is located [here](sre-server.go).
You will extend this implementation by adding features described below depending
on the level you are applying to. Start by creating a new Github repository. Then
let the interview panel know the repository's location by pasting a link in your
interview Slack channel. Invite interview panel participants as contributors to
the new repository if you prefer to keep your submission private.

## Design Doc

Start with a brief design document that covers the edge cases and design approach. At
Teleport, we prefer Markdown for
[our designs](https://github.com/gravitational/teleport/blob/master/rfd/0000-rfds.md).

Be sure to cover the following in your design:

* Developer workflow
* Level 3+: Build, Release
* Level 4+: State management and Delivery
* Level 5+: Reconciliation, Conflicts, and Automation

Open a pull request with your design and share a link with the reviewers via Slack.
After the document is approved by at least 2 reviewers, move on to the implementation.

A few notes about the design document:

* We expect the design document to be completed roughly within the first week.
  This is to ensure you have enough time to work on the implementation.
* Avoid writing an overly detailed design document. 500-1500 words is
  sufficient.
* Avoid sending us draft design documents, it is difficult to evaluate what
  parts are draft and which parts are complete. Instead we encourage asking
  questions in Slack and sharing a design document that is ready to be
  reviewed.

## Server & Automation

Once the design document is approved, begin working on adding features to the
server implementation that includes automation. A key feature of this challenge 
is to produce a self-contained GitHub project that automates away as much as 
possible. This should include code quality, testing, build, and deployment.

Expect the interview panel to clone the repository and execute one or more 
Make targets that build and test a working solution. 

Add a couple of high quality tests that cover happy and unhappy scenarios. Keep
external dependencies as low or clear as possible. Expect the reviewers to
checkout the repository and execute the tests themselves.

Do not try to achieve full test coverage. This will take too long. Write enough
to exercise the different key components and show they are working as
intended, as well as to demonstrate your approach to automation.

## Deployment

Levels 3+ include a deployment step which should deploy the solution to a local
Kubernetes cluster. The choice of which Kubernetes cluster is up to you, but 
please choose one that has the ability to run on macOS and Linux.

## Key Dependencies

You are welcome to solve this challenge using tools you are most familiar with,
but, at a minimum, please include the following:

* Go
* make
* Docker

## Level 1

### Server

* HTTP API to retrieve the replica count of the Kubernetes Deployment

### Automation

* Ability to build the server in a Docker container
* Ability to run the server

## Level 2

### Server

* HTTP API to retrieve the replica count of the Kubernetes Deployment
* HTTP API to set the replica count of the Kubernetes Deployment

### Automation

* Ability to build the server in a Docker container
* Ability to run the server
* Ability to execute integration tests against the local Kubernetes cluster

## Level 3

### Server

* HTTP API to retrieve the replica count of the Kubernetes Deployment
* HTTP API to set the replica count of a the Kubernetes Deployment
* HTTP API to get the list of available Deployments in the Kubernetes cluster
* HTTP health check verifying Kubernetes connectivity

### Automation

* Ability to build the server in a Docker container
* Ability to run the server
* Ability to execute integration tests against the local Kubernetes cluster
* GitHub Actions for every commit that verifies code quality, ensures the software builds, and all tests pass

### Deployment

* Create a helm chart for the service

## Level 4

### Design

In addition to including details about the requirements below, the design
document should include a section describing how this solution could be deployed
to an AWS EKS cluster. Please note that you are not expected to implement any
deployment automation for an EKS target, or create any AWS resources.

### Server

* HTTP API to retrieve the replica count of the Kubernetes Deployment
* HTTP API to set the replica count of a the Kubernetes Deployment
* HTTP API to get the list of available Deployments in the Kubernetes cluster
* HTTP health check verifying Kubernetes connectivity
* Extend the API to support a difference between current and desired state
* Extend the Server to support the ability to show the differences between current and desired state
* Secure connections with mTLS

### Automation

* Ability to build the server in a Docker container
* Ability to run the server
* Ability to execute integration tests against the local Kubernetes cluster
* GitHub Actions for every commit that verifies code quality, ensures the software builds, and all tests pass
* Produce production ready releases (binaries and docker image)

### Deployment

* Create a helm chart for the service
* Deploy releases of the API Server to a Kubernetes cluster (and possible dependencies)
* Upgrade releases of the API Server to a Kubernetes cluster (and possible dependencies)
* Include a production Deployment of this service, including but not limited to: Deployment, ServiceAccount, Service


## Level 5

### Design

In addition to including details about the requirements below, the design
document should include a section describing how this solution could be deployed
to an AWS EKS cluster. Please include security considerations as well.

### Server

* HTTP API to retrieve the replica count of the Kubernetes Deployment
* HTTP API to set the replica count of a the Kubernetes Deployment
* HTTP API to get the list of available Deployments in the Kubernetes cluster
* HTTP health check verifying Kubernetes connectivity
* Extend the API to support a difference between current and desired state
* Extend the Server to support the ability to show the differences between current and desired state
* Secure connections with mTLS
* Replace the HTTP API with gRPC
* Extend the Server to support reconciling cluster state (i.e. an external actor changed the replica count manually)

### Automation

* Ability to build the server in a Docker container
* Ability to run the server
* Ability to execute integration tests against the local Kubernetes cluster
* GitHub Actions for every commit that verifies code quality, ensures the software builds, and all tests pass
* Produce production ready releases (binaries and docker image)
* Deploy and manage a local Kubernetes Cluster
* Deploy and upgrade to the local Kubernetes Cluster with no service interruption
* GitHub Action driven workflows for Deployment

### Deployment

* Create a helm chart for the service
* Deploy releases of the API Server to a Kubernetes cluster (and possible dependencies)
* Upgrade releases of the API Server to a Kubernetes cluster (and possible dependencies)
* Include a production Deployment of this service, including but not limited to: Deployment, Role, RoleBinding, ServiceAccount, Service

# Guidance

## Interview process

The interview team will join the Slack channel. The team consists of the
engineers who will be working with you. Ask them about the engineering culture,
work and life balance, or anything else that you would like to learn about
Teleport.

Before writing the actual code, create a brief design document in markdown in Github
and share with the team.

This document should cover at least: design approach, trade-offs, scope,
proposed API, resiliency, scalability and deployment considerations.

Please avoid writing an overly detailed design document. Use this document to
make sure the team could provide feedback on your design and demonstrate that
you've investigated the problem space.

Split your code submission using pull requests and give the team an opportunity
to review the PRs. A good "rule of thumb" to follow is that the final PR
submission is a formality adding a small feature set - it means that the team
had an opportunity to contribute the feedback during multiple well defined
stages of your work.

Our team will do their best to provide a high quality review of the submitted
pull requests in a reasonable time frame. You are spending your time on this, we
are going to contribute our time too.

After the final submission, the interview team will assemble and vote using a
"+1, -2" anonymous voting system: +1 is submitted whenever a team member accepts
the submission, -2 otherwise.

In case of a positive result, we will connect you to our HR and recruiting
teams, who will work out the details and present an offer.

In case of a negative score result, hiring manager will contact you and share a
list of the key observations from the team that affected the result.

## Code and project ownership

This is a test challenge and we have no intent of using the code you've
submitted in production. This is your work, and you are free to do whatever you
feel  is reasonable with it. In the scenario where you don't pass, you can open
source it with any license and use it as a portfolio project.

## Areas of focus

Teleport focuses on simplicity, automation and robustness.

These are the areas we will be evaluating in the submission:

* Use consistent coding style. We follow
  [Go Coding Style](https://github.com/golang/go/wiki/CodeReviewComments) for
  the Go language. If you are going to other tools and/or languages, please pick
  coding style guidelines and let us know what they are.
* At the minimum, create tests for reading and writing, and an unhappy/error scenario.
* Make sure builds are reproducible. Pick any vendoring/packaging system that
  will allow us to get consistent build results.
* Ensure error handling and error reporting is consistent. The system should
  report clear errors and not crash under non-critical conditions.
* Production readiness. Once completed, the code itself, even if incomplete, should
  be sufficiently solid and robust to make it to a real production cluster.

## Trade-offs

Write as little code as possible, otherwise this task will consume too much time
and code quality will suffer.

Please cut corners, for example configuration tends to take a lot of time, and
is not important for this task.

Use hardcoded values as much as possible and simply add TODO items showing your
thinking, for example:

```
  // TODO: Add configuration system.
  // environment variables and reasonable default values,
  // for example https://github.com/alecthomas/kingpin
  // TODO: Add retry logic
```

Comments like this one are really helpful to us. They save yourself a lot of
time and demonstrate that you've spent time thinking about this problem and
provide a clear path to a solution.

Consider making other reasonable trade-offs. Make sure you communicate them to
the interview team.

Here are some other trade-offs that will help you to spend less time on the task:

* Do not implement a system that scales or is highly performing. Describe which
  performance improvements you would add in the future.
* It is OK if the system is not highly available. Write down how you would make
  the system highly-available and why your system is not.

## Pitfalls and Gotchas

To help you out, we've composed a list of things that previously resulted in a no-pass from the interview team:

* Scope creep. Candidates have tried to implement too much and ran out of time
  and energy. To avoid this pitfall, use the simplest solution that will work.
  Avoid writing too much code. We've seen candidates' code introducing caching
  and making many mistakes in the caching layer validation logic. Not having
  caching would have solved this problem.
  
* Overly complex designs. Keep things simple and try and eliminate as many moving
  parts as possible. This is not only going to help in reviewing the solution,
  but is also often a way to distill a design to its essential parts.
  
* Unstructured code. We've seen candidates leaving commented chunks of code,
  having one large file with all the code, not having code structure at all.
  
* Not communicating. Some candidates have submitted all their code to the master
  branch without raising pull requests, which does not give us the ability to
  provide feedback on the various implementation phases. We are a distributed
  team, so structured, asynchronous communication is critical to us.

## Questions
 
It is OK to ask the interview team questions. Some folks stay away from asking
questions to avoid appearing less experienced, so we provide examples of
questions to ask and questions we expect candidates to figure out on their own.

Here is a great question to ask:

> Is it OK to assume there will be only a single target kubernetes cluster for
  this service? I will add a note on how support for multiple clusters could be
  implemented, but it feels like an unnecessary complexity.

This is the question we expect candidates to figure out on their own:

> What version of Go should I use? What dependency manager should I use? What
  framework/tool should I use to automate testing and deployment ?

Unless specified in the requirements, pick the solution that works best for you.

# Tools

This task should be implemented in Go and should work on a x86 64-bit Linux or MacOS
machine.

It's safe to assume a working Docker environment will be available locally as well.

Additional external dependencies are acceptable such as Terraform and minikube/k3d,
but please ensure your Makefile has targets to install these dependencies.

# Timing

You can split coding over a couple of weekdays or weekends and find time to ask
questions and receive feedback.

Once you join the Slack channel, you have up to 2 weeks to complete the
challenge.

Within this time frame, we don't give higher scores to challenges submitted more
quickly. We only evaluate the quality of the submission.
