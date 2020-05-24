# znaczek

[![Go Report Card](https://goreportcard.com/badge/github.com/jpalczewski/znaczek)](https://goreportcard.com/report/github.com/jpalczewski/znaczek)
[![Maintainability](https://api.codeclimate.com/v1/badges/cbe4df9ac2272d3b03db/maintainability)](https://codeclimate.com/github/jpalczewski/znaczek/maintainability)
[![Build Status](https://travis-ci.com/jpalczewski/znaczek.svg?branch=master)](https://travis-ci.com/jpalczewski/znaczek)

Manage github repo labels with a breeze.

## Usage

```bash
znaczek e -r [repository name] -o [owner]
```

## Synopsis

As of v0.1.3:

```bash
NAME:
   znaczek - Managing github labels have never been easier!

USAGE:
   znaczek [global options] command [command options] [arguments...]

COMMANDS:
   export, e  export labels to file
   nuke, n    Delete all label in existing repo
   apply, a   Apply json file to repo
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -d, -v  Show further info (default: false)
   --quiet, -q      Hide unnecesary info (default: false)
   --help, -h       show help (default: false)
```

## Other notes

App make use of env variable `GH_TOKEN` with repo scope in order to manage labels.

## Shotouts, thanks, further readings

- First idea came out from <https://douglascayers.com/2019/08/01/how-to-export-and-import-github-issue-labels-between-projects/> - My googling also ended up in nothing useful, so there that simple tool came in.
- Makefile is inspired by <https://danishpraka.sh/2019/12/07/using-makefiles-for-go.html>
