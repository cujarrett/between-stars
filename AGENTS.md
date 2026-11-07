# Between Stars

A VS Code color theme. [themes/between-stars.json](./themes/between-stars.json) is the whole theme. Edit it by hand.

- `colors`: the editor chrome.
- `tokenColors`: syntax colors, copied from Dracula Purple (MIT).

## Rules

- **Never run `git add`, `git commit`, `git push`, or any git command that writes to the index, history or a remote.** Output the commands instead. Staging is the human's review.
- **Never add a `Co-Authored-By` trailer or a "Generated with Claude Code" line** to commit messages or PR descriptions.
- **Always give a suggested commit message**, with `git add` and the commit as two separate steps, listing every file by name. Never `git add .` or a bare directory.
- **Never output a `git push` command.**
- Run `/security-review` before suggesting a commit. Report one line when clean.

## Build

- `just build`: validates the theme and packages the `.vsix`.
- `just run`: installs that build into the local editor.

## Publishing

CI publishes on a push to main and skips when the version is already live. A release is a version bump in `package.json` plus a `CHANGELOG.md` entry. CI needs the `VSCE_PAT` repo secret, an Azure DevOps token with Marketplace publish scope.
