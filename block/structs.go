package block

type Block struct {
	Kind  BlockKind
	Nest  *Block
	Level int
	Depth int
	Value string
}

type BlockKind int

const (
	ND_END    BlockKind = 0
	ND_VALUE  BlockKind = 1
	ND_HEADER BlockKind = 2
	ND_LIST   BlockKind = 3
)

func NewBlock(kind BlockKind, value string, level, depth int, nest *Block) *Block {
	return &Block{
		Kind:  kind,
		Nest:  nest,
		Level: level,
		Depth: depth,
		Value: value,
	}
}
