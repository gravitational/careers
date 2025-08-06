# Summary

Implement a CLI tool that detects reconnaisance activity in a body of CloudTrail
logs.

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

# Levels

There are 6 engineering levels at Teleport. This challenge only supports L4.

Check [Systems Engineering Levels](../../levels/systems.pdf) for more details.

# Interview Process

The interview process will start with you receiving an invite to a private
Slack channel. That channel will contain the interview panel. You can ask them
about the engineering culture, work-life balance, or anything else that you
would like to learn about Teleport.

## Design Doc

Before writing any actual code, we ask that you write a brief design document.
The design document should cover: design approach, scope, proposed API, security
considerations, CLI UX, and implementation details where appropriate.

Include approaches you have evaluated and reasoning for picking the approach
you're planning to go with.

Please submit the design document and all code in a private GitHub repository.
Please submit the design document written in markdown as a Pull Request to allow
us to provide you feedback on the proposed design.

A few notes about the design document:

* Try to get the design document approved within the first week. This is to
  ensure you have enough time to work on the implementation.
* Avoid sending us draft design documents. It is difficult to evaluate which
  parts are draft and which parts are complete. Instead we encourage asking
  questions in Slack and sharing a design document that is ready to be reviewed.

Be sure to cover the following in your design:

* Your approach to solving the challenge problem.
* CLI user experience.

Once the design document has been approved by two reviewers, move on to the
implementation.

## Implementation

Split your code submission into roughly 2-3 Pull Requests to give the team an
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

### Tools

This task should be implemented in Python, Go or Rust and should work on 64-bit
Linux or macOS machines.

### Testing

Key components of the challenge should have tests that cover the happy and
unhappy scenarios. Do not try to achieve 100% test coverage as that will take
too long.

# Requirements

## Level 4

Use the public dataset of CloudTrail logs from flaws.cloud to implement a CLI
tool that is able to detect a cloud account reconnaissance activity using
anomaly detection and/or machine learning techniques.

https://summitroute.com/blog/2020/10/09/public_dataset_of_cloudtrail_logs_from_flaws_cloud/
https://summitroute.com/downloads/flaws_cloudtrail_logs.tar

Reconnaisance activity may include enumeration of AWS services and resources,
permissions probing and brute-force attempts, unusual API call frequencies.

After analyzing the logs, the tool should output the following information:

- Whether the suspicious activity was detected and the confidence level.
- The suspected identity.
- The time frame of the suspected recon activity.
- Key indicators (example logs/API calls).

Include a README documentation with the tool.

# Guidance

## Code and project ownership

This is a test challenge and we have no intent of using the code you've
submitted in production. This is your work, and you are free to do whatever you
feel is reasonable with it. In the scenario where you don't pass, you can open
source it with any license and use it as a portfolio project.

## Areas of focus

Teleport focuses on networking, infrastructure and security.

These are the areas we will be evaluating in the submission:

* Use consistent coding style.
* Tests exist for happy path and error scenarios for key components of the challenge.
* Make sure builds are reproducible.
* Ensure error handling and error reporting is consistent. The system should
  report clear errors and not crash under non-critical conditions.

The primary factor in the team's decision is overall code quality. We are looking for
the highest possible quality with the smallest possible scope that meets the requirements
of the challenge.

## Trade-offs

Cut corners, for example configuration tends to take a lot of time, and is not
important for this task.

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

* Persistency. Do not rely on any external databases.
* High-availability. This is out of scope for this challenge.

## Pitfalls and Gotchas

To help you out, we've composed a list of things that previously resulted in a
no-pass from the interview team:

* Scope creep. Candidates have tried to implement too much and ran out of time
  and energy. To avoid this pitfall, use the simplest solution that will work.
  Avoid writing too much code. We've seen candidates' code introducing caching
  and making many mistakes in the caching layer validation logic. Not having
  caching would have solved this problem.
* Unstructured code. We've seen candidates leaving commented chunks of code,
  having one large file with all the code, not having code structure at all.
* Not communicating. Some candidates have submitted all their code to the master
  branch without raising pull requests, which does not give us the ability to
  provide feedback on the various implementation phases. We are a distributed
  team, so structured, asynchronous communication is critical to us.

# Timing

You can split coding over a couple of weekdays or weekends and find time to ask
questions and receive feedback.

Once you join the Slack channel, you have 2 weeks complete the challenge.

Within this timeframe, we don't give higher scores to challenges submitted more
quickly. We only evaluate the quality of the submission.
