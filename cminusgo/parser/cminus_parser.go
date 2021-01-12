// Code generated from cminus.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // cminus

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 60, 225,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 5, 2, 53, 10, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 7, 3, 60, 10, 3, 12, 3, 14, 3, 63, 11, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 85, 10, 7, 12, 7, 14, 7, 88, 11,
	7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 97, 10, 8, 12, 8, 14,
	8, 100, 11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 107, 10, 9, 12, 9, 14,
	9, 110, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 117, 10, 10, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14,
	5, 14, 130, 10, 14, 3, 14, 5, 14, 133, 10, 14, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 153, 10, 16, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 167,
	10, 18, 3, 18, 5, 18, 170, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 182, 10, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 196,
	10, 19, 12, 19, 14, 19, 199, 11, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22,
	3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 214, 10,
	24, 12, 24, 14, 24, 217, 11, 24, 3, 25, 5, 25, 220, 10, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 2, 7, 4, 14, 16, 36, 46, 26, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 2, 5, 4,
	2, 34, 34, 36, 36, 3, 2, 38, 40, 3, 2, 26, 31, 2, 224, 2, 50, 3, 2, 2,
	2, 4, 54, 3, 2, 2, 2, 6, 64, 3, 2, 2, 2, 8, 69, 3, 2, 2, 2, 10, 71, 3,
	2, 2, 2, 12, 79, 3, 2, 2, 2, 14, 91, 3, 2, 2, 2, 16, 101, 3, 2, 2, 2, 18,
	116, 3, 2, 2, 2, 20, 118, 3, 2, 2, 2, 22, 120, 3, 2, 2, 2, 24, 125, 3,
	2, 2, 2, 26, 127, 3, 2, 2, 2, 28, 136, 3, 2, 2, 2, 30, 145, 3, 2, 2, 2,
	32, 154, 3, 2, 2, 2, 34, 169, 3, 2, 2, 2, 36, 181, 3, 2, 2, 2, 38, 200,
	3, 2, 2, 2, 40, 202, 3, 2, 2, 2, 42, 204, 3, 2, 2, 2, 44, 206, 3, 2, 2,
	2, 46, 208, 3, 2, 2, 2, 48, 219, 3, 2, 2, 2, 50, 52, 5, 6, 4, 2, 51, 53,
	5, 4, 3, 2, 52, 51, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 3, 3, 2, 2, 2,
	54, 55, 8, 3, 1, 2, 55, 56, 5, 10, 6, 2, 56, 61, 3, 2, 2, 2, 57, 58, 12,
	3, 2, 2, 58, 60, 5, 10, 6, 2, 59, 57, 3, 2, 2, 2, 60, 63, 3, 2, 2, 2, 61,
	59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 5, 3, 2, 2, 2, 63, 61, 3, 2, 2,
	2, 64, 65, 7, 17, 2, 2, 65, 66, 7, 20, 2, 2, 66, 67, 7, 21, 2, 2, 67, 68,
	5, 26, 14, 2, 68, 7, 3, 2, 2, 2, 69, 70, 5, 20, 11, 2, 70, 9, 3, 2, 2,
	2, 71, 72, 7, 16, 2, 2, 72, 73, 5, 8, 5, 2, 73, 74, 7, 20, 2, 2, 74, 75,
	5, 46, 24, 2, 75, 76, 7, 21, 2, 2, 76, 77, 5, 44, 23, 2, 77, 78, 5, 26,
	14, 2, 78, 11, 3, 2, 2, 2, 79, 80, 5, 8, 5, 2, 80, 81, 7, 20, 2, 2, 81,
	86, 5, 20, 11, 2, 82, 83, 7, 51, 2, 2, 83, 85, 5, 20, 11, 2, 84, 82, 3,
	2, 2, 2, 85, 88, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87,
	89, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 89, 90, 7, 21, 2, 2, 90, 13, 3, 2,
	2, 2, 91, 92, 8, 8, 1, 2, 92, 93, 5, 18, 10, 2, 93, 98, 3, 2, 2, 2, 94,
	95, 12, 3, 2, 2, 95, 97, 5, 18, 10, 2, 96, 94, 3, 2, 2, 2, 97, 100, 3,
	2, 2, 2, 98, 96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 15, 3, 2, 2, 2, 100,
	98, 3, 2, 2, 2, 101, 102, 8, 9, 1, 2, 102, 103, 5, 28, 15, 2, 103, 108,
	3, 2, 2, 2, 104, 105, 12, 3, 2, 2, 105, 107, 5, 28, 15, 2, 106, 104, 3,
	2, 2, 2, 107, 110, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2,
	2, 109, 17, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 111, 117, 5, 26, 14, 2, 112,
	117, 5, 30, 16, 2, 113, 117, 5, 32, 17, 2, 114, 117, 5, 34, 18, 2, 115,
	117, 5, 22, 12, 2, 116, 111, 3, 2, 2, 2, 116, 112, 3, 2, 2, 2, 116, 113,
	3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 116, 115, 3, 2, 2, 2, 117, 19, 3, 2,
	2, 2, 118, 119, 7, 53, 2, 2, 119, 21, 3, 2, 2, 2, 120, 121, 5, 20, 11,
	2, 121, 122, 5, 24, 13, 2, 122, 123, 5, 36, 19, 2, 123, 124, 7, 50, 2,
	2, 124, 23, 3, 2, 2, 2, 125, 126, 7, 52, 2, 2, 126, 25, 3, 2, 2, 2, 127,
	129, 7, 24, 2, 2, 128, 130, 5, 16, 9, 2, 129, 128, 3, 2, 2, 2, 129, 130,
	3, 2, 2, 2, 130, 132, 3, 2, 2, 2, 131, 133, 5, 14, 8, 2, 132, 131, 3, 2,
	2, 2, 132, 133, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 135, 7, 25, 2, 2,
	135, 27, 3, 2, 2, 2, 136, 137, 7, 7, 2, 2, 137, 138, 5, 20, 11, 2, 138,
	139, 7, 52, 2, 2, 139, 140, 5, 44, 23, 2, 140, 141, 7, 20, 2, 2, 141, 142,
	7, 54, 2, 2, 142, 143, 7, 21, 2, 2, 143, 144, 7, 50, 2, 2, 144, 29, 3,
	2, 2, 2, 145, 146, 7, 6, 2, 2, 146, 147, 7, 20, 2, 2, 147, 148, 5, 36,
	19, 2, 148, 149, 7, 21, 2, 2, 149, 152, 5, 26, 14, 2, 150, 151, 7, 5, 2,
	2, 151, 153, 5, 26, 14, 2, 152, 150, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2,
	153, 31, 3, 2, 2, 2, 154, 155, 7, 15, 2, 2, 155, 156, 7, 20, 2, 2, 156,
	157, 5, 36, 19, 2, 157, 158, 7, 21, 2, 2, 158, 159, 5, 26, 14, 2, 159,
	33, 3, 2, 2, 2, 160, 161, 7, 4, 2, 2, 161, 170, 7, 50, 2, 2, 162, 163,
	7, 3, 2, 2, 163, 170, 7, 50, 2, 2, 164, 166, 7, 13, 2, 2, 165, 167, 5,
	36, 19, 2, 166, 165, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 168, 3, 2,
	2, 2, 168, 170, 7, 50, 2, 2, 169, 160, 3, 2, 2, 2, 169, 162, 3, 2, 2, 2,
	169, 164, 3, 2, 2, 2, 170, 35, 3, 2, 2, 2, 171, 172, 8, 19, 1, 2, 172,
	182, 5, 12, 7, 2, 173, 174, 7, 36, 2, 2, 174, 182, 5, 36, 19, 9, 175, 176,
	7, 20, 2, 2, 176, 177, 5, 36, 19, 2, 177, 178, 7, 21, 2, 2, 178, 182, 3,
	2, 2, 2, 179, 182, 7, 54, 2, 2, 180, 182, 5, 20, 11, 2, 181, 171, 3, 2,
	2, 2, 181, 173, 3, 2, 2, 2, 181, 175, 3, 2, 2, 2, 181, 179, 3, 2, 2, 2,
	181, 180, 3, 2, 2, 2, 182, 197, 3, 2, 2, 2, 183, 184, 12, 7, 2, 2, 184,
	185, 5, 40, 21, 2, 185, 186, 5, 36, 19, 8, 186, 196, 3, 2, 2, 2, 187, 188,
	12, 6, 2, 2, 188, 189, 5, 38, 20, 2, 189, 190, 5, 36, 19, 7, 190, 196,
	3, 2, 2, 2, 191, 192, 12, 3, 2, 2, 192, 193, 5, 42, 22, 2, 193, 194, 5,
	36, 19, 4, 194, 196, 3, 2, 2, 2, 195, 183, 3, 2, 2, 2, 195, 187, 3, 2,
	2, 2, 195, 191, 3, 2, 2, 2, 196, 199, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2,
	197, 198, 3, 2, 2, 2, 198, 37, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 200, 201,
	9, 2, 2, 2, 201, 39, 3, 2, 2, 2, 202, 203, 9, 3, 2, 2, 203, 41, 3, 2, 2,
	2, 204, 205, 9, 4, 2, 2, 205, 43, 3, 2, 2, 2, 206, 207, 7, 8, 2, 2, 207,
	45, 3, 2, 2, 2, 208, 209, 8, 24, 1, 2, 209, 210, 5, 48, 25, 2, 210, 215,
	3, 2, 2, 2, 211, 212, 12, 3, 2, 2, 212, 214, 5, 48, 25, 2, 213, 211, 3,
	2, 2, 2, 214, 217, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2,
	2, 216, 47, 3, 2, 2, 2, 217, 215, 3, 2, 2, 2, 218, 220, 7, 51, 2, 2, 219,
	218, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 221, 3, 2, 2, 2, 221, 222,
	5, 44, 23, 2, 222, 223, 5, 20, 11, 2, 223, 49, 3, 2, 2, 2, 18, 52, 61,
	86, 98, 108, 116, 129, 132, 152, 166, 169, 181, 195, 197, 215, 219,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'break'", "'continue'", "'else'", "'if'", "'var'", "'int'", "'boolean'",
	"'string'", "'rune'", "'array'", "'return'", "'void'", "'while'", "'function'",
	"'main'", "'true'", "'false'", "'('", "')'", "'['", "']'", "'{'", "'}'",
	"'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'<<'", "'>>'", "'+'", "'++'",
	"'-'", "'--'", "'*'", "'/'", "'%'", "'&'", "'|'", "'&&'", "'||'", "'^'",
	"'!'", "'~'", "'?'", "':'", "';'", "','", "'='",
}
var symbolicNames = []string{
	"", "Break", "Continue", "Else", "If", "Var", "Int", "Bool", "StringType",
	"Rune", "Array", "Return", "Void", "While", "Func", "Main", "True", "False",
	"Paren", "Thesis", "LeftBracket", "RightBracket", "LeftBrace", "RightBrace",
	"EqualEqual", "NotEqual", "Less", "LessEqual", "Greater", "GreaterEqual",
	"LeftShift", "RightShift", "Plus", "PlusPlus", "Minus", "MinusMinus", "Star",
	"Div", "Mod", "And", "Or", "AndAnd", "OrOr", "Caret", "Not", "Tilde", "Question",
	"Colon", "Semi", "Comma", "Assign", "Identifier", "Constant", "DigitSequence",
	"StringLiteral", "Whitespace", "Newline", "BlockComment", "LineComment",
}

