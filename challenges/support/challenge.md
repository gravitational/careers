# Summary

Investigate, troubleshoot, and remediate several issues in a simulated customer
environment.

# Rationale

* It helps us understand what to expect from you as a Support Engineer, how
  you reason about and troubleshoot production systems and how you communicate
  with customers.
* It helps you get a feel for what it would be like to work at Teleport, as
  this exercise aims to simulate our day-as-usual and expose you to the type of
  work we're doing here.

We believe this technique is not only better, but also is more fun compared to
whiteboard/quiz interviews so common in the industry. It's not without the
downsides - it could take longer than traditional interviews.

We appreciate your time and are looking forward to hacking on this project
together.

# Interview Process

## Before the Interview

Several days before the interview process begins, you will receive an
invitation to a private Slack channel. This channel will contain the interview
panel. You can ask the interview panel about the engineering culture, work-life
balance, or anything else that you would like to learn about Teleport.

## During the Interview

During the 2 hours allotted for the interview, you will be given access to the
hosts you will be debugging. Please ensure you are in the private Slack channel
during this time, as this is where you will get the invite link.

You will connect to the two faulty hosts using Teleport. You can access these
hosts in your web browser or the CLI. If you want to use the CLI, make sure you
[download and install](https://goteleport.com/download/) `tsh` ahead of time.
Your cluster will be running Teleport 14.3, make sure to download the correct
version of Teleport.

You will use the three commands below during the interview process.

* `tsh --proxy=interview-NAME.teleport.careers --user=interview-NAME login` to
  login to the Teleport cluster.
* `tsh ls` to see the list of (faulty) hosts you can connect to.
* `tsh ssh user@interview-NAME-0` or `tsh ssh user@interview-NAME-1` to connect
  to the faulty hosts.

Each system has two issues that you will need to investigate, troubleshoot, and
resolve. They will be outlined in `problem-*.txt` files in the user's home
directory on each host.

The interview panel will be in this channel during the interview process for
you to ask questions if you need help debugging the systems.

## After the Interview

After the interview, for each problem you were able to resolve, you will write
a short explanation (half a page each) that covers the following:

* What was wrong with the system(s)
* How you resolved the issue(s)
* Any future action items for the customer

The panel will review your document and the session recording (captured by
Teleport) and privately/anonymously submit +1/-2 to the hiring manager.
Afterward, the panel will meet to discuss how the session went.

In the case of a positive result, we will connect you to our HR team, collect
one or two references, and work out other details. You can start the reference
collection process in parallel if you would like to speed up the process. After
reference collection, our ops team will send you an offer.

In the case of a negative score result, the hiring manager will contact you and
share a list of the team's key observations that affected the result.

# Areas of Focus

We are not looking for any specific application knowledge. The test is designed
to troubleshoot generic configuration or system faults on any software deployed
to a generic Linux host. While this does not require any deployed application
knowledge, some familiarity with collecting logs or investigating unfamiliar
software is part of the process.

The faults may be present in the following areas.

* kernel / systemd / cgroups
* filesystem or volumes
* networking
* TLS connectivity

We use the following software to represent applications on the host, which are
experiencing faults.

* nginx
* etcd

As this is meant to test the way we normally work, you may use any resource you
would normally use when troubleshooting a system; for example, using Google as
a reference is normal and expected.

## Pitfalls and Gotchas

While this is a simulated environment, you should treat the challenge as
troubleshooting a live customer or production system.

It's OK to ask the interview team questions; we'll do our best to provide hints
or help steer the investigation if it's going in the wrong direction.
