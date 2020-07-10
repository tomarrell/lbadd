# Contributing

This project is open to contributors, and we'd like to make it as easy as possible to share your ideas, and your code.

Firstly, no contribution is too small. If you see something that's not right, please feel free to submit a spontaneous PR. It will very likely be reviewed same day.

## Where to start

If you don't know where to start, have a look at the list of [issues](./issues). If you find something interesting, feel free to say so in a comment, and you can be assigned to the issue. If you don't know specifically what you'd like to work on, feel free to open an issue with the label `work-wanted`, and someone will reach out to discuss the areas that are in need of contributions.

If you don't find something in the issues, don't worry! There is **plenty** that is *likely not listed there*. If you think you have a clear idea about something that should be done, feel free to open an issue yourself and start a discussion. A maintainer will get back to you, and most likely assign you to the task or clarify some requirements beforehand.


## Project Setup

In order to begin contributing, please fork the project to your personal Github account. You can then clone the project to your machine, and begin work on a new branch.

You'll need to have installed:

- Go compiler `version >=1.13`
- golint `master // TODO`
- errcheck `master // TODO`
- gosec `master // TODO`

It's recommended to join the slack organization to discuss your plans for changes early on. You can also ask any questions you might have throughout the process, or get help if you get stuck. To join, use the [invite link](https://join.slack.com/t/lbadd/shared_invite/zt-fk2eswyf-kbtIiXcJpQIWTHqb4jQhbA).

## Contributing

On this separate branch, make the minimum set of changes that are required to fulfill the task. If possible, please also write clear, package prefixed commit messages in order to make reviews easier. Once you're happy with the changes you've made, please make sure that you've added tests, and that the existing tests pass. We aim for high test coverage, so you'll likely be asked to add tests around areas that you've missed.

```bash
$ make test
```

This will verify that you're not causing any regressions in functionality.

You'll also need to run the linters. The linters that are run are listed above under **Project Setup**.

```bash
$ make lint
```

If both `exit 0` then you're good to go.

## Pull Request

Once you feel that you've finished the tasks, and the tests are passing, you can open a pull-request against the main repository. Please write a relatively descriptive pull-request message to make it easier to review your request.

> Note: for larger pieces of work, it's recommended to open a Draft PR in order to get early feedback.

Tests and linting will be automatically run on your PR using Github Actions.

Please assign one of the contributors, who will review your code and leave any feedback if necessary. If you do receive feedback, please do not be offended! We're trying to keep the quality standards high, and feedback is common, and will always be constructive. After a round of review and after any changes are made, then your PR will be approved, and merged into the `master` branch.

Congratulations!

## Becoming a Collaborator

If you'd like to become a maintainer, please first strive to make a meaningful contribution. Once you've done this and would like to work more actively on the project, you're more than welcome to ask to be added as a maintainer.

