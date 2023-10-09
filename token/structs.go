package token

type Token struct {
	Kind   TokenKind
	Value  string
	Next   *Token
	Lenght int
}

type TokenKind int

const (
	END        TokenKind = 0
	START      TokenKind = 1
	RESERVED   TokenKind = 2
	PLAIN_TEXT TokenKind = 3
	SEPARATE   TokenKind = 99
)

var RESERVED_RUNES = []rune{
	'#', '*', '-', '_', '`', '[', '(', '!', '>',
}

func InitToken() *Token {
	return &Token{
		Kind:   START,
		Value:  "",
		Next:   nil,
		Lenght: 0,
	}
}

func NewToken(cur *Token, kind TokenKind, value string, len int) *Token {
	token := &Token{
		Kind:   kind,
		Value:  value,
		Next:   nil,
		Lenght: len,
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
