package main

type TokenType string

const (
	BracesOpen   TokenType = "BracesOpen"
	BracesClose  TokenType = "BracesClose"
	Comma        TokenType = "Comma"
	Colon        TokenType = "Colon"
	Number       TokenType = "Number"
	String       TokenType = "String"
	Bool         TokenType = "Bool"
	Null         TokenType = "Null"
	BracketOpen  TokenType = "BracketOpen"
	BracketClose TokenType = "BracketClose"
	True         TokenType = "True"
	False        TokenType = "False"
)

type Token struct {
	Type  TokenType
	Value string
}

type ASTNode interface{}

type ObjectType struct {
	Value map[string]ASTNode
}

type ArrayType struct {
	Value []ASTNode
}

type StringType struct {
	Value string
}

type NumberType struct {
	Value float64
}

type BooleanType struct {
	Value bool
}

type NullType struct{}
