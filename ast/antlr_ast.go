package ast

import (
	"lang/parsing"
	"strconv"
)

type CSTWalker struct {
	parsing.BaselangParserVisitor
}

func (c *CSTWalker) VisitInteger(ctx *parsing.IntegerContext) Integer {
	val, _ := strconv.Atoi(ctx.INTEGER().GetText())
	return Integer{val}
}

func (c *CSTWalker) VisitSum(ctx *parsing.SumContext) Infix {
	return Infix{"+", c.Visit(ctx.Exp(0)), c.Visit(ctx.Exp(1))}
}

func (c *CSTWalker) VisitProduct(ctx *parsing.ProductContext) Infix {
	return Infix{"*", c.Visit(ctx.Exp(0)), c.Visit(ctx.Exp(1))}
}

func (c *CSTWalker) VisitParens(ctx *parsing.ParensContext) interface{} {
	return c.Visit(ctx.Exp())
}

func (c *CSTWalker) VisitIdent(ctx *parsing.IdentContext) Identifier {
	return Identifier{ctx.IDENTIFIER().GetText()}
}

func (c *CSTWalker) VisitReal(ctx *parsing.RealContext) Real {
	val, _ := strconv.ParseFloat(ctx.REAL().GetText(), 64)
	return Real{val}
}

func (c *CSTWalker) VisitLet_statement(ctx *parsing.Let_statementContext) LetStatement {
	return LetStatement{ctx.IDENTIFIER(0).GetText(), ctx.IDENTIFIER(1).GetText(), c.Visit(ctx.Exp())}
}

func (c *CSTWalker) VisitProgram(ctx *parsing.ProgramContext) Program {
	list := make([]interface{}, 0)
	return Program{list}
}
