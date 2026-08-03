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

## Design Doc

Before writing any actual code, we ask that you write a brief design document.
The design document should cover: your data model, pipeline architecture,
your approach to data validation and PII handling, and implementation
details where appropriate.

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

Within a level, requirements are grouped into the two pull requests you'll
submit, in order: an ingestion PR (load the provided raw data, normalize it,
and handle bad records), followed by a business queries PR (build new tables
on top of the ingested data to answer specific business questions).

## Level 3

### Ingestion

* Ingest the provided raw data into DuckDB using a Go/SQL pipeline
* Normalize the data, handling records with missing or incorrect-type fields
  without crashing the pipeline
* Pipeline re-runs must be idempotent - re-running against the same input must
  not create duplicate rows
* Tests covering happy path, malformed records, and duplicate-run scenarios

### Business Queries

* Build tables that support basic business queries (for example, revenue by
  sales rep, or units sold by product)
* Design efficient tables and queries for analysis
* Document the grain of each table you build (what one row represents)
* The entire pipeline, from ingestion through the final queryable tables,
  runs via a single Makefile target and is reproducible from a clean
  checkout

## Level 4

### Ingestion

* Ingest the provided raw data into DuckDB using a Go/SQL pipeline
* Normalize the data, handling records with missing or incorrect-type fields
  without crashing the pipeline
* Pipeline re-runs must be idempotent - re-running against the same input must
  not create duplicate rows
* Validate raw data to detect incorrect types, duplicates, missing data, and
  magnitude errors (for example, an implausibly large order quantity)
* Invalid records are quarantined for review rather than crashing the pipeline
  or being silently dropped
* Safely handle any PII present in the raw data (for example, employee names)
  using a simple technique like hashing, so it is not exposed unprotected in
  downstream tables
* Tests covering happy path, malformed records, duplicate-run, and validation
  edge cases

### Business Queries

* Build tables that support basic business queries (for example, revenue by
  sales rep, or units sold by product)
* Design efficient tables and queries for analysis
* Document the grain of each table you build (what one row represents)
* Automated data quality checks (for example: row count thresholds, null
  checks, referential checks between tables) that fail the pipeline loudly
  rather than silently producing bad data
* The entire pipeline, from ingestion through the final queryable tables,
  runs via a single Makefile target and is reproducible from a clean
  checkout

## Level 5

### Ingestion

* Ingest the provided raw data into DuckDB using a Go/SQL pipeline
* Normalize the data, handling records with missing or incorrect-type fields
  without crashing the pipeline
* Pipeline re-runs must be idempotent - re-running against the same input must
  not create duplicate rows
* Validate raw data to detect incorrect types, duplicates, missing data, and
  magnitude errors (for example, an implausibly large order quantity)
* Invalid records are quarantined for review rather than crashing the pipeline
  or being silently dropped
* Safely handle any PII present in the raw data (for example, employee names)
  using a simple technique like hashing, so it is not exposed unprotected in
  downstream tables
* Support ingesting a new batch of data where the schema has evolved (for
  example, a new column or a new product category), without breaking the
  pipeline or losing data
* Design the ingestion layer so that adding a new data source would not
  require rewriting the existing pipeline. Show this through a documented
  interface or abstraction; a working new source is not required
* Include performance considerations for scaling the pipeline to
  substantially larger data volumes
* Tests covering happy path, malformed records, duplicate-run, and validation
  edge cases

### Business Queries

* Build tables that support basic business queries (for example, revenue by
  sales rep, or units sold by product)
* Design efficient tables and queries for analysis
* Document the grain of each table you build (what one row represents)
* Automated data quality checks (for example: row count thresholds, null
  checks, referential checks between tables) that fail the pipeline loudly
  rather than silently producing bad data
* The entire pipeline, from ingestion through the final queryable tables,
  runs via a single Makefile target and is reproducible from a clean
  checkout

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
* Data models. Design performant tables. Keep queries fast, plan for schema
  migration, and document the grain of each table so joins and aggregations
  don't silently double- or under-count.
* Data quality. Add validation and test coverage to all stages of your
  pipeline.
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
* *Suggesting custom security algorithms* for handling PII is always a bad
  idea unless you are a trained security researcher/engineer. Stick to
  industry proven methods.

## Questions

It is OK to ask the interview team questions. Some folks stay away from asking
questions to avoid appearing less experienced, so we provide examples of
questions to ask and questions we expect candidates to figure out on their
own.

Here is a great question to ask:

> Is it OK for everything to run locally? I would build a real data warehouse
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

It should take you from 4 and no more than 24 full hours to complete the
challenge. You can split coding over a couple of weekdays or weekends and find
time to ask questions and receive feedback.

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