var ruleNames = []string{
	"program", "functionList", "mainFunction", "fname", "functionDefinition",
	"functionCall", "statementList", "declarationList", "statement", "variable",
	"assignmentStatement", "assignmentOperator", "compoundStatement", "declaration",
	"selectionStatement", "iterationStatement", "jumpStatement", "expression",
	"addop", "mulop", "relop", "typeSpecifier", "paramList", "param",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type cminusParser struct {
	*antlr.BaseParser
}

func NewcminusParser(input antlr.TokenStream) *cminusParser {
	this := new(cminusParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "cminus.g4"

	return this
}

// cminusParser tokens.
const (
	cminusParserEOF           = antlr.TokenEOF
	cminusParserBreak         = 1
	cminusParserContinue      = 2
	cminusParserElse          = 3
	cminusParserIf            = 4
	cminusParserVar           = 5
	cminusParserInt           = 6
	cminusParserBool          = 7
	cminusParserStringType    = 8
	cminusParserRune          = 9
	cminusParserArray         = 10
	cminusParserReturn        = 11
	cminusParserVoid          = 12
	cminusParserWhile         = 13
	cminusParserFunc          = 14
	cminusParserMain          = 15
	cminusParserTrue          = 16
	cminusParserFalse         = 17
	cminusParserParen         = 18
	cminusParserThesis        = 19
	cminusParserLeftBracket   = 20
	cminusParserRightBracket  = 21
	cminusParserLeftBrace     = 22
	cminusParserRightBrace    = 23
	cminusParserEqualEqual    = 24
	cminusParserNotEqual      = 25
	cminusParserLess          = 26
	cminusParserLessEqual     = 27
	cminusParserGreater       = 28
	cminusParserGreaterEqual  = 29
	cminusParserLeftShift     = 30
	cminusParserRightShift    = 31
	cminusParserPlus          = 32
	cminusParserPlusPlus      = 33
	cminusParserMinus         = 34
	cminusParserMinusMinus    = 35
	cminusParserStar          = 36
	cminusParserDiv           = 37
	cminusParserMod           = 38
	cminusParserAnd           = 39
	cminusParserOr            = 40
	cminusParserAndAnd        = 41
	cminusParserOrOr          = 42
	cminusParserCaret         = 43
	cminusParserNot           = 44
	cminusParserTilde         = 45
	cminusParserQuestion      = 46
	cminusParserColon         = 47
	cminusParserSemi          = 48
	cminusParserComma         = 49
	cminusParserAssign        = 50
	cminusParserIdentifier    = 51
	cminusParserConstant      = 52
	cminusParserDigitSequence = 53
	cminusParserStringLiteral = 54
	cminusParserWhitespace    = 55
	cminusParserNewline       = 56
	cminusParserBlockComment  = 57
	cminusParserLineComment   = 58
)

// cminusParser rules.
const (
	cminusParserRULE_program             = 0
	cminusParserRULE_functionList        = 1
	cminusParserRULE_mainFunction        = 2
	cminusParserRULE_fname               = 3
	cminusParserRULE_functionDefinition  = 4
	cminusParserRULE_functionCall        = 5
	cminusParserRULE_statementList       = 6
	cminusParserRULE_declarationList     = 7
	cminusParserRULE_statement           = 8
	cminusParserRULE_variable            = 9
	cminusParserRULE_assignmentStatement = 10
	cminusParserRULE_assignmentOperator  = 11
	cminusParserRULE_compoundStatement   = 12
	cminusParserRULE_declaration         = 13
	cminusParserRULE_selectionStatement  = 14
	cminusParserRULE_iterationStatement  = 15
	cminusParserRULE_jumpStatement       = 16
	cminusParserRULE_expression          = 17
	cminusParserRULE_addop               = 18
	cminusParserRULE_mulop               = 19
	cminusParserRULE_relop               = 20
	cminusParserRULE_typeSpecifier       = 21
	cminusParserRULE_paramList           = 22
	cminusParserRULE_param               = 23
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) MainFunction() IMainFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMainFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMainFunctionContext)
}

