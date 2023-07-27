package evaluator

import "github.com/emilioziniades/interpreterbook/object"

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: len(arg.Value)}
			default:
				return newError("argument to `len` is not supported. got %s", arg.Type())
			}
		},
	},
}
