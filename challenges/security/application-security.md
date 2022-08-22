# Summary

We use a synchronous, remote, one day hiring challenge for application
security engineers. The challenge consists of the following sections:

1. Kick off with the hiring manager. 45 minutes. 09:15 - 10:00 PST

2. Peer review. You'll receive a brief Request for Discussion (RFD) and
architecture diagram of a service. You will have 90 minutes to
provide written PR feedback. 10:00 - 11:30 PST.

3. After the peer review, section, there is a 30 minute break for lunch.

4. Development. You'll receive prototype code for the service reviewed
in part one. Building off this prototype codebase, we ask you develop
some changes and open a pull request over 3 hours. 12:00 - 15:00 PST

5. Attack & defense discussion.  This is a 60 minute video call with future
peers where we ask you to show your depth of knowledge in hypothetically
attacking a service, and defending against those attacks.
You will also have a chance to ask us questions about working at Teleport.
15:00 - 16:00 PST.

Throughout the day, the interview team will join you in a slack channel, where
we'll provide relevant links, and you're welcome to ask any questions.

# Rationale

These exercises have several goals:

* It helps us to understand what to expect from you as a developer, what
  security issues you find important, how you communicate those issues to peers
  and how you write production code.
* It helps you get a feel for what it would be like to work at Teleport, as this
  exercise aims to simulate our day-as-usual and expose you to the type of work
  and communication you can expect if you join Teleport.
* We aim to keep the process light. While there are time limits, it is to keep
  exercises from consuming undue amounts of your time, and not to test
  performance under pressure.

We believe this technique is not only better, but also is more fun compared to
whiteboard/quiz interviews so common in the industry.

