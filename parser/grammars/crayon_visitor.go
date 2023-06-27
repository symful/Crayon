// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Crayon

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by CrayonParser.
type CrayonVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CrayonParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by CrayonParser#numberLiteral.
	VisitNumberLiteral(ctx *NumberLiteralContext) interface{}

	// Visit a parse tree produced by CrayonParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by CrayonParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by CrayonParser#none.
	VisitNone(ctx *NoneContext) interface{}

	// Visit a parse tree produced by CrayonParser#bool.
	VisitBool(ctx *BoolContext) interface{}

	// Visit a parse tree produced by CrayonParser#pI.
	VisitPI(ctx *PIContext) interface{}

	// Visit a parse tree produced by CrayonParser#infinity.
	VisitInfinity(ctx *InfinityContext) interface{}

	// Visit a parse tree produced by CrayonParser#object.
	VisitObject(ctx *ObjectContext) interface{}

	// Visit a parse tree produced by CrayonParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by CrayonParser#commandValue.
	VisitCommandValue(ctx *CommandValueContext) interface{}

	// Visit a parse tree produced by CrayonParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by CrayonParser#end.
	VisitEnd(ctx *EndContext) interface{}

	// Visit a parse tree produced by CrayonParser#keywordPath.
	VisitKeywordPath(ctx *KeywordPathContext) interface{}

	// Visit a parse tree produced by CrayonParser#valueTag.
	VisitValueTag(ctx *ValueTagContext) interface{}

	// Visit a parse tree produced by CrayonParser#valuePath.
	VisitValuePath(ctx *ValuePathContext) interface{}

	// Visit a parse tree produced by CrayonParser#path.
	VisitPath(ctx *PathContext) interface{}

	// Visit a parse tree produced by CrayonParser#command.
	VisitCommand(ctx *CommandContext) interface{}

	// Visit a parse tree produced by CrayonParser#scope.
	VisitScope(ctx *ScopeContext) interface{}

	// Visit a parse tree produced by CrayonParser#exp.
	VisitExp(ctx *ExpContext) interface{}

	// Visit a parse tree produced by CrayonParser#tag.
	VisitTag(ctx *TagContext) interface{}

	// Visit a parse tree produced by CrayonParser#mainTag.
	VisitMainTag(ctx *MainTagContext) interface{}

	// Visit a parse tree produced by CrayonParser#frameTag.
	VisitFrameTag(ctx *FrameTagContext) interface{}

	// Visit a parse tree produced by CrayonParser#loopTag.
	VisitLoopTag(ctx *LoopTagContext) interface{}

	// Visit a parse tree produced by CrayonParser#customTag.
	VisitCustomTag(ctx *CustomTagContext) interface{}

	// Visit a parse tree produced by CrayonParser#script.
	VisitScript(ctx *ScriptContext) interface{}
}
