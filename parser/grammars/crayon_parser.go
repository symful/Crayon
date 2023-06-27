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
		"", "'$'", "'='", "'#{'", "','", "'}'", "'#['", "']'", "'@('", "')'",
		"'('", "'@'", "'{'", "';'", "'['", "", "", "", "", "", "", "", "", "",
		"", "'~'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "NUMBER",
		"MAIN", "FRAME", "LOOP", "STRING", "NONE", "INFINITY", "YES", "NO",
		"PI", "END", "WS", "IDENTIFIER",
	}
	staticData.ruleNames = []string{
		"variable", "numberLiteral", "pair", "stringLiteral", "none", "bool",
		"pI", "infinity", "object", "array", "commandValue", "value", "end",
		"keywordPath", "valueTag", "valuePath", "path", "command", "scope",
		"exp", "tag", "mainTag", "frameTag", "loopTag", "customTag", "script",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 27, 198, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 1, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 76, 8,
		8, 5, 8, 78, 8, 8, 10, 8, 12, 8, 81, 9, 8, 3, 8, 83, 8, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 91, 8, 9, 5, 9, 93, 8, 9, 10, 9, 12, 9, 96,
		9, 9, 3, 9, 98, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 3, 11, 121, 8, 11, 1, 12, 1, 12, 3, 12, 125, 8, 12,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 3, 15, 134, 8, 15, 1,
		16, 1, 16, 3, 16, 138, 8, 16, 1, 17, 4, 17, 141, 8, 17, 11, 17, 12, 17,
		142, 1, 18, 1, 18, 1, 18, 5, 18, 148, 8, 18, 10, 18, 12, 18, 151, 9, 18,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 3, 19, 158, 8, 19, 1, 19, 4, 19, 161,
		8, 19, 11, 19, 12, 19, 162, 4, 19, 165, 8, 19, 11, 19, 12, 19, 166, 3,
		19, 169, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 176, 8, 20, 1,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 24, 1, 25, 4, 25, 192, 8, 25, 11, 25, 12, 25, 193, 1, 25, 1,
		25, 1, 25, 0, 0, 26, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 0, 1, 1, 0, 22, 23, 202,
		0, 52, 1, 0, 0, 0, 2, 55, 1, 0, 0, 0, 4, 57, 1, 0, 0, 0, 6, 61, 1, 0, 0,
		0, 8, 63, 1, 0, 0, 0, 10, 65, 1, 0, 0, 0, 12, 67, 1, 0, 0, 0, 14, 69, 1,
		0, 0, 0, 16, 71, 1, 0, 0, 0, 18, 86, 1, 0, 0, 0, 20, 101, 1, 0, 0, 0, 22,
		120, 1, 0, 0, 0, 24, 122, 1, 0, 0, 0, 26, 126, 1, 0, 0, 0, 28, 128, 1,
		0, 0, 0, 30, 133, 1, 0, 0, 0, 32, 137, 1, 0, 0, 0, 34, 140, 1, 0, 0, 0,
		36, 144, 1, 0, 0, 0, 38, 168, 1, 0, 0, 0, 40, 170, 1, 0, 0, 0, 42, 180,
		1, 0, 0, 0, 44, 182, 1, 0, 0, 0, 46, 185, 1, 0, 0, 0, 48, 187, 1, 0, 0,
		0, 50, 191, 1, 0, 0, 0, 52, 53, 5, 1, 0, 0, 53, 54, 5, 27, 0, 0, 54, 1,
		1, 0, 0, 0, 55, 56, 5, 15, 0, 0, 56, 3, 1, 0, 0, 0, 57, 58, 5, 27, 0, 0,
		58, 59, 5, 2, 0, 0, 59, 60, 3, 22, 11, 0, 60, 5, 1, 0, 0, 0, 61, 62, 5,
		19, 0, 0, 62, 7, 1, 0, 0, 0, 63, 64, 5, 20, 0, 0, 64, 9, 1, 0, 0, 0, 65,
		66, 7, 0, 0, 0, 66, 11, 1, 0, 0, 0, 67, 68, 5, 24, 0, 0, 68, 13, 1, 0,
		0, 0, 69, 70, 5, 21, 0, 0, 70, 15, 1, 0, 0, 0, 71, 82, 5, 3, 0, 0, 72,
		79, 3, 4, 2, 0, 73, 75, 5, 4, 0, 0, 74, 76, 3, 4, 2, 0, 75, 74, 1, 0, 0,
		0, 75, 76, 1, 0, 0, 0, 76, 78, 1, 0, 0, 0, 77, 73, 1, 0, 0, 0, 78, 81,
		1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 83, 1, 0, 0, 0,
		81, 79, 1, 0, 0, 0, 82, 72, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 1,
		0, 0, 0, 84, 85, 5, 5, 0, 0, 85, 17, 1, 0, 0, 0, 86, 97, 5, 6, 0, 0, 87,
		94, 3, 22, 11, 0, 88, 90, 5, 4, 0, 0, 89, 91, 3, 22, 11, 0, 90, 89, 1,
		0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 93, 1, 0, 0, 0, 92, 88, 1, 0, 0, 0, 93,
		96, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 98, 1, 0, 0,
		0, 96, 94, 1, 0, 0, 0, 97, 87, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99,
		1, 0, 0, 0, 99, 100, 5, 7, 0, 0, 100, 19, 1, 0, 0, 0, 101, 102, 5, 8, 0,
		0, 102, 103, 3, 34, 17, 0, 103, 104, 5, 9, 0, 0, 104, 21, 1, 0, 0, 0, 105,
		121, 3, 12, 6, 0, 106, 121, 3, 10, 5, 0, 107, 121, 3, 28, 14, 0, 108, 121,
		3, 20, 10, 0, 109, 121, 3, 0, 0, 0, 110, 121, 3, 8, 4, 0, 111, 121, 3,
		14, 7, 0, 112, 121, 3, 6, 3, 0, 113, 121, 3, 2, 1, 0, 114, 121, 3, 16,
		8, 0, 115, 121, 3, 18, 9, 0, 116, 117, 5, 10, 0, 0, 117, 118, 3, 22, 11,
		0, 118, 119, 5, 9, 0, 0, 119, 121, 1, 0, 0, 0, 120, 105, 1, 0, 0, 0, 120,
		106, 1, 0, 0, 0, 120, 107, 1, 0, 0, 0, 120, 108, 1, 0, 0, 0, 120, 109,
		1, 0, 0, 0, 120, 110, 1, 0, 0, 0, 120, 111, 1, 0, 0, 0, 120, 112, 1, 0,
		0, 0, 120, 113, 1, 0, 0, 0, 120, 114, 1, 0, 0, 0, 120, 115, 1, 0, 0, 0,
		120, 116, 1, 0, 0, 0, 121, 23, 1, 0, 0, 0, 122, 124, 5, 25, 0, 0, 123,
		125, 3, 22, 11, 0, 124, 123, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 25,
		1, 0, 0, 0, 126, 127, 5, 27, 0, 0, 127, 27, 1, 0, 0, 0, 128, 129, 5, 11,
		0, 0, 129, 130, 5, 27, 0, 0, 130, 29, 1, 0, 0, 0, 131, 134, 3, 22, 11,
		0, 132, 134, 3, 36, 18, 0, 133, 131, 1, 0, 0, 0, 133, 132, 1, 0, 0, 0,
		134, 31, 1, 0, 0, 0, 135, 138, 3, 26, 13, 0, 136, 138, 3, 30, 15, 0, 137,
		135, 1, 0, 0, 0, 137, 136, 1, 0, 0, 0, 138, 33, 1, 0, 0, 0, 139, 141, 3,
		32, 16, 0, 140, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 140, 1, 0,
		0, 0, 142, 143, 1, 0, 0, 0, 143, 35, 1, 0, 0, 0, 144, 149, 5, 12, 0, 0,
		145, 148, 3, 38, 19, 0, 146, 148, 3, 24, 12, 0, 147, 145, 1, 0, 0, 0, 147,
		146, 1, 0, 0, 0, 148, 151, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150,
		1, 0, 0, 0, 150, 152, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 152, 153, 5, 5,
		0, 0, 153, 37, 1, 0, 0, 0, 154, 169, 3, 36, 18, 0, 155, 158, 3, 34, 17,
		0, 156, 158, 3, 24, 12, 0, 157, 155, 1, 0, 0, 0, 157, 156, 1, 0, 0, 0,
		158, 160, 1, 0, 0, 0, 159, 161, 5, 13, 0, 0, 160, 159, 1, 0, 0, 0, 161,
		162, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 165,
		1, 0, 0, 0, 164, 157, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 164, 1, 0,
		0, 0, 166, 167, 1, 0, 0, 0, 167, 169, 1, 0, 0, 0, 168, 154, 1, 0, 0, 0,
		168, 164, 1, 0, 0, 0, 169, 39, 1, 0, 0, 0, 170, 175, 5, 14, 0, 0, 171,
		176, 3, 42, 21, 0, 172, 176, 3, 44, 22, 0, 173, 176, 3, 46, 23, 0, 174,
		176, 3, 48, 24, 0, 175, 171, 1, 0, 0, 0, 175, 172, 1, 0, 0, 0, 175, 173,
		1, 0, 0, 0, 175, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 178, 5, 7,
		0, 0, 178, 179, 3, 38, 19, 0, 179, 41, 1, 0, 0, 0, 180, 181, 5, 16, 0,
		0, 181, 43, 1, 0, 0, 0, 182, 183, 5, 17, 0, 0, 183, 184, 5, 15, 0, 0, 184,
		45, 1, 0, 0, 0, 185, 186, 5, 18, 0, 0, 186, 47, 1, 0, 0, 0, 187, 188, 5,
		11, 0, 0, 188, 189, 5, 27, 0, 0, 189, 49, 1, 0, 0, 0, 190, 192, 3, 40,
		20, 0, 191, 190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0,
		193, 194, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 196, 5, 0, 0, 1, 196,
		51, 1, 0, 0, 0, 19, 75, 79, 82, 90, 94, 97, 120, 124, 133, 137, 142, 147,
		149, 157, 162, 166, 168, 175, 193,
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
	CrayonParserT__13      = 14
	CrayonParserNUMBER     = 15
	CrayonParserMAIN       = 16
	CrayonParserFRAME      = 17
	CrayonParserLOOP       = 18
	CrayonParserSTRING     = 19
	CrayonParserNONE       = 20
	CrayonParserINFINITY   = 21
	CrayonParserYES        = 22
	CrayonParserNO         = 23
	CrayonParserPI         = 24
	CrayonParserEND        = 25
	CrayonParserWS         = 26
	CrayonParserIDENTIFIER = 27
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
	CrayonParserRULE_end           = 12
	CrayonParserRULE_keywordPath   = 13
	CrayonParserRULE_valueTag      = 14
	CrayonParserRULE_valuePath     = 15
	CrayonParserRULE_path          = 16
	CrayonParserRULE_command       = 17
	CrayonParserRULE_scope         = 18
	CrayonParserRULE_exp           = 19
	CrayonParserRULE_tag           = 20
	CrayonParserRULE_mainTag       = 21
	CrayonParserRULE_frameTag      = 22
	CrayonParserRULE_loopTag       = 23
	CrayonParserRULE_customTag     = 24
	CrayonParserRULE_script        = 25
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
		p.SetState(52)
		p.Match(CrayonParserT__0)
	}
	{
		p.SetState(53)
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
		p.SetState(55)
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
		p.SetState(57)
		p.Match(CrayonParserIDENTIFIER)
	}
	{
		p.SetState(58)
		p.Match(CrayonParserT__1)
	}
	{
		p.SetState(59)
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
		p.SetState(61)
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
		p.SetState(63)
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
		p.SetState(65)
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
		p.SetState(67)
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
		p.SetState(69)
		p.Match(CrayonParserINFINITY)
	}

	return localctx
}

