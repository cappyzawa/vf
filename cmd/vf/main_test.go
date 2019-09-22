package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestRun(t *testing.T) {
	cases := []struct {
		name     string
		testFile string
		expect   int
	}{
		{name: "format result of golint", testFile: "../../testdata/golint.txt", expect: 0},
		{name: "format result of go test", testFile: "../../testdata/gotest.txt", expect: 0},
	}

	t.Helper()
	for _, c := range cases {
		reader, _ := ioutil.ReadFile(c.testFile)
		cli := &CLI{
			reader:    bytes.NewBuffer(reader),
			outWriter: new(bytes.Buffer),
			errWriter: new(bytes.Buffer),
		}
		t.Run(c.name, func(t *testing.T) {
			actual := cli.Run([]string{})
			if actual != c.expect {
				t.Errorf("expect %d, but actual is %d", c.expect, actual)
			}
		})
	}
}
