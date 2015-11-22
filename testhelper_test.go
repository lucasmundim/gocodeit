package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

type TestRunner struct {
	fixture string
}

func (r TestRunner) Run(command string, args ...string) ([]byte, error) {
	if r.fixture == "" {
		r.fixture = "default"
	}
	cs := []string{"-test.run=TestHelperProcess", "--", r.fixture, command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	out, err := cmd.CombinedOutput()
	return out, err
}

func TestHelperProcess(*testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	defer os.Exit(0)

	args := os.Args
	for len(args) > 0 {
		if args[0] == "--" {
			args = args[1:]
			break
		}
		args = args[1:]
	}
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "No command\n")
		os.Exit(2)
	}

	fixture, cmd, args := args[0], args[1], args[2:]

	if _, found := outputFixture[cmd]; !found {
		fmt.Fprintf(os.Stderr, "Unknown command %q\n", cmd)
		os.Exit(2)
	}

	fmt.Fprintln(os.Stdout, outputFixture[cmd][fixture])
}
