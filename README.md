# tomles

`tomles` is a tool for modifying toml files via command line, scripts and automation.

[![GoDoc](https://godoc.org/github.com/n3wscott/tomles?status.svg)](https://godoc.org/github.com/n3wscott/tomles)
[![Go Report Card](https://goreportcard.com/badge/n3wscott/tomles)](https://goreportcard.com/report/n3wscott/tomles)

_Work in progress._

## Installation

`tomles` can be installed via:

```shell
go get github.com/n3wscott/tomles/cmd/tomles
```

To update your installation:

```shell
go get -u github.com/n3wscott/tomles/cmd/tomles
```

## Usage

`tomles` has one command at the moment: `update`

### Update 

```shell
Update changes the branch, version, or revision used for an import.

Examples:

  To change the dep on github.com/n3wscott/sockeye to the master branch in the local Gopkg.toml file,
  $ tomles update github.com/n3wscott/sockeye -b master -f ./Gopkg.toml

  Or the revision,
  $ tomles update github.com/n3wscott/sockeye -r 17fc779daec193476ff79fbe535ed00d426b08cb -f ./Gopkg.toml

  Or the version,
  $ tomles update github.com/n3wscott/sockeye -v v0.4.0 -f ./Gopkg.toml

  To see what the tomles command will produce,
  $ tomles update github.com/n3wscott/sockeye -b master -f ./Gopkg.toml --dry-run

  You can chain tomles output with a diff to see what will change,
  $ tomles update github.com/n3wscott/sockeye -b master -f ./Gopkg.toml --dry-run | diff ./Gopkg.toml -


Flags:
  -b, --branch string     The value of the branch to use.
  -D, --dry-run           Output what the new file would be.
  -f, --filename string   The path to the .toml file to use.
  -h, --help              help for update
  -r, --revision string   The value of the revision to use.
  -V, --verbose           Output more debug info to stderr
  -v, --version string    The value of the version to use.
```

