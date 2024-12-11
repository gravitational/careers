# SRE - Security

The interview process is divided into two main sections: a walkthrough and a challenge. In the walkthrough, led by the hiring manager, you’ll have an opportunity to share your work history, review our leveling matrix, and discuss the upcoming challenge. Following the walkthrough, the challenge will assess your skills and fit for this position.

We will use the challenge in order to evaluate your skill in the following areas:
* Translating high-level requirements into a simple, functional design.
* Writing production level code that does not have extensive or insecure dependencies.
* Understanding of Go, Kubernetes, release engineering, and security.
* Communicating with the team and handling feedback.

We believe this technique is not only better but also more fun compared to
whiteboard/quiz interviews so common in the industry. It’s not without the
downsides - it could take longer than traditional interviews. That said, it's
our view that this type of challenge gives us a more accurate assessment of your
ability to work well on the types of projects we’re working on day-to-day here
at Teleport. [Some of the best teams use coding
challenges](https://sockpuppet.org/blog/2015/03/06/the-hiring-post/). We
appreciate your time and are looking forward to hacking on this project
together.

Please come prepared to the walkthrough having reviewed:
* [Site Reliability Engineering (SRE) Levels](../../levels/sre.pdf)
* [Challenge documentation](challenge-security.md)

# Challenge

In this challenge, you will build a Go API server interacting with Redis and a Go service to report on its SLOs. This solution will be deployed to Kubernetes with automated builds, containerization, testing, and secure practices.

## Getting Started

First, create a new private GitHub repository, invite the interview panel as contributors, and share the link with the team in Slack.

> [!IMPORTANT]
> You’ll have up to 2 weeks from the agreed start date to complete the challenge. Please allow 24-48 hours for each PR submission to give the team ample time for review and feedback during business hours.

### PR Submissions
The repository you created be used for each PR submission during the challenge.  For each submission, please ensure the following:

* Your submissions meets the requirements of the level you are applying for.
* Split the work into 3 pull requests: one for the design document, one for the server code, and one covering service levels and deployment.
* After submitting, the team will review and provide feedback. Open one PR at a time and wait for two approvals before proceeding with submitting another. You can continue working locally while waiting for feedback on the current PR.

## Design Doc
The first pull request must be a brief design document that describes how you plan to implement the solution. At Teleport, we prefer Markdown for [our designs](https://github.com/gravitational/teleport/blob/master/rfd/0000-rfds.md).

### Requirements

Please be sure to cover the following design topics:

* API structure
* Service Levels
  * Service design
  * Selected objectives and indicators
* Developer workflow
  * Ease of building, running and testing the server
* Build, security, and deployment
  * Certificates and ciphers
  * RBAC and network policy
  * Selected invariants

### Suggestions
* Complete the design document within the first week.  This is to ensure you have enough time to work on the implementation.
* Avoid writing an overly detailed design document. 500-1500 words is a good target.
* We encourage you to ask questions in Slack.
* Avoid sending draft PRs for feedback as it is difficult to evaluate which parts are draft or complete.

## Implementation

### Server
Once the design document is approved, the next PR should include your server code, along with build automation to simplify testing.  Use the [redis-cli](https://redis.io/learn/guides/import) to import the [test data](files/data.redis) provided as this is needed for the final solution. The interview panel will clone the repository and run Make targets to build and test your server code. Be sure to include a couple of high-quality tests that cover both happy and unhappy scenarios, and minimize external dependencies.

### Service Levels & Deployment

After the server code is approved, submit a final PR including service levels, deployment, and any remaining requirements.

Your chosen service levels should focus on meaningful indicators, be set at realistic thresholds, and align with the system's functionality and user experience.  The presentation of these SLOs is up to the candidate's discretion.  Consider a simple solution and include comments on how a more robust solution could be implemented (see [trade-offs](#trade-offs)).

When selecting [security invariants](https://d1.awsstatic.com/events/aws-reinforce-2022/NIS234_Learn-to-create-end-to-end-security-invariants-using-AWS-services.pdf), consider which properties you expect your system to always maintain.  The chosen invariants should cover attack vectors against the most critical parts of your system design. How these properties are evaluated and reported is up to the candidate.

When securing your solution, assess the trust model for service-to-service communication and interactions with end users, ensuring all connections are protected with mTLS. The final implementation should incorporate zero-trust principles across the Helm chart, build artifacts, API design, etc.

### Tooling
For evaluation purposes, your solution should be written in Go and is deployable to a local [KIND](https://kind.sigs.k8s.io/) Kubernetes cluster. You may use additional external dependencies, but ensure that detecting or installing these are straightforward for the reviewer. At a minimum, your solution should include the following dependencies:

* Go
* make
* Docker
* Kind

## Level 4

### Server
* HTTPS API to create, get, or list users with Redis datastore
* HTTPS API to add a new payment with Redis datastore
* HTTPS API to get payments by user with Redis datastore
* HTTPS Health check ensuring availability
* Secure connections between the HTTPS API and caller with mTLS
* Include one or two tests that cover happy and unhappy path scenarios

### Service Levels
* Define three Service Level Objectives (SLOs) for the API service
* Create a service to periodically evaluate the defined SLOs and report on their status
* SLO service communication is secure

### Deployment
* Create a Helm chart for the API server
* Single make target to build and deploy all artifacts and certificates to a Kubernetes cluster
* Ensure artifacts and services are production ready
* Evaluate the state of three security invariants.

## Level 5

### Server
* gRPC API to create, get, or list users with Redis datastore
* gRPC API to add a new payment with Redis datastore
* gRPC API to get payments by user with Redis datastore
* gRPC or HTTP Health check ensuring availability
* Secure connections between the HTTP API and caller with mTLS
* Include one or two tests that cover happy and unhappy path scenarios
* Service should handle autoscaling without disruption or data duplication

### Service Levels
* Define five Service Level Objectives (SLOs) for the API service
* Create a service to periodically evaluate the defined SLOs and report on their status
* SLO service communication is secure

### Deployment
* Create a Helm chart for the API and SLO services
* Single make target to build and deploy all artifacts and certificates to a Kubernetes cluster
* Ensure artifacts and services are production ready
* Evaluate the state of five security invariants
* Automatically autoscale API server
* Deploy a zero-trust implementation of Prometheus and Loki

# Guidance

## Areas of focus

The primary factor in the team's decision is overall code quality. We are looking for the highest possible quality with the smallest possible scope that meets the requirements of the challenge.

* Use consistent coding style. Internally we follow [Go Coding Style](https://github.com/golang/go/wiki/CodeReviewComments).
* Make sure builds are reproducible and allow consistent build results.
* Ensure error handling and error reporting is consistent. The system should report clear errors and not crash under non-critical conditions.
* Production readiness. Once completed, the code itself, even if incomplete, should be sufficiently solid and robust to make it to a real production cluster.
* API design. Please include your proposed HTTP API or gRPC API in the design doc. For the gRPC API, you should include a complete proto file in the design doc.
* Security. Describe your mTLS setup in the design doc, including chosen cipher suites. Ensure service to service communication and service to end users are adequately secured. Consider how your service could be exploited.
* Project management and scope. Manage your time wisely and ensure that the project scope aligns with the criteria for the level you're applying for. Avoid unnecessary complexity.


## Trade-offs
Write as little code as possible to avoid letting the project consume too much time, which can impact code quality. Cut corners where appropriate; for example, avoid complex configuration and use hardcoded values. Add TODO items to indicate future enhancements or considerations, such as:

```
  // TODO: Add configuration system.
  // environment variables and reasonable default values,
  // for example https://github.com/alecthomas/kingpin
  // TODO: Add retry logic
```

Comments like this one are really helpful. They save yourself time and demonstrate that you've spent time thinking about this problem and provide a clear path toward a longer-term solution.

Consider making other reasonable trade-offs. Make sure you communicate them to the interview team.

### Pitfalls and Gotchas
To help you out, we've composed a list of things that previously resulted in a no-pass from the interview team:

* Scope creep. Candidates have tried to implement too much and ran out of time
  and energy. To avoid this pitfall, use the simplest solution that will work.
  Avoid writing too much code.
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

## Code and project ownership

This is a test challenge, and we have no intention of using the code in production. The work is yours, and you’re free to handle it as you see fit. If you don’t pass, you’re welcome to open-source it under any license and use it as a portfolio project.
