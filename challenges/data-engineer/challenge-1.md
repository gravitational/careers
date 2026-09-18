# Summary

Design and build an analytics warehouse for a fictional business.

# Rationale

This exercise has two goals:

* It helps us understand what to expect from you as a data engineer, how you
  reason about data model design and business definitions, and how you
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

You will receive data from a fictional B2B SaaS company called Northstar,
including CRM accounts, subscription contracts, application workspaces,
product usage, and cloud costs. The source data is already loaded into a
DuckDB database. You must model it so that it can be queried and analyzed.

Northstar sells monthly and annual subscriptions. A customer account can own
several workspaces, and finance wants to understand annualized subscription
value, product usage, cloud cost, and gross margin by customer.

## Design Doc

Before writing any actual code, we ask that you write a brief design document.
The design document should cover: your data model, production ingestion and
pipeline architecture, where each transformation happens and why, your
approach to data validation and PII handling, and implementation details where
appropriate. If you are targeting Level 5, also cover how you'd reconcile a
later batch that isn't purely additive, including schema changes.

Include approaches you have evaluated and reasoning for picking the approach
you're planning to go with. Be explicit about what you will build, defer, or
intentionally omit to complete the exercise within two weeks.

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
Canonical Models and Business Queries requirements below, to give the team an
opportunity to review your code and provide feedback. This is in addition to
the design document Pull Request submitted earlier. Feel free to merge
each PR after you have two approvals.

Our team will do their best to provide a high quality review of the submitted
Pull Requests in a reasonable time frame. You are spending your time on this,
we are going to contribute our time too.

After the final submission, we will schedule a synchronous walkthrough call 
where you explain your data model and the trade-offs you made. The call includes 
a short Q&A about the design and implementation choices you made.

After the walkthrough call, the interview team will assemble and vote using a
"+1, -2" anonymous voting system: +1 is submitted whenever a team member
accepts the submission, -2 otherwise.

In case of a positive result, we will connect you to our HR and recruiting
teams, who will work out the details and present an offer.

In case of a negative score result, the hiring manager will contact you and
share a list of the key observations from the team that affected the result.

### Tools

Use dbt for transformations, testing, and model documentation. Use DuckDB as
your data warehouse. Everything should run locally, driven by a Makefile.

The starter repository contains a pinned dbt environment, working dbt
configuration and source declarations, original synthetic source files, and a
DuckDB database with the source data already loaded in the `raw` schema. Do not
modify the supplied source files or raw tables. You do not need to write
ingestion code; describe the production ingestion approach in the design
document instead.

### Testing

Key components of the models should have tests that cover the happy and
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

Within a level, requirements are grouped into the two implementation pull
requests you'll submit, in order: a canonical models PR that creates the
reusable foundation, followed by a business queries PR that builds the table an
analyst would actually query.

## Inputs

You will receive synthetic exports from the fictional company's operational
systems: CRM accounts, subscription contracts, application workspaces, product
usage, and cloud costs. The files are already loaded into source tables in the
supplied DuckDB database, with every source column stored as text.

This data came out of operational systems. Assume nothing about its quality.
Do not assume fields are populated, that types or formats are consistent, that
identifiers are unique, that keys join cleanly across sources, or that every
value is plausible. Working out what is actually wrong with this data, and
deciding how to handle each case, is a substantial part of the challenge.

You will have all extracts from the start. If you are targeting Level 3 or 4,
the later contract and usage extracts are out of scope. Level 5 candidates
should treat them as arriving after the first extracts. Later data is not
guaranteed to be purely additive, and one usage extract has an evolved schema.

## Level 3

### Canonical Models

* Build staging models for the initial account, contract, workspace, and usage
  sources and shape them into canonical tables that support the business
  queries below
* Running the project twice against the same input must leave the warehouse in
  the same state as running it once
* Every input row in scope must be accounted for after a run: you should be
  able to say what happened to any given row and why
* If a downstream table is deleted or truncated, the project must be able to
  restore it correctly
