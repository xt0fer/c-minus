// Code generated from cminus.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // cminus

import "github.com/antlr/antlr4/runtime/Go/antlr"

// cminusListener is a complete listener for a parse tree produced by cminusParser.
type cminusListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterDeclarationList is called when entering the declarationList production.
	EnterDeclarationList(c *DeclarationListContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssignmentStatement is called when entering the assignmentStatement production.
	EnterAssignmentStatement(c *AssignmentStatementContext)

	// EnterCompoundStatement is called when entering the compoundStatement production.
	EnterCompoundStatement(c *CompoundStatementContext)

	// EnterBlockItemList is called when entering the blockItemList production.
	EnterBlockItemList(c *BlockItemListContext)

	// EnterBlockItem is called when entering the blockItem production.
	EnterBlockItem(c *BlockItemContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterSelectionStatement is called when entering the selectionStatement production.
	EnterSelectionStatement(c *SelectionStatementContext)

	// EnterIterationStatement is called when entering the iterationStatement production.
	EnterIterationStatement(c *IterationStatementContext)

	// EnterJumpStatement is called when entering the jumpStatement production.
	EnterJumpStatement(c *JumpStatementContext)

	// EnterUnaryOperator is called when entering the unaryOperator production.
	EnterUnaryOperator(c *UnaryOperatorContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAddop is called when entering the addop production.
	EnterAddop(c *AddopContext)

	// EnterMulop is called when entering the mulop production.
	EnterMulop(c *MulopContext)

	// EnterConditionalExpression is called when entering the conditionalExpression production.
	EnterConditionalExpression(c *ConditionalExpressionContext)

	// EnterAssignmentExpression is called when entering the assignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterAssignmentOperator is called when entering the assignmentOperator production.
	EnterAssignmentOperator(c *AssignmentOperatorContext)

	// EnterConstantExpression is called when entering the constantExpression production.
	EnterConstantExpression(c *ConstantExpressionContext)

	// EnterShiftExpression is called when entering the shiftExpression production.
	EnterShiftExpression(c *ShiftExpressionContext)

	// EnterRelationalExpression is called when entering the relationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterEqualityExpression is called when entering the equalityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterAndExpression is called when entering the andExpression production.
	EnterAndExpression(c *AndExpressionContext)

	// EnterExclusiveOrExpression is called when entering the exclusiveOrExpression production.
	EnterExclusiveOrExpression(c *ExclusiveOrExpressionContext)

	// EnterInclusiveOrExpression is called when entering the inclusiveOrExpression production.
	EnterInclusiveOrExpression(c *InclusiveOrExpressionContext)

	// EnterLogicalAndExpression is called when entering the logicalAndExpression production.
	EnterLogicalAndExpression(c *LogicalAndExpressionContext)

	// EnterLogicalOrExpression is called when entering the logicalOrExpression production.
	EnterLogicalOrExpression(c *LogicalOrExpressionContext)

	// EnterTypeSpecifier is called when entering the typeSpecifier production.
	EnterTypeSpecifier(c *TypeSpecifierContext)

	// EnterFunctionDefinition is called when entering the functionDefinition production.
	EnterFunctionDefinition(c *FunctionDefinitionContext)

	// EnterParamList is called when entering the paramList production.
	EnterParamList(c *ParamListContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitDeclarationList is called when exiting the declarationList production.
	ExitDeclarationList(c *DeclarationListContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssignmentStatement is called when exiting the assignmentStatement production.
	ExitAssignmentStatement(c *AssignmentStatementContext)

	// ExitCompoundStatement is called when exiting the compoundStatement production.
	ExitCompoundStatement(c *CompoundStatementContext)

	// ExitBlockItemList is called when exiting the blockItemList production.
	ExitBlockItemList(c *BlockItemListContext)

	// ExitBlockItem is called when exiting the blockItem production.
	ExitBlockItem(c *BlockItemContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitSelectionStatement is called when exiting the selectionStatement production.
	ExitSelectionStatement(c *SelectionStatementContext)

	// ExitIterationStatement is called when exiting the iterationStatement production.
	ExitIterationStatement(c *IterationStatementContext)

	// ExitJumpStatement is called when exiting the jumpStatement production.
	ExitJumpStatement(c *JumpStatementContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAddop is called when exiting the addop production.
	ExitAddop(c *AddopContext)

	// ExitMulop is called when exiting the mulop production.
	ExitMulop(c *MulopContext)

	// ExitConditionalExpression is called when exiting the conditionalExpression production.
	ExitConditionalExpression(c *ConditionalExpressionContext)

	// ExitAssignmentExpression is called when exiting the assignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitAssignmentOperator is called when exiting the assignmentOperator production.
	ExitAssignmentOperator(c *AssignmentOperatorContext)

	// ExitConstantExpression is called when exiting the constantExpression production.
	ExitConstantExpression(c *ConstantExpressionContext)

	// ExitShiftExpression is called when exiting the shiftExpression production.
	ExitShiftExpression(c *ShiftExpressionContext)

	// ExitRelationalExpression is called when exiting the relationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitEqualityExpression is called when exiting the equalityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitAndExpression is called when exiting the andExpression production.
	ExitAndExpression(c *AndExpressionContext)

	// ExitExclusiveOrExpression is called when exiting the exclusiveOrExpression production.
	ExitExclusiveOrExpression(c *ExclusiveOrExpressionContext)

	// ExitInclusiveOrExpression is called when exiting the inclusiveOrExpression production.
	ExitInclusiveOrExpression(c *InclusiveOrExpressionContext)

	// ExitLogicalAndExpression is called when exiting the logicalAndExpression production.
	ExitLogicalAndExpression(c *LogicalAndExpressionContext)

	// ExitLogicalOrExpression is called when exiting the logicalOrExpression production.
	ExitLogicalOrExpression(c *LogicalOrExpressionContext)

	// ExitTypeSpecifier is called when exiting the typeSpecifier production.
	ExitTypeSpecifier(c *TypeSpecifierContext)

	// ExitFunctionDefinition is called when exiting the functionDefinition production.
	ExitFunctionDefinition(c *FunctionDefinitionContext)

	// ExitParamList is called when exiting the paramList production.
	ExitParamList(c *ParamListContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)
}
