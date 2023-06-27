package main

import (
	"os"

	engine "github.com/NekoMaru76/Crayon/engine"
	parser "github.com/NekoMaru76/Crayon/parser/grammars"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewCrayonLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCrayonParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	ch := make(chan interface{})
	eng := engine.NewEngine(ch)

	eng.AddCommands(
		engine.WriteCommand,
		engine.ByeCommand,
		engine.DefineVariableCommand,
		engine.AssignVariableCommand,
		engine.GlobalVariableCommand,
		engine.DeleteVariableCommand,
		engine.IfCommand,
		engine.IsEqualCommand,
		engine.IsNotEqualCommand,
		engine.NotCommand,
		engine.AndCommand,
		engine.OrCommand,
		engine.IsLessThanCommand,
		engine.IsLessThanOrEqualCommand,
		engine.IsMoreThanCommand,
		engine.IsMoreThanOrEqualCommand,
		engine.SubCommand,
		engine.SumCommand,
		engine.DivideCommand,
		engine.MultiplyCommand,
		engine.CallCustomTagCommand,
		engine.GetPropCommand,
		engine.SetPropCommand,
		engine.GetTypeofCommand,
	)

	script := p.Script().(*parser.ScriptContext)
	parentScope := engine.NewScope(nil)
	v := engine.NewCrayonVisitor(eng, engine.NewScope(parentScope), make(map[string]*engine.CustomTag))
	tags := v.VisitScript(script)

	v.RunAllAvailableTags(tags, make([][]any, len(tags)))

	defer func() {
		for range ch {
		}
	}()
}
