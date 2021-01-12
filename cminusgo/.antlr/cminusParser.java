// Generated from /home/kristofer/p/c-minus/cminusgo/cminus.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class cminusParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		Break=1, Continue=2, Else=3, If=4, Var=5, Int=6, Bool=7, StringType=8, 
		Rune=9, Array=10, Return=11, Void=12, While=13, Func=14, Main=15, True=16, 
		False=17, Paren=18, Thesis=19, LeftBracket=20, RightBracket=21, LeftBrace=22, 
		RightBrace=23, EqualEqual=24, NotEqual=25, Less=26, LessEqual=27, Greater=28, 
		GreaterEqual=29, LeftShift=30, RightShift=31, Plus=32, PlusPlus=33, Minus=34, 
		MinusMinus=35, Star=36, Div=37, Mod=38, And=39, Or=40, AndAnd=41, OrOr=42, 
		Caret=43, Not=44, Tilde=45, Question=46, Colon=47, Semi=48, Comma=49, 
		Assign=50, Identifier=51, Constant=52, DigitSequence=53, StringLiteral=54, 
		Whitespace=55, Newline=56, BlockComment=57, LineComment=58;
	public static final int
		RULE_program = 0, RULE_functionList = 1, RULE_mainFunction = 2, RULE_fname = 3, 
		RULE_functionDefinition = 4, RULE_functionCall = 5, RULE_statementList = 6, 
		RULE_declarationList = 7, RULE_statement = 8, RULE_variable = 9, RULE_assignmentStatement = 10, 
		RULE_assignmentOperator = 11, RULE_compoundStatement = 12, RULE_declaration = 13, 
		RULE_selectionStatement = 14, RULE_iterationStatement = 15, RULE_jumpStatement = 16, 
		RULE_expression = 17, RULE_addop = 18, RULE_mulop = 19, RULE_relop = 20, 
		RULE_typeSpecifier = 21, RULE_paramList = 22, RULE_param = 23;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "functionList", "mainFunction", "fname", "functionDefinition", 
			"functionCall", "statementList", "declarationList", "statement", "variable", 
			"assignmentStatement", "assignmentOperator", "compoundStatement", "declaration", 
			"selectionStatement", "iterationStatement", "jumpStatement", "expression", 
			"addop", "mulop", "relop", "typeSpecifier", "paramList", "param"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'break'", "'continue'", "'else'", "'if'", "'var'", "'int'", "'boolean'", 
			"'string'", "'rune'", "'array'", "'return'", "'void'", "'while'", "'function'", 
			"'main'", "'true'", "'false'", "'('", "')'", "'['", "']'", "'{'", "'}'", 
			"'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'<<'", "'>>'", "'+'", 
			"'++'", "'-'", "'--'", "'*'", "'/'", "'%'", "'&'", "'|'", "'&&'", "'||'", 
			"'^'", "'!'", "'~'", "'?'", "':'", "';'", "','", "'='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "Break", "Continue", "Else", "If", "Var", "Int", "Bool", "StringType", 
			"Rune", "Array", "Return", "Void", "While", "Func", "Main", "True", "False", 
			"Paren", "Thesis", "LeftBracket", "RightBracket", "LeftBrace", "RightBrace", 
			"EqualEqual", "NotEqual", "Less", "LessEqual", "Greater", "GreaterEqual", 
			"LeftShift", "RightShift", "Plus", "PlusPlus", "Minus", "MinusMinus", 
			"Star", "Div", "Mod", "And", "Or", "AndAnd", "OrOr", "Caret", "Not", 
			"Tilde", "Question", "Colon", "Semi", "Comma", "Assign", "Identifier", 
			"Constant", "DigitSequence", "StringLiteral", "Whitespace", "Newline", 
			"BlockComment", "LineComment"
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
	public String getGrammarFileName() { return "cminus.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public cminusParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ProgramContext extends ParserRuleContext {
		public MainFunctionContext mainFunction() {
			return getRuleContext(MainFunctionContext.class,0);
		}
		public FunctionListContext functionList() {
			return getRuleContext(FunctionListContext.class,0);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(48);
			mainFunction();
			setState(50);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Func) {
				{
				setState(49);
				functionList(0);
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

	public static class FunctionListContext extends ParserRuleContext {
		public FunctionDefinitionContext functionDefinition() {
			return getRuleContext(FunctionDefinitionContext.class,0);
		}
		public FunctionListContext functionList() {
			return getRuleContext(FunctionListContext.class,0);
		}
		public FunctionListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionList; }
	}

	public final FunctionListContext functionList() throws RecognitionException {
		return functionList(0);
	}

	private FunctionListContext functionList(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		FunctionListContext _localctx = new FunctionListContext(_ctx, _parentState);
		FunctionListContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_functionList, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(53);
			functionDefinition();
			}
			_ctx.stop = _input.LT(-1);
			setState(59);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new FunctionListContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_functionList);
					setState(55);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(56);
					functionDefinition();
					}
					} 
				}
				setState(61);
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

	public static class MainFunctionContext extends ParserRuleContext {
		public TerminalNode Main() { return getToken(cminusParser.Main, 0); }
		public TerminalNode Paren() { return getToken(cminusParser.Paren, 0); }
		public TerminalNode Thesis() { return getToken(cminusParser.Thesis, 0); }
		public CompoundStatementContext compoundStatement() {
			return getRuleContext(CompoundStatementContext.class,0);
		}
		public MainFunctionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mainFunction; }
	}

	public final MainFunctionContext mainFunction() throws RecognitionException {
		MainFunctionContext _localctx = new MainFunctionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_mainFunction);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(62);
			match(Main);
			setState(63);
			match(Paren);
			setState(64);
			match(Thesis);
			setState(65);
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

	public static class FnameContext extends ParserRuleContext {
		public VariableContext variable() {
			return getRuleContext(VariableContext.class,0);
		}
		public FnameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fname; }
	}

	public final FnameContext fname() throws RecognitionException {
		FnameContext _localctx = new FnameContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_fname);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(67);
			variable();
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
		public TerminalNode Func() { return getToken(cminusParser.Func, 0); }
		public FnameContext fname() {
			return getRuleContext(FnameContext.class,0);
		}
		public TerminalNode Paren() { return getToken(cminusParser.Paren, 0); }
		public ParamListContext paramList() {
			return getRuleContext(ParamListContext.class,0);
		}
		public TerminalNode Thesis() { return getToken(cminusParser.Thesis, 0); }
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
		enterRule(_localctx, 8, RULE_functionDefinition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(69);
			match(Func);
			setState(70);
			fname();
			setState(71);
			match(Paren);
			setState(72);
			paramList(0);
			setState(73);
			match(Thesis);
			setState(74);
			typeSpecifier();
			setState(75);
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

	public static class FunctionCallContext extends ParserRuleContext {
		public FnameContext fname() {
			return getRuleContext(FnameContext.class,0);
		}
		public TerminalNode Paren() { return getToken(cminusParser.Paren, 0); }
		public List<VariableContext> variable() {
			return getRuleContexts(VariableContext.class);
		}
		public VariableContext variable(int i) {
			return getRuleContext(VariableContext.class,i);
		}
		public TerminalNode Thesis() { return getToken(cminusParser.Thesis, 0); }
		public List<TerminalNode> Comma() { return getTokens(cminusParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(cminusParser.Comma, i);
		}
		public FunctionCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionCall; }
	}

	public final FunctionCallContext functionCall() throws RecognitionException {
		FunctionCallContext _localctx = new FunctionCallContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_functionCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(77);
			fname();
			setState(78);
			match(Paren);
			setState(79);
			variable();
			setState(84);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==Comma) {
				{
				{
				setState(80);
				match(Comma);
				setState(81);
				variable();
				}
				}
				setState(86);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(87);
			match(Thesis);
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
		int _startState = 12;
		enterRecursionRule(_localctx, 12, RULE_statementList, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(90);
			statement();
			}
			_ctx.stop = _input.LT(-1);
			setState(96);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new StatementListContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_statementList);
					setState(92);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(93);
					statement();
					}
					} 
				}
				setState(98);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
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
		int _startState = 14;
		enterRecursionRule(_localctx, 14, RULE_declarationList, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(100);
			declaration();
			}
			_ctx.stop = _input.LT(-1);
			setState(106);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new DeclarationListContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_declarationList);
					setState(102);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(103);
					declaration();
					}
					} 
				}
				setState(108);
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

	public static class StatementContext extends ParserRuleContext {
		public CompoundStatementContext compoundStatement() {
			return getRuleContext(CompoundStatementContext.class,0);
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
		enterRule(_localctx, 16, RULE_statement);
		try {
			setState(114);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LeftBrace:
				enterOuterAlt(_localctx, 1);
				{
				setState(109);
				compoundStatement();
				}
				break;
			case If:
				enterOuterAlt(_localctx, 2);
				{
				setState(110);
				selectionStatement();
				}
				break;
			case While:
				enterOuterAlt(_localctx, 3);
				{
				setState(111);
				iterationStatement();
				}
				break;
			case Break:
			case Continue:
			case Return:
				enterOuterAlt(_localctx, 4);
				{
				setState(112);
				jumpStatement();
				}
				break;
			case Identifier:
				enterOuterAlt(_localctx, 5);
				{
				setState(113);
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

	public static class VariableContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(cminusParser.Identifier, 0); }
		public VariableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variable; }
	}

	public final VariableContext variable() throws RecognitionException {
		VariableContext _localctx = new VariableContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_variable);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(116);
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

	public static class AssignmentStatementContext extends ParserRuleContext {
		public VariableContext variable() {
			return getRuleContext(VariableContext.class,0);
		}
		public AssignmentOperatorContext assignmentOperator() {
			return getRuleContext(AssignmentOperatorContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode Semi() { return getToken(cminusParser.Semi, 0); }
		public AssignmentStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignmentStatement; }
	}

	public final AssignmentStatementContext assignmentStatement() throws RecognitionException {
		AssignmentStatementContext _localctx = new AssignmentStatementContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_assignmentStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(118);
			variable();
			setState(119);
			assignmentOperator();
			setState(120);
			expression(0);
			setState(121);
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

	public static class AssignmentOperatorContext extends ParserRuleContext {
		public TerminalNode Assign() { return getToken(cminusParser.Assign, 0); }
		public AssignmentOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignmentOperator; }
	}

	public final AssignmentOperatorContext assignmentOperator() throws RecognitionException {
		AssignmentOperatorContext _localctx = new AssignmentOperatorContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_assignmentOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(123);
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

	public static class CompoundStatementContext extends ParserRuleContext {
		public TerminalNode LeftBrace() { return getToken(cminusParser.LeftBrace, 0); }
		public TerminalNode RightBrace() { return getToken(cminusParser.RightBrace, 0); }
		public DeclarationListContext declarationList() {
			return getRuleContext(DeclarationListContext.class,0);
		}
		public StatementListContext statementList() {
			return getRuleContext(StatementListContext.class,0);
		}
		public CompoundStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_compoundStatement; }
	}

	public final CompoundStatementContext compoundStatement() throws RecognitionException {
		CompoundStatementContext _localctx = new CompoundStatementContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_compoundStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(125);
			match(LeftBrace);
			setState(127);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Var) {
				{
				setState(126);
				declarationList(0);
				}
			}

			setState(130);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Break) | (1L << Continue) | (1L << If) | (1L << Return) | (1L << While) | (1L << LeftBrace) | (1L << Identifier))) != 0)) {
				{
				setState(129);
				statementList(0);
				}
			}

			setState(132);
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

	public static class DeclarationContext extends ParserRuleContext {
		public TerminalNode Var() { return getToken(cminusParser.Var, 0); }
		public VariableContext variable() {
			return getRuleContext(VariableContext.class,0);
		}
		public TerminalNode Assign() { return getToken(cminusParser.Assign, 0); }
		public TypeSpecifierContext typeSpecifier() {
			return getRuleContext(TypeSpecifierContext.class,0);
		}
		public TerminalNode Paren() { return getToken(cminusParser.Paren, 0); }
		public TerminalNode Constant() { return getToken(cminusParser.Constant, 0); }
		public TerminalNode Thesis() { return getToken(cminusParser.Thesis, 0); }
		public TerminalNode Semi() { return getToken(cminusParser.Semi, 0); }
		public DeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaration; }
	}

	public final DeclarationContext declaration() throws RecognitionException {
		DeclarationContext _localctx = new DeclarationContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_declaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(134);
			match(Var);
			setState(135);
			variable();
			setState(136);
			match(Assign);
			setState(137);
			typeSpecifier();
			setState(138);
			match(Paren);
			setState(139);
			match(Constant);
			setState(140);
			match(Thesis);
			setState(141);
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
		public TerminalNode If() { return getToken(cminusParser.If, 0); }
		public TerminalNode Paren() { return getToken(cminusParser.Paren, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode Thesis() { return getToken(cminusParser.Thesis, 0); }
		public List<CompoundStatementContext> compoundStatement() {
			return getRuleContexts(CompoundStatementContext.class);
		}
		public CompoundStatementContext compoundStatement(int i) {
			return getRuleContext(CompoundStatementContext.class,i);
		}
		public TerminalNode Else() { return getToken(cminusParser.Else, 0); }
		public SelectionStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selectionStatement; }
	}

	public final SelectionStatementContext selectionStatement() throws RecognitionException {
		SelectionStatementContext _localctx = new SelectionStatementContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_selectionStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(143);
			match(If);
			setState(144);
			match(Paren);
			setState(145);
			expression(0);
			setState(146);
			match(Thesis);
			setState(147);
			compoundStatement();
			setState(150);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				{
				setState(148);
				match(Else);
				setState(149);
				compoundStatement();
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
		public TerminalNode While() { return getToken(cminusParser.While, 0); }
		public TerminalNode Paren() { return getToken(cminusParser.Paren, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode Thesis() { return getToken(cminusParser.Thesis, 0); }
		public CompoundStatementContext compoundStatement() {
			return getRuleContext(CompoundStatementContext.class,0);
		}
		public IterationStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_iterationStatement; }
	}

	public final IterationStatementContext iterationStatement() throws RecognitionException {
		IterationStatementContext _localctx = new IterationStatementContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_iterationStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(152);
			match(While);
			setState(153);
			match(Paren);
			setState(154);
			expression(0);
			setState(155);
			match(Thesis);
			setState(156);
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

	public static class JumpStatementContext extends ParserRuleContext {
		public TerminalNode Continue() { return getToken(cminusParser.Continue, 0); }
		public TerminalNode Semi() { return getToken(cminusParser.Semi, 0); }
		public TerminalNode Break() { return getToken(cminusParser.Break, 0); }
		public TerminalNode Return() { return getToken(cminusParser.Return, 0); }
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
		enterRule(_localctx, 32, RULE_jumpStatement);
		int _la;
		try {
			setState(167);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Continue:
				enterOuterAlt(_localctx, 1);
				{
				setState(158);
				match(Continue);
				setState(159);
				match(Semi);
				}
				break;
			case Break:
				enterOuterAlt(_localctx, 2);
				{
				setState(160);
				match(Break);
				setState(161);
				match(Semi);
				}
				break;
			case Return:
				enterOuterAlt(_localctx, 3);
				{
				setState(162);
				match(Return);
				setState(164);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Paren) | (1L << Minus) | (1L << Identifier) | (1L << Constant))) != 0)) {
					{
					setState(163);
					expression(0);
					}
				}

				setState(166);
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

	public static class ExpressionContext extends ParserRuleContext {
		public FunctionCallContext functionCall() {
			return getRuleContext(FunctionCallContext.class,0);
		}
		public TerminalNode Minus() { return getToken(cminusParser.Minus, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode Paren() { return getToken(cminusParser.Paren, 0); }
		public TerminalNode Thesis() { return getToken(cminusParser.Thesis, 0); }
		public TerminalNode Constant() { return getToken(cminusParser.Constant, 0); }
		public VariableContext variable() {
			return getRuleContext(VariableContext.class,0);
		}
		public MulopContext mulop() {
			return getRuleContext(MulopContext.class,0);
		}
		public AddopContext addop() {
			return getRuleContext(AddopContext.class,0);
		}
		public RelopContext relop() {
			return getRuleContext(RelopContext.class,0);
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
		int _startState = 34;
		enterRecursionRule(_localctx, 34, RULE_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(179);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				{
				setState(170);
				functionCall();
				}
				break;
			case 2:
				{
				setState(171);
				match(Minus);
				setState(172);
				expression(7);
				}
				break;
			case 3:
				{
				setState(173);
				match(Paren);
				setState(174);
				expression(0);
				setState(175);
				match(Thesis);
				}
				break;
			case 4:
				{
				setState(177);
				match(Constant);
				}
				break;
			case 5:
				{
				setState(178);
				variable();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(195);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(193);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
					case 1:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(181);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(182);
						mulop();
						setState(183);
						expression(6);
						}
						break;
					case 2:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(185);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(186);
						addop();
						setState(187);
						expression(5);
						}
						break;
					case 3:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(189);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(190);
						relop();
						setState(191);
						expression(2);
						}
						break;
					}
					} 
				}
				setState(197);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
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
		public TerminalNode Plus() { return getToken(cminusParser.Plus, 0); }
		public TerminalNode Minus() { return getToken(cminusParser.Minus, 0); }
		public AddopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_addop; }
	}

	public final AddopContext addop() throws RecognitionException {
		AddopContext _localctx = new AddopContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_addop);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(198);
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
		public TerminalNode Star() { return getToken(cminusParser.Star, 0); }
		public TerminalNode Div() { return getToken(cminusParser.Div, 0); }
		public TerminalNode Mod() { return getToken(cminusParser.Mod, 0); }
		public MulopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mulop; }
	}

	public final MulopContext mulop() throws RecognitionException {
		MulopContext _localctx = new MulopContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_mulop);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(200);
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

	public static class RelopContext extends ParserRuleContext {
		public TerminalNode EqualEqual() { return getToken(cminusParser.EqualEqual, 0); }
		public TerminalNode NotEqual() { return getToken(cminusParser.NotEqual, 0); }
		public TerminalNode Less() { return getToken(cminusParser.Less, 0); }
		public TerminalNode LessEqual() { return getToken(cminusParser.LessEqual, 0); }
		public TerminalNode Greater() { return getToken(cminusParser.Greater, 0); }
		public TerminalNode GreaterEqual() { return getToken(cminusParser.GreaterEqual, 0); }
		public RelopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relop; }
	}

	public final RelopContext relop() throws RecognitionException {
		RelopContext _localctx = new RelopContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_relop);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(202);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << EqualEqual) | (1L << NotEqual) | (1L << Less) | (1L << LessEqual) | (1L << Greater) | (1L << GreaterEqual))) != 0)) ) {
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

	public static class TypeSpecifierContext extends ParserRuleContext {
		public TerminalNode Int() { return getToken(cminusParser.Int, 0); }
		public TypeSpecifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeSpecifier; }
	}

	public final TypeSpecifierContext typeSpecifier() throws RecognitionException {
		TypeSpecifierContext _localctx = new TypeSpecifierContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_typeSpecifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(204);
			match(Int);
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
		int _startState = 44;
		enterRecursionRule(_localctx, 44, RULE_paramList, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(207);
			param();
			}
			_ctx.stop = _input.LT(-1);
			setState(213);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ParamListContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_paramList);
					setState(209);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(210);
					param();
					}
					} 
				}
				setState(215);
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

	public static class ParamContext extends ParserRuleContext {
		public TypeSpecifierContext typeSpecifier() {
			return getRuleContext(TypeSpecifierContext.class,0);
		}
		public VariableContext variable() {
			return getRuleContext(VariableContext.class,0);
		}
		public TerminalNode Comma() { return getToken(cminusParser.Comma, 0); }
		public ParamContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_param; }
	}

	public final ParamContext param() throws RecognitionException {
		ParamContext _localctx = new ParamContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_param);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(217);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Comma) {
				{
				setState(216);
				match(Comma);
				}
			}

			setState(219);
			typeSpecifier();
			setState(220);
			variable();
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
			return functionList_sempred((FunctionListContext)_localctx, predIndex);
		case 6:
			return statementList_sempred((StatementListContext)_localctx, predIndex);
		case 7:
			return declarationList_sempred((DeclarationListContext)_localctx, predIndex);
		case 17:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		case 22:
			return paramList_sempred((ParamListContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean functionList_sempred(FunctionListContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean statementList_sempred(StatementListContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean declarationList_sempred(DeclarationListContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 5);
		case 4:
			return precpred(_ctx, 4);
		case 5:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean paramList_sempred(ParamListContext _localctx, int predIndex) {
		switch (predIndex) {
		case 6:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3<\u00e1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\3\2\3\2\5\2\65\n\2\3\3\3\3\3\3\3\3\3\3\7\3<\n\3\f\3\16\3?\13\3\3\4\3"+
		"\4\3\4\3\4\3\4\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7"+
		"\3\7\7\7U\n\7\f\7\16\7X\13\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\7\ba\n\b\f\b"+
		"\16\bd\13\b\3\t\3\t\3\t\3\t\3\t\7\tk\n\t\f\t\16\tn\13\t\3\n\3\n\3\n\3"+
		"\n\3\n\5\nu\n\n\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\16\3\16\5\16\u0082"+
		"\n\16\3\16\5\16\u0085\n\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\5\20\u0099\n\20\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u00a7\n\22\3\22"+
		"\5\22\u00aa\n\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\5\23"+
		"\u00b6\n\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23"+
		"\7\23\u00c4\n\23\f\23\16\23\u00c7\13\23\3\24\3\24\3\25\3\25\3\26\3\26"+
		"\3\27\3\27\3\30\3\30\3\30\3\30\3\30\7\30\u00d6\n\30\f\30\16\30\u00d9\13"+
		"\30\3\31\5\31\u00dc\n\31\3\31\3\31\3\31\3\31\2\7\4\16\20$.\32\2\4\6\b"+
		"\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\2\5\4\2\"\"$$\3\2&(\3\2\32"+
		"\37\2\u00e0\2\62\3\2\2\2\4\66\3\2\2\2\6@\3\2\2\2\bE\3\2\2\2\nG\3\2\2\2"+
		"\fO\3\2\2\2\16[\3\2\2\2\20e\3\2\2\2\22t\3\2\2\2\24v\3\2\2\2\26x\3\2\2"+
		"\2\30}\3\2\2\2\32\177\3\2\2\2\34\u0088\3\2\2\2\36\u0091\3\2\2\2 \u009a"+
		"\3\2\2\2\"\u00a9\3\2\2\2$\u00b5\3\2\2\2&\u00c8\3\2\2\2(\u00ca\3\2\2\2"+
		"*\u00cc\3\2\2\2,\u00ce\3\2\2\2.\u00d0\3\2\2\2\60\u00db\3\2\2\2\62\64\5"+
		"\6\4\2\63\65\5\4\3\2\64\63\3\2\2\2\64\65\3\2\2\2\65\3\3\2\2\2\66\67\b"+
		"\3\1\2\678\5\n\6\28=\3\2\2\29:\f\3\2\2:<\5\n\6\2;9\3\2\2\2<?\3\2\2\2="+
		";\3\2\2\2=>\3\2\2\2>\5\3\2\2\2?=\3\2\2\2@A\7\21\2\2AB\7\24\2\2BC\7\25"+
		"\2\2CD\5\32\16\2D\7\3\2\2\2EF\5\24\13\2F\t\3\2\2\2GH\7\20\2\2HI\5\b\5"+
		"\2IJ\7\24\2\2JK\5.\30\2KL\7\25\2\2LM\5,\27\2MN\5\32\16\2N\13\3\2\2\2O"+
		"P\5\b\5\2PQ\7\24\2\2QV\5\24\13\2RS\7\63\2\2SU\5\24\13\2TR\3\2\2\2UX\3"+
		"\2\2\2VT\3\2\2\2VW\3\2\2\2WY\3\2\2\2XV\3\2\2\2YZ\7\25\2\2Z\r\3\2\2\2["+
		"\\\b\b\1\2\\]\5\22\n\2]b\3\2\2\2^_\f\3\2\2_a\5\22\n\2`^\3\2\2\2ad\3\2"+
		"\2\2b`\3\2\2\2bc\3\2\2\2c\17\3\2\2\2db\3\2\2\2ef\b\t\1\2fg\5\34\17\2g"+
		"l\3\2\2\2hi\f\3\2\2ik\5\34\17\2jh\3\2\2\2kn\3\2\2\2lj\3\2\2\2lm\3\2\2"+
		"\2m\21\3\2\2\2nl\3\2\2\2ou\5\32\16\2pu\5\36\20\2qu\5 \21\2ru\5\"\22\2"+
		"su\5\26\f\2to\3\2\2\2tp\3\2\2\2tq\3\2\2\2tr\3\2\2\2ts\3\2\2\2u\23\3\2"+
		"\2\2vw\7\65\2\2w\25\3\2\2\2xy\5\24\13\2yz\5\30\r\2z{\5$\23\2{|\7\62\2"+
		"\2|\27\3\2\2\2}~\7\64\2\2~\31\3\2\2\2\177\u0081\7\30\2\2\u0080\u0082\5"+
		"\20\t\2\u0081\u0080\3\2\2\2\u0081\u0082\3\2\2\2\u0082\u0084\3\2\2\2\u0083"+
		"\u0085\5\16\b\2\u0084\u0083\3\2\2\2\u0084\u0085\3\2\2\2\u0085\u0086\3"+
		"\2\2\2\u0086\u0087\7\31\2\2\u0087\33\3\2\2\2\u0088\u0089\7\7\2\2\u0089"+
		"\u008a\5\24\13\2\u008a\u008b\7\64\2\2\u008b\u008c\5,\27\2\u008c\u008d"+
		"\7\24\2\2\u008d\u008e\7\66\2\2\u008e\u008f\7\25\2\2\u008f\u0090\7\62\2"+
		"\2\u0090\35\3\2\2\2\u0091\u0092\7\6\2\2\u0092\u0093\7\24\2\2\u0093\u0094"+
		"\5$\23\2\u0094\u0095\7\25\2\2\u0095\u0098\5\32\16\2\u0096\u0097\7\5\2"+
		"\2\u0097\u0099\5\32\16\2\u0098\u0096\3\2\2\2\u0098\u0099\3\2\2\2\u0099"+
		"\37\3\2\2\2\u009a\u009b\7\17\2\2\u009b\u009c\7\24\2\2\u009c\u009d\5$\23"+
		"\2\u009d\u009e\7\25\2\2\u009e\u009f\5\32\16\2\u009f!\3\2\2\2\u00a0\u00a1"+
		"\7\4\2\2\u00a1\u00aa\7\62\2\2\u00a2\u00a3\7\3\2\2\u00a3\u00aa\7\62\2\2"+
		"\u00a4\u00a6\7\r\2\2\u00a5\u00a7\5$\23\2\u00a6\u00a5\3\2\2\2\u00a6\u00a7"+
		"\3\2\2\2\u00a7\u00a8\3\2\2\2\u00a8\u00aa\7\62\2\2\u00a9\u00a0\3\2\2\2"+
		"\u00a9\u00a2\3\2\2\2\u00a9\u00a4\3\2\2\2\u00aa#\3\2\2\2\u00ab\u00ac\b"+
		"\23\1\2\u00ac\u00b6\5\f\7\2\u00ad\u00ae\7$\2\2\u00ae\u00b6\5$\23\t\u00af"+
		"\u00b0\7\24\2\2\u00b0\u00b1\5$\23\2\u00b1\u00b2\7\25\2\2\u00b2\u00b6\3"+
		"\2\2\2\u00b3\u00b6\7\66\2\2\u00b4\u00b6\5\24\13\2\u00b5\u00ab\3\2\2\2"+
		"\u00b5\u00ad\3\2\2\2\u00b5\u00af\3\2\2\2\u00b5\u00b3\3\2\2\2\u00b5\u00b4"+
		"\3\2\2\2\u00b6\u00c5\3\2\2\2\u00b7\u00b8\f\7\2\2\u00b8\u00b9\5(\25\2\u00b9"+
		"\u00ba\5$\23\b\u00ba\u00c4\3\2\2\2\u00bb\u00bc\f\6\2\2\u00bc\u00bd\5&"+
		"\24\2\u00bd\u00be\5$\23\7\u00be\u00c4\3\2\2\2\u00bf\u00c0\f\3\2\2\u00c0"+
		"\u00c1\5*\26\2\u00c1\u00c2\5$\23\4\u00c2\u00c4\3\2\2\2\u00c3\u00b7\3\2"+
		"\2\2\u00c3\u00bb\3\2\2\2\u00c3\u00bf\3\2\2\2\u00c4\u00c7\3\2\2\2\u00c5"+
		"\u00c3\3\2\2\2\u00c5\u00c6\3\2\2\2\u00c6%\3\2\2\2\u00c7\u00c5\3\2\2\2"+
		"\u00c8\u00c9\t\2\2\2\u00c9\'\3\2\2\2\u00ca\u00cb\t\3\2\2\u00cb)\3\2\2"+
		"\2\u00cc\u00cd\t\4\2\2\u00cd+\3\2\2\2\u00ce\u00cf\7\b\2\2\u00cf-\3\2\2"+
		"\2\u00d0\u00d1\b\30\1\2\u00d1\u00d2\5\60\31\2\u00d2\u00d7\3\2\2\2\u00d3"+
		"\u00d4\f\3\2\2\u00d4\u00d6\5\60\31\2\u00d5\u00d3\3\2\2\2\u00d6\u00d9\3"+
		"\2\2\2\u00d7\u00d5\3\2\2\2\u00d7\u00d8\3\2\2\2\u00d8/\3\2\2\2\u00d9\u00d7"+
		"\3\2\2\2\u00da\u00dc\7\63\2\2\u00db\u00da\3\2\2\2\u00db\u00dc\3\2\2\2"+
		"\u00dc\u00dd\3\2\2\2\u00dd\u00de\5,\27\2\u00de\u00df\5\24\13\2\u00df\61"+
		"\3\2\2\2\22\64=Vblt\u0081\u0084\u0098\u00a6\u00a9\u00b5\u00c3\u00c5\u00d7"+
		"\u00db";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}