// IObjectContext is an interface to support dynamic dispatch.
type IObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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
		p.SetState(71)
		p.Match(CrayonParserT__2)
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CrayonParserIDENTIFIER {
		{
			p.SetState(72)
			p.Pair()
		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CrayonParserT__3 {
			{
				p.SetState(73)
				p.Match(CrayonParserT__3)
			}
			p.SetState(75)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == CrayonParserIDENTIFIER {
				{
					p.SetState(74)
					p.Pair()
				}

			}

			p.SetState(81)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(84)
		p.Match(CrayonParserT__4)
	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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
		p.SetState(86)
		p.Match(CrayonParserT__5)
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33066314) != 0 {
		{
			p.SetState(87)
			p.Value()
		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CrayonParserT__3 {
			{
				p.SetState(88)
				p.Match(CrayonParserT__3)
			}
			p.SetState(90)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33066314) != 0 {
				{
					p.SetState(89)
					p.Value()
				}

			}

			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(99)
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
		p.SetState(101)
		p.Match(CrayonParserT__7)
	}
	{
		p.SetState(102)
		p.Command()
	}
	{
		p.SetState(103)
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

func (s *ValueContext) ValueTag() IValueTagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueTagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueTagContext)
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

	p.SetState(120)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CrayonParserPI:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.PI()
		}

	case CrayonParserYES, CrayonParserNO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
			p.Bool_()
		}

	case CrayonParserT__10:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(107)
			p.ValueTag()
		}

	case CrayonParserT__7:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(108)
			p.CommandValue()
		}

	case CrayonParserT__0:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(109)
			p.Variable()
		}

	case CrayonParserNONE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(110)
			p.None()
		}

	case CrayonParserINFINITY:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(111)
			p.Infinity()
		}

	case CrayonParserSTRING:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(112)
			p.StringLiteral()
		}

	case CrayonParserNUMBER:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(113)
			p.NumberLiteral()
		}

	case CrayonParserT__2:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(114)
			p.Object()
		}

	case CrayonParserT__5:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(115)
			p.Array()
		}

	case CrayonParserT__9:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(116)
			p.Match(CrayonParserT__9)
		}
		{
			p.SetState(117)
			p.Value()
		}
		{
			p.SetState(118)
			p.Match(CrayonParserT__8)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEndContext is an interface to support dynamic dispatch.
type IEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEndContext differentiates from other interfaces.
	IsEndContext()
}

type EndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndContext() *EndContext {
	var p = new(EndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_end
	return p
}

func (*EndContext) IsEndContext() {}

func NewEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndContext {
	var p = new(EndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_end

	return p
}

func (s *EndContext) GetParser() antlr.Parser { return s.parser }

func (s *EndContext) END() antlr.TerminalNode {
	return s.GetToken(CrayonParserEND, 0)
}

func (s *EndContext) Value() IValueContext {
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

func (s *EndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterEnd(s)
	}
}

func (s *EndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitEnd(s)
	}
}

func (s *EndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitEnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) End() (localctx IEndContext) {
	this := p
	_ = this

	localctx = NewEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CrayonParserRULE_end)

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
		p.SetState(122)
		p.Match(CrayonParserEND)
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(123)
			p.Value()
		}

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
	p.EnterRule(localctx, 26, CrayonParserRULE_keywordPath)

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
		p.SetState(126)
		p.Match(CrayonParserIDENTIFIER)
	}

	return localctx
}

// IValueTagContext is an interface to support dynamic dispatch.
type IValueTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueTagContext differentiates from other interfaces.
	IsValueTagContext()
}

type ValueTagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueTagContext() *ValueTagContext {
	var p = new(ValueTagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_valueTag
	return p
}

func (*ValueTagContext) IsValueTagContext() {}

func NewValueTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueTagContext {
	var p = new(ValueTagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_valueTag

	return p
}

func (s *ValueTagContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueTagContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CrayonParserIDENTIFIER, 0)
}

func (s *ValueTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterValueTag(s)
	}
}

