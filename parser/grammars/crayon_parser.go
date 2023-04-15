// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Crayon

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CrayonParser struct {
	*antlr.BaseParser
}

var crayonParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func crayonParserInit() {
	staticData := &crayonParserStaticData
	staticData.literalNames = []string{
		"", "'$'", "'='", "'{'", "','", "'}'", "'['", "']'", "'@('", "')'",
		"'('", "' '", "';'", "'.'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "NUMBER", "MAIN",
		"FRAME", "LOOP", "STRING", "NONE", "INFINITY", "YES", "NO", "PI", "IDENTIFIER",
		"WS",
	}
	staticData.ruleNames = []string{
		"variable", "numberLiteral", "pair", "stringLiteral", "none", "bool",
		"pI", "infinity", "object", "array", "commandValue", "value", "keywordPath",
		"valuePath", "path", "command", "scope", "exp", "tag", "script",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 25, 180, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 64, 8, 8, 5, 8,
		66, 8, 8, 10, 8, 12, 8, 69, 9, 8, 3, 8, 71, 8, 8, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 3, 9, 79, 8, 9, 5, 9, 81, 8, 9, 10, 9, 12, 9, 84, 9, 9,
		3, 9, 86, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 3, 11, 108, 8, 11, 1, 12, 1, 12, 1, 13, 1, 13, 3, 13, 114, 8, 13,
		1, 14, 1, 14, 3, 14, 118, 8, 14, 1, 15, 1, 15, 4, 15, 122, 8, 15, 11, 15,
		12, 15, 123, 1, 15, 5, 15, 127, 8, 15, 10, 15, 12, 15, 130, 9, 15, 1, 16,
		1, 16, 5, 16, 134, 8, 16, 10, 16, 12, 16, 137, 9, 16, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 17, 1, 17, 3, 17, 145, 8, 17, 5, 17, 147, 8, 17, 10, 17,
		12, 17, 150, 9, 17, 3, 17, 152, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 4, 18,
		158, 8, 18, 11, 18, 12, 18, 159, 1, 18, 1, 18, 3, 18, 164, 8, 18, 1, 18,
		1, 18, 5, 18, 168, 8, 18, 10, 18, 12, 18, 171, 9, 18, 1, 19, 4, 19, 174,
		8, 19, 11, 19, 12, 19, 175, 1, 19, 1, 19, 1, 19, 0, 0, 20, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 0, 2, 1,
		0, 21, 22, 2, 0, 4, 4, 12, 13, 188, 0, 40, 1, 0, 0, 0, 2, 43, 1, 0, 0,
		0, 4, 45, 1, 0, 0, 0, 6, 49, 1, 0, 0, 0, 8, 51, 1, 0, 0, 0, 10, 53, 1,
		0, 0, 0, 12, 55, 1, 0, 0, 0, 14, 57, 1, 0, 0, 0, 16, 59, 1, 0, 0, 0, 18,
		74, 1, 0, 0, 0, 20, 89, 1, 0, 0, 0, 22, 107, 1, 0, 0, 0, 24, 109, 1, 0,
		0, 0, 26, 113, 1, 0, 0, 0, 28, 117, 1, 0, 0, 0, 30, 119, 1, 0, 0, 0, 32,
		131, 1, 0, 0, 0, 34, 151, 1, 0, 0, 0, 36, 153, 1, 0, 0, 0, 38, 173, 1,
		0, 0, 0, 40, 41, 5, 1, 0, 0, 41, 42, 5, 24, 0, 0, 42, 1, 1, 0, 0, 0, 43,
		44, 5, 14, 0, 0, 44, 3, 1, 0, 0, 0, 45, 46, 5, 24, 0, 0, 46, 47, 5, 2,
		0, 0, 47, 48, 3, 22, 11, 0, 48, 5, 1, 0, 0, 0, 49, 50, 5, 18, 0, 0, 50,
		7, 1, 0, 0, 0, 51, 52, 5, 19, 0, 0, 52, 9, 1, 0, 0, 0, 53, 54, 7, 0, 0,
		0, 54, 11, 1, 0, 0, 0, 55, 56, 5, 23, 0, 0, 56, 13, 1, 0, 0, 0, 57, 58,
		5, 20, 0, 0, 58, 15, 1, 0, 0, 0, 59, 70, 5, 3, 0, 0, 60, 67, 3, 4, 2, 0,
		61, 63, 5, 4, 0, 0, 62, 64, 3, 4, 2, 0, 63, 62, 1, 0, 0, 0, 63, 64, 1,
		0, 0, 0, 64, 66, 1, 0, 0, 0, 65, 61, 1, 0, 0, 0, 66, 69, 1, 0, 0, 0, 67,
		65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67, 1, 0, 0,
		0, 70, 60, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 73,
		5, 5, 0, 0, 73, 17, 1, 0, 0, 0, 74, 85, 5, 6, 0, 0, 75, 82, 3, 22, 11,
		0, 76, 78, 5, 4, 0, 0, 77, 79, 3, 22, 11, 0, 78, 77, 1, 0, 0, 0, 78, 79,
		1, 0, 0, 0, 79, 81, 1, 0, 0, 0, 80, 76, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0,
		82, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82, 1,
		0, 0, 0, 85, 75, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87,
		88, 5, 7, 0, 0, 88, 19, 1, 0, 0, 0, 89, 90, 5, 8, 0, 0, 90, 91, 3, 30,
		15, 0, 91, 92, 5, 9, 0, 0, 92, 21, 1, 0, 0, 0, 93, 108, 3, 12, 6, 0, 94,
		108, 3, 10, 5, 0, 95, 108, 3, 20, 10, 0, 96, 108, 3, 0, 0, 0, 97, 108,
		3, 8, 4, 0, 98, 108, 3, 14, 7, 0, 99, 108, 3, 6, 3, 0, 100, 108, 3, 2,
		1, 0, 101, 108, 3, 16, 8, 0, 102, 108, 3, 18, 9, 0, 103, 104, 5, 10, 0,
		0, 104, 105, 3, 22, 11, 0, 105, 106, 5, 9, 0, 0, 106, 108, 1, 0, 0, 0,
		107, 93, 1, 0, 0, 0, 107, 94, 1, 0, 0, 0, 107, 95, 1, 0, 0, 0, 107, 96,
		1, 0, 0, 0, 107, 97, 1, 0, 0, 0, 107, 98, 1, 0, 0, 0, 107, 99, 1, 0, 0,
		0, 107, 100, 1, 0, 0, 0, 107, 101, 1, 0, 0, 0, 107, 102, 1, 0, 0, 0, 107,
		103, 1, 0, 0, 0, 108, 23, 1, 0, 0, 0, 109, 110, 5, 24, 0, 0, 110, 25, 1,
		0, 0, 0, 111, 114, 3, 22, 11, 0, 112, 114, 3, 32, 16, 0, 113, 111, 1, 0,
		0, 0, 113, 112, 1, 0, 0, 0, 114, 27, 1, 0, 0, 0, 115, 118, 3, 24, 12, 0,
		116, 118, 3, 26, 13, 0, 117, 115, 1, 0, 0, 0, 117, 116, 1, 0, 0, 0, 118,
		29, 1, 0, 0, 0, 119, 128, 3, 28, 14, 0, 120, 122, 5, 11, 0, 0, 121, 120,
		1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0,
		0, 0, 124, 125, 1, 0, 0, 0, 125, 127, 3, 28, 14, 0, 126, 121, 1, 0, 0,
		0, 127, 130, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129,
		31, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 131, 135, 5, 3, 0, 0, 132, 134, 3,
		34, 17, 0, 133, 132, 1, 0, 0, 0, 134, 137, 1, 0, 0, 0, 135, 133, 1, 0,
		0, 0, 135, 136, 1, 0, 0, 0, 136, 138, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0,
		138, 139, 5, 5, 0, 0, 139, 33, 1, 0, 0, 0, 140, 152, 3, 32, 16, 0, 141,
		148, 3, 30, 15, 0, 142, 144, 7, 1, 0, 0, 143, 145, 3, 30, 15, 0, 144, 143,
		1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 147, 1, 0, 0, 0, 146, 142, 1, 0,
		0, 0, 147, 150, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0,
		149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 151, 140, 1, 0, 0, 0, 151,
		141, 1, 0, 0, 0, 152, 35, 1, 0, 0, 0, 153, 163, 5, 6, 0, 0, 154, 164, 5,
		15, 0, 0, 155, 157, 5, 16, 0, 0, 156, 158, 5, 11, 0, 0, 157, 156, 1, 0,
		0, 0, 158, 159, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0,
		160, 161, 1, 0, 0, 0, 161, 164, 5, 14, 0, 0, 162, 164, 5, 17, 0, 0, 163,
		154, 1, 0, 0, 0, 163, 155, 1, 0, 0, 0, 163, 162, 1, 0, 0, 0, 164, 165,
		1, 0, 0, 0, 165, 169, 5, 7, 0, 0, 166, 168, 3, 34, 17, 0, 167, 166, 1,
		0, 0, 0, 168, 171, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 170, 1, 0, 0,
		0, 170, 37, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 172, 174, 3, 36, 18, 0, 173,
		172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 175, 176,
		1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 178, 5, 0, 0, 1, 178, 39, 1, 0,
		0, 0, 19, 63, 67, 70, 78, 82, 85, 107, 113, 117, 123, 128, 135, 144, 148,
		151, 159, 163, 169, 175,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CrayonParserInit initializes any static state used to implement CrayonParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCrayonParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CrayonParserInit() {
	staticData := &crayonParserStaticData
	staticData.once.Do(crayonParserInit)
}

// NewCrayonParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCrayonParser(input antlr.TokenStream) *CrayonParser {
	CrayonParserInit()
	this := new(CrayonParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &crayonParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// CrayonParser tokens.
const (
	CrayonParserEOF        = antlr.TokenEOF
	CrayonParserT__0       = 1
	CrayonParserT__1       = 2
	CrayonParserT__2       = 3
	CrayonParserT__3       = 4
	CrayonParserT__4       = 5
	CrayonParserT__5       = 6
	CrayonParserT__6       = 7
	CrayonParserT__7       = 8
	CrayonParserT__8       = 9
	CrayonParserT__9       = 10
	CrayonParserT__10      = 11
	CrayonParserT__11      = 12
	CrayonParserT__12      = 13
	CrayonParserNUMBER     = 14
	CrayonParserMAIN       = 15
	CrayonParserFRAME      = 16
	CrayonParserLOOP       = 17
	CrayonParserSTRING     = 18
	CrayonParserNONE       = 19
	CrayonParserINFINITY   = 20
	CrayonParserYES        = 21
	CrayonParserNO         = 22
	CrayonParserPI         = 23
	CrayonParserIDENTIFIER = 24
	CrayonParserWS         = 25
)

// CrayonParser rules.
const (
	CrayonParserRULE_variable      = 0
	CrayonParserRULE_numberLiteral = 1
	CrayonParserRULE_pair          = 2
	CrayonParserRULE_stringLiteral = 3
	CrayonParserRULE_none          = 4
	CrayonParserRULE_bool          = 5
	CrayonParserRULE_pI            = 6
	CrayonParserRULE_infinity      = 7
	CrayonParserRULE_object        = 8
	CrayonParserRULE_array         = 9
	CrayonParserRULE_commandValue  = 10
	CrayonParserRULE_value         = 11
	CrayonParserRULE_keywordPath   = 12
	CrayonParserRULE_valuePath     = 13
	CrayonParserRULE_path          = 14
	CrayonParserRULE_command       = 15
	CrayonParserRULE_scope         = 16
	CrayonParserRULE_exp           = 17
	CrayonParserRULE_tag           = 18
	CrayonParserRULE_script        = 19
)

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CrayonParserIDENTIFIER, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CrayonParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(CrayonParserT__0)
	}
	{
		p.SetState(41)
		p.Match(CrayonParserIDENTIFIER)
	}

	return localctx
}

// INumberLiteralContext is an interface to support dynamic dispatch.
type INumberLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberLiteralContext differentiates from other interfaces.
	IsNumberLiteralContext()
}

type NumberLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberLiteralContext() *NumberLiteralContext {
	var p = new(NumberLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_numberLiteral
	return p
}

func (*NumberLiteralContext) IsNumberLiteralContext() {}

func NewNumberLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberLiteralContext {
	var p = new(NumberLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_numberLiteral

	return p
}

func (s *NumberLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberLiteralContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CrayonParserNUMBER, 0)
}

func (s *NumberLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterNumberLiteral(s)
	}
}

func (s *NumberLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitNumberLiteral(s)
	}
}

func (s *NumberLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitNumberLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) NumberLiteral() (localctx INumberLiteralContext) {
	this := p
	_ = this

	localctx = NewNumberLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CrayonParserRULE_numberLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(CrayonParserNUMBER)
	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CrayonParserIDENTIFIER, 0)
}

func (s *PairContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitPair(s)
	}
}

func (s *PairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) Pair() (localctx IPairContext) {
	this := p
	_ = this

	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CrayonParserRULE_pair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(CrayonParserIDENTIFIER)
	}
	{
		p.SetState(46)
		p.Match(CrayonParserT__1)
	}
	{
		p.SetState(47)
		p.Value()
	}

	return localctx
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(CrayonParserSTRING, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) StringLiteral() (localctx IStringLiteralContext) {
	this := p
	_ = this

	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CrayonParserRULE_stringLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(CrayonParserSTRING)
	}

	return localctx
}

