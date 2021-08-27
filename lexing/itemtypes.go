package lexing

//go:generate stringer -type=ItemType
type ItemType int

const (
	// Meta
	ItemStartOfInput ItemType = iota
	ItemEndOfInput
	ItemError

	// Operators
	ItemAddition
	ItemAdditionAssign
	ItemSubtraction
	ItemSubtractionAssign
	ItemMultiplication
	ItemMultiplicationAssign
	ItemDivision
	ItemDivisionAssign
	ItemAssignment
	ItemLooseAssignment
	ItemEquals
	ItemNot
	ItemNotEqual
	ItemLessThan
	ItemLessThanOrEqual
	ItemGreaterThan
	ItemGreaterThanOrEqual
	ItemAnd
	ItemOr
	ItemSpread
	ItemIncrement
	ItemDecrement
	ItemModulo
	ItemModuloAssign

	// Punctuation
	ItemLeftParentheses
	ItemRightParentheses
	ItemLeftBracket
	ItemRightBracket
	ItemLeftCurly
	ItemRightCurly
	ItemColon
	ItemColonColon
	ItemComma
	ItemDot
	ItemSubstitutionStart
	ItemSubstitutionEnd
	ItemDoubleQuote

	// Keywords
	ItemKeywordPackage
	ItemKeywordImport
	ItemKeywordFunc
	ItemKeywordRule
	ItemKeywordExport
	ItemKeywordIf
	ItemKeywordFor
	ItemKeywordIn
	ItemKeywordElse
	ItemKeywordReturn
	ItemKeywordLet
	ItemKeywordShell
	ItemKeywordAlias

	// Identifiers
	ItemIdentifier

	// Whitespace
	ItemNewline
	ItemWhitespace

	// Strings etc.
	ItemRawString
	ItemStringPart

	// Numbers
	ItemInteger

	// Boleans
	ItemBoolean

	// Comments
	ItemComment
)
