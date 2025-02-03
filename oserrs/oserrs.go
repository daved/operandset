package oserrs

import (
	"fmt"
	"reflect"
)

type Error struct {
	child error
}

func NewError(child error) *Error {
	return &Error{child}
}

func (e *Error) Error() string {
	return fmt.Sprintf("operandset: %v", e.child)
}

func (e *Error) Unwrap() error {
	return e.child
}

func (e *Error) Is(err error) bool {
	return reflect.TypeOf(e) == reflect.TypeOf(err)
}

type ParseError struct {
	child error
}

func NewParseError(child error) *ParseError {
	return &ParseError{child}
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("parse: %v", e.child)
}

func (e *ParseError) Unwrap() error {
	return e.child
}

func (e *ParseError) Is(err error) bool {
	return reflect.TypeOf(e) == reflect.TypeOf(err)
}

type ResolveError struct {
	child error
}

func NewResolveError(child error) *ResolveError {
	return &ResolveError{child}
}

func (e *ResolveError) Error() string {
	return fmt.Sprintf("resolve: %v", e.child)
}

func (e *ResolveError) Unwrap() error {
	return e.child
}

func (e *ResolveError) Is(err error) bool {
	return reflect.TypeOf(e) == reflect.TypeOf(err)
}

type HydrateError struct {
	Operand string
	child   error
}

func NewHydrateError(operand string, child error) *HydrateError {
	return &HydrateError{operand, child}
}

func (e *HydrateError) Error() string {
	return fmt.Sprintf("hydrate (%s): %v", e.Operand, e.child)
}

func (e *HydrateError) Unwrap() error {
	return e.child
}

func (e *HydrateError) Is(err error) bool {
	return reflect.TypeOf(e) == reflect.TypeOf(err)
}

type OperandMissingError struct {
	name string
}

func NewOperandMissingError(name string) *OperandMissingError {
	return &OperandMissingError{name}
}

func (e *OperandMissingError) Error() string {
	return fmt.Sprintf("missing an expected operand: %s", e.name)
}

func (e *OperandMissingError) Is(err error) bool {
	return reflect.TypeOf(e) == reflect.TypeOf(err)
}
