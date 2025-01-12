package operandset

// Operand manages operand option data. The exported fields are for easy
// post-construction configuration.
type Operand struct {
	val  any
	name string
	desc string
	req  bool
	Meta map[string]any
}

func newOperand(val any, req bool, name, desc string) *Operand {
	return &Operand{
		val:  val,
		name: name,
		desc: desc,
		req:  req,
		Meta: map[string]any{},
	}
}

// Name returns the name.
func (o *Operand) Name() string {
	return o.name
}

// IsRequired returns whether the operand is required.
func (o *Operand) IsRequired() bool {
	return o.req
}

// Description returns the description string.
func (o *Operand) Description() string {
	return o.desc
}
