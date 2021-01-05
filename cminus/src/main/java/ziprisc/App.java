package ziprisc;

/**
 * Hello world!
 *
 */
public class App 
{
    public static void main( String[] args )
    {
        String javaClassContent = "public class SampleClass { void DoSomething(){} }";
        CminusLexer CminusLexer = new CminusLexer(CharStreams.fromString(javaClassContent));
        }
}
