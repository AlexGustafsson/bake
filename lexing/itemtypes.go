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
	ItemSubtraction
	ItemMultiplication
	ItemDivision
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

	// Keywords
	ItemKeywordPackage
	ItemKeywordImport
	ItemKeywordFunc
	ItemKeywordRule
	ItemKeywordExport
	ItemKeywordIf
	ItemKeywordElse
	ItemKeywordReturn
	ItemKeywordLet
	ItemKeywordShell

	// Identifiers
	ItemIdentifier

	// Whitespace
	ItemNewline
	ItemWhitespace

	// Strings etc.
	ItemRawString
	ItemInterpretedString
	ItemShellString

	// Numbers
	ItemInteger

	// Boleans
	ItemBoolean

	// Comments
	ItemComment
)
