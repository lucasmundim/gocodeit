package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func set(paths []string, project string) {
	if isProjectInPath(paths, project) {
		fmt.Println("Current project's vendor is already at lib-path. Nothing to do.")
	} else {
		addProjectToLibpath(paths, project)
		fmt.Println("Current project's vendor successfully added to gocode's lib-path")
	}
}

func show(paths []string, project string) {
	fmt.Println("Current project's vendor:", currentProjectVendor(project))
	fmt.Println("gocode lib-path:")
	for _, path := range paths {
		fmt.Println("  -", path)
	}
}

func status(paths []string, project string) {
	if isProjectInPath(paths, project) {
		fmt.Println("Current project's vendor is already at lib-path.")
	} else {
		fmt.Println("Current project's vendor is not at lib-path")
	}
}

func reset(project string) {
	setLibPath([]string{currentProjectVendor(project)})
	fmt.Println("Current project's vendor successfully reset to gocode's lib-path")
}

func unset(paths []string, project string) {
	if isProjectInPath(paths, project) {
		removeProjectFromLibpath(paths, project)
		fmt.Println("Current project's vendor successfully removed to gocode's lib-path")
	} else {
		fmt.Println("Current project's vendor is not at lib-path")
	}
}

func isProjectInPath(paths []string, project string) bool {
	var found bool
	for _, path := range paths {
		if strings.HasSuffix(path, filepath.Join(project, "vendor")) {
			found = true
		}
	}
	return found
}

func addProjectToLibpath(paths []string, project string) {
	paths = append(paths, currentProjectVendor(project))
	setLibPath(paths)
}

func removeProjectFromLibpath(paths []string, project string) {
	var index int
	for i, path := range paths {
		if path == currentProjectVendor(project) {
			index = i
			break
		}
	}
	paths = append(paths[:index], paths[index+1:]...)
	setLibPath(paths)
}

func currentProjectVendor(project string) string {
	osArch := runtime.GOOS + "_" + runtime.GOARCH
	return filepath.Join(currentGoPath(), "pkg", osArch, project, "vendor")
}

func currentGoPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, path := range strings.Split(os.Getenv("GOPATH"), ":") {
		if strings.HasPrefix(pwd, path) {
			return path
		}
	}
	return ""
}
