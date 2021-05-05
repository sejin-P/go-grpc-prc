.PHONY: format
## format: format files
format:
	@go get golang.org/x/tools/cmd/goimports
	goimports -local github.com/sejin-P -w .
	gofmt -s -w .
	go mod tidy
