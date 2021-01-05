// Generated from /home/kristofer/p/c-minus/cminusgo/parser/cminus.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class cminusLexer extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "Break", "Continue", "Else", "If", "Int", "Bool", 
			"StringType", "Rune", "Array", "Return", "Void", "While", "Func", "True", 
			"False", "Paren", "Thesis", "LeftBracket", "RightBracket", "LeftBrace", 
			"RightBrace", "Less", "LessEqual", "Greater", "GreaterEqual", "LeftShift", 
			"RightShift", "Plus", "PlusPlus", "Minus", "MinusMinus", "Star", "Div", 
			"Mod", "And", "Or", "AndAnd", "OrOr", "Caret", "Not", "Tilde", "Question", 
			"Colon", "Semi", "Comma", "Assign", "Identifier", "IdentifierNondigit", 
			"Nondigit", "Digit", "UniversalCharacterName", "HexQuad", "Constant", 
			"IntegerConstant", "BinaryConstant", "BooleanConstant", "DecimalConstant", 
			"OctalConstant", "HexadecimalConstant", "HexadecimalPrefix", "NonzeroDigit", 
			"OctalDigit", "HexadecimalDigit", "Sign", "DigitSequence", "HexadecimalDigitSequence", 
			"CharacterConstant", "CCharSequence", "CChar", "EscapeSequence", "SimpleEscapeSequence", 
			"OctalEscapeSequence", "HexadecimalEscapeSequence", "StringLiteral", 
			"SCharSequence", "SChar", "Whitespace", "Newline", "BlockComment", "LineComment"
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


	public cminusLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "cminus.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2;\u0230\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\4J\tJ\4K\tK\4L\tL\4M\tM\4N\tN\4O\tO\4P\tP\4Q\tQ\4R\tR\4S\tS\4T\tT"+
		"\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\t\3\t"+
		"\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22"+
		"\3\23\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\25\3\25\3\26\3\26\3\27\3\27"+
		"\3\30\3\30\3\31\3\31\3\32\3\32\3\33\3\33\3\33\3\34\3\34\3\35\3\35\3\35"+
		"\3\36\3\36\3\36\3\37\3\37\3\37\3 \3 \3!\3!\3!\3\"\3\"\3#\3#\3#\3$\3$\3"+
		"%\3%\3&\3&\3\'\3\'\3(\3(\3)\3)\3)\3*\3*\3*\3+\3+\3,\3,\3-\3-\3.\3.\3/"+
		"\3/\3\60\3\60\3\61\3\61\3\62\3\62\3\63\3\63\3\63\7\63\u0158\n\63\f\63"+
		"\16\63\u015b\13\63\3\64\3\64\5\64\u015f\n\64\3\65\3\65\3\66\3\66\3\67"+
		"\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\5\67\u016f\n\67\38\38\3"+
		"8\38\38\39\39\59\u0178\n9\3:\3:\3:\3:\3:\5:\u017f\n:\3;\3;\3;\6;\u0184"+
		"\n;\r;\16;\u0185\3<\3<\5<\u018a\n<\3=\3=\7=\u018e\n=\f=\16=\u0191\13="+
		"\3>\3>\7>\u0195\n>\f>\16>\u0198\13>\3?\3?\6?\u019c\n?\r?\16?\u019d\3@"+
		"\3@\3@\3A\3A\3B\3B\3C\3C\3D\3D\3E\6E\u01ac\nE\rE\16E\u01ad\3F\6F\u01b1"+
		"\nF\rF\16F\u01b2\3G\3G\3G\3G\3G\3G\3G\3G\3G\3G\3G\3G\3G\3G\3G\3G\3G\3"+
		"G\3G\3G\3G\3G\5G\u01cb\nG\3H\6H\u01ce\nH\rH\16H\u01cf\3I\3I\5I\u01d4\n"+
		"I\3J\3J\3J\3J\5J\u01da\nJ\3K\3K\3K\3L\3L\3L\3L\3L\3L\3L\3L\3L\3L\3L\5"+
		"L\u01ea\nL\3M\3M\3M\3M\6M\u01f0\nM\rM\16M\u01f1\3N\3N\5N\u01f6\nN\3N\3"+
		"N\3O\6O\u01fb\nO\rO\16O\u01fc\3P\3P\3P\3P\3P\3P\3P\5P\u0206\nP\3Q\6Q\u0209"+
		"\nQ\rQ\16Q\u020a\3Q\3Q\3R\3R\5R\u0211\nR\3R\5R\u0214\nR\3R\3R\3S\3S\3"+
		"S\3S\7S\u021c\nS\fS\16S\u021f\13S\3S\3S\3S\3S\3S\3T\3T\3T\3T\7T\u022a"+
		"\nT\fT\16T\u022d\13T\3T\3T\3\u021d\2U\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21"+
		"\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30"+
		"/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.["+
		"/]\60_\61a\62c\63e\64g\2i\2k\2m\2o\2q\65s\2u\2w\2y\2{\2}\2\177\2\u0081"+
		"\2\u0083\2\u0085\2\u0087\2\u0089\66\u008b\2\u008d\2\u008f\2\u0091\2\u0093"+
		"\2\u0095\2\u0097\2\u0099\2\u009b\67\u009d\2\u009f\2\u00a18\u00a39\u00a5"+
		":\u00a7;\3\2\20\5\2C\\aac|\3\2\62;\4\2DDdd\3\2\62\63\4\2ZZzz\3\2\63;\3"+
		"\2\629\5\2\62;CHch\4\2--//\6\2\f\f\17\17))^^\f\2$$))AA^^cdhhppttvvxx\6"+
		"\2\f\f\17\17$$^^\4\2\13\13\"\"\4\2\f\f\17\17\2\u023a\2\3\3\2\2\2\2\5\3"+
		"\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2"+
		"\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3"+
		"\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'"+
		"\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63"+
		"\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2"+
		"?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3"+
		"\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2"+
		"\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2"+
		"e\3\2\2\2\2q\3\2\2\2\2\u0089\3\2\2\2\2\u009b\3\2\2\2\2\u00a1\3\2\2\2\2"+
		"\u00a3\3\2\2\2\2\u00a5\3\2\2\2\2\u00a7\3\2\2\2\3\u00a9\3\2\2\2\5\u00ad"+
		"\3\2\2\2\7\u00b0\3\2\2\2\t\u00b3\3\2\2\2\13\u00b9\3\2\2\2\r\u00c2\3\2"+
		"\2\2\17\u00c7\3\2\2\2\21\u00ca\3\2\2\2\23\u00ce\3\2\2\2\25\u00d6\3\2\2"+
		"\2\27\u00dd\3\2\2\2\31\u00e2\3\2\2\2\33\u00e8\3\2\2\2\35\u00ef\3\2\2\2"+
		"\37\u00f4\3\2\2\2!\u00fa\3\2\2\2#\u0103\3\2\2\2%\u0108\3\2\2\2\'\u010e"+
		"\3\2\2\2)\u0110\3\2\2\2+\u0112\3\2\2\2-\u0114\3\2\2\2/\u0116\3\2\2\2\61"+
		"\u0118\3\2\2\2\63\u011a\3\2\2\2\65\u011c\3\2\2\2\67\u011f\3\2\2\29\u0121"+
		"\3\2\2\2;\u0124\3\2\2\2=\u0127\3\2\2\2?\u012a\3\2\2\2A\u012c\3\2\2\2C"+
		"\u012f\3\2\2\2E\u0131\3\2\2\2G\u0134\3\2\2\2I\u0136\3\2\2\2K\u0138\3\2"+
		"\2\2M\u013a\3\2\2\2O\u013c\3\2\2\2Q\u013e\3\2\2\2S\u0141\3\2\2\2U\u0144"+
		"\3\2\2\2W\u0146\3\2\2\2Y\u0148\3\2\2\2[\u014a\3\2\2\2]\u014c\3\2\2\2_"+
		"\u014e\3\2\2\2a\u0150\3\2\2\2c\u0152\3\2\2\2e\u0154\3\2\2\2g\u015e\3\2"+
		"\2\2i\u0160\3\2\2\2k\u0162\3\2\2\2m\u016e\3\2\2\2o\u0170\3\2\2\2q\u0177"+
		"\3\2\2\2s\u017e\3\2\2\2u\u0180\3\2\2\2w\u0189\3\2\2\2y\u018b\3\2\2\2{"+
		"\u0192\3\2\2\2}\u0199\3\2\2\2\177\u019f\3\2\2\2\u0081\u01a2\3\2\2\2\u0083"+
		"\u01a4\3\2\2\2\u0085\u01a6\3\2\2\2\u0087\u01a8\3\2\2\2\u0089\u01ab\3\2"+
		"\2\2\u008b\u01b0\3\2\2\2\u008d\u01ca\3\2\2\2\u008f\u01cd\3\2\2\2\u0091"+
		"\u01d3\3\2\2\2\u0093\u01d9\3\2\2\2\u0095\u01db\3\2\2\2\u0097\u01e9\3\2"+
		"\2\2\u0099\u01eb\3\2\2\2\u009b\u01f3\3\2\2\2\u009d\u01fa\3\2\2\2\u009f"+
		"\u0205\3\2\2\2\u00a1\u0208\3\2\2\2\u00a3\u0213\3\2\2\2\u00a5\u0217\3\2"+
		"\2\2\u00a7\u0225\3\2\2\2\u00a9\u00aa\7x\2\2\u00aa\u00ab\7c\2\2\u00ab\u00ac"+
		"\7t\2\2\u00ac\4\3\2\2\2\u00ad\u00ae\7?\2\2\u00ae\u00af\7?\2\2\u00af\6"+
		"\3\2\2\2\u00b0\u00b1\7#\2\2\u00b1\u00b2\7?\2\2\u00b2\b\3\2\2\2\u00b3\u00b4"+
		"\7d\2\2\u00b4\u00b5\7t\2\2\u00b5\u00b6\7g\2\2\u00b6\u00b7\7c\2\2\u00b7"+
		"\u00b8\7m\2\2\u00b8\n\3\2\2\2\u00b9\u00ba\7e\2\2\u00ba\u00bb\7q\2\2\u00bb"+
		"\u00bc\7p\2\2\u00bc\u00bd\7v\2\2\u00bd\u00be\7k\2\2\u00be\u00bf\7p\2\2"+
		"\u00bf\u00c0\7w\2\2\u00c0\u00c1\7g\2\2\u00c1\f\3\2\2\2\u00c2\u00c3\7g"+
		"\2\2\u00c3\u00c4\7n\2\2\u00c4\u00c5\7u\2\2\u00c5\u00c6\7g\2\2\u00c6\16"+
		"\3\2\2\2\u00c7\u00c8\7k\2\2\u00c8\u00c9\7h\2\2\u00c9\20\3\2\2\2\u00ca"+
		"\u00cb\7k\2\2\u00cb\u00cc\7p\2\2\u00cc\u00cd\7v\2\2\u00cd\22\3\2\2\2\u00ce"+
		"\u00cf\7d\2\2\u00cf\u00d0\7q\2\2\u00d0\u00d1\7q\2\2\u00d1\u00d2\7n\2\2"+
		"\u00d2\u00d3\7g\2\2\u00d3\u00d4\7c\2\2\u00d4\u00d5\7p\2\2\u00d5\24\3\2"+
		"\2\2\u00d6\u00d7\7u\2\2\u00d7\u00d8\7v\2\2\u00d8\u00d9\7t\2\2\u00d9\u00da"+
		"\7k\2\2\u00da\u00db\7p\2\2\u00db\u00dc\7i\2\2\u00dc\26\3\2\2\2\u00dd\u00de"+
		"\7t\2\2\u00de\u00df\7w\2\2\u00df\u00e0\7p\2\2\u00e0\u00e1\7g\2\2\u00e1"+
		"\30\3\2\2\2\u00e2\u00e3\7c\2\2\u00e3\u00e4\7t\2\2\u00e4\u00e5\7t\2\2\u00e5"+
		"\u00e6\7c\2\2\u00e6\u00e7\7{\2\2\u00e7\32\3\2\2\2\u00e8\u00e9\7t\2\2\u00e9"+
		"\u00ea\7g\2\2\u00ea\u00eb\7v\2\2\u00eb\u00ec\7w\2\2\u00ec\u00ed\7t\2\2"+
		"\u00ed\u00ee\7p\2\2\u00ee\34\3\2\2\2\u00ef\u00f0\7x\2\2\u00f0\u00f1\7"+
		"q\2\2\u00f1\u00f2\7k\2\2\u00f2\u00f3\7f\2\2\u00f3\36\3\2\2\2\u00f4\u00f5"+
		"\7y\2\2\u00f5\u00f6\7j\2\2\u00f6\u00f7\7k\2\2\u00f7\u00f8\7n\2\2\u00f8"+
		"\u00f9\7g\2\2\u00f9 \3\2\2\2\u00fa\u00fb\7h\2\2\u00fb\u00fc\7w\2\2\u00fc"+
		"\u00fd\7p\2\2\u00fd\u00fe\7e\2\2\u00fe\u00ff\7v\2\2\u00ff\u0100\7k\2\2"+
		"\u0100\u0101\7q\2\2\u0101\u0102\7p\2\2\u0102\"\3\2\2\2\u0103\u0104\7v"+
		"\2\2\u0104\u0105\7t\2\2\u0105\u0106\7w\2\2\u0106\u0107\7g\2\2\u0107$\3"+
		"\2\2\2\u0108\u0109\7h\2\2\u0109\u010a\7c\2\2\u010a\u010b\7n\2\2\u010b"+
		"\u010c\7u\2\2\u010c\u010d\7g\2\2\u010d&\3\2\2\2\u010e\u010f\7*\2\2\u010f"+
		"(\3\2\2\2\u0110\u0111\7+\2\2\u0111*\3\2\2\2\u0112\u0113\7]\2\2\u0113,"+
		"\3\2\2\2\u0114\u0115\7_\2\2\u0115.\3\2\2\2\u0116\u0117\7}\2\2\u0117\60"+
		"\3\2\2\2\u0118\u0119\7\177\2\2\u0119\62\3\2\2\2\u011a\u011b\7>\2\2\u011b"+
		"\64\3\2\2\2\u011c\u011d\7>\2\2\u011d\u011e\7?\2\2\u011e\66\3\2\2\2\u011f"+
		"\u0120\7@\2\2\u01208\3\2\2\2\u0121\u0122\7@\2\2\u0122\u0123\7?\2\2\u0123"+
		":\3\2\2\2\u0124\u0125\7>\2\2\u0125\u0126\7>\2\2\u0126<\3\2\2\2\u0127\u0128"+
		"\7@\2\2\u0128\u0129\7@\2\2\u0129>\3\2\2\2\u012a\u012b\7-\2\2\u012b@\3"+
		"\2\2\2\u012c\u012d\7-\2\2\u012d\u012e\7-\2\2\u012eB\3\2\2\2\u012f\u0130"+
		"\7/\2\2\u0130D\3\2\2\2\u0131\u0132\7/\2\2\u0132\u0133\7/\2\2\u0133F\3"+
		"\2\2\2\u0134\u0135\7,\2\2\u0135H\3\2\2\2\u0136\u0137\7\61\2\2\u0137J\3"+
		"\2\2\2\u0138\u0139\7\'\2\2\u0139L\3\2\2\2\u013a\u013b\7(\2\2\u013bN\3"+
		"\2\2\2\u013c\u013d\7~\2\2\u013dP\3\2\2\2\u013e\u013f\7(\2\2\u013f\u0140"+
		"\7(\2\2\u0140R\3\2\2\2\u0141\u0142\7~\2\2\u0142\u0143\7~\2\2\u0143T\3"+
		"\2\2\2\u0144\u0145\7`\2\2\u0145V\3\2\2\2\u0146\u0147\7#\2\2\u0147X\3\2"+
		"\2\2\u0148\u0149\7\u0080\2\2\u0149Z\3\2\2\2\u014a\u014b\7A\2\2\u014b\\"+
		"\3\2\2\2\u014c\u014d\7<\2\2\u014d^\3\2\2\2\u014e\u014f\7=\2\2\u014f`\3"+
		"\2\2\2\u0150\u0151\7.\2\2\u0151b\3\2\2\2\u0152\u0153\7?\2\2\u0153d\3\2"+
		"\2\2\u0154\u0159\5g\64\2\u0155\u0158\5g\64\2\u0156\u0158\5k\66\2\u0157"+
		"\u0155\3\2\2\2\u0157\u0156\3\2\2\2\u0158\u015b\3\2\2\2\u0159\u0157\3\2"+
		"\2\2\u0159\u015a\3\2\2\2\u015af\3\2\2\2\u015b\u0159\3\2\2\2\u015c\u015f"+
		"\5i\65\2\u015d\u015f\5m\67\2\u015e\u015c\3\2\2\2\u015e\u015d\3\2\2\2\u015f"+
		"h\3\2\2\2\u0160\u0161\t\2\2\2\u0161j\3\2\2\2\u0162\u0163\t\3\2\2\u0163"+
		"l\3\2\2\2\u0164\u0165\7^\2\2\u0165\u0166\7w\2\2\u0166\u0167\3\2\2\2\u0167"+
		"\u016f\5o8\2\u0168\u0169\7^\2\2\u0169\u016a\7W\2\2\u016a\u016b\3\2\2\2"+
		"\u016b\u016c\5o8\2\u016c\u016d\5o8\2\u016d\u016f\3\2\2\2\u016e\u0164\3"+
		"\2\2\2\u016e\u0168\3\2\2\2\u016fn\3\2\2\2\u0170\u0171\5\u0085C\2\u0171"+
		"\u0172\5\u0085C\2\u0172\u0173\5\u0085C\2\u0173\u0174\5\u0085C\2\u0174"+
		"p\3\2\2\2\u0175\u0178\5s:\2\u0176\u0178\5\u008dG\2\u0177\u0175\3\2\2\2"+
		"\u0177\u0176\3\2\2\2\u0178r\3\2\2\2\u0179\u017f\5y=\2\u017a\u017f\5{>"+
		"\2\u017b\u017f\5}?\2\u017c\u017f\5u;\2\u017d\u017f\5w<\2\u017e\u0179\3"+
		"\2\2\2\u017e\u017a\3\2\2\2\u017e\u017b\3\2\2\2\u017e\u017c\3\2\2\2\u017e"+
		"\u017d\3\2\2\2\u017ft\3\2\2\2\u0180\u0181\7\62\2\2\u0181\u0183\t\4\2\2"+
		"\u0182\u0184\t\5\2\2\u0183\u0182\3\2\2\2\u0184\u0185\3\2\2\2\u0185\u0183"+
		"\3\2\2\2\u0185\u0186\3\2\2\2\u0186v\3\2\2\2\u0187\u018a\5#\22\2\u0188"+
		"\u018a\5%\23\2\u0189\u0187\3\2\2\2\u0189\u0188\3\2\2\2\u018ax\3\2\2\2"+
		"\u018b\u018f\5\u0081A\2\u018c\u018e\5k\66\2\u018d\u018c\3\2\2\2\u018e"+
		"\u0191\3\2\2\2\u018f\u018d\3\2\2\2\u018f\u0190\3\2\2\2\u0190z\3\2\2\2"+
		"\u0191\u018f\3\2\2\2\u0192\u0196\7\62\2\2\u0193\u0195\5\u0083B\2\u0194"+
		"\u0193\3\2\2\2\u0195\u0198\3\2\2\2\u0196\u0194\3\2\2\2\u0196\u0197\3\2"+
		"\2\2\u0197|\3\2\2\2\u0198\u0196\3\2\2\2\u0199\u019b\5\177@\2\u019a\u019c"+
		"\5\u0085C\2\u019b\u019a\3\2\2\2\u019c\u019d\3\2\2\2\u019d\u019b\3\2\2"+
		"\2\u019d\u019e\3\2\2\2\u019e~\3\2\2\2\u019f\u01a0\7\62\2\2\u01a0\u01a1"+
		"\t\6\2\2\u01a1\u0080\3\2\2\2\u01a2\u01a3\t\7\2\2\u01a3\u0082\3\2\2\2\u01a4"+
		"\u01a5\t\b\2\2\u01a5\u0084\3\2\2\2\u01a6\u01a7\t\t\2\2\u01a7\u0086\3\2"+
		"\2\2\u01a8\u01a9\t\n\2\2\u01a9\u0088\3\2\2\2\u01aa\u01ac\5k\66\2\u01ab"+
		"\u01aa\3\2\2\2\u01ac\u01ad\3\2\2\2\u01ad\u01ab\3\2\2\2\u01ad\u01ae\3\2"+
		"\2\2\u01ae\u008a\3\2\2\2\u01af\u01b1\5\u0085C\2\u01b0\u01af\3\2\2\2\u01b1"+
		"\u01b2\3\2\2\2\u01b2\u01b0\3\2\2\2\u01b2\u01b3\3\2\2\2\u01b3\u008c\3\2"+
		"\2\2\u01b4\u01b5\7)\2\2\u01b5\u01b6\5\u008fH\2\u01b6\u01b7\7)\2\2\u01b7"+
		"\u01cb\3\2\2\2\u01b8\u01b9\7N\2\2\u01b9\u01ba\7)\2\2\u01ba\u01bb\3\2\2"+
		"\2\u01bb\u01bc\5\u008fH\2\u01bc\u01bd\7)\2\2\u01bd\u01cb\3\2\2\2\u01be"+
		"\u01bf\7w\2\2\u01bf\u01c0\7)\2\2\u01c0\u01c1\3\2\2\2\u01c1\u01c2\5\u008f"+
		"H\2\u01c2\u01c3\7)\2\2\u01c3\u01cb\3\2\2\2\u01c4\u01c5\7W\2\2\u01c5\u01c6"+
		"\7)\2\2\u01c6\u01c7\3\2\2\2\u01c7\u01c8\5\u008fH\2\u01c8\u01c9\7)\2\2"+
		"\u01c9\u01cb\3\2\2\2\u01ca\u01b4\3\2\2\2\u01ca\u01b8\3\2\2\2\u01ca\u01be"+
		"\3\2\2\2\u01ca\u01c4\3\2\2\2\u01cb\u008e\3\2\2\2\u01cc\u01ce\5\u0091I"+
		"\2\u01cd\u01cc\3\2\2\2\u01ce\u01cf\3\2\2\2\u01cf\u01cd\3\2\2\2\u01cf\u01d0"+
		"\3\2\2\2\u01d0\u0090\3\2\2\2\u01d1\u01d4\n\13\2\2\u01d2\u01d4\5\u0093"+
		"J\2\u01d3\u01d1\3\2\2\2\u01d3\u01d2\3\2\2\2\u01d4\u0092\3\2\2\2\u01d5"+
		"\u01da\5\u0095K\2\u01d6\u01da\5\u0097L\2\u01d7\u01da\5\u0099M\2\u01d8"+
		"\u01da\5m\67\2\u01d9\u01d5\3\2\2\2\u01d9\u01d6\3\2\2\2\u01d9\u01d7\3\2"+
		"\2\2\u01d9\u01d8\3\2\2\2\u01da\u0094\3\2\2\2\u01db\u01dc\7^\2\2\u01dc"+
		"\u01dd\t\f\2\2\u01dd\u0096\3\2\2\2\u01de\u01df\7^\2\2\u01df\u01ea\5\u0083"+
		"B\2\u01e0\u01e1\7^\2\2\u01e1\u01e2\5\u0083B\2\u01e2\u01e3\5\u0083B\2\u01e3"+
		"\u01ea\3\2\2\2\u01e4\u01e5\7^\2\2\u01e5\u01e6\5\u0083B\2\u01e6\u01e7\5"+
		"\u0083B\2\u01e7\u01e8\5\u0083B\2\u01e8\u01ea\3\2\2\2\u01e9\u01de\3\2\2"+
		"\2\u01e9\u01e0\3\2\2\2\u01e9\u01e4\3\2\2\2\u01ea\u0098\3\2\2\2\u01eb\u01ec"+
		"\7^\2\2\u01ec\u01ed\7z\2\2\u01ed\u01ef\3\2\2\2\u01ee\u01f0\5\u0085C\2"+
		"\u01ef\u01ee\3\2\2\2\u01f0\u01f1\3\2\2\2\u01f1\u01ef\3\2\2\2\u01f1\u01f2"+
		"\3\2\2\2\u01f2\u009a\3\2\2\2\u01f3\u01f5\7$\2\2\u01f4\u01f6\5\u009dO\2"+
		"\u01f5\u01f4\3\2\2\2\u01f5\u01f6\3\2\2\2\u01f6\u01f7\3\2\2\2\u01f7\u01f8"+
		"\7$\2\2\u01f8\u009c\3\2\2\2\u01f9\u01fb\5\u009fP\2\u01fa\u01f9\3\2\2\2"+
		"\u01fb\u01fc\3\2\2\2\u01fc\u01fa\3\2\2\2\u01fc\u01fd\3\2\2\2\u01fd\u009e"+
		"\3\2\2\2\u01fe\u0206\n\r\2\2\u01ff\u0206\5\u0093J\2\u0200\u0201\7^\2\2"+
		"\u0201\u0206\7\f\2\2\u0202\u0203\7^\2\2\u0203\u0204\7\17\2\2\u0204\u0206"+
		"\7\f\2\2\u0205\u01fe\3\2\2\2\u0205\u01ff\3\2\2\2\u0205\u0200\3\2\2\2\u0205"+
		"\u0202\3\2\2\2\u0206\u00a0\3\2\2\2\u0207\u0209\t\16\2\2\u0208\u0207\3"+
		"\2\2\2\u0209\u020a\3\2\2\2\u020a\u0208\3\2\2\2\u020a\u020b\3\2\2\2\u020b"+
		"\u020c\3\2\2\2\u020c\u020d\bQ\2\2\u020d\u00a2\3\2\2\2\u020e\u0210\7\17"+
		"\2\2\u020f\u0211\7\f\2\2\u0210\u020f\3\2\2\2\u0210\u0211\3\2\2\2\u0211"+
		"\u0214\3\2\2\2\u0212\u0214\7\f\2\2\u0213\u020e\3\2\2\2\u0213\u0212\3\2"+
		"\2\2\u0214\u0215\3\2\2\2\u0215\u0216\bR\2\2\u0216\u00a4\3\2\2\2\u0217"+
		"\u0218\7\61\2\2\u0218\u0219\7,\2\2\u0219\u021d\3\2\2\2\u021a\u021c\13"+
		"\2\2\2\u021b\u021a\3\2\2\2\u021c\u021f\3\2\2\2\u021d\u021e\3\2\2\2\u021d"+
		"\u021b\3\2\2\2\u021e\u0220\3\2\2\2\u021f\u021d\3\2\2\2\u0220\u0221\7,"+
		"\2\2\u0221\u0222\7\61\2\2\u0222\u0223\3\2\2\2\u0223\u0224\bS\2\2\u0224"+
		"\u00a6\3\2\2\2\u0225\u0226\7\61\2\2\u0226\u0227\7\61\2\2\u0227\u022b\3"+
		"\2\2\2\u0228\u022a\n\17\2\2\u0229\u0228\3\2\2\2\u022a\u022d\3\2\2\2\u022b"+
		"\u0229\3\2\2\2\u022b\u022c\3\2\2\2\u022c\u022e\3\2\2\2\u022d\u022b\3\2"+
		"\2\2\u022e\u022f\bT\2\2\u022f\u00a8\3\2\2\2\36\2\u0157\u0159\u015e\u016e"+
		"\u0177\u017e\u0185\u0189\u018f\u0196\u019d\u01ad\u01b2\u01ca\u01cf\u01d3"+
		"\u01d9\u01e9\u01f1\u01f5\u01fc\u0205\u020a\u0210\u0213\u021d\u022b\3\b"+
		"\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}