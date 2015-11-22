package main

import "testing"

var outputFixture = map[string]map[string]string{
	"glide": {
		"blank":   "",
		"default": "github.com/someuser/somepackage",
	},
	"gocode": {
		"blank":   "lib-path \"\"",
		"default": "lib-path \"/home/user/gopath/pkg/linux_amd64/github.com/someuser/somepackage/vendor\"",
	},
}

func TestProjectName(t *testing.T) {
	runner = TestRunner{}
	defer func() { runner = RealRunner{} }()

	expected := "github.com/someuser/somepackage"

	out := projectName()
	if string(out) != expected {
		t.Errorf("Expected %q, got %q", expected, out)
	}
}

func TestBlankProjectName(t *testing.T) {
	runner = TestRunner{fixture: "blank"}
	defer func() { runner = RealRunner{} }()

	expected := ""

	out := projectName()
	if string(out) != expected {
		t.Errorf("Expected %q, got %q", expected, out)
	}
}
