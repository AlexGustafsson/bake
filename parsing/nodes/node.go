package nodes

type Node interface {
	Type() NodeType
	String() string
	Position() NodePosition
}

type NodeType int

func (nodeType NodeType) Type() NodeType {
	return nodeType
}

type NodePosition int

func (nodePosition NodePosition) Position() NodePosition {
	return nodePosition
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
