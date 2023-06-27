package engine

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	parser "github.com/NekoMaru76/Crayon/parser/grammars"
)

type CrayonVisitor struct {
	*parser.BaseCrayonVisitor
	*Scope
	*Engine
	CustomTags map[string]*CustomTag
}

type Expression struct {
	IsContinue bool
	Value      any
}

func NewCrayonVisitor(engine *Engine, scope *Scope, customTags map[string]*CustomTag) *CrayonVisitor {
	return &CrayonVisitor{
		Scope:      scope,
		Engine:     engine,
		CustomTags: customTags,
	}
}

func (v *CrayonVisitor) VisitNumberLiteral(ctx *parser.NumberLiteralContext) float64 {
	literal := strings.TrimPrefix(ctx.GetText(), "$")
	val, _ := strconv.ParseFloat(literal, 64)
	return val
}

func (v *CrayonVisitor) VisitStringLiteral(ctx *parser.StringLiteralContext) any {
	text := strings.TrimPrefix(strings.Trim(ctx.GetText(), "\""), "$\"")
	return text
}

func (v *CrayonVisitor) VisitVariable(ctx *parser.VariableContext) Variable {
	identifier := ctx.IDENTIFIER().GetText()

	return Variable{
		Name: identifier,
	}
}

func (v *CrayonVisitor) VisitObject(ctx *parser.ObjectContext) any {
	obj := map[string]any{}
	for _, pairI := range ctx.AllPair() {
		pair := pairI.(*parser.PairContext)
		key := pair.IDENTIFIER().GetText()
		value := v.VisitValue(pair.Value().(*parser.ValueContext))
		obj[key] = value
	}
	return obj
}

func (v *CrayonVisitor) VisitArray(ctx *parser.ArrayContext) any {
	arr := []any{}
	for _, value := range ctx.AllValue() {
		arr = append(arr, v.VisitValue(value.(*parser.ValueContext)))
	}
	return arr
}

func (v *CrayonVisitor) VisitTag(ctx *parser.TagContext) Tag {
	expNode := ctx.Exp().(*parser.ExpContext)
	node := ctx.GetChild(1)

	switch node := node.(type) {
	case *parser.CustomTagContext:
		return v.VisitCustomTag(node, expNode)
	case *parser.MainTagContext:
		return v.VisitMainTag(node, expNode)
	case *parser.FrameTagContext:
		return v.VisitFrameTag(node, expNode)
	case *parser.LoopTagContext:
		return v.VisitLoopTag(node, expNode)
	}

	panic(fmt.Errorf("Invalid tag name %s", ctx.GetText()))
}

func (v *CrayonVisitor) VisitMainTag(ctx *parser.MainTagContext, expr *parser.ExpContext) Tag {
	return NewMainTag(expr)
}

func (v *CrayonVisitor) VisitLoopTag(ctx *parser.LoopTagContext, expr *parser.ExpContext) Tag {
	return NewLoopTag(expr)
}

func (v *CrayonVisitor) VisitFrameTag(ctx *parser.FrameTagContext, expr *parser.ExpContext) Tag {
	f, _ := strconv.ParseFloat(ctx.NUMBER().GetText(), 64)
	return NewFrameTag(expr, f)
}

func (v *CrayonVisitor) VisitCustomTag(ctx *parser.CustomTagContext, expr *parser.ExpContext) Tag {
	name := ctx.IDENTIFIER().GetText()
	_, ok := v.CustomTags[name]

	if ok {
		panic(fmt.Errorf("Custom tag with name %s is already exist", name))
	}

	tag := NewCustomTag(expr, name)
	v.CustomTags[name] = tag

	return tag
}

func (v *CrayonVisitor) VisitCommand(ctx *parser.CommandContext) any {
	return v.Engine.ParseCommand(v, ctx)
}

func (v *CrayonVisitor) VisitEnd(ctx *parser.EndContext) any {
	if ctx.Value() != nil {
		return v.VisitValue(ctx.Value().(*parser.ValueContext))
	}

	return nil
}

func (v *CrayonVisitor) VisitScope(ctx *parser.ScopeContext) any {
	scope := NewScope(v.Scope)
	newVisitor := NewCrayonVisitor(v.Engine, scope, v.CustomTags)
	nodes := ctx.GetChildren()
	len := len(nodes)

	for i, node := range nodes {
		if i == 0 || i-1 >= len {
			break
		}

		switch node := node.(type) {
		case *parser.ExpContext:
			exp := newVisitor.VisitExp(node)
			if !exp.IsContinue {
				return exp.Value
			}
		case *parser.EndContext:
			return newVisitor.VisitEnd(node)
		}

	}

	return nil
}

