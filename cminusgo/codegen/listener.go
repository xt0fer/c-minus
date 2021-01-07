package codegen

import (
	"fmt"
	"github.com/xt0fer/cminus/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type CminusListener struct {
	*parser.BasecminusListener
	GlobalRuleNames []string
}

func (c *CminusListener)  EnterEveryRule(ctx antlr.ParserRuleContext) {
	ruleNum := ctx.GetRuleIndex()
	//parser := ctx.Parser
	fmt.Printf("Enter ctx RuleNum: %d Text:<<%s>> Rule: %+v\n", 
		ruleNum, ctx.GetText(), c.GlobalRuleNames[ruleNum])
}
func (c *CminusListener)  ExitEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Printf("Exit ctx Text:<<%s>>\n", ctx.GetText())
}