func (s *ValueTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitValueTag(s)
	}
}

func (s *ValueTagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitValueTag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) ValueTag() (localctx IValueTagContext) {
	this := p
	_ = this

	localctx = NewValueTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CrayonParserRULE_valueTag)

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
		p.SetState(128)
		p.Match(CrayonParserT__10)
	}
	{
		p.SetState(129)
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
	p.EnterRule(localctx, 30, CrayonParserRULE_valuePath)

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

	p.SetState(133)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CrayonParserT__0, CrayonParserT__2, CrayonParserT__5, CrayonParserT__7, CrayonParserT__9, CrayonParserT__10, CrayonParserNUMBER, CrayonParserSTRING, CrayonParserNONE, CrayonParserINFINITY, CrayonParserYES, CrayonParserNO, CrayonParserPI:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.Value()
		}

	case CrayonParserT__11:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
			p.Scope()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 32, CrayonParserRULE_path)

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

	p.SetState(137)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CrayonParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.KeywordPath()
		}

	case CrayonParserT__0, CrayonParserT__2, CrayonParserT__5, CrayonParserT__7, CrayonParserT__9, CrayonParserT__10, CrayonParserT__11, CrayonParserNUMBER, CrayonParserSTRING, CrayonParserNONE, CrayonParserINFINITY, CrayonParserYES, CrayonParserNO, CrayonParserPI:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(136)
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

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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
	p.EnterRule(localctx, 34, CrayonParserRULE_command)
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
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&167288138) != 0 {
		{
			p.SetState(139)
			p.Path()
		}

		p.SetState(142)
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

func (s *ScopeContext) AllEnd() []IEndContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEndContext); ok {
			len++
		}
	}

	tst := make([]IEndContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEndContext); ok {
			tst[i] = t.(IEndContext)
			i++
		}
	}

	return tst
}

