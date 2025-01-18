package vtype

import "encoding"

// OperandFunc describes functions that can be called when an operand is
// succesfully parsed.
type OperandFunc func(string) error

// TextMarshalUnmarshaler descibes types that are able to be marshaled to and
// unmarshaled from text.
type TextMarshalUnmarshaler interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}
