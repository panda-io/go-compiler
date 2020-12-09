package main

import (
	"os/exec"
	"testing"
)

func TestSample(t *testing.T) {
	c := NewCompiler([]string{"cpp"})

	c.ParseFile("./sample/foobar.pd")
	c.Generate("./sample/main.cpp")
	cmd := exec.Command("g++", "-o", "./sample/main", "./sample/main.cpp")
	err := cmd.Run()
	if err != nil {
		t.Error(err)
	}
}
