package main

import (
	"fmt"
	"regexp"
)

func Tokenizer(input string) []Token {
	tokens := []Token{}

	for i := 0; i < len(input); i++ {
		char := string(input[i])
		switch char {
		case "{":
			newToken := Token{
				Type:  BracesOpen,
				Value: "{",
			}
			tokens = append(tokens, newToken)
		case "}":
			newToken := Token{
				Type:  BracesClose,
				Value: "}",
			}
			tokens = append(tokens, newToken)
		case "[":
			newToken := Token{
				Type:  BracketOpen,
				Value: "[",
			}
			tokens = append(tokens, newToken)
		case "]":
			newToken := Token{
				Type:  BracketClose,
				Value: "]",
			}
			tokens = append(tokens, newToken)
		case ":":
			newToken := Token{
				Type:  Colon,
				Value: ":",
			}
			tokens = append(tokens, newToken)
		case ",":
			newToken := Token{
				Type:  Comma,
				Value: ",",
			}
			tokens = append(tokens, newToken)
		case "\"":
			value := ""
			i += 1
			char = string(input[i])
			for char != "\"" {
				value = value + string(char)
				i += 1
				char = string(input[i])
			}
			newToken := Token{
				Type:  String,
				Value: value,
			}
			tokens = append(tokens, newToken)
		default:
			numorBool, err := regexp.Compile(`^[a-zA-Z0-9]+$`)
			if err != nil {
				fmt.Print("Unexpected token")
				return nil
			}
			space, err := regexp.Compile(`\s`)
			if err != nil {
				fmt.Print("Unexpected token")
				return nil
			}
			if numorBool.MatchString(string(char)) {
				value := ""
				for numorBool.MatchString(string(char)) {
					value += string(char)
					i += 1
					char = string(input[i])
				}
				i -= 1
				if IsNumber(value) {
					newToken := Token{
						Type:  Number,
						Value: value,
					}
					tokens = append(tokens, newToken)
				} else if IsBooleanTrue(value) {
					newToken := Token{
						Type:  True,
						Value: value,
					}
					tokens = append(tokens, newToken)
				} else if IsBooleanFalse(value) {
					newToken := Token{
						Type:  False,
						Value: value,
					}
					tokens = append(tokens, newToken)
				} else if IsNull(value) {
					newToken := Token{
						Type:  Null,
						Value: value,
					}
					tokens = append(tokens, newToken)
				}
			} else if space.MatchString(string(char)) {
				continue
			} else {
				fmt.Print("Unexpected Character")
				return nil
			}

		}

	}
	return tokens
}
