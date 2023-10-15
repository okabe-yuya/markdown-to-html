package generator

import (
	"fmt"

	"github.com/okabe-yuya/makrdown-to-html/parser"
)

func Html(nodes []*parser.Node) string {
	html := "<html>\n"
	html += header()
	html += "<body>\n"
	for _, node := range nodes {
		switch node.Kind {
		case parser.ND_HEADER:
			html += fmt.Sprintf("<h%d>%s</h%d>\n", node.Level, node.Value, node.Level)
		case parser.ND_LIST:
			h, _ := listToHtml(node)
			html += h
		case parser.ND_VALUE:
			h := valueToHtml(node)
			html += h
		case parser.ND_NEW_LINE:
			html += "<br />\n"
		case parser.ND_QUOTE:
			h, _ := quoteToHtml(node)
			html += h
		case parser.ND_BACKQUOTE:
			if node.Level == 1 {
				html += fmt.Sprintf("<code>%s</code>\n", node.Value)
			} else if node.Level == 3 {
				html += fmt.Sprintf("<pre><code>%s</code></pre>\n", node.Value)
			}
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

func quoteToHtml(node *parser.Node) (string, *parser.Node) {
	curNode := node
	res := "<blockquote>\n"

	for {
		if curNode == nil {
			break
		}

		if node.Depth == curNode.Depth {
			res += valueToHtml(curNode.Nest)
			curNode = curNode.Sub
		} else {
			if node.Depth > curNode.Depth {
				break
			}
			r, bk := quoteToHtml(curNode)
			curNode = bk
			res += r
		}
	}
	res += "</blockquote>\n"
	return res, curNode
}

func listToHtml(node *parser.Node) (string, *parser.Node) {
	curNode := node
	res := "<ul>\n"

	for {
		if curNode == nil {
			break
		}

		if node.Depth == curNode.Depth {
			res += fmt.Sprintf("<li>%s</li>\n", valueToHtmlForList(curNode.Nest))
			curNode = curNode.Sub
		} else {
			if node.Depth > curNode.Depth {
				break
			}
			r, bk := listToHtml(curNode)
			curNode = bk
			res += r
		}
	}

	res += "</ul>\n"
	return res, curNode
}

func valueToHtmlForList(node *parser.Node) string {
	if node.Nest == nil {
		return node.Value
	} else {
		html := node.Value
		html += _valueToHtml(node.Nest)
		return html
	}
}

func valueToHtml(node *parser.Node) string {
	if node.Nest == nil {
		return fmt.Sprintf("<p>%s</p>\n", node.Value)
	} else {
		html := "<p>"
		html += node.Value
		html += _valueToHtml(node.Nest)
		html += "</p>\n"
		return html
	}
}

func _valueToHtml(node *parser.Node) string {
	html := ""
	switch node.Kind {
	case parser.ND_VALUE:
		html += node.Value
	case parser.ND_WEIGHT:
		html += fmt.Sprintf("<b>%s</b>", node.Value)
	case parser.ND_ITALIC:
		html += fmt.Sprintf("<i>%s</i>", node.Value)
	}

	if node.Nest != nil {
		html += _valueToHtml(node.Nest)
	}
	return html
}
