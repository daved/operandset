package operandset

import (
	er "github.com/daved/operandset/oserrs"
	"github.com/daved/vtype"
)

func resolve(ops []*Operand, args []string) error {
	wrap := er.NewResolveError

	for i, op := range ops {
		if len(args) <= i {
			if !op.req {
				continue
			}

			return wrap(er.ErrOperandRequired, op.name)
		}

		if err := vtype.Hydrate(op.val, args[i]); err != nil {
			return wrap(err, op.name)
		}
	}

	return nil
}
