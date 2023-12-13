package common

import (
	"fmt"
	"io"
	"strings"
)

type LinePrinter struct {
	Prefix        string
	MaxLineLength int
	linelen       int       // current linelen
	target        io.Writer // if nil, uses stdout
	truncated     bool
}

func (l *LinePrinter) Printf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	l.Print(s)
}
func (l *LinePrinter) Print(s string) {
	// our own split function, because we must retain the "\n"
	sx := l.lines(s)
	for _, line := range sx {
		l.printLine(line)
	}
}

// like strings.Split but does not remove the seperator
func (l *LinePrinter) lines(txt string) []string {
	if txt == "" {
		return []string{}
	}
	if !strings.Contains(txt, "\n") {
		return []string{txt}
	}
	var res []string
	s := txt
	for {
		i := strings.Index(s, "\n")
		if i == -1 {
			break
		}
		res = append(res, s[:i+1])
		s = s[i+1:]
	}
	if len(s) > 0 {
		res = append(res, s)
	}
	return res
}

// print text, but cut lines to max line
// line must be EITHER a line terminated by \n OR a partial line not containing any \n
func (l *LinePrinter) printLine(line string) {
	if line == "" {
		return
	}
	if l.MaxLineLength == 0 {
		fmt.Print(line)
		return
	}

	if l.truncated {
		// do not print more partial lines if previous one was cut short
		// but if this line is a complete line with terminator, reset flag
		if line[len(line)-1] == '\n' {
			l.truncated = false
			l.linelen = 0
		}
		return
	}
	s := line
	if l.linelen+len(line) > l.MaxLineLength {
		l.truncated = true
		l.linelen = 0
		s = line[:l.MaxLineLength-l.linelen] + "...\n"
	}

	fmt.Print(l.Prefix + s)

	if line[len(line)-1] == '\n' {
		l.truncated = false
		l.linelen = 0
	}
}



