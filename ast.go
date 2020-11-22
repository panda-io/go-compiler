package main

import (
	"bytes"
	"fmt"
	"strconv"
)

const (
	MetaCpp = "cpp"

	/*
		MetaSerialilzer = "serializer"
		MetaRef         = "ref"
		MetaCall        = "call"
		MetaReturn      = "return"
		MetaInclude     = "include"
		MetaMacro       = "macro"
	*/

	TabSize = 4
)

var (
	indents = []byte(" ")
)

func WriteIndent(buffer *bytes.Buffer, indent int) {
	if indent > len(indents) {
		for i := indent - len(indents); i > -1; i-- {
			indents = append(indents, byte(' '))
		}
	}
	buffer.Write(indents[:indent])
}

type Node interface {
	Pos() int // position of first character belonging to the node
}

type Expr interface {
	Node
	Print(buffer *bytes.Buffer)
	exprNode()
}

type Stmt interface {
	Node
	Print(buffer *bytes.Buffer, indent int)
	stmtNode()
}

type Decl interface {
	Node
	Print(buffer *bytes.Buffer, indent int, onlyDeclare bool)
	declNode()
}

// ----------------------------------------------------------------------------
// Meta

// A Metadata node represents a single //-style or /*-style comment.
type Metadata struct {
	Start  int // position of "@" starting the comment
	Name   string
	Text   string
	Values map[string]*BasicLit
}

func (meta *Metadata) Pos() int { return meta.Start }

// ----------------------------------------------------------------------------
// Modifier

// A Modifier node represents public or static to var, function, class, enum
type Modifier struct {
	Start  int
	Public bool
	Static bool
	Async  bool
}

func (modifier *Modifier) Pos() int { return modifier.Start }

// ----------------------------------------------------------------------------
// Field

type Field struct {
	Name    *Ident // field/method/parameter names; or nil
	Type    Expr   // field/method/parameter type
	Default Expr   // default value
}

func (f *Field) Pos() int {
	if f.Name != nil {
		return f.Name.Pos()
	}
	return f.Type.Pos()
}

func (f *Field) Print(buffer *bytes.Buffer) {
	f.Type.Print(buffer)
	if f.Name != nil {
		buffer.WriteString(" ")
		f.Name.Print(buffer)
		if f.Default != nil {
			buffer.WriteString(" =")
			f.Default.Print(buffer)
		}
	}
}

type FieldList struct {
	Start  int      // position of opening parenthesis/brace, if any
	Fields []*Field // field list; or nil
}

func (f *FieldList) Pos() int {
	if f.Start > 0 {
		return f.Start
	}
	if len(f.Fields) > 0 {
		return f.Fields[0].Pos()
	}
	return 0
}

func (f *FieldList) Print(buffer *bytes.Buffer) {
	for i, v := range f.Fields {
		if i != 0 {
			buffer.WriteString(", ")
		}
		v.Print(buffer)
	}
}

// ----------------------------------------------------------------------------
// Expressions

type BadExpr struct {
	Start int
}

func (x *BadExpr) Pos() int { return x.Start }

func (*BadExpr) exprNode() {}

func (x *BadExpr) Print(buffer *bytes.Buffer) {
	buffer.WriteString("/* Bad expr declared here */")
}

type Scalar struct {
	Start int
	Token Token
}

func (x *Scalar) Pos() int { return x.Start }

func (*Scalar) exprNode() {}

func (x *Scalar) Print(buffer *bytes.Buffer) {
	x.Token.Print(buffer)
}

type Ident struct {
	Start int     // identifier position
	Name  string  // identifier name
	Obj   *Object // denoted object; or nil
}

func (x *Ident) Pos() int { return x.Start }

func (*Ident) exprNode() {}

func (x *Ident) Print(buffer *bytes.Buffer) {
	buffer.WriteString(x.Name)
}

type EllipsisLit struct {
	Start int  // position of "..."
	Expr  Expr // ellipsis element type (parameter lists only); or nil
}

func (x *EllipsisLit) Pos() int { return x.Start }

