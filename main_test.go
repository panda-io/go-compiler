package main

/*
import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"testing"
)

func TestPanda(t *testing.T) {
	parser := NewParser(true, nil)
	parser.ParseFile("../compiler/main.pd")
	parser.ParseFile("../compiler/token/token.pd")
	parser.ParseFile("../panda/collection/vector.pd")
	p := parser.GetProgram()

	buff := bytes.NewBuffer(nil)
	p.Print(buff)

	ioutil.WriteFile("../compiler/panda.cpp", buff.Bytes(), 0644)

	cmd := exec.Command("g++", "-o", "../compiler/panda", "../compiler/panda.cpp")
	err := cmd.Run()
	if err != nil {
		t.Error("compile failed:", err)
	}
}*/
