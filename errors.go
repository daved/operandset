package operandset

import (
	"github.com/daved/operandset/oserrs"
	"github.com/daved/vtypes"
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
	ErrOperandRequired  = oserrs.ErrOperandRequired
	ErrTypeUnsupported  = vtypes.ErrTypeUnsupported
	ErrValueUnsupported = vtypes.ErrValueUnsupported
)