func (*EllipsisLit) exprNode() {}

func (x *EllipsisLit) Print(buffer *bytes.Buffer) {
	buffer.WriteString("...")
	x.Expr.Print(buffer)
}

type BasicLit struct {
	Start int    // literal position
	Kind  Token  // INT, FLOAT, CHAR, or STRING
	Value string // literal string; e.g. 42, 0x7f, 3.14, 1e-9, 2.4i, 'a', '\x7f', "foo" or `\m\n\o`
}

func (*BasicLit) exprNode() {}

func (x *BasicLit) Pos() int { return x.Start }

func (x *BasicLit) Print(buffer *bytes.Buffer) {
	switch x.Kind {
	case STRING:
		if x.Value[0] == '"' {
			buffer.WriteString(x.Value)
		} else if x.Value[0] == '`' {
			buffer.WriteString(strconv.Quote(x.Value[1 : len(x.Value)-1]))
		}

	case INT, FLOAT:
		buffer.WriteString(x.Value)

	case CHAR:
		//TO-DO convert to unicode char
		buffer.WriteString(x.Value)

	case True, False, Void, Null:
		x.Kind.Print(buffer)

	default:
	}
}

type CompositeLit struct {
	Start  int
	Type   Expr   // literal type; or nil
	Values []Expr // list of composite elements; or nil
}

func (x *CompositeLit) Pos() int {
	if x.Type != nil {
		return x.Type.Pos()
	}
	return x.Start
}
func (*CompositeLit) exprNode() {}

func (x *CompositeLit) Print(buffer *bytes.Buffer) {
	buffer.WriteString("{")
	for i, v := range x.Values {
		if i != 0 {
			buffer.WriteString(", ")
		}
		v.Print(buffer)
	}
	buffer.WriteString("}")
}

type GenericLit struct {
	Start int    // < position
	Types []Expr // <int, int> <T>
}

func (x *GenericLit) Pos() int { return x.Start }

func (x *GenericLit) Print(buffer *bytes.Buffer) {
	buffer.WriteString("<")
	for i, v := range x.Types {
		if i != 0 {
			buffer.WriteString(", ")
		}
		v.Print(buffer)
	}
	buffer.WriteString(">")
}

type ParenExpr struct {
	Start int  // position of "("
	Expr  Expr // parenthesized expression
}

func (x *ParenExpr) Pos() int { return x.Start }

func (*ParenExpr) exprNode() {}

func (x *ParenExpr) Print(buffer *bytes.Buffer) {
	buffer.WriteString("(")
	x.Expr.Print(buffer)
	buffer.WriteString(")")
}

type SelectorExpr struct {
	Expr     Expr   // expression
	Selector *Ident // field selector
}

func (x *SelectorExpr) Pos() int { return x.Expr.Pos() }

func (*SelectorExpr) exprNode() {}

func (x *SelectorExpr) Print(buffer *bytes.Buffer) {
	x.Expr.Print(buffer)
	buffer.WriteString(".")
	x.Selector.Print(buffer)
}

type IndexExpr struct {
	Expr  Expr // expression
	Index Expr // index expression
}

func (x *IndexExpr) Pos() int { return x.Expr.Pos() }

func (*IndexExpr) exprNode() {}

func (x *IndexExpr) Print(buffer *bytes.Buffer) {
	x.Expr.Print(buffer)
	buffer.WriteString("[")
	x.Index.Print(buffer)
	buffer.WriteString("]")
}

type CallExpr struct {
	Func     Expr   // function expression
	Args     []Expr // function arguments; or nil
	Ellipsis int    // position of Args array, -1 not ellipsis
}

func (*CallExpr) exprNode() {}

func (x *CallExpr) Pos() int { return x.Func.Pos() }

func (x *CallExpr) Print(buffer *bytes.Buffer) {
	x.Func.Print(buffer)
	buffer.WriteString("(")
	for i, v := range x.Args {
		if i != 0 {
			buffer.WriteString(", ")
		}
		v.Print(buffer)
	}
	buffer.WriteString(")")
}

