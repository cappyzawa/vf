# vf
[![Actions Status](https://github.com/cappyzawa/vf/workflows/CI/badge.svg)](https://github.com/cappyzawa/vf/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/cappyzawa/vf)](https://goreportcard.com/report/github.com/cappyzawa/vf)
[![codecov](https://codecov.io/gh/cappyzawa/vf/branch/master/graph/badge.svg)](https://codecov.io/gh/cappyzawa/vf)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

vf is formatter for vim.

## Install
```bash
$ go get github.com/cappyzawa/vf/cmd/vf
```

## Example
```bash
$ golint ./...
cmd/vf/main.go:30:1: exported method CLI.Run should have comment or be unexported

# with vf
$ golint ./... | vf
cmd/vf/main.go+30:1: exported method CLI.Run should have comment or be unexported
```

Vim can open a file and jump the line by using `+<num>`.
```bash
$ vim cmd/vf/main.go+30
```
