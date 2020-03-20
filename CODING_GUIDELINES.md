# Coding Guidelines
This document describes coding guidelines.

## panic
No `panic` must be used.
It endangeres the stability of the whole application.

## Comments
If couse we want comments.
However, make sure that they are wrapped neatly, and don't cause crazy long lines.

## Committing changes
Whatever you work on, create a new branch for it.
If the task is long running, the branch should have a meaningful name, otherwise, something like `Username-patch-1` is sufficient.
After finishing your work, create a PR.
If you want early reviews and feedback, create a draft PR.
See the codeowners file, to see what reviewers make sense.

## Reviews
Before anything is merged into `master`, there is at least one review approval needed.

# VSCode
Suggested plugins:
* `Go`
* `Rewrap`
* `vscode-icons`