type NewExpr struct {
	Class Expr   // function expression
	Args  []Expr // function arguments; or nil
	//TO-DO elllipsis
	//Ellipsis
}

func (*NewExpr) exprNode() {}

func (x *NewExpr) Pos() int { return x.Class.Pos() }

func (x *NewExpr) Print(buffer *bytes.Buffer) {
	buffer.WriteString("std::make_shared<//TO-DO>(")
	for i, v := range x.Args {
		if i != 0 {
			buffer.WriteString(", ")
			v.Print(buffer)
		}
	}
	buffer.WriteString(")")
}

type EmitExpr struct {
	Meta *Metadata
}

func (x *EmitExpr) Pos() int { return x.Meta.Pos() }

func (*EmitExpr) exprNode() {}

func (x *EmitExpr) Print(buffer *bytes.Buffer) {
	buffer.WriteString(x.Meta.Text)
}

type UnaryExpr struct {
	Start int   // position of Op
	Op    Token // operator
	Expr  Expr  // operand
}

func (x *UnaryExpr) Pos() int { return x.Start }

func (*UnaryExpr) exprNode() {}

func (x *UnaryExpr) Print(buffer *bytes.Buffer) {
	x.Op.Print(buffer)
	x.Expr.Print(buffer)
}

type BinaryExpr struct {
	Left  Expr  // left operand
	Op    Token // operator
	Right Expr  // right operand
}

func (x *BinaryExpr) Pos() int { return x.Left.Pos() }

func (*BinaryExpr) exprNode() {}

func (x *BinaryExpr) Print(buffer *bytes.Buffer) {
	x.Left.Print(buffer)
	buffer.WriteString(" ")
	x.Op.Print(buffer)
	buffer.WriteString(" ")
	x.Right.Print(buffer)
}

type TernaryExpr struct {
	Condition Expr // operator
	First     Expr // left operand
	Second    Expr // right operand
}

func (x *TernaryExpr) Pos() int { return x.Condition.Pos() }

func (*TernaryExpr) exprNode() {}

func (x *TernaryExpr) Print(buffer *bytes.Buffer) {
	x.Condition.Print(buffer)
	buffer.WriteString(" ? ")
	x.First.Print(buffer)
	buffer.WriteString(" : ")
	x.Second.Print(buffer)
}

type KeyValueExpr struct {
	Key   Expr
	Value Expr
}

func (x *KeyValueExpr) Pos() int { return x.Key.Pos() }

func (*KeyValueExpr) exprNode() {}

func (x *KeyValueExpr) Print(buffer *bytes.Buffer) {
	x.Key.Print(buffer)
	buffer.WriteString(" : ")
	x.Value.Print(buffer)
}

// ----------------------------------------------------------------------------
// Statements

// A BadStmt node is a placeholder for statements containing
// syntax errors for which no correct statement nodes can be
// created.
//

type BadStmt struct {
	Start int // position range of bad statement
}

func (s *BadStmt) Pos() int { return s.Start }

func (*BadStmt) stmtNode() {}

func (s *BadStmt) Print(buffer *bytes.Buffer, indent int) {
	WriteIndent(buffer, indent)
	buffer.WriteString("/* bad statement here. */")
}

// A DeclStmt node represents a declaration in a statement list.
type DeclStmt struct {
	Decl *ValueDecl // *GenDecl with CONST, or VAR token
}

func (s *DeclStmt) Pos() int { return s.Decl.Pos() }

func (*DeclStmt) stmtNode() {}

func (s *DeclStmt) Print(buffer *bytes.Buffer, indent int) {
	s.Decl.Print(buffer, indent, false)
}

// An EmptyStmt node represents an empty statement.
// The "position" of the empty statement is the position
// of the immediately following (explicit or implicit) semicolon.
//
type EmptyStmt struct {
	Start int // position of following ";"
}

func (s *EmptyStmt) Pos() int { return s.Start }

func (*EmptyStmt) stmtNode() {}

