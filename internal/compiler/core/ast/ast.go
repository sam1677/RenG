package ast

import (
	"bytes"
	"internal/compiler/core/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

/*
Statement

- ExpressionStatement
- BlockStatement
- IfStatement
- ForStatement
- ReturnStatement
*/

type Statement interface {
	Node
	statementNode()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type FunctionStatement struct {
	Token      token.Token
	Name       *Identifier
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fs *FunctionStatement) statementNode()       {}
func (fs *FunctionStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *FunctionStatement) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fs.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fs.TokenLiteral() + " ")
	out.WriteString(fs.Name.String())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(") {")
	out.WriteString(fs.Body.String())
	out.WriteString(" }")

	return out.String()
}

type IfStatement struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Elif        []*IfStatement
	Else        *BlockStatement
}

func (is *IfStatement) statementNode() {}
func (is *IfStatement) TokenLiteral() string {
	var out bytes.Buffer

	out.WriteString("if " + is.Condition.String())
	out.WriteString(" { " + is.Consequence.String() + " }")

	for _, ee := range is.Elif {
		if ee != nil {
			out.WriteString(" elif ")
			out.WriteString(ee.Condition.String())
			out.WriteString(" ")
			out.WriteString(is.Consequence.String())
		}
	}

	if is.Else != nil {
		out.WriteString("else ")
		out.WriteString(is.Else.String())
	}

	return out.String()
}
func (is *IfStatement) String() string {
	var out bytes.Buffer

	out.WriteString("if " + is.Condition.String())
	out.WriteString(" { " + is.Consequence.String() + " }")

	for _, ee := range is.Elif {
		if ee != nil {
			out.WriteString(" elif ")
			out.WriteString(ee.Condition.String())
			out.WriteString(" ")
			out.WriteString("{ ")
			out.WriteString(is.Consequence.String())
			out.WriteString(" }")
		}
	}

	if is.Else != nil {
		out.WriteString("else ")
		out.WriteString(is.Else.String())
	}

	return out.String()
}

type ForStatement struct {
	Token          token.Token
	Initialization Expression
	Condition      Expression
	Increment      Expression
	Loop           *BlockStatement
}

func (fs *ForStatement) statementNode()       {}
func (fs *ForStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *ForStatement) String() string {
	var out bytes.Buffer

	out.WriteString("for ( ")

	if fs.Initialization != nil {
		out.WriteString(fs.Initialization.String() + "; ")
	} else {
		out.WriteString("; ")
	}

	if fs.Condition != nil {
		out.WriteString(fs.Condition.String() + "; ")
	} else {
		out.WriteString("; ")
	}

	if fs.Increment != nil {
		out.WriteString(fs.Increment.String() + ") ")
	} else {
		out.WriteString(") ")
	}

	out.WriteString("{ " + fs.Loop.String() + " }")

	return out.String()
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(" ")

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	return out.String()
}

type LabelStatement struct {
	Token token.Token
	Name  *Identifier
	Body  *BlockStatement
}

func (ls *LabelStatement) statementNode()       {}
func (ls *LabelStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LabelStatement) String() string {
	var out bytes.Buffer

	out.WriteString("label ")
	out.WriteString(ls.Name.String())
	out.WriteString(" { ")
	out.WriteString(ls.Body.String())
	out.WriteString(" } ")

	return out.String()
}

type ScreenStatement struct {
	Token token.Token
	Name  *Identifier
	Body  *BlockStatement
}

func (ss *ScreenStatement) statementNode()       {}
func (ss *ScreenStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *ScreenStatement) String() string {
	var out bytes.Buffer

	out.WriteString("screen ")
	out.WriteString(ss.Name.String())
	out.WriteString(" { ")
	out.WriteString(ss.Body.String())
	out.WriteString(" } ")

	return out.String()
}

/*
Expression

-PrefixExpression
-InfixExpression
-PostfixExpression
-IndexExpression
-CallFunctionExpression
*/

type Expression interface {
	Node
	expressionNode()
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("] ")

	return out.String()
}

type CallFunctionExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (cfe *CallFunctionExpression) expressionNode()      {}
func (cfe *CallFunctionExpression) TokenLiteral() string { return cfe.Token.Literal }
func (cfe *CallFunctionExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range cfe.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(cfe.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ","))
	out.WriteString(")")

	return out.String()
}

/*
literal

-Identifier
-Boolean
-IntegerLiteral
-FloatLiteral
-StringLiteral
-ArrayLiteral
*/

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

type BooleanLiteral struct {
	Token token.Token
	Value bool
}

func (bl *BooleanLiteral) expressionNode()      {}
func (bl *BooleanLiteral) TokenLiteral() string { return bl.Token.Literal }
func (bl *BooleanLiteral) String() string       { return bl.Token.Literal }

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (fl *FloatLiteral) expressionNode()      {}
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FloatLiteral) String() string       { return fl.Token.Literal }

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return "\"" + sl.Token.Literal + "\"" }

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
