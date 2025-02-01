# operandset [![GoDoc](https://pkg.go.dev/badge/github.com/daved/operandset.svg)](https://pkg.go.dev/github.com/daved/operandset)

```go
go get github.com/daved/operandset
```

Package operandset is similar to the standard library flag package. Instead of flags, operands are
the focus.

## Usage

```go
type Operand
    func (o *Operand) Description() string
    func (o *Operand) IsRequired() bool
    func (o *Operand) Name() string
type OperandSet
    func New(name string) *OperandSet
    func (os *OperandSet) Name() string
    func (os *OperandSet) Operand(val any, req bool, name, desc string) *Operand
    func (os *OperandSet) Operands() []*Operand
    func (os *OperandSet) Parse(args []string) error
    func (os *OperandSet) Parsed() []string
    func (os *OperandSet) SetUsageTemplating(tmplCfg *TmplConfig)
    func (os *OperandSet) Usage() string
// see package docs for more
```

### Setup

```go
func main() {
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
}
```

## More Info

### Operands

Operands are the non-flag, non-command args in a CLI command that are at the end of the arg set.
Operands are normally treated as the important values used in the behavior being executed by a
particular CLI command.

### Supported Operand Value Types

- builtin: *string, *bool, *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32,
*uint64, *float32, *float64
- stdlib: *time.Duration, flag.Value
- vtype: vtype.TextMarshalUnmarshaler, vtype.OperandFunc

#### `vtype` Types

```go
type OperandFunc
type TextMarshalUnmarshaler
```

TextMarshalUnmarshaler describes types which satisfy both the encoding.TextMarshaler and
encoding.TextUnmarshaler interfaces, and is offered so that callers can easily use standard library
compatible types. OperandFunc describes functions that are able to be called when the related
operand is parsed. Compatible functions will be automatically converted to OperandFunc.

```go
func main() {
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
}
```
Output:
```txt
Operand Value: something
```

### Default Templating

`os.Usage()` value from the usage example above:

```txt
Operands for app:

    number  (required)
        Number for printing.

    information
        Info to use.
```

### Custom Templating

Custom templates and template behaviors (i.e. template function maps) can be set. Custom data can be
attached to instances of OperandSet, and Operand using their Meta fields for access from custom
templates.
