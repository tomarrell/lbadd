package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
	"sort"
	"strings"
)

var (
	tmpl = template.Must(template.New("_scanKeyword").Funcs(template.FuncMap{
		"lower":    strings.ToLower,
		"sanitize": sanitize,
	}).Parse(`

func scanKeyword{{ sanitize .path }}(s RuneScanner) (token.Type, bool) {
	{{ if .isLeaf }}return {{ .tokenType }}, true{{ else }}
	{{- if .nextRunes }}next{{ else }}_{{ end }}, ok := s.Lookahead()
	if !ok {
		{{ if .tokenType }}return {{ .tokenType }}, true{{ else }}return token.Unknown, false{{ end }}
	}{{ if .nextRunes }}
	switch next { {{- range .nextRunes }}
	{{ $low := lower . }}
	case '{{ . }}'{{ if eq . $low }}{{ else }}, '{{ $low }}'{{ end }}:
		s.ConsumeRune()
		return scanKeyword{{ sanitize $.path }}{{ sanitize . }}(s){{ end }}
	}{{ end }}
	{{ if .hasValue }}return {{ .tokenType }}, true{{ else }}return token.Unknown, false{{ end }}{{ end }}
}`))
	header = `// Code generated with internal/tool/generate/keywordtrie; DO NOT EDIT.

package ruleset

import "github.com/tomarrell/lbadd/internal/parser/scanner/token"

func defaultKeywordsRule(s RuneScanner) (token.Type, bool) {
	tok, ok := scanKeyword(s)
	if !ok {
		return token.Unknown, false
	}
	peek, noEof := s.Lookahead()
	if noEof && defaultLiteral.Matches(peek) { // keywords must be terminated with a whitespace
		return token.Unknown, false
	}
	return tok, ok
}`
)

func main() {
	if len(os.Args[1:]) != 1 {
		exitErr(fmt.Errorf("output file must be specified (eaxctly 1 argument required)"))
	}

	t := newTrie()
	for k, v := range keywordTokens {
		t.Put(k, v)
	}

	var buf bytes.Buffer
	buf.WriteString(header)
	genTrie([]rune{}, t, &buf)

	f, err := os.Create(os.Args[1])
	if err != nil {
		exitErr(err)
	}

	if err := f.Truncate(0); err != nil {
		exitErr(err)
	}

	if _, err = f.Seek(0, io.SeekStart); err != nil {
		exitErr(err)
	}

	rd := bytes.NewReader(buf.Bytes())
	if _, err = rd.WriteTo(f); err != nil {
		exitErr(err)
	}
}

func exitErr(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

func genTrie(path []rune, t *trie, buf io.Writer) {
	data := map[string]interface{}{
		"path": string(path),
	}
	if t.val != nil {
		data["tokenType"] = fmt.Sprintf("token.%s", t.val)
		if len(t.sub) == 0 {
			data["isLeaf"] = true
		} else {
			data["hasValue"] = true
		}
	}
	nextRunes := []string{}
	for k := range t.sub {
		nextRunes = append(nextRunes, string(k))
	}
	sort.Strings(nextRunes)
	data["nextRunes"] = nextRunes

	if err := tmpl.Execute(buf, data); err != nil {
		exitErr(err)
	}

	for _, nextRune := range nextRunes {
		r := []rune(nextRune)[0]
		genTrie(append(path, r), t.sub[r], buf)
	}
}

func sanitize(s string) string {
	return strings.ReplaceAll(s, "_", "x")
}
