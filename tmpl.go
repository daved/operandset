package operandset

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

// TmplData is the structure used for usage output templating. Custom template
// string values should be based on this type.
type TmplData struct {
	OperandSet *OperandSet
}

// TmplConfig tracks the template string and function map used for usage output
// templating.
type TmplConfig struct {
	Text string
	FMap template.FuncMap
}

// NewDefaultTmplConfig returns the default TmplConfig value. This can be used
// as an example of how to setup custom usage output templating.
func NewDefaultTmplConfig() *TmplConfig {
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

	tmplFMap := template.FuncMap{
		"NameHint": nameHintFn,
	}

	tmplText := strings.TrimSpace(`
{{- if .OperandSet.Operands -}}
Operands for {{.OperandSet.Name}}:
{{range $i, $op := .OperandSet.Operands}}
  {{if .}}  {{end}}{{if $op.Name}}{{NameHint $op}}{{end}}
        {{$op.Description}}
{{end}}{{else}}{{- end}}
`)

	return &TmplConfig{
		Text: tmplText,
		FMap: tmplFMap,
	}
}

func executeTmpl(tc *TmplConfig, data any) string {
	tmpl := template.New("operandset").Funcs(tc.FMap)

	buf := &bytes.Buffer{}

	tmpl, err := tmpl.Parse(tc.Text)
	if err != nil {
		fmt.Fprintf(buf, "%v\n", err)
		return buf.String()
	}

	if err := tmpl.Execute(buf, data); err != nil {
		fmt.Fprintf(buf, "%v\n", err)
		return buf.String()
	}

	return buf.String()
}
