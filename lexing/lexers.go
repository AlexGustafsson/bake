package lexing

import (
	"unicode"
)

func lexRoot(lexer *Lexer) stateModifier {
	switch rune := lexer.Peek(); rune {
	case ' ', '\t':
		// Ignore whitespace
		lexer.Next()
		lexer.Ignore()
		return lexRoot
	case '\n':
		lexer.Next()
		lexer.Emit(ItemNewline)
		return lexRoot
	case eof:
		lexer.Next()
		lexer.Emit(ItemEndOfInput)
		return nil
	case '/':
		lexer.Next()
		rune := lexer.Peek()
		if rune == '/' {
			lexer.Next()
			for {
				rune := lexer.Peek()
				if rune == '\n' || rune == eof {
					break
				} else {
					lexer.Next()
				}
			}
			lexer.Emit(ItemComment)
		} else if rune == '=' {
			lexer.Next()
			lexer.Emit(ItemDivisionAssign)
		} else {
			lexer.Emit(ItemDivision)
		}
		return lexRoot
	case '+':
		lexer.Next()
		rune := lexer.Peek()
		if rune == '+' {
			lexer.Next()
			lexer.Emit(ItemIncrement)
		} else if rune == '=' {
			lexer.Next()
			lexer.Emit(ItemAdditionAssign)
		} else {
			lexer.Emit(ItemAddition)
		}
		return lexRoot
	case '-':
		lexer.Next()
		rune := lexer.Peek()
		if rune == '-' {
			lexer.Next()
			lexer.Emit(ItemDecrement)
		} else if rune == '=' {
			lexer.Next()
			lexer.Emit(ItemSubtractionAssign)
		} else {
			lexer.Emit(ItemSubtraction)
		}
		return lexRoot
	case '*':
		lexer.Next()
		if rune := lexer.Peek(); rune == '=' {
			lexer.Next()
			lexer.Emit(ItemMultiplicationAssign)
		} else {
			lexer.Emit(ItemMultiplication)
		}
		return lexRoot
	case '%':
		lexer.Next()
		if rune := lexer.Peek(); rune == '=' {
			lexer.Next()
			lexer.Emit(ItemModuloAssign)
		} else {
			lexer.Emit(ItemModulo)
		}
		return lexRoot
	case '=':
		lexer.Next()
		if rune := lexer.Peek(); rune == '=' {
			lexer.Next()
			lexer.Emit(ItemEquals)
			return lexRoot
		} else {
			lexer.Emit(ItemAssignment)
			return lexRoot
		}
	case '!':
		lexer.Next()
		if rune := lexer.Peek(); rune == '=' {
			lexer.Next()
			lexer.Emit(ItemNotEqual)
			return lexRoot
		} else {
			lexer.Emit(ItemNot)
			return lexRoot
		}
	case '<':
		lexer.Next()
		if rune := lexer.Peek(); rune == '=' {
			lexer.Next()
			lexer.Emit(ItemLessThanOrEqual)
			return lexRoot
		} else {
			lexer.Emit(ItemLessThan)
			return lexRoot
		}
	case '>':
		lexer.Next()
		if rune := lexer.Peek(); rune == '=' {
			lexer.Next()
			lexer.Emit(ItemGreaterThanOrEqual)
			return lexRoot
		} else {
			lexer.Emit(ItemGreaterThan)
			return lexRoot
		}
	case '?':
		lexer.Next()
		if rune := lexer.Peek(); rune == '=' {
			lexer.Next()
			lexer.Emit(ItemLooseAssignment)
			return lexRoot
		} else {
			lexer.errorf("unexpected token '%c'", rune)
			return nil
		}
	case '&':
		lexer.Next()
		if rune := lexer.Peek(); rune == '&' {
			lexer.Next()
			lexer.Emit(ItemAnd)
			return lexRoot
		} else {
			lexer.errorf("unexpected token '%c'", rune)
			return nil
		}
	case '|':
		lexer.Next()
		if rune := lexer.Peek(); rune == '|' {
			lexer.Next()
			lexer.Emit(ItemOr)
			return lexRoot
		} else {
			lexer.errorf("unexpected token '%c'", rune)
			return nil
		}
	case '.':
		lexer.Next()
		if rune := lexer.Peek(); rune == '.' {
			lexer.Next()
			if rune := lexer.Peek(); rune == '.' {
				lexer.Next()
				lexer.Emit(ItemSpread)
				return lexRoot
			} else {
				lexer.errorf("unexpected token '%c'", rune)
				return nil
			}
		} else {
			lexer.Emit(ItemDot)
			return lexRoot
		}
	case '(':
		lexer.Next()
		lexer.Emit(ItemLeftParentheses)
		if lexer.Mode == ModeEvaluatedString || lexer.Mode == ModeShellString || lexer.Mode == ModeMultilineShellString {
			lexer.parenthesesDepth++
		}
		return lexRoot
	case ')':
		lexer.Next()
		if lexer.parenthesesDepth == 0 && (lexer.Mode == ModeEvaluatedString || lexer.Mode == ModeShellString || lexer.Mode == ModeMultilineShellString) {
			lexer.Emit(ItemSubstitutionEnd)
			lexer.substitutionDepth--
			if lexer.substitutionDepth < 0 {
				lexer.substitutionDepth = 0
			}

			next := lexRoot
			if lexer.Mode == ModeEvaluatedString {
				next = lexEvaluatedString
			} else if lexer.Mode == ModeShellString {
				next = lexShellString
			} else if lexer.Mode == ModeMultilineShellString {
				next = lexMultilineShellString
			}

			if lexer.substitutionDepth == 0 {
				lexer.Mode = ModeRoot
			}
			return next
		} else {
			lexer.Emit(ItemRightParentheses)
		}
		lexer.parenthesesDepth--
		if lexer.parenthesesDepth < 0 {
			lexer.parenthesesDepth = 0
		}
		return lexRoot
	case '[':
		lexer.Next()
		lexer.Emit(ItemLeftBracket)
		return lexRoot
	case ']':
		lexer.Next()
		lexer.Emit(ItemRightBracket)
		return lexRoot
	case '{':
		lexer.Next()
		lexer.Emit(ItemLeftCurly)
		return lexRoot
	case '}':
		lexer.Next()
		lexer.Emit(ItemRightCurly)
		return lexRoot
	case ':':
		lexer.Next()
		if rune := lexer.Peek(); rune == ':' {
			lexer.Next()
			lexer.Emit(ItemColonColon)
			return lexRoot
		} else {
			lexer.Emit(ItemColon)
			return lexRoot
		}
	case ',':
		lexer.Next()
		lexer.Emit(ItemComma)
		return lexRoot
	case '`':
		lexer.Next()
		lexer.Emit(ItemBacktick)
		for {
			rune := lexer.Peek()
			if rune == '`' {
				lexer.Emit(ItemStringPart)
				lexer.Next()
				lexer.Emit(ItemBacktick)
				break
			} else if rune == eof {
				lexer.errorf("unexpected end of file")
				return nil
			} else {
				lexer.Next()
			}
		}
		return lexRoot
	case '"':
		lexer.Next()
		lexer.Emit(ItemDoubleQuote)
		return lexEvaluatedString
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		lexer.Next()
		for {
			rune := lexer.Peek()
			if unicode.IsDigit(rune) {
				lexer.Next()
			} else {
				break
			}
		}
		lexer.Emit(ItemInteger)
		return lexRoot
	default:
		// Parse words such as "package" and "import" or identifiers
		if unicode.IsLetter(rune) {
			lexer.Next()
			word := string(rune)
			for {
				rune = lexer.Peek()
				if unicode.IsLetter(rune) || unicode.IsDigit(rune) || rune == '_' {
					lexer.Next()
					word += string(rune)
				} else {
					break
				}
			}

			if len(word) > 0 {
				switch word {
				case "package":
					lexer.Emit(ItemKeywordPackage)
				case "import":
					lexer.Emit(ItemKeywordImport)
				case "func":
					lexer.Emit(ItemKeywordFunc)
				case "rule":
					lexer.Emit(ItemKeywordRule)
				case "export":
					lexer.Emit(ItemKeywordExport)
				case "if":
					lexer.Emit(ItemKeywordIf)
				case "for":
					lexer.Emit(ItemKeywordFor)
				case "in":
					lexer.Emit(ItemKeywordIn)
				case "else":
					lexer.Emit(ItemKeywordElse)
				case "return":
					lexer.Emit(ItemKeywordReturn)
				case "break":
					lexer.Emit(ItemKeywordBreak)
				case "let":
					lexer.Emit(ItemKeywordLet)
				case "shell":
					lexer.Emit(ItemKeywordShell)
					return lexShell
				case "true", "false":
					lexer.Emit(ItemBoolean)
				case "alias":
					lexer.Emit(ItemKeywordAlias)
				default:
					lexer.Emit(ItemIdentifier)
				}

				return lexRoot
			}
		}

		lexer.errorf("unexpected token '%c'", rune)
		return nil
	}
}

