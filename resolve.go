package operandset

import (
	er "github.com/daved/operandset/oserrs"
	"github.com/daved/vtypes"
)

func resolve(ops []*Operand, args []string) ([]string, error) {
	wrap := er.NewResolveError

	for i, op := range ops {
		if len(args) <= i {
			if !op.req {
				return nil, nil
			}

			return nil, wrap(er.ErrOperandRequired, op.name)
		}

		if err := vtypes.Hydrate(op.val, args[i]); err != nil {
			return args[i:], wrap(err, op.name)
		}
	}

	return args[len(ops):], nil
}
