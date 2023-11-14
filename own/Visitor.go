package own

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"smell-hunter/parsing"
)

type Visitor struct {
	parsing.BaselangParserVisitor
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *Visitor) VisitProgram(ctx *parsing.ProgramContext) interface{} {
	fmt.Println("Visit program")
	return v.VisitChildren(ctx)
}