func (v *CrayonVisitor) VisitPath(ctx *parser.PathContext) any {
	if ctx.KeywordPath() != nil {
		return v.VisitKeywordPath(ctx.KeywordPath().(*parser.KeywordPathContext))
	} else if ctx.ValuePath() != nil {
		return v.VisitValuePath(ctx.ValuePath().(*parser.ValuePathContext))
	} else {
		panic("Unknown path type")
	}
}

func (v *CrayonVisitor) VisitKeywordPath(ctx *parser.KeywordPathContext) any {
	keyword := ctx.IDENTIFIER().GetText()
	return strings.ToUpper(keyword)
}

func (v *CrayonVisitor) VisitValuePath(ctx *parser.ValuePathContext) any {
	if ctx.Scope() != nil {
		return func() any {
			return v.VisitScope(ctx.Scope().(*parser.ScopeContext))
		}
	}

	value := v.VisitValue(ctx.Value().(*parser.ValueContext))
	return value
}

func (v *CrayonVisitor) VisitExp(ctx *parser.ExpContext) Expression {
	if ctx.Command(0) != nil || ctx.End(0) != nil {
		for _, node := range ctx.GetChildren() {
			switch node := node.(type) {
			case *parser.CommandContext:
				v.VisitCommand(node)
			case *parser.EndContext:
				return Expression{
					Value:      v.VisitEnd(node),
					IsContinue: false,
				}
			}

		}

		return Expression{
			Value:      nil,
			IsContinue: true,
		}
	} else if ctx.Scope() != nil {
		return Expression{
			Value:      v.VisitScope(ctx.Scope().(*parser.ScopeContext)),
			IsContinue: true,
		}
	}

	panic("Unknown expression type")

}

func (v *CrayonVisitor) VisitNone(ctx *parser.NoneContext) any {
	return nil
}

func (v *CrayonVisitor) VisitPI(ctx *parser.PIContext) any {
	return math.Pi
}

func (v *CrayonVisitor) VisitInfinity(ctx *parser.InfinityContext) any {
	return math.Inf(0)
}

func (v *CrayonVisitor) VisitCommandValue(ctx *parser.CommandValueContext) any {
	return v.VisitCommand(ctx.Command().(*parser.CommandContext))
}

func (v *CrayonVisitor) VisitBool(ctx *parser.BoolContext) any {
	if ctx.YES() != nil {
		return true
	} else {
		return false
	}
}

func (v *CrayonVisitor) VisitValue(ctx *parser.ValueContext) any {
	if ctx.NumberLiteral() != nil {
		return v.VisitNumberLiteral(ctx.NumberLiteral().(*parser.NumberLiteralContext))
	} else if ctx.StringLiteral() != nil {
		return v.VisitStringLiteral(ctx.StringLiteral().(*parser.StringLiteralContext))
	} else if ctx.Variable() != nil {
		return v.VisitVariable(ctx.Variable().(*parser.VariableContext))
	} else if ctx.Object() != nil {
		return v.VisitObject(ctx.Object().(*parser.ObjectContext))
	} else if ctx.Array() != nil {
		return v.VisitArray(ctx.Array().(*parser.ArrayContext))
	} else if ctx.None() != nil {
		return v.VisitNone(ctx.None().(*parser.NoneContext))
	} else if ctx.Infinity() != nil {
		return v.VisitInfinity(ctx.Infinity().(*parser.InfinityContext))
	} else if ctx.CommandValue() != nil {
		return v.VisitCommandValue(ctx.CommandValue().(*parser.CommandValueContext))
	} else if ctx.Bool_() != nil {
		return v.VisitBool(ctx.Bool_().(*parser.BoolContext))
	} else if ctx.PI() != nil {
		return v.VisitPI(ctx.PI().(*parser.PIContext))
	} else if ctx.ValueTag() != nil {
		return v.VisitValueTag(ctx.ValueTag().(*parser.ValueTagContext))
	} else {
		panic("Unknown value type")
	}
}

func (v *CrayonVisitor) VisitValueTag(ctx *parser.ValueTagContext) any {
	name := ctx.IDENTIFIER().GetText()
	tag, ok := v.CustomTags[name]

	if !ok {
		panic(fmt.Errorf("Invalid custom tag name %s", name))
	}

	return tag
}

func (v *CrayonVisitor) VisitScript(ctx *parser.ScriptContext) []Tag {
	tags := []Tag{}

	for _, tagCtx := range ctx.AllTag() {
		tags = append(tags, v.VisitTag(tagCtx.(*parser.TagContext)))
	}

	return tags
}

func (v *CrayonVisitor) RunTag(tag Tag, args []any) any {
	return tag.Run(v, args)
}

func (v *CrayonVisitor) RunAllAvailableTags(tags []Tag, tagsArgs [][]any) []any {
	v.StartTag()

	vals := []any{}

	for i, tag := range tags {
		if tag.Type() == "Custom" {
			continue
		}

		vals = append(vals, v.RunTag(tag, tagsArgs[i]))
	}

	go v.EndTag()

	return vals
}
