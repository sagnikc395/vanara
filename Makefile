build:
	go build main.go -o dist/monkey
clean:
	rm -rf dist/
run:
	go run main.go
test:
	go test .