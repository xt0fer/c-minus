# c-minus
a pico C for ZipRISC1
return-type methodname()
take a look in Spec.md for the very basic language definition/specification.

ANTLR for lex/parser? (it can generate Java).


Some sed lines for changing the Export status of the Parser inside of the generated code.

# fix gen'd file sed -i "s/[[:space:]]parser antlr.Parser/\ Parser antlr.Parser/g" *.go
# sed -i "s/[[:space:]]p.parser =/\ p.Parser =/g" parser/*.go
# sed -i "s/return s.parser/return s.Parser/g" parser/*.go
