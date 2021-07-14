package lexing

func lexRoot(lexer *Lexer) stateModifier {
	switch rune := lexer.Peek(); rune {
	case 'i':
		return lexImport
	case '(':
		lexer.Next()
		lexer.Emit(ItemLeftParentheses)
		return lexRoot
	case ')':
		lexer.Next()
		lexer.Emit(ItemRightParentheses)
		return lexRoot
	case '"':
		return lexString
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
		lexer.Emit(ItemEndOfFile)
		return nil
	default:
		lexer.Errorf("Unexpected token '%c'", rune)
		return nil
	}
}

func lexImport(lexer *Lexer) stateModifier {
	if ok := lexer.NextString("import"); ok {
		lexer.Emit(ItemImport)
		return lexRoot
	} else {
		lexer.Errorf("Expected token 'import'")
		return nil
	}
}

func lexString(lexer *Lexer) stateModifier {
	rune := lexer.Next()
	if rune == '"' {
		for {
			rune = lexer.Next()
			if rune == '"' {
				lexer.Emit(ItemString)
				return lexRoot
			} else if rune == eof {
				lexer.Errorf("Unexpected end of file in string")
				return nil
			} else if rune == '\n' {
				lexer.Errorf("Unexpected end of line in string")
				return nil
			}
		}
	} else {
		lexer.Errorf("Unexpected token '%c', expected '\"'", rune)
		return nil
	}
}
