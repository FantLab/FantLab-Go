package scheme

import (
	"strings"
	"unicode/utf8"
)

const indent = "  "

type line struct {
	depth   int
	builder strings.Builder
	text    string
	comment string
}

type lines struct {
	list    []*line
	current *line
}

func (ls *lines) new(depth int) {
	if ls.current != nil {
		ls.current.text = ls.current.builder.String()
		ls.current.builder.Reset()
		ls.list = append(ls.list, ls.current)
	}
	ls.current = &line{depth: depth}
}

func (ls *lines) join(prefix, postfix string) string {
	indentLen := utf8.RuneCountInString(indent)

	var maxLineLen int
	{
		for i, line := range ls.list {
			lineLen := line.depth*indentLen + utf8.RuneCountInString(line.text)

			if i == 0 || lineLen > maxLineLen {
				maxLineLen = lineLen
			}
		}
	}

	var sb strings.Builder

	sb.WriteString(prefix)

	for i, line := range ls.list {
		if i > 0 {
			sb.WriteRune('\n')
		}

		if indentLen > 0 {
			for j := 0; j < line.depth; j++ {
				sb.WriteString(indent)
			}
		}

		sb.WriteString(line.text)

		if "" != line.comment {
			lineLen := line.depth*indentLen + utf8.RuneCountInString(line.text)

			for j := 0; j < maxLineLen-lineLen+1; j++ {
				sb.WriteRune(' ')
			}

			sb.WriteString(line.comment)
		}
	}

	sb.WriteString(postfix)

	return sb.String()
}
