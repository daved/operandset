package operandset

import (
	"github.com/daved/operandset/oserrs"
)

// Error types forward basic error types from the oserrs package for access and
// documentation. If an error has interesting behavior, it should be defined
// directly in this package.
type (
	Error               = oserrs.Error
	ParseError          = oserrs.ParseError
	OperandMissingError = oserrs.OperandMissingError
	ConvertRawError     = oserrs.ConvertRawError
)
