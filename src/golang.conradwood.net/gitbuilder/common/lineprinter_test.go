package common

import (
	"testing"
)

func TestLines(t *testing.T) {
	checkLineBreak(t, "line\nplusnoline", []string{"line\n", "plusnoline"})
	checkLineBreak(t, "line\n", []string{"line\n"})
	checkLineBreak(t, "noline", []string{"noline"})
	checkLineBreak(t, "\n", []string{"\n"})
	checkLineBreak(t, "", []string{})
	checkLineBreak(t, "a", []string{"a"})
	checkLineBreak(t, "line\nline\nline\n", []string{"line\n", "line\n", "line\n"})
	lp := &LinePrinter{Prefix: "[test]", MaxLineLength: 30}
	lp.Printf("first test")
	lp.Printf("second test\n")
	lp.Printf("third test\n")
	lp.Printf("a really long line that should be truncated in the output, we will see if it does. The quick brown fox jumps over the lazy dog.")
	lp.Printf("should not be printed - continuing a really long line that should be truncated in the output\n")
	lp.Printf("should be printed, is asecond really long line that should be truncated in the output, we will see if it does\n")
}

func checkLineBreak(t *testing.T, txt string, result []string) {
	lp := &LinePrinter{Prefix: "test", MaxLineLength: 50}
	lines := lp.lines(txt)
	if len(lines) != len(result) {
		t.Errorf("text \"%s\": len mismatch (expected %d vs actual %d)", txt, len(result), len(lines))
	}
	for i, _ := range lines {
		if lines[i] != result[i] {
			t.Errorf(" text \"%s\": did not split properly. line %d expected: \"%s\" vs actual \"%s\"\n", txt, i, result[i], lines[i])
		}
	}
}