func lexShell(lexer *Lexer) stateModifier {
	// Require at least one whitespace after the shell keyword. This solves cases
	// such as when "shell" is used as an identifier (context.shell.stdout.string)
	if rune := lexer.Peek(); rune != ' ' && rune != '\t' {
		return lexRoot
	}

	// Consume leading whitespace
	for {
		rune := lexer.Peek()
		if rune == ' ' || rune == '\t' {
			lexer.Next()
			lexer.Ignore()
		} else {
			break
		}
	}

	if rune := lexer.Peek(); rune == '{' {
		lexer.Next()
		lexer.Emit(ItemLeftCurly)
		return lexMultilineShellString
	} else {
		// Handle shell line
		return lexShellString
	}
}

func lexShellString(lexer *Lexer) stateModifier {
	rune := lexer.Peek()
	switch rune {
	case '\n':
		lexer.Emit(ItemStringPart)
		return lexRoot
	case '\\':
		lexer.Next()
		lexer.Next()
	case '$':
		lexer.Next()
		if rune := lexer.Peek(); rune == '(' {
			lexer.Backtrack()
			lexer.Emit(ItemStringPart)
			lexer.Next()
			lexer.Next()
			lexer.Emit(ItemSubstitutionStart)
			lexer.Mode = ModeShellString
			lexer.substitutionDepth++
			return lexRoot
		} else {
			lexer.Next()
		}
	default:
		lexer.Next()
	}

	return lexShellString
}

