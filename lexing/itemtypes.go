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
	ItemEqualEqual
	ItemNotEqual
	ItemSpread

	// Punctuation
	ItemLeftParentheses
	ItemRightParentheses
	ItemLeftBracket
	ItemRightBracket
	ItemLeftCurly
	ItemRightCurly
	ItemColon
	ItemComma
	ItemDot
	ItemDollar

	// Keywords
	ItemKeywordPackage
	ItemKeywordImport
	ItemKeywordFunc
	ItemKeywordRule
	ItemKeywordExport
	ItemKeywordReturn

	// Identifiers
	ItemIdentifier

	// Whitespace
	ItemNewline
	ItemWhitespace

	// Characters
	ItemLetter
	ItemUnicodeCharacter
	ItemDecimalDigit

	// Strings etc.
	ItemRawString

	// Comments
	ItemComment
)