// INoneContext is an interface to support dynamic dispatch.
type INoneContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNoneContext differentiates from other interfaces.
	IsNoneContext()
}

type NoneContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoneContext() *NoneContext {
	var p = new(NoneContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_none
	return p
}

func (*NoneContext) IsNoneContext() {}

func NewNoneContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoneContext {
	var p = new(NoneContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_none

	return p
}

func (s *NoneContext) GetParser() antlr.Parser { return s.parser }

func (s *NoneContext) NONE() antlr.TerminalNode {
	return s.GetToken(CrayonParserNONE, 0)
}

func (s *NoneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoneContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterNone(s)
	}
}

func (s *NoneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitNone(s)
	}
}

func (s *NoneContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitNone(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) None() (localctx INoneContext) {
	this := p
	_ = this

	localctx = NewNoneContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CrayonParserRULE_none)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(CrayonParserNONE)
	}

	return localctx
}

// IBoolContext is an interface to support dynamic dispatch.
type IBoolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolContext differentiates from other interfaces.
	IsBoolContext()
}

type BoolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolContext() *BoolContext {
	var p = new(BoolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_bool
	return p
}

func (*BoolContext) IsBoolContext() {}

func NewBoolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolContext {
	var p = new(BoolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_bool

	return p
}

func (s *BoolContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolContext) YES() antlr.TerminalNode {
	return s.GetToken(CrayonParserYES, 0)
}

func (s *BoolContext) NO() antlr.TerminalNode {
	return s.GetToken(CrayonParserNO, 0)
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterBool(s)
	}
}

func (s *BoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitBool(s)
	}
}

