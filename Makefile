build:
	@go build -o bin/templ-starter cmd/main.go
run:
	@templ generate
	@go run cmd/main.go