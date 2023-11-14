package own

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"smell-hunter/parsing"
)

type Listener struct {
	*parsing.BaseJavaScriptParserListener
}

func (s *Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText(), "Every rule implementation")
}
