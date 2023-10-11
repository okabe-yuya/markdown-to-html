package generator

import (
	"fmt"

	"github.com/okabe-yuya/makrdown-to-html/parser"
)

func Html(blocks []*parser.Node) string {
	html := "<html>\n"
	html += header()
	html += "<body>\n"
	for _, b := range blocks {
		switch b.Kind {
		case parser.ND_HEADER:
			html += fmt.Sprintf("<h%d>%s</h%d>\n", b.Level, b.Value, b.Level)
		case parser.ND_LIST:
			h, _ := listToHtml(b)
			html += h
		case parser.ND_VALUE:
			html += fmt.Sprintf("<p>%s</p>\n", b.Value)
		default:
			continue
		}
	}
	html += "</body>\n"
	html += "</html>"
	return html
}

func header() string {
	head := "<head>\n"
	head += "<title>Generate html from markdown!</title>\n"
	head += "</head>\n"
	return head
}

func listToHtml(b *parser.Node) (string, *parser.Node) {
	curBlock := b
	res := "<ul>\n"

	for {
		if curBlock == nil {
			break
		}

		if b.Depth == curBlock.Depth {
			res += fmt.Sprintf("<li>%s</li>\n", curBlock.Value)
			curBlock = curBlock.Nest
		} else {
			if b.Depth > curBlock.Depth {
				break
			}
			r, bk := listToHtml(curBlock)
			curBlock = bk
			res += r
		}
	}

	res += "</ul>\n"
	return res, curBlock
}
