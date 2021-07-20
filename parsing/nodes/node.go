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
	NodeImportType NodeType = iota
)