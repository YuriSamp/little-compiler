package object

import (
	"bytes"
	"fmt"
	"monkey/ast"
	"strings"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	NULL_OBJ = "NULL"
	ERROR_OBJ = "ERROR"
	FUNCTION_OBJ = "FUNCTION"
)

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store : s}
}

type Environment struct {
	store map[string]Object
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

func(e *Environment) Get(name string) (Object, bool){
	obj, ok := e.store[name]
	return obj, ok
}

type Integer struct {
	Value int64
}

func(i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func(i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

type Boolean struct {
	Value bool
}

func(b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func(b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

type Null struct {}

func(n *Null) Inspect() string {
	return "null"
}

func(n *Null) Type() ObjectType {
	return NULL_OBJ
}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

func (e *Error) Inspect() string {
	return "ERROR" + e.Message
}

type Function struct {
	Parameters []*ast.Identifier
	Body * ast.BlockStatement
	Env *Environment
}

func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

func(f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}

	for _, p := range f.Parameters{
		params = append(params,p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}
