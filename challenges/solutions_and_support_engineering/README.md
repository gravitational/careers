# Summary

The Solutions and Support Engineer sysadmin challenge is a simulated troubleshooting session where several faults have been created in a
sample environment requiring the candidate to investigate and solve the problems. 


> **Note:** For SREs/Devops without a backend development background, the on-call sysadmin challenges are used, but instead of remote hands, shell access to the nodes will be provided by provisioning a public ssh key.

# Rationale

This exercise has two goals.

- It helps us understand what to expect from you as an engineer,  how you reason about production systems and investigate faults.
- It helps you get a feel for what it would be like to work at Teleport, as this exercise aims to simulate some of our day-as-usual and expose you to the type of work we're doing here.

We believe this technique is better and is more fun compared to whiteboard/quiz interviews so common in the industry. It's not without the downsides as it could take longer than traditional interviews. [Some of the best teams use coding challenges.](https://sockpuppet.org/blog/2015/03/06/the-hiring-post/)

We appreciate your time and are looking forward to hacking on this challenge together.

# Requirements
The challenge is set up as a remote troubleshooting session. The candidate and Teleport engineers in similar roles will join a live troubleshooting session via zoom screen sharing. The candidate will be required to solve several faults created in a simulated customer environment. 

Because the hardest troubleshooting we have to do in customer environments is using remote hands without direct access to the system, this challenge is set up to simulate that worst-case experience of walking remote hands by troubleshooting the problems in the environment with only a shared screen.

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


