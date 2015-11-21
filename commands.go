package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func projectName() string {
	output, err := exec.Command("glide", "name").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(output))
}

func getLibPath() []string {
	output, err := exec.Command("gocode", "set", "lib-path").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	r := regexp.MustCompile("lib-path \"(.*)\"")
	paths := r.FindStringSubmatch(string(output))[1]
	if len(paths) > 0 {
		return strings.Split(strings.TrimSpace(paths), ":")
	}
	return []string{}
}

func setLibPath(paths []string) {
	_, err := exec.Command("gocode", "set", "lib-path", strings.Join(paths, ":")).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
