package ast

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
	walker := &CSTWalker{}

	assert.Equal(t, Integer{1}, walker.VisitInteger(tree), "'1' should give a Integer node with value 1")
}
