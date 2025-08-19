# Summary

Build an MVP for workflow automation that can securely configure, add web
applications, and users to an Auth0 tenant.

# Rationale

This exercise has two goals:

* It helps us to understand what to expect from you as a Security and Automation
  Engineer, how you securely configure infrastructure and write code
  and scripts.
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

There are 6 levels for Security and Automation at Teleport. We are only
interviewing for P3 and P4 at the moment.

# Interview Process

The interview process will start with you receiving an invite to a private Slack
channel. That channel will contain the interview panel. You can ask the panel
about company culture, work-life balance, or anything else that you would like
to learn about Teleport.

## Design Doc

Before working on the implementation, we ask that you write a brief design
document. The design document should cover: scope, infrastructure
configuration, APIs you will use, security considerations, edge cases, and
implementation details where appropriate.

At Teleport, we prefer Markdown for [our
designs](https://github.com/gravitational/teleport/blob/master/rfd/0000-rfds.md).

Please submit the design document and all code in a GitHub repository. Public
or private is your choice. Please submit the design document as a Pull Request
(PR) to allow us to provide you feedback on the proposed design.

A few notes about the design document:

* Try to get the design document approved within the first 2-3 days. This is to
  ensure you have enough time to work on the implementation.
* Avoid writing an overly detailed design document. One to two pages is
  sufficient.
* Avoid sending us draft design documents. It is difficult to evaluate which
  parts are draft and which parts are complete. Instead we encourage asking
  questions in Slack and sharing a design document that is ready to be reviewed.

Once the design document has been approved by two reviewers, move on to the
implementation.

## Implementation

Split your submission into roughly 3-5 Pull Requests to give the team an
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

### Dependencies

Please write as much of your own code as possible. Avoid relying on non-first
party dependencies. When needed, use reputable sources for key components of the
challenge.

# Requirements

## Level 3

Use [GitHub Actions](https://github.com/features/actions) to build MVP workflows
that can do the following.

* Add a workflow that triggers [Terraform](https://www.terraform.io) to
  configure an Auth0 tenant. Enforce strong authentication for all web
  applications within tenant.
* Add a workflow that triggers Terraform to add a web application to the Auth0
  tenant. Use [OSS Grafana](https://grafana.com/oss) with OIDC authentication
  as the sample application.
* Add a workflow that executes a shell script (or Go program) to register
  users within the Auth0 tenant. Add a sample user to the Auth0 tenant.
* Securely store and scope all credentials used.

Link to and share successful workflow runs after each PR has been approved and
merged.

## Level 4

Use [Temporal](https://temporal.io) to build MVP workflows that can do the
following. Use Go (or Python) Temporal SDK to the implement workflows.

* Create a [Docker Compose](https://docs.docker.com/compose) stack that will
  be used to run Temporal.
* Create a CLI tool that can be used to trigger Temporal workflows.
* Add a workflow that triggers [Terraform](https://www.terraform.io) to
  configure an Auth0 tenant. Enforce strong authentication for all web
  applications within tenant.
* Add a workflow that adds a web application to the Auth0 tenant. Use OSS
  Grafana with OIDC authentication as the sample application.
* Add a workflow that registers users within the Auth0 tenant. Add a sample user
  to the Auth0 tenant.
* Ensure workflows are covered by end-to-end tests.
* Securely store and scope all credentials used.
* Include README file explaining how to try out the project.

# Guidance

## Code and project ownership

This is a test challenge and we have no intent of using anything you have
written in production. This is your work, and you are free to do whatever you
feel is reasonable with it. In the scenario where you don't pass, you can open
source it with any license and use it as a portfolio project.

## Areas of focus

Teleport focuses on networking, infrastructure, and security.

These are the areas we will be evaluating in the submission:

* Secure configuration of infrastructure.
* High quality code and scripts written in a consistent coding style.

The primary factor in the team's decision is submission quality. We are looking
for the highest possible quality with the smallest possible scope that meets
the requirements of the challenge.

## Trade-offs

Don't enforce every security feature available. Think about your attack vectors
and the minimal set of features that provide the highest security benefits.
Write as little code as possible, otherwise this task will consume too much time
and submission quality will suffer.

Please cut corners where appropriate, for example, you don't need to persist
Terraform state for this task. Simply add a `TODO` item to show your thinking.
For example:

```bash
# TODO: Terraform state is wiped out after each run. In a production environment
# state would be persisted across multiple runs by storing it in a S3 bucket
# with locked down access.
```

Comments like this one are really helpful. They save you a lot of time and
demonstrate that you've spent time thinking about this problem and provide a
clear path to a solution.

Another corner to cut is using hardcoded values where appropriate, don't overly
abstract your code.

Make reasonable trade-offs and communicate them to the interview team.

## Pitfalls and Gotchas

To help you out, we've composed a list of things that previously resulted in a
no-pass from the interview team:

* Scope creep. Candidates have tried to implement too much and ran out of time
  and energy. To avoid this pitfall, use the simplest solution that will work.
* Buggy submissions. You are writing a MVP, but that does not mean your
  submission should be poor quality. Put your best foot forward.
* Unstructured submissions. We've seen candidates leaving in commented chunks of
  code, having one large file with all the code, not having code structure at
  all.
* Not communicating. Some candidates have submitted all their code to the
  `master` branch without raising pull requests, which does not give us the
  ability to provide feedback on the various implementation phases. We are a
  distributed team, so structured, asynchronous communication is critical to us.
* Implementing custom security algorithms/authentication schemes is always a
  bad idea unless you are a trained security researcher/engineer. It is
  definitely a bad idea for this task - try to stick to industry proven security
  methods as much as possible.

## Questions

It is OK to ask the interview team questions. Some people stay away from asking
questions to avoid appearing less experienced, so we provide examples of
questions to ask and questions we expect candidates to figure out on their own.

Here is a great question to ask:

> Is it OK to enforce strong unphishable multi-factor authentication (like
> WebAuthn) even if it requires the purchase of additional hardware?

It demonstrates that you understand the problem domain (phishing risks of
passwords and one-time passwords), have a solution (unphishable authentication),
and the trade-offs (additional costs). Our answer would be: do it.

This is the question we expect candidates to figure out on their own:

> What version of the Auth0 Terraform Provider should I use? What container
> image should I use?

Unless specified in the requirements, pick the solution that works best for
you.

# Tools

This task should be implemented using GitHub Actions or Temporal/Docker,
Terraform, and shell scripts or Go/Python.

# Timing

You can split coding over a couple of weekdays or weekends and find time to ask
questions and receive feedback.

Once you join the Slack channel, you have 2 weeks complete the challenge.

Within this timeframe, we don't give higher scores to challenges submitted more
quickly. We only evaluate the quality of the submission.
