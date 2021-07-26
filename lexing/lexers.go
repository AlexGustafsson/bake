package lexing

func lexRoot(lexer *Lexer) stateModifier {
	switch rune := lexer.Peek(); rune {
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
	case ' ', '\t':
		lexer.Next()
		lexer.Emit(ItemWhitespace)
		return lexRoot
	case '\n':
		lexer.Next()
		lexer.Emit(ItemNewline)
		return lexRoot
	case eof:
		lexer.Next()
		lexer.Emit(ItemEndOfInput)
		return nil
	default:
		lexer.errorf("unexpected token '%c'", rune)
		return nil
	}
}
