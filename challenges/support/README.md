# Summary

The Solutions and Technical Support Engineer sysadmin challenge is a simulated troubleshooting
session where you will investigate and remediate several faults in a simulated customer environment.

# Rationale

This exercise has two goals.

- It helps us understand what to expect from you as an engineer, how you reason about production systems and investigate faults.
- It helps you get a feel for what it would be like to work at Teleport, as this exercise aims to simulate some of our day-as-usual and expose you to the type of work we're doing here.

We believe this technique is better and is more fun compared to whiteboard/quiz interviews so common in the industry. [Some of the best teams use coding challenges.](https://sockpuppet.org/blog/2015/03/06/the-hiring-post/)

We appreciate your time and are looking forward to hacking on this challenge together.

# Requirements
The challenge is set up as a remote troubleshooting session.
Several Teleport engineers will join you in a Zoom call where you will
work to solve several faults in a simulated customer environment.

We'll ask you to share your terminal with us and connect over SSH to some hosts which have problems
that need fixing. We'll play the part of the customer that the hosts belong to, and will answer any
questions you have.

# Guidance
## Interview Process
The troubleshooting session is scheduled for 1.5 hours, with the following agenda:

- 5m - Intro and Setup
- 35m - Troubleshooting
- 5m - Optional built-in break
- 35m - Troubleshooting continued
- 10m - Wrap up, Q&A, etc.

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

As this is meant to test the way we normally work, you may use any resource you would normally use when troubleshooting a system; for example, using Google as a reference is normal and expected.

## Pitfalls and Gotchas
Some of the challenges exercise some of the hardest problems we've seen in the wild, so it's common to get stuck on one problem. If you do find you're not making progress on an issue, it's fine to move on and skip a problem and come back to it later.

While this is a simulated environment, you should treat the challenge as troubleshooting a live customer or production system.

It's OK to ask the interview team questions; we'll do our best to provide hints or help steer the investigation if the investigation is going in the wrong direction.


