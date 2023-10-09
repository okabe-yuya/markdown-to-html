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

func TestH1Only(t *testing.T) {
	res, expect, err := ResultAndExpect("h1_only")
	if err != nil {
		t.Errorf("[ERROR] Caused error: %e", err)
		t.Fail()
	}

	if res != expect {
		t.Errorf("[ERROR] Failed Test (Expect=%s, Result=%s): \nerror: %e", res, expect, err)
		t.Fail()
	}
}

func TestH1ToH6(t *testing.T) {
	res, expect, err := ResultAndExpect("h1_h6")
	if err != nil {
		t.Errorf("[ERROR] Caused error: %e", err)
		t.Fail()
	}

	if res != expect {
		t.Errorf("[ERROR] Failed Test (Expect=%s, Result=%s): \nerror: %e", res, expect, err)
		t.Fail()
	}
}

func TestEmpty(t *testing.T) {
	res, expect, err := ResultAndExpect("empty")
	if err != nil {
		t.Errorf("[ERROR] Caused error: %e", err)
		t.Fail()
	}

	if res != expect {
		t.Errorf("[ERROR] Failed Test (Expect=%s, Result=%s): \nerror: %e", res, expect, err)
		t.Fail()
	}
}