func (s *BoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitBool(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) Bool_() (localctx IBoolContext) {
	this := p
	_ = this

	localctx = NewBoolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CrayonParserRULE_bool)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CrayonParserYES || _la == CrayonParserNO) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPIContext is an interface to support dynamic dispatch.
type IPIContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPIContext differentiates from other interfaces.
	IsPIContext()
}

type PIContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPIContext() *PIContext {
	var p = new(PIContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_pI
	return p
}

func (*PIContext) IsPIContext() {}

func NewPIContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PIContext {
	var p = new(PIContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_pI

	return p
}

func (s *PIContext) GetParser() antlr.Parser { return s.parser }

func (s *PIContext) PI() antlr.TerminalNode {
	return s.GetToken(CrayonParserPI, 0)
}

func (s *PIContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PIContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PIContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterPI(s)
	}
}

func (s *PIContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitPI(s)
	}
}

func (s *PIContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitPI(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) PI() (localctx IPIContext) {
	this := p
	_ = this

	localctx = NewPIContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CrayonParserRULE_pI)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Match(CrayonParserPI)
	}

	return localctx
}

// IInfinityContext is an interface to support dynamic dispatch.
type IInfinityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInfinityContext differentiates from other interfaces.
	IsInfinityContext()
}

type InfinityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInfinityContext() *InfinityContext {
	var p = new(InfinityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_infinity
	return p
}

func (*InfinityContext) IsInfinityContext() {}

func NewInfinityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InfinityContext {
	var p = new(InfinityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_infinity

	return p
}

func (s *InfinityContext) GetParser() antlr.Parser { return s.parser }

func (s *InfinityContext) INFINITY() antlr.TerminalNode {
	return s.GetToken(CrayonParserINFINITY, 0)
}

func (s *InfinityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InfinityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InfinityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterInfinity(s)
	}
}

func (s *InfinityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitInfinity(s)
	}
}

