package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	re = regexp.MustCompile(`([a-zA-Z_/])+.\.[a-z]+:`)
)

// CLI has reader and writer
type CLI struct {
	reader    io.Reader
	outWriter io.Writer
	errWriter io.Writer
}

const (
	// ExitOK exits successfully
	ExitOK = iota
	// ScanError exits unsuccessfully by occuring scanning error
	ScanError
)

// Run runs command
func (c *CLI) Run(args []string) int {
	sc := bufio.NewScanner(c.reader)
	for sc.Scan() {
		if re.MatchString(sc.Text()) {
			fmt.Fprintln(c.outWriter, strings.Replace(sc.Text(), ":", "+", 1))
			continue
		}
		fmt.Fprintln(c.outWriter, sc.Text())
	}
	if sc.Err() != nil {
		fmt.Fprintf(c.errWriter, "failed to scan text: %v\n", sc.Err())
		return ScanError
	}

	return ExitOK
}

func main() {
	c := &CLI{
		reader:    os.Stdin,
		outWriter: os.Stdout,
		errWriter: os.Stderr,
	}
	os.Exit(c.Run(os.Args[1:]))
}
