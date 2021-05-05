# Summary

This exercise simulates a troubleshooting session where several faults and failures have been created in a sample environment, requiring the candidate to investigate and solve problems. 


# Requirements
The challenge is set up as a remote troubleshooting session. The candidate and Teleport engineers in similar roles will join a live troubleshooting session via zoom screen sharing. The candidate will be required to solve several faults created in a simulated customer environment. 

These problems represent the hardest type of troubleshooting we have to do, solving systems problems with unfamiliar software and simulating the worst-case experience of solving difficult problems we've witnessed in the field. These problems come from our experience solving unexpected problems and miss-configurations in a production environment.

# Guidance
## Interview Process
The interview team and candidate will join a scheduled zoom call for the troubleshooting session, scheduled for 2.5 hours.

- 5m - Intro and Setup
- 60m - Troubleshooting
- 10m - Built-in break
- 60m - Troubleshooting continued
- 15m - Wrap up, Q&A, etc.

After the session, the interview panel will privately/anonymously submit +1/-2 to the hiring manager, after which the panel will meet to discuss how the session went.

In the case of a positive result, we will connect you to our HR team, collect one-two references, and work out other details.  You can start the reference collection process in parallel if you would like to speed up the process. After reference collection, our ops team will send you an offer. 

In case of a negative score result, the hiring manager will contact you and share a list of the team's key observations that affected the result.

## Areas of Focus
We are not looking for any specific application knowledge. The test is designed to troubleshoot generic configuration or system faults on any software deployed to a generic Linux host. While this does not require any deployed application knowledge, some familiarity with collecting logs or investigating unfamiliar software is part of the process.

The faults are present in:
- kernel / systemd / cgroups
- filesystem or volumes
- networking
- SELinux
- TLS connectivity

We use the following software to represent applications on the host, which are experiencing faults:
- nginx
- etcd
- docker
- ipvs

As this is meant to test the way we normally work, you may use any resource you would normally use when troubleshooting a system; for example, using google as a reference is normal and expected.

## Pitfalls and Gotchas
Some of the challenges exercise some of the hardest problems we've seen in the wild, so it's common to get stuck on one problem. If you do find you're not making progress on an issue, it's fine to move on and skip a problem and come back to it later.

While this is a simulated environment, the candidate should treat the challenge as troubleshooting a live customer or production system.

It's OK to ask the interview team questions; we'll do our best to provide hints or help steer the investigation if the investigation is going in the wrong direction.

# Scoring
The session will be attended by a panel of team member, who will participate in the troubleshooting session. After the session, each team member will privately submit a vote of +1 or -2 within our tracking system along with comments. With your permission, we will record the zoom meeting incase any panel member cannot attend or any panel member needs to leave part way through the meeting. Any recordings will be deleted once each panel member has submitted their feedback.

## What we look for
- How the candidate reacts under pressure.
- How the candidate communicates their thought process while troubleshooting.
- The troubleshooting approach, compared to the level being applied for (the team will expect more from a senior candidate than a junior candidate).
- How the problems are solved.


# Tips
- It's not necessary to complete every problem. These are difficult problems to solve.
- It's easy to get hung up or stuck on an individual problem, don't be afraid to move on.
- Use any tools or resources you would use during your normal work day.
- We'll provide tips during the session if you get off course.