func (s *ProgramContext) FunctionList() IFunctionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionListContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *cminusParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, cminusParserRULE_program)
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
		p.SetState(48)
		p.MainFunction()
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == cminusParserFunc {
		{
			p.SetState(49)
			p.functionList(0)
		}

	}

	return localctx
}

// IFunctionListContext is an interface to support dynamic dispatch.
type IFunctionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionListContext differentiates from other interfaces.
	IsFunctionListContext()
}

type FunctionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionListContext() *FunctionListContext {
	var p = new(FunctionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_functionList
	return p
}

func (*FunctionListContext) IsFunctionListContext() {}

func NewFunctionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionListContext {
	var p = new(FunctionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_functionList

	return p
}

func (s *FunctionListContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionListContext) FunctionDefinition() IFunctionDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDefinitionContext)
}

func (s *FunctionListContext) FunctionList() IFunctionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionListContext)
}

func (s *FunctionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterFunctionList(s)
	}
}

func (s *FunctionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitFunctionList(s)
	}
}

func (p *cminusParser) FunctionList() (localctx IFunctionListContext) {
	return p.functionList(0)
}

func (p *cminusParser) functionList(_p int) (localctx IFunctionListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFunctionListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFunctionListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, cminusParserRULE_functionList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
		p.SetState(53)
		p.FunctionDefinition()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFunctionListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_functionList)
			p.SetState(55)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(56)
				p.FunctionDefinition()
			}

		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IMainFunctionContext is an interface to support dynamic dispatch.
type IMainFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMainFunctionContext differentiates from other interfaces.
	IsMainFunctionContext()
}

type MainFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainFunctionContext() *MainFunctionContext {
	var p = new(MainFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_mainFunction
	return p
}

func (*MainFunctionContext) IsMainFunctionContext() {}

func NewMainFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainFunctionContext {
	var p = new(MainFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_mainFunction

	return p
}

func (s *MainFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *MainFunctionContext) Main() antlr.TerminalNode {
	return s.GetToken(cminusParserMain, 0)
}

func (s *MainFunctionContext) Paren() antlr.TerminalNode {
	return s.GetToken(cminusParserParen, 0)
}

func (s *MainFunctionContext) Thesis() antlr.TerminalNode {
	return s.GetToken(cminusParserThesis, 0)
}

func (s *MainFunctionContext) CompoundStatement() ICompoundStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompoundStatementContext)
}

func (s *MainFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterMainFunction(s)
	}
}

func (s *MainFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitMainFunction(s)
	}
}

func (p *cminusParser) MainFunction() (localctx IMainFunctionContext) {
	localctx = NewMainFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, cminusParserRULE_mainFunction)

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
		p.SetState(62)
		p.Match(cminusParserMain)
	}
	{
		p.SetState(63)
		p.Match(cminusParserParen)
	}
	{
		p.SetState(64)
		p.Match(cminusParserThesis)
	}
	{
		p.SetState(65)
		p.CompoundStatement()
	}

	return localctx
}

// IFnameContext is an interface to support dynamic dispatch.
type IFnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFnameContext differentiates from other interfaces.
	IsFnameContext()
}

type FnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnameContext() *FnameContext {
	var p = new(FnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_fname
	return p
}

func (*FnameContext) IsFnameContext() {}

func NewFnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnameContext {
	var p = new(FnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_fname

	return p
}

func (s *FnameContext) GetParser() antlr.Parser { return s.parser }

func (s *FnameContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *FnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterFname(s)
	}
}

func (s *FnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitFname(s)
	}
}

func (p *cminusParser) Fname() (localctx IFnameContext) {
	localctx = NewFnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, cminusParserRULE_fname)

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
		p.Variable()
	}

	return localctx
}

// IFunctionDefinitionContext is an interface to support dynamic dispatch.
type IFunctionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDefinitionContext differentiates from other interfaces.
	IsFunctionDefinitionContext()
}

type FunctionDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefinitionContext() *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_functionDefinition
	return p
}

func (*FunctionDefinitionContext) IsFunctionDefinitionContext() {}

func NewFunctionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_functionDefinition

	return p
}

func (s *FunctionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefinitionContext) Func() antlr.TerminalNode {
	return s.GetToken(cminusParserFunc, 0)
}

func (s *FunctionDefinitionContext) Fname() IFnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFnameContext)
}

func (s *FunctionDefinitionContext) Paren() antlr.TerminalNode {
	return s.GetToken(cminusParserParen, 0)
}

func (s *FunctionDefinitionContext) ParamList() IParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *FunctionDefinitionContext) Thesis() antlr.TerminalNode {
	return s.GetToken(cminusParserThesis, 0)
}

func (s *FunctionDefinitionContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *FunctionDefinitionContext) CompoundStatement() ICompoundStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompoundStatementContext)
}

func (s *FunctionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterFunctionDefinition(s)
	}
}

func (s *FunctionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitFunctionDefinition(s)
	}
}

func (p *cminusParser) FunctionDefinition() (localctx IFunctionDefinitionContext) {
	localctx = NewFunctionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, cminusParserRULE_functionDefinition)

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
		p.Match(cminusParserFunc)
	}
	{
		p.SetState(70)
		p.Fname()
	}
	{
		p.SetState(71)
		p.Match(cminusParserParen)
	}
	{
		p.SetState(72)
		p.paramList(0)
	}
	{
		p.SetState(73)
		p.Match(cminusParserThesis)
	}
	{
		p.SetState(74)
		p.TypeSpecifier()
	}
	{
		p.SetState(75)
		p.CompoundStatement()
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) Fname() IFnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFnameContext)
}

func (s *FunctionCallContext) Paren() antlr.TerminalNode {
	return s.GetToken(cminusParserParen, 0)
}

func (s *FunctionCallContext) AllVariable() []IVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableContext)(nil)).Elem())
	var tst = make([]IVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableContext)
		}
	}

	return tst
}

func (s *FunctionCallContext) Variable(i int) IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *FunctionCallContext) Thesis() antlr.TerminalNode {
	return s.GetToken(cminusParserThesis, 0)
}

func (s *FunctionCallContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(cminusParserComma)
}

func (s *FunctionCallContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(cminusParserComma, i)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *cminusParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, cminusParserRULE_functionCall)
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
		p.SetState(77)
		p.Fname()
	}
	{
		p.SetState(78)
		p.Match(cminusParserParen)
	}
	{
		p.SetState(79)
		p.Variable()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == cminusParserComma {
		{
			p.SetState(80)
			p.Match(cminusParserComma)
		}
		{
			p.SetState(81)
			p.Variable()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(87)
		p.Match(cminusParserThesis)
	}

	return localctx
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_statementList
	return p
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitStatementList(s)
	}
}

func (p *cminusParser) StatementList() (localctx IStatementListContext) {
	return p.statementList(0)
}

func (p *cminusParser) statementList(_p int) (localctx IStatementListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IStatementListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, cminusParserRULE_statementList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
		p.SetState(90)
		p.Statement()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewStatementListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_statementList)
			p.SetState(92)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(93)
				p.Statement()
			}

		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IDeclarationListContext is an interface to support dynamic dispatch.
type IDeclarationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationListContext differentiates from other interfaces.
	IsDeclarationListContext()
}

type DeclarationListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationListContext() *DeclarationListContext {
	var p = new(DeclarationListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_declarationList
	return p
}

func (*DeclarationListContext) IsDeclarationListContext() {}

func NewDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationListContext {
	var p = new(DeclarationListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_declarationList

	return p
}

func (s *DeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationListContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *DeclarationListContext) DeclarationList() IDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationListContext)
}

func (s *DeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterDeclarationList(s)
	}
}

func (s *DeclarationListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitDeclarationList(s)
	}
}

func (p *cminusParser) DeclarationList() (localctx IDeclarationListContext) {
	return p.declarationList(0)
}

func (p *cminusParser) declarationList(_p int) (localctx IDeclarationListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDeclarationListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDeclarationListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, cminusParserRULE_declarationList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
		p.SetState(100)
		p.Declaration()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewDeclarationListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_declarationList)
			p.SetState(102)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(103)
				p.Declaration()
			}

		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CompoundStatement() ICompoundStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompoundStatementContext)
}

func (s *StatementContext) SelectionStatement() ISelectionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionStatementContext)
}

func (s *StatementContext) IterationStatement() IIterationStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIterationStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIterationStatementContext)
}

func (s *StatementContext) JumpStatement() IJumpStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJumpStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJumpStatementContext)
}

func (s *StatementContext) AssignmentStatement() IAssignmentStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *cminusParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, cminusParserRULE_statement)

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

	p.SetState(114)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case cminusParserLeftBrace:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.CompoundStatement()
		}

	case cminusParserIf:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.SelectionStatement()
		}

	case cminusParserWhile:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(111)
			p.IterationStatement()
		}

	case cminusParserBreak, cminusParserContinue, cminusParserReturn:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(112)
			p.JumpStatement()
		}

	case cminusParserIdentifier:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(113)
			p.AssignmentStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

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
	p.RuleIndex = cminusParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) Identifier() antlr.TerminalNode {
	return s.GetToken(cminusParserIdentifier, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *cminusParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, cminusParserRULE_variable)

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
		p.SetState(116)
		p.Match(cminusParserIdentifier)
	}

	return localctx
}

// IAssignmentStatementContext is an interface to support dynamic dispatch.
type IAssignmentStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentStatementContext differentiates from other interfaces.
	IsAssignmentStatementContext()
}

type AssignmentStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentStatementContext() *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_assignmentStatement
	return p
}

func (*AssignmentStatementContext) IsAssignmentStatementContext() {}

func NewAssignmentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_assignmentStatement

	return p
}

func (s *AssignmentStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentStatementContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AssignmentStatementContext) AssignmentOperator() IAssignmentOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentOperatorContext)
}

func (s *AssignmentStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStatementContext) Semi() antlr.TerminalNode {
	return s.GetToken(cminusParserSemi, 0)
}

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitAssignmentStatement(s)
	}
}

func (p *cminusParser) AssignmentStatement() (localctx IAssignmentStatementContext) {
	localctx = NewAssignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, cminusParserRULE_assignmentStatement)

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
		p.SetState(118)
		p.Variable()
	}
	{
		p.SetState(119)
		p.AssignmentOperator()
	}
	{
		p.SetState(120)
		p.expression(0)
	}
	{
		p.SetState(121)
		p.Match(cminusParserSemi)
	}

	return localctx
}

// IAssignmentOperatorContext is an interface to support dynamic dispatch.
type IAssignmentOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentOperatorContext differentiates from other interfaces.
	IsAssignmentOperatorContext()
}

type AssignmentOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentOperatorContext() *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_assignmentOperator
	return p
}

func (*AssignmentOperatorContext) IsAssignmentOperatorContext() {}

func NewAssignmentOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_assignmentOperator

	return p
}

func (s *AssignmentOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentOperatorContext) Assign() antlr.TerminalNode {
	return s.GetToken(cminusParserAssign, 0)
}

func (s *AssignmentOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterAssignmentOperator(s)
	}
}

func (s *AssignmentOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitAssignmentOperator(s)
	}
}

func (p *cminusParser) AssignmentOperator() (localctx IAssignmentOperatorContext) {
	localctx = NewAssignmentOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, cminusParserRULE_assignmentOperator)

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
		p.SetState(123)
		p.Match(cminusParserAssign)
	}

	return localctx
}

// ICompoundStatementContext is an interface to support dynamic dispatch.
type ICompoundStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompoundStatementContext differentiates from other interfaces.
	IsCompoundStatementContext()
}

type CompoundStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompoundStatementContext() *CompoundStatementContext {
	var p = new(CompoundStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_compoundStatement
	return p
}

func (*CompoundStatementContext) IsCompoundStatementContext() {}

func NewCompoundStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompoundStatementContext {
	var p = new(CompoundStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_compoundStatement

	return p
}

func (s *CompoundStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CompoundStatementContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(cminusParserLeftBrace, 0)
}

func (s *CompoundStatementContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(cminusParserRightBrace, 0)
}

func (s *CompoundStatementContext) DeclarationList() IDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationListContext)
}

func (s *CompoundStatementContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *CompoundStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompoundStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompoundStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterCompoundStatement(s)
	}
}

func (s *CompoundStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitCompoundStatement(s)
	}
}

func (p *cminusParser) CompoundStatement() (localctx ICompoundStatementContext) {
	localctx = NewCompoundStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, cminusParserRULE_compoundStatement)
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
		p.SetState(125)
		p.Match(cminusParserLeftBrace)
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == cminusParserVar {
		{
			p.SetState(126)
			p.declarationList(0)
		}

	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<cminusParserBreak)|(1<<cminusParserContinue)|(1<<cminusParserIf)|(1<<cminusParserReturn)|(1<<cminusParserWhile)|(1<<cminusParserLeftBrace))) != 0) || _la == cminusParserIdentifier {
		{
			p.SetState(129)
			p.statementList(0)
		}

	}
	{
		p.SetState(132)
		p.Match(cminusParserRightBrace)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) Var() antlr.TerminalNode {
	return s.GetToken(cminusParserVar, 0)
}

func (s *DeclarationContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *DeclarationContext) Assign() antlr.TerminalNode {
	return s.GetToken(cminusParserAssign, 0)
}

func (s *DeclarationContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *DeclarationContext) Paren() antlr.TerminalNode {
	return s.GetToken(cminusParserParen, 0)
}

func (s *DeclarationContext) Constant() antlr.TerminalNode {
	return s.GetToken(cminusParserConstant, 0)
}

func (s *DeclarationContext) Thesis() antlr.TerminalNode {
	return s.GetToken(cminusParserThesis, 0)
}

func (s *DeclarationContext) Semi() antlr.TerminalNode {
	return s.GetToken(cminusParserSemi, 0)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *cminusParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, cminusParserRULE_declaration)

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
		p.SetState(134)
		p.Match(cminusParserVar)
	}
	{
		p.SetState(135)
		p.Variable()
	}
	{
		p.SetState(136)
		p.Match(cminusParserAssign)
	}
	{
		p.SetState(137)
		p.TypeSpecifier()
	}
	{
		p.SetState(138)
		p.Match(cminusParserParen)
	}
	{
		p.SetState(139)
		p.Match(cminusParserConstant)
	}
	{
		p.SetState(140)
		p.Match(cminusParserThesis)
	}
	{
		p.SetState(141)
		p.Match(cminusParserSemi)
	}

	return localctx
}

// ISelectionStatementContext is an interface to support dynamic dispatch.
type ISelectionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionStatementContext differentiates from other interfaces.
	IsSelectionStatementContext()
}

type SelectionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionStatementContext() *SelectionStatementContext {
	var p = new(SelectionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_selectionStatement
	return p
}

func (*SelectionStatementContext) IsSelectionStatementContext() {}

func NewSelectionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionStatementContext {
	var p = new(SelectionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_selectionStatement

	return p
}

func (s *SelectionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionStatementContext) If() antlr.TerminalNode {
	return s.GetToken(cminusParserIf, 0)
}

func (s *SelectionStatementContext) Paren() antlr.TerminalNode {
	return s.GetToken(cminusParserParen, 0)
}

func (s *SelectionStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SelectionStatementContext) Thesis() antlr.TerminalNode {
	return s.GetToken(cminusParserThesis, 0)
}

func (s *SelectionStatementContext) AllCompoundStatement() []ICompoundStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICompoundStatementContext)(nil)).Elem())
	var tst = make([]ICompoundStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICompoundStatementContext)
		}
	}

	return tst
}

func (s *SelectionStatementContext) CompoundStatement(i int) ICompoundStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICompoundStatementContext)
}

func (s *SelectionStatementContext) Else() antlr.TerminalNode {
	return s.GetToken(cminusParserElse, 0)
}

func (s *SelectionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterSelectionStatement(s)
	}
}

func (s *SelectionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitSelectionStatement(s)
	}
}

func (p *cminusParser) SelectionStatement() (localctx ISelectionStatementContext) {
	localctx = NewSelectionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, cminusParserRULE_selectionStatement)

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
		p.SetState(143)
		p.Match(cminusParserIf)
	}
	{
		p.SetState(144)
		p.Match(cminusParserParen)
	}
	{
		p.SetState(145)
		p.expression(0)
	}
	{
		p.SetState(146)
		p.Match(cminusParserThesis)
	}
	{
		p.SetState(147)
		p.CompoundStatement()
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(148)
			p.Match(cminusParserElse)
		}
		{
			p.SetState(149)
			p.CompoundStatement()
		}

	}

	return localctx
}

// IIterationStatementContext is an interface to support dynamic dispatch.
type IIterationStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIterationStatementContext differentiates from other interfaces.
	IsIterationStatementContext()
}

type IterationStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIterationStatementContext() *IterationStatementContext {
	var p = new(IterationStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_iterationStatement
	return p
}

func (*IterationStatementContext) IsIterationStatementContext() {}

func NewIterationStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IterationStatementContext {
	var p = new(IterationStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_iterationStatement

	return p
}

func (s *IterationStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IterationStatementContext) While() antlr.TerminalNode {
	return s.GetToken(cminusParserWhile, 0)
}

func (s *IterationStatementContext) Paren() antlr.TerminalNode {
	return s.GetToken(cminusParserParen, 0)
}

func (s *IterationStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IterationStatementContext) Thesis() antlr.TerminalNode {
	return s.GetToken(cminusParserThesis, 0)
}

func (s *IterationStatementContext) CompoundStatement() ICompoundStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompoundStatementContext)
}

func (s *IterationStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IterationStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IterationStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterIterationStatement(s)
	}
}

func (s *IterationStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitIterationStatement(s)
	}
}

func (p *cminusParser) IterationStatement() (localctx IIterationStatementContext) {
	localctx = NewIterationStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, cminusParserRULE_iterationStatement)

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
		p.SetState(152)
		p.Match(cminusParserWhile)
	}
	{
		p.SetState(153)
		p.Match(cminusParserParen)
	}
	{
		p.SetState(154)
		p.expression(0)
	}
	{
		p.SetState(155)
		p.Match(cminusParserThesis)
	}
	{
		p.SetState(156)
		p.CompoundStatement()
	}

	return localctx
}

