# vf
vf is formatter for vim.

## Install
```bash
$ go get github.com/cappyzawa/vf/cmd/vf
```

## Example
```bash
$ golint ./... | vf
cmd/vf/main.go:30:1: exported method CLI.Run should have comment or be unexported

# with vf
$ golint ./... | vf
cmd/vf/main.go+30:1: exported method CLI.Run should have comment or be unexported
```

Vim can open a file and jump the line by using `+<num>`.
```bash
$ vim cmd/vf/main.go+30
```
