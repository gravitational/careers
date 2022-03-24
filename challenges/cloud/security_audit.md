# Summary

Security engineers are often dropped into new environments and asked to figure out what's going on. And so we're going to do that to you. 

This challenge attempts to be a reasonable simulation of the work involved in securing a complicated application.

# Guidance

## Process
When you start the process, you will get an email with more detailed instructions, with links to reference material, and instructions on how to access the environment set up for this interview.

You'll be given a set of AWS Credentials, which will provide CLI access to the environment, as well as CodeCommit (AWS hosted git) access to the application source code. 

You are free to take as much or as little time as you want with the challenge, but we don't think it should take more than 4 hours. 

You are also free to take breaks any time after starting.

At the end, you'll submit a report of what you found.

## The Application
You’ll be looking at an internal corporate intranet app (a bad social network) running  in AWS.

The application is built on Django; in fact, it’s pretty close to vanilla Django. You’ll get both source code and default credentials to access it. It allows users to post messages, send private messages to each other, upload files, and stuff like that. 

We’re certain that the application has several serious vulnerabilities, and we’d like you to identify as many as you can.

We’re also certain that the AWS environment is misconfigured and missing many opportunities for improved security. Please consider:

## Interviewer
This interview is run by [Latacora](https://www.latacora.com/) which is one of our security partners, and based on how they interview for their teams.

# Tips
* You are free to take breaks any time after starting.
* The assessment shouldn't take more than 4 hours, this is because we do not want to waste too much of YOUR time. You are free to take as much or as little time as you want to.
* The report does not need to be formal - some of the best reports we have gotten were written in markdown.
* You are welcome to ask questions, but after the challenge starts we do role play as the client - and just like real life clients we may not always have the answer (:
* You will have as much time as you like, but please keep in touch or else the environment may get taken down.
* You will get more details in the email that kicks off the interview.

# Additional Resources
* [PortSwigger Web Security Academy](https://portswigger.net/web-security) labs is a cool site that will help you boost your cybersecurity skills. Also check out: https://owasp.org/www-project-top-ten/
* For cloud security we like [Summit Route](https://summitroute.com/blog). If you want a cloud book to reference, "Securing Devops "by Julien Vehent is good.
* We use Terraform a bunch, and we like [A Comprehensive Guide to Terraform](https://blog.gruntwork.io/a-comprehensive-guide-to-terraform-b3d32832baca).
* The horribly named [scrty.io](https://scrty.io/) is a great foundational resource for thinking seriously about starting security practices.
