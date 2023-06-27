package engine

import (
	"fmt"
	"os"
	"time"
)

var DelayCommandRoute = Route{
	Paths: []Path{NewKeywordPath("wait"), NewKeywordPath("for"), NumberValuePath, NewKeywordPath("seconds")},
	Run: func(a []any, cv *CrayonVisitor) any {
		time.Sleep(time.Duration(a[0].(float64) * float64(time.Second)))

		return a[0]
	},
}
var DelayCommand = Command{
	Routes: []Route{DelayCommandRoute},
}

var IfCommandRoute = Route{
	Paths: []Path{NewKeywordPath("if"), BoolValuePath, NewKeywordPath("do"), ScopePath},
	Run: func(a []any, cv *CrayonVisitor) any {
		if a[0] == true {
			return a[1].(ScopeAsValue)()
		}

		return nil
	},
}
var IfElseCommandRoute = Route{
	Paths: []Path{NewKeywordPath("if"), BoolValuePath, NewKeywordPath("do"), ScopePath, NewKeywordPath("else"), NewKeywordPath("do"), ScopePath},
	Run: func(a []any, cv *CrayonVisitor) any {
		if a[0] == true {
			return a[1].(ScopeAsValue)()
		} else {
			return a[2].(ScopeAsValue)()
		}
	},
}
var IfCommand = Command{
	Routes: []Route{IfCommandRoute, IfElseCommandRoute},
}

var IsEqualCommandRoute = Route{
	Paths: []Path{NewKeywordPath("is"), AnyValuePath, NewKeywordPath("equal"), NewKeywordPath("with"), AnyValuePath},
	Run: func(a []any, cv *CrayonVisitor) any {
		return a[0] == a[1]
	},
}
var IsEqualCommand = Command{
	Routes: []Route{IsEqualCommandRoute},
}

var OrCommandRoute = Route{
	Paths: []Path{BoolValuePath, NewKeywordPath("or"), BoolValuePath},
	Run: func(a []any, cv *CrayonVisitor) any {
		return a[0].(bool) || a[1].(bool)
	},
}
var OrCommand = Command{
	Routes: []Route{OrCommandRoute},
}

var AndCommandRoute = Route{
	Paths: []Path{BoolValuePath, NewKeywordPath("and"), BoolValuePath},
	Run: func(a []any, cv *CrayonVisitor) any {
		return a[0].(bool) && a[1].(bool)
	},
}
var AndCommand = Command{
	Routes: []Route{AndCommandRoute},
}

var IsLessThanCommandRoute = Route{
	Paths: []Path{NewKeywordPath("is"), BoolValuePath, NewKeywordPath("less"), NewKeywordPath("than"), BoolValuePath},
	Run: func(a []any, cv *CrayonVisitor) any {
		return a[0].(float64) < a[1].(float64)
	},
}
var IsLessThanCommand = Command{
	Routes: []Route{IsLessThanCommandRoute},
}

var IsMoreThanCommandRoute = Route{
	Paths: []Path{NewKeywordPath("is"), BoolValuePath, NewKeywordPath("more"), NewKeywordPath("than"), BoolValuePath},
	Run: func(a []any, cv *CrayonVisitor) any {
		return a[0].(float64) > a[1].(float64)
	},
}
var IsMoreThanCommand = Command{
	Routes: []Route{IsMoreThanCommandRoute},
}

var IsLessThanOrEqualCommandRoute = Route{
	Paths: []Path{NewKeywordPath("is"), BoolValuePath, NewKeywordPath("less"), NewKeywordPath("than"), NewKeywordPath("or"), NewKeywordPath("equal"), NewKeywordPath("with"), BoolValuePath},
	Run: func(a []any, cv *CrayonVisitor) any {
		return a[0].(float64) <= a[1].(float64)
	},
}
var IsLessThanOrEqualCommand = Command{
	Routes: []Route{IsLessThanOrEqualCommandRoute},
}

