package main

import (
	"fmt"
	"os"
	"strconv"
)

func parseValue(tokens []Token, index *int) ASTNode {
	token := tokens[*index]
	switch token.Type {
	case String:
		return StringType{
			Value: token.Value,
		}

	case Number:
		val, _ := strconv.Atoi(token.Value)
		return NumberType{
			Value: float64(val),
		}
	case Null:
		return NullType{}
	case True:
		return BooleanType{
			Value: true,
		}
	case False:
		return BooleanType{
			Value: false,
		}
	case BracesOpen:
		return parseObject(tokens, index)
	case BracketOpen:
		return parseArray(tokens, index)
	default:
		fmt.Print("Unexpexted Error")
		os.Exit(0)
	}
	return nil
}

func parseObject(tokens []Token, index *int) ObjectType {
	obj :=
		ObjectType{
			Value: map[string]ASTNode{},
		}
	*index++
	token := tokens[*index]
	for token.Type != BracesClose {
		if token.Type == String {
			key := tokens[*index].Value
			*index++
			token = tokens[*index]
			if token.Type != Colon {
				fmt.Print("Expected : in key-value pair")
				os.Exit(0)
			}
			*index++
			token = tokens[*index]
			value := parseValue(tokens, index)
			obj.Value[key] = value
			*index++
			token = tokens[*index]
		} else {
			fmt.Print("Expected : in key-value pair")
			os.Exit(0)
		}

		if token.Type == Comma {
			*index++
			token = tokens[*index]
		}
	}
	return obj
}

func parseArray(tokens []Token, index *int) ArrayType {
	array := ArrayType{
		Value: []ASTNode{},
	}
	*index++
	token := tokens[*index]
	for token.Type != BracketClose {
		if token.Type != Comma {
			value := parseValue(tokens, index)
			*index++
			token = tokens[*index]
			array.Value = append(array.Value, value)
		} else if token.Type == Comma {
			*index++
			token = tokens[*index]
			continue
		} else {
			fmt.Print("Expected : in key-value pair")
			os.Exit(0)
		}
	}
	return array
}

func Parser(tokens []Token) ASTNode {
	if len(tokens) == 0 {
		os.Exit(0)
	}

	index := 0

	AST := parseValue(tokens, &index)

	return AST
}