func (s *InfinityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitInfinity(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) Infinity() (localctx IInfinityContext) {
	this := p
	_ = this

	localctx = NewInfinityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CrayonParserRULE_infinity)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(CrayonParserINFINITY)
	}

	return localctx
}

// IObjectContext is an interface to support dynamic dispatch.
type IObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_pair returns the _pair rule contexts.
	Get_pair() IPairContext

	// Set_pair sets the _pair rule contexts.
	Set_pair(IPairContext)

	// GetPairs returns the Pairs rule context list.
	GetPairs() []IPairContext

	// SetPairs sets the Pairs rule context list.
	SetPairs([]IPairContext)

	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_pair  IPairContext
	Pairs  []IPairContext
}

func NewEmptyObjectContext() *ObjectContext {
	var p = new(ObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_object
	return p
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_object

	return p
}

func (s *ObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectContext) Get_pair() IPairContext { return s._pair }

func (s *ObjectContext) Set_pair(v IPairContext) { s._pair = v }

func (s *ObjectContext) GetPairs() []IPairContext { return s.Pairs }

func (s *ObjectContext) SetPairs(v []IPairContext) { s.Pairs = v }

func (s *ObjectContext) AllPair() []IPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPairContext); ok {
			len++
		}
	}

	tst := make([]IPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPairContext); ok {
			tst[i] = t.(IPairContext)
			i++
		}
	}

	return tst
}

func (s *ObjectContext) Pair(i int) IPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *ObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterObject(s)
	}
}

func (s *ObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitObject(s)
	}
}

func (s *ObjectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitObject(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) Object() (localctx IObjectContext) {
	this := p
	_ = this

	localctx = NewObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CrayonParserRULE_object)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(CrayonParserT__2)
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CrayonParserIDENTIFIER {
		{
			p.SetState(60)

			var _x = p.Pair()

			localctx.(*ObjectContext)._pair = _x
		}
		localctx.(*ObjectContext).Pairs = append(localctx.(*ObjectContext).Pairs, localctx.(*ObjectContext)._pair)
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CrayonParserT__3 {
			{
				p.SetState(61)
				p.Match(CrayonParserT__3)
			}
			p.SetState(63)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == CrayonParserIDENTIFIER {
				{
					p.SetState(62)

					var _x = p.Pair()

					localctx.(*ObjectContext)._pair = _x
				}
				localctx.(*ObjectContext).Pairs = append(localctx.(*ObjectContext).Pairs, localctx.(*ObjectContext)._pair)

			}

			p.SetState(69)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(72)
		p.Match(CrayonParserT__4)
	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_value returns the _value rule contexts.
	Get_value() IValueContext

	// Set_value sets the _value rule contexts.
	Set_value(IValueContext)

	// GetValues returns the Values rule context list.
	GetValues() []IValueContext

	// SetValues sets the Values rule context list.
	SetValues([]IValueContext)

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_value IValueContext
	Values []IValueContext
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) Get_value() IValueContext { return s._value }

func (s *ArrayContext) Set_value(v IValueContext) { s._value = v }

func (s *ArrayContext) GetValues() []IValueContext { return s.Values }

func (s *ArrayContext) SetValues(v []IValueContext) { s.Values = v }

func (s *ArrayContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ArrayContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitArray(s)
	}
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) Array() (localctx IArrayContext) {
	this := p
	_ = this

	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CrayonParserRULE_array)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(CrayonParserT__5)
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16532810) != 0 {
		{
			p.SetState(75)

			var _x = p.Value()

			localctx.(*ArrayContext)._value = _x
		}
		localctx.(*ArrayContext).Values = append(localctx.(*ArrayContext).Values, localctx.(*ArrayContext)._value)
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CrayonParserT__3 {
			{
				p.SetState(76)
				p.Match(CrayonParserT__3)
			}
			p.SetState(78)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16532810) != 0 {
				{
					p.SetState(77)

					var _x = p.Value()

					localctx.(*ArrayContext)._value = _x
				}
				localctx.(*ArrayContext).Values = append(localctx.(*ArrayContext).Values, localctx.(*ArrayContext)._value)

			}

			p.SetState(84)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(87)
		p.Match(CrayonParserT__6)
	}

	return localctx
}

