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
	fmt.Printf("Enter ctx Text:<<%s>>\n", ctx.GetText())
}
func (c *CminusListener)  ExitEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("Exit ctx Text:<<%s>>\n", ctx.GetText())
}