func (s *ScopeContext) End(i int) IEndContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndContext); ok {
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

	return t.(IEndContext)
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
	p.EnterRule(localctx, 36, CrayonParserRULE_scope)
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
		p.SetState(144)
		p.Match(CrayonParserT__11)
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&200842570) != 0 {
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(145)
				p.Exp()
			}

		case 2:
			{
				p.SetState(146)
				p.End()
			}

		}

		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(152)
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

func (s *ExpContext) AllEnd() []IEndContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEndContext); ok {
			len++
		}
	}

	tst := make([]IEndContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEndContext); ok {
			tst[i] = t.(IEndContext)
			i++
		}
	}

	return tst
}

func (s *ExpContext) End(i int) IEndContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndContext); ok {
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

	return t.(IEndContext)
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
	p.EnterRule(localctx, 38, CrayonParserRULE_exp)
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

	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(154)
			p.Scope()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(157)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case CrayonParserT__0, CrayonParserT__2, CrayonParserT__5, CrayonParserT__7, CrayonParserT__9, CrayonParserT__10, CrayonParserT__11, CrayonParserNUMBER, CrayonParserSTRING, CrayonParserNONE, CrayonParserINFINITY, CrayonParserYES, CrayonParserNO, CrayonParserPI, CrayonParserIDENTIFIER:
					{
						p.SetState(155)
						p.Command()
					}

				case CrayonParserEND:
					{
						p.SetState(156)
						p.End()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}
				p.SetState(160)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for ok := true; ok; ok = _la == CrayonParserT__12 {
					{
						p.SetState(159)
						p.Match(CrayonParserT__12)
					}

					p.SetState(162)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(166)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
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

func (s *TagContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *TagContext) MainTag() IMainTagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMainTagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMainTagContext)
}

func (s *TagContext) FrameTag() IFrameTagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrameTagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrameTagContext)
}

func (s *TagContext) LoopTag() ILoopTagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopTagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopTagContext)
}

func (s *TagContext) CustomTag() ICustomTagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICustomTagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICustomTagContext)
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
	p.EnterRule(localctx, 40, CrayonParserRULE_tag)

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
		p.SetState(170)
		p.Match(CrayonParserT__13)
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CrayonParserMAIN:
		{
			p.SetState(171)
			p.MainTag()
		}

	case CrayonParserFRAME:
		{
			p.SetState(172)
			p.FrameTag()
		}

	case CrayonParserLOOP:
		{
			p.SetState(173)
			p.LoopTag()
		}

	case CrayonParserT__10:
		{
			p.SetState(174)
			p.CustomTag()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(177)
		p.Match(CrayonParserT__6)
	}
	{
		p.SetState(178)
		p.Exp()
	}

	return localctx
}

// IMainTagContext is an interface to support dynamic dispatch.
type IMainTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMainTagContext differentiates from other interfaces.
	IsMainTagContext()
}

type MainTagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainTagContext() *MainTagContext {
	var p = new(MainTagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_mainTag
	return p
}

func (*MainTagContext) IsMainTagContext() {}

func NewMainTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainTagContext {
	var p = new(MainTagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_mainTag

	return p
}

func (s *MainTagContext) GetParser() antlr.Parser { return s.parser }

func (s *MainTagContext) MAIN() antlr.TerminalNode {
	return s.GetToken(CrayonParserMAIN, 0)
}

func (s *MainTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterMainTag(s)
	}
}

func (s *MainTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitMainTag(s)
	}
}

