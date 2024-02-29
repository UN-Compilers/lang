package ast

type Identifier struct {
	val string
}

type Infix struct {
	op  string
	lhs interface{}
	rhs interface{}
}

type Integer struct {
	val int
}

type LetStatement struct {
	kind string
	name string
	val  interface{}
}

type Program struct {
	sentences []interface{}
}

type Real struct {
	val float64
}
