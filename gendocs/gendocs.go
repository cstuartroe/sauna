package gendocs

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"

	"github.com/cstuartroe/sauna/gloss"
)

func gl(s string) string {
	words, err := gloss.ParseGloss(s)
	if err != nil {
		panic(err)
	}
	return gloss.Text(words)
}

func glossedSentence(s string, translation string) string {
	wrong := false
	if s[0] == '*' {
		s = s[1:]
		wrong = true
	}

	words, err := gloss.ParseGloss(s)
	if err != nil {
		panic(err)
	}

	wordsLine := ""
	morphemesLine := ""
	glossLine := ""

	if wrong {
		wordsLine += "*"
		morphemesLine += "*"
		glossLine += " "
	}

	for i := 0; i < len(words); i++ {
		rom := words[i].Word.Romanization()
		morphemes := ""
		glosses := ""

		for _, gw := range words[i].Pieces {
			maxLen := max(len(gw.Form), len(gw.Gloss))
			morphemes += fmt.Sprintf("%-*s", maxLen, gw.Form)
			glosses += fmt.Sprintf("%-*s", maxLen, gw.Gloss)
		}

		maxLen := max(len(rom), len(morphemes))

		wordsLine += fmt.Sprintf("%-*s", maxLen+1, rom)
		morphemesLine += fmt.Sprintf("%-*s", maxLen+1, morphemes)
		glossLine += fmt.Sprintf("%-*s", maxLen+1, glosses)
	}

	wordsLine = strings.TrimSpace(wordsLine)
	morphemesLine = strings.TrimSpace(morphemesLine)
	glossLine = strings.TrimSpace(glossLine)

	return fmt.Sprintf("```\n%s\n%s\n%s\n%q\n```", wordsLine, morphemesLine, glossLine, translation)
}

func applyToChildren(n ast.Node, f func(child ast.Node)) {
	child := n.FirstChild()
	for i := 0; i < n.ChildCount(); i++ {
		f(child)
		child = child.NextSibling()
	}
}

func renderList(n *ast.List, source []byte, indent int) []byte {
	var out []byte

	applyToChildren(n, func(child ast.Node) {
		li := child.(*ast.ListItem)
		if li.HasBlankPreviousLines() {
			out = append(out, '\n')
		}
		if li.Offset > 0 {
			out = append(out, slices.Repeat([]byte{' '}, indent)...)
		}
		out = append(out, n.Marker)
		out = append(out, ' ')
		anyListChildren := false
		applyToChildren(li, func(child ast.Node) {
			switch p := child.(type) {
			case *ast.List:
				anyListChildren = true
				out = append(out, '\n')
				out = append(out, renderList(p, source, indent+li.Offset)...)
			default:
				out = append(out, RenderAsMarkdown(child, source)...)
			}
		})
		if !anyListChildren {
			out = append(out, '\n')
		}
	})

	return out
}

func RenderAsMarkdown(n ast.Node, source []byte) []byte {
	var out []byte

	switch p := n.(type) {
	case *ast.Document:
		applyToChildren(n, func(child ast.Node) {
			out = append(out, RenderAsMarkdown(child, source)...)
		})
	case *ast.Heading:
		out = append(out, slices.Repeat([]byte{'#'}, p.Level)...)
		out = append(out, ' ')
		lines := p.Lines()
		for i := 0; i < lines.Len(); i++ {
			line := lines.At(i)
			out = append(out, source[line.Start:line.Stop]...)
		}
		out = append(out, []byte("\n\n")...)
	case *ast.Text:
		s := p.Segment
		out = append(out, source[s.Start:s.Stop]...)
		if p.SoftLineBreak() {
			out = append(out, '\n')
		}
	case *ast.Paragraph:
		applyToChildren(n, func(child ast.Node) {
			out = append(out, RenderAsMarkdown(child, source)...)
		})
		if n.Parent().Kind().String() != "ListItem" {
			if n.NextSibling().Kind().String() != "List" {
				out = append(out, '\n')
			}
			out = append(out, '\n')
		}
	case *ast.Emphasis:
		stars := slices.Repeat([]byte{'*'}, p.Level)
		out = append(out, stars...)
		applyToChildren(n, func(child ast.Node) {
			out = append(out, RenderAsMarkdown(child, source)...)
		})
		out = append(out, stars...)
	case *ast.RawHTML:
		segments := p.Segments
		for i := 0; i < segments.Len(); i++ {
			s := segments.At(i)
			out = append(out, source[s.Start:s.Stop]...)
		}
	case *ast.List:
		out = append(out, renderList(p, source, 0)...)
		out = append(out, []byte("\n")...)
	case *ast.TextBlock:
		applyToChildren(n, func(child ast.Node) {
			out = append(out, RenderAsMarkdown(child, source)...)
		})
	case *ast.CodeSpan:
		outline := RenderAsMarkdown(n.FirstChild(), source)
		result := gl(string(outline))

		out = append(out, '*')
		out = append(out, []byte(result)...)
		out = append(out, '*')
	case *ast.FencedCodeBlock:
		lines := p.Lines()
		l0, l1 := lines.At(0), lines.At(1)
		outline := strings.TrimSuffix(string(source[l0.Start:l0.Stop]), "\n")
		translation := strings.TrimSuffix(string(source[l1.Start:l1.Stop]), "\n")

		out = append(out, []byte(glossedSentence(outline, translation)+"\n\n")...)
	default:
		panic(fmt.Errorf("unknown node kind: %q", p.Kind().String()))
	}

	return out
}

func GenerateDocs() {
	filesToMove := []string{"README", "North Wind"}

	for _, filename := range filesToMove {
		source, err := os.ReadFile(fmt.Sprintf("templates/%s.md", filename))
		if err != nil {
			panic(err)
		}
		reader := text.NewReader(source)
		p := goldmark.DefaultParser()
		n := p.Parse(reader)

		os.WriteFile(fmt.Sprintf("%s.md", filename), []byte(RenderAsMarkdown(n, source)), 0)
	}
}