var IsMoreThanOrEqualCommandRoute = Route{
	Paths: []Path{NewKeywordPath("is"), BoolValuePath, NewKeywordPath("more"), NewKeywordPath("than"), NewKeywordPath("or"), NewKeywordPath("equal"), NewKeywordPath("with"), BoolValuePath},
	Run: func(a []any, cv *CrayonVisitor) any {
		return a[0].(float64) >= a[1].(float64)
	},
}
var IsMoreThanOrEqualCommand = Command{
	Routes: []Route{IsMoreThanOrEqualCommandRoute},
}

var SumCommandRoute = Route{
	Paths: []Path{NewKeywordPath("add"), NumberValuePath, NewKeywordPath("by"), NumberValuePath},
	Run: func(a []any, cv *CrayonVisitor) any {
		return a[0].(float64) + a[1].(float64)
	},
}
var SumCommand = Command{
	Routes: []Route{SumCommandRoute},
}

var SubCommandRoute = Route{
	Paths: []Path{NewKeywordPath("sub"), NumberValuePath, NewKeywordPath("by"), NumberValuePath},
	Run: func(a []any, cv *CrayonVisitor) any {
		return a[0].(float64) - a[1].(float64)
	},
}
var SubCommand = Command{
	Routes: []Route{SubCommandRoute},
}

var DivideCommandRoute = Route{
	Paths: []Path{NewKeywordPath("divide"), NumberValuePath, NewKeywordPath("by"), NumberValuePath},
	Run: func(a []any, cv *CrayonVisitor) any {
		return a[0].(float64) / a[1].(float64)
	},
}
var DivideCommand = Command{
	Routes: []Route{DivideCommandRoute},
}

var MultiplyCommandRoute = Route{
	Paths: []Path{NewKeywordPath("multiply"), NumberValuePath, NewKeywordPath("by"), NumberValuePath},
	Run: func(a []any, cv *CrayonVisitor) any {
		return a[0].(float64) * a[1].(float64)
	},
}
var MultiplyCommand = Command{
	Routes: []Route{MultiplyCommandRoute},
}

var IsNotEqualCommandRoute = Route{
	Paths: []Path{NewKeywordPath("is"), AnyValuePath, NewKeywordPath("not"), NewKeywordPath("equal"), NewKeywordPath("with"), AnyValuePath},
	Run: func(a []any, cv *CrayonVisitor) any {
		return a[0] != a[1]
	},
}
var IsNotEqualCommand = Command{
	Routes: []Route{IsNotEqualCommandRoute},
}

var NotCommandRoute = Route{
	Paths: []Path{NewKeywordPath("not"), BoolValuePath},
	Run: func(a []any, cv *CrayonVisitor) any {
		return !a[0].(bool)
	},
}
var NotCommand = Command{
	Routes: []Route{NotCommandRoute},
}

var DefineVariableCommandRoute = Route{
	Paths: []Path{NewKeywordPath("define"), NewKeywordPath("variable"), VariablePath},
	Run: func(a []any, cv *CrayonVisitor) any {
		name := a[0].(Variable).Name
		err := cv.Scope.DefineVariable(name, nil)

		if err != nil {
			panic(err)
		}

		return name
	},
}
var DefineVariableCommand = Command{
	Routes: []Route{DefineVariableCommandRoute},
}

var GlobalVariableCommandRoute = Route{
	Paths: []Path{NewKeywordPath("make"), NewKeywordPath("variable"), VariablePath, NewKeywordPath("global")},
	Run: func(a []any, cv *CrayonVisitor) any {
		name := a[0].(Variable).Name
		err := cv.Scope.GlobalVariable(name)

		if err != nil {
			panic(err)
		}

		return name
	},
}
var GlobalVariableCommand = Command{
	Routes: []Route{GlobalVariableCommandRoute},
}

var AssignVariableCommandRoute = Route{
	Paths: []Path{NewKeywordPath("assign"), NewKeywordPath("value"), AnyValuePath, NewKeywordPath("to"), NewKeywordPath("variable"), VariablePath},
	Run: func(a []any, cv *CrayonVisitor) any {
		name := a[1].(Variable).Name
		err := cv.Scope.UpdateVariable(name, a[0])

		if err != nil {
			panic(err)
		}

		return name
	},
}
var AssignVariableCommand = Command{
	Routes: []Route{AssignVariableCommandRoute},
}

