package playground

import (
	"os"
	"strings"
	"testing"
)

func TestIndexHTML(t *testing.T) {
	page, err := os.ReadFile("index.html")
	if err != nil {
		t.Fatalf("read index.html: %v", err)
	}
	got := string(page)
	if !strings.HasPrefix(got, "<!DOCTYPE html>") {
		t.Fatalf("index.html has no doctype")
	}
	if !strings.Contains(got, "<title>playground</title>") {
		t.Fatalf("index.html has no title")
	}
}
