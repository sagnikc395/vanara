package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/sagnikc395/vanara/ast"
)

type ObjectType string

const (
	IntegerObj          = "Integer"
	BooleanObj          = "Boolean"
	NullObj             = "Null"
	ReturnValueObj      = "ReturnValue"
	ErrorObj            = "Error"
	FunctionObj         = "Function"
	StringObj           = "String"
	BuiltinObj          = "Builtin"
	ArraysObj           = "Array"
	HashObj             = "Hash"
	CompiledFunctionObj = "CompiledFunction"
	ClosureObj          = "Closure"
)

type HashKey struct {
	Type  ObjectType
	Value uint64
}

type Hashable interface {
	HashKey() HashKey
}

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType {
	return IntegerObj
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BooleanObj }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

type Null struct{}

func (n *Null) Type() ObjectType {
	return NullObj
}
func (n *Null) Inspect() string {
	return "null"
}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return ReturnValueObj }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType {
	return ErrorObj
}

func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType {
	return FunctionObj
}

func (f *Function) Inspect() string {
	var out bytes.Buffer

	var params []string
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}
