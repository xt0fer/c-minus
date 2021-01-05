// Generated from /home/kristofer/p/c-minus/cminusgo/Cminus.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class CminusParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, Break=4, Continue=5, Else=6, If=7, Int=8, Bool=9, 
		StringType=10, Rune=11, Array=12, Return=13, Void=14, While=15, Func=16, 
		True=17, False=18, Paren=19, Thesis=20, LeftBracket=21, RightBracket=22, 
		LeftBrace=23, RightBrace=24, Less=25, LessEqual=26, Greater=27, GreaterEqual=28, 
		LeftShift=29, RightShift=30, Plus=31, PlusPlus=32, Minus=33, MinusMinus=34, 
		Star=35, Div=36, Mod=37, And=38, Or=39, AndAnd=40, OrOr=41, Caret=42, 
		Not=43, Tilde=44, Question=45, Colon=46, Semi=47, Comma=48, Assign=49, 
		Identifier=50, Constant=51, DigitSequence=52, StringLiteral=53, Whitespace=54, 
		Newline=55, BlockComment=56, LineComment=57;
	public static final int
		RULE_program = 0, RULE_statementList = 1, RULE_declarationList = 2, RULE_statement = 3, 
		RULE_assignmentStatement = 4, RULE_compoundStatement = 5, RULE_blockItemList = 6, 
		RULE_blockItem = 7, RULE_declaration = 8, RULE_expressionStatement = 9, 
		RULE_selectionStatement = 10, RULE_iterationStatement = 11, RULE_jumpStatement = 12, 
		RULE_unaryOperator = 13, RULE_primaryExpression = 14, RULE_expression = 15, 
		RULE_addop = 16, RULE_mulop = 17, RULE_conditionalExpression = 18, RULE_assignmentExpression = 19, 
		RULE_assignmentOperator = 20, RULE_constantExpression = 21, RULE_shiftExpression = 22, 
		RULE_relationalExpression = 23, RULE_equalityExpression = 24, RULE_andExpression = 25, 
		RULE_exclusiveOrExpression = 26, RULE_inclusiveOrExpression = 27, RULE_logicalAndExpression = 28, 
		RULE_logicalOrExpression = 29, RULE_typeSpecifier = 30, RULE_functionDefinition = 31, 
		RULE_paramList = 32, RULE_param = 33;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "statementList", "declarationList", "statement", "assignmentStatement", 
			"compoundStatement", "blockItemList", "blockItem", "declaration", "expressionStatement", 
			"selectionStatement", "iterationStatement", "jumpStatement", "unaryOperator", 
			"primaryExpression", "expression", "addop", "mulop", "conditionalExpression", 
			"assignmentExpression", "assignmentOperator", "constantExpression", "shiftExpression", 
			"relationalExpression", "equalityExpression", "andExpression", "exclusiveOrExpression", 
			"inclusiveOrExpression", "logicalAndExpression", "logicalOrExpression", 
			"typeSpecifier", "functionDefinition", "paramList", "param"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'var'", "'=='", "'!='", "'break'", "'continue'", "'else'", "'if'", 
			"'int'", "'boolean'", "'string'", "'rune'", "'array'", "'return'", "'void'", 
			"'while'", "'function'", "'true'", "'false'", "'('", "')'", "'['", "']'", 
			"'{'", "'}'", "'<'", "'<='", "'>'", "'>='", "'<<'", "'>>'", "'+'", "'++'", 
			"'-'", "'--'", "'*'", "'/'", "'%'", "'&'", "'|'", "'&&'", "'||'", "'^'", 
			"'!'", "'~'", "'?'", "':'", "';'", "','", "'='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, "Break", "Continue", "Else", "If", "Int", "Bool", 
			"StringType", "Rune", "Array", "Return", "Void", "While", "Func", "True", 
			"False", "Paren", "Thesis", "LeftBracket", "RightBracket", "LeftBrace", 
			"RightBrace", "Less", "LessEqual", "Greater", "GreaterEqual", "LeftShift", 
			"RightShift", "Plus", "PlusPlus", "Minus", "MinusMinus", "Star", "Div", 
			"Mod", "And", "Or", "AndAnd", "OrOr", "Caret", "Not", "Tilde", "Question", 
			"Colon", "Semi", "Comma", "Assign", "Identifier", "Constant", "DigitSequence", 
			"StringLiteral", "Whitespace", "Newline", "BlockComment", "LineComment"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Cminus.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public CminusParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ProgramContext extends ParserRuleContext {
		public StatementListContext statementList() {
			return getRuleContext(StatementListContext.class,0);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(68);
			statementList(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementListContext extends ParserRuleContext {
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public StatementListContext statementList() {
			return getRuleContext(StatementListContext.class,0);
		}
		public StatementListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statementList; }
	}

	public final StatementListContext statementList() throws RecognitionException {
		return statementList(0);
	}

	private StatementListContext statementList(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		StatementListContext _localctx = new StatementListContext(_ctx, _parentState);
		StatementListContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_statementList, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(71);
			statement();
			}
			_ctx.stop = _input.LT(-1);
			setState(77);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new StatementListContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_statementList);
					setState(73);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(74);
					statement();
					}
					} 
				}
				setState(79);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class DeclarationListContext extends ParserRuleContext {
		public DeclarationContext declaration() {
			return getRuleContext(DeclarationContext.class,0);
		}
		public DeclarationListContext declarationList() {
			return getRuleContext(DeclarationListContext.class,0);
		}
		public DeclarationListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declarationList; }
	}

	public final DeclarationListContext declarationList() throws RecognitionException {
		return declarationList(0);
	}

	private DeclarationListContext declarationList(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		DeclarationListContext _localctx = new DeclarationListContext(_ctx, _parentState);
		DeclarationListContext _prevctx = _localctx;
		int _startState = 4;
		enterRecursionRule(_localctx, 4, RULE_declarationList, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(81);
			declaration();
			}
			_ctx.stop = _input.LT(-1);
			setState(87);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new DeclarationListContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_declarationList);
					setState(83);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(84);
					declaration();
					}
					} 
				}
				setState(89);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class StatementContext extends ParserRuleContext {
		public CompoundStatementContext compoundStatement() {
			return getRuleContext(CompoundStatementContext.class,0);
		}
		public ExpressionStatementContext expressionStatement() {
			return getRuleContext(ExpressionStatementContext.class,0);
		}
		public SelectionStatementContext selectionStatement() {
			return getRuleContext(SelectionStatementContext.class,0);
		}
		public IterationStatementContext iterationStatement() {
			return getRuleContext(IterationStatementContext.class,0);
		}
		public JumpStatementContext jumpStatement() {
			return getRuleContext(JumpStatementContext.class,0);
		}
		public AssignmentStatementContext assignmentStatement() {
			return getRuleContext(AssignmentStatementContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_statement);
		try {
			setState(96);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LeftBrace:
				enterOuterAlt(_localctx, 1);
				{
				setState(90);
				compoundStatement();
				}
				break;
			case Paren:
			case Minus:
			case Semi:
			case Constant:
				enterOuterAlt(_localctx, 2);
				{
				setState(91);
				expressionStatement();
				}
				break;
			case If:
				enterOuterAlt(_localctx, 3);
				{
				setState(92);
				selectionStatement();
				}
				break;
			case While:
				enterOuterAlt(_localctx, 4);
				{
				setState(93);
				iterationStatement();
				}
				break;
			case Break:
			case Continue:
			case Return:
				enterOuterAlt(_localctx, 5);
				{
				setState(94);
				jumpStatement();
				}
				break;
			case Identifier:
				enterOuterAlt(_localctx, 6);
				{
				setState(95);
				assignmentStatement();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AssignmentStatementContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(CminusParser.Identifier, 0); }
		public TerminalNode Assign() { return getToken(CminusParser.Assign, 0); }
		public ExpressionStatementContext expressionStatement() {
			return getRuleContext(ExpressionStatementContext.class,0);
		}
		public AssignmentStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignmentStatement; }
	}

	public final AssignmentStatementContext assignmentStatement() throws RecognitionException {
		AssignmentStatementContext _localctx = new AssignmentStatementContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_assignmentStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(98);
			match(Identifier);
			setState(99);
			match(Assign);
			setState(100);
			expressionStatement();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CompoundStatementContext extends ParserRuleContext {
		public TerminalNode LeftBrace() { return getToken(CminusParser.LeftBrace, 0); }
		public TerminalNode RightBrace() { return getToken(CminusParser.RightBrace, 0); }
		public BlockItemListContext blockItemList() {
			return getRuleContext(BlockItemListContext.class,0);
		}
		public CompoundStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_compoundStatement; }
	}

	public final CompoundStatementContext compoundStatement() throws RecognitionException {
		CompoundStatementContext _localctx = new CompoundStatementContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_compoundStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(102);
			match(LeftBrace);
			setState(104);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << Break) | (1L << Continue) | (1L << If) | (1L << Return) | (1L << While) | (1L << Paren) | (1L << LeftBrace) | (1L << Minus) | (1L << Semi) | (1L << Identifier) | (1L << Constant))) != 0)) {
				{
				setState(103);
				blockItemList(0);
				}
			}

			setState(106);
			match(RightBrace);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BlockItemListContext extends ParserRuleContext {
		public BlockItemContext blockItem() {
			return getRuleContext(BlockItemContext.class,0);
		}
		public BlockItemListContext blockItemList() {
			return getRuleContext(BlockItemListContext.class,0);
		}
		public BlockItemListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_blockItemList; }
	}

	public final BlockItemListContext blockItemList() throws RecognitionException {
		return blockItemList(0);
	}

	private BlockItemListContext blockItemList(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		BlockItemListContext _localctx = new BlockItemListContext(_ctx, _parentState);
		BlockItemListContext _prevctx = _localctx;
		int _startState = 12;
		enterRecursionRule(_localctx, 12, RULE_blockItemList, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(109);
			blockItem();
			}
			_ctx.stop = _input.LT(-1);
			setState(115);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new BlockItemListContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_blockItemList);
					setState(111);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(112);
					blockItem();
					}
					} 
				}
				setState(117);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class BlockItemContext extends ParserRuleContext {
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public DeclarationContext declaration() {
			return getRuleContext(DeclarationContext.class,0);
		}
		public BlockItemContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_blockItem; }
	}

	public final BlockItemContext blockItem() throws RecognitionException {
		BlockItemContext _localctx = new BlockItemContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_blockItem);
		try {
			setState(120);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Break:
			case Continue:
			case If:
			case Return:
			case While:
			case Paren:
			case LeftBrace:
			case Minus:
			case Semi:
			case Identifier:
			case Constant:
				enterOuterAlt(_localctx, 1);
				{
				setState(118);
				statement();
				}
				break;
			case T__0:
				enterOuterAlt(_localctx, 2);
				{
				setState(119);
				declaration();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DeclarationContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(CminusParser.Identifier, 0); }
		public TerminalNode Assign() { return getToken(CminusParser.Assign, 0); }
		public TypeSpecifierContext typeSpecifier() {
			return getRuleContext(TypeSpecifierContext.class,0);
		}
		public TerminalNode Paren() { return getToken(CminusParser.Paren, 0); }
		public TerminalNode Constant() { return getToken(CminusParser.Constant, 0); }
		public TerminalNode Thesis() { return getToken(CminusParser.Thesis, 0); }
		public TerminalNode Semi() { return getToken(CminusParser.Semi, 0); }
		public DeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaration; }
	}

	public final DeclarationContext declaration() throws RecognitionException {
		DeclarationContext _localctx = new DeclarationContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_declaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(122);
			match(T__0);
			setState(123);
			match(Identifier);
			setState(124);
			match(Assign);
			setState(125);
			typeSpecifier();
			setState(126);
			match(Paren);
			setState(127);
			match(Constant);
			setState(128);
			match(Thesis);
			setState(129);
			match(Semi);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionStatementContext extends ParserRuleContext {
		public TerminalNode Semi() { return getToken(CminusParser.Semi, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ExpressionStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionStatement; }
	}

	public final ExpressionStatementContext expressionStatement() throws RecognitionException {
		ExpressionStatementContext _localctx = new ExpressionStatementContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_expressionStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(132);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Paren) | (1L << Minus) | (1L << Constant))) != 0)) {
				{
				setState(131);
				expression(0);
				}
			}

			setState(134);
			match(Semi);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SelectionStatementContext extends ParserRuleContext {
		public TerminalNode If() { return getToken(CminusParser.If, 0); }
		public TerminalNode Paren() { return getToken(CminusParser.Paren, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode Thesis() { return getToken(CminusParser.Thesis, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public TerminalNode Else() { return getToken(CminusParser.Else, 0); }
		public SelectionStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selectionStatement; }
	}

	public final SelectionStatementContext selectionStatement() throws RecognitionException {
		SelectionStatementContext _localctx = new SelectionStatementContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_selectionStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(136);
			match(If);
			setState(137);
			match(Paren);
			setState(138);
			expression(0);
			setState(139);
			match(Thesis);
			setState(140);
			statement();
			setState(143);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				{
				setState(141);
				match(Else);
				setState(142);
				statement();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IterationStatementContext extends ParserRuleContext {
		public TerminalNode While() { return getToken(CminusParser.While, 0); }
		public TerminalNode Paren() { return getToken(CminusParser.Paren, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode Thesis() { return getToken(CminusParser.Thesis, 0); }
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public IterationStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_iterationStatement; }
	}

	public final IterationStatementContext iterationStatement() throws RecognitionException {
		IterationStatementContext _localctx = new IterationStatementContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_iterationStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(145);
			match(While);
			setState(146);
			match(Paren);
			setState(147);
			expression(0);
			setState(148);
			match(Thesis);
			setState(149);
			statement();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class JumpStatementContext extends ParserRuleContext {
		public TerminalNode Continue() { return getToken(CminusParser.Continue, 0); }
		public TerminalNode Semi() { return getToken(CminusParser.Semi, 0); }
		public TerminalNode Break() { return getToken(CminusParser.Break, 0); }
		public TerminalNode Return() { return getToken(CminusParser.Return, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public JumpStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_jumpStatement; }
	}

	public final JumpStatementContext jumpStatement() throws RecognitionException {
		JumpStatementContext _localctx = new JumpStatementContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_jumpStatement);
		int _la;
		try {
			setState(160);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Continue:
				enterOuterAlt(_localctx, 1);
				{
				setState(151);
				match(Continue);
				setState(152);
				match(Semi);
				}
				break;
			case Break:
				enterOuterAlt(_localctx, 2);
				{
				setState(153);
				match(Break);
				setState(154);
				match(Semi);
				}
				break;
			case Return:
				enterOuterAlt(_localctx, 3);
				{
				setState(155);
				match(Return);
				setState(157);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Paren) | (1L << Minus) | (1L << Constant))) != 0)) {
					{
					setState(156);
					expression(0);
					}
				}

				setState(159);
				match(Semi);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class UnaryOperatorContext extends ParserRuleContext {
		public TerminalNode And() { return getToken(CminusParser.And, 0); }
		public TerminalNode Star() { return getToken(CminusParser.Star, 0); }
		public TerminalNode Plus() { return getToken(CminusParser.Plus, 0); }
		public TerminalNode Minus() { return getToken(CminusParser.Minus, 0); }
		public TerminalNode Tilde() { return getToken(CminusParser.Tilde, 0); }
		public TerminalNode Not() { return getToken(CminusParser.Not, 0); }
		public UnaryOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unaryOperator; }
	}

	public final UnaryOperatorContext unaryOperator() throws RecognitionException {
		UnaryOperatorContext _localctx = new UnaryOperatorContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_unaryOperator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(162);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Plus) | (1L << Minus) | (1L << Star) | (1L << And) | (1L << Not) | (1L << Tilde))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PrimaryExpressionContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(CminusParser.Identifier, 0); }
		public TerminalNode Constant() { return getToken(CminusParser.Constant, 0); }
		public List<TerminalNode> StringLiteral() { return getTokens(CminusParser.StringLiteral); }
		public TerminalNode StringLiteral(int i) {
			return getToken(CminusParser.StringLiteral, i);
		}
		public TerminalNode Paren() { return getToken(CminusParser.Paren, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode Thesis() { return getToken(CminusParser.Thesis, 0); }
		public PrimaryExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primaryExpression; }
	}

	public final PrimaryExpressionContext primaryExpression() throws RecognitionException {
		PrimaryExpressionContext _localctx = new PrimaryExpressionContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_primaryExpression);
		int _la;
		try {
			setState(175);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(164);
				match(Identifier);
				}
				break;
			case Constant:
				enterOuterAlt(_localctx, 2);
				{
				setState(165);
				match(Constant);
				}
				break;
			case StringLiteral:
				enterOuterAlt(_localctx, 3);
				{
				setState(167); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(166);
					match(StringLiteral);
					}
					}
					setState(169); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==StringLiteral );
				}
				break;
			case Paren:
				enterOuterAlt(_localctx, 4);
				{
				setState(171);
				match(Paren);
				setState(172);
				expression(0);
				setState(173);
				match(Thesis);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionContext extends ParserRuleContext {
		public TerminalNode Minus() { return getToken(CminusParser.Minus, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode Paren() { return getToken(CminusParser.Paren, 0); }
		public TerminalNode Thesis() { return getToken(CminusParser.Thesis, 0); }
		public TerminalNode Constant() { return getToken(CminusParser.Constant, 0); }
		public MulopContext mulop() {
			return getRuleContext(MulopContext.class,0);
		}
		public AddopContext addop() {
			return getRuleContext(AddopContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 30;
		enterRecursionRule(_localctx, 30, RULE_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(185);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Minus:
				{
				setState(178);
				match(Minus);
				setState(179);
				expression(5);
				}
				break;
			case Paren:
				{
				setState(180);
				match(Paren);
				setState(181);
				expression(0);
				setState(182);
				match(Thesis);
				}
				break;
			case Constant:
				{
				setState(184);
				match(Constant);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(197);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(195);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
					case 1:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(187);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(188);
						mulop();
						setState(189);
						expression(5);
						}
						break;
					case 2:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(191);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(192);
						addop();
						setState(193);
						expression(4);
						}
						break;
					}
					} 
				}
				setState(199);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class AddopContext extends ParserRuleContext {
		public TerminalNode Plus() { return getToken(CminusParser.Plus, 0); }
		public TerminalNode Minus() { return getToken(CminusParser.Minus, 0); }
		public AddopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_addop; }
	}

	public final AddopContext addop() throws RecognitionException {
		AddopContext _localctx = new AddopContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_addop);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(200);
			_la = _input.LA(1);
			if ( !(_la==Plus || _la==Minus) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MulopContext extends ParserRuleContext {
		public TerminalNode Star() { return getToken(CminusParser.Star, 0); }
		public TerminalNode Div() { return getToken(CminusParser.Div, 0); }
		public TerminalNode Mod() { return getToken(CminusParser.Mod, 0); }
		public MulopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mulop; }
	}

	public final MulopContext mulop() throws RecognitionException {
		MulopContext _localctx = new MulopContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_mulop);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(202);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Star) | (1L << Div) | (1L << Mod))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConditionalExpressionContext extends ParserRuleContext {
		public LogicalOrExpressionContext logicalOrExpression() {
			return getRuleContext(LogicalOrExpressionContext.class,0);
		}
		public TerminalNode Question() { return getToken(CminusParser.Question, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode Colon() { return getToken(CminusParser.Colon, 0); }
		public ConditionalExpressionContext conditionalExpression() {
			return getRuleContext(ConditionalExpressionContext.class,0);
		}
		public ConditionalExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditionalExpression; }
	}

	public final ConditionalExpressionContext conditionalExpression() throws RecognitionException {
		ConditionalExpressionContext _localctx = new ConditionalExpressionContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_conditionalExpression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(204);
			logicalOrExpression(0);
			setState(210);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Question) {
				{
				setState(205);
				match(Question);
				setState(206);
				expression(0);
				setState(207);
				match(Colon);
				setState(208);
				conditionalExpression();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AssignmentExpressionContext extends ParserRuleContext {
		public ConditionalExpressionContext conditionalExpression() {
			return getRuleContext(ConditionalExpressionContext.class,0);
		}
		public TerminalNode DigitSequence() { return getToken(CminusParser.DigitSequence, 0); }
		public AssignmentExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignmentExpression; }
	}

	public final AssignmentExpressionContext assignmentExpression() throws RecognitionException {
		AssignmentExpressionContext _localctx = new AssignmentExpressionContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_assignmentExpression);
		try {
			setState(214);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Paren:
			case Minus:
			case Constant:
				enterOuterAlt(_localctx, 1);
				{
				setState(212);
				conditionalExpression();
				}
				break;
			case DigitSequence:
				enterOuterAlt(_localctx, 2);
				{
				setState(213);
				match(DigitSequence);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AssignmentOperatorContext extends ParserRuleContext {
		public TerminalNode Assign() { return getToken(CminusParser.Assign, 0); }
		public AssignmentOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignmentOperator; }
	}

	public final AssignmentOperatorContext assignmentOperator() throws RecognitionException {
		AssignmentOperatorContext _localctx = new AssignmentOperatorContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_assignmentOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(216);
			match(Assign);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConstantExpressionContext extends ParserRuleContext {
		public ConditionalExpressionContext conditionalExpression() {
			return getRuleContext(ConditionalExpressionContext.class,0);
		}
		public ConstantExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constantExpression; }
	}

	public final ConstantExpressionContext constantExpression() throws RecognitionException {
		ConstantExpressionContext _localctx = new ConstantExpressionContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_constantExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(218);
			conditionalExpression();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ShiftExpressionContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ShiftExpressionContext shiftExpression() {
			return getRuleContext(ShiftExpressionContext.class,0);
		}
		public TerminalNode LeftShift() { return getToken(CminusParser.LeftShift, 0); }
		public TerminalNode RightShift() { return getToken(CminusParser.RightShift, 0); }
		public ShiftExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_shiftExpression; }
	}

	public final ShiftExpressionContext shiftExpression() throws RecognitionException {
		return shiftExpression(0);
	}

	private ShiftExpressionContext shiftExpression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ShiftExpressionContext _localctx = new ShiftExpressionContext(_ctx, _parentState);
		ShiftExpressionContext _prevctx = _localctx;
		int _startState = 44;
		enterRecursionRule(_localctx, 44, RULE_shiftExpression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(221);
			expression(0);
			}
			_ctx.stop = _input.LT(-1);
			setState(231);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(229);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
					case 1:
						{
						_localctx = new ShiftExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_shiftExpression);
						setState(223);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(224);
						match(LeftShift);
						setState(225);
						expression(0);
						}
						break;
					case 2:
						{
						_localctx = new ShiftExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_shiftExpression);
						setState(226);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(227);
						match(RightShift);
						setState(228);
						expression(0);
						}
						break;
					}
					} 
				}
				setState(233);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class RelationalExpressionContext extends ParserRuleContext {
		public ShiftExpressionContext shiftExpression() {
			return getRuleContext(ShiftExpressionContext.class,0);
		}
		public RelationalExpressionContext relationalExpression() {
			return getRuleContext(RelationalExpressionContext.class,0);
		}
		public TerminalNode Less() { return getToken(CminusParser.Less, 0); }
		public TerminalNode Greater() { return getToken(CminusParser.Greater, 0); }
		public TerminalNode LessEqual() { return getToken(CminusParser.LessEqual, 0); }
		public TerminalNode GreaterEqual() { return getToken(CminusParser.GreaterEqual, 0); }
		public RelationalExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationalExpression; }
	}

	public final RelationalExpressionContext relationalExpression() throws RecognitionException {
		return relationalExpression(0);
	}

	private RelationalExpressionContext relationalExpression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		RelationalExpressionContext _localctx = new RelationalExpressionContext(_ctx, _parentState);
		RelationalExpressionContext _prevctx = _localctx;
		int _startState = 46;
		enterRecursionRule(_localctx, 46, RULE_relationalExpression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(235);
			shiftExpression(0);
			}
			_ctx.stop = _input.LT(-1);
			setState(251);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(249);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
					case 1:
						{
						_localctx = new RelationalExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_relationalExpression);
						setState(237);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(238);
						match(Less);
						setState(239);
						shiftExpression(0);
						}
						break;
					case 2:
						{
						_localctx = new RelationalExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_relationalExpression);
						setState(240);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(241);
						match(Greater);
						setState(242);
						shiftExpression(0);
						}
						break;
					case 3:
						{
						_localctx = new RelationalExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_relationalExpression);
						setState(243);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(244);
						match(LessEqual);
						setState(245);
						shiftExpression(0);
						}
						break;
					case 4:
						{
						_localctx = new RelationalExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_relationalExpression);
						setState(246);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(247);
						match(GreaterEqual);
						setState(248);
						shiftExpression(0);
						}
						break;
					}
					} 
				}
				setState(253);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class EqualityExpressionContext extends ParserRuleContext {
		public RelationalExpressionContext relationalExpression() {
			return getRuleContext(RelationalExpressionContext.class,0);
		}
		public EqualityExpressionContext equalityExpression() {
			return getRuleContext(EqualityExpressionContext.class,0);
		}
		public EqualityExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_equalityExpression; }
	}

	public final EqualityExpressionContext equalityExpression() throws RecognitionException {
		return equalityExpression(0);
	}

	private EqualityExpressionContext equalityExpression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		EqualityExpressionContext _localctx = new EqualityExpressionContext(_ctx, _parentState);
		EqualityExpressionContext _prevctx = _localctx;
		int _startState = 48;
		enterRecursionRule(_localctx, 48, RULE_equalityExpression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(255);
			relationalExpression(0);
			}
			_ctx.stop = _input.LT(-1);
			setState(265);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,22,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(263);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
					case 1:
						{
						_localctx = new EqualityExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_equalityExpression);
						setState(257);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(258);
						match(T__1);
						setState(259);
						relationalExpression(0);
						}
						break;
					case 2:
						{
						_localctx = new EqualityExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_equalityExpression);
						setState(260);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(261);
						match(T__2);
						setState(262);
						relationalExpression(0);
						}
						break;
					}
					} 
				}
				setState(267);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,22,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class AndExpressionContext extends ParserRuleContext {
		public EqualityExpressionContext equalityExpression() {
			return getRuleContext(EqualityExpressionContext.class,0);
		}
		public AndExpressionContext andExpression() {
			return getRuleContext(AndExpressionContext.class,0);
		}
		public TerminalNode And() { return getToken(CminusParser.And, 0); }
		public AndExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_andExpression; }
	}

	public final AndExpressionContext andExpression() throws RecognitionException {
		return andExpression(0);
	}

	private AndExpressionContext andExpression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		AndExpressionContext _localctx = new AndExpressionContext(_ctx, _parentState);
		AndExpressionContext _prevctx = _localctx;
		int _startState = 50;
		enterRecursionRule(_localctx, 50, RULE_andExpression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(269);
			equalityExpression(0);
			}
			_ctx.stop = _input.LT(-1);
			setState(276);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,23,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new AndExpressionContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_andExpression);
					setState(271);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(272);
					match(And);
					setState(273);
					equalityExpression(0);
					}
					} 
				}
				setState(278);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,23,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ExclusiveOrExpressionContext extends ParserRuleContext {
		public AndExpressionContext andExpression() {
			return getRuleContext(AndExpressionContext.class,0);
		}
		public ExclusiveOrExpressionContext exclusiveOrExpression() {
			return getRuleContext(ExclusiveOrExpressionContext.class,0);
		}
		public TerminalNode Caret() { return getToken(CminusParser.Caret, 0); }
		public ExclusiveOrExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exclusiveOrExpression; }
	}

	public final ExclusiveOrExpressionContext exclusiveOrExpression() throws RecognitionException {
		return exclusiveOrExpression(0);
	}

	private ExclusiveOrExpressionContext exclusiveOrExpression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExclusiveOrExpressionContext _localctx = new ExclusiveOrExpressionContext(_ctx, _parentState);
		ExclusiveOrExpressionContext _prevctx = _localctx;
		int _startState = 52;
		enterRecursionRule(_localctx, 52, RULE_exclusiveOrExpression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(280);
			andExpression(0);
			}
			_ctx.stop = _input.LT(-1);
			setState(287);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,24,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ExclusiveOrExpressionContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_exclusiveOrExpression);
					setState(282);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(283);
					match(Caret);
					setState(284);
					andExpression(0);
					}
					} 
				}
				setState(289);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,24,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class InclusiveOrExpressionContext extends ParserRuleContext {
		public ExclusiveOrExpressionContext exclusiveOrExpression() {
			return getRuleContext(ExclusiveOrExpressionContext.class,0);
		}
		public InclusiveOrExpressionContext inclusiveOrExpression() {
			return getRuleContext(InclusiveOrExpressionContext.class,0);
		}
		public TerminalNode Or() { return getToken(CminusParser.Or, 0); }
		public InclusiveOrExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_inclusiveOrExpression; }
	}

	public final InclusiveOrExpressionContext inclusiveOrExpression() throws RecognitionException {
		return inclusiveOrExpression(0);
	}

	private InclusiveOrExpressionContext inclusiveOrExpression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		InclusiveOrExpressionContext _localctx = new InclusiveOrExpressionContext(_ctx, _parentState);
		InclusiveOrExpressionContext _prevctx = _localctx;
		int _startState = 54;
		enterRecursionRule(_localctx, 54, RULE_inclusiveOrExpression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(291);
			exclusiveOrExpression(0);
			}
			_ctx.stop = _input.LT(-1);
			setState(298);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,25,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new InclusiveOrExpressionContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_inclusiveOrExpression);
					setState(293);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(294);
					match(Or);
					setState(295);
					exclusiveOrExpression(0);
					}
					} 
				}
				setState(300);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,25,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class LogicalAndExpressionContext extends ParserRuleContext {
		public InclusiveOrExpressionContext inclusiveOrExpression() {
			return getRuleContext(InclusiveOrExpressionContext.class,0);
		}
		public LogicalAndExpressionContext logicalAndExpression() {
			return getRuleContext(LogicalAndExpressionContext.class,0);
		}
		public TerminalNode AndAnd() { return getToken(CminusParser.AndAnd, 0); }
		public LogicalAndExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_logicalAndExpression; }
	}

	public final LogicalAndExpressionContext logicalAndExpression() throws RecognitionException {
		return logicalAndExpression(0);
	}

	private LogicalAndExpressionContext logicalAndExpression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		LogicalAndExpressionContext _localctx = new LogicalAndExpressionContext(_ctx, _parentState);
		LogicalAndExpressionContext _prevctx = _localctx;
		int _startState = 56;
		enterRecursionRule(_localctx, 56, RULE_logicalAndExpression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(302);
			inclusiveOrExpression(0);
			}
			_ctx.stop = _input.LT(-1);
			setState(309);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new LogicalAndExpressionContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_logicalAndExpression);
					setState(304);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(305);
					match(AndAnd);
					setState(306);
					inclusiveOrExpression(0);
					}
					} 
				}
				setState(311);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class LogicalOrExpressionContext extends ParserRuleContext {
		public LogicalAndExpressionContext logicalAndExpression() {
			return getRuleContext(LogicalAndExpressionContext.class,0);
		}
		public LogicalOrExpressionContext logicalOrExpression() {
			return getRuleContext(LogicalOrExpressionContext.class,0);
		}
		public TerminalNode OrOr() { return getToken(CminusParser.OrOr, 0); }
		public LogicalOrExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_logicalOrExpression; }
	}

	public final LogicalOrExpressionContext logicalOrExpression() throws RecognitionException {
		return logicalOrExpression(0);
	}

	private LogicalOrExpressionContext logicalOrExpression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		LogicalOrExpressionContext _localctx = new LogicalOrExpressionContext(_ctx, _parentState);
		LogicalOrExpressionContext _prevctx = _localctx;
		int _startState = 58;
		enterRecursionRule(_localctx, 58, RULE_logicalOrExpression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(313);
			logicalAndExpression(0);
			}
			_ctx.stop = _input.LT(-1);
			setState(320);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new LogicalOrExpressionContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_logicalOrExpression);
					setState(315);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(316);
					match(OrOr);
					setState(317);
					logicalAndExpression(0);
					}
					} 
				}
				setState(322);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class TypeSpecifierContext extends ParserRuleContext {
		public TerminalNode Void() { return getToken(CminusParser.Void, 0); }
		public TerminalNode Rune() { return getToken(CminusParser.Rune, 0); }
		public TerminalNode Int() { return getToken(CminusParser.Int, 0); }
		public TerminalNode StringType() { return getToken(CminusParser.StringType, 0); }
		public TerminalNode Bool() { return getToken(CminusParser.Bool, 0); }
		public TypeSpecifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeSpecifier; }
	}

	public final TypeSpecifierContext typeSpecifier() throws RecognitionException {
		TypeSpecifierContext _localctx = new TypeSpecifierContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_typeSpecifier);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(323);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Int) | (1L << Bool) | (1L << StringType) | (1L << Rune) | (1L << Void))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionDefinitionContext extends ParserRuleContext {
		public TerminalNode Func() { return getToken(CminusParser.Func, 0); }
		public TerminalNode Identifier() { return getToken(CminusParser.Identifier, 0); }
		public TerminalNode Paren() { return getToken(CminusParser.Paren, 0); }
		public ParamListContext paramList() {
			return getRuleContext(ParamListContext.class,0);
		}
		public TerminalNode Thesis() { return getToken(CminusParser.Thesis, 0); }
		public TypeSpecifierContext typeSpecifier() {
			return getRuleContext(TypeSpecifierContext.class,0);
		}
		public CompoundStatementContext compoundStatement() {
			return getRuleContext(CompoundStatementContext.class,0);
		}
		public FunctionDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionDefinition; }
	}

	public final FunctionDefinitionContext functionDefinition() throws RecognitionException {
		FunctionDefinitionContext _localctx = new FunctionDefinitionContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_functionDefinition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(325);
			match(Func);
			setState(326);
			match(Identifier);
			setState(327);
			match(Paren);
			setState(328);
			paramList(0);
			setState(329);
			match(Thesis);
			setState(330);
			typeSpecifier();
			setState(331);
			compoundStatement();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParamListContext extends ParserRuleContext {
		public ParamContext param() {
			return getRuleContext(ParamContext.class,0);
		}
		public ParamListContext paramList() {
			return getRuleContext(ParamListContext.class,0);
		}
		public ParamListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paramList; }
	}

	public final ParamListContext paramList() throws RecognitionException {
		return paramList(0);
	}

	private ParamListContext paramList(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ParamListContext _localctx = new ParamListContext(_ctx, _parentState);
		ParamListContext _prevctx = _localctx;
		int _startState = 64;
		enterRecursionRule(_localctx, 64, RULE_paramList, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(334);
			param();
			}
			_ctx.stop = _input.LT(-1);
			setState(340);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ParamListContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_paramList);
					setState(336);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(337);
					param();
					}
					} 
				}
				setState(342);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ParamContext extends ParserRuleContext {
		public TypeSpecifierContext typeSpecifier() {
			return getRuleContext(TypeSpecifierContext.class,0);
		}
		public TerminalNode Identifier() { return getToken(CminusParser.Identifier, 0); }
		public TerminalNode Comma() { return getToken(CminusParser.Comma, 0); }
		public ParamContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_param; }
	}

	public final ParamContext param() throws RecognitionException {
		ParamContext _localctx = new ParamContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_param);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(344);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Comma) {
				{
				setState(343);
				match(Comma);
				}
			}

			setState(346);
			typeSpecifier();
			setState(347);
			match(Identifier);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 1:
			return statementList_sempred((StatementListContext)_localctx, predIndex);
		case 2:
			return declarationList_sempred((DeclarationListContext)_localctx, predIndex);
		case 6:
			return blockItemList_sempred((BlockItemListContext)_localctx, predIndex);
		case 15:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		case 22:
			return shiftExpression_sempred((ShiftExpressionContext)_localctx, predIndex);
		case 23:
			return relationalExpression_sempred((RelationalExpressionContext)_localctx, predIndex);
		case 24:
			return equalityExpression_sempred((EqualityExpressionContext)_localctx, predIndex);
		case 25:
			return andExpression_sempred((AndExpressionContext)_localctx, predIndex);
		case 26:
			return exclusiveOrExpression_sempred((ExclusiveOrExpressionContext)_localctx, predIndex);
		case 27:
			return inclusiveOrExpression_sempred((InclusiveOrExpressionContext)_localctx, predIndex);
		case 28:
			return logicalAndExpression_sempred((LogicalAndExpressionContext)_localctx, predIndex);
		case 29:
			return logicalOrExpression_sempred((LogicalOrExpressionContext)_localctx, predIndex);
		case 32:
			return paramList_sempred((ParamListContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean statementList_sempred(StatementListContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean declarationList_sempred(DeclarationListContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean blockItemList_sempred(BlockItemListContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 4);
		case 4:
			return precpred(_ctx, 3);
		}
		return true;
	}
	private boolean shiftExpression_sempred(ShiftExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 5:
			return precpred(_ctx, 2);
		case 6:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean relationalExpression_sempred(RelationalExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 7:
			return precpred(_ctx, 4);
		case 8:
			return precpred(_ctx, 3);
		case 9:
			return precpred(_ctx, 2);
		case 10:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean equalityExpression_sempred(EqualityExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 11:
			return precpred(_ctx, 2);
		case 12:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean andExpression_sempred(AndExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 13:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean exclusiveOrExpression_sempred(ExclusiveOrExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 14:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean inclusiveOrExpression_sempred(InclusiveOrExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 15:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean logicalAndExpression_sempred(LogicalAndExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 16:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean logicalOrExpression_sempred(LogicalOrExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 17:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean paramList_sempred(ParamListContext _localctx, int predIndex) {
		switch (predIndex) {
		case 18:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3;\u0160\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\3\2\3\2\3\3\3\3\3\3\3\3\3\3\7\3N\n\3\f\3\16\3Q\13\3"+
		"\3\4\3\4\3\4\3\4\3\4\7\4X\n\4\f\4\16\4[\13\4\3\5\3\5\3\5\3\5\3\5\3\5\5"+
		"\5c\n\5\3\6\3\6\3\6\3\6\3\7\3\7\5\7k\n\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\7"+
		"\bt\n\b\f\b\16\bw\13\b\3\t\3\t\5\t{\n\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\13\5\13\u0087\n\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\5\f"+
		"\u0092\n\f\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\5\16"+
		"\u00a0\n\16\3\16\5\16\u00a3\n\16\3\17\3\17\3\20\3\20\3\20\6\20\u00aa\n"+
		"\20\r\20\16\20\u00ab\3\20\3\20\3\20\3\20\5\20\u00b2\n\20\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\5\21\u00bc\n\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\7\21\u00c6\n\21\f\21\16\21\u00c9\13\21\3\22\3\22\3\23\3\23"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\5\24\u00d5\n\24\3\25\3\25\5\25\u00d9\n"+
		"\25\3\26\3\26\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\7"+
		"\30\u00e8\n\30\f\30\16\30\u00eb\13\30\3\31\3\31\3\31\3\31\3\31\3\31\3"+
		"\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\7\31\u00fc\n\31\f\31\16\31"+
		"\u00ff\13\31\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\7\32\u010a\n"+
		"\32\f\32\16\32\u010d\13\32\3\33\3\33\3\33\3\33\3\33\3\33\7\33\u0115\n"+
		"\33\f\33\16\33\u0118\13\33\3\34\3\34\3\34\3\34\3\34\3\34\7\34\u0120\n"+
		"\34\f\34\16\34\u0123\13\34\3\35\3\35\3\35\3\35\3\35\3\35\7\35\u012b\n"+
		"\35\f\35\16\35\u012e\13\35\3\36\3\36\3\36\3\36\3\36\3\36\7\36\u0136\n"+
		"\36\f\36\16\36\u0139\13\36\3\37\3\37\3\37\3\37\3\37\3\37\7\37\u0141\n"+
		"\37\f\37\16\37\u0144\13\37\3 \3 \3!\3!\3!\3!\3!\3!\3!\3!\3\"\3\"\3\"\3"+
		"\"\3\"\7\"\u0155\n\"\f\"\16\"\u0158\13\"\3#\5#\u015b\n#\3#\3#\3#\3#\2"+
		"\17\4\6\16 .\60\62\64\668:<B$\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 "+
		"\"$&(*,.\60\62\64\668:<>@BD\2\6\7\2!!##%%((-.\4\2!!##\3\2%\'\4\2\n\r\20"+
		"\20\2\u0165\2F\3\2\2\2\4H\3\2\2\2\6R\3\2\2\2\bb\3\2\2\2\nd\3\2\2\2\fh"+
		"\3\2\2\2\16n\3\2\2\2\20z\3\2\2\2\22|\3\2\2\2\24\u0086\3\2\2\2\26\u008a"+
		"\3\2\2\2\30\u0093\3\2\2\2\32\u00a2\3\2\2\2\34\u00a4\3\2\2\2\36\u00b1\3"+
		"\2\2\2 \u00bb\3\2\2\2\"\u00ca\3\2\2\2$\u00cc\3\2\2\2&\u00ce\3\2\2\2(\u00d8"+
		"\3\2\2\2*\u00da\3\2\2\2,\u00dc\3\2\2\2.\u00de\3\2\2\2\60\u00ec\3\2\2\2"+
		"\62\u0100\3\2\2\2\64\u010e\3\2\2\2\66\u0119\3\2\2\28\u0124\3\2\2\2:\u012f"+
		"\3\2\2\2<\u013a\3\2\2\2>\u0145\3\2\2\2@\u0147\3\2\2\2B\u014f\3\2\2\2D"+
		"\u015a\3\2\2\2FG\5\4\3\2G\3\3\2\2\2HI\b\3\1\2IJ\5\b\5\2JO\3\2\2\2KL\f"+
		"\3\2\2LN\5\b\5\2MK\3\2\2\2NQ\3\2\2\2OM\3\2\2\2OP\3\2\2\2P\5\3\2\2\2QO"+
		"\3\2\2\2RS\b\4\1\2ST\5\22\n\2TY\3\2\2\2UV\f\3\2\2VX\5\22\n\2WU\3\2\2\2"+
		"X[\3\2\2\2YW\3\2\2\2YZ\3\2\2\2Z\7\3\2\2\2[Y\3\2\2\2\\c\5\f\7\2]c\5\24"+
		"\13\2^c\5\26\f\2_c\5\30\r\2`c\5\32\16\2ac\5\n\6\2b\\\3\2\2\2b]\3\2\2\2"+
		"b^\3\2\2\2b_\3\2\2\2b`\3\2\2\2ba\3\2\2\2c\t\3\2\2\2de\7\64\2\2ef\7\63"+
		"\2\2fg\5\24\13\2g\13\3\2\2\2hj\7\31\2\2ik\5\16\b\2ji\3\2\2\2jk\3\2\2\2"+
		"kl\3\2\2\2lm\7\32\2\2m\r\3\2\2\2no\b\b\1\2op\5\20\t\2pu\3\2\2\2qr\f\3"+
		"\2\2rt\5\20\t\2sq\3\2\2\2tw\3\2\2\2us\3\2\2\2uv\3\2\2\2v\17\3\2\2\2wu"+
		"\3\2\2\2x{\5\b\5\2y{\5\22\n\2zx\3\2\2\2zy\3\2\2\2{\21\3\2\2\2|}\7\3\2"+
		"\2}~\7\64\2\2~\177\7\63\2\2\177\u0080\5> \2\u0080\u0081\7\25\2\2\u0081"+
		"\u0082\7\65\2\2\u0082\u0083\7\26\2\2\u0083\u0084\7\61\2\2\u0084\23\3\2"+
		"\2\2\u0085\u0087\5 \21\2\u0086\u0085\3\2\2\2\u0086\u0087\3\2\2\2\u0087"+
		"\u0088\3\2\2\2\u0088\u0089\7\61\2\2\u0089\25\3\2\2\2\u008a\u008b\7\t\2"+
		"\2\u008b\u008c\7\25\2\2\u008c\u008d\5 \21\2\u008d\u008e\7\26\2\2\u008e"+
		"\u0091\5\b\5\2\u008f\u0090\7\b\2\2\u0090\u0092\5\b\5\2\u0091\u008f\3\2"+
		"\2\2\u0091\u0092\3\2\2\2\u0092\27\3\2\2\2\u0093\u0094\7\21\2\2\u0094\u0095"+
		"\7\25\2\2\u0095\u0096\5 \21\2\u0096\u0097\7\26\2\2\u0097\u0098\5\b\5\2"+
		"\u0098\31\3\2\2\2\u0099\u009a\7\7\2\2\u009a\u00a3\7\61\2\2\u009b\u009c"+
		"\7\6\2\2\u009c\u00a3\7\61\2\2\u009d\u009f\7\17\2\2\u009e\u00a0\5 \21\2"+
		"\u009f\u009e\3\2\2\2\u009f\u00a0\3\2\2\2\u00a0\u00a1\3\2\2\2\u00a1\u00a3"+
		"\7\61\2\2\u00a2\u0099\3\2\2\2\u00a2\u009b\3\2\2\2\u00a2\u009d\3\2\2\2"+
		"\u00a3\33\3\2\2\2\u00a4\u00a5\t\2\2\2\u00a5\35\3\2\2\2\u00a6\u00b2\7\64"+
		"\2\2\u00a7\u00b2\7\65\2\2\u00a8\u00aa\7\67\2\2\u00a9\u00a8\3\2\2\2\u00aa"+
		"\u00ab\3\2\2\2\u00ab\u00a9\3\2\2\2\u00ab\u00ac\3\2\2\2\u00ac\u00b2\3\2"+
		"\2\2\u00ad\u00ae\7\25\2\2\u00ae\u00af\5 \21\2\u00af\u00b0\7\26\2\2\u00b0"+
		"\u00b2\3\2\2\2\u00b1\u00a6\3\2\2\2\u00b1\u00a7\3\2\2\2\u00b1\u00a9\3\2"+
		"\2\2\u00b1\u00ad\3\2\2\2\u00b2\37\3\2\2\2\u00b3\u00b4\b\21\1\2\u00b4\u00b5"+
		"\7#\2\2\u00b5\u00bc\5 \21\7\u00b6\u00b7\7\25\2\2\u00b7\u00b8\5 \21\2\u00b8"+
		"\u00b9\7\26\2\2\u00b9\u00bc\3\2\2\2\u00ba\u00bc\7\65\2\2\u00bb\u00b3\3"+
		"\2\2\2\u00bb\u00b6\3\2\2\2\u00bb\u00ba\3\2\2\2\u00bc\u00c7\3\2\2\2\u00bd"+
		"\u00be\f\6\2\2\u00be\u00bf\5$\23\2\u00bf\u00c0\5 \21\7\u00c0\u00c6\3\2"+
		"\2\2\u00c1\u00c2\f\5\2\2\u00c2\u00c3\5\"\22\2\u00c3\u00c4\5 \21\6\u00c4"+
		"\u00c6\3\2\2\2\u00c5\u00bd\3\2\2\2\u00c5\u00c1\3\2\2\2\u00c6\u00c9\3\2"+
		"\2\2\u00c7\u00c5\3\2\2\2\u00c7\u00c8\3\2\2\2\u00c8!\3\2\2\2\u00c9\u00c7"+
		"\3\2\2\2\u00ca\u00cb\t\3\2\2\u00cb#\3\2\2\2\u00cc\u00cd\t\4\2\2\u00cd"+
		"%\3\2\2\2\u00ce\u00d4\5<\37\2\u00cf\u00d0\7/\2\2\u00d0\u00d1\5 \21\2\u00d1"+
		"\u00d2\7\60\2\2\u00d2\u00d3\5&\24\2\u00d3\u00d5\3\2\2\2\u00d4\u00cf\3"+
		"\2\2\2\u00d4\u00d5\3\2\2\2\u00d5\'\3\2\2\2\u00d6\u00d9\5&\24\2\u00d7\u00d9"+
		"\7\66\2\2\u00d8\u00d6\3\2\2\2\u00d8\u00d7\3\2\2\2\u00d9)\3\2\2\2\u00da"+
		"\u00db\7\63\2\2\u00db+\3\2\2\2\u00dc\u00dd\5&\24\2\u00dd-\3\2\2\2\u00de"+
		"\u00df\b\30\1\2\u00df\u00e0\5 \21\2\u00e0\u00e9\3\2\2\2\u00e1\u00e2\f"+
		"\4\2\2\u00e2\u00e3\7\37\2\2\u00e3\u00e8\5 \21\2\u00e4\u00e5\f\3\2\2\u00e5"+
		"\u00e6\7 \2\2\u00e6\u00e8\5 \21\2\u00e7\u00e1\3\2\2\2\u00e7\u00e4\3\2"+
		"\2\2\u00e8\u00eb\3\2\2\2\u00e9\u00e7\3\2\2\2\u00e9\u00ea\3\2\2\2\u00ea"+
		"/\3\2\2\2\u00eb\u00e9\3\2\2\2\u00ec\u00ed\b\31\1\2\u00ed\u00ee\5.\30\2"+
		"\u00ee\u00fd\3\2\2\2\u00ef\u00f0\f\6\2\2\u00f0\u00f1\7\33\2\2\u00f1\u00fc"+
		"\5.\30\2\u00f2\u00f3\f\5\2\2\u00f3\u00f4\7\35\2\2\u00f4\u00fc\5.\30\2"+
		"\u00f5\u00f6\f\4\2\2\u00f6\u00f7\7\34\2\2\u00f7\u00fc\5.\30\2\u00f8\u00f9"+
		"\f\3\2\2\u00f9\u00fa\7\36\2\2\u00fa\u00fc\5.\30\2\u00fb\u00ef\3\2\2\2"+
		"\u00fb\u00f2\3\2\2\2\u00fb\u00f5\3\2\2\2\u00fb\u00f8\3\2\2\2\u00fc\u00ff"+
		"\3\2\2\2\u00fd\u00fb\3\2\2\2\u00fd\u00fe\3\2\2\2\u00fe\61\3\2\2\2\u00ff"+
		"\u00fd\3\2\2\2\u0100\u0101\b\32\1\2\u0101\u0102\5\60\31\2\u0102\u010b"+
		"\3\2\2\2\u0103\u0104\f\4\2\2\u0104\u0105\7\4\2\2\u0105\u010a\5\60\31\2"+
		"\u0106\u0107\f\3\2\2\u0107\u0108\7\5\2\2\u0108\u010a\5\60\31\2\u0109\u0103"+
		"\3\2\2\2\u0109\u0106\3\2\2\2\u010a\u010d\3\2\2\2\u010b\u0109\3\2\2\2\u010b"+
		"\u010c\3\2\2\2\u010c\63\3\2\2\2\u010d\u010b\3\2\2\2\u010e\u010f\b\33\1"+
		"\2\u010f\u0110\5\62\32\2\u0110\u0116\3\2\2\2\u0111\u0112\f\3\2\2\u0112"+
		"\u0113\7(\2\2\u0113\u0115\5\62\32\2\u0114\u0111\3\2\2\2\u0115\u0118\3"+
		"\2\2\2\u0116\u0114\3\2\2\2\u0116\u0117\3\2\2\2\u0117\65\3\2\2\2\u0118"+
		"\u0116\3\2\2\2\u0119\u011a\b\34\1\2\u011a\u011b\5\64\33\2\u011b\u0121"+
		"\3\2\2\2\u011c\u011d\f\3\2\2\u011d\u011e\7,\2\2\u011e\u0120\5\64\33\2"+
		"\u011f\u011c\3\2\2\2\u0120\u0123\3\2\2\2\u0121\u011f\3\2\2\2\u0121\u0122"+
		"\3\2\2\2\u0122\67\3\2\2\2\u0123\u0121\3\2\2\2\u0124\u0125\b\35\1\2\u0125"+
		"\u0126\5\66\34\2\u0126\u012c\3\2\2\2\u0127\u0128\f\3\2\2\u0128\u0129\7"+
		")\2\2\u0129\u012b\5\66\34\2\u012a\u0127\3\2\2\2\u012b\u012e\3\2\2\2\u012c"+
		"\u012a\3\2\2\2\u012c\u012d\3\2\2\2\u012d9\3\2\2\2\u012e\u012c\3\2\2\2"+
		"\u012f\u0130\b\36\1\2\u0130\u0131\58\35\2\u0131\u0137\3\2\2\2\u0132\u0133"+
		"\f\3\2\2\u0133\u0134\7*\2\2\u0134\u0136\58\35\2\u0135\u0132\3\2\2\2\u0136"+
		"\u0139\3\2\2\2\u0137\u0135\3\2\2\2\u0137\u0138\3\2\2\2\u0138;\3\2\2\2"+
		"\u0139\u0137\3\2\2\2\u013a\u013b\b\37\1\2\u013b\u013c\5:\36\2\u013c\u0142"+
		"\3\2\2\2\u013d\u013e\f\3\2\2\u013e\u013f\7+\2\2\u013f\u0141\5:\36\2\u0140"+
		"\u013d\3\2\2\2\u0141\u0144\3\2\2\2\u0142\u0140\3\2\2\2\u0142\u0143\3\2"+
		"\2\2\u0143=\3\2\2\2\u0144\u0142\3\2\2\2\u0145\u0146\t\5\2\2\u0146?\3\2"+
		"\2\2\u0147\u0148\7\22\2\2\u0148\u0149\7\64\2\2\u0149\u014a\7\25\2\2\u014a"+
		"\u014b\5B\"\2\u014b\u014c\7\26\2\2\u014c\u014d\5> \2\u014d\u014e\5\f\7"+
		"\2\u014eA\3\2\2\2\u014f\u0150\b\"\1\2\u0150\u0151\5D#\2\u0151\u0156\3"+
		"\2\2\2\u0152\u0153\f\3\2\2\u0153\u0155\5D#\2\u0154\u0152\3\2\2\2\u0155"+
		"\u0158\3\2\2\2\u0156\u0154\3\2\2\2\u0156\u0157\3\2\2\2\u0157C\3\2\2\2"+
		"\u0158\u0156\3\2\2\2\u0159\u015b\7\62\2\2\u015a\u0159\3\2\2\2\u015a\u015b"+
		"\3\2\2\2\u015b\u015c\3\2\2\2\u015c\u015d\5> \2\u015d\u015e\7\64\2\2\u015e"+
		"E\3\2\2\2 OYbjuz\u0086\u0091\u009f\u00a2\u00ab\u00b1\u00bb\u00c5\u00c7"+
		"\u00d4\u00d8\u00e7\u00e9\u00fb\u00fd\u0109\u010b\u0116\u0121\u012c\u0137"+
		"\u0142\u0156\u015a";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}