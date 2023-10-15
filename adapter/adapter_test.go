package adapter

import (
	"fmt"
	"os"
	"testing"
)

func ResultAndExpect(filename string) (string, string, error) {
	md, err := os.Open(fmt.Sprintf("../static/md/%s.md", filename))
	if err != nil {
		return "", "", err
	}
	html, err := GenerateHtml(md)
	if err != nil {
		return "", "", err
	}

	bytes, err := os.ReadFile(fmt.Sprintf("../static/html/%s.html", filename))
	if err != nil {
		return "", "", err
	}
	return html, string(bytes), nil
}

func CommonExecuter(path string, t *testing.T) {
	res, expect, err := ResultAndExpect(path)
	if err != nil {
		t.Errorf("[ERROR] Caused error: %v", err)
		t.Fail()
	}

	if res != expect {
		t.Errorf("[ERROR] Failed Test (Result=\n%s\n, Expect=\n%s\n): \nerror: %v", res, expect, err)
		t.Fail()
	}
}

// common
func TestEmpty(t *testing.T) {
	CommonExecuter("empty", t)
}

func TestVariety(t *testing.T) {
	CommonExecuter("variety", t)
}

// head
func TestHeadH1_Only(t *testing.T) {
	CommonExecuter("head/h1_only", t)
}

func TestHeadH1ToH6(t *testing.T) {
	CommonExecuter("head/h1_h6", t)
}

// list
func TestListAnyNest(t *testing.T) {
	CommonExecuter("list/any_nest", t)
}

func TestListNest(t *testing.T) {
	CommonExecuter("list/nest", t)
}

func TestListNoNest(t *testing.T) {
	CommonExecuter("list/no_nest", t)
}

func TestListTwoNest(t *testing.T) {
	CommonExecuter("list/two_nest", t)
}

func TestListWeight(t *testing.T) {
	CommonExecuter("list/weight", t)
}

// weight
func TestWeightSimple(t *testing.T) {
	CommonExecuter("weight/simple", t)
}

func TestWeightInText(t *testing.T) {
	CommonExecuter("weight/in_text", t)
}

// italic
func TestItalicSimple(t *testing.T) {
	CommonExecuter("italic/simple", t)
}

func TestItalicInText(t *testing.T) {
	CommonExecuter("italic/in_text", t)
}

func TestItalicAnyItalic(t *testing.T) {
	CommonExecuter("italic/any_italic", t)
}

// quote
func TestQuoteSingle(t *testing.T) {
	CommonExecuter("quote/single", t)
}

func TestQuoteNest(t *testing.T) {
	CommonExecuter("quote/nest", t)
}

func TestQuoteTwoNest(t *testing.T) {
	CommonExecuter("quote/two_nest", t)
}

func TestQuoteDecorate(t *testing.T) {
	CommonExecuter("quote/decorate", t)
}

func TestQuoteVariety(t *testing.T) {
	CommonExecuter("quote/variety", t)
}

// backquote
func TestBackQuoteSingle(t *testing.T) {
	CommonExecuter("backquote/single", t)
}

func TestBackQuoteTriple(t *testing.T) {
	CommonExecuter("backquote/triple", t)
}

func TestBackQuoteAny(t *testing.T) {
	CommonExecuter("backquote/any", t)
}

// link
func TestLinkSingle(t *testing.T) {
	CommonExecuter("link/single", t)
}

func TestLinkTwo(t *testing.T) {
	CommonExecuter("link/two", t)
}

func TestLinkTextInLink(t *testing.T) {
	CommonExecuter("link/text_in_link", t)
}

func TestLinkLinkAndText(t *testing.T) {
	CommonExecuter("link/link_and_text", t)
}
