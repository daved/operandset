package operandset_test

import (
	"fmt"

	"github.com/daved/operandset"
)

func Example() {
	var (
		num  int
		info = "default-value"
	)

	os := operandset.New("app")
	os.Operand(&num, true, "number", "Number for printing.")
	os.Operand(&info, false, "information", "Info to use.")

	args := []string{"42", "non-default"}

	if err := os.Parse(args); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Num: %d, Info: %s\n", num, info)
	fmt.Println()
	fmt.Println(os.Usage())
	// Output:
	// Num: 42, Info: non-default
	//
	// Operands for app:
	//
	//     number  (required)
	//         Number for printing.
	//
	//     information
	//         Info to use.
}

func ExampleOperandSet_Unresolved() {
	var num int

	os := operandset.New("app")
	os.Operand(&num, true, "number", "Number for printing.")

	args := []string{"42", "unresolved-A", "unresolved-B"}

	if err := os.Parse(args); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Num: %d, Unresolved: %v\n", num, os.Unresolved())
	// Output:
	// Num: 42, Unresolved: [unresolved-A unresolved-B]
}
