package lexing

import "unicode"

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
		if rune := lexer.Peek(); rune == '/' {
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
			return lexRoot
		} else {
			lexer.Emit(ItemDivision)
			return lexRoot
		}
	case '+':
		lexer.Next()
		lexer.Emit(ItemAddition)
		return lexRoot
	case '-':
		lexer.Next()
		lexer.Emit(ItemSubtraction)
		return lexRoot
	case '*':
		lexer.Next()
		lexer.Emit(ItemMultiplication)
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
			lexer.Emit(ItemAnd)
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
		return lexRoot
	case ')':
		lexer.Next()
		lexer.Emit(ItemRightParentheses)
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
		for {
			rune := lexer.Next()
			if rune == '`' {
				break
			} else if rune == eof {
				lexer.errorf("unexpected end of file")
				return nil
			}
		}
		lexer.Emit(ItemRawString)
		return lexRoot
	case '"':
		lexer.Next()
		for {
			rune := lexer.Next()
			if rune == '\\' {
				escaped := lexer.Next()
				switch escaped {
				case '"', '\\':
					// Do nothing
				default:
					lexer.errorf("unexpected escape sequence '%c'", rune)
					return nil
				}
			} else if rune == '"' {
				break
			} else if rune == '\n' {
				lexer.errorf("unexpected newline")
				return nil
			} else if rune == eof {
				lexer.errorf("unexpected end of file")
				return nil
			}
		}
		lexer.Emit(ItemInterpretedString)
		return lexRoot
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
				case "else":
					lexer.Emit(ItemKeywordElse)
				case "return":
					lexer.Emit(ItemKeywordReturn)
				case "let":
					lexer.Emit(ItemKeywordLet)
				case "shell":
					lexer.Emit(ItemKeywordShell)
					return lexShell
				case "true", "false":
					lexer.Emit(ItemBoolean)
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
		// Handle shell block
		lexer.Next()
		lexer.Emit(ItemLeftCurly)

		curlyDepth := 1
		for {
			rune := lexer.Peek()
			if rune == '{' {
				lexer.Next()
				curlyDepth++
			} else if rune == '}' {
				curlyDepth--
				if curlyDepth == 0 {
					break
				} else {
					lexer.Next()
				}
			} else if rune == eof {
				lexer.errorf("unexpected end of file")
				return nil
			} else {
				lexer.Next()
			}
		}
		lexer.Emit(ItemShellString)

		// Terminating right curly bracket
		rune = lexer.Next()
		if rune == '}' {
			lexer.Emit(ItemRightCurly)
		} else {
			lexer.errorf("unexpected token '%c'", rune)
			return nil
		}

		return lexRoot
	} else {
		// Handle shell line
		for {
			rune := lexer.Peek()
			if rune == '\n' || rune == eof {
				break
			} else {
				lexer.Next()
			}
		}
		lexer.Emit(ItemShellString)
	}

	return lexRoot
}
