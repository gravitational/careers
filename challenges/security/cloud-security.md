# Summary

We use a two part hiring challenge for cloud security engineers:

Part One: Peer review. You'll receive a Request for Discussion (RFD) and
architecture diagram for a service similar to docker. You have two hours to provide
written PR feedback.

Part Two: Development. You'll receive prototype code for the service reviewed
in part one. Implement either AWS or GCP Terraform to securely distribute and
audit artifacts generated from this codebase.

# Rationale

These exercises have several goals:

* It helps us to understand what to expect from you as a developer, what security
  issues you find important, how you communicate those issues to peers and how you
  write production code.
* It helps you get a feel for what it would be like to work at Teleport, as this
  exercise aims to simulate our day-as-usual and expose you to the type of work
  we're doing here.
* We aim to keep the process light. When there are time limits, it is to keep
  exercises from consuming undue amounts of your time, and not to test performance
  under pressure.

We believe this technique is not only better, but also is more fun compared to
whiteboard/quiz interviews so common in the industry.

[Some of the best teams use coding challenges.](https://sockpuppet.org/blog/2015/03/06/the-hiring-post/)

We appreciate your time and are looking forward to hack on this project together.

# Leveling
There are 6 engineering levels at Teleport. It's possible to earn level 4-5 via
this interview process.

We are not currently hiring L1-L3 security engineers and level 6 is only available
via internal promotion. See the [engineering levels document](https://raw.githubusercontent.com/gravitational/careers/main/levels.pdf)
for more details.

We expect different performance at different engineering levels. These expectations
are explicitly documented below.


# Part 1: Peer Review
You will be given a GitHub repo containing a PR that includes a RFD and architecture
diagram for a service similar to docker.

You will have two hours to read through the RFD and provide feedback. Please
focus on security concerns, though you're welcome to provide input in any area
where you see room for improvement.

## Level 4
Review the service as written, assuming it will run on a single server.

## Level 5
The service will run in multiple geographies with unreliable network connections
between them. Please address the security and business tradeoffs between
authentication and authorization consistency architectures.


# Part 2: Development
You will be given a prototype implementation of the RFD reviewed in part 1.

Write a single PR implementing AWS or GCP Terraform that securely distributes
and provides an audit trail for artifacts generated from this codebase.

We will review the pull request and leave feedback, and provide you a chance to
respond to or incorporate the feedback.

After PR discussion, we will apply the terraform in a sandboxed account, test it
and get back to you.

## Level 4

A Make target should be able to store built artifacts using a least-privilege
service account.

Authenticated users should be able to download any of these blobs, but not change
any data.

Configure your cloud provider to log artifact any artifact write and retain these
logs for at least 90 days.

## Level 5

TBD.


# Guidance

## Interview process

The interview team will join the Slack channel. The team consists of the
engineers who will be working with you. Ask them about the engineering culture,
work and life balance, or anything else that you would like to learn about
at Teleport.

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

These are the areas we will be evaluating in the submission:

* Ensure the Terraform applies cleanly from scratch and contains
  all resources needed.
* Use consistent coding style. We follow `terraform fmt` and
  [tfsec](https://github.com/aquasecurity/tfsec) for our Terraform.
* Follow the [principle of least privilege](https://en.wikipedia.org/wiki/Principle_of_least_privilege)
  when configuring IAMs.
* Show your breadth and depth of knowledge for either AWS or GCP.

## Trade-offs

Write as little code as possible, otherwise this task will consume too much time
and code quality will suffer.

Please cut corners, for example cross account IAMs tend to take a lot of time, and
are not important for this task.

Use hardcoded values as much as possible and simply add TODO items showing your
thinking, for example:

```
  // TODO: Configure AWS SSO role bindings
  // In a production environment we'd want to get users from an IDP
  // but for demo purposes we use IAM users.
```

Comments like this one are really helpful to us. They save yourself a lot of
time and demonstrate that you've spent time thinking about this problem and
provide a clear path to a solution.

Consider making other reasonable trade-offs. Make sure you communicate them to
the interview team.

Here are some other trade-offs that will help you to spend less time on the task:

* Do not implement a system that has multi-geography fault tolerance. Describe
  which improvements you would add in the future.
* 

## Pitfalls and Gotchas

To help you out, we've composed a list of things that previously resulted in a
no-pass from the interview team:

* Scope creep. Candidates have tried to implement too much and ran out of time
  and energy. To avoid this pitfall, use the simplest solution that will work.
  Avoid writing too much code.
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

> Is it OK to pre-generate secret data and put the secrets in the repository for
> a proof of concept? I will add a note that we will auto-generate secrets in
> the future.

It demonstrates that you thought about this problem domain, recognize the trade
off and are saving you and the team time by not implementing it.

This is the question we expect candidates to figure out on their own:

> Which Terraform providers should I use?

Unless specified in the requirements, pick the solution that works best for you.

# Tools

Use Terraform 1.0+. The final version should build and run on 64-bit Mac or
Linux machines.

# Timing

You can split coding over a couple of weekdays or weekend and find time to ask
questions and receive feedback.

Once you join the Slack channel, you have 1 week complete the challenge.

Within this timeframe, we don't give higher scores to challenges submitted more
quickly. We only evaluate the quality of the submission.