func (s *EmptyStmt) Print(buffer *bytes.Buffer, indent int) {
	buffer.WriteString(";")
}

// An ExprStmt node represents a (stand-alone) expression
// in a statement list.
//
type ExprStmt struct {
	Expr Expr // expression
}

func (s *ExprStmt) Pos() int { return s.Expr.Pos() }

func (*ExprStmt) stmtNode() {}

func (s *ExprStmt) Print(buffer *bytes.Buffer, indent int) {
	s.Expr.Print(buffer)
}

// An ExprStmt node represents a (stand-alone) expression
// in a statement list.
//
type EmitStmt struct {
	Start   int
	Content string // expression
}

func (s *EmitStmt) Pos() int { return s.Start }

func (*EmitStmt) stmtNode() {}

func (s *EmitStmt) Print(buffer *bytes.Buffer, indent int) {
	buffer.WriteString(s.Content[1 : len(s.Content)-1])
}

// An IncDecStmt node represents an increment or decrement statement. ++ --
type IncDecStmt struct {
	Expr Expr
	Tok  Token // INC or DEC
}

func (s *IncDecStmt) Pos() int { return s.Expr.Pos() }

func (*IncDecStmt) stmtNode() {}

func (s *IncDecStmt) Print(buffer *bytes.Buffer, indent int) {
	s.Expr.Print(buffer)
	s.Tok.Print(buffer)
}

type EnumStmt struct {
	Name  *Ident
	Value *BasicLit
}

func (s *EnumStmt) Pos() int { return s.Name.Pos() }

func (*EnumStmt) stmtNode() {}

func (s *EnumStmt) Print(buffer *bytes.Buffer, indent int) {
	WriteIndent(buffer, indent)
	s.Name.Print(buffer)
	if s.Value != nil {
		buffer.WriteString(" = ")
		s.Value.Print(buffer)
	}
	buffer.WriteString(",\n")
}

// An AssignStmt node represents an assignment or
// a short variable declaration.
//
type AssignStmt struct {
	Left  Expr
	Tok   Token // assignment token, DEFINE
	Right Expr
}

func (s *AssignStmt) Pos() int { return s.Left.Pos() }

func (*AssignStmt) stmtNode() {}

func (s *AssignStmt) Print(buffer *bytes.Buffer, indent int) {
	s.Left.Print(buffer)
	buffer.WriteString(" = ")
	s.Right.Print(buffer)
}

// A ReturnStmt node represents a return statement.
type ReturnStmt struct {
	Start  int  // position of "return" keyword
	Result Expr // result expressions; or nil
}

func (s *ReturnStmt) Pos() int { return s.Start }

func (*ReturnStmt) stmtNode() {}

func (s *ReturnStmt) Print(buffer *bytes.Buffer, indent int) {
	buffer.WriteString("return")
	if s.Result != nil {
		buffer.WriteString(" ")
		s.Result.Print(buffer)
	}
}

// A BranchStmt node represents a break, continue, goto,
// or fallthrough statement.
//
type BranchStmt struct {
	Start int   // position of Tok
	Tok   Token // keyword token (BREAK, CONTINUE)
}

func (s *BranchStmt) Pos() int { return s.Start }

func (*BranchStmt) stmtNode() {}

func (s *BranchStmt) Print(buffer *bytes.Buffer, indent int) {
	s.Tok.Print(buffer)
}

// A BlockStmt node represents a braced statement list.
type BlockStmt struct {
	Start int // position of "{"
	Stmts []Stmt
}

func (s *BlockStmt) Pos() int { return s.Start }

func (*BlockStmt) stmtNode() {}

func (s *BlockStmt) Print(buffer *bytes.Buffer, indent int) {
	WriteIndent(buffer, indent)
	buffer.WriteString("{\n")
	for _, v := range s.Stmts {
		WriteIndent(buffer, indent+TabSize)
		v.Print(buffer, indent+TabSize)
		if _, ok := v.(*EmitStmt); !ok {
			buffer.WriteString(";")
		}
		buffer.WriteString("\n")
	}
	WriteIndent(buffer, indent)
	buffer.WriteString("}\n")
}

