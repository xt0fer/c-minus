package codegen

import (
	"fmt"
	"github.com/xt0fer/cminus/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type CminusListener struct {
	*parser.BasecminusListener
}

func (c *CminusListener)  EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("ctx Text:<<%s>>\n", ctx.GetText())
}
