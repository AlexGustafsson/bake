package nodes

type PrimaryExpression struct {
	NodeType
	NodePosition
	Expression Node // TODO - document types
}

type PrimaryOperator int

const (
	PrimaryOperatorSubtraction PrimaryOperator = iota
	PrimaryOperatorNot
	PrimaryOperatorSpread
)

func CreatePrimaryExpression(position NodePosition, operator PrimaryOperator, expression Node) *PrimaryExpression {
	return &PrimaryExpression{
		NodeType:     NodeTypePrimaryExpression,
		NodePosition: position,
		Expression:   expression,
	}
}

func (node *PrimaryExpression) String() string {
	return ""
}
