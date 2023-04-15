package engine

import (
	"math"
	"strconv"
	"strings"
	"time"

	parser "github.com/NekoMaru76/Crayon/parser/grammars"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type CrayonVisitor struct {
	*parser.BaseCrayonVisitor
	*Scope
	*Engine
}

func NewCrayonVisitor(engine *Engine, scope *Scope) *CrayonVisitor {
	return &CrayonVisitor{
		Scope:  scope,
		Engine: engine,
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
	for _, pairI := range ctx.Pairs {
		pair := pairI.(*parser.PairContext)
		key := pair.IDENTIFIER().GetText()
		value := v.VisitValue(pair.Value().(*parser.ValueContext))
		obj[key] = value
	}
	return obj
}

func (v *CrayonVisitor) VisitArray(ctx *parser.ArrayContext) any {
	arr := []any{}
	for _, value := range ctx.Values {
		arr = append(arr, v.VisitValue(value.(*parser.ValueContext)))
	}
	return arr
}

type Tag interface {
	Type() string
	GetExpContexts() []*parser.ExpContext
	Run(*CrayonVisitor) []any
}

type MainTag struct {
	Tag
	exprCtx []*parser.ExpContext
}

func NewMainTag(exprCtx []*parser.ExpContext) MainTag {
	return MainTag{
		exprCtx: exprCtx,
	}
}

func (tag MainTag) Run(v *CrayonVisitor) []any {
	v.StartTag()
	var res []any
	for _, ctx := range tag.GetExpContexts() {
		res = append(res, v.VisitExp(ctx))
	}
	v.EndTag()
	return res
}

func (MainTag) Type() string {
	return "Main"
}

func (t MainTag) GetExpContexts() []*parser.ExpContext {
	return t.exprCtx
}

type FrameTag struct {
	Tag
	Time    float64
	exprCtx []*parser.ExpContext
}

func (tag FrameTag) Run(v *CrayonVisitor) []any {
	v.StartTag()

	now := time.Now().UnixMilli()

	go func() {
		time.Sleep(time.Duration(int64((float64(now-v.Engine.Start) + tag.Time) * float64(time.Second))))
		for _, ctx := range tag.GetExpContexts() {
			v.VisitExp(ctx)
		}
		v.EndTag()
	}()

	return nil
}

func NewFrameTag(time float64, exprCtx []*parser.ExpContext) FrameTag {
	return FrameTag{
		Time:    time,
		exprCtx: exprCtx,
	}
}

func (FrameTag) Type() string {
	return "Frame"
}

func (t FrameTag) GetExpContexts() []*parser.ExpContext {
	return t.exprCtx
}

type LoopTag struct {
	Tag
	exprCtx []*parser.ExpContext
}

func (tag LoopTag) Run(v *CrayonVisitor) []any {
	v.StartTag()

	go func() {
		for {
			for _, ctx := range tag.GetExpContexts() {
				v.VisitExp(ctx)
			}
		}

		//v.EndTag()
	}()

	return nil
}

func NewLoopTag(exprCtx []*parser.ExpContext) LoopTag {
	return LoopTag{
		exprCtx: exprCtx,
	}
}

func (LoopTag) Type() string {
	return "Loop"
}

func (t LoopTag) GetExpContexts() []*parser.ExpContext {
	return t.exprCtx
}

func (v *CrayonVisitor) VisitTag(ctx *parser.TagContext) Tag {
	expNodes := ctx.AllExp()
	exps := make([]*parser.ExpContext, len(expNodes))
	for i, node := range expNodes {
		exps[i] = node.(*parser.ExpContext)
	}

	tagContent := strings.ToUpper(ctx.GetChild(1).(antlr.TerminalNode).GetText())

	if tagContent == "MAIN" {
		return NewMainTag(exps)
	} else if tagContent == "FRAME" {
		time, _ := strconv.ParseFloat(ctx.NUMBER().GetText(), 64)
		return NewFrameTag(time, exps)
	} else if tagContent == "LOOP" {
		return NewLoopTag(exps)
	} else {
		panic("Unknown tag type")
	}
}

func (v *CrayonVisitor) VisitCommand(ctx *parser.CommandContext) any {
	return v.Engine.ParseCommand(v, ctx)
}

func (v *CrayonVisitor) VisitScope(ctx *parser.ScopeContext) any {
	scope := NewScope(v.Scope)
	newVisitor := NewCrayonVisitor(v.Engine, scope)

	for _, exp := range ctx.AllExp() {
		newVisitor.VisitExp(exp.(*parser.ExpContext))
	}

	return scope
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

func (v *CrayonVisitor) VisitExp(ctx *parser.ExpContext) any {
	var val any

	if ctx.Command(0) != nil {
		vals := []any{}

		for _, cmd := range ctx.AllCommand() {
			vals = append(vals, v.VisitCommand(cmd.(*parser.CommandContext)))
		}

		val = vals
	} else if ctx.Scope() != nil {
		val = v.VisitScope(ctx.Scope().(*parser.ScopeContext))
	} else {
		panic("Unknown expression type")
	}

	return val
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
	} else {
		panic("Unknown value type")
	}
}

func (v *CrayonVisitor) VisitScript(ctx *parser.ScriptContext) []Tag {
	tags := []Tag{}

	for _, tagCtx := range ctx.AllTag() {
		tags = append(tags, v.VisitTag(tagCtx.(*parser.TagContext)))
	}

	return tags
}

func (v *CrayonVisitor) RunTag(tag Tag) any {
	return tag.Run(v)
}

func (v *CrayonVisitor) RunAllTags(tags []Tag) []any {
	v.StartTag()

	vals := []any{}

	for _, tag := range tags {
		vals = append(vals, v.RunTag(tag))
	}

	go v.EndTag()

	return vals
}
