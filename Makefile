build:
	go build ./dist/donkey 

test:
	go test -v ./...

test-lexer:
	go test -v ./lexer 

test-ast:
	go test -v ./ast 

test-parser:
	go test -v ./parser 