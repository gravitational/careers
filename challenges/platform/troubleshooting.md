# Summary

This exercise simulates a troubleshooting session; The candidate must investigate and solve several faults and failures across Linux servers.

The challenge is set up as a remote troubleshooting session. Teleport engineers will follow along via Zoom screen sharing as the candidate SSHs into Linux servers and debugs a range of problems.

The sandbox problems are designed to simulate real-world troubleshooting encountered in our jobs. To solve them, the candidate will need to interact with unfamiliar software, debug systems level problems, and make changes to remediate the issues -- similar to responding to unexpected downtime. The problems stem from our experience debugging unexpected problems and misconfigurations in both our own and our customer's production environments.

# Guidance
## Interview Process
The interview team and candidate will join a scheduled Zoom call for the troubleshooting session, scheduled for 2.5 hours.

- 5m - Intro and Setup
- 60m - Troubleshooting
- 10m - Built-in break
- 60m - Troubleshooting continued
- 15m - Wrap up, Q&A, etc.

## Areas of Focus
Our test is designed to cover a range of configuration errors and system faults encountered on a Linux server. Successful candidates will be familiar with collecting logs from and investigating unfamiliar software in order to diagnose underlying system problems. We avoid challenges that require application specific expertise; Instead, we use applications as a jumping off point for diagnosing underlying system issues.

The faults are present in:
- kernel / systemd / cgroups
- filesystems and volumes
- networking
- TLS connectivity
- Application behavior

We use the following software to represent applications on the host, which are experiencing faults:
- nginx
- etcd
- docker
- kubernetes
- ipvs

While it is not required to be familiar with these applications, familiarity will lend a leg up when you need to querying logs or judging expected behavior.

This is meant to test the way we normally work; please use any resource you would use when troubleshooting a system at home or in your job. For example, it is encouraged to Google if you encounter a promising, but unfamiliar error message. We look favorably on candidates who know how to formulate good queries when they're stumped.

Similarly, you're welcome to install troubleshooting packages on the servers. While we've installed many common debugging tools, if you're used to having `htop`, `tldr`, or `execsnoop` at your fingertips, don't hesitate to install them.

## Pitfalls and Gotchas
Some challenges reproduce the hardest problems we've seen in the wild, so it's common to get stuck. If you don't know where to start or find you're not making progress on an issue, please skip that problem.  You may return to it later, after you complete the other problems.

While this is a simulated environment, the candidate should treat the challenge as troubleshooting a live production or customer system. If you'd like to try something risky, ask the interview team and we'll act as your peers. We hold it against candidates that "cowboy" in a potentially dangerous guess at a fix without understanding the underlying problem.

It's OK to ask the interview team questions; we'll do our best to provide hints or help steer the investigation if the investigation is going in the wrong direction.

# Scoring
The session is scored by a panel of peer Engineers attending the troubleshooting session.  With your permission, we will record the Zoom meeting in case a panel member cannot attend or any panel member needs to leave part way through the meeting. All recordings are deleted immediately after feedback is collected.

After participating in or viewing the session, each panelist will privately submit a vote of +1 or -2 to our applicant tracking system along with comments. In the case of a positive result, we will connect you to our HR team, who will reach out to discuss an offer. In case of a negative score result, the hiring manager will contact you and share a list of the team's key observations that affected the result of both interviews.

## What we look for
- Good debugging instincts under pressure. E.g. Determining which error messages are promising, and which are unrelated to the problem at hand.
- Clear communication of thought process while troubleshooting.
- The troubleshooting approach, compared to the level being applied for (the team will expect more from a senior candidate than a junior candidate).
- How the problems are solved.

# Tips
- It's not necessary to complete every problem. These are difficult problems to solve.
- It's easy to get hung up or stuck on an individual problem, don't be afraid to move on.
- Use any tools or resources you would use during your normal work day.
- We'll provide tips during the session if you get off course.
