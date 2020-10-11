# conventionalcommits-git-hook

Simple Go code to ensure you commit message respect https://www.conventionalcommits.org/en/v1.0.0/.

## Install

```shell
$ # Clone this repository and cd into
$ go build -o conventionalcommits .

$ # Create a Git template directory
$ mkdir -p ~/.git_template/hooks

$ # Move the bin into this folder
$ mv conventionalcommits ~/.git_template/hooks/commit-msg

$ # Tell Git to use thiss directory
$ git config --global init.templatedir '~/.git_template'
```

This hook will be triggered AFTER you set the commit message but BEFORE validating the commit itself.
Each `git init` will creates a GIt repository with this template directory and ths hook.

More info about Git hooks : https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks

