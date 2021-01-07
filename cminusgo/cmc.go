package main

// antlr4 -Dlanguage=Go -o parser cminus.g4

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/xt0fer/cminus/codegen"
	"github.com/xt0fer/cminus/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	fmt.Println("// *** cminus v1.0")
	// Setup the input
	path, err := os.Getwd()
	inputFilename := "testfiles/two.cminus"
	fileContents, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		panic("input file not found " + inputFilename + " CWD is " + path)
	}
	is := antlr.NewInputStream(string(fileContents))
	fmt.Println("// *** input: " + inputFilename)

	fmt.Println("// *** Create the Lexer")
	// Create the Lexer
	lexer := parser.NewcminusLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// for {
	// 	t := lexer.NextToken()
	// 	if t.GetTokenType() == antlr.TokenEOF {
	// 		break
	// 	}
	// 	fmt.Printf("%s (%q)\n",
	// 		lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	// }

	// after this loop, the stream is EMPTY!

	fmt.Println("// *** Create the Parser")
	// Create the Parser
	p := parser.NewcminusParser(stream)

	//fmt.Println("// *** Program")
	//tree := p.Program()

	ruleNames := p.RuleNames
	fmt.Println("// *** ParseTreeWalkerDefault.Walk")
	// Finally parse the expression

	antlr.ParseTreeWalkerDefault.Walk(&codegen.CminusListener{GlobalRuleNames: ruleNames}, p.Program())
}
