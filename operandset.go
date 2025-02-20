// Package operandset provides simple, POSIX-friendly operand parsing. It is
// similar to and intended to be used with the [github.com/daved/flagset]
// package.
//
// CLI command arguments that are not subcommands or flag-related (no hyphen
// prefix, not a flag value) should be provided to be processed as operands.
// Typically, operands are the most interesting values used in the behavior
// being executed by a command.
package operandset

import (
	er "github.com/daved/operandset/oserrs"
	"github.com/daved/vtype"
)

// OperandSet contains operand options and usage-related values. Exported fields
// are used for easy post-construction configuraiton.
type OperandSet struct {
	// Templating
	Tmpl *Tmpl // set to NewUsageTmpl by default
	Meta map[string]any

	name string
	ops  []*Operand
	raws []string
}

// New constructs an OperandSet. Package convention is to name the operandset
// after the command that the operands are being associated with.
func New(name string) *OperandSet {
	os := &OperandSet{
		name: name,
		Meta: map[string]any{},
	}

	os.Tmpl = NewUsageTmpl(os)

	return os
}

// Name returns the name of the OperandSet  set during construction.
func (os *OperandSet) Name() string {
	return os.name
}

// Operands returns all operand options that have been set.
func (os *OperandSet) Operands() []*Operand {
	return os.ops
}

// Operand adds an operand option to the OperandSet. See [vtype.Hydrate] for
// details about which value types are supported. Functions compatible with
// [vtype] typed functions will be auto-converted.
func (os *OperandSet) Operand(val any, req bool, name, desc string) *Operand {
	val = vtype.ConvertCompatible(val)

	o := newOperand(val, req, name, desc)
	os.ops = append(os.ops, o)

	return o
}

// Parse processes operand values from the argument list, which must not include
// the initial command name. Parse must be called after all operands in the
// OperandSet are defined and before operand value access.
func (os *OperandSet) Parse(args []string) error {
	os.raws = args

	if err := resolve(os.ops, args); err != nil {
		return er.NewError(er.NewParseError(err))
	}

	return nil
}

// Parsed returns the args provided to Parse.
func (os *OperandSet) Parsed() []string {
	return os.raws
}

// Usage returns usage text. The default template construction function
// ([NewUsageTmpl]) can be used as a reference for custom templates which should
// be used to set the "Tmpl" field on OperandSet.
func (os *OperandSet) Usage() string {
	return os.Tmpl.String()
}
