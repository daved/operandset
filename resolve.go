package operandset

import (
	"flag"
	"strconv"
	"time"

	er "github.com/daved/operandset/oserrs"
	"github.com/daved/operandset/vtype"
)

func resolve(ops []*Operand, args []string) error {
	wrap := er.NewResolveError

	for i, op := range ops {
		if len(args) <= i {
			if !op.req {
				continue
			}

			return wrap(er.NewOperandMissingError(op.name))
		}

		if err := hydrate(op, args[i]); err != nil {
			return wrap(err)
		}
	}

	return nil
}

func hydrate(op *Operand, raw string) error {
	wrap := er.NewOperandHydrateError

	switch v := op.val.(type) {
	case *string:
		*v = raw

	case *bool:
		b, err := strconv.ParseBool(raw)
		if err != nil {
			return wrap(op.name, err)
		}
		*v = b

	case *int:
		n, err := strconv.Atoi(raw)
		if err != nil {
			return wrap(op.name, err)
		}
		*v = n

	case *int64:
		n, err := strconv.ParseInt(raw, 10, 0)
		if err != nil {
			return wrap(op.name, err)
		}
		*v = n

	case *int8:
		n, err := strconv.ParseInt(raw, 10, 8)
		if err != nil {
			return wrap(op.name, err)
		}
		*v = int8(n)

	case *int16:
		n, err := strconv.ParseInt(raw, 10, 16)
		if err != nil {
			return wrap(op.name, err)
		}
		*v = int16(n)

	case *int32:
		n, err := strconv.ParseInt(raw, 10, 32)
		if err != nil {
			return wrap(op.name, err)
		}
		*v = int32(n)

	case *uint:
		n, err := strconv.ParseUint(raw, 10, 0)
		if err != nil {
			return wrap(op.name, err)
		}
		*v = uint(n)

	case *uint64:
		n, err := strconv.ParseUint(raw, 10, 0)
		if err != nil {
			return wrap(op.name, err)
		}
		*v = n

	case *uint8:
		n, err := strconv.ParseUint(raw, 10, 8)
		if err != nil {
			return wrap(op.name, err)
		}
		*v = uint8(n)

	case *uint16:
		n, err := strconv.ParseUint(raw, 10, 16)
		if err != nil {
			return wrap(op.name, err)
		}
		*v = uint16(n)

	case *uint32:
		n, err := strconv.ParseUint(raw, 10, 32)
		if err != nil {
			return wrap(op.name, err)
		}
		*v = uint32(n)

	case *float64:
		f, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return wrap(op.name, err)
		}
		*v = f

	case *float32:
		f, err := strconv.ParseFloat(raw, 32)
		if err != nil {
			return wrap(op.name, err)
		}
		*v = float32(f)

	case *time.Duration:
		d, err := time.ParseDuration(raw)
		if err != nil {
			return wrap(op.name, err)
		}
		*v = d

	case vtype.TextMarshalUnmarshaler:
		if err := v.UnmarshalText([]byte(raw)); err != nil {
			return wrap(op.name, err)
		}

	case flag.Value:
		if err := v.Set(raw); err != nil {
			return wrap(op.name, err)
		}

	case vtype.OperandFunc:
		if err := v(raw); err != nil {
			return wrap(op.name, err)
		}
	}

	return nil
}