// IJumpStatementContext is an interface to support dynamic dispatch.
type IJumpStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJumpStatementContext differentiates from other interfaces.
	IsJumpStatementContext()
}

type JumpStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJumpStatementContext() *JumpStatementContext {
	var p = new(JumpStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_jumpStatement
	return p
}

func (*JumpStatementContext) IsJumpStatementContext() {}

func NewJumpStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JumpStatementContext {
	var p = new(JumpStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_jumpStatement

	return p
}

func (s *JumpStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *JumpStatementContext) Continue() antlr.TerminalNode {
	return s.GetToken(cminusParserContinue, 0)
}

func (s *JumpStatementContext) Semi() antlr.TerminalNode {
	return s.GetToken(cminusParserSemi, 0)
}

func (s *JumpStatementContext) Break() antlr.TerminalNode {
	return s.GetToken(cminusParserBreak, 0)
}

func (s *JumpStatementContext) Return() antlr.TerminalNode {
	return s.GetToken(cminusParserReturn, 0)
}

func (s *JumpStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *JumpStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JumpStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JumpStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterJumpStatement(s)
	}
}

func (s *JumpStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitJumpStatement(s)
	}
}

func (p *cminusParser) JumpStatement() (localctx IJumpStatementContext) {
	localctx = NewJumpStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, cminusParserRULE_jumpStatement)
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

	p.SetState(167)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case cminusParserContinue:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Match(cminusParserContinue)
		}
		{
			p.SetState(159)
			p.Match(cminusParserSemi)
		}

	case cminusParserBreak:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(160)
			p.Match(cminusParserBreak)
		}
		{
			p.SetState(161)
			p.Match(cminusParserSemi)
		}

	case cminusParserReturn:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(162)
			p.Match(cminusParserReturn)
		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == cminusParserParen || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(cminusParserMinus-34))|(1<<(cminusParserIdentifier-34))|(1<<(cminusParserConstant-34)))) != 0) {
			{
				p.SetState(163)
				p.expression(0)
			}

		}
		{
			p.SetState(166)
			p.Match(cminusParserSemi)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExpressionContext) Minus() antlr.TerminalNode {
	return s.GetToken(cminusParserMinus, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) Paren() antlr.TerminalNode {
	return s.GetToken(cminusParserParen, 0)
}

func (s *ExpressionContext) Thesis() antlr.TerminalNode {
	return s.GetToken(cminusParserThesis, 0)
}

func (s *ExpressionContext) Constant() antlr.TerminalNode {
	return s.GetToken(cminusParserConstant, 0)
}

func (s *ExpressionContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ExpressionContext) Mulop() IMulopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMulopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMulopContext)
}

func (s *ExpressionContext) Addop() IAddopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddopContext)
}

func (s *ExpressionContext) Relop() IRelopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelopContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *cminusParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *cminusParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, cminusParserRULE_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(170)
			p.FunctionCall()
		}

	case 2:
		{
			p.SetState(171)
			p.Match(cminusParserMinus)
		}
		{
			p.SetState(172)
			p.expression(7)
		}

	case 3:
		{
			p.SetState(173)
			p.Match(cminusParserParen)
		}
		{
			p.SetState(174)
			p.expression(0)
		}
		{
			p.SetState(175)
			p.Match(cminusParserThesis)
		}

	case 4:
		{
			p.SetState(177)
			p.Match(cminusParserConstant)
		}

	case 5:
		{
			p.SetState(178)
			p.Variable()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(193)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_expression)
				p.SetState(181)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(182)
					p.Mulop()
				}
				{
					p.SetState(183)
					p.expression(6)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_expression)
				p.SetState(185)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(186)
					p.Addop()
				}
				{
					p.SetState(187)
					p.expression(5)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_expression)
				p.SetState(189)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(190)
					p.Relop()
				}
				{
					p.SetState(191)
					p.expression(2)
				}

			}

		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}

	return localctx
}

// IAddopContext is an interface to support dynamic dispatch.
type IAddopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddopContext differentiates from other interfaces.
	IsAddopContext()
}

type AddopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddopContext() *AddopContext {
	var p = new(AddopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_addop
	return p
}

func (*AddopContext) IsAddopContext() {}

func NewAddopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddopContext {
	var p = new(AddopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_addop

	return p
}

func (s *AddopContext) GetParser() antlr.Parser { return s.parser }

func (s *AddopContext) Plus() antlr.TerminalNode {
	return s.GetToken(cminusParserPlus, 0)
}

func (s *AddopContext) Minus() antlr.TerminalNode {
	return s.GetToken(cminusParserMinus, 0)
}

func (s *AddopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterAddop(s)
	}
}

func (s *AddopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitAddop(s)
	}
}

func (p *cminusParser) Addop() (localctx IAddopContext) {
	localctx = NewAddopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, cminusParserRULE_addop)
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
		p.SetState(198)
		_la = p.GetTokenStream().LA(1)

		if !(_la == cminusParserPlus || _la == cminusParserMinus) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMulopContext is an interface to support dynamic dispatch.
type IMulopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMulopContext differentiates from other interfaces.
	IsMulopContext()
}

type MulopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulopContext() *MulopContext {
	var p = new(MulopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_mulop
	return p
}

func (*MulopContext) IsMulopContext() {}

func NewMulopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulopContext {
	var p = new(MulopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_mulop

	return p
}

func (s *MulopContext) GetParser() antlr.Parser { return s.parser }

func (s *MulopContext) Star() antlr.TerminalNode {
	return s.GetToken(cminusParserStar, 0)
}

func (s *MulopContext) Div() antlr.TerminalNode {
	return s.GetToken(cminusParserDiv, 0)
}

func (s *MulopContext) Mod() antlr.TerminalNode {
	return s.GetToken(cminusParserMod, 0)
}

func (s *MulopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MulopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterMulop(s)
	}
}

func (s *MulopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitMulop(s)
	}
}

func (p *cminusParser) Mulop() (localctx IMulopContext) {
	localctx = NewMulopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, cminusParserRULE_mulop)
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
		p.SetState(200)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(cminusParserStar-36))|(1<<(cminusParserDiv-36))|(1<<(cminusParserMod-36)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRelopContext is an interface to support dynamic dispatch.
type IRelopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelopContext differentiates from other interfaces.
	IsRelopContext()
}

type RelopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelopContext() *RelopContext {
	var p = new(RelopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_relop
	return p
}

func (*RelopContext) IsRelopContext() {}

func NewRelopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelopContext {
	var p = new(RelopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_relop

	return p
}

func (s *RelopContext) GetParser() antlr.Parser { return s.parser }

func (s *RelopContext) EqualEqual() antlr.TerminalNode {
	return s.GetToken(cminusParserEqualEqual, 0)
}

func (s *RelopContext) NotEqual() antlr.TerminalNode {
	return s.GetToken(cminusParserNotEqual, 0)
}

func (s *RelopContext) Less() antlr.TerminalNode {
	return s.GetToken(cminusParserLess, 0)
}

func (s *RelopContext) LessEqual() antlr.TerminalNode {
	return s.GetToken(cminusParserLessEqual, 0)
}

func (s *RelopContext) Greater() antlr.TerminalNode {
	return s.GetToken(cminusParserGreater, 0)
}

func (s *RelopContext) GreaterEqual() antlr.TerminalNode {
	return s.GetToken(cminusParserGreaterEqual, 0)
}

func (s *RelopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterRelop(s)
	}
}

func (s *RelopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitRelop(s)
	}
}

func (p *cminusParser) Relop() (localctx IRelopContext) {
	localctx = NewRelopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, cminusParserRULE_relop)
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
		p.SetState(202)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<cminusParserEqualEqual)|(1<<cminusParserNotEqual)|(1<<cminusParserLess)|(1<<cminusParserLessEqual)|(1<<cminusParserGreater)|(1<<cminusParserGreaterEqual))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITypeSpecifierContext is an interface to support dynamic dispatch.
type ITypeSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecifierContext differentiates from other interfaces.
	IsTypeSpecifierContext()
}

type TypeSpecifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecifierContext() *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_typeSpecifier
	return p
}

func (*TypeSpecifierContext) IsTypeSpecifierContext() {}

func NewTypeSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_typeSpecifier

	return p
}

func (s *TypeSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecifierContext) Int() antlr.TerminalNode {
	return s.GetToken(cminusParserInt, 0)
}

func (s *TypeSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterTypeSpecifier(s)
	}
}

func (s *TypeSpecifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitTypeSpecifier(s)
	}
}

func (p *cminusParser) TypeSpecifier() (localctx ITypeSpecifierContext) {
	localctx = NewTypeSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, cminusParserRULE_typeSpecifier)

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
		p.SetState(204)
		p.Match(cminusParserInt)
	}

	return localctx
}

// IParamListContext is an interface to support dynamic dispatch.
type IParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamListContext differentiates from other interfaces.
	IsParamListContext()
}

type ParamListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamListContext() *ParamListContext {
	var p = new(ParamListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_paramList
	return p
}

func (*ParamListContext) IsParamListContext() {}

func NewParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListContext {
	var p = new(ParamListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_paramList

	return p
}

func (s *ParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListContext) Param() IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamListContext) ParamList() IParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterParamList(s)
	}
}

func (s *ParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitParamList(s)
	}
}

func (p *cminusParser) ParamList() (localctx IParamListContext) {
	return p.paramList(0)
}

func (p *cminusParser) paramList(_p int) (localctx IParamListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewParamListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IParamListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, cminusParserRULE_paramList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
		p.SetState(207)
		p.Param()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewParamListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_paramList)
			p.SetState(209)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(210)
				p.Param()
			}

		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}

	return localctx
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *ParamContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ParamContext) Comma() antlr.TerminalNode {
	return s.GetToken(cminusParserComma, 0)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *cminusParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, cminusParserRULE_param)
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
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == cminusParserComma {
		{
			p.SetState(216)
			p.Match(cminusParserComma)
		}

	}
	{
		p.SetState(219)
		p.TypeSpecifier()
	}
	{
		p.SetState(220)
		p.Variable()
	}

	return localctx
}

func (p *cminusParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *FunctionListContext = nil
		if localctx != nil {
			t = localctx.(*FunctionListContext)
		}
		return p.FunctionList_Sempred(t, predIndex)

	case 6:
		var t *StatementListContext = nil
		if localctx != nil {
			t = localctx.(*StatementListContext)
		}
		return p.StatementList_Sempred(t, predIndex)

	case 7:
		var t *DeclarationListContext = nil
		if localctx != nil {
			t = localctx.(*DeclarationListContext)
		}
		return p.DeclarationList_Sempred(t, predIndex)

	case 17:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 22:
		var t *ParamListContext = nil
		if localctx != nil {
			t = localctx.(*ParamListContext)
		}
		return p.ParamList_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *cminusParser) FunctionList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *cminusParser) StatementList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *cminusParser) DeclarationList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *cminusParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *cminusParser) ParamList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
