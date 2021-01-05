package main

// antlr -Dlanguage=Go -o parser gminus.g4

import (
	"github.com/xt0fer/cminus/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type cminusListener struct {
	*parser.BasecminusListener
}

func main() {
	// Setup the input
	is := antlr.NewInputStream("while (5) { ; }")

	// Create the Lexer
	lexer := parser.NewcminusLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewcminusParser(stream)
	
	tree := p.Program()

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&cminusListener{}, tree)
}