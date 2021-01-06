// Code generated from cminus.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // cminus

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasecminusListener is a complete listener for a parse tree produced by cminusParser.
type BasecminusListener struct{}

var _ cminusListener = &BasecminusListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasecminusListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasecminusListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasecminusListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasecminusListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasecminusListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasecminusListener) ExitProgram(ctx *ProgramContext) {}

// EnterFunctionList is called when production functionList is entered.
func (s *BasecminusListener) EnterFunctionList(ctx *FunctionListContext) {}

// ExitFunctionList is called when production functionList is exited.
func (s *BasecminusListener) ExitFunctionList(ctx *FunctionListContext) {}

// EnterMainFunction is called when production mainFunction is entered.
func (s *BasecminusListener) EnterMainFunction(ctx *MainFunctionContext) {}

// ExitMainFunction is called when production mainFunction is exited.
func (s *BasecminusListener) ExitMainFunction(ctx *MainFunctionContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BasecminusListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BasecminusListener) ExitStatementList(ctx *StatementListContext) {}

// EnterDeclarationList is called when production declarationList is entered.
func (s *BasecminusListener) EnterDeclarationList(ctx *DeclarationListContext) {}

// ExitDeclarationList is called when production declarationList is exited.
func (s *BasecminusListener) ExitDeclarationList(ctx *DeclarationListContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasecminusListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasecminusListener) ExitStatement(ctx *StatementContext) {}

// EnterAssignmentStatement is called when production assignmentStatement is entered.
func (s *BasecminusListener) EnterAssignmentStatement(ctx *AssignmentStatementContext) {}

// ExitAssignmentStatement is called when production assignmentStatement is exited.
func (s *BasecminusListener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *BasecminusListener) EnterCompoundStatement(ctx *CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *BasecminusListener) ExitCompoundStatement(ctx *CompoundStatementContext) {}

// EnterBlockItemList is called when production blockItemList is entered.
func (s *BasecminusListener) EnterBlockItemList(ctx *BlockItemListContext) {}

// ExitBlockItemList is called when production blockItemList is exited.
func (s *BasecminusListener) ExitBlockItemList(ctx *BlockItemListContext) {}

// EnterBlockItem is called when production blockItem is entered.
func (s *BasecminusListener) EnterBlockItem(ctx *BlockItemContext) {}

// ExitBlockItem is called when production blockItem is exited.
func (s *BasecminusListener) ExitBlockItem(ctx *BlockItemContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BasecminusListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BasecminusListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BasecminusListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BasecminusListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterSelectionStatement is called when production selectionStatement is entered.
func (s *BasecminusListener) EnterSelectionStatement(ctx *SelectionStatementContext) {}

// ExitSelectionStatement is called when production selectionStatement is exited.
func (s *BasecminusListener) ExitSelectionStatement(ctx *SelectionStatementContext) {}

// EnterIterationStatement is called when production iterationStatement is entered.
func (s *BasecminusListener) EnterIterationStatement(ctx *IterationStatementContext) {}

// ExitIterationStatement is called when production iterationStatement is exited.
func (s *BasecminusListener) ExitIterationStatement(ctx *IterationStatementContext) {}

// EnterJumpStatement is called when production jumpStatement is entered.
func (s *BasecminusListener) EnterJumpStatement(ctx *JumpStatementContext) {}

// ExitJumpStatement is called when production jumpStatement is exited.
func (s *BasecminusListener) ExitJumpStatement(ctx *JumpStatementContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BasecminusListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BasecminusListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BasecminusListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BasecminusListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasecminusListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasecminusListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAddop is called when production addop is entered.
func (s *BasecminusListener) EnterAddop(ctx *AddopContext) {}

// ExitAddop is called when production addop is exited.
func (s *BasecminusListener) ExitAddop(ctx *AddopContext) {}

// EnterMulop is called when production mulop is entered.
func (s *BasecminusListener) EnterMulop(ctx *MulopContext) {}

// ExitMulop is called when production mulop is exited.
func (s *BasecminusListener) ExitMulop(ctx *MulopContext) {}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *BasecminusListener) EnterConditionalExpression(ctx *ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *BasecminusListener) ExitConditionalExpression(ctx *ConditionalExpressionContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *BasecminusListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *BasecminusListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BasecminusListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BasecminusListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterConstantExpression is called when production constantExpression is entered.
func (s *BasecminusListener) EnterConstantExpression(ctx *ConstantExpressionContext) {}

// ExitConstantExpression is called when production constantExpression is exited.
func (s *BasecminusListener) ExitConstantExpression(ctx *ConstantExpressionContext) {}

// EnterShiftExpression is called when production shiftExpression is entered.
func (s *BasecminusListener) EnterShiftExpression(ctx *ShiftExpressionContext) {}

// ExitShiftExpression is called when production shiftExpression is exited.
func (s *BasecminusListener) ExitShiftExpression(ctx *ShiftExpressionContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BasecminusListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BasecminusListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BasecminusListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BasecminusListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BasecminusListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BasecminusListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterExclusiveOrExpression is called when production exclusiveOrExpression is entered.
func (s *BasecminusListener) EnterExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {}

// ExitExclusiveOrExpression is called when production exclusiveOrExpression is exited.
func (s *BasecminusListener) ExitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {}

// EnterInclusiveOrExpression is called when production inclusiveOrExpression is entered.
func (s *BasecminusListener) EnterInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {}

// ExitInclusiveOrExpression is called when production inclusiveOrExpression is exited.
func (s *BasecminusListener) ExitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {}

// EnterLogicalAndExpression is called when production logicalAndExpression is entered.
func (s *BasecminusListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production logicalAndExpression is exited.
func (s *BasecminusListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterLogicalOrExpression is called when production logicalOrExpression is entered.
func (s *BasecminusListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production logicalOrExpression is exited.
func (s *BasecminusListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *BasecminusListener) EnterTypeSpecifier(ctx *TypeSpecifierContext) {}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *BasecminusListener) ExitTypeSpecifier(ctx *TypeSpecifierContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *BasecminusListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *BasecminusListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BasecminusListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BasecminusListener) ExitParamList(ctx *ParamListContext) {}

// EnterParam is called when production param is entered.
func (s *BasecminusListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BasecminusListener) ExitParam(ctx *ParamContext) {}
