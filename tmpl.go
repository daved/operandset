package operandset

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

// Tmpl holds template configuration details.
type Tmpl struct {
	Text string
	FMap template.FuncMap
	Data any
}

// Execute parses the template text and funcmap, then executes it using the set
// data.
func (t *Tmpl) Execute() (string, error) {
	tmpl := template.New("clic").Funcs(t.FMap)

	buf := &bytes.Buffer{}

	tmpl, err := tmpl.Parse(t.Text)
	if err != nil {
		return "", err
	}

	if err := tmpl.Execute(buf, t.Data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// String calls the Execute method returning either the validly executed
// template output or error message text.
func (t *Tmpl) String() string {
	s, err := t.Execute()
	if err != nil {
		s = fmt.Sprintf("%v\n", err)
	}
	return s
}

// NewUsageTmpl returns the default template configuration. This can be used as
// an example of how to setup custom usage output templating.
func NewUsageTmpl(os *OperandSet) *Tmpl {
	type tmplData struct {
		OperandSet *OperandSet
	}

	data := &tmplData{
		OperandSet: os,
	}

	nameHintFn := func(o *Operand) string {
		if o.name == "" {
			return ""
		}

		var post string
		if o.req {
			post = "  (required)"
		}

		return o.name + post
	}

	fMap := template.FuncMap{
		"NameHint": nameHintFn,
	}

	text := strings.TrimSpace(`
{{- if .OperandSet.Operands -}}
Operands for {{.OperandSet.Name}}:
{{range $i, $op := .OperandSet.Operands}}
  {{if .}}  {{end}}{{if $op.Name}}{{NameHint $op}}{{end}}
        {{$op.Description}}
{{end}}{{else}}{{- end}}
`)

	return &Tmpl{text, fMap, data}
}
