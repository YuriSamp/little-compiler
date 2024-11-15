package object

import "fmt"

var Builtins = []struct {
	Name    string
	BuiltIn *BuiltIn
}{
	{
		"len",
		&BuiltIn{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got=%d, want=1", len(args))
				}
				switch arg := args[0].(type) {
				case *Array:
					return &Integer{Value: int64(len(arg.Elements))}
				case *String:
					return &Integer{Value: int64(len(arg.Value))}
				default:
					return newError("argument to `len` not suported, got %s", args[0].Type())
				}
			},
		},
	},
	{
		"puts",
		&BuiltIn{
			Fn: func(args ...Object) Object {
				for _, arg := range args {
					fmt.Print(arg.Inspect())
				}
				return nil
			},
		},
	},
	{
		"first",
		&BuiltIn{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got=%d, want 1", len(args))
				}
				if args[0].Type() != ARRAY_OBJ {
					return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
				}

				arr := args[0].(*Array)
				if len(arr.Elements) > 0 {
					return arr.Elements[0]
				}
				return nil
			},
		},
	},
	{
		"last",
		&BuiltIn{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got=%d, want 1", len(args))
				}
				if args[0].Type() != ARRAY_OBJ {
					return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
				}

				arr := args[0].(*Array)
				length := len(arr.Elements)
				if length > 0 {
					return arr.Elements[length-1]
				}
				return nil
			},
		},
	},
	{
		"rest",
		&BuiltIn{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got=%d, want 1", len(args))
				}
				if args[0].Type() != ARRAY_OBJ {
					return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
				}

				arr := args[0].(*Array)
				length := len(arr.Elements)
				if length > 0 {
					newElements := make([]Object, length-1)
					copy(newElements, arr.Elements[1:length])
					return &Array{Elements: newElements}
				}
				return nil
			},
		},
	},
	{
		"push",
		&BuiltIn{
			Fn: func(args ...Object) Object {
				if args[0].Type() != ARRAY_OBJ {
					return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
				}

				arr := args[0].(*Array)
				length := len(arr.Elements)

				newElements := make([]Object, length+1)

				copy(newElements, arr.Elements)

				newElements[length] = args[1]

				return &Array{Elements: newElements}
			},
		},
	},
}

func GetBuiltInByname(name string) *BuiltIn {
	for _, def := range Builtins {
		if def.Name == name {
			return def.BuiltIn
		}
	}

	return nil
}

func newError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}
