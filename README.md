# JSON Parser in Go

![GoJson](https://www.google.com/imgres?q=GO%20Json&imgurl=https%3A%2F%2Fuser-images.githubusercontent.com%2F209884%2F92572337-42b42900-f2bf-11ea-973a-c74a359553a5.png&imgrefurl=https%3A%2F%2Fgithub.com%2Fgoccy%2Fgo-json&docid=OEM91_JPlDxi5M&tbnid=4sTsqlteMn9mTM&vet=12ahUKEwi-uPnA_vmKAxWzb_UHHaOXKnYQM3oECBYQAA..i&w=500&h=533&hcb=2&ved=2ahUKEwi-uPnA_vmKAxWzb_UHHaOXKnYQM3oECBYQAA)

A simple and efficient JSON parser written in Go. This parser takes JSON input and converts it into a structured Abstract Syntax Tree (AST) that can be further processed, analyzed, or manipulated.

## Features
- Parses JSON objects, arrays, strings, numbers, booleans, and null values.
- Supports nested objects and arrays.
- Converts parsed JSON data into Go structs for easy manipulation.

## Installation

   ```bash
   git clone https://github.com/vr-varad/GoJsonator.git
   cd json-parser-go
   go build
   ./main
   ```

## Example

### Input
```
{
  "name": "John Doe",
  "age": 30,
  "isStudent": false,
  "courses": ["Go", "JSON Parsing"]
}
```

### Output
```
Object:
Key: isStudent, Value: Boolean: false
Key: courses, Value: Array:
Index 0: String: Go
Index 1: String: JSON Parsing
Key: name, Value: String: John Doe
Key: age, Value: Number: 30
```

## How It Works
1. The parser breaks the JSON string into tokens.
2. It constructs an AST by recursively processing the tokens
mapping them to their respective Go types (e.g., String, Number, Boolean).
3. The result is an AST that represents the structure of the JSON data.


### Flow Diagram

Here's a simple flow for your JSON parser:

```plaintext
+-------------------+
|  Start Parsing    |
+-------------------+
        |
        v
+--------------------+
|  Tokenize Input    |
| (String, Number,   |
|  Braces, etc.)     |
+--------------------+
        |
        v
+--------------------+
|  Parse Tokens      |
| (Objects, Arrays,  |
|  Strings, Numbers) |
+--------------------+
        |
        v
+--------------------+
|  Construct AST     |
| (Nested Structures)|
+--------------------+
        |
        v
+--------------------+
|  Output AST        |
+--------------------+

