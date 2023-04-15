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

	eng.AddCommand(engine.WriteCommand)
	eng.AddCommand(engine.ByeCommand)
	eng.AddCommand(engine.DefineVariableCommand)
	eng.AddCommand(engine.AssignVariableCommand)
	eng.AddCommand(engine.DeleteVariableCommand)
	eng.AddCommand(engine.IfCommand)
	eng.AddCommand(engine.IsEqualCommand)
	eng.AddCommand(engine.IsNotEqualCommand)
	eng.AddCommand(engine.NotCommand)
	eng.AddCommand(engine.AndCommand)
	eng.AddCommand(engine.OrCommand)
	eng.AddCommand(engine.IsEqualCommand)
	eng.AddCommand(engine.IsLessThanCommand)
	eng.AddCommand(engine.IsLessThanOrEqualCommand)
	eng.AddCommand(engine.IsMoreThanCommand)
	eng.AddCommand(engine.IsMoreThanOrEqualCommand)
	eng.AddCommand(engine.IsNotEqualCommand)
	eng.AddCommand(engine.SubCommand)
	eng.AddCommand(engine.SumCommand)
	eng.AddCommand(engine.DivideCommand)
	eng.AddCommand(engine.MultiplyCommand)

	script := p.Script().(*parser.ScriptContext)
	v := engine.NewCrayonVisitor(eng, engine.NewScope(nil))
	tags := v.VisitScript(script)

	defer func() {
		for range ch {

		}
	}()
	v.RunAllTags(tags)
}
