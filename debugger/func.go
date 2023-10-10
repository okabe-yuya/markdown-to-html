package debugger

import (
	"fmt"

	"github.com/okabe-yuya/makrdown-to-html/block"
	"github.com/okabe-yuya/makrdown-to-html/token"
)

func Token(token *token.Token) {
	c := token
	for c.Next != nil {
		fmt.Printf(format(), c.Kind, c.Value, c.Depth)
		c = c.Next
	}
}

func Block(blocks []*block.Block) {
	for _, b := range blocks {
		fmt.Printf(format(), b.Kind, b.Value)
	}
}

func format() string {
	return "kind: %d, value: %s, size: %d \n"
}
