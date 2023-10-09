package token

import (
	"bufio"
	"io"
	"os"
	"unicode"
)

func Generate(f *os.File) (*Token, error) {
	var curToken *Token
	head := InitToken()
	curToken = head
	spaceCnt := 0

	br := bufio.NewReader(f)
	for {
		c, _, err := br.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if unicode.IsSpace(c) {
			spaceCnt++
			if IsSeparete(c) {
				curToken = NewToken(curToken, SEPARATE, string(c), 0)
				spaceCnt = 0
			}
			continue
		}

		if IsReserve(c) {
			curToken = NewToken(curToken, RESERVED, string(c), spaceCnt)
			spaceCnt = 0
			continue
		}
		curToken = NewToken(curToken, PLAIN_TEXT, string(c), 0)
		spaceCnt = 0
	}
	return head.Next, nil
}
