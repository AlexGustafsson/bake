package lexing

import "unicode"

func lexRoot(lexer *Lexer) stateModifier {
	switch rune := lexer.Peek(); rune {
	case ' ', '\t':
		lexer.Next()
		// lexer.Emit(ItemWhitespace)
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
				rune := lexer.Next()
				if rune == '\n' {
					break
				}
			}
			lexer.Emit(ItemComment)
			return lexRoot
		} else {
			lexer.errorf("unexpected token '%c'", rune)
			return nil
		}
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
