package parser

import (
	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/token"
)

func (p *Parser) parseIdentifier() *ast.Identifier {
	position := p.position
	name := ""
	if p.token == token.IDENT {
		name = p.literal
		p.next()
	} else {
		p.expect(token.IDENT)
	}
	return &ast.Identifier{
		Position: position,
		Name:     name,
	}
}

/*
// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------
func (p *Parser) parseType() Expr {
	typ := p.tryType()

	if typ == nil {
		pos := p.pos
		p.errorExpected(pos, "type")
		return &BadExpr{Start: pos}
	}

	return typ
}

// If the result is an identifier, it is not resolved.
func (p *Parser) parseTypeName() Expr {
	ident := p.parseIdent()
	// don't resolve ident yet - it may be a parameter or field name
	if p.tok == Dot {
		// ident is a package name
		p.next()
		sel := p.parseIdent()
		return &SelectorExpr{Expr: ident, Selector: sel}
	}
	return ident
}

// If the result is an identifier, it is not resolved.
func (p *Parser) tryVarType(isParam bool) Expr {
	if isParam && p.tok == Ellipsis {
		pos := p.pos
		p.next()
		typ := p.tryType() // don't use parseType so we can provide better error message
		if typ == nil {
			p.error(pos, "'...' parameter is missing type")
			typ = &BadExpr{Start: pos}
		}
		return &EllipsisLit{Start: pos, Expr: typ}
	}
	return p.tryType()
}

// If the result is an identifier, it is not resolved.
func (p *Parser) parseVarType(isParam bool) Expr {
	typ := p.tryVarType(isParam)
	if typ == nil {
		pos := p.pos
		p.errorExpected(pos, "type")
		p.next() // make progress
		typ = &BadExpr{Start: pos}
	}
	return typ
}

func (p *Parser) parseParameterList() (params []*Field) {
	for p.tok != RightParen {
		field := &Field{}
		typ := p.parseVarType(true)
		if p.tok == Comma {
			field.Name = nil
			field.Type = typ
			p.expect(Comma)
			params = append(params, field)
			continue
		}
		ok := false
		field.Name, ok = typ.(*Ident)
		if !ok {
			p.expect(IDENT)
		}
		field.Type = p.parseVarType(true)
		if p.tok == Equal {
			p.next()
			field.Default = p.parseExpr(false)
		}
		params = append(params, field)
		if p.tok != RightParen {
			p.expect(Comma)
		}
	}
	//TP-DO ...
	//TO-DO check default
	return
}

func (p *Parser) parseParameters() *FieldList {
	var params []*Field
	start := p.expect(LeftParen)
	if p.tok != RightParen {
		params = p.parseParameterList()
	}
	p.expect(RightParen)

	return &FieldList{Start: start, Fields: params}
}

func (p *Parser) parseResult() *Field {
	typ := p.tryType()
	if typ != nil {
		return &Field{Type: typ}
	}
	return &Field{Type: &Scalar{Token: Void}}
}

// If the result is an identifier, it is not resolved.
func (p *Parser) tryType() Expr {
	if p.tok.IsScalar() {
		scalar := &Scalar{
			Start: p.pos,
			Token: p.tok,
		}
		p.next()
		return scalar

	} else if p.tok == IDENT {
		typ := p.parseTypeName()
		return typ
	}
	return nil
}
*/
/*
func (p *Parser) parseGeneric() *GenericLit {
	if p.tok == Less {
		g := &GenericLit{
			Start: p.pos,
		}
		p.next()
		g.Types = append(g.Types, p.tryType())

		for p.tok == Comma {
			p.next()
			g.Types = append(g.Types, p.tryType())
		}
		p.expect(Greater)
		return g

	}
	return nil
}
*/
