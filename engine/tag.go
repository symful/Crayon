package engine

import (
	"time"

	parser "github.com/NekoMaru76/Crayon/parser/grammars"
)

func NewTagVisitor(v *CrayonVisitor, args []any) *CrayonVisitor {
	newScope := NewScope(v.Scope)
	_ = newScope.DefineVariable("args", args)

	return NewCrayonVisitor(v.Engine, newScope, v.CustomTags)
}

type Tag interface {
	Type() string
	GetExpContext() *parser.ExpContext
	Run(visitor *CrayonVisitor, args []any) any
}

type MainTag struct {
	Tag
	exprCtx *parser.ExpContext
}

func NewMainTag(exprCtx *parser.ExpContext) *MainTag {
	return &MainTag{
		exprCtx: exprCtx,
	}
}

func (tag *MainTag) Run(v *CrayonVisitor, args []any) any {
	v.StartTag()

	visitor := NewTagVisitor(v, args)
	res := visitor.VisitExp(tag.exprCtx)

	v.EndTag()
	return res.Value
}

func (*MainTag) Type() string {
	return "Main"
}

func (t *MainTag) GetExpContexts() *parser.ExpContext {
	return t.exprCtx
}

type FrameTag struct {
	Tag
	Time    float64
	exprCtx *parser.ExpContext
}

func (tag *FrameTag) Run(v *CrayonVisitor, args []any) any {
	v.StartTag()

	now := time.Now().UnixMilli()
	visitor := NewTagVisitor(v, args)

	go func() {
		time.Sleep(time.Duration(int64((float64(now-v.Engine.Start) + tag.Time) * float64(time.Second))))
		_ = visitor.VisitExp(tag.exprCtx)
		v.EndTag()
	}()

	return nil
}

func NewFrameTag(exprCtx *parser.ExpContext, time float64) *FrameTag {
	return &FrameTag{
		Time:    time,
		exprCtx: exprCtx,
	}
}

func (*FrameTag) Type() string {
	return "Frame"
}

func (t *FrameTag) GetExpContext() *parser.ExpContext {
	return t.exprCtx
}

type LoopTag struct {
	Tag
	exprCtx *parser.ExpContext
}

func (tag *LoopTag) Run(v *CrayonVisitor, args []any) any {
	v.StartTag()

	visitor := NewTagVisitor(v, args)

	go func() {
		for {
			_ = visitor.VisitExp(tag.exprCtx)
		}

		//v.EndTag()
	}()

	return nil
}

func NewLoopTag(exprCtx *parser.ExpContext) *LoopTag {
	return &LoopTag{
		exprCtx: exprCtx,
	}
}

func (*LoopTag) Type() string {
	return "Loop"
}

func (t *LoopTag) GetExpContext() *parser.ExpContext {
	return t.exprCtx
}

type CustomTag struct {
	Tag
	exprCtx *parser.ExpContext
	Name    string
}

func (tag *CustomTag) Run(v *CrayonVisitor, args []any) any {
	visitor := NewTagVisitor(v, args)
	res := visitor.VisitExp(tag.exprCtx)

	return res.Value
}

func NewCustomTag(exprCtx *parser.ExpContext, name string) *CustomTag {
	return &CustomTag{
		exprCtx: exprCtx,
		Name:    name,
	}
}

func (*CustomTag) Type() string {
	return "Custom"
}

func (t *CustomTag) GetExpContext() *parser.ExpContext {
	return t.exprCtx
}
