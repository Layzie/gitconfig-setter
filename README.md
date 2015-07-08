# gitconfig-setter

## What's this?

This command (`gitconfig-setter`) sets local *user.name* and *user.email* of git.

You can set GitHub account in git global config, and set GHE account in git local config.

## Usage

The command is avelable on Mac. Other os needs `go build main.go`.

### Install

```sh
go install github.com/Layzie/git
```

### Make config file

Make `config.toml` such as:

```toml
name = "YOUR NAME" # local git config: user.name
email = "your@example.com" # local git config: user.email
```

### Run command
```sh
$ cd /path/to/set/config
$ git config --local -l
core.repositoryformatversion=0
core.filemode=true
core.bare=false
core.logallrefupdates=true
core.ignorecase=true
$ cd /path/to/gitconfig-setter
$ gitconfig-setter /path/to/repository /path/to/config
2015/08/07 18:41:13 ./'s local git config name has changed! Using ./config.toml
2015/08/07 18:41:13 ./'s local git config email has changed! Using ./config.toml
$ git config --local -l
core.repositoryformatversion=0
core.filemode=true
core.bare=false
core.logallrefupdates=true
core.ignorecase=true
user.name=YOUR NAME
user.email=your@example.com
```

## LICENCE

MIT
