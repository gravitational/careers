# Summary

Implement a Go server that interacts with a Kubernetes cluster, incorporating
automated builds, containerization, deployment, and testing.

# Rationale

This exercise has two goals:

* It helps us to understand what to expect from you as a developer, how you
  write production code, how you reason about API design and how you
  communicate when trying to understand a problem before you solve it.
* It helps you get a feel for what it would be like to work at Teleport, as
  this exercise aims to simulate our day-as-usual and expose you to the type of
  work we're doing here.

We believe this technique is not only better, but also is more fun compared to
whiteboard/quiz interviews so common in the industry. It's not without the
downsides - it could take longer than traditional interviews.

[Some of the best teams use coding challenges.](https://sockpuppet.org/blog/2015/03/06/the-hiring-post/)

We appreciate your time and are looking forward to hack on this project
together.

# Levels

There are 6 engineering levels at Teleport. It's possible to score on level 1-5
through coding challenge.

Level 6 is only for internal promotions. Check
[Systems Engineering Levels](../../levels/systems.pdf) for more details.

# Interview Process

The interview process will start with you receiving an invite to a private
Slack channel. That channel will contain the interview panel. You can ask them
about the engineering culture, work-life balance, or anything else that you
would like to learn about Teleport.

## Design Doc

Before writing any actual code, we ask that you write a brief design document.
The design document should cover: design approach, scope, proposed APIs,
security considerations, CLI UX, and implementation details where appropriate.
Start with a brief doc that covers the edge cases and design approach.

Please submit the design document and all code in a GitHub repository. Public
or private is your choice. Please submit the design document written in
markdown as a Pull Request to allow us to provide you feedback on the proposed
design.

A few notes about the design document:

* Try to get the design document approved within the first 2-3 days. This is to
  ensure you have enough time to work on the implementation.
* Avoid writing an overly detailed design document. Two to three pages is
  sufficient.
* Avoid sending us draft design documents. It is difficult to evaluate which
  parts are draft and which parts are complete. Instead we encourage asking
  questions in Slack and sharing a design document that is ready to be reviewed.

Be sure to cover the following in your design:

* API structure
* Pod Lifecycle
* TLS Configuration
* Developer workflow
  * Ease of contributing to the project from a fresh clone
  * Ease of building, running and testing the server
* Level 3+: Build and Release
* Level 4+: Caching and mTLS
* Level 5+: Reconciliation, Conflicts, and Automation

Once the design document has been approved by two reviewers, move on to the
implementation.

## Implementation

Split your code submission into roughly 3-5 Pull Requests to give the team an
opportunity to review your code and provide feedback. Feel free to merge each
PR after you have two approvals.

Our team will do their best to provide a high quality review of the submitted
Pull Requests in a reasonable time frame. You are spending your time on this,
we are going to contribute our time too.

After the final submission, the interview team will assemble and vote using a
"+1, -2" anonymous voting system: +1 is submitted whenever a team member
accepts the submission, -2 otherwise.

In case of a positive result, we will connect you to our HR and recruiting
teams, who will work out the details and present an offer.

In case of a negative score result, the hiring manager will contact you and
share a list of the key observations from the team that affected the result.

### Tools

This task should be written in Go and is deployable to a local Kubernetes
cluster. The choice of which local Kubernetes cluster is up to you, but please
ensure compatibility with both macOS and Linux. We suggest
[KIND](https://kind.sigs.k8s.io/).

You may use additional external dependencies, but ensure that detecting or
installing these are straightforward for the reviewer. At a minimum, your
solution should include the following dependencies:

* Go
* make
* Docker

### Testing

Key components of the challenge should have tests that cover the happy and
unhappy scenarios. Do not try to achieve 100% test coverage as that will take
too long.

# Requirements

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
* Ability to execute integration tests against the local Kubernetes cluster by
  following documentation

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
* Ability to execute integration tests against the local Kubernetes cluster by
  following documentation

### Deployment

* Create a Helm chart for the service that includes at least: a Deployment,
  ServiceAccount and Service
* Upgrading the Helm chart should not result in unavailability of the service

## Level 4

### Server

* HTTP API to retrieve the replica count of the Kubernetes Deployment
* HTTP API to set the replica count of the Kubernetes Deployment
* HTTP API to get the list of available Deployments in the Kubernetes cluster
* HTTP API must cache the replica count by watching for changes to Deployments.
  Read-only requests should not each trigger a request to the cluster.  It is
  acceptable to use either client-go or controller-runtime to implement this
* HTTP health check verifying Kubernetes connectivity
* Secure connections between the HTTP API and caller with mTLS
* One or two tests that cover happy and unhappy scenarios

### Automation

* Ability to build and deploy all artifacts to a Kubernetes cluster using make
* Ability to execute integration tests against the local Kubernetes cluster
  using make

### Deployment

* Create a configurable Helm chart for the service
* Includes at a minimum: a Deployment, ServiceAccount, and Service
* Upgrading the Helm chart should not result in unavailability of the service

## Level 5

### Server

* gRPC API to retrieve the replica count of the Kubernetes Deployment
* gRPC API to set the replica count of the Kubernetes Deployment
* gRPC API to get the list of available Deployments in the Kubernetes cluster
* gRPC API must cache the replica count by watching for changes to Deployments.
  Read-only requests should not each trigger a request to the cluster.  It is
  acceptable to use either client-go or controller-runtime to implement this.
* Implement a controller to store the desired state in a CRD (per-deployment)
  and reconcile the deployment to that state.
* gRPC or HTTP health check verifying Kubernetes connectivity
* Secure connections between the gRPC API and caller with mTLS
* One or two tests that cover happy and unhappy scenarios

### Automation

* Ability to build and deploy all artifacts to a Kubernetes cluster using make
* Ability to execute integration tests against the local Kubernetes cluster
  using make

### Deployment

* Create a configurable Helm chart for the service
* Include production-level packaging for this service, including but not
  limited to: Deployment, Role, RoleBinding, ServiceAccount, and Service
* Upgrading the Helm chart should not result in unavailability of the HTTP API
  portion of the service.

# Guidance

## Code and project ownership

This is a test challenge and we have no intent of using the code you've
submitted in production. This is your work, and you are free to do whatever you
feel is reasonable with it. In the scenario where you don't pass, you can open
source it with any license and use it as a portfolio project.

## Areas of focus

Teleport focuses on networking, infrastructure and security.

These are the areas we will be evaluating in the submission:

* Use consistent coding style. We follow [Go Coding
  Style](https://go.dev/wiki/CodeReviewComments) for the Go
  language.
* Tests exist for happy path and error scenarios for key components of the
  challenge.
* Make sure builds are reproducible.
* Ensure error handling and error reporting is consistent. The system should
  report clear errors and not crash under non-critical conditions.
* Production readiness. Once completed, the code itself should be bug free and
  reliable. Cut scope, but not quality.
* API design. Please include your proposed HTTP API or gRPC API in the design
  doc. For the gRPC API, you should include a complete proto file in the design
  doc.
* Security. Describe your mTLS setup in the design doc, including chosen cipher
  suites. Ensure that your implementation is secure.
* Project management and scope. Manage your time wisely and ensure that the
  project scope aligns with the criteria for the level you're applying for.
  Avoid unnecessary complexity.

The primary factor in the team's decision is overall code quality. We are
looking for the highest possible quality with the smallest possible scope that
meets the requirements of the challenge.

## Trade-offs

Write as little code as possible, otherwise this task will consume too much
time and code quality will suffer.

Please cut corners, for example configuration tends to take a lot of time, and
is not important for this task.

Use hardcoded values as much as possible and simply add TODO items showing your
thinking, for example:

```go
// TODO: Add configuration system.
// Consider using CLI library to support both environment variables and
// reasonable default values, for example https://github.com/alecthomas/kingpin

// TODO: Add retry logic.
```

Comments like this one are really helpful to us. They save yourself a lot of
time and demonstrate that you've spent time thinking about this problem and
provide a clear path to a solution.

Consider making other reasonable trade-offs. Make sure you communicate them to
the interview team.

## Pitfalls and Gotchas

To help you out, we've composed a list of things that previously resulted in a
no-pass from the interview team:

* Use of AI. Don't outsource your thinking to an AI. We recommend using AI for
  use cases like learning about a new problem space, exploring APIs, and
  finding missing edge cases. However, we strongly recommend you write the design
  document and all code yourself.
* Scope creep. Candidates have tried to implement too much and ran out of time
  and energy. To avoid this pitfall, use the simplest solution that will work.
  Avoid writing too much code. For example, it is not necessary to deploy a
  shared cached or support deployment to multiple regions or AZs.
* Overly complex designs. Keep things simple and try and eliminate as many
  moving parts as possible. This is not only going to help in reviewing the
  solution, but is also often a way to distill a design to its essential parts.
* Unstructured code. We've seen candidates leaving commented chunks of code,
  having one large file with all the code, not having code structure at all.
* Not communicating. Some candidates have submitted all their code to the
  master branch without raising pull requests, which does not give us the
  ability to provide feedback on the various implementation phases. We are a
  distributed team, so structured, asynchronous communication is critical to us.

## Questions

It is OK to ask the interview team questions. Some folks stay away from asking
questions to avoid appearing less experienced, so we provide examples of
questions to ask and questions we expect candidates to figure out on their own.

Here is a great question to ask:

> Is it OK to pre-generate secret data and put the secrets in the repository
> for a proof of concept? I will add a note that we will auto-generate secrets
> in the future.

It demonstrates that you thought about this problem domain, recognize the trade
off and are saving you and the team time by not implementing it.

This is the question we expect candidates to figure out on their own:

> What version of Go should I use? What dependency manager should I use? What
> framework/tool should I use to automate testing and deployment?

Unless specified in the requirements, pick the solution that works best for
you.

# Timing

You can split coding over a couple of weekdays or weekends and find time to ask
questions and receive feedback.

Once you join the Slack channel, you will have 2 weeks complete the challenge.

Within this timeframe, we don't give higher scores to challenges submitted more
quickly. We only evaluate the quality of the submission.

We always aim to provide 1-2 rounds of feedback on all work that is submitted. 
In order to be respectful of your time, we may opt to end the challenge early
if the submission does not improve after this feedback is suggested or if we
identify a large number of issues.
