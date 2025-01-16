package main

func GoJsonator(input string) ASTNode {
	return Parser(Tokenizer(input))
}

func main() {
	ans := GoJsonator(`{
  "name": "John Doe",
  "age": 30,
  "isStudent": false,
  "courses": ["Go", "JSON Parsing"]
}
`)
	TraverseAST(ans)
}
