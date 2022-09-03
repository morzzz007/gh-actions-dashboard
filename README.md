# GitHub Actions Dashboard TUI

A GitHub CLI extension written in Go to help you keep track of your GitHub action statuses without ever leaving your terminal.

## Installation

1. Install the `gh` CLI - see the [installation.](https://github.com/cli/cli#installation)

   _Installation requires a minimum version (2.0.0) of the the GitHub CLI that supports extensions._

2. Install this extension:

   ```sh
   gh extension install morzzz007/gh-actions-dashboard
   ```

## Configuring

Configuration is provided within a `config.yml` file under the extension's directory (usually `~/.config/gh-actions-dashboard/`)

An example `config.yml` file contains:

```yml
repoPaths: ["morzzz007/gh-actions-dashboard", "..."]
```

## Usage

Run:

```sh
gh actions-dashboard
```
