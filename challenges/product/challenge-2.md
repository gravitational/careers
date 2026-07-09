# Summary

Interview process for a Director of Product, Data Platform that selects for the
skills needed to build data pipelines and a data warehouse.

## Rationale

This exercise has two goals:

It helps us understand what to expect from you as a technical leader and
whether you have the technical aptitude to engage with highly technical
customers. It gives you a feel for what it's like to work at Teleport. The
exercise simulates day-to-day work, helping you better understand the team and
our hiring process.

We believe this approach is not only more effective, but also more enjoyable
than the whiteboard or quiz-style interviews common in the industry.

We appreciate your time and look forward to hacking on this little project
together.

## Objective

Design and build a data warehouse for a fictional business.

You will receive raw data about the business and must build a data pipeline to process
and import the data into a data warehouse. Then you'll use the data to
write a 1-page document that provides insight and suggests changes to that
fictional business.

## Requirements

* Determine implementation, quality, security, and scope requirements to write
  a technical design document
* Build data pipelines and data warehouse using the sample data we provide to you
* Write a business document that outlines the insights and changes to the
  fictional business you are proposing
* Explain the business and technical value of the proposed solution on a short
  call

Write your data pipeline scripts in Go or Python. Use DuckDB as your data
warehouse.

## Guidance

### Interview Process

You will join a Slack channel with the interview team that consists of
peers who will be working with you.

Ask them about the culture, work and life balance, or anything else that you
would like to learn about Teleport.

Before writing the actual code, create a brief design document in Markdown and
share with the team via a GitHub Pull Request.

There is no required format or word count for the design document. The design
should convey your plan to satisfy the requirements, possible alternatives, and
the tradeoffs informing your approach. The team will review your design and
provide feedback by commenting on the pull request.

After the team has approved your design document, you may begin submitting pull
requests with the implementation.

After the implementation is complete, prepare a presentation and demo, then
schedule a 45 minute call with the interview team. During this call you will
present your solution to the interview team and answer any questions.

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

* Reproducible pipelines. Scripts should be written in a way to allow
  reproduction of environment.
* Data models. Design performant tables. Keep queries fast and plan for schema
  migration.
* Data quality. Add validation and test coverage to all stages of your pipeline
* Security. Data should be processed and stored in a way not to leak sensitive
  data.
* Demo. Product leaders have to communicate complex technical deployments in
  simple terms. The demo should be interesting to watch and have a good
  delivery.
* Communication. The technical and business documents should be simple to
  understand and communicate key insights to someone who sees it the first
  time.

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

> Is it OK for everything to run locally? I would build a real data warehouse
> using something like Redshift or Snowflake, but for this challenge everything
> will be driven by a Makefile and run locally.

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

We only start the challenge if there is a positions available and let all
candidates finish the submission.

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
