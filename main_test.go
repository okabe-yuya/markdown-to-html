package main

import (
	"fmt"
	"os"
	"testing"
)

func ResultAndExpect(filename string) (string, string, error) {
	md, err := os.Open(fmt.Sprintf("static/md/%s.md", filename))
	if err != nil {
		return "", "", err
	}
	html, err := GenrateHtml(md)
	if err != nil {
		return "", "", err
	}

	bytes, err := os.ReadFile(fmt.Sprintf("static/html/%s.html", filename))
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
