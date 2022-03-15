# Summary

Implement a prototype job worker service that provides an API to run arbitrary Linux processes.

# Rationale

This exercise has two goals:

* It helps us to understand what to expect from you as a developer, how you
  write production code, how you reason about API design and how you communicate
  when trying to understand a problem before you solve it.
* It helps you get a feel for what it would be like to work at Teleport, as this
  exercise aims to simulate our day-as-usual and expose you to the type of work
  we're doing here.

We believe this technique is not only better, but also is more fun compared to
whiteboard/quiz interviews so common in the industry. It's not without the
downsides - it could take longer than traditional interviews.

[Some of the best teams use coding challenges.](https://sockpuppet.org/blog/2015/03/06/the-hiring-post/)

We appreciate your time and are looking forward to hack on this project together.

# Requirements

There are 6 engineering levels at Teleport. It's possible to score on level 1-5 through coding challenge.

Level 6 is only for internal promotions. Check the
[engineering levels document](https://raw.githubusercontent.com/gravitational/careers/main/levels.pdf)
for more details.

Split the submission into 2-3 pull requests for us to review. We will review
every pull request and provide our feedback.

We are going to compile the program, test it and get back to you.

## Design Doc

Start with a brief doc that covers the edge cases and design approach. At
Teleport, we prefer Markdown for
[our designs](https://github.com/gravitational/teleport/blob/master/rfd/0000-rfds.md).

Be sure to cover the following in your design:

* CLI user experience: a couple examples of what it may look like to invoke a
  command. This allows us an opportunity to envision how we will run the program.
* Level 3+: Streaming approach, TLS setup (version, cipher suites, etc.)
* Level 4+: Process execution lifecycle (how will you start and stop jobs, add
  them to cgroups, etc.)

Open a pull request with your design and share a link with the reviewers via
Slack. After the doc is approved, implement interfaces and an example program
using the library.

A few notes about the design document:

* We expect the design document to be complete roughly within the first week.
  This is to ensure you have enough time to work on the implementation.
* Avoid writing an overly detailed design document. Two to three pages is
  sufficient.
* Avoid sending us draft design documents, it is difficult to evaluate what
  parts are draft and which parts are complete. Instead we encourage asking
  questions in Slack and sharing a design document that is ready to be
  reviewed.

## Testing

Add a couple of high quality tests that cover happy and unhappy scenarios.

Do not try to achieve full test coverage. This will take too long. Take two key
components, e.g. authentication/authorization layer and networking and implement
one or two test cases that demonstrate your approach to testing.

## Level 1

### Library

* Worker library with methods to start/stop/query status and get the output of a job.

### API

* HTTPS API to start/stop/get status of a running process.
* Use HTTP Basic Authentication.
* Use a simple authorization scheme.

### CLI

* CLI should be able to connect to worker service and start, stop, get status, and output of a job.

## Level 2

### Library

* Worker library with methods to start/stop/query status and get the output of a job.

### API

* HTTPS API to start/stop/get status of a running process.
* Use mTLS authentication and verify client certificate. Set up strong set of
  cipher suites for TLS and good crypto setup for certificates. Do not use any
  other authentication protocols on top of mTLS.
* Use a simple authorization scheme.

### CLI

* CLI should be able to connect to worker service and start, stop, get status, and output of a job.

## Level 3

### Library

* Worker library with methods to start/stop/query status and get the output of a job.
* Library should be able to stream the output of a running job.
  * Output should be from start of process execution.
  * Multiple concurrent clients should be supported.

### API

* [GRPC](https://grpc.io) API to start/stop/get status/stream output of a running process.
* Use mTLS authentication and verify client certificate. Set up strong set of
  cipher suites for TLS and good crypto setup for certificates. Do not use any
  other authentication protocols on top of mTLS.
* Use a simple authorization scheme.

### Client

* CLI should be able to connect to worker service and start, stop, get status, and stream output of a job.

## Level 4

### Library

* Worker library with methods to start/stop/query status and get the output of a job.
* Library should be able to stream the output of a running job.
  * Output should be from start of process execution.
  * Multiple concurrent clients should be supported.
* Add resource control for CPU, Memory and Disk IO per job using cgroups.

### API

* [GRPC](https://grpc.io) API to start/stop/get status/stream output of a running process.
* Use mTLS authentication and verify client certificate. Set up strong set of
  cipher suites for TLS and good crypto setup for certificates. Do not use any
  other authentication protocols on top of mTLS.
* Use a simple authorization scheme.

### Client

* CLI should be able to connect to worker service and start, stop, get status, and stream output of a job.

## Level 5

### Library

* Worker library with methods to start/stop/query status and get the output of a job.
* Library should be able to stream the output of a running job.
  * Output should be from start of process execution.
  * Multiple concurrent clients should be supported.
* Add resource control for CPU, Memory and Disk IO per job using cgroups.
* Add resource isolation for using PID, mount, and networking namespaces.

### API

* [GRPC](https://grpc.io) API to start/stop/get status/stream output of a running process.
* Use mTLS authentication and verify client certificate. Set up strong set of
  cipher suites for TLS and good crypto setup for certificates. Do not use any
  other authentication protocols on top of mTLS.
* Use a simple authorization scheme.

### Client

* CLI should be able to connect to worker service and start, stop, get status, and stream output of a job.

# Guidance

## Interview process

The interview team will join the Slack channel. The team consists of the
engineers who will be working with you. Ask them about the engineering culture,
work and life balance, or anything else that you would like to learn about
Teleport.

Before writing the actual code, create a brief design document in Google Docs or
markdown in Github and share with the team.

This document should cover at least: design approach, trade-offs, scope,
proposed API, and security considerations. For security, please include
considerations for transport layer, authentication, and authorization.

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
feel is reasonable with it. In the scenario where you don't pass, you can open
source it with any license and use it as a portfolio project.

## Areas of focus

Teleport focuses on networking, infrastructure and security.

These are the areas we will be evaluating in the submission:

* Use consistent coding style. We follow
  [Go Coding Style](https://github.com/golang/go/wiki/CodeReviewComments) for
  the Go language. If you are going to use a different language, please pick
  coding style guidelines and let us know what they are.
* At the minimum, create tests for authentication, networking, and an
  unhappy/error scenario.
* Make sure builds are reproducible. Pick any vendoring/packaging system that
  will allow us to get consistent build results.
* Ensure error handling and error reporting is consistent. The system should
  report clear errors and not crash under non-critical conditions.
* Avoid concurrency and networking errors. Most of the issues we've seen in
  production are related to data races, networking error handling or goroutine
  leaks. We will be looking for those errors in your code.
* Security. Use strong authentication and simplest, but robust authorization.
  Set up the strongest transport encryption you can. Test it.

## Trade-offs

Write as little code as possible, otherwise this task will consume too much time
and code quality will suffer.

Please cut corners, for example configuration tends to take a lot of time, and
is not important for this task.

Use hardcoded values as much as possible and simply add TODO items showing your
thinking, for example:

```
  // TODO: Add configuration system.
  // Consider using CLI library to support both
  // environment variables and reasonable default values,
  // for example https://github.com/alecthomas/kingpin
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
  the system highly available and why your system is not.

## Pitfalls and Gotchas

To help you out, we've composed a list of things that previously resulted in a no-pass from the interview team:

* Scope creep. Candidates have tried to implement too much and ran out of time
  and energy. To avoid this pitfall, use the simplest solution that will work.
  Avoid writing too much code. We've seen candidates' code introducing caching
  and making many mistakes in the caching layer validation logic. Not having
  caching would have solved this problem.
* Data races. We will scan the code with a race detector and do our best to find
  data races in the code. Avoid global state as much as possible; if using
  global state, write down a good description why it is necessary and protect it
  against data races.
* Deadlocks. When using mutexes, channels or any other synchronization
  primitives, make sure the system won't deadlock. We've seen candidates' code
  holding a mutex and making a network call without timeouts in place. Be extra
  careful with networking and sync primitives.
* Unstructured code. We've seen candidates leaving commented chunks of code,
  having one large file with all the code, not having code structure at all.
* Not communicating. Some candidates have submitted all their code to the master
  branch without raising pull requests, which does not give us the ability to
  provide feedback on the various implementation phases. We are a distributed
  team, so structured, asynchronous communication is critical to us.
* Implementing custom security algorithms/authentication schemes is always a bad
  idea unless you are a trained security researcher/engineer. It is definitely a
  bad idea for this task - try to stick to industry proven security methods as
  much as possible.

## Questions

It is OK to ask the interview team questions. Some folks stay away from asking
questions to avoid appearing less experienced, so we provide examples of
questions to ask and questions we expect candidates to figure out on their own.

Here is a great question to ask:

> Is it OK to pre-generate secret data and put the secrets in the repository for
> a proof of concept? I will add a note that we will auto-generate secrets in
> the future.

It demonstrates that you thought about this problem domain, recognize the trade
off and are saving you and the team time by not implementing it.

This is the question we expect candidates to figure out on their own:

> What version of Go should I use? What dependency manager should I use?

Unless specified in the requirements, pick the solution that works best for you.

# Tools

This task should be implemented in Go or Rust and should work on 64-bit Linux machines.

# Timing

You can split coding over a couple of weekdays or weekends and find time to ask
questions and receive feedback.

Once you join the Slack channel, you have between 1 to 2 weeks complete the
challenge depending on the challenge you choose.

Within this timeframe, we don't give higher scores to challenges submitted more
quickly. We only evaluate the quality of the submission.

We only start the coding challenge if there are several open positions available.
