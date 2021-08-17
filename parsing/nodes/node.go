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
	NodeTypeBinaryExpression
	NodeTypeIdentifier
	NodeTypeInteger
	NodeTypePrimaryExpression
	NodeTypeUnaryExpression
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
)
