build:
	@templ fmt .
	@templ generate
	@go mod tidy
	@go fmt ./...	
	@go build -o bin/templ-starter cmd/main.go
run:
	@templ fmt .
	@templ generate
	@go mod tidy
	@go fmt ./...	
	@go run cmd/main.go