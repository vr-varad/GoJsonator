package main

import (
	"fmt"
	"strconv"
)

func IsNumber(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

func IsBooleanTrue(input string) bool {
	return input == "true"
}

func IsBooleanFalse(input string) bool {
	return input == "false"
}

func IsNull(input string) bool {
	return input == "null"
}

func TraverseAST(node ASTNode) {
	switch v := node.(type) {
	case ObjectType:
		fmt.Println("Object:")
		for key, value := range v.Value {
			fmt.Printf("Key: %s, Value: ", key)
			TraverseAST(value)
		}
	case ArrayType:
		fmt.Println("Array:")
		for i, value := range v.Value {
			fmt.Printf("Index %d: ", i)
			TraverseAST(value)
		}
	case StringType:
		fmt.Println("String:", v.Value)
	case NumberType:
		fmt.Println("Number:", v.Value)
	case BooleanType:
		fmt.Println("Boolean:", v.Value)
	case NullType:
		fmt.Println("Null")
	default:
		fmt.Println("Unknown Type")
	}
}
