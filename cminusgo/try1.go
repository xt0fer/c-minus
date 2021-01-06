package main

// antlr4 -Dlanguage=Go -o parser cminus.g4

import (
	"fmt"
	"io/ioutil"
	"os"
	"github.com/xt0fer/cminus/parser"
	"github.com/xt0fer/cminus/codegen"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	fmt.Println("// *** cminus v1.0")
	// Setup the input
	path, err := os.Getwd()
	inputFilename := "testfiles/one.cminus"
	fileContents, err := ioutil.ReadFile(inputFilename)
	if (err != nil){
		panic("input file not found " + inputFilename+ " dir is "+path)
	}
	is := antlr.NewInputStream(string(fileContents))

	// Create the Lexer
	lexer := parser.NewcminusLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewcminusParser(stream)
	
	tree := p.Program()

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&codegen.CminusListener{}, tree)
}