package lexer

type Token struct {
	Kind  TokenKind
	Value string
	Next  *Token
	Depth int
}

type TokenKind int

const (
	END        TokenKind = 0
	START      TokenKind = 1
	BLANK      TokenKind = 2
	SEPARATE   TokenKind = 3
	RESERVED   TokenKind = 4
	PLAIN_TEXT TokenKind = 5
)

var RESERVED_RUNES = []rune{
	'#', '*', '-', '_', '`', '[', ']', '(', ')', '!', '>',
}

func InitToken() *Token {
	return &Token{
		Kind:  START,
		Value: "",
		Next:  nil,
		Depth: 0,
	}
}

func NewToken(cur *Token, kind TokenKind, value string, depth int) *Token {
	token := &Token{
		Kind:  kind,
		Value: value,
		Next:  nil,
		Depth: depth,
	}
	cur.Next = token
	return token
}

func IsReserve(r rune) bool {
	for i := 0; i < len(RESERVED_RUNES); i++ {
		if r == RESERVED_RUNES[i] {
			return true
		}
	}
	return false
}

func IsSeparete(r rune) bool {
	return r == '\n'
}
