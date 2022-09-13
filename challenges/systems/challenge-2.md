# Summary

Implement a basic TCP load balancer. Load balancers are used to distribute
network load across multiple hosts (or upstreams).

# Rationale

This exercise has two goals:

* It helps us to understand what to expect from you as a developer, how you
  write code, how you reason about API design and how you communicate when
  trying to understand a problem before you solve it.
* It helps you get a feel for what it would be like to work at Teleport, as this
  exercise aims to simulate our day-as-usual and expose you to the type of work
  we're doing here.

We believe this technique is not only better, but also is more fun compared to
whiteboard/quiz interviews so common in the industry. It's not without downsides
- it could take longer than traditional interviews.

[Some of the best teams use coding challenges.](https://sockpuppet.org/blog/2015/03/06/the-hiring-post/)

We appreciate your time and are looking forward to hacking on this project
together.

# Levels

There are 6 engineering levels at Teleport. It's possible to score on level 1-5
through coding challenge.

Level 6 is only for internal promotions. Check
[Systems Engineering Levels](../../levels/systems.pdf) for more details.

# Interview process

The interview process will start with you receiving an invite to a private
Slack channel. That channel will contain the interview panel. You can ask them
about the engineering culture, work-life balance, or anything else that you
would like to learn about Teleport.

## Design Doc

Before writing any actual code, we ask that you write a brief design document.
The design document should cover: design approach, scope, proposed API, security
considerations, CLI UX, and implementation details where appropriate. Start with
a brief doc that covers the edge cases and design approach. At Teleport, we
prefer Markdown for
[our designs](https://github.com/gravitational/teleport/blob/master/rfd/0000-rfds.md).

Please submit the design document and all code in a GitHub repository. Public
or private is your choice. Please submit the design document as a Pull Request
to allow us to provide you feedback on the proposed design.

A few notes about the design document:

* Try to get the design document approved within the first 2-3 days. This is to
  ensure you have enough time to work on the implementation.
* Avoid writing an overly detailed design document. Two to three pages is
  sufficient.
* Avoid sending us draft design documents. It is difficult to evaluate which
  parts are draft and which parts are complete. Instead we encourage asking
  questions in Slack and sharing a design document that is ready to be reviewed.

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

In case of a negative score result, the hiring manager will contact you and share a
list of the key observations from the team that affected the result.

### Testing

Add a couple of high quality tests that cover happy and unhappy scenarios.

Do not try to achieve full test coverage. This will take too long. Take two key
components, e.g. authentication/authorization layer and networking and implement
one or two test cases that demonstrate your approach to testing.

# Requirements

The project is broken down into 2 components:

1. A reusable library implementing the functionality of the load balancer.
2. A server that exposes the library functionality while providing secure
   communication to clients and an authorization layer.

## Level 1

### Library

* Implement a connection forwarder to forward connections to an upstream server.
  Your connection forwarder should copy data in both directions (from the LB to
  the upstream server, and from the upstream server back to the client of the LB)

### Server

* Secure communication between client and server using TLS and an [opaque
  bearer token](https://latacora.micro.blog/2018/06/12/a-childs-garden.html)
  for authentication
* Develop a simple authorization scheme that defines what clients can use the
  load balancer; this scheme can be statically defined in code
* Accept and forward connections to upstreams using the library

## Level 2

### Library

* Implement a round-robin connection forwarder to forward connections to
  multiple upstreams.

### Server

* Use mTLS authentication to verify the client's identity
* Develop a simple authorization scheme that defines what upstreams are
  available to which clients; this scheme can be statically defined in code
* Accept and forward connections to upstreams using library

## Level 3

### Library

* Implement a least connections connection forwarder that tracks the number of
  connections per upstream.

### Server

* Use mTLS authentication to have the server verify identity of the client and
  client of the server
* Develop a simple authorization scheme that defines what upstreams are
  available to which clients; this scheme can be statically defined in code
* Accept and forward connections to upstreams using library

## Level 4

### Library

* Implement a least connections connection forwarder that tracks the number of
  connections per upstream.
* Implement a per-client connection rate limiter

### Server

* Use mTLS authentication to have the server verify identity of the client and
  client of the server
* Develop a simple authorization scheme that defines what upstreams are
  available to which clients; this scheme can be statically defined in code
* Accept and forward connections to upstreams using library

## Level 5

## Library

* Implement a least connections forwarder that tracks the number of
  connections per upstream.
* Implement a per-client connection rate limiter that tracks the number of
  client connections.
* Add health checking: unhealthy upstreams should not be eligible to receive
  connections until they are determined to be healthy.

## Server

* Use mTLS authentication to have the server verify identity of the client and
  client of the server
* Develop a simple authorization scheme that defines what upstreams are
  available to which clients; this scheme can be statically defined in code
* Accept and forward connections to upstreams using library

# Guidance

## Code and project ownership

This is a test challenge and we have no intent of using the code you've
submitted in production. This is your work, and you are free to do whatever you
feel is reasonable with it. In the scenario where you don't pass, you can open
source it with any license and use it as a portfolio project.

## Areas of focus

Teleport focuses on networking, infrastructure and security.

These are the areas we will be evaluating in the submission:

* Use consistent coding style. We follow [Go Coding Style]
  (https://github.com/golang/go/wiki/CodeReviewComments) for the Go
  language. If you are going to use a different language, please pick coding
  style guidelines and let us know what they are.
* Add a couple of high quality tests that cover happy and unhappy scenarios. Do
  not try to achieve full test coverage. This will take too long. Take two key
  components, e.g. authentication/authorization layer and networking and
  implement one or two test cases that demonstrate your approach to testing.
* Make sure builds are reproducible. Pick any vendoring/packaging system that
  will allow us to get consistent build results.
* Ensure error handling and error reporting is consistent. The system should
  report clear errors and not crash under non-critical conditions.
* Avoid concurrency and networking errors. Most of the issues we've seen in
  production are related to data races, networking error handling or goroutine
  leaks. We will be looking for those errors in your code.
* Security. Use strong authentication and simple but robust authorization.
  Set up the strongest transport encryption you can. Test it.

## Trade-offs

Write as little code as possible, otherwise this task will consume too much
time and code quality will suffer.

Please cut corners. For example, configuration tends to take a lot of time and
is not important for this task.

Use hardcoded values as much as possible and simply add TODO items showing your
thinking, for example:

```
  // TODO: Add configuration system.
  // Consider using CLI library to support both
  // environment variables and reasonable default values,
  // for example https://github.com/alecthomas/kingpin
```

Comments like this one are really helpful to us. They save you a lot of
time, demonstrate that you've spent time thinking about this problem, and
provide a clear path to a solution.

Consider making other reasonable trade-offs. Make sure you communicate them to
the interview team.

Here are some other trade-offs that will help you to spend less time on the
task:

* Do not implement a system that scales or is highly performing. Describe which
  performance improvements you would add in the future.
* It is okay if the system is not highly available. Write down how you would make
  the system highly available and why your system is not.

## Pitfalls and Gotchas

To help you out, we've composed a list of things that previously resulted in a
no-pass from the interview team:

* Scope creep. Candidates have tried to implement too much and they ran out of time
  and energy. To avoid this pitfall, use the simplest solution that will work.
  Avoid writing too much code. We've seen candidates' code introduce caching,
  which resulted in many mistakes in the caching layer validation logic. Not
  implementing a caching layer would have solved this problem.
* Data races. We will scan the code with a race detector and do our best to
  find data races in the code. Avoid global state as much as possible; if using
  global state, write down a good description why it is necessary and protect it
  against data races.
* Deadlocks. When using mutexes, channels or any other synchronization
  primitives, make sure the system won't deadlock. We've seen candidates' code
  holding a mutex and making a network call without timeouts in place. Be extra
  careful with networking and sync primitives.
* Unstructured code. We've seen candidates leaving commented chunks of code,
  having one large file with all the code, or not having code structure at all.
* Not communicating. Some candidates have submitted all their code to the
  master branch without raising pull requests, which does not give us the
  ability to provide feedback on the various implementation phases. We are a
  distributed team, so structured, asynchronous communication is critical to us.
* Implementing custom security algorithms/authentication schemes is always a
  bad idea unless you are a trained security researcher/engineer. It is
  definitely a bad idea for this task - try to stick to industry proven security
  methods as much as possible.

## Questions

It is okay to ask the interview team questions. Some folks stay away from asking
questions to avoid appearing less experienced, so we provide examples of
questions to ask and questions we expect candidates to figure out on their own.

Here is a great question to ask:

> Is it OK to pre-generate secret data and put the secrets in the repository
> for a proof of concept? I will add a note that we will auto-generate secrets
> in the future.

It demonstrates that you thought about this problem domain, recognize the trade
off and are saving you and the team time by not implementing it.

This is the question we expect candidates to figure out on their own:

> What version of Go should I use? What dependency manager should I use?

Unless specified in the requirements, pick the solution that works best for
you.

# Tools

This task should be implemented in Go or Rust and should work on 64-bit Linux
machines.

# Timing

You can split coding over a couple of weekdays or weekends and find time to ask
questions and receive feedback.

Once you join the Slack channel, you have 2 weeks to complete the challenge.

Within this timeframe, we don't give higher scores to challenges submitted more
quickly. We only evaluate the quality of the submission.

We only start the coding challenge if there are several open positions
available.
