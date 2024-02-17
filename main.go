package main

import (
	"lang/own"
	"lang/parsing"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parsing.NewlangLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parsing.NewlangParser(stream)
	tree := p.Program()
	visitor := &own.Visitor{}
	visitor.Visit(tree)
	antlr.ParseTreeWalkerDefault.Walk(&own.Listener{}, tree)
}