[Some of the best teams use coding challenges.](https://sockpuppet.org/blog/2015/03/06/the-hiring-post/)

We appreciate your time and are looking forward to the interview.

# Leveling
There are 6 engineering levels at Teleport. It's possible to earn level 4-5 via
this interview process.

We are not currently hiring L1-L3 security engineers and level 6 is only
available via internal promotion. See our
[security engineering levels](https://github.com/gravitational/careers/blob/main/levels/security.pdf)
for more details.

We expect different performance at different engineering levels. These
expectations are explicitly documented below.

# Part 1: Hiring Manager Kick Off
The hiring manager will discuss your background and approach to security with
you.


# Part 2: Peer Review
You will be given a GitHub repo containing a PR that includes a
[RFD](https://github.com/gravitational/teleport/blob/master/rfd/0000-rfds.md)
and architecture diagram for a service similar to docker.

You will have 90 minutes of asynchronous time to read through the RFD and
provide written PR feedback in GitHub.

Please be sure to cover security concerns, though you're welcome to provide
input in any area where you see room for improvement.

## Level 4
Review the service as written, assuming it will run on a single server.

## Level 5
The service will run in multiple geographies with unreliable network connections
between them. Please address the security and business tradeoffs between
authentication and authorization consistency architectures.

# Part 3: Break
Grab a bite of food and some water. Take a breather before we continue to the
next section.

# Part 4: Development
You will be given a prototype go implementation of the RFD reviewed in
part 1. You will have 3 hours to develop and submit a single PR implementing
an authorization and authentication scheme.

This is asynchronous. You'll develop on your machine and you're welcome to
use whatever tools you're comfortable with to develop the changes. We'll be
available in the interview slack channel to answer any questions.

We will review your PR, including compiling and run the program and testing
it behavior. Make sure to write good commit messages, as well as a helpful
PR description.

This prototype is expected to run on only a single machine. The distributed
state considerations of Part 2 L5 need not be considered.

## Level 4
Add authentication and authorization to the Library, API, and client codebase.

Separate actors should not be able to view or interact with other actors jobs.

Log all actions and the actor performing them.

## Level 5

Add authentication and authorization to the Library, API, and client codebase.

Separate users should not be able to view or interact with other user's jobs.

Log all actions and the actor performing them.

Certain actors with an "admin" permission should be able to impersonate
another user.  In this case, log both the impersonator and the account being
impersonated.

# Part 5: Red Team & Blue Team discussion
After your development PR is posted, you'll join a synchronous video call with
future peers to finish the day. During this hour long discussion, we'll present
you with a hypothetical service, and ask the following:

1. How would you attack this service, specifically to exfiltrate the core data?
2. How would you defend against the attacks you outlined earlier.

At the end of this discussion, you'll have an opportunity to ask peers
questions about working at Teleport.

# Guidance

## Interview process

Although the interview is scheduled for a single day, the only in-person
discussions are the video calls at the beginning and end of the day. Both the
peer review and development portion are on your own machine without anyone
hovering over your shoulder--as we work in our day-to-day at Teleport.

Please make sure you have an environment set up for go development. The tools
needed can be found in the [Tools](#Tools) section below.

Should you have any questions during the day, please ask the interview team
Slack channel.

After your interview, a panel of engineers will review the submitted
pull requests and vote using a "+1, -2" anonymous voting system: +1 is
submitted whenever a team member accepts the submission, -2 otherwise.

In case of a positive result, we will connect you to our HR and recruiting
teams, who will work out the details and present an offer.

In case of a negative score result, hiring manager will contact you and share a
list of the key observations from the team that affected the result.

We're actively improving this interview process, and we would be glad to hear
your feedback about the experience.

## Code and project ownership

This is a contrived challenge based off Teleport's
[Job Worker challenge](https://github.com/gravitational/careers/blob/a660d1bf4c327ec01e0c03d903eecd29522b7c6e/challenges/systems/challenge-1.md),
and we have no intent of using the code you've submitted in production. However,
we retain all rights to *our example code and RFD*. We ask that you do not
publish our code to avoid giving future candidates an unfair advantage by
allowing them develop a solution ahead of time.

## Areas of focus

These are the areas we will be evaluating in the submission:

* Use consistent coding style. We follow
  [Go Coding Style](https://github.com/golang/go/wiki/CodeReviewComments) for
  the Go language.
* Ensure error handling and error reporting is consistent. The system should
  report clear errors and not crash under non-critical conditions.
* Avoid concurrency and networking errors. Most of the issues we've seen in
  production are related to data races, networking error handling or goroutine
  leaks. We will be looking for those errors in your code.
* Security. Use strong authentication and robust authorization.
  Set up the strongest transport encryption you can.

## Trade-offs

Write as little code as possible, otherwise this task will take too much time
and PR quality will suffer.

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

Comments like this one are helpful to us. They save you time and demonstrate
that you've thought about this problem and can provide a clear path to a
solution.

Making other reasonable trade-offs. Please communicate them to the interview
team.

Here are some other trade-offs that will help you:

* Do not implement a system that scales or is highly performing. Note which
  performance improvements you would add in the future.
* It is OK if the system is not highly available. Write down how you would make
  the system highly available and why your system is not.

## Pitfalls and Gotchas

To help you out, we've composed a list of things that previously resulted in a
no-pass from the interview team:

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

## Questions

It is encouraged to ask the interview team questions. Some folks stay away from
asking questions to avoid appearing less experienced, so we provide examples of
questions to ask and questions we expect candidates to figure out on their own.

Here is a great question to ask:

> Is it OK to pre-generate secret data and put the secrets in the repository for
> a proof of concept? I will add a note that we will auto-generate secrets in
> the future.

It demonstrates that you thought about this problem domain, recognize the trade
off and are saving you and the team time by not implementing it.

This is the question we expect candidates to figure out on their own:

> Which logging library should I use?

Unless specified in the requirements, pick the solution that works best for you.

# Tools

The prototype codebase will be provided in go and the final version should build
and run on 64-bit Linux machines.

Please have the following available in your development environment:

* [Go](https://go.dev/doc/install) 1.18+
* Make
* [protoc](https://grpc.io/docs/protoc-installation/) v3.13+
* [protoc-gen-go](https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers) v1.28+
