// Package operandset is similar to the standard library flag package. Instead
// of flags, operands are the focus. Operands are the non-flag, non-subcommand
// args in a CLI command that are typically at the end of the arg list. Operands
// are normally treated as the important values used by the behavior being
// executed by the particular CLI command.
package operandset

import (
	"flag"
	"strconv"
	"time"

	er "github.com/daved/operandset/oserrs"
	"github.com/daved/operandset/vtype"
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
//   - builtin: *string, *bool, *int, *int64, *uint, *uint64, *float64
//   - stdlib: *[time.Duration], [flag.Value]
//   - vtype: [vtype.TextMarshalUnmarshaler], [vtype.OperandFunc]
func (os *OperandSet) Operand(val any, req bool, name, desc string) *Operand {
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

	if err := parse(os.ops, args); err != nil {
		return er.NewError(err)
	}

	return nil
}

func parse(ops []*Operand, args []string) error {
	newError := er.NewParseError

	for i, op := range ops {
		if len(args) <= i {
			if !op.req {
				continue
			}

			return newError(er.NewOperandMissingError(op.name))
		}

		raw := args[i]

		switch v := op.val.(type) {
		case *string:
			*v = raw

		case *bool:
			b, err := strconv.ParseBool(raw)
			if err != nil {
				return newError(er.NewConvertRawError(err))
			}
			*v = b

		case *int:
			n, err := strconv.Atoi(raw)
			if err != nil {
				return newError(er.NewConvertRawError(err))
			}
			*v = n

		case *int64:
			n, err := strconv.ParseInt(raw, 10, 0)
			if err != nil {
				return newError(er.NewConvertRawError(err))
			}
			*v = n

		case *uint:
			n, err := strconv.ParseUint(raw, 10, 0)
			if err != nil {
				return newError(er.NewConvertRawError(err))
			}
			*v = uint(n)

		case *uint64:
			n, err := strconv.ParseUint(raw, 10, 0)
			if err != nil {
				return newError(er.NewConvertRawError(err))
			}
			*v = n

		case *float64:
			f, err := strconv.ParseFloat(raw, 64)
			if err != nil {
				return newError(er.NewConvertRawError(err))
			}
			*v = f

		case *time.Duration:
			d, err := time.ParseDuration(raw)
			if err != nil {
				return newError(er.NewConvertRawError(err))
			}
			*v = d

		case vtype.TextMarshalUnmarshaler:
			if err := v.UnmarshalText([]byte(raw)); err != nil {
				return newError(er.NewConvertRawError(err))
			}

		case flag.Value:
			if err := v.Set(raw); err != nil {
				return newError(er.NewConvertRawError(err))
			}

		case vtype.OperandFunc:
			if err := v(raw); err != nil {
				return newError(er.NewConvertRawError(err))
			}
		}
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
