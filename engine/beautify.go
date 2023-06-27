package engine

import (
	"strings"

	parser "github.com/NekoMaru76/Crayon/parser/grammars"
)

func BeautifyCommandContext(ctx *parser.CommandContext) string {
	ss := []string{}

	for _, node := range ctx.AllPath() {
		ss = append(ss, BeautifyPathContext(node.(*parser.PathContext)))
	}

	return strings.Join(ss, " ")
}

func BeautifyPathContext(ctx *parser.PathContext) string {
	if ctx.KeywordPath() != nil {
		return ctx.KeywordPath().GetText()
	} else {
		return BeautifyValuePathContext(ctx.ValuePath().(*parser.ValuePathContext))
	}
}

func BeautifyValuePathContext(ctx *parser.ValuePathContext) string {
	if ctx.Scope() != nil {
		return BeautifyScopeContext(ctx.Scope().(*parser.ScopeContext))
	}

	return BeautifyValueContext(ctx.Value().(*parser.ValueContext))
}

func BeautifyScopeContext(ctx *parser.ScopeContext) string {
	ss := []string{}

	for _, exp := range ctx.AllExp() {
		ss = append(ss, BeautifyExpContext(exp.(*parser.ExpContext)))
	}

	return "{\n" + strings.Join(ss, "\n\t") + "\n}"
}

func BeautifyValueContext(ctx *parser.ValueContext) string {
	return ctx.GetText()
}

func BeautifyExpContext(ctx *parser.ExpContext) string {
	if ctx.Scope() != nil {
		return BeautifyScopeContext(ctx.Scope().(*parser.ScopeContext))
	}

	ss := []string{}

	for _, cmd := range ctx.AllCommand() {
		ss = append(ss, BeautifyCommandContext(cmd.(*parser.CommandContext)))
	}

	return strings.Join(ss, ";")
}
