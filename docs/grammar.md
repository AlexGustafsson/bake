# Bake grammar

## Notation

This document describes the grammar of Bake using a similar notation to [Go](https://golang.org/ref/spec) and in turn  [EBNF](https://en.wikipedia.org/wiki/Extended_Backus–Naur_form).

The following (slightly modified) is taken from [the golang documentation](https://golang.org/ref/spec).

The syntax is specified using Extended Backus-Naur Form (EBNF):

```
Production  = production_name "=" [ Expression ] "." .
Expression  = Alternative { "|" Alternative } .
Alternative = Term { Term } .
Term        = production_name | token [ "…" token ] | Group | Option | Repetition .
Group       = "(" Expression ")" .
Option      = "[" Expression "]" .
Repetition  = "{" Expression "}" .
```

Productions are expressions constructed from terms and the following operators, in increasing precedence:

```
|   alternation
()  grouping
[]  option (0 or 1 times)
{}  repetition (0 to n times)
```

Lower-case production names are used to identify lexical tokens. Non-terminals are in CamelCase. Lexical tokens are enclosed in double quotes `""`.

Lastly, `...` denotes a range of characters.

## Grammar

### Source file

```
SourceFile = [ PackageDeclaration ] [ ImportsDeclaration ] { TopLevelDeclarations } .

PackageDeclaration = "package" identifier .

ImportsDeclaration = "import" "(" { string_literal } ")" .
```

### Declarations

```
TopLevelDeclaration = Declaration | FunctionDeclaration | RuleFunctionDeclaration | AliasDeclaration | RuleDeclaration .

Declaration = VarDeclaration .

VarDeclaration = "let" identifier [ "=" Expression ] .

FunctionDeclaration = [ "export" ] "func" identifier [Signature] Block .

RuleFunctionDeclaration = [ "export" ] "rule" identifier [Signature] Block .

AliasDeclaration = "alias" identifier ( file_path | FileList ) .

RuleDeclaration = ( file_path | FileList ) [ FileList ] ( ":" PrimaryExpression [ Block ] | Block ) .
FileList = "[" file_path { "," file_path } "]" .
file_path = string_literal .

Signature = "(" [ ParameterList ] ")" .
ParameterList = identifier { "," identifier } .
Block = "{" Statements "}" .
```

### Statements

```
Statements = Statement { "\n" Statement } .

Statement = Declaration
          | SimpleStatement
          | IfStatement
          | ForStatement
          .

SimpleStatement = EmptyStatement
                | ExpressionStatement
                | IncreaseDecreaseStatement
                | ShellStatement
                | Assignment
                .

EmptyStatement = .
ExpressionStatement = Expression .
IncreaseDecreaseStatement = Expression ( "++" | "--" ) .
ShellStatement = "shell" { unicode_char } .
Assignment = Expression assignment_operand ExpressionList .

assignment_operand = [ additive_operator | multiplicative_operator | "?" ] "=" .
```

### Expressions


```ebnf
Expression = UnaryExpression
           | Expression binary_operator Expression
           .

UnaryExpression = PrimaryExpression | binary_operator UnaryExpression .

binary_operator = "||" | "&&" | comparison_operator | additive_operator | multiplicative_operator .
comparison_operator = "==" | "!=" | "<" | "<=" | ">" | ">=" .
additive_operator = "+" | "-" | "|" | "^" .
multiplicative_operator = "*" | "/" | "%" | "<<" | ">>" | "&" .
unary_operator = "-" | "!" | "..." .

PrimaryExpression = Operand
                  | PrimaryExpression Selector
                  | PrimaryExpression ImportSelector
                  | PrimaryExpression Index
                  | PrimaryExpression Arguments
                  .

Operand = Literal | identifier | "(" Expression ")" .
Literal = integer_literal | string_literal .

Selector = "." identifier .
ImportSelector = "::" identifier .
Index = "[" Expression "]" .
Arguments = "(" [ ExpressionList ] ")" .
ExpressionList = Expression { "," Expression } .
```

## Tokens

#### Letters and digits

```
letter = "a" ... "z" .
unicode_char = /* an arbitrary Unicode code point except newline */
decimal_digit = "0" ... "9" .
```

### Identifier

An identifier is specified by a letter, followed by a sequence of zero or more alpha numeric characters.

```
identifier = letter { letter | decimal_digit}
```

### Keywords

```
package
import
func
rule
export
if
else
return
context
shell
let
```

### Operators

```
+
-
*
/
==
!=
...
```

### Punctuation

```
(
)
[
]
{
}
:
,
.
$
@
```

### Literals

#### Integers

Decimal integers are supported and may be written with one zero as a prefix.

```
integer_literal = decimal_literal .
decimal_literal = "0" | "1" ... "9" [ decimal_digits ] .
decimal_digits = decimal_digit { decimal_digit } .
```

#### Strings

```
string_literal = raw_string_literal | interpreted_string_literal .
raw_string_literal = "`" { unicode_char | newline } "`" .
interpreted_string_literal = `"` { variable_substitution | unicode_char } `"` .
variable_substitution = "$" "(" identifier ")" .
```
