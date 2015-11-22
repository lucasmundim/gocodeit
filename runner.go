package main

import "os/exec"

// Runner first argument is the command, like cat or echo,
// the second is the list of args to pass to it
type Runner interface {
	Run(string, ...string) ([]byte, error)
}

// RealRunner the real one
type RealRunner struct{}

var runner Runner

// Run the real runner for the actual program, actually execs the command
func (r RealRunner) Run(command string, args ...string) ([]byte, error) {
	out, err := exec.Command(command, args...).CombinedOutput()
	return out, err
}
