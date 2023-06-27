package engine

import (
	"fmt"
	"strings"
	"time"

	parser "github.com/NekoMaru76/Crayon/parser/grammars"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Route struct {
	Paths []Path
	Run   func([]any, *CrayonVisitor) any
}

type Variable struct {
	Name string
}

type ScopeAsValue = func() any

type Path interface{}

type KeywordPath struct {
	Path
	Keyword string
}

type ValuePath struct {
	Path
	Type string
	Test func(any) bool
}

func NewKeywordPath(keyword string) Path {
	return KeywordPath{
		Keyword: keyword,
	}
}

func NewValuePath(typeName string, test func(any) bool) Path {
	return ValuePath{
		Type: typeName,
		Test: test,
	}
}

var AnyValuePath ValuePath = NewValuePath("Any", func(a any) bool {
	return true
}).(ValuePath)
var StringValuePath ValuePath = NewValuePath("String", func(a any) bool {
	_, ok := a.(string)

	return ok
}).(ValuePath)
var NumberValuePath ValuePath = NewValuePath("Number", func(a any) bool {
	_, ok := a.(float64)

	return ok
}).(ValuePath)
var ObjectValuePath ValuePath = NewValuePath("Object", func(a any) bool {
	_, ok := a.(map[string]any)

	return ok
}).(ValuePath)
var ArrayValuePath ValuePath = NewValuePath("Array", func(a any) bool {
	_, ok := a.([]any)

	return ok
}).(ValuePath)
var VariablePath ValuePath = NewValuePath("Variable", func(a any) bool {
	_, ok := a.(Variable)

	return ok
}).(ValuePath)
var BoolValuePath ValuePath = NewValuePath("Bool", func(a any) bool {
	_, ok := a.(bool)

	return ok
}).(ValuePath)
var ScopePath ValuePath = NewValuePath("Scope", func(a any) bool {
	_, ok := a.(ScopeAsValue)

	return ok
}).(ValuePath)
var ValueTagPath ValuePath = NewValuePath("Tag", func(a any) bool {
	_, ok := a.(*CustomTag)

	return ok
}).(ValuePath)

type Engine struct {
	Commands []Command
	Start    int64
	Chan     chan interface{}
	CountTag int64
	MaxTag   int64
}

type Command struct {
	Routes []Route
}

func NewEngine(Chan chan interface{}) *Engine {
	return &Engine{
		Start:    time.Now().UnixMilli(),
		Chan:     Chan,
		CountTag: 0,
		MaxTag:   0,
	}
}

func (e *Engine) StartTag() {
	e.MaxTag++
}

func (e *Engine) EndTag() {
	e.CountTag++

	if e.CountTag >= e.MaxTag {
		close(e.Chan)
	}
}

func (e *Engine) AddCommand(cmd Command) {
	e.Commands = append(e.Commands, cmd)
}

func (e *Engine) AddCommands(cmds ...Command) {
	e.Commands = append(e.Commands, cmds...)
}

func (e *Engine) ParseCommand(visitor *CrayonVisitor, commandCtx *parser.CommandContext) any {
	ctxPaths := commandCtx.AllPath()

	for _, cmd := range e.Commands {
		for _, route := range cmd.Routes {
			if len(ctxPaths) == len(route.Paths) {
				args := []any{}
				skip := false

				for i, ctxPath := range ctxPaths {
					path := route.Paths[i]

					switch p := path.(type) {
					case KeywordPath:
						word := strings.ToUpper(ctxPath.GetText())

						if word == strings.ToUpper(p.Keyword) {
							continue
						} else {
							skip = true
							break
						}
					case ValuePath:
						vP := ctxPath.(*parser.PathContext).ValuePath()

						if vP == nil {
							skip = true
							break
						}

						val := visitor.VisitValuePath(vP.(*parser.ValuePathContext))

						if v, ok := val.(Variable); ok && p.Type != VariablePath.Type {
							var er error
							val, er = visitor.GetVariable(v.Name)

							if er != nil {
								panic(er)
							}
						}

						ok := p.Test(val)

						if !ok {
							skip = true
							break
						}

						args = append(args, val)
					}
				}

				if !skip {
					return route.Run(args, visitor)
				}
			}
		}
	}

	panic(fmt.Errorf("Unknown command: '%s'", BeautifyCommandContext(commandCtx)))
}

type CrayonErrorListener struct {
	*antlr.DefaultErrorListener
}

func (el *CrayonErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	panic(fmt.Sprintf("Syntax error at line %d:%d, %s", line, column, msg))
}
