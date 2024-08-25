package object

import "fmt"

const (
	BuiltFuncNameLen     = "len"
	BuiltInFuncNameFirst = "first"
	BuiltInFuncNameLast  = "last"
	BuiltInFuncNameRest  = "rest"
	BuiltinFuncNamePush  = "push"
	BuiltinFuncNamePuts  = "puts"
)

func newerror(format string, a ...interface{}) *Error {
	return &Error{
		Message: fmt.Sprintf(format, a...),
	}
}

func GetBuiltinByName(name string) *Builtin {
	for _, def := range Builtins {
		if def.Name == name {
			return def.Builtin
		}
	}
	return nil
}

var builtins = []struct {
	Name    string
	Builtin *Builtin
}{
	{
		BuiltinFuncNameLen,
		&Builtin{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newerror("wrong number of arguments. got=%d, want=1",
						len(args))
				}
			},
		},
	},
}
