package github

import (
	"strings"
	"testing"
)

var md = []string{
	`prefix text
	| a | b |
	|---|---|
	| | |`,
	`| Commit | Description |
	| -- | -- |
	| 7a4d9805ea | fix(compiler-cli): use '' for the source map URL of indirect templates (#41973) |`,
}

func TestMarkdownTable(t *testing.T) {
	for _, m := range md {
		mkd := parseMarkdown([]byte(m))

		if !strings.Contains(mkd, "table") {
			t.Fatalf("Expected a table, got %s", mkd)
		}
	}
}
