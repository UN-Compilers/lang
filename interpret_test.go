package main

import (
	"lang/parsing"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
)

func TestIntEval(t *testing.T) {
	input := antlr.NewInputStream("1")
	lexer := parsing.NewlangLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parsing.NewlangParser(stream)
	tree := p.Exp().(*parsing.IntegerContext)
	interp := &Interpreter{}

	assert.Equal(t, 1, interp.VisitInteger(tree), "'1' should eval to 1")
}

func TestIntSumEval(t *testing.T) {
	input := antlr.NewInputStream("1+1")
	lexer := parsing.NewlangLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parsing.NewlangParser(stream)
	tree := p.Exp().(*parsing.SumContext)
	interp := &Interpreter{}

	assert.Equal(t, 2, interp.VisitSum(tree), "'1+1' should eval to 2")
}
