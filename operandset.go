// Package operandset is similar to the [github.com/daved/flagset] package.
// Instead of flags, operands are the focus.
//
// Operands are the non-flag, non-command args in a CLI command that are at the
// end of the arg set. Operands are normally treated as the important values
// used in the behavior being executed by a particular CLI command.
package operandset

import (
	er "github.com/daved/operandset/oserrs"
	"github.com/daved/vtype"
)

// OperandSet contains operand options and related information used for usage
// output. The exported fields are for easy post-construction configuraiton.
type OperandSet struct {
	name    string
	ops     []*Operand
	raws    []string
	tmplCfg *TmplConfig
	Meta    map[string]any
}

// New constructs an OperandSet. In this package, it is conventional to name the
// operandset after the command that the options are being associated with.
func New(name string) *OperandSet {
	return &OperandSet{
		name:    name,
		tmplCfg: NewDefaultTmplConfig(),
		Meta:    map[string]any{},
	}
}

// Name returns the name set during construction.
func (os *OperandSet) Name() string {
	return os.name
}

// Operands returns all operand options that have been set.
func (os *OperandSet) Operands() []*Operand {
	return os.ops
}

// Operand adds an operand option to the OperandSet.
// Valid values are:
//   - builtin: *string, *bool, *int, *int8, *int16, *int32, *int64, *uint,
//     *uint8, *uint16, *uint32, *uint64, *float32, *float64
//   - stdlib: *[time.Duration], [flag.Value]
//   - vtype: [vtype.TextMarshalUnmarshaler], [vtype.OperandFunc]
func (os *OperandSet) Operand(val any, req bool, name, desc string) *Operand {
	val = vtype.ConvertCompatible(val)

	o := newOperand(val, req, name, desc)
	os.ops = append(os.ops, o)

	return o
}

// Parse parses operand definitions from the argument list, which must not
// include the initial command name. Parse must be called after all operands in
// the OperandSet are defined and before operand values are accessed by the
// program.
func (os *OperandSet) Parse(args []string) error {
	os.raws = args

	if err := resolve(os.ops, args); err != nil {
		return er.NewError(er.NewParseError(err))
	}

	return nil
}

// Parsed returns the args provided when Parse was called. The returned value
// can be helpful for debugging.
func (os *OperandSet) Parsed() []string {
	return os.raws
}

// SetUsageTemplating is used to override the base template text, and provide a
// custom FuncMap. If a nil FuncMap is provided, no change will be made to the
// existing value.
func (os *OperandSet) SetUsageTemplating(tmplCfg *TmplConfig) {
	os.tmplCfg = tmplCfg
}

// Usage returns the executed usage template. Each Operand type's Meta field can
// be leveraged to convey detailed info/behavior in a custom template.
func (os *OperandSet) Usage() string {
	return executeTmpl(os.tmplCfg, &TmplData{OperandSet: os})
}
