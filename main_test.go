package main

import (
	"os/exec"
	"testing"
)

func TestVector(t *testing.T) {
	c := NewCompiler([]string{"cpp"})

	c.ParseFile("../panda/core/console.pd")
	c.ParseFile("../panda/collection/vector.pd")
	c.ParseFile("./sample/vector.pd")
	c.Compile("./sample/vector.cpp")
	cmd := exec.Command("g++", "-o", "./sample/vector", "./sample/vector.cpp")
	err := cmd.Run()
	if err != nil {
		t.Error(err)
	}
}

func TestSample(t *testing.T) {
	c := NewCompiler([]string{"cpp"})

	c.ParseFile("./sample/foobar.pd")
	c.Compile("./sample/foobar.cpp")
	cmd := exec.Command("g++", "-o", "./sample/foobar", "./sample/foobar.cpp")
	err := cmd.Run()
	if err != nil {
		t.Error(err)
	}
}

func TestBasic(t *testing.T) {
	c := NewCompiler(nil)

	c.ParseFile("../panda/libc/libc.pd")
	//c.ParseFile("../panda/core/allocator.pd")
	//c.ParseFile("../panda/core/string.pd")
	//c.ParseFile("../panda/core/counter.pd")
	c.ParseFile("./sample/basic.pd")
	c.Compile("./sample/basic")

	//TO-DO vector[any] for generic function call
}
