package operandset

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

func (o *Operand) Name() string {
	return o.name
}

func (o *Operand) Required() bool {
	return o.req
}

func (o *Operand) Description() string {
	return o.desc
}
