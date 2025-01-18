package operandset_test

import (
	"fmt"

	"github.com/daved/operandset"
)

func Example_operandFunc() {
	do := func(operandVal string) error {
		fmt.Println("Operand Value:", operandVal)
		return nil
	}

	os := operandset.New("app")
	os.Operand(do, true, "first_operand", "Run callback.")

	args := []string{"something"}

	if err := os.Parse(args); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println(os.Usage())

	// Output:
	// Operand Value: something
	//
	// Operands for app:
	//
	//     first_operand  (required)
	//         Run callback.
}