* The source data contains personal information. Identifying it is part of the
  task. No table that an analyst or any downstream consumer can query may
  expose it in readable form. Your design doc must state what you classified
  as sensitive, what you plan to do about it, and what risk remains

### Business Queries

* Publish documented tables that answer, at minimum:
  * annualized subscription value by customer
  * product usage by customer
* An analyst must be able to answer each of those with a single SELECT against
  one of your tables. They should not have to join across your tables, filter
  out records you decided were untrustworthy, or know anything about how the
  data was cleaned
* Someone who has never read your model code must be able to use your published
  tables and get the right numbers
* Your published totals must reconcile against the source. It must be possible
  to explain the difference between the totals in your tables and the totals in
  the raw input
* Provide a simple way to run each required business query and see its result,
  for example an additional `make` target
* A single `make` target takes a clean checkout to a queryable warehouse and
  runs the tests

## Level 4

### Canonical Models

* Build staging models for the initial account, contract, workspace, and usage
  sources and shape them into canonical tables that support the business
  queries below
* Running the project twice against the same input must leave the warehouse in
  the same state as running it once
* Every input row in scope must be accounted for after a run: you should be
  able to say what happened to any given row and why
* If a downstream table is deleted or truncated, the project must be able to
  restore it correctly
* The source data contains personal information. Identifying it is part of the
  task. No table that an analyst or any downstream consumer can query may
  expose it in readable form. Your design doc must state what you classified
  as sensitive, what you plan to do about it, and what risk remains
* Also build staging models for the initial cloud-cost source
* Build a canonical customer dimension, workspace dimension,
  customer-to-workspace mapping, and current subscription or contract fact
* Normalize deterministic identifier differences, but keep identities that
  cannot be resolved safely in an explicit unresolved or quarantine output

### Business Queries

* Publish documented tables that answer, at minimum:
  * annualized subscription value by customer
  * product usage by customer
  * cloud cost by customer
  * gross margin by customer
* An analyst must be able to answer each of those with a single SELECT against
  one of your tables. They should not have to join across your tables, filter
  out records you decided were untrustworthy, or know anything about how the
  data was cleaned
* Someone who has never read your model code must be able to use your published
  tables and get the right numbers
* Your published totals must reconcile against the source. It must be possible
  to explain the difference between the totals in your tables and the totals in
  the raw input
* Provide a simple way to run each required business query and see its result,
  for example an additional `make` target
* A single `make` target takes a clean checkout to a queryable warehouse and
  runs the tests
* All four business questions must be answerable from a single published table
* Unallocated costs and unresolved customer mappings must remain visible rather
  than being silently dropped or assigned to a customer without evidence
* Define and implement verification checks for conditions that would make the
  published output untrustworthy. Deciding what those conditions are for this
  data is part of the design

## Level 5

### Canonical Models

* Build staging models for the initial account, contract, workspace, and usage
  sources and shape them into canonical tables that support the business
  queries below
* Running the project twice against the same input must leave the warehouse in
  the same state as running it once
* Every input row in scope must be accounted for after a run: you should be
  able to say what happened to any given row and why
* If a downstream table is deleted or truncated, the project must be able to
  restore it correctly
* The source data contains personal information. Identifying it is part of the
  task. No table that an analyst or any downstream consumer can query may
  expose it in readable form. Your design doc must state what you classified
  as sensitive, what you plan to do about it, and what risk remains
* Also build staging models for the initial cloud-cost source
* Build a canonical customer dimension, workspace dimension,
  customer-to-workspace mapping, and current subscription or contract fact
* Normalize deterministic identifier differences, but keep identities that
  cannot be resolved safely in an explicit unresolved or quarantine output
* Later extracts introduce corrections and a schema change. Processing them
  must not lose data from any batch; decide on and document a deterministic
  policy for reconciling overlap and preserve source lineage
* In the design document only, describe how these sources would enter a
  production warehouse. Explain what you would use managed, source-native, or
  custom ingestion for; you do not need to build it
