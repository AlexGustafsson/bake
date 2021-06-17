package lexing

func lexRoot(lexer *Lexer) stateModifier {
	return lexSpace
}

func lexSpace(lexer *Lexer) stateModifier {
	rune := lexer.Next()
	if rune == ' ' {
		lexer.Emit(ItemSpace)
		return lexSpace
	} else if rune == eof {
		lexer.Emit(ItemEndOfFile)
		return nil
	} else {
		lexer.Errorf("Unexpected token '%c', expected ' '", rune)
		return nil
	}
}
