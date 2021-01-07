all: compile

./parser/*.go : ./cminus.g4
	antlr4 -Dlanguage=Go -o parser cminus.g4
	#go build ./...

compile: ./parser/*.go ./codegen/*.go
	go run ./cmc.go

clean:
	go clean ./...