// An IfStmt node represents an if statement.
type IfStmt struct {
	Start     int  // position of "if" keyword
	Condition Expr // condition
	Body      *BlockStmt
	Else      Stmt // else branch; or nil
}

func (s *IfStmt) Pos() int { return s.Start }

func (*IfStmt) stmtNode() {}

func (s *IfStmt) Print(buffer *bytes.Buffer, indent int) {
	WriteIndent(buffer, indent)
	buffer.WriteString("if (")
	s.Condition.Print(buffer)
	buffer.WriteString(")\n")
	WriteIndent(buffer, indent)
	buffer.WriteString("{\n")
	s.Body.Print(buffer, indent+4)
	WriteIndent(buffer, indent)
	buffer.WriteString("}\n")
}

// A CaseClause represents a case of an expression or type switch statement.
type CaseClause struct {
	Start int    // position of "case" or "default" keyword
	Expr  Expr   // list of expressions or types; nil means default case
	Body  []Stmt // statement list; or nil
}

func (s *CaseClause) Pos() int { return s.Start }
func (*CaseClause) stmtNode()  {}
func (s *CaseClause) Print(buffer *bytes.Buffer, indent int) {
	WriteIndent(buffer, indent)
	buffer.WriteString("case ")
	s.Expr.Print(buffer)
	buffer.WriteString(":\n")

	for _, v := range s.Body {
		v.Print(buffer, indent+4)
	}
}

// A SwitchStmt node represents an expression switch statement.
type SwitchStmt struct {
	Start int        // position of "switch" keyword
	Tag   Expr       // tag expression
	Body  *BlockStmt // CaseClauses only
}

func (s *SwitchStmt) Pos() int { return s.Start }

func (*SwitchStmt) stmtNode() {}

func (s *SwitchStmt) Print(buffer *bytes.Buffer, indent int) {
	WriteIndent(buffer, indent)
	buffer.WriteString("switch (")
	s.Tag.Print(buffer)
	buffer.WriteString(")\n")
	s.Body.Print(buffer, indent+4)
}

// A ForStmt represents a for statement.
type ForStmt struct {
	For       int  // position of "for" keyword
	Init      Stmt // initialization statement; or nil
	Condition Expr // condition; or nil
	Post      Stmt // post iteration statement; or nil
	Body      *BlockStmt
}

func (s *ForStmt) Pos() int { return s.For }

func (*ForStmt) stmtNode() {}

func (s *ForStmt) Print(buffer *bytes.Buffer, indent int) {
	WriteIndent(buffer, indent)
	buffer.WriteString("for (")
	if s.Init != nil {
		s.Init.Print(buffer, 0)
	}
	buffer.WriteString(" ;")
	if s.Condition != nil {
		s.Condition.Print(buffer)
	}
	buffer.WriteString(" ;")
	if s.Post != nil {
		s.Post.Print(buffer, 0)
	}
	buffer.WriteString(" )\n")
	s.Body.Print(buffer, indent)
}

type ForInStmt struct {
	For      int  // position of "for" keyword
	Init     Stmt // initialization statement; or nil
	Iterator Expr // condition; or nil
	Body     *BlockStmt
}

func (s *ForInStmt) Pos() int { return s.For }

func (*ForInStmt) stmtNode() {}

func (s *ForInStmt) Print(buffer *bytes.Buffer, indent int) {
	WriteIndent(buffer, indent)
	buffer.WriteString("for (")
	s.Init.Print(buffer, 0)
	buffer.WriteString(" : ")
	s.Iterator.Print(buffer)
	buffer.WriteString(" )\n")
	s.Body.Print(buffer, indent)
}

// ----------------------------------------------------------------------------
// Declarations
// A declaration is represented by one of the following declaration nodes.

// A BadDecl node is a placeholder for declarations containing
// syntax errors for which no correct declaration nodes can be
// created.
//
type BadDecl struct {
	Start int // position range of bad declaration
}

