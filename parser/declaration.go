package parser

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/token"
)

func (p *Parser) parseVariable() *ast.Variable {
	d := &ast.Variable{}
	d.Name = p.parseIdentifier()
	d.Type = p.parseType()

	if p.token == token.Assign {
		p.next()
		if !p.token.IsLiteral() || p.token == token.IDENT {
			p.error(p.position, "variable can only be initialized by const value (string, char, float, int)")
		}
		d.Value = &ast.Literal{
			Position: p.position,
			Type:     p.token,
			Value:    p.literal,
		}
		p.next()
	}
	p.expect(token.Semi)
	return d
}

func (p *Parser) parseFunction() *ast.Function {
	return nil
}

func (p *Parser) parseEnum() *ast.Enum {
	d := &ast.Enum{
		Members: make(map[string]*ast.Variable),
	}
	d.Name = p.parseIdentifier()
	p.expect(token.LeftBrace)
	for p.token != token.RightBrace {
		v := &ast.Variable{}
		v.Name = p.parseIdentifier()
		if p.token == token.Assign {
			p.next()
			if p.token != token.INT {
				p.error(p.position, "value can only be int")
			}
			v.Value = &ast.Literal{
				Position: p.position,
				Type:     p.token,
				Value:    p.literal,
			}
			p.next()
		}
		if p.token != token.Comma {
			break
		}
		p.next()
	}
	p.expect(token.RightBrace)
	return d
}

func (p *Parser) parseInterface() *ast.Interface {
	d := &ast.Interface{
		Functions: make(map[string]*ast.Function),
	}
	d.Name = p.parseIdentifier()
	if p.token == token.Less {
		d.TypeParameters = p.parseTypeParameters()
	}
	if p.token == token.Colon {
		d.Base = p.parseBaseTypes()
	}
	p.expect(token.LeftBrace)
	for p.token != token.RightBrace {
		m := p.parseMetadata()
		f := p.parseFunction()
		f.Custom = append(f.Custom, m...)
		name := f.Name.Name
		if _, ok := d.Functions[name]; ok {
			p.error(f.Name.Position, fmt.Sprintf("function %s redeclared", name))
		}
		d.Functions[name] = f
	}
	p.expect(token.RightBrace)
	return d
}

func (p *Parser) parseClass() *ast.Class {
	d := &ast.Class{
		Variables: make(map[string]*ast.Variable),
		Functions: make(map[string]*ast.Function),
	}
	d.Name = p.parseIdentifier()
	if p.token == token.Less {
		d.TypeParameters = p.parseTypeParameters()
	}
	if p.token == token.Colon {
		d.Base = p.parseBaseTypes()
	}
	p.expect(token.LeftBrace)
	for p.token != token.RightBrace {
		m := p.parseMetadata()
		modifier := p.parseModifier()
		switch p.token {
		case token.Const, token.Var:
			p.next()
			v := p.parseVariable()
			v.Custom = append(v.Custom, m...)
			v.Modifier = modifier
			name := v.Name.Name
			if _, ok := d.Variables[name]; ok {
				p.error(v.Name.Position, fmt.Sprintf("variable %s redeclared", name))
			}
			d.Variables[name] = v

		case token.Function:
			p.next()
			f := p.parseFunction()
			f.Custom = append(f.Custom, m...)
			f.Modifier = modifier
			name := f.Name.Name
			if _, ok := d.Functions[name]; ok {
				p.error(f.Name.Position, fmt.Sprintf("function %s redeclared", name))
			}
			d.Functions[name] = f

		default:
			p.unexpected(p.position, "declaration")
		}
	}
	p.expect(token.RightBrace)
	return d
}

/*
func (p *Parser) parseFuncDecl(m *Modifier, onlyDeclare bool) *FuncDecl {
	p.expect(Function)

	//Tilde
	tilde := false
	if p.tok == Tilde {
		tilde = true
		p.next()
	}
	ident := p.parseIdent()
	if tilde {
		//TO-DO check in class
		ident.Name = "~" + ident.Name
	}
	generic := p.parseGeneric()
	params := p.parseParameters()
	result := p.parseResult()

	decl := &FuncDecl{
		Modifier: m,
		Name:     ident,
		Params:   params,
		Result:   result,
		Generic:  generic,
	}

	if onlyDeclare {
		//TO-DO check later call.delare ?
		return decl
	}

	if p.tok == LeftBrace {
		decl.Body = p.parseBody()
	}

	return decl
}
*/
