package main

import (
	"os/exec"
	"testing"
)

func TestVector(t *testing.T) {
	c := NewCompiler([]string{"cpp"})

	c.ParseFile("../panda/console.pd")
	//c.ParseFile("../panda/collection/vector.pd")
	c.ParseFile("./test/vector.pd")
	c.Generate("./test/main.cpp")
}

func TestSample(t *testing.T) {
	c := NewCompiler([]string{"cpp"})

	c.ParseFile("./sample/foobar.pd")
	c.Generate("./sample/foobar.cpp")
	cmd := exec.Command("g++", "-o", "./sample/foobar", "./sample/foobar.cpp")
	err := cmd.Run()
	if err != nil {
		t.Error(err)
	}
}
