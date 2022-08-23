# Summary

We use a synchronous, remote, one-day hiring challenge for cloud
security engineers. The challenge consists of the following sections:

1. Kick off with the hiring manager. 45 minutes. 09:15 - 10:00 PT

2. Peer Review. 90 minutes. 10:00 - 11:30 PT. You'll receive a brief Request
for Discussion (RFD) and an architecture diagram of a service. You will
provide written PR feedback.

3. Break/Lunch. 30 minutes. 11:30 - 12:00 PT

4. Development. 3 hours. 12:00 - 15:00 PT. You'll receive prototype code for
the service you reviewed earlier. Building off this prototype codebase, we ask
you develop some changes and open a pull request.

5. Attack & Defense Discussion. 1 hour. 15:00 - 16:00 PT. A video call with
future peers where we ask you to show your depth of knowledge in hypothetically
attacking a service, and defending against those attacks. You will also have a
chance to ask us questions about working at Teleport.

Throughout the day, the interview team will join you in a Slack channel, where
we'll provide relevant links, and you're welcome to ask any questions.

# Rationale

These exercises have several goals:

* It helps us to understand what to expect from you as a developer, what
  security issues you find important, how you communicate those issues to peers,
  and how you write production code.
* It helps you get a feel for what it would be like to work at Teleport, as this
  exercise aims to simulate our day-as-usual and exposes you to the type of work
  and communication you can expect if you join Teleport.
* We aim to keep the process light. While there are time limits, it is to keep
  exercises from consuming undue amounts of your time and not to test
  performance under pressure.

We believe this technique is not only better, but is also more fun compared to
whiteboard/quiz interviews so common in the industry.

[Some of the best teams use coding challenges.](https://sockpuppet.org/blog/2015/03/06/the-hiring-post/)

We appreciate your time and are looking forward to the interview.

# Leveling

There are 6 engineering levels at Teleport. It's possible to earn level 4-5 via
this interview process.

We are not currently hiring L1-L3 security engineers, and level 6 is only
available via internal promotion. See our
[security engineering levels](https://raw.githubusercontent.com/gravitational/careers/main/levels/security.pdf)
for more details.

We expect different performance at different engineering levels. These
expectations are explicitly documented below.

# Part 1: Hiring Manager Kick Off
The hiring manager will discuss your background and approach to security with
you.

# Part 2: Peer Review
You will be given a GitHub repo containing a PR that includes a
[RFD](https://github.com/gravitational/teleport/blob/master/rfd/0000-rfds.md)
and architecture diagram for a cloud-hosted binary distribution service.

You will have 90 minutes of asynchronous time to read through the RFD and
provide written PR feedback in GitHub.

Please be sure to cover security concerns, though you're welcome to provide
input in any area where you see room for improvement.

## Level 4
Review the service as written, assuming it will run in a single availability
zone of a single region.

## Level 5
The service will run in multiple regions with unreliable network connections
between them. Please address the security and business trade-offs of different
consistency architectures.

# Part 3: Break
Grab a bite of food and some water. Take a breather before we continue to the
next section.

# Part 4: Development
You will be given a prototype AWS Terraform implementation of the RFD you
reviewed earlier. You will have 3 hours to develop and submit a single PR
securing the service. Please implement as many of the following security
features as you find time for. Implement the ones that you believe provide the
best security return on investment.

 - Tamper-resitant audit logs
 - Principle of least privilege, role-based IAM for:
   - Uploading new artifacts
   - Reading artifacts
   - Viewing audit logs
 - Encryption-at-rest
 - Encryption-in-transit
 - [Write Once Read Many](https://en.wikipedia.org/wiki/Write_once_read_many)
   or similar data retention protections to avoid data loss
 - Cloud provider specific security services such as:
   - AWS: CloudTrail, Config, GuardDuty, WAF

This is asynchronous. You'll develop on your machine, and you're welcome to
use whatever tools you're comfortable with to develop the changes. We'll be
available in the interview Slack channel to answer any questions.

We will review your PR, including running `terraform plan`. Make sure to write
good commit messages, as well as a helpful PR description.

# Part 5: Red Team & Blue Team discussion
After your development PR is posted, you'll join a synchronous video call with
future peers to finish the day. During this hour-long discussion, we'll present
you with a hypothetical service, and ask the following:

1. How would you attack this service, specifically to exfiltrate the core data?
2. How would you defend against the attacks you outlined earlier.

At the end of this discussion, you'll have an opportunity to ask peers
questions about working at Teleport.

# Guidance

## Interview Process

Although the interview is scheduled for a single day, the only in-person
discussions are the video calls at the beginning and end of the day. Both the
peer review and development portion are on your own machine without anyone
hovering over your shoulder--as we work in our day-to-day at Teleport.

Please make sure you have an environment set up for go development. The tools
needed can be found in the [Tools](#Tools) section below.

Should you have any questions during the day, please ask in the interview team
Slack channel.

After your interview, a panel of engineers will review the submitted
pull requests and vote using a "+1, -2" anonymous voting system: +1 is
submitted whenever a team member accepts the submission, -2 otherwise.

In case of a positive result, we will connect you to our HR and recruiting
teams, who will work out the details and present an offer.

In case of a negative score result, the hiring manager will contact you and
share a list of the key observations from the team that affected the result.

We're actively improving this interview process, and we would be glad to hear
your feedback about the experience.

## Code and Project Ownership

This is a contrived challenge and we have no intent of using the code you've
submitted in production. However, we retain all rights to
*our example code and RFD*. We ask that you do not publish our code to avoid
giving future candidates an unfair advantage by allowing them develop a solution
ahead of time.

## Areas of Focus

These are the areas we will be evaluating in the submission:

* Ensure the Terraform contains all resources needed.
* Use consistent coding style. We follow `terraform fmt` and
  [tfsec](https://github.com/aquasecurity/tfsec) for our Terraform.
* Follow the [principle of least privilege](https://en.wikipedia.org/wiki/Principle_of_least_privilege)
  when configuring IAM.
* Show your breadth and depth of knowledge for either AWS or GCP.

## Trade-offs

Write as little code as possible. Otherwise this task will take too much time
and PR quality will suffer.

Please cut corners. For example, cross-account IAM tends to take a lot of time
and is not important for this task.

Use hardcoded values as much as possible and simply add TODO items showing your
thinking, for example:

```
  // TODO: Configure AWS SSO role bindings
  // In a production environment, we'd want to get users from an IdP
  // but for demo purposes, we use IAM users.
```

Comments like this one are helpful to us. They save you time and demonstrate
that you've thought about this problem and can provide a clear path to a
solution.

Making other reasonable trade-offs. Please communicate them to the interview
team.

Here are some other trade-offs that will help you:

* Do not implement a system that has Multi-AZ or region-fault tolerance. Describe
  which improvements you would add in the future.

## Pitfalls and Gotchas

To help you out, we've composed a list of things that previously resulted in a
no-pass from the interview team:

* Scope creep. Candidates have tried to implement too much and ran out of time
  and energy. To avoid this pitfall, use the simplest solution that will work.
  Avoid writing too much code.
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

> Which Terraform providers should I use?

Unless specified in the requirements, pick the solution that works best for you.

# Tools

Your codebase should target 64-bit Linux or Mac machines.

Please have the following available in your development environment:

* [Terraform](https://www.terraform.io/downloads) 1.2+
