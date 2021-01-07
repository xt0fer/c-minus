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

// EnterFname is called when production fname is entered.
func (s *BasecminusListener) EnterFname(ctx *FnameContext) {}

// ExitFname is called when production fname is exited.
func (s *BasecminusListener) ExitFname(ctx *FnameContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *BasecminusListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *BasecminusListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BasecminusListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BasecminusListener) ExitFunctionCall(ctx *FunctionCallContext) {}

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

// EnterVariable is called when production variable is entered.
func (s *BasecminusListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BasecminusListener) ExitVariable(ctx *VariableContext) {}

// EnterAssignmentStatement is called when production assignmentStatement is entered.
func (s *BasecminusListener) EnterAssignmentStatement(ctx *AssignmentStatementContext) {}

// ExitAssignmentStatement is called when production assignmentStatement is exited.
func (s *BasecminusListener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BasecminusListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BasecminusListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *BasecminusListener) EnterCompoundStatement(ctx *CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *BasecminusListener) ExitCompoundStatement(ctx *CompoundStatementContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BasecminusListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BasecminusListener) ExitDeclaration(ctx *DeclarationContext) {}

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

// EnterRelop is called when production relop is entered.
func (s *BasecminusListener) EnterRelop(ctx *RelopContext) {}

// ExitRelop is called when production relop is exited.
func (s *BasecminusListener) ExitRelop(ctx *RelopContext) {}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *BasecminusListener) EnterTypeSpecifier(ctx *TypeSpecifierContext) {}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *BasecminusListener) ExitTypeSpecifier(ctx *TypeSpecifierContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BasecminusListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BasecminusListener) ExitParamList(ctx *ParamListContext) {}

// EnterParam is called when production param is entered.
func (s *BasecminusListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BasecminusListener) ExitParam(ctx *ParamContext) {}
