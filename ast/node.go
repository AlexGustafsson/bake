package ast

type Node interface {
	Type() NodeType
	String() string
	Start() Position
	End() Position
}

//go:generate stringer -type=NodeType
type NodeType int

func (nodeType NodeType) Type() NodeType {
	return nodeType
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
	NodeTypeRuleFunctionDeclaration
	NodeTypeBlock
	NodeTypeSelector
	NodeTypeImportSelector
	NodeTypeIndex
	NodeTypeInvocation
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
	NodeTypeAliasDeclaration
	NodeTypeRuleDeclaration
	NodeTypeReturnStatement
	NodeTypeIfStatement
)