func (d *BadDecl) Pos() int { return d.Start }

func (*BadDecl) declNode() {}

func (*BadDecl) Print(buffer *bytes.Buffer, indent int, onlyDeclare bool) {}

type NamespaceDecl struct {
	Doc  *Metadata // associated documentation; or nil
	Path Expr      // namespace path
}

func (s *NamespaceDecl) Pos() int { return s.Path.Pos() }

func (*NamespaceDecl) declNode() {}

type ImportDecl struct {
	Doc  *Metadata // associated documentation; or nil
	Name *Ident    // local name; or nil
	Path Expr      // import path
}

func (s *ImportDecl) Pos() int {
	if s.Name != nil {
		return s.Name.Pos()
	}
	return s.Path.Pos()
}
func (*ImportDecl) declNode() {}

func (*ImportDecl) Print(buffer *bytes.Buffer, indent int, onlyDeclare bool) {}

type ValueDecl struct {
	Doc      *Metadata // associated documentation; or nil
	Modifier *Modifier
	Name     *Ident // value names (len(Names) > 0)
	Generic  *GenericLit
	Type     Expr // value type; or nil
	Value    Expr // initial values; or nil
}

func (s *ValueDecl) Pos() int { return s.Name.Pos() }

func (*ValueDecl) declNode() {}

func (v *ValueDecl) Print(buffer *bytes.Buffer, indent int, onlyDeclare bool) {
	WriteIndent(buffer, indent)
	v.Type.Print(buffer)
	buffer.WriteString(" ")
	v.Name.Print(buffer)
	if v.Value != nil {
		buffer.WriteString(" = ")
		v.Value.Print(buffer)
	}
	buffer.WriteString(";\n")
}

type ClassDecl struct {
	Doc       *Metadata // associated documentation; or nil
	Modifier  *Modifier
	Parents   []Expr
	Name      *Ident // type name
	Generic   *GenericLit
	Values    []*ValueDecl
	Functions []*FuncDecl
}

func (c *ClassDecl) Pos() int { return c.Name.Pos() }

func (*ClassDecl) declNode() {}

func (c *ClassDecl) Print(buffer *bytes.Buffer, indent int, onlyDeclare bool) {
	WriteIndent(buffer, indent)
	buffer.WriteString("class ")
	c.Name.Print(buffer)
	//TO-DO inheritance
	if onlyDeclare {
		buffer.WriteString(";\n")
		return
	}
	buffer.WriteString("\n")
	WriteIndent(buffer, indent)
	buffer.WriteString("{\n")
	buffer.WriteString("public:\n")
	for _, v := range c.Values {
		v.Print(buffer, indent+TabSize, false)
		buffer.WriteString("\n")
	}
	for _, f := range c.Functions {
		f.Print(buffer, indent+TabSize, false)
		buffer.WriteString("\n")
	}
	buffer.WriteString("};\n")
}

type EnumDecl struct {
	Doc      *Metadata // associated documentation; or nil
	Modifier *Modifier
	Name     *Ident // type name
	List     []*EnumStmt
}

func (e *EnumDecl) Pos() int { return e.Name.Pos() }

func (*EnumDecl) declNode() {}

func (e *EnumDecl) Print(buffer *bytes.Buffer, indent int, onlyDeclare bool) {
	WriteIndent(buffer, indent)
	buffer.WriteString("enum class ")
	e.Name.Print(buffer)
	buffer.WriteString("\n")
	WriteIndent(buffer, indent)
	buffer.WriteString("{\n")
	for _, item := range e.List {
		item.Print(buffer, indent+4)
	}
	buffer.WriteString("};\n")
}

type InterfaceDecl struct {
	Doc       *Metadata // associated documentation; or nil
	Modifier  *Modifier
	Name      *Ident // type name
	Generic   *GenericLit
	Values    []*ValueDecl
	Functions []*FuncDecl // position of '=', if any
}

func (c *InterfaceDecl) Pos() int { return c.Name.Pos() }

