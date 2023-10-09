package block

import "github.com/okabe-yuya/makrdown-to-html/token"

func Generate(t *token.Token) ([]*Block, error) {
	var block *Block
	var blocks []*Block
	curToken := t

	// Avoid empty markdown file
	if t == nil {
		return blocks, nil
	}

	for curToken.Next != nil {
		switch curToken.Kind {
		case token.RESERVED:
			block, curToken = reserved(curToken)
			blocks = append(blocks, block)
		case token.PLAIN_TEXT:
			block, curToken = plaintext(curToken)
			blocks = append(blocks, block)
		case token.SEPARATE:
			curToken = curToken.Next
			continue
		}
	}
	return blocks, nil
}

func reserved(token *token.Token) (*Block, *token.Token) {
	switch token.Value {
	case "#":
		return header(token)
	default:
		panic(1)
	}
}

func header(t *token.Token) (*Block, *token.Token) {
	hdrLvl := 0
	value := ""
	curToken := t

	for curToken.Kind == token.RESERVED && curToken.Value == "#" {
		hdrLvl++
		curToken = curToken.Next
	}

	if curToken.Kind == token.PLAIN_TEXT {
		for {
			if curToken.Kind == token.SEPARATE {
				curToken = curToken.Next
				break
			}
			value += curToken.Value
			curToken = curToken.Next
		}
	}

	res := NewBlock(ND_HEADER, value, hdrLvl, nil)
	return res, curToken
}

func plaintext(t *token.Token) (*Block, *token.Token) {
	res := NewBlock(ND_VALUE, t.Value, 1, nil)
	nextToken := t.Next

	for nextToken.Kind == token.PLAIN_TEXT {
		res.Value += nextToken.Value
		nextToken = nextToken.Next
	}
	return res, nextToken
}
