build:
	@templ fmt .
	@templ generate
	@go mod tidy
	@go fmt ./...	
	@go build -o bin/templ-starter cmd/*.go
dev:
	@templ fmt .
	@templ generate
	@go mod tidy
	@go fmt ./...	
	@go build -o bin/templ-starter cmd/*.go
	@ENV=dev ./bin/templ-starter
prod:
	@templ fmt .
	@templ generate
	@go mod tidy
	@go fmt ./...
	@go build -o bin/templ-starter cmd/*.go
	@flyctl deploy