package parser

type Node struct {
	Kind  NodeKind
	Nest  *Node
	Level int
	Depth int
	Value string
}

type NodeKind int

const (
	ND_END    NodeKind = 0
	ND_VALUE  NodeKind = 1
	ND_HEADER NodeKind = 2
	ND_LIST   NodeKind = 3
)

func NewNode(kind NodeKind, value string, level, depth int, nest *Node) *Node {
	return &Node{
		Kind:  kind,
		Nest:  nest,
		Level: level,
		Depth: depth,
		Value: value,
	}
}