* In the design document only, describe what would change if product usage
  grew substantially: what breaks first, what you would change, and what you
  would keep as is

### Business Queries

* Publish documented tables that answer, at minimum:
  * annualized subscription value by customer
  * product usage by customer
  * cloud cost by customer
  * gross margin by customer
* An analyst must be able to answer each of those with a single SELECT against
  one of your tables. They should not have to join across your tables, filter
  out records you decided were untrustworthy, or know anything about how the
  data was cleaned
* Someone who has never read your model code must be able to use your published
  tables and get the right numbers
* Your published totals must reconcile against the source. It must be possible
  to explain the difference between the totals in your tables and the totals in
  the raw input
* Provide a simple way to run each required business query and see its result,
  for example an additional `make` target
* A single `make` target takes a clean checkout to a queryable warehouse and
  runs the tests
* All four business questions must be answerable from a single published table
* Unallocated costs and unresolved customer mappings must remain visible rather
  than being silently dropped or assigned to a customer without evidence
* Define and implement verification checks for conditions that would make the
  published output untrustworthy. Deciding what those conditions are for this
  data is part of the design

# Guidance

## Code and project ownership

This is a test challenge and we have no intent of using the code you've
submitted in production. This is your work, and you are free to do whatever
you feel is reasonable with it. In the scenario where you don't pass, you can
open source it with any license and use it as a portfolio project.

## Areas of focus

Teleport focuses on networking, infrastructure and security.

These are the areas we will be evaluating in the submission:

* Business reasoning. Define material business concepts and make assumptions
  explicit.
* Reproducible projects. Scripts should be written in a way that allows
  reproduction of the environment.
* Data models. Design tables that are fast to query and clear enough to use
  without reading your model code.
* Data quality and reconciliation. Bad data should fail loudly instead of
  quietly reaching downstream tables, and published totals should be
  explainable from their sources.
* Security. Data should be processed and stored in a way that does not leak
  sensitive data.
* Corrections and schema evolution. Later data should be handled deliberately
  without silent loss or duplication.
* Architecture and scope. Use an appropriately simple design and make clear
  trade-offs about what not to build.
* Communication. The design document and pull requests should be simple to
  understand and communicate key decisions to someone seeing them for the
  first time.

We evaluate these areas holistically. A serious issue in any one of them can
determine the outcome. We are looking for the highest possible quality with
the smallest possible scope that meets the requirements.

## Trade-offs

Write as little code as possible, otherwise this task will consume too much
time and quality will suffer.

Please cut corners, for example configuration tends to take a lot of time, and
is not important for this task.

Use hardcoded values as much as possible and simply add TODO items showing
your thinking, for example:

```
-- TODO: Move the reporting date to governed configuration.
-- TODO: Add source freshness alerting in the production scheduler.
```

Comments like this one are really helpful to us. They save yourself a lot of
time and demonstrate that you've spent time thinking about this problem and
provide a clear path to a solution.

The system does not need to be perfect or production-ready. Prioritize a small,
complete solution and document lower-priority work instead of implementing it.
It is okay if your models are not optimized to handle very large data volumes.
Describe what you would change to handle substantially more data rather than
building it.

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
  requirements and identify the right scope during the design phase.
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
* *Not communicating.* Submitting all your code in a single PR, or splitting
  PRs along different lines than the Canonical Models/Business Queries split we
  ask for, makes it harder for us to give you incremental feedback. We are a
  distributed team, so structured, asynchronous communication is critical to
  us.

## Questions

It is okay to ask the interview team questions. Some folks stay away from asking
questions to avoid appearing less experienced, so we provide examples of
questions to ask and questions we expect candidates to figure out on their
own.

Here is a great question to ask:

> As of what date should annualized subscription value be measured, and should
> cancelled or future contracts contribute to it?

It demonstrates that you identified a material business definition that the
source data alone cannot answer.

This is the question we expect candidates to figure out on their own:

> Should I deploy a production warehouse or implement source connectors for
> this exercise?

No. Use the supplied local environment and describe the production approach in
the design document.

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
