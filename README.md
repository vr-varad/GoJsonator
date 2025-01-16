# JSON Parser in Go

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

