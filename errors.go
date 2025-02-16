package operandset

import (
	"github.com/daved/operandset/oserrs"
	"github.com/daved/vtype"
)

// Error types forward basic error types from sub-packages for access and
// documentation. If an error has interesting behavior, it should be defined
// directly in this package.
type (
	Error        = oserrs.Error
	ParseError   = oserrs.ParseError
	ResolveError = oserrs.ResolveError
)

var (
	ErrOperandMissing  = oserrs.ErrOperandMissing
	ErrTypeUnsupported = vtype.ErrTypeUnsupported
)
