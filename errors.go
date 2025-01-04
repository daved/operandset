package operandset

import (
	"github.com/daved/operandset/oserrs"
)

type (
	Error               = oserrs.Error
	ParseError          = oserrs.ParseError
	OperandMissingError = oserrs.OperandMissingError
	ConvertRawError     = oserrs.ConvertRawError
)
