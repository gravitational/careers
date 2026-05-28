# Summary

Investigate and exploit security vulnerabilities in a cloud-hosted Kubernetes cluster.

# Rationale

This exercise has two goals:

* It helps us to understand what to expect from you as an Offensive Security
  Engineer and how you find security flaws in an environment.
* It helps you get a feel for what it would be like to work at Teleport, as
  this exercise aims to simulate our day-as-usual and expose you to the type of
  work we're doing here.

We believe this technique is not only better, but also is more fun compared to
whiteboard/quiz interviews so common in the industry. It's not without the
downsides - it could take longer than traditional interviews.

[Some of the best teams use coding challenges](https://sockpuppet.org/blog/2015/03/06/the-hiring-post/) — and we apply the same philosophy to non-coding roles.

We appreciate your time and are looking forward to hack on this project
together.

# Interview Process

## Before the Interview

On the day of the interview, you will be added to a private Slack channel with the interview panel. Questions about the challenge can be asked there.

We recommend installing the Teleport Client Tools on your local computer before this exercise.
See here: https://goteleport.com/download/client-tools/?os=linux

## During the Interview
During the 2 hours allotted for the interview, you will be given SSH access to a server via Teleport.

Your objective is to identify and exploit vulnerabilities in the environment. There are flags hidden across the environment, capture as many as you can. Your session will be recorded automatically by Teleport.

# Report

After you complete the challenge, you will write a red team report covering:

* The attack chain you followed and the techniques you used
* The vulnerabilities and flags you found
* Brief remediation or prevention suggestions for each finding

The panel will review your report alongside the session recording and
privately/anonymously submit +1/-2 to the hiring manager.

## Debrief

Following the interview, we will schedule a call to walk through your report
and discuss your approach with the panel.


## Areas of Focus

Your scope is the Kubernetes environment running on the lab server. The following are out of scope:

* Teleport infrastructure: the proxy, auth service, and web UI
* Any systems or infrastructure outside the lab server

## Pitfalls and Gotchas

* Use of AI. Don't outsource your thinking to an AI. While it is important to find the flags, its equally as important to be able to articulate what they were, how you found them and what the remediations should be.