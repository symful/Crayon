// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Crayon

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// CrayonListener is a complete listener for a parse tree produced by CrayonParser.
type CrayonListener interface {
	antlr.ParseTreeListener

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterNumberLiteral is called when entering the numberLiteral production.
	EnterNumberLiteral(c *NumberLiteralContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterNone is called when entering the none production.
	EnterNone(c *NoneContext)

	// EnterBool is called when entering the bool production.
	EnterBool(c *BoolContext)

	// EnterPI is called when entering the pI production.
	EnterPI(c *PIContext)

	// EnterInfinity is called when entering the infinity production.
	EnterInfinity(c *InfinityContext)

	// EnterObject is called when entering the object production.
	EnterObject(c *ObjectContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterCommandValue is called when entering the commandValue production.
	EnterCommandValue(c *CommandValueContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterKeywordPath is called when entering the keywordPath production.
	EnterKeywordPath(c *KeywordPathContext)

	// EnterValuePath is called when entering the valuePath production.
	EnterValuePath(c *ValuePathContext)

	// EnterPath is called when entering the path production.
	EnterPath(c *PathContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterScope is called when entering the scope production.
	EnterScope(c *ScopeContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterTag is called when entering the tag production.
	EnterTag(c *TagContext)

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitNumberLiteral is called when exiting the numberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitNone is called when exiting the none production.
	ExitNone(c *NoneContext)

	// ExitBool is called when exiting the bool production.
	ExitBool(c *BoolContext)

	// ExitPI is called when exiting the pI production.
	ExitPI(c *PIContext)

	// ExitInfinity is called when exiting the infinity production.
	ExitInfinity(c *InfinityContext)

	// ExitObject is called when exiting the object production.
	ExitObject(c *ObjectContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitCommandValue is called when exiting the commandValue production.
	ExitCommandValue(c *CommandValueContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitKeywordPath is called when exiting the keywordPath production.
	ExitKeywordPath(c *KeywordPathContext)

	// ExitValuePath is called when exiting the valuePath production.
	ExitValuePath(c *ValuePathContext)

	// ExitPath is called when exiting the path production.
	ExitPath(c *PathContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitScope is called when exiting the scope production.
	ExitScope(c *ScopeContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitTag is called when exiting the tag production.
	ExitTag(c *TagContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)
}
