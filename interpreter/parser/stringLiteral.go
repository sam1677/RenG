package parser

import (
	"RenG/interpreter/ast"
	"RenG/interpreter/token"
)

func (p *Parser) parseStringLiteral() ast.Expression {
	for p.peekTokenIs(token.LBRACKET) {
		p.nextToken()
		p.nextToken()

	}
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

// 항상 STRING이 여러개 이상이면 그 사이에 표현식이 들어가야 함.
// 즉 예로는, "hello[a]word[b]" STRING + EXPRESSION + STRING + EXPRESSION + STRING 형식
// 한 번 이것을 이용해보자.
