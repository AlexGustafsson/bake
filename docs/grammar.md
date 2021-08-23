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
SourceFile = { Prolog } .

Prolog = "\n" | PackageDeclaration | ImportsDeclaration | TopLevelDeclaration .

PackageDeclaration = "package" identifier .

ImportsDeclaration = "import" "(" { string_literal } ")" .
```

### Declarations

```
TopLevelDeclaration = RuleFunctionDeclaration | AliasDeclaration | RuleDeclaration | Statement .

VarDeclaration = "let" identifier [ "=" Expression ] .

FunctionDeclaration = [ "export" ] "func" identifier [Signature] Block .

RuleFunctionDeclaration = [ "export" ] "rule" identifier [Signature] Block .

AliasDeclaration = [ "export" ] "alias" identifier ":" Expression .

RuleDeclaration = ( string_literal | Array ) [ Array ] ( ":" PrimaryExpression [ Block ] | Block ) .

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
          | ReturnStatement
          .

Declaration = FunctionDeclaration | VarDeclaration .

ReturnStatement = "return" Expression .

IfStatement = "if" Expression Block [ "else" ( IfStatement | Block ) ] .

SimpleStatement = ExpressionStatement
                | IncreaseDecreaseStatement
                | ShellStatement
                | Assignment
                .

ExpressionStatement = Expression .
IncreaseDecreaseStatement = Expression ( "++" | "--" ) .
ShellStatement = ShellLine | ShellBlock .
ShellLine = "shell" { unicode_char } .
ShellBlock = "shell" "{" { unicode_char } "}" . /* may contain matching pairs of "{" and "}" */
Assignment = Expression assignment_operand Expression .

assignment_operand = [ additive_operator | multiplicative_operator | "?" ] "=" .
```

### Expressions


```ebnf
Expression = Equality .

Equality = Comparison { equality_operator Comparison } .
equality_operator = "||" | "&&" .

Comparison = Term { comparison_operator Term } .
comparison_operator = "==" | "!=" | "<" | "<=" | ">" | ">=" .

Term = Factor { additive_operator Factor } .
additive_operator = "+" | "-" .

Factor = Unary { multiplicative_operator Unary } .
multiplicative_operator = "*" | "/" .

Unary = [ unary_operator ] Primary .
unary_operator = "-" | "!" | "..." .

Primary = Operand { Selector | Index | Arguments } .

Operand = Literal | Array | ImportSelector | identifier | "(" Expression ")" .
Literal = boolean_literal | integer_literal | string_literal .
Array = "[" [ ExpressionList ] "]" .
ImportSelector = identifier "::" identifier .

Selector = "." identifier .
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
identifier = letter { letter | decimal_digit} .
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
let
shell
true
false
alias
```

### Operators

```
+
-
*
/
=
==
!
!=
<
<=
>
>=
?=
&&
||
...
++
--
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

#### Booleans

```
boolean_literal = "true" | "false" .
```

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
variable_substitution = "$" "(" Expression ")" .
```

Note that the interpreted string literal is a runtime feature, meaning the parser and lexer treats it as a single token / node.

