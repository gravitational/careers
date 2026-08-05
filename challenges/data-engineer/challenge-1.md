# Summary

Design and build a data pipeline and warehouse for a fictional business.

# Rationale

This exercise has two goals:

* It helps us understand what to expect from you as a developer, how you write
  production code, how you reason about data model design, and how you
  communicate when trying to understand a problem before you solve it.
* It helps you get a feel for what it would be like to work at Teleport, as
  this exercise aims to simulate our day-as-usual and expose you to the type
  of work we're doing here.

We believe this technique is not only better, but also is more fun compared to
whiteboard/quiz interviews so common in the industry. It's not without the
downsides - it could take longer than traditional interviews.

[Some of the best teams use coding challenges.](https://sockpuppet.org/blog/2015/03/06/the-hiring-post/)

We appreciate your time and are looking forward to hack on this project
together.

# Levels

There are 6 engineering levels at Teleport. This challenge supports L3-L5.

Level 6 is only for internal promotions. Check [Engineering
Levels](../../levels/systems.pdf) for more details.

# Interview Process

The interview process will start with you receiving an invite to a private
Slack channel. That channel will contain the interview panel. You can ask them
about the engineering culture, work-life balance, or anything else that you
would like to learn about Teleport.

## Scenario

You will receive raw data about a fictional business, including records like
employee/HR exports, sales transactions, and pricing information. You must
build a data pipeline that processes and imports this data into a data
warehouse, modeling it so that it can be queried and analyzed.

The company sells a small catalog of widget SKUs. Sales reps are credited for
the deals they close, and finance wants to understand product profitability
and rep performance on a monthly basis.

## Design Doc

Before writing any actual code, we ask that you write a brief design document.
The design document should cover: your data model, pipeline architecture,
where each transformation happens and why, your approach to data validation
and PII handling, and implementation details where appropriate. If you are
targeting Level 4 or 5, also cover your plan for ingesting data that arrives
in more than one batch over time. If you are targeting Level 5, also cover
how you'd reconcile a later batch that isn't purely additive, including any
schema changes.

Include approaches you have evaluated and reasoning for picking the approach
you're planning to go with.

Please submit the design document and all code in a GitHub repository. Public
or private is your choice. Please submit the design document written in
markdown as a Pull Request to allow us to provide you feedback on the proposed
design.

A few notes about the design document:

* Try to get the design document approved within the first 2-3 days. This is
  to ensure you have enough time to work on the implementation.
* There is no required format or word count for the design document. The
  design should convey your plan to satisfy the requirements, possible
  alternatives, and the tradeoffs informing your approach.
* Avoid sending us draft design documents. It is difficult to evaluate which
  parts are draft and which parts are complete. Instead we encourage asking
  questions in Slack and sharing a design document that is ready to be
  reviewed.

Once the design document has been approved by two reviewers, move on to the
implementation.

## Implementation

Split your implementation into at least two Pull Requests, matching the
Ingestion and Business Queries requirements below, to give the team an
opportunity to review your code and provide feedback. This is in addition to
the design document Pull Request submitted earlier. Feel free to merge
each PR after you have two approvals.

Our team will do their best to provide a high quality review of the submitted
Pull Requests in a reasonable time frame. You are spending your time on this,
we are going to contribute our time too.

After the final submission, we will schedule a synchronous walkthrough call
where you walk the interview team through your data model, pipeline design,
and the trade-offs you made. The call includes a short Q&A about the design and
implementation choices you made.

After the walkthrough call, the interview team will assemble and vote using a
"+1, -2" anonymous voting system: +1 is submitted whenever a team member
accepts the submission, -2 otherwise.

In case of a positive result, we will connect you to our HR and recruiting
teams, who will work out the details and present an offer.

In case of a negative score result, the hiring manager will contact you and
share a list of the key observations from the team that affected the result.

### Tools

Write your data pipeline in Go. Use DuckDB as your data warehouse.
Everything should run locally, driven by a Makefile.

### Testing

Key components of the pipeline should have tests that cover the happy and
unhappy scenarios. Do not try to achieve 100% test coverage as that will take
too long.

# Requirements

Each level below lists the full set of requirements for that level. Higher
levels do not build implicitly on lower levels - if you are targeting Level 5,
read the Level 5 section for the complete scope, not just what's new relative
to Level 4.

Requirements are stated as the inputs you will be given, the outputs we
expect, and the constraints those outputs have to satisfy. How you get from
one to the other is your design decision. We are more interested in your
reasoning and in whether the result holds up than in any particular
implementation.

Within a level, requirements are grouped into the two pull requests you'll
submit, in order: an ingestion PR that gets the source data into the
warehouse, followed by a business queries PR that builds the tables an
analyst would actually query.

## Inputs (all levels)

You will receive three files exported from the fictional company's
operational systems: an HR/employee export, a CRM sales export, and a
product pricing sheet.

This data came out of systems that humans type into. Assume nothing about
its quality. Do not assume fields are populated, that types or formats are
consistent, that identifiers are unique, that keys join cleanly across
files, or that every value is plausible. Working out what is actually wrong
with this data, and deciding how to handle each case, is a substantial part 
of the challenge.

Level 4+: you will also receive a second CRM export covering a later
period. You'll have it from the start, but treat it as arriving after the
first batch has already been ingested - your pipeline should be able to
take it on without reprocessing or duplicating what's already loaded.

Level 5 only: you will also receive a third export covering a further
period still, produced by a newer version of the upstream system. Treat it,
too, as arriving after the earlier batches. It does not have the same
schema as the first two, and later exports are not guaranteed to be purely
additive relative to what came before.

## Level 3

### Ingestion

* Load the provided source data into DuckDB and shape it into tables that
  support the business queries below
* Running the pipeline twice against the same input must leave the warehouse
  in the same state as running it once.
* Every input row must be accounted for after a run: you should be able to
  say what happened to any given row and why. No single malformed input row
  should crash or halt the run
* If a downstream table is deleted or truncated, the pipeline must be able
  to restore it correctly
* The source data contains personal information. Identifying it is part of
  the task. No table that an analyst or any downstream consumer can query
  may expose it in readable form. Your design doc must state what you
  classified as sensitive, what you plan to do about it, and what risk
  remains

### Business Queries

* Publish tables that answer, at minimum:
  * revenue and units sold per product, per month
  * revenue per sales rep, per month
* An analyst must be able to answer each of those with a single SELECT
  against one of your tables. They should not have to join across your
  tables, filter out records you decided were untrustworthy, or know
  anything about how the data was cleaned
* Someone who has never read your pipeline code must be able to use your
  published tables and get the right numbers
* Your published totals must reconcile against the source. For any month,
  it must be possible to explain the difference between the totals in your
  tables and the totals in the raw input
* Provide a simple way to run each required business query and see its
  result, for example an additional `make` target
* A single `make` target takes a clean checkout to a queryable warehouse

## Level 4

### Ingestion

* Load the provided source data into DuckDB and shape it into tables that
  support the business queries below
* Running the pipeline twice against the same input must leave the warehouse
  in the same state as running it once.
* Every input row must be accounted for after a run: you should be able to
  say what happened to any given row and why. No single malformed input row
  should crash or halt the run
* If a downstream table is deleted or truncated, the pipeline must be able
  to restore it correctly
* The source data contains personal information. Identifying it is part of
  the task. No table that an analyst or any downstream consumer can query
  may expose it in readable form. Your design doc must state what you
  classified as sensitive, what you plan to do about it, and what risk
  remains
* A second batch covering a later period must be ingested without
  reprocessing or duplicating the first. After it arrives, both batches
  must be queryable together, and your tables must make clear which batch
  any given row came from

### Business Queries

* Publish tables that answer, at minimum:
  * revenue and units sold per product, per month
  * revenue per sales rep, per month
  * gross margin per product
* An analyst must be able to answer each of those with a single SELECT
  against one of your tables. They should not have to join across your
  tables, filter out records you decided were untrustworthy, or know
  anything about how the data was cleaned
* Someone who has never read your pipeline code must be able to use your
  published tables and get the right numbers
* Your published totals must reconcile against the source. For any month,
  it must be possible to explain the difference between the totals in your
  tables and the totals in the raw input
* Provide a simple way to run each required business query and see its
  result, for example an additional `make` target
* A single `make` target takes a clean checkout to a queryable warehouse
* When the pipeline produces output that can't be trusted, the run must
  fail: non-zero exit, and a message saying what went wrong. Deciding what
  "can't be trusted" means for this data is part of the design. A run that
  quietly publishes wrong numbers is a failed run

## Level 5

### Ingestion

* Load the provided source data into DuckDB and shape it into tables that
  support the business queries below
* Running the pipeline twice against the same input must leave the warehouse
  in the same state as running it once.
* Every input row must be accounted for after a run: you should be able to
  say what happened to any given row and why. No single malformed input row
  should crash or halt the run
* If a downstream table is deleted or truncated, the pipeline must be able
  to restore it correctly
* The source data contains personal information. Identifying it is part of
  the task. No table that an analyst or any downstream consumer can query
  may expose it in readable form. Your design doc must state what you
  classified as sensitive, what you plan to do about it, and what risk
  remains
* A second batch covering a later period must be ingested without
  reprocessing or duplicating the first. After it arrives, both batches
  must be queryable together, and your tables must make clear which batch
  any given row came from
* A third batch introduces a schema change and is not guaranteed to be
  purely additive relative to the batches that came before it. Ingesting it
  must not break the pipeline or lose data from any batch; decide on and
  document a policy for reconciling any overlap
* In the design document only, describe how a new data source could plug
  into your ingestion layer without a rewrite - you don't need to build one
* In the design document only, describe what happens to this pipeline at
  roughly 10^8 to 10^9 sales rows per year: what breaks first, what you
  would change, and what you would keep as is

### Business Queries

* Publish tables that answer, at minimum:
  * revenue and units sold per product, per month
  * revenue per sales rep, per month
  * gross margin per product
* An analyst must be able to answer each of those with a single SELECT
  against one of your tables. They should not have to join across your
  tables, filter out records you decided were untrustworthy, or know
  anything about how the data was cleaned
* Someone who has never read your pipeline code must be able to use your
  published tables and get the right numbers
* Your published totals must reconcile against the source. For any month,
  it must be possible to explain the difference between the totals in your
  tables and the totals in the raw input
* Provide a simple way to run each required business query and see its
  result, for example an additional `make` target
* A single `make` target takes a clean checkout to a queryable warehouse
* When the pipeline produces output that can't be trusted, the run must
  fail: non-zero exit, and a message saying what went wrong. Deciding what
  "can't be trusted" means for this data is part of the design. A run that
  quietly publishes wrong numbers is a failed run

# Guidance

## Code and project ownership

This is a test challenge and we have no intent of using the code you've
submitted in production. This is your work, and you are free to do whatever
you feel is reasonable with it. In the scenario where you don't pass, you can
open source it with any license and use it as a portfolio project.

## Areas of focus

Teleport focuses on networking, infrastructure and security.

These are the areas we will be evaluating in the submission:

* Reproducible pipelines. Scripts should be written in a way that allows
  reproduction of the environment.
* Data models. Design tables that are fast to query and clear enough to use
  without reading your pipeline code.
* Data quality. Bad data should fail loudly instead of quietly reaching
  downstream tables.
* Security. Data should be processed and stored in a way that does not leak
  sensitive data.
* Demo. The walkthrough should clearly explain your data model, pipeline
  design, and trade-offs to a technical audience.
* Communication. The design document and pull requests should be simple to
  understand and communicate key decisions to someone seeing them for the
  first time.

The primary factor in the team's decision is overall code quality. We are
looking for the highest possible quality with the smallest possible scope that
meets the requirements of the level you're applying for.

## Trade-offs

Write as little code as possible, otherwise this task will consume too much
time and code quality will suffer.

Please cut corners, for example configuration tends to take a lot of time, and
is not important for this task.

Use hardcoded values as much as possible and simply add TODO items showing
your thinking, for example:

```
// TODO: Add configuration system.
// Consider using a CLI library to support both environment variables and
// reasonable default values.

// TODO: Add retry logic.
```

Comments like this one are really helpful to us. They save yourself a lot of
time and demonstrate that you've spent time thinking about this problem and
provide a clear path to a solution.

It is okay if your pipeline is not optimized to handle very large data volumes.
Where relevant (Level 5), describe what you would change to handle
substantially more data rather than building it.

Consider making other reasonable trade-offs. Make sure you communicate them to
the interview team.

## Pitfalls & Gotchas

To help you prepare, here are some common reasons candidates have failed to
pass our interviews:

* *Use of AI.* Don't outsource your thinking to an AI. We recommend using AI
  for use cases like learning about a new problem space, exploring APIs, and
  finding missing edge cases. However, we strongly recommend you write the
  design document and all code yourself.
* *Jumping into implementation without clarifying requirements or narrowing
  the scope.* We expect you to ask clarifying questions about the data and
  requirements and identify the right scope for your target level during the
  design phase.
* *Scope creep.* Candidates have tried to design too much and ran out of time
  and energy. To avoid this pitfall, use the simplest solution that will work.
* *Overly complex designs.* Keep things simple and try to eliminate as many
  moving parts as possible. This is not only going to help in reviewing the
  solution, but is also often a way to distill a design to its essential
  parts.
* *Custom Security Algorithms.* Implementing custom security
  algorithms/authentication schemes is always a bad idea unless you are a trained
  security researcher/engineer. It is definitely a bad idea for this task.
  Try to stick to industry proven security methods as much as possible.
* *Not communicating.* Submitting all your code in a single PR, or
  splitting PRs along different lines than the Ingestion/Business Queries
  split we ask for, makes it harder for us to give you incremental
  feedback. We are a distributed team, so structured, asynchronous
  communication is critical to us.

## Questions

It is okay to ask the interview team questions. Some folks stay away from asking
questions to avoid appearing less experienced, so we provide examples of
questions to ask and questions we expect candidates to figure out on their
own.

Here is a great question to ask:

> Is it okay for everything to run locally? I would build a real data warehouse
> using something like Redshift or Snowflake, but for this challenge
> everything will be driven by a Makefile and run locally.

It demonstrates that you thought about this problem domain, recognized the
trade off and are saving you and the team time by not implementing it.

This is the question we expect candidates to figure out on their own:

> What version of Go should I use? What dependency manager should I use?

Unless specified in the requirements, pick the version that works best for
you.

Do not hesitate to reach out in case you get stuck or have any kind of general
questions or concerns. Please remember that communication is just as important
for this exercise as the code.

# Timing

This should take between 4 and 24 hours of focused work to complete. You can
split coding over a couple of weekdays or weekends and find time to ask
questions and receive feedback.

Once you join the Slack channel, you have a maximum of 2 weeks to complete the
challenge.

Within this timeframe, we don't give higher scores to challenges submitted
more quickly. We only evaluate the quality of the submission.

We only start the challenge if there is a position available and let all
candidates finish the submission.

We always aim to provide 1-2 rounds of feedback on all work that is submitted.
In order to be respectful of your time, we may opt to end the challenge early
if the submission does not improve after this feedback is suggested or if we
identify a large number of issues.
