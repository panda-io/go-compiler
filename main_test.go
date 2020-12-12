package main

import (
	"os/exec"
	"testing"
)

func TestVector(t *testing.T) {
	c := NewCompiler([]string{"cpp"})

	c.ParseFile("../panda/console.pd")
	//c.ParseFile("../panda/collection/vector.pd")
	c.ParseFile("./sample/vector.pd")
	c.Generate("./sample/vector.cpp")
	cmd := exec.Command("g++", "-o", "./sample/vector", "./sample/vector.cpp")
	err := cmd.Run()
	if err != nil {
		t.Error(err)
	}
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
