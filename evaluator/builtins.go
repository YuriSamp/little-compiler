package evaluator

import (
	"monkey/object"
)

var builtins = map[string]*object.BuiltIn{
	"len":   object.GetBuiltInByname("len"),
	"puts":  object.GetBuiltInByname("puts"),
	"first": object.GetBuiltInByname("first"),
	"last":  object.GetBuiltInByname("last"),
	"rest":  object.GetBuiltInByname("rest"),
	"push":  object.GetBuiltInByname("push"),
}
