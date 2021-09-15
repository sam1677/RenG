package ast

import (
	"RenG/interpreter/token"
	"bytes"
)

type StringLiteral struct {
	Token  token.Token
	Value  string
	Values []string
	Exp    []Expression
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string {
	var out bytes.Buffer

	out.WriteString("\"" + sl.Value)

	for i := 0; i < len(sl.Values); i++ {
		out.WriteString("[" + sl.Exp[i].String() + "]")
		out.WriteString(sl.Values[i])
	}

	out.WriteString("\"")

	return out.String()
}
