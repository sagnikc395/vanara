build:
	go build  -o dist/monkey
clean:
	rm -rf dist/
run:
	go run main.go
test-lexer:
	go test ./lexer 
test-token:
	go test ./token 