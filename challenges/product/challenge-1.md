# Product Challenge

## Rationale

This exercise has two goals:

- It helps us understand what to expect from you as a technical leader and
  whether you have the technical aptitude to engage with highly technical
  customers.
- It gives you a feel for what it's like to work at Teleport. The exercise
  simulates day-to-day work, helping you better understand the team and our
  hiring process.

We believe this approach is not only more effective but also more enjoyable
than the whiteboard or quiz-style interviews common in the industry.

We appreciate your time and look forward to hacking on this project together.

## Objective

Create a design document and implement a service that lets users authenticate
through an SSO Identity Provider and then access a
protected resource over a WireGuard VPN.

After authenticating, the client and server exchange WireGuard keys, configure
the tunnel, and the client should be able to reach a service that
is only available on the server's WireGuard network, not on the public internet.

## Requirements

### Design Document

Before writing code, produce a concise design document (Markdown) and submit it
as a GitHub pull request for the team to review. The document should cover:

- **Problem statement & success criteria.** Why does this product exist? What
  user problem does it solve? How will you know the solution works?
- **User journey.** Think through the end-to-end experience from the user's
  perspective.
- **API design.** Define the client-gateway interface.
- **Key management.** Explain your strategy for WireGuard key lifecycle and
  what trade-offs it introduces for security and user experience.
- **Security model.** Identify your threat model and how your design addresses
  common attack vectors.
- **Scope & trade-offs.** State what is in and out of scope for your MVP.

Here is an example command-line tool that will let users sign in with SSO. After
a successful login users get a WireGuard key registered in the provisioned
WireGuard network.

```
$ vpn login
# ...opens browser with login page
# ...on successful login
You are logged in. WireGuard is activated.
```

### Implementation

- **Service.** Build the service and include automation that sets up a fully
  working demo environment.
- **Reproducibility.** A reviewer should be able to clone, build, and run the
  demo with minimal steps.
- **Presentation.** Present the design and demo to the interview team (45 min)
  - explain the value proposition and answer technical questions from both
  technical and business audience members.

### Evaluation Criteria

We evaluate submissions holistically across these dimensions:

- **Product thinking.** Clear problem definition, user empathy, defined success criteria.
- **CLI UX.** Precise, intuitive commands with helpful output and clear error messages.
- **API design.** Well-defined contract with proper status codes, error handling, versioning.
- **Key management.** Justified design with clear rationale for lifecycle, storage, and threat model.
- **Security.** Correct use of proven protocols, defense in depth, no custom crypto.
- **Error handling.** Graceful failures, actionable messages, no silent drops.
- **Architecture.** Clean separation of concerns, reproducible builds, automation.
- **Communication.** Concise design doc, clear demo delivery, good Q&A.

## Guidance

### Interview Process

You will join a Slack channel with the interview team - the peers who will be
working with you. Feel free to ask them about culture, work-life balance, or
anything else you'd like to know about Teleport.

1. Share your design document via a GitHub pull request. Keep it focused on key
   trade-offs and design approaches - avoid excessive detail. The team will
   review and comment on the PR.
2. Once the design is approved, submit pull requests with the implementation.
3. When the implementation is complete, schedule a 45-minute call with the
   interview team for your presentation and demo.

After the demo the team votes using a +1 / -2 anonymous system. On a positive
result, our HR team will collect one or two references and extend an offer. On a
negative result, the hiring manager will share the key observations that
affected the outcome. You may start the reference collection process in parallel
to speed things up.

### Timing

The challenge should take between 4 and 24 hours of focused work. You can split
it over weekdays or weekends and use the time to ask questions and receive
feedback.

Once you join the Slack channel you have a maximum of 2 weeks to complete the
challenge. We do not give higher scores to faster submissions - we only evaluate
quality. We aim to provide 1–2 rounds of feedback on all work. To be respectful
of your time, we may end the challenge early if the submission does not improve
after feedback or if we identify a large number of issues.

### Setup

Create a GitHub repository - either a throwaway or in your personal account.
Feel free to open-source it and use it as a portfolio project later.

Do not hesitate to reach out if you get stuck or have questions. Communication
is just as important for this exercise as the code.

### Areas of Focus

Teleport focuses on networking, infrastructure, and security. Beyond the
requirements above, keep the following in mind:

- **Demo delivery.** Product leaders communicate complex deployments in simple
  terms. The demo should be clear, concise, and interesting to watch.
- **Security justification.** Be prepared to justify your design choices in
  terms of specific threats they address.
- **Design-first approach.** The design doc should make the user journey and
  API contract obvious before any code is written.

### Pitfalls

To help you prepare, here are common reasons candidates have not passed:

- **Use of AI.** Don't outsource your thinking to an AI. We recommend using AI
  for learning about a new problem space, exploring APIs, and finding edge
  cases. However, write the design document and all code yourself.
- **Jumping into code without a design.** We expect you to clarify
  requirements, negotiate scope, and think through the user experience and API
  before writing implementation code.
- **Scope creep.** Candidates have tried to design too much and ran out of
  time. Use the simplest solution that works, and clearly state what is out of
  scope.
- **Custom crypto.** Never invent security algorithms or authentication
  schemes. Stick to industry-proven protocols.

### Questions

It is OK to ask the interview team questions. Some folks stay away from asking
questions to avoid appearing less experienced, so here are examples.

A great question to ask:

> Is it OK to pre-generate secret data and put the secrets in the repository for
> a proof of concept? I will add a note that we will auto-generate secrets in the
> future.

It demonstrates that you thought about the problem domain, recognized the
trade-off, and are saving time by not implementing it yet.

A question we expect candidates to figure out on their own:

> What version of Go should I use? What dependency manager should I use?

Unless specified in the requirements, pick the version that works best for you.

### Code and Project Ownership

This is a test challenge and we have no intent of using your submission in
production. This is your work, and you are free to do whatever you wish with it.
If you don't pass, you can open-source it with any license and use it as a
portfolio project.