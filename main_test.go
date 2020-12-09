package main

import (
	"os/exec"
	"testing"
)

func TestVector(t *testing.T) {
	c := NewCompiler([]string{"cpp"})

	c.ParseFile("../panda/console.pd")
	c.ParseFile("../panda/collection/vector.pd")
	c.ParseFile("./test/vector.pd")
	c.Generate("./test/main.cpp")

	t.Fail()
}

func TestSample(t *testing.T) {
	c := NewCompiler([]string{"cpp"})

	c.ParseFile("./sample/all_types.pd")
	c.Generate("./sample/all_types.cpp")
	cmd := exec.Command("g++", "-o", "./sample/all_types", "./sample/all_types.cpp")
	err := cmd.Run()
	if err != nil {
		t.Error(err)
	}
}