var DeleteVariableCommandRoute = Route{
	Paths: []Path{NewKeywordPath("delete"), NewKeywordPath("variable"), VariablePath},
	Run: func(a []any, cv *CrayonVisitor) any {
		name := a[0].(Variable).Name
		err := cv.Scope.DeleteVariable(name)

		if err != nil {
			panic(err)
		}

		return name
	},
}
var DeleteVariableCommand = Command{
	Routes: []Route{DeleteVariableCommandRoute},
}

var WriteCommandRoute = Route{
	Paths: []Path{NewKeywordPath("write"), AnyValuePath, NewKeywordPath("to"), NewKeywordPath("console")},
	Run: func(a []any, _ *CrayonVisitor) any {
		fmt.Println(a[0])

		return a
	},
}
var WriteCommand = Command{
	Routes: []Route{WriteCommandRoute},
}

var ByeCommandRoute = Route{
	Paths: []Path{NewKeywordPath("k"), NewKeywordPath("thx"), NewKeywordPath("bye")},
	Run: func(a []any, v *CrayonVisitor) any {
		os.Exit(0)

		return nil
	},
}
var ByeCommand = Command{
	Routes: []Route{ByeCommandRoute},
}

var CallCustomTagCommandRoute = Route{
	Paths: []Path{NewKeywordPath("call"), NewKeywordPath("tag"), ValueTagPath, NewKeywordPath("with"), NewKeywordPath("args"), ArrayValuePath},
	Run: func(a []any, v *CrayonVisitor) any {
		return a[0].(*CustomTag).Run(v, a[1].([]any))
	},
}
var CallCustomTagCommand = Command{
	Routes: []Route{CallCustomTagCommandRoute},
}

var GetPropArrayCommandRoute = Route{
	Paths: []Path{NewKeywordPath("get"), NewKeywordPath("property"), NumberValuePath, NewKeywordPath("from"), ArrayValuePath},
	Run: func(a []any, v *CrayonVisitor) any {
		return a[1].([]any)[int64(a[0].(float64))]
	},
}
var GetPropObjectCommandRoute = Route{
	Paths: []Path{NewKeywordPath("get"), NewKeywordPath("property"), StringValuePath, NewKeywordPath("from"), ObjectValuePath},
	Run: func(a []any, v *CrayonVisitor) any {
		return a[1].(map[string]any)[a[0].(string)]
	},
}
var GetPropCommand = Command{
	Routes: []Route{GetPropArrayCommandRoute, GetPropObjectCommandRoute},
}

var SetPropArrayCommandRoute = Route{
	Paths: []Path{NewKeywordPath("set"), AnyValuePath, NewKeywordPath("as"), NewKeywordPath("property"), NumberValuePath, NewKeywordPath("from"), ArrayValuePath},
	Run: func(a []any, v *CrayonVisitor) any {
		return a[1].([]any)[int64(a[0].(float64))]
	},
}
var SetPropObjectCommandRoute = Route{
	Paths: []Path{NewKeywordPath("set"), AnyValuePath, NewKeywordPath("as"), NewKeywordPath("property"), StringValuePath, NewKeywordPath("from"), ObjectValuePath},
	Run: func(a []any, v *CrayonVisitor) any {
		return a[1].(map[string]any)[a[0].(string)]
	},
}
var SetPropCommand = Command{
	Routes: []Route{SetPropArrayCommandRoute, SetPropObjectCommandRoute},
}

var GetTypeofCommandRoute = Route{
	Paths: []Path{NewKeywordPath("get"), NewKeywordPath("type"), StringValuePath, NewKeywordPath("of"), AnyValuePath},
	Run: func(a []any, v *CrayonVisitor) any {
		switch a[0].(type) {
		case []any:
			return "array"
		case CustomTag:
			return "tag"
		case bool:
			return "bool"
		case string:
			return "string"
		case float64:
			return "number"
		case map[string]any:
			return "object"
		case nil:
			return "none"
		}

		return "unknown"
	},
}
var GetTypeofCommand = Command{
	Routes: []Route{GetTypeofCommandRoute},
}
