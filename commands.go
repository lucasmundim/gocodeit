package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func projectName() string {
	output, err := runner.Run("glide", "name")
	if err != nil {
		fmt.Println("Error while executing 'glide name':", err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(output))
}

func getLibPath() []string {
	output, err := runner.Run("gocode", "set", "lib-path")
	if err != nil {
		fmt.Println("Error while executing 'gocode set lib-path':", err)
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
	_, err := runner.Run("gocode", "set", "lib-path", strings.Join(paths, ":"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
