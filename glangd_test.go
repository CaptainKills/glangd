package main

import (
	"testing"

	cmd "github.com/CaptainKills/glangd/cmd"
)

func TestCommandToJson(t *testing.T) {
	var cmd cmd.CompileCommand

	cmd.Command = "gcc main.c"
	cmd.Directory = "/home/programmer/project-x/"
	cmd.File = "main.c"

	got := cmd.ToJson()
	want := `	{
		"directory": "/home/programmer/project-x/",
		"command": "gcc main.c",
		"file": "main.c"
	}`

	if got != want {
		t.Fatalf("got %q but wanted %q\n", got, want)
	}
}
