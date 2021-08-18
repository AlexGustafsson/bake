package nodes

type Node interface {
	Type() NodeType
	String() string
	DotString() string
	Position() NodePosition
	// Start is the offset to the start of the token in bytes
	Start() int
	// Line is the zero-based line of the input on which the token is found
	Line() int
	// Column is the zero-based column (in runes) of the line in which the token is found
	Column() int
}

type NodeType int

func (nodeType NodeType) Type() NodeType {
	return nodeType
}

type NodePosition struct {
	start  int
	line   int
	column int
}

func (position NodePosition) Position() NodePosition {
	return position
}

func (position NodePosition) Start() int {
	return position.start
}

func (position NodePosition) Line() int {
	return position.line
}

func (position NodePosition) Column() int {
	return position.column
}

func CreateNodePosition(start int, line int, column int) NodePosition {
	return NodePosition{
		start:  start,
		line:   line,
		column: column,
	}
}

const (
	NodeTypeSourceFile NodeType = iota
	NodeTypePackageDeclaration
	NodeTypeImportsDeclaration
	NodeTypeComment
	NodeTypeInterpretedString
	NodeTypeRawString
	NodeTypeVariableDeclaration
	NodeTypeIdentifier
	NodeTypeInteger
	NodeTypeSignature
	NodeTypeFunctionDeclaration
	NodeTypeBlock
	NodeTypeSelector
	NodeTypeImportSelector
	NodeTypeIndex
	NodeTypeInvokation
	NodeTypeIncrement
	NodeTypeDecrement
	NodeTypeLooseAssignment
	NodeTypeAdditionAssignment
	NodeTypeSubtractionAssignment
	NodeTypeMultiplicationAssignment
	NodeTypeDivisionAssignment
	NodeTypeShellStatement
	NodeTypeAssignment
	NodeTypeComparison
	NodeTypeEquality
	NodeTypeFactor
	NodeTypePrimary
	NodeTypeTerm
	NodeTypeUnary
)
