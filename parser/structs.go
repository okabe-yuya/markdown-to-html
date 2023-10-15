package parser

type Node struct {
	Kind  NodeKind
	Nest  *Node
	Sub   *Node // for list
	Level int
	Depth int
	Value string
}

type NodeKind int

const (
	ND_END       NodeKind = 0
	ND_NEW_LINE  NodeKind = 1
	ND_VALUE     NodeKind = 2
	ND_HEADER    NodeKind = 3
	ND_LIST      NodeKind = 4
	ND_WEIGHT    NodeKind = 5
	ND_ITALIC    NodeKind = 6
	ND_QUOTE     NodeKind = 7
	ND_BACKQUOTE NodeKind = 8
	ND_LINK      NodeKind = 9
)

func NewNode(kind NodeKind, value string, level, depth int) *Node {
	return &Node{
		Kind:  kind,
		Level: level,
		Depth: depth,
		Value: value,
	}
}