// assumes the left curly brace has been consumed as the start of the evaluated string
func lexMultilineShellString(lexer *Lexer) stateModifier {
	rune := lexer.Peek()
	switch rune {
	case '\\':
		lexer.Next()
		lexer.Next()
	case '$':
		lexer.Next()
		if rune := lexer.Peek(); rune == '(' {
			lexer.Backtrack()
			lexer.Emit(ItemStringPart)
			lexer.Next()
			lexer.Next()
			lexer.Emit(ItemSubstitutionStart)
			lexer.Mode = ModeMultilineShellString
			lexer.substitutionDepth++
			return lexRoot
		} else {
			lexer.Next()
		}
	case '}':
		lexer.Emit(ItemStringPart)
		return lexRoot
	default:
		lexer.Next()
	}

	return lexMultilineShellString
}

// assumes one quote has been consumed as the start of the evaluated string
func lexEvaluatedString(lexer *Lexer) stateModifier {
	rune := lexer.Peek()
	switch rune {
	case '"':
		lexer.Emit(ItemStringPart)
		lexer.Next()
		lexer.Emit(ItemDoubleQuote)
		return lexRoot
	case '\\':
		lexer.Next()
		escaped := lexer.Next()
		switch escaped {
		case '"', '\\', '$':
			// Do nothing
		default:
			lexer.errorf("unexpected escape sequence '%c'", rune)
			return nil
		}
	case '$':
		lexer.Next()
		if rune := lexer.Peek(); rune == '(' {
			lexer.Backtrack()
			lexer.Emit(ItemStringPart)
			lexer.Next()
			lexer.Next()
			lexer.Emit(ItemSubstitutionStart)
			lexer.Mode = ModeEvaluatedString
			lexer.substitutionDepth++
			return lexRoot
		} else {
			lexer.Next()
		}
	default:
		lexer.Next()
	}

	return lexEvaluatedString
}
