// Generated from ./cminus.g4 by ANTLR 4.7.2
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link cminusParser}.
 */
public interface cminusListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link cminusParser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(cminusParser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(cminusParser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#statementList}.
	 * @param ctx the parse tree
	 */
	void enterStatementList(cminusParser.StatementListContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#statementList}.
	 * @param ctx the parse tree
	 */
	void exitStatementList(cminusParser.StatementListContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#declarationList}.
	 * @param ctx the parse tree
	 */
	void enterDeclarationList(cminusParser.DeclarationListContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#declarationList}.
	 * @param ctx the parse tree
	 */
	void exitDeclarationList(cminusParser.DeclarationListContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(cminusParser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(cminusParser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void enterAssignmentStatement(cminusParser.AssignmentStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void exitAssignmentStatement(cminusParser.AssignmentStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#compoundStatement}.
	 * @param ctx the parse tree
	 */
	void enterCompoundStatement(cminusParser.CompoundStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#compoundStatement}.
	 * @param ctx the parse tree
	 */
	void exitCompoundStatement(cminusParser.CompoundStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#blockItemList}.
	 * @param ctx the parse tree
	 */
	void enterBlockItemList(cminusParser.BlockItemListContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#blockItemList}.
	 * @param ctx the parse tree
	 */
	void exitBlockItemList(cminusParser.BlockItemListContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#blockItem}.
	 * @param ctx the parse tree
	 */
	void enterBlockItem(cminusParser.BlockItemContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#blockItem}.
	 * @param ctx the parse tree
	 */
	void exitBlockItem(cminusParser.BlockItemContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#declaration}.
	 * @param ctx the parse tree
	 */
	void enterDeclaration(cminusParser.DeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#declaration}.
	 * @param ctx the parse tree
	 */
	void exitDeclaration(cminusParser.DeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterExpressionStatement(cminusParser.ExpressionStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitExpressionStatement(cminusParser.ExpressionStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#selectionStatement}.
	 * @param ctx the parse tree
	 */
	void enterSelectionStatement(cminusParser.SelectionStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#selectionStatement}.
	 * @param ctx the parse tree
	 */
	void exitSelectionStatement(cminusParser.SelectionStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#iterationStatement}.
	 * @param ctx the parse tree
	 */
	void enterIterationStatement(cminusParser.IterationStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#iterationStatement}.
	 * @param ctx the parse tree
	 */
	void exitIterationStatement(cminusParser.IterationStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#jumpStatement}.
	 * @param ctx the parse tree
	 */
	void enterJumpStatement(cminusParser.JumpStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#jumpStatement}.
	 * @param ctx the parse tree
	 */
	void exitJumpStatement(cminusParser.JumpStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#unaryOperator}.
	 * @param ctx the parse tree
	 */
	void enterUnaryOperator(cminusParser.UnaryOperatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#unaryOperator}.
	 * @param ctx the parse tree
	 */
	void exitUnaryOperator(cminusParser.UnaryOperatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void enterPrimaryExpression(cminusParser.PrimaryExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void exitPrimaryExpression(cminusParser.PrimaryExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(cminusParser.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(cminusParser.ExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#addop}.
	 * @param ctx the parse tree
	 */
	void enterAddop(cminusParser.AddopContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#addop}.
	 * @param ctx the parse tree
	 */
	void exitAddop(cminusParser.AddopContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#mulop}.
	 * @param ctx the parse tree
	 */
	void enterMulop(cminusParser.MulopContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#mulop}.
	 * @param ctx the parse tree
	 */
	void exitMulop(cminusParser.MulopContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#conditionalExpression}.
	 * @param ctx the parse tree
	 */
	void enterConditionalExpression(cminusParser.ConditionalExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#conditionalExpression}.
	 * @param ctx the parse tree
	 */
	void exitConditionalExpression(cminusParser.ConditionalExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#assignmentExpression}.
	 * @param ctx the parse tree
	 */
	void enterAssignmentExpression(cminusParser.AssignmentExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#assignmentExpression}.
	 * @param ctx the parse tree
	 */
	void exitAssignmentExpression(cminusParser.AssignmentExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#assignmentOperator}.
	 * @param ctx the parse tree
	 */
	void enterAssignmentOperator(cminusParser.AssignmentOperatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#assignmentOperator}.
	 * @param ctx the parse tree
	 */
	void exitAssignmentOperator(cminusParser.AssignmentOperatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#constantExpression}.
	 * @param ctx the parse tree
	 */
	void enterConstantExpression(cminusParser.ConstantExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#constantExpression}.
	 * @param ctx the parse tree
	 */
	void exitConstantExpression(cminusParser.ConstantExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#shiftExpression}.
	 * @param ctx the parse tree
	 */
	void enterShiftExpression(cminusParser.ShiftExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#shiftExpression}.
	 * @param ctx the parse tree
	 */
	void exitShiftExpression(cminusParser.ShiftExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#relationalExpression}.
	 * @param ctx the parse tree
	 */
	void enterRelationalExpression(cminusParser.RelationalExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#relationalExpression}.
	 * @param ctx the parse tree
	 */
	void exitRelationalExpression(cminusParser.RelationalExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#equalityExpression}.
	 * @param ctx the parse tree
	 */
	void enterEqualityExpression(cminusParser.EqualityExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#equalityExpression}.
	 * @param ctx the parse tree
	 */
	void exitEqualityExpression(cminusParser.EqualityExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#andExpression}.
	 * @param ctx the parse tree
	 */
	void enterAndExpression(cminusParser.AndExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#andExpression}.
	 * @param ctx the parse tree
	 */
	void exitAndExpression(cminusParser.AndExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#exclusiveOrExpression}.
	 * @param ctx the parse tree
	 */
	void enterExclusiveOrExpression(cminusParser.ExclusiveOrExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#exclusiveOrExpression}.
	 * @param ctx the parse tree
	 */
	void exitExclusiveOrExpression(cminusParser.ExclusiveOrExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#inclusiveOrExpression}.
	 * @param ctx the parse tree
	 */
	void enterInclusiveOrExpression(cminusParser.InclusiveOrExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#inclusiveOrExpression}.
	 * @param ctx the parse tree
	 */
	void exitInclusiveOrExpression(cminusParser.InclusiveOrExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#logicalAndExpression}.
	 * @param ctx the parse tree
	 */
	void enterLogicalAndExpression(cminusParser.LogicalAndExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#logicalAndExpression}.
	 * @param ctx the parse tree
	 */
	void exitLogicalAndExpression(cminusParser.LogicalAndExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#logicalOrExpression}.
	 * @param ctx the parse tree
	 */
	void enterLogicalOrExpression(cminusParser.LogicalOrExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#logicalOrExpression}.
	 * @param ctx the parse tree
	 */
	void exitLogicalOrExpression(cminusParser.LogicalOrExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#typeSpecifier}.
	 * @param ctx the parse tree
	 */
	void enterTypeSpecifier(cminusParser.TypeSpecifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#typeSpecifier}.
	 * @param ctx the parse tree
	 */
	void exitTypeSpecifier(cminusParser.TypeSpecifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#functionDefinition}.
	 * @param ctx the parse tree
	 */
	void enterFunctionDefinition(cminusParser.FunctionDefinitionContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#functionDefinition}.
	 * @param ctx the parse tree
	 */
	void exitFunctionDefinition(cminusParser.FunctionDefinitionContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#paramList}.
	 * @param ctx the parse tree
	 */
	void enterParamList(cminusParser.ParamListContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#paramList}.
	 * @param ctx the parse tree
	 */
	void exitParamList(cminusParser.ParamListContext ctx);
	/**
	 * Enter a parse tree produced by {@link cminusParser#param}.
	 * @param ctx the parse tree
	 */
	void enterParam(cminusParser.ParamContext ctx);
	/**
	 * Exit a parse tree produced by {@link cminusParser#param}.
	 * @param ctx the parse tree
	 */
	void exitParam(cminusParser.ParamContext ctx);
}