package own

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"lang/parsing"
)

type Listener struct {
	*parsing.BaselangParserListener
}

func (s *Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText(), "Every rule implementation")
}
