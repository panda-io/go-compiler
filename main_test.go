package main

import (
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
}