// ICommandValueContext is an interface to support dynamic dispatch.
type ICommandValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandValueContext differentiates from other interfaces.
	IsCommandValueContext()
}

type CommandValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandValueContext() *CommandValueContext {
	var p = new(CommandValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_commandValue
	return p
}

func (*CommandValueContext) IsCommandValueContext() {}

func NewCommandValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandValueContext {
	var p = new(CommandValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_commandValue

	return p
}

func (s *CommandValueContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandValueContext) Command() ICommandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *CommandValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterCommandValue(s)
	}
}

func (s *CommandValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitCommandValue(s)
	}
}

func (s *CommandValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitCommandValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) CommandValue() (localctx ICommandValueContext) {
	this := p
	_ = this

	localctx = NewCommandValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CrayonParserRULE_commandValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(CrayonParserT__7)
	}
	{
		p.SetState(90)
		p.Command()
	}
	{
		p.SetState(91)
		p.Match(CrayonParserT__8)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) PI() IPIContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPIContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPIContext)
}

func (s *ValueContext) Bool_() IBoolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolContext)
}

func (s *ValueContext) CommandValue() ICommandValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandValueContext)
}

func (s *ValueContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ValueContext) None() INoneContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoneContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoneContext)
}

func (s *ValueContext) Infinity() IInfinityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInfinityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInfinityContext)
}

func (s *ValueContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *ValueContext) NumberLiteral() INumberLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberLiteralContext)
}

func (s *ValueContext) Object() IObjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *ValueContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ValueContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CrayonParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(107)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CrayonParserPI:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.PI()
		}

	case CrayonParserYES, CrayonParserNO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.Bool_()
		}

	case CrayonParserT__7:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(95)
			p.CommandValue()
		}

	case CrayonParserT__0:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(96)
			p.Variable()
		}

	case CrayonParserNONE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(97)
			p.None()
		}

	case CrayonParserINFINITY:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(98)
			p.Infinity()
		}

	case CrayonParserSTRING:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(99)
			p.StringLiteral()
		}

	case CrayonParserNUMBER:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(100)
			p.NumberLiteral()
		}

	case CrayonParserT__2:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(101)
			p.Object()
		}

	case CrayonParserT__5:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(102)
			p.Array()
		}

	case CrayonParserT__9:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(103)
			p.Match(CrayonParserT__9)
		}
		{
			p.SetState(104)
			p.Value()
		}
		{
			p.SetState(105)
			p.Match(CrayonParserT__8)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IKeywordPathContext is an interface to support dynamic dispatch.
type IKeywordPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywordPathContext differentiates from other interfaces.
	IsKeywordPathContext()
}

type KeywordPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordPathContext() *KeywordPathContext {
	var p = new(KeywordPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_keywordPath
	return p
}

func (*KeywordPathContext) IsKeywordPathContext() {}

func NewKeywordPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordPathContext {
	var p = new(KeywordPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_keywordPath

	return p
}

func (s *KeywordPathContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordPathContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CrayonParserIDENTIFIER, 0)
}

func (s *KeywordPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterKeywordPath(s)
	}
}

func (s *KeywordPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitKeywordPath(s)
	}
}

func (s *KeywordPathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitKeywordPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) KeywordPath() (localctx IKeywordPathContext) {
	this := p
	_ = this

	localctx = NewKeywordPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CrayonParserRULE_keywordPath)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(CrayonParserIDENTIFIER)
	}

	return localctx
}

