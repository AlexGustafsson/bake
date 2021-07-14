package lexing

//go:generate stringer -type=ItemType
type ItemType int

const (
	ItemNewline ItemType = iota
	ItemEndOfFile
	ItemError
	ItemImport
	ItemLeftParentheses
	ItemRightParentheses
	ItemDoubleQuote
	ItemString
	ItemWhitespace
)
