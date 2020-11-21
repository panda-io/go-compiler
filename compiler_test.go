package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"testing"
)

func TestAllTypes(t *testing.T) {
	p, err := ParseFile("./test/all_types.pd", true, nil)
	if err != nil {
		t.Error(err)
	}

	if p.Imports[0].Name.Name != "sys" {
		t.Error("parse [import system] failed")
	}

	if p.Imports[0].Doc.Text != `"import_doc"` {
		t.Error("parse [import system] failed")
	}

	if p.Imports[0].Path.(*Ident).Name != "system" {
		t.Error("parse [import system] failed")
	}

	if p.Imports[1].Name != nil {
		t.Error("parse [import collection.test.sub] failed")
	}

	if p.Imports[1].Path.(*SelectorExpr).Selector.Name != "sub" {
		t.Error("parse [import collection.test.sub] failed")
	}

	if p.Imports[1].Path.(*SelectorExpr).Expr.(*SelectorExpr).Selector.Name != "test" {
		t.Error("parse [import collection.test.sub] failed")
	}

	if p.Values[0].Name.Name != "b" || !p.Values[0].Modifier.Public {
		t.Error("parse [public const b bool = true] failed")
	}

	if v, ok := p.Values[0].Value.(*BasicLit); !ok || v.Value != "true" {
		t.Error("parse [public const b bool = true] failed")
	}

	if p.Values[0].Doc.Text != `"const_doc"` {
		t.Error("parse doc for [public const b bool = true] failed")
	}
}

func TestNameSpace(t *testing.T) {
	s := "@doc `namespace_doc`\nnamespace test;\n"

	p, err := ParseString(s, true, nil)
	if err != nil {
		t.Error(err)
	}

	if p.Namespace.Path.(*Ident).Name != "test" || p.Namespace.Doc.Text != "`namespace_doc`" {
		t.Error("parse [namespace test] failed")
	}

	s = "@doc `namespace_doc`\nnamespace test.sub;\n"

	p, err = ParseString(s, true, nil)
	if err != nil {
		t.Error(err)
	}

	if p.Namespace.Path.(*SelectorExpr).Selector.Name != "sub" || p.Namespace.Doc.Text != "`namespace_doc`" {
		t.Error("parse [namespace test.sub] failed")
	}

	if p.Namespace.Path.(*SelectorExpr).Expr.(*Ident).Name != "test" {
		t.Error("parse [namespace test.sub] failed")
	}
}

//Test some errors

func TestGenerate(t *testing.T) {
	b, _ := ioutil.ReadFile("./test/all_types.pd")
	p, err := ParseString(string(b), true, nil)
	if err != nil {
		t.Error(err)
	}

	buff := bytes.NewBuffer(nil)
	p.Print(buff, true)

	ioutil.WriteFile("./test/all_types.cpp", buff.Bytes(), 0644)

	cmd := exec.Command("g++", "-o", "test/all_types", "test/all_types.cpp")
	err = cmd.Run()
	if err != nil {
		t.Error("compile failed:", err)
	}
}