// IValuePathContext is an interface to support dynamic dispatch.
type IValuePathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValuePathContext differentiates from other interfaces.
	IsValuePathContext()
}

type ValuePathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValuePathContext() *ValuePathContext {
	var p = new(ValuePathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_valuePath
	return p
}

func (*ValuePathContext) IsValuePathContext() {}

func NewValuePathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValuePathContext {
	var p = new(ValuePathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_valuePath

	return p
}

func (s *ValuePathContext) GetParser() antlr.Parser { return s.parser }

func (s *ValuePathContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValuePathContext) Scope() IScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *ValuePathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValuePathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValuePathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterValuePath(s)
	}
}

func (s *ValuePathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitValuePath(s)
	}
}

func (s *ValuePathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitValuePath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) ValuePath() (localctx IValuePathContext) {
	this := p
	_ = this

	localctx = NewValuePathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CrayonParserRULE_valuePath)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.Value()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Scope()
		}

	}

	return localctx
}

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_path
	return p
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) KeywordPath() IKeywordPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordPathContext)
}

func (s *PathContext) ValuePath() IValuePathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValuePathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValuePathContext)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterPath(s)
	}
}

func (s *PathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitPath(s)
	}
}

func (s *PathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) Path() (localctx IPathContext) {
	this := p
	_ = this

	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CrayonParserRULE_path)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(117)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CrayonParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)
			p.KeywordPath()
		}

	case CrayonParserT__0, CrayonParserT__2, CrayonParserT__5, CrayonParserT__7, CrayonParserT__9, CrayonParserNUMBER, CrayonParserSTRING, CrayonParserNONE, CrayonParserINFINITY, CrayonParserYES, CrayonParserNO, CrayonParserPI:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.ValuePath()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_path returns the _path rule contexts.
	Get_path() IPathContext

	// Set_path sets the _path rule contexts.
	Set_path(IPathContext)

	// GetPaths returns the Paths rule context list.
	GetPaths() []IPathContext

	// SetPaths sets the Paths rule context list.
	SetPaths([]IPathContext)

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_path  IPathContext
	Paths  []IPathContext
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_command
	return p
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) Get_path() IPathContext { return s._path }

func (s *CommandContext) Set_path(v IPathContext) { s._path = v }

func (s *CommandContext) GetPaths() []IPathContext { return s.Paths }

func (s *CommandContext) SetPaths(v []IPathContext) { s.Paths = v }

func (s *CommandContext) AllPath() []IPathContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPathContext); ok {
			len++
		}
	}

	tst := make([]IPathContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPathContext); ok {
			tst[i] = t.(IPathContext)
			i++
		}
	}

	return tst
}

func (s *CommandContext) Path(i int) IPathContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (s *CommandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitCommand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) Command() (localctx ICommandContext) {
	this := p
	_ = this

	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CrayonParserRULE_command)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)

		var _x = p.Path()

		localctx.(*CommandContext)._path = _x
	}
	localctx.(*CommandContext).Paths = append(localctx.(*CommandContext).Paths, localctx.(*CommandContext)._path)
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CrayonParserT__10 {
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == CrayonParserT__10 {
			{
				p.SetState(120)
				p.Match(CrayonParserT__10)
			}

			p.SetState(123)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(125)

			var _x = p.Path()

			localctx.(*CommandContext)._path = _x
		}
		localctx.(*CommandContext).Paths = append(localctx.(*CommandContext).Paths, localctx.(*CommandContext)._path)

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IScopeContext is an interface to support dynamic dispatch.
type IScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScopeContext differentiates from other interfaces.
	IsScopeContext()
}

type ScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScopeContext() *ScopeContext {
	var p = new(ScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_scope
	return p
}

func (*ScopeContext) IsScopeContext() {}

func NewScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScopeContext {
	var p = new(ScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_scope

	return p
}

func (s *ScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *ScopeContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ScopeContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterScope(s)
	}
}

func (s *ScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitScope(s)
	}
}

