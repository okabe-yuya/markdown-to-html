package block

type Block struct {
	Kind  BlockKind
	Nest  *Block
	Level int
	Value string
}

type BlockKind int

const (
	ND_END    BlockKind = 0
	ND_VALUE  BlockKind = 1
	ND_HEADER BlockKind = 2
)

func NewBlock(kind BlockKind, value string, level int, nest *Block) *Block {
	return &Block{
		Kind:  kind,
		Nest:  nest,
		Level: level,
		Value: value,
	}
}
