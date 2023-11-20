# Summary

Monitor, deploy, and extend a prototype service that provides an API for
Kubernetes resources.

# Rationale

This exercise has two goals:

* It helps us to understand what to expect from you as a Site Reliability Engineer,
  how you reason about a production service, how you reason about system design, your
  opinions on automation & tooling, and how you communicate when trying to solve problems.
* It helps you get a feel for what it would be like to work at Teleport, as this
  exercise aims to simulate our day-as-usual and expose you to the type of work
  we're doing here.

We believe this technique is not only better, but also is more fun compared to
whiteboard/quiz interviews so common in the industry. It's not without the
downsides - it could take longer than traditional interviews.

[Some of the best teams use coding challenges.](https://sockpuppet.org/blog/2015/03/06/the-hiring-post/)

We appreciate your time and are looking forward to hacking on this project together.

# Requirements

In this challenge you'll extend a Go server by adding additional APIs that interact
with a Kubernetes environment. You'll also add automation to build the server,
create a container image, run the server, and execute tests.

The requirements vary depending on the level you are applying to. This
challenge covers 5 engineering levels at Teleport. Level 6 is only for internal
promotions. Check [Site Reliability Engineering (SRE) Levels](../../levels/sre.pdf) for more details.

An incomplete implementation of a Go server is located at
[sre-server/main.go](./sre-server/main.go).

Start by creating a new GitHub repository. Then let the interview panel know the
repository's location by pasting a link in your interview Slack channel. Invite
interview panel participants as contributors to the new repository if you prefer
to keep your submission private.

* Your submission will need to meet the requirements of the level you are applying for.
* Split the submission into 2-3 pull requests for us to review. We will review
  every pull request and provide feedback. Open a single pull request at a time
  and wait for at least 2 approvals before merging.
* Share a link with the interview panel on Slack each time you open a new pull request.
* The interview panel will clone your repository and test it after the last pull
  request is merged.

## Design Doc

The first pull request must be a brief design document that describes how you plan
to implement the solution. At Teleport, we prefer Markdown for [our designs](https://github.com/gravitational/teleport/blob/master/rfd/0000-rfds.md).

Be sure to cover the following in your design:

* API structure
* Developer workflow
  * Ease of contributing to the project from a fresh clone
  * Ease of building, running and testing the server
* Level 3+: Build, Release
* Level 4+: Caching, mTLS, and Delivery
* Level 5+: Reconciliation, Conflicts, and Automation

A few notes about the design document:

* We expect the design document to be completed roughly within the first week.
  This is to ensure you have enough time to work on the implementation.
* Avoid writing an overly detailed design document. 500-1500 words is
  sufficient.
* Avoid sending us draft design documents, because it is difficult to evaluate which
  parts are draft and which parts are complete. Instead we encourage asking
  questions in Slack.

## Server & Automation

Once the design document is approved, begin working on adding features to the
server implementation that includes automation. A key feature of this challenge
is to produce a self-contained GitHub project that automates away as much as
possible. This should include code quality, testing, build, and deployment.

Expect the interview panel to clone the repository and execute one or more
Make targets that build and test a working solution. Add a couple of high quality
tests that cover happy and unhappy scenarios. Keep the number of external
dependencies low.

Do not try to achieve full test coverage. This will take too long. Write enough
to exercise the different key components to show they are working as intendend 
while demonstrating your approach to automation.

For level 4+, please research Kubernetes controllers and the recommended Go
libraries carefully. Understanding how existing controllers cache resources
is key to implementing a straightforward solution.

## Deployment

For evaluation purposes, your solution should be deployable to a local
Kubernetes cluster. The choice of which Kubernetes cluster is up to you, but
please choose one that has the ability to run on macOS and Linux.
We suggest [KIND](https://kind.sigs.k8s.io/).

## Key Dependencies

You are welcome to solve this challenge using tools you are most familiar with,
but, at a minimum, please include the following:

* Go
* make
* Docker

## Level 1

### Server

* HTTP API to retrieve the replica count of the Kubernetes Deployment
* One or two tests that cover happy and unhappy scenarios

### Automation

* Write a Dockerfile to build an image for the server
* Ability to deploy the server to Kubernetes by following documentation

## Level 2

### Server

* HTTP API to retrieve the replica count of the Kubernetes Deployment
* HTTP API to set the replica count of the Kubernetes Deployment
* One or two tests that cover happy and unhappy scenarios

### Automation

* Write a Dockerfile to build an image for the server
* Ability to deploy the server to Kubernetes by following documentation
* Ability to execute integration tests against the local Kubernetes cluster by following documentation

## Level 3

### Server

* HTTP API to retrieve the replica count of the Kubernetes Deployment
* HTTP API to set the replica count of the Kubernetes Deployment
* HTTP API to get the list of available Deployments in the Kubernetes cluster
* HTTP health check verifying Kubernetes connectivity
* One or two tests that cover happy and unhappy scenarios

### Automation

* Write a Dockerfile to build an image for the server
* Ability to deploy the server to Kubernetes by following documentation
* Ability to execute integration tests against the local Kubernetes cluster by following documentation

### Deployment

* Create a helm chart for the service that includes at least: a Deployment, ServiceAccount and Service

## Level 4

### Server

* HTTP API to retrieve the replica count of the Kubernetes Deployment
* HTTP API to set the replica count of the Kubernetes Deployment
* HTTP API to get the list of available Deployments in the Kubernetes cluster
* HTTP health check verifying Kubernetes connectivity
* HTTP API must cache the replica count by watching for changes to Deployments.
  Read-only requests should not each trigger a request to the cluster.
  It is acceptable to use either client-go or controller-runtime to implement this.
* Secure connections between the HTTP API and caller with mTLS
* One or two tests that cover happy and unhappy scenarios

### Automation

* Ability to build and deploy all artifacts to a Kubernetes cluster using make
* Ability to execute integration tests against the local Kubernetes cluster using make

### Deployment

* Create a helm chart for the service that includes at least: a Deployment, ServiceAccount and Service

## Level 5

### Server

* gRPC API to retrieve the replica count of the Kubernetes Deployment
* gRPC API to set the replica count of the Kubernetes Deployment
* gRPC API to get the list of available Deployments in the Kubernetes cluster
* gRPC or HTTP health check verifying Kubernetes connectivity
* gRPC API must cache the replica count by watching for changes to Deployments.
  Read-only requests should not each trigger a request to the cluster.
  It is acceptable to use either client-go or controller-runtime to implement this.
* Secure connections between the gRPC API and caller with mTLS
* Server must store the desired state in a CRD (per-deployment) and reconcile the deployment to that state.
  (gRPC API endpoints only need to read the real, current value.) 
* One or two tests that cover happy and unhappy scenarios

### Automation

* Ability to build and deploy all artifacts to a Kubernetes cluster using make
* Ability to execute integration tests against the local Kubernetes cluster using make

### Deployment

* Create a helm chart for the service
* Include production-level packaging for this service, including but not limited to: Deployment, Role, RoleBinding, ServiceAccount, Service

# Guidance

## Interview process

The interview team will join the Slack channel. The team consists of the
engineers who will be working with you. Ask them about the engineering culture,
work and life balance, or anything else that you would like to learn about
Teleport.

Before writing the actual code, create a brief design document in markdown in GitHub
and share with the team.

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
  the Go language.
* At the minimum, create tests for reading and writing, and an unhappy/error scenario.
* Make sure builds are reproducible. Pick any vendoring/packaging system that
  will allow us to get consistent build results.
* Ensure error handling and error reporting is consistent. The system should
  report clear errors and not crash under non-critical conditions.
* Production readiness. Once completed, the code itself, even if incomplete, should
  be sufficiently solid and robust to make it to a real production cluster.
* API design. Please include your proposed HTTP API or gRPC API in the design doc.
  For the gRPC API, you should include a complete proto file in the design doc.
* Security. Describe your mTLS setup in the design doc, including choosen cipher suites.
  Ensure that your implementation is secure.

The primary factor in the team's decision is overall code quality. We are looking for
the highest possible quality with the smallest possible scope that meets the requirements
of the challenge.
  
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

Additional external dependencies are acceptable,
but please ensure detecting or installing the required/missing dependencies
is as low friction as possible for the user/reviewer.

# Timing

You can split coding over a couple of weekdays or weekends and find time to ask
questions and receive feedback.

Once you join the Slack channel, you have up to 2 weeks to complete the
challenge.

Within this time frame, we don't give higher scores to challenges submitted more
quickly. We only evaluate the quality of the submission.