func (s *ScopeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitScope(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) Scope() (localctx IScopeContext) {
	this := p
	_ = this

	localctx = NewScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CrayonParserRULE_scope)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(CrayonParserT__2)
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33310026) != 0 {
		{
			p.SetState(132)
			p.Exp()
		}

		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(138)
		p.Match(CrayonParserT__4)
	}

	return localctx
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) Scope() IScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *ExpContext) AllCommand() []ICommandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommandContext); ok {
			len++
		}
	}

	tst := make([]ICommandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommandContext); ok {
			tst[i] = t.(ICommandContext)
			i++
		}
	}

	return tst
}

func (s *ExpContext) Command(i int) ICommandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitExp(s)
	}
}

func (s *ExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) Exp() (localctx IExpContext) {
	this := p
	_ = this

	localctx = NewExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CrayonParserRULE_exp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)
			p.Scope()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(141)
			p.Command()
		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&12304) != 0 {
			{
				p.SetState(142)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&12304) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			p.SetState(144)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(143)
					p.Command()
				}

			}

			p.SetState(150)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// ITagContext is an interface to support dynamic dispatch.
type ITagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagContext differentiates from other interfaces.
	IsTagContext()
}

type TagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagContext() *TagContext {
	var p = new(TagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_tag
	return p
}

func (*TagContext) IsTagContext() {}

func NewTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagContext {
	var p = new(TagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_tag

	return p
}

func (s *TagContext) GetParser() antlr.Parser { return s.parser }

func (s *TagContext) MAIN() antlr.TerminalNode {
	return s.GetToken(CrayonParserMAIN, 0)
}

func (s *TagContext) LOOP() antlr.TerminalNode {
	return s.GetToken(CrayonParserLOOP, 0)
}

func (s *TagContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *TagContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *TagContext) FRAME() antlr.TerminalNode {
	return s.GetToken(CrayonParserFRAME, 0)
}

func (s *TagContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CrayonParserNUMBER, 0)
}

func (s *TagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterTag(s)
	}
}

func (s *TagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitTag(s)
	}
}

func (s *TagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitTag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) Tag() (localctx ITagContext) {
	this := p
	_ = this

	localctx = NewTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CrayonParserRULE_tag)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(CrayonParserT__5)
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CrayonParserMAIN:
		{
			p.SetState(154)
			p.Match(CrayonParserMAIN)
		}

	case CrayonParserFRAME:
		{
			p.SetState(155)
			p.Match(CrayonParserFRAME)
		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == CrayonParserT__10 {
			{
				p.SetState(156)
				p.Match(CrayonParserT__10)
			}

			p.SetState(159)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(161)
			p.Match(CrayonParserNUMBER)
		}

	case CrayonParserLOOP:
		{
			p.SetState(162)
			p.Match(CrayonParserLOOP)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(165)
		p.Match(CrayonParserT__6)
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(166)
				p.Exp()
			}

		}
		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	return localctx
}

// IScriptContext is an interface to support dynamic dispatch.
type IScriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScriptContext differentiates from other interfaces.
	IsScriptContext()
}

type ScriptContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScriptContext() *ScriptContext {
	var p = new(ScriptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_script
	return p
}

func (*ScriptContext) IsScriptContext() {}

func NewScriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptContext {
	var p = new(ScriptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_script

	return p
}

func (s *ScriptContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptContext) EOF() antlr.TerminalNode {
	return s.GetToken(CrayonParserEOF, 0)
}

func (s *ScriptContext) AllTag() []ITagContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITagContext); ok {
			len++
		}
	}

	tst := make([]ITagContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITagContext); ok {
			tst[i] = t.(ITagContext)
			i++
		}
	}

	return tst
}

func (s *ScriptContext) Tag(i int) ITagContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *ScriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterScript(s)
	}
}

func (s *ScriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitScript(s)
	}
}

func (s *ScriptContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitScript(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) Script() (localctx IScriptContext) {
	this := p
	_ = this

	localctx = NewScriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CrayonParserRULE_script)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CrayonParserT__5 {
		{
			p.SetState(172)
			p.Tag()
		}

		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(177)
		p.Match(CrayonParserEOF)
	}

	return localctx
}
