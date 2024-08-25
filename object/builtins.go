package object

const (
	BuiltFuncNameLen     = "len"
	BuiltInFuncNameFirst = "first"
	BuiltInFuncNameLast  = "last"
	BuiltInFuncNameRest  = "rest"
	BuiltinFuncNamePush  = "push"
	BuiltinFuncNamePuts  = "puts"
)

var builtins = []struct{
	Name string 
	Builtin *Builtin 
}{
	
}