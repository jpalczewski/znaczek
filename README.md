# znaczek
[![Go Report Card](https://goreportcard.com/badge/github.com/jpalczewski/znaczek)](https://goreportcard.com/report/github.com/jpalczewski/znaczek)
[![Maintainability](https://api.codeclimate.com/v1/badges/cbe4df9ac2272d3b03db/maintainability)](https://codeclimate.com/github/jpalczewski/znaczek/maintainability)


Manage github repo labels with a breeze.

## Usage

```bash
znaczek e -r [repository name] -o [owner] 
```

## Synopsis
As of v0.1.3:
```
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

# Other notes
App make use of env variable `GH_TOKEN` with repo scope in order to manage labels.