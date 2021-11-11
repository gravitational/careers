# Frontend Application Developer

## Summary
Implement an application that allows a user to browse directory content on a remote server.

## Rationale
We would like to evaluate your skill in the following areas:
- Taking existing high-level requirements and translating them to a functional application
- Writing production level code that does not depend on gigabytes of npm packages
- Communicating with the team when working on the challenge
- Handling feedback

We believe this technique is not only better but also more fun compared to whiteboard/quiz interviews so common in the industry. It’s not without the downsides - it could take longer than traditional interviews. That said, it's our view that this type of challenge gives us a more accurate assessment of your ability to work well on the types of projects we’re working on day-to-day here at Teleport. [Some of the best teams use coding challenges](https://sockpuppet.org/blog/2015/03/06/the-hiring-post/). We appreciate your time and are looking forward to hacking on this project together.

## Tools
- CSS: any
- Frontend: React, React Router
- Backend: NodeJS, Golang, Rust
- Version Control: Github

## Requirements
  Implement an application that allows a user to browse directory content on a remote server.
  This application should have the following functionality:

   - A file/folder viewer with a client-side search and sorting capabilities (level 1-2)
   - Authentication with a login screen and backend API to browse a directory  (levels 2-4)

  There are 6 engineering levels at Teleport. It’s possible to score on levels 1-4 through this coding challenge; each level has its own set of requirements.

### Level 1
  This level focuses on JS/CSS/DOM knowledge and does not require any server-side logic to implement.

  Create a directory viewer with the following functionality:
  - Viewing and browsing of directory content
  - Viewing a file will show you metadata about the file (preview is not required)
  - Sorting and filtering on all attributes (within the current directory)
  - Allowing to bookmark URLs for quick navigation

  Directions:
  - Build your own components
  - Display the `name`, `size`, and `type` file attributes
  - Use human readable format to display the `size` of a file
  - You can use a basic server (even nginx) to serve the assets
  - Use below JSON structure to mock real directory

```js
const directory = {
  name: "teleport",
  sizeKb: 0,
  type: "dir",
  items: [
    {
      name: "lib",
      sizeKb: 0,
      type: "dir",
      items: [
        {
          name: "teleport.go",
          sizeKb: 320,
          type: "file",
        },
        {
          name: "test.go",
          sizeKb: 3320,
          type: "file",
        },
      ],
    },
    {
      name: "README.md",
      sizeKb: 4340,
      type: "file",
    },
  ],
};
```

  - You can use a design of similar viewers like Google Drive or Github
    as a starting point and then add missing pieces to it. For example, you can borrow CSS from these design systems:
    [github](https://primer.style/components/getting-started) or
    [google](https://material-ui.com/getting-started/usage/) if you like.

<img src="./assets/google.jpg" height="300" />

<img src="./assets/github.jpg" height="200" />

### Level 2-3
  Take level 1 requirements and instead of a static JSON create your own service with the following functionality:
  - A session management for login/logout
  - A login screen where an unauthenticated user is automatically redirected to (and then taken back to original URL)
  - An API that allows authenticated users to browse a server directory 1-level deep at a time

  Directions:
  - You can store user sessions in the memory or on disk
  - You can hardcode some parameters such as a directory location, or username/hash.
  - Use native/standard libraries. Avoid using out-of-the-box solutions like passportjs, or lodash
  - Use native browser API for networking.

### Level 4-5
  Add file upload/delete/download functionality to the level 2-3 requirements.

# Guidance

## Interview process
The interview team joins the Slack channel. The team consists of the engineers who will be working with you. Ask them about the engineering culture, work and life balance, or anything else that you would like to learn about Teleport.

Before writing the actual code, create a brief design document in Google Docs or markdown in Github and share with the team.

This document should include proposed UX of the app, tools and libraries that will be used, proposed API, implementation details where appropriate (for example: session management), and security.

Please avoid writing an overly detailed design document. Use this document to make sure the team could provide feedback on your design and demonstrate that you've investigated the problem space.

Split your code submission using pull requests and give the team an opportunity to review the PRs. A good “rule of thumb” to follow is that the final PR submission is a formality adding a small feature set - it means that the team had an opportunity to contribute the feedback during multiple well defined stages of your work.

Our team will do their best to provide a high quality review of the submitted pull requests in a reasonable time frame. You are spending your time on this, we are going to contribute our time too.

After the final submission, the interview team will assemble and vote using +1, -2 anonymous voting system: +1 is submitted whenever a team member accepts the submission, -2 otherwise.
In case of a positive result, we will connect you to our HR team who will collect one-two references and will work out other details. You can start the reference collection process in parallel if you would like to speed up the process.

After reference collection, our ops team will send you an offer.
In case of a negative score result, the hiring manager will contact you and send a list of key observations from the team that affected the result.

## Code and project ownership
This is a test challenge and we have no intent of using the code you’ve submitted in production. This is your work, and you are free to do whatever you feel is reasonable with it. In the scenario when you don’t pass, you can open source it with any license and use it as a portfolio project.

## Areas of focus
These are the areas we will be evaluating in the submission:

- Use consistent coding style. We follow [AirBnB Coding Style](https://github.com/airbnb/javascript) for JS and use [prettier](https://prettier.io/) code formatter.
- Create a few unit-tests for scenarios you think make sense.
- Make sure builds are reproducible. Pick any vendoring/packaging system that will allow us to get consistent build results.
- Ensure error handling and error reporting is consistent. The app should report clear errors and not crash under non-critical conditions.
- Ensure that your app is secure.

## Pitfalls and Gotchas
To help you out, we’ve composed a list of things that previously resulted in a no-pass from the interview team:

- Scope creep. Candidates have tried to implement too much and ran out of time, energy.
- Avoid using unnecessary 3rd party dependencies.
- Error handling. We pay extra attention to error handling. Make sure that they are properly handled and not ignored.
- URL navigation is a big part of great UX, do not ignore it.
- Keep your CSS simple but not simpler. Do not waste your time on animations, instead make sure that an extra character does not destroy your layout.
- Make sure that your code is secured and your application is not vulnerable to common web security vulnerabilities. For a senior level, make sure you have a good crypto setup and secure session management.

## Scoring

We want to be as transparent as possible on how we will be scoring your submission. The following table describes areas you will be evaluated on and how they will affect your overall score.

| Description                                                                                                          | Possible Points Awarded | Possible Points Subtracted |
|----------------------------------------------------------------------------------------------------------------------|-------------------------|----------------------------|
| The submitted code has a clear and modular structure.                                                                | +1                      | -1                         |
| The candidate outlined the key design points in the design document                                                  | +1                      | -1                         |
| The code provides examples of tests covering key components                                                          | +1                      | -1                         |
| The code provides clear error handling and reporting                                                                 | +1                      | -1                         |
| The app works according to the specifications, no bugs                                                               | +1                      | -1                         |
| The candidate demonstrates an ability to handle and apply feedback                                                   | +1                      | -1                         |
| The code is not vulnerable to common web security vulnerabilities                                                    | +1                      | -1                         |
