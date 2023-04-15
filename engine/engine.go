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

type any interface{}

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

func (e *Engine) ParseCommand(visitor *CrayonVisitor, commandCtx *parser.CommandContext) any {
	ctxPaths := commandCtx.Paths

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
						val := visitor.VisitValuePath(ctxPath.(*parser.PathContext).ValuePath().(*parser.ValuePathContext))

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

	panic(fmt.Errorf("unknown command: '%s'", commandCtx.GetText()))
}

type CrayonErrorListener struct {
	*antlr.DefaultErrorListener
}

func (el *CrayonErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	panic(fmt.Sprintf("syntax error at line %d:%d, %s", line, column, msg))
}

type Scope struct {
	Variables   map[string]any
	ParentScope *Scope
}

func NewScope(parentScope *Scope) *Scope {
	return &Scope{
		Variables:   make(map[string]any),
		ParentScope: parentScope,
	}
}

func (s *Scope) GetVariable(name string) (any, error) {
	value, ok := s.Variables[name]
	if ok {
		return value, nil
	}

	if s.ParentScope != nil {
		return s.ParentScope.GetVariable(name)
	}

	return nil, fmt.Errorf("variable '%s' not found", name)
}

func (s *Scope) SetVariable(name string, value any) {
	s.Variables[name] = value
}

func (s *Scope) DefineVariable(name string, value any) error {
	if _, ok := s.Variables[name]; ok {
		return fmt.Errorf("variable '%s' already defined", name)
	}

	s.Variables[name] = value
	return nil
}

func (s *Scope) UpdateVariable(name string, value any) error {
	if _, ok := s.Variables[name]; !ok {
		if s.ParentScope != nil {
			return s.ParentScope.UpdateVariable(name, value)
		} else {
			return fmt.Errorf("variable '%s' not found", name)
		}
	}

	s.Variables[name] = value
	return nil
}

func (s *Scope) DeleteVariable(name string) error {
	if _, ok := s.Variables[name]; !ok {
		if s.ParentScope != nil {
			return s.ParentScope.DeleteVariable(name)
		} else {
			return fmt.Errorf("variable '%s' not found", name)
		}
	}

	delete(s.Variables, name)
	return nil
}
