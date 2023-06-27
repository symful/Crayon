// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Crayon

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseCrayonListener is a complete listener for a parse tree produced by CrayonParser.
type BaseCrayonListener struct{}

var _ CrayonListener = &BaseCrayonListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCrayonListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCrayonListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCrayonListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCrayonListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseCrayonListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseCrayonListener) ExitVariable(ctx *VariableContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseCrayonListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseCrayonListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseCrayonListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseCrayonListener) ExitPair(ctx *PairContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseCrayonListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseCrayonListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterNone is called when production none is entered.
func (s *BaseCrayonListener) EnterNone(ctx *NoneContext) {}

// ExitNone is called when production none is exited.
func (s *BaseCrayonListener) ExitNone(ctx *NoneContext) {}

// EnterBool is called when production bool is entered.
func (s *BaseCrayonListener) EnterBool(ctx *BoolContext) {}

// ExitBool is called when production bool is exited.
func (s *BaseCrayonListener) ExitBool(ctx *BoolContext) {}

// EnterPI is called when production pI is entered.
func (s *BaseCrayonListener) EnterPI(ctx *PIContext) {}

// ExitPI is called when production pI is exited.
func (s *BaseCrayonListener) ExitPI(ctx *PIContext) {}

// EnterInfinity is called when production infinity is entered.
func (s *BaseCrayonListener) EnterInfinity(ctx *InfinityContext) {}

// ExitInfinity is called when production infinity is exited.
func (s *BaseCrayonListener) ExitInfinity(ctx *InfinityContext) {}

// EnterObject is called when production object is entered.
func (s *BaseCrayonListener) EnterObject(ctx *ObjectContext) {}

// ExitObject is called when production object is exited.
func (s *BaseCrayonListener) ExitObject(ctx *ObjectContext) {}

// EnterArray is called when production array is entered.
func (s *BaseCrayonListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseCrayonListener) ExitArray(ctx *ArrayContext) {}

// EnterCommandValue is called when production commandValue is entered.
func (s *BaseCrayonListener) EnterCommandValue(ctx *CommandValueContext) {}

// ExitCommandValue is called when production commandValue is exited.
func (s *BaseCrayonListener) ExitCommandValue(ctx *CommandValueContext) {}

// EnterValue is called when production value is entered.
func (s *BaseCrayonListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseCrayonListener) ExitValue(ctx *ValueContext) {}

// EnterEnd is called when production end is entered.
func (s *BaseCrayonListener) EnterEnd(ctx *EndContext) {}

// ExitEnd is called when production end is exited.
func (s *BaseCrayonListener) ExitEnd(ctx *EndContext) {}

// EnterKeywordPath is called when production keywordPath is entered.
func (s *BaseCrayonListener) EnterKeywordPath(ctx *KeywordPathContext) {}

// ExitKeywordPath is called when production keywordPath is exited.
func (s *BaseCrayonListener) ExitKeywordPath(ctx *KeywordPathContext) {}

// EnterValueTag is called when production valueTag is entered.
func (s *BaseCrayonListener) EnterValueTag(ctx *ValueTagContext) {}

// ExitValueTag is called when production valueTag is exited.
func (s *BaseCrayonListener) ExitValueTag(ctx *ValueTagContext) {}

// EnterValuePath is called when production valuePath is entered.
func (s *BaseCrayonListener) EnterValuePath(ctx *ValuePathContext) {}

// ExitValuePath is called when production valuePath is exited.
func (s *BaseCrayonListener) ExitValuePath(ctx *ValuePathContext) {}

// EnterPath is called when production path is entered.
func (s *BaseCrayonListener) EnterPath(ctx *PathContext) {}

// ExitPath is called when production path is exited.
func (s *BaseCrayonListener) ExitPath(ctx *PathContext) {}

// EnterCommand is called when production command is entered.
func (s *BaseCrayonListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseCrayonListener) ExitCommand(ctx *CommandContext) {}

// EnterScope is called when production scope is entered.
func (s *BaseCrayonListener) EnterScope(ctx *ScopeContext) {}

// ExitScope is called when production scope is exited.
func (s *BaseCrayonListener) ExitScope(ctx *ScopeContext) {}

// EnterExp is called when production exp is entered.
func (s *BaseCrayonListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseCrayonListener) ExitExp(ctx *ExpContext) {}

// EnterTag is called when production tag is entered.
func (s *BaseCrayonListener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BaseCrayonListener) ExitTag(ctx *TagContext) {}

// EnterMainTag is called when production mainTag is entered.
func (s *BaseCrayonListener) EnterMainTag(ctx *MainTagContext) {}

// ExitMainTag is called when production mainTag is exited.
func (s *BaseCrayonListener) ExitMainTag(ctx *MainTagContext) {}

// EnterFrameTag is called when production frameTag is entered.
func (s *BaseCrayonListener) EnterFrameTag(ctx *FrameTagContext) {}

// ExitFrameTag is called when production frameTag is exited.
func (s *BaseCrayonListener) ExitFrameTag(ctx *FrameTagContext) {}

// EnterLoopTag is called when production loopTag is entered.
func (s *BaseCrayonListener) EnterLoopTag(ctx *LoopTagContext) {}

// ExitLoopTag is called when production loopTag is exited.
func (s *BaseCrayonListener) ExitLoopTag(ctx *LoopTagContext) {}

// EnterCustomTag is called when production customTag is entered.
func (s *BaseCrayonListener) EnterCustomTag(ctx *CustomTagContext) {}

// ExitCustomTag is called when production customTag is exited.
func (s *BaseCrayonListener) ExitCustomTag(ctx *CustomTagContext) {}

// EnterScript is called when production script is entered.
func (s *BaseCrayonListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseCrayonListener) ExitScript(ctx *ScriptContext) {}