func (*InterfaceDecl) declNode() {}

func (i *InterfaceDecl) Print(buffer *bytes.Buffer, indent int, onlyDeclare bool) {
	WriteIndent(buffer, indent)
	buffer.WriteString("class ")
	i.Name.Print(buffer)
	buffer.WriteString("\n")
	WriteIndent(buffer, indent)
	buffer.WriteString("{\n")
	buffer.WriteString("public:\n")
	for _, f := range i.Functions {
		f.Print(buffer, indent+TabSize, true)
		buffer.WriteString("\n")
	}
	buffer.WriteString("};\n")
}

// A FuncDecl node represents a function declaration.
type FuncDecl struct {
	Doc           *Metadata // associated documentation; or nil
	Modifier      *Modifier
	Name          *Ident     // function/method name
	Params        *FieldList // (incoming) parameters; non-nil
	Result        *Field     // (outgoing) results; or nil
	Body          *BlockStmt // function body; or nil for external (non-Go) function
	Generic       *GenericLit
	IsMember      bool
	IsConstructor bool
	IsDestructor  bool
}

func (d *FuncDecl) Pos() int { return d.Name.Pos() }

func (*FuncDecl) declNode() {}

func (f *FuncDecl) Print(buffer *bytes.Buffer, indent int, onlyDeclare bool) {
	//TO-DO indent
	if f.Generic != nil {
		WriteIndent(buffer, indent)
		//template <class T, int N>
		buffer.WriteString("template <")
		for i, t := range f.Generic.Types {
			if i > 0 {
				buffer.WriteString(", ")
			}
			buffer.WriteString("class ")
			t.Print(buffer)
		}
		buffer.WriteString(">\n")
	}
	WriteIndent(buffer, indent)
	if f.IsMember && !f.IsConstructor {
		buffer.WriteString("virtual ")
	}
	if !(f.IsMember && (f.IsDestructor || f.IsConstructor)) {
		f.Result.Type.Print(buffer)
		buffer.WriteString(" ")
	}
	f.Name.Print(buffer)
	buffer.WriteString("(")
	f.Params.Print(buffer)
	buffer.WriteString(")")
	if onlyDeclare {
		if f.IsMember {
			buffer.WriteString(" = 0")
		}
		buffer.WriteString(";\n")
	} else {
		buffer.WriteString("\n")
		f.Body.Print(buffer, indent)
	}
}

// ----------------------------------------------------------------------------
// Files and packages
type ProgramFile struct {
	Namespace  *NamespaceDecl // position of "namespace" keyword
	Imports    []*ImportDecl  // imports in this file
	Values     []*ValueDecl
	Functions  []*FuncDecl
	Classes    []*ClassDecl
	Enums      []*EnumDecl
	Interfaces []*InterfaceDecl
	EndPos     int
	Unresolved []*Ident // unresolved identifiers in this file

	//TO-DO add other program file as children
	//one namespace has one program file // use tree to store
}

func (f *ProgramFile) Pos() int { return f.Namespace.Pos() }

func (f *ProgramFile) End() int {
	return f.EndPos
}
func (f *ProgramFile) Print(buffer *bytes.Buffer, header bool) {
	if header {
		buffer.WriteString("#include <cinttypes>\n")
		buffer.WriteString("#include <iostream>\n")
		buffer.WriteString("#include <string>\n\n")
	}

	for _, v := range f.Functions {
		v.Print(buffer, 0, true)
		buffer.WriteString("\n")
	}

	for _, v := range f.Enums {
		v.Print(buffer, 0, true)
		buffer.WriteString("\n")
	}

	for _, v := range f.Interfaces {
		v.Print(buffer, 0, true)
		buffer.WriteString("\n")
	}

	for _, v := range f.Classes {
		v.Print(buffer, 0, true)
		buffer.WriteString("\n")
	}

	for _, v := range f.Values {
		v.Print(buffer, 0, false)
		buffer.WriteString("\n")
	}

	for _, v := range f.Functions {
		v.Print(buffer, 0, false)
		buffer.WriteString("\n")
	}

	for _, v := range f.Classes {
		v.Print(buffer, 0, false)
		buffer.WriteString("\n")
	}
}

