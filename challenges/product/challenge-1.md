# Product Challenge

## Rationale

This exercise has two goals:

It helps us understand what to expect from you as a technical leader, and
whether you have the technical aptitude to engage with highly technical
customers. It gives you a feel for what it's like to work at Teleport. The
exercise simulates day-to-day work, helping you better understand the team and
our hiring process.

We believe this approach is not only more effective, but also more enjoyable
than the whiteboard or quiz-style interviews common in the industry.

We appreciate your time and look forward to hacking on this little project
together.

## Objective

Create a design document and implement a service that will let users sign on
through SSO IDP of of any kind in order to access a WireGuard network.

After authenticating, the client and server exchange keys and configure Wireguard.
The client should then be able to access a service that is only available on the
server's WireGuard interface (nginx or any other web service will do).

## Requirements

- Determine implementation, security and scope requirements for the document.
- Present the design to the interview team - a mix of technical and business audience,
  explain value proposition and answer technical questions to both.
- Build the service. Your solution should include automation using Terraform, helm charts,
  or docker-compose that sets up a demo environment.
- Explain the business and technical value of the proposed solution.

Here is an example command-line tool that will let users sign in with SSO. After
a successful login users get a WireGuard key registered in the provisioned
WireGuard network.

```
$ vpn login
# ...opens browser with login page
# ...on successful login
You are logged in. WireGuard is activated.
```

## Guidance

### Interview Process

You will join a Slack channel with the interview team. The team consists of the
peers who will be working with you.

Ask them about the culture, work and life balance, or anything else that you
would like to learn about Teleport.

Before writing the actual code, create a brief design document in markdown and
share with the team via a GitHub pull request.

This document should consist of key trade-offs and key design approaches. Please
avoid writing an overly detailed design document. Use this document to make sure
demonstrate that you've investigated the problem space. The team will review your
design and provide feedback by commenting on the pull request.

After the team has approved your design document, you may begin submitting pull
requests with the implementation.

After the implementation is complete, prepare a presentation and demo and
schedule a 45 minute call with the interview team. During this call you will
present your solution to the audience and answer any questions.

After the demo, the interview team will assemble and vote using +1, -2 anonymous
voting system: +1 is submitted whenever a team member accepts the submission, -2
otherwise.

In case of a positive result, we will connect you to our HR team who will
collect one-two references and will work out other details. You can start the
reference collection process in parallel if you would like to speed up the
process.

After reference collection, our recruiting team will send you an offer.

In case of a negative score result, the hiring manager will contact you and
share a list of the key observations from the team that affected the result.

### Code and project ownership

This is a test challenge and we have no intent of using the design you've
submitted in production. This is your work, and you are free to do whatever you
feel is reasonable with it. In the scenario when you don't pass, you can open
source it with any license and use it as a portfolio project.

### Areas of focus

Teleport focuses on networking, infrastructure and security.

- Reproducible builds. The team should be able to apply the configuration and reproduce your example easily.
- Demo. Product leaders have to communicate complex technical deployments in simple terms. The demo should be interesting to watch and have a good delivery.
- Security. Design the login flow that exchanges username/password for a wire guard keypair with security best practices in mind.  Consider common attack vectors, such as man-in-the-middle and replay attacks.
- User Experience. The solution should be simple to understand and use to someone who sees it the first time.

### Pitfalls & Gotchas

To help you prepare, here are some common reasons candidates have failed to pass
our interviews:

* *Use of AI.* Don't outsource your thinking to an AI. We recommend using AI for
  use cases like learning about a new problem space, exploring APIs, and
  finding missing edge cases. However, we strongly recommend you write the design
  document and all code yourself.
* *Jumping into implementation without clarifying requirements or narrowing the
  scope.* We expect product leaders to conduct customer interviews, negotiate
  scope, and identify the right MVP. We want you to take the same approach
  during the design phase.
* *Scope creep.* Candidates have tried to design too much and ran out of time
  and energy.  To avoid this pitfall, use the simplest solution that will work.
* *Suggesting custom security algorithms/authentication schemes* is always a
  bad idea unless you are a trained security researcher/engineer. It is
  definitely a bad idea for this task - try to stick to industry proven security
  methods as much as possible.

### Questions

It is OK to ask the interview team questions. Some folks stay away from asking
questions to avoid appearing less experienced, so we provide examples of
questions to ask and questions we expect candidates to figure out on their own.

Here is a great question to ask:

> Is it OK to pre-generate secret data and put the secrets in the repository for
> a proof of concept? I will add a note that we will auto-generate secrets in the
> future.

It demonstrates that you thought about this problem domain, recognized the trade
off and are saving you and the team time by not implementing it.

This is the question we expect candidates to figure out on their own:

> What version of Go should I use? What dependency manager should I use?

Unless specified in the requirements, pick the version that works best for you.

### Timing

It should take you from 4 and no more than 24 full hours to complete the
challenge. You can split coding over a couple of weekdays or weekends and find
time to ask questions and receive feedback.

Once you join the Slack channel, you have a maximum of 2 weeks to complete the
challenge.

Within this timeframe, we don't give higher scores to challenges submitted more
quickly. We only evaluate the quality of the submission.

We only start the challenge if there are several open positions available and let all candidates finish the submission.

We always aim to provide 1-2 rounds of feedback on all work that is submitted.
In order to be respectful of your time, we may opt to end the challenge early
if the submission does not improve after this feedback is suggested or if we
identify a large number of issues.

### Setup

Create a GitHub repository, either throwaway or in your GitHub account, feel
free to open source it and use it later.

We understand that you will spend a couple of days on this project, so we know
that it will take you 1-2 weeks to reserve the time for it.

Do not hesitate to reach out in case you get stuck or have any kind of general
questions or concerns. Please remember that communication is just as important
for this exercise as the code.
