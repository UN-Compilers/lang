package main

import (
	"lang/parsing"
	"strconv"
)

type Interpreter struct {
	parsing.BaselangParserVisitor
}

func (i *Interpreter) VisitProgram(ctx *parsing.ProgramContext) interface{} {
	return i.VisitChildren(ctx)
}

func (i *Interpreter) VisitInteger(ctx *parsing.IntegerContext) int {
	res, _ := strconv.Atoi(ctx.INTEGER().GetText())
	return res
}

func (i *Interpreter) VisitSum(ctx *parsing.SumContext) interface{} {
	return i.VisitInteger(ctx.Exp(0).(*parsing.IntegerContext)) + i.VisitInteger(ctx.Exp(1).(*parsing.IntegerContext))
}

func (i *Interpreter) VisitParens(ctx *parsing.ParensContext) interface{} {
	return i.VisitInteger(ctx.Exp().(*parsing.IntegerContext))
}
