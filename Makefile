all: gotool
	@go build -o analyst
gotool:
	gofmt -w .
	go mod tidy
