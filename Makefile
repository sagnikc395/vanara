build:
	mkdir -p ./dist 
	go build  -o ./dist/vanara
run:
	./dist/vanara

test:
	go test -v ./...

test-lexer:
	go test -v ./lexer 

test-ast:
	go test -v ./ast 

test-parser:
	go test -v ./parser 

clean:
	rm -rf ./build