all: compile

./parser/*.java : ./Cminus2.g4
	antlr4 -Dlanguage=Java -o parser Cminus2.g4
	#go build ./...

./codegen/*.java :
	touch ZRGen.java
	
compile: ./parser/*.java 
	javac ./parser/*.java

clean:
	# go clean ./...