// A Package node represents a set of source files
// collectively building a Go package.
//
type ProgramPackage struct {
	Name    string                  // package name
	Scope   *Scope                  // package scope across all files
	Imports map[string]*Object      // map of package id -> package object
	Files   map[string]*ProgramFile // Go source files by filename
}

func (p *ProgramPackage) Pos() int { return 0 }

// A Scope maintains the set of named language entities declared
// in the scope and a link to the immediately surrounding (outer)
// scope.
//
type Scope struct {
	Outer   *Scope
	Objects map[string]*Object
}

// NewScope creates a new scope nested in the outer scope.
func NewScope(outer *Scope) *Scope {
	const n = 4 // initial scope capacity
	return &Scope{outer, make(map[string]*Object, n)}
}

// Find returns the object with the given name if it is
// found in scope s, otherwise it returns nil. Outer scopes
// are ignored.
//
func (s *Scope) Find(name string) *Object {
	return s.Objects[name]
}

// Insert attempts to insert a named object obj into the scope s.
// If the scope already contains an object alt with the same name,
// Insert leaves the scope unchanged and returns alt. Otherwise
// it inserts obj and returns nil.
//
func (s *Scope) Insert(obj *Object) (alt *Object) {
	if alt = s.Objects[obj.Name]; alt == nil {
		s.Objects[obj.Name] = obj
	}
	return
}

// Debugging support
func (s *Scope) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "scope %p {", s)
	if s != nil && len(s.Objects) > 0 {
		fmt.Fprintln(&buf)
		for _, obj := range s.Objects {
			fmt.Fprintf(&buf, "\t%s %s\n", obj.Kind, obj.Name)
		}
	}
	fmt.Fprintf(&buf, "}\n")
	return buf.String()
}

// ----------------------------------------------------------------------------
// Objects

// An Object describes a named language entity such as a package,
// constant, type, variable, function (incl. methods), or label.
//
type Object struct {
	Kind ObjKind
	Name string      // declared name
	Decl interface{} // corresponding Field, XxxSpec, FuncDecl, AssignStmt, Scope; or nil
	Type interface{} // placeholder for type information; may be nil
}

// NewObj creates a new object of a given kind and name.
func NewObject(kind ObjKind, name string) *Object {
	return &Object{Kind: kind, Name: name}
}

// Pos computes the source position of the declaration of an object name.
// The result may be an invalid position if it cannot be computed
// (obj.Decl may be nil or not correct).
func (obj *Object) Pos() int {
	name := obj.Name
	switch d := obj.Decl.(type) {
	case *Field:
		if d.Name.Name == name {
			return d.Pos()
		}
	case *NamespaceDecl:
		return d.Pos()
	case *ImportDecl:
		if d.Name != nil && d.Name.Name == name {
			return d.Pos()
		}
		return d.Path.Pos()
	case *ValueDecl:
		if d.Name.Name == name {
			return d.Name.Pos()
		}
	//TO-DO class enum interface
	case *FuncDecl:
		if d.Name.Name == name {
			return d.Name.Pos()
		}
	case *AssignStmt:
		if ident, isIdent := d.Left.(*Ident); isIdent && ident.Name == name {
			return ident.Pos()
		}
	}
	return 0
}

// ObjKind describes what an object represents.
type ObjKind int

// The list of possible Object kinds.
const (
	Bad          ObjKind = iota // for error handling
	ConstObj                    // constant
	VarObj                      // variable
	ClassObj                    // class
	InterfaceObj                //interface
	EnumObj                     // enum
	FunctionObj                 // function or method
)

var objKindStrings = [...]string{
	Bad:         "bad",
	ConstObj:    "const",
	VarObj:      "var",
	ClassObj:    "class",
	EnumObj:     "enum",
	FunctionObj: "function",
}

func (kind ObjKind) String() string { return objKindStrings[kind] }