func (s *MainTagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitMainTag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) MainTag() (localctx IMainTagContext) {
	this := p
	_ = this

	localctx = NewMainTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CrayonParserRULE_mainTag)

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
		p.SetState(180)
		p.Match(CrayonParserMAIN)
	}

	return localctx
}

// IFrameTagContext is an interface to support dynamic dispatch.
type IFrameTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFrameTagContext differentiates from other interfaces.
	IsFrameTagContext()
}

type FrameTagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrameTagContext() *FrameTagContext {
	var p = new(FrameTagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_frameTag
	return p
}

func (*FrameTagContext) IsFrameTagContext() {}

func NewFrameTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FrameTagContext {
	var p = new(FrameTagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_frameTag

	return p
}

func (s *FrameTagContext) GetParser() antlr.Parser { return s.parser }

func (s *FrameTagContext) FRAME() antlr.TerminalNode {
	return s.GetToken(CrayonParserFRAME, 0)
}

func (s *FrameTagContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CrayonParserNUMBER, 0)
}

func (s *FrameTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FrameTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FrameTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterFrameTag(s)
	}
}

func (s *FrameTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitFrameTag(s)
	}
}

func (s *FrameTagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitFrameTag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) FrameTag() (localctx IFrameTagContext) {
	this := p
	_ = this

	localctx = NewFrameTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CrayonParserRULE_frameTag)

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
		p.SetState(182)
		p.Match(CrayonParserFRAME)
	}
	{
		p.SetState(183)
		p.Match(CrayonParserNUMBER)
	}

	return localctx
}

// ILoopTagContext is an interface to support dynamic dispatch.
type ILoopTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopTagContext differentiates from other interfaces.
	IsLoopTagContext()
}

type LoopTagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopTagContext() *LoopTagContext {
	var p = new(LoopTagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_loopTag
	return p
}

func (*LoopTagContext) IsLoopTagContext() {}

func NewLoopTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopTagContext {
	var p = new(LoopTagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_loopTag

	return p
}

func (s *LoopTagContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopTagContext) LOOP() antlr.TerminalNode {
	return s.GetToken(CrayonParserLOOP, 0)
}

func (s *LoopTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterLoopTag(s)
	}
}

func (s *LoopTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitLoopTag(s)
	}
}

func (s *LoopTagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitLoopTag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) LoopTag() (localctx ILoopTagContext) {
	this := p
	_ = this

	localctx = NewLoopTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CrayonParserRULE_loopTag)

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
		p.SetState(185)
		p.Match(CrayonParserLOOP)
	}

	return localctx
}

// ICustomTagContext is an interface to support dynamic dispatch.
type ICustomTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCustomTagContext differentiates from other interfaces.
	IsCustomTagContext()
}

type CustomTagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCustomTagContext() *CustomTagContext {
	var p = new(CustomTagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CrayonParserRULE_customTag
	return p
}

func (*CustomTagContext) IsCustomTagContext() {}

func NewCustomTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CustomTagContext {
	var p = new(CustomTagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CrayonParserRULE_customTag

	return p
}

func (s *CustomTagContext) GetParser() antlr.Parser { return s.parser }

func (s *CustomTagContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CrayonParserIDENTIFIER, 0)
}

func (s *CustomTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CustomTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CustomTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.EnterCustomTag(s)
	}
}

func (s *CustomTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CrayonListener); ok {
		listenerT.ExitCustomTag(s)
	}
}

func (s *CustomTagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CrayonVisitor:
		return t.VisitCustomTag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CrayonParser) CustomTag() (localctx ICustomTagContext) {
	this := p
	_ = this

	localctx = NewCustomTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CrayonParserRULE_customTag)

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
		p.SetState(187)
		p.Match(CrayonParserT__10)
	}
	{
		p.SetState(188)
		p.Match(CrayonParserIDENTIFIER)
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
	p.EnterRule(localctx, 50, CrayonParserRULE_script)
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
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CrayonParserT__13 {
		{
			p.SetState(190)
			p.Tag()
		}

		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(195)
		p.Match(CrayonParserEOF)
	}

	return localctx
}
