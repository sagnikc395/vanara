build:
	mkdir -p ./dist 
	go build -o ./dist/vanara 

run:
	./dist/vanara 

tests:
	go test -v ./...

# run the tests for the specific parts 
test-lexer:
	go test -v ./lexer 

test-ast:
	go test -v ./ast 

test-parser:
	go test -v  ./parser 

clean:
	rm -rf ./build 
