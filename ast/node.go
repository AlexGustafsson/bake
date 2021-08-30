package ast

type baseNode struct {
	nodeType  NodeType
	nodeRange *Range
}

type Node interface {
	Type() NodeType
	Range() *Range
}

func (node *baseNode) Type() NodeType {
	return node.nodeType
}

func (node *baseNode) Range() *Range {
	return node.nodeRange
}

//go:generate stringer -type=NodeType
type NodeType int

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
	NodeTypeEvaluatedString
	NodeTypeStringPart
	NodeTypeForStatement
	NodeTypeModuloAssignment
	NodeTypeObject
)
