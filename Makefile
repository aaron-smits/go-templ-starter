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
watch:
	@templ fmt .
	@templ generate
	@go mod tidy
	@go fmt ./...	
	@go build -o bin/templ-starter cmd/*.go
	@air 
prod:
	@templ fmt .
	@templ generate
	@go mod tidy
	@go fmt ./...
	@go build -o bin/templ-starter cmd/*.go
	@flyctl deploy --dockerfile Dockerfile.prod
docker:
	@templ fmt .
	@templ generate
	@go mod tidy
	@go fmt ./...	
	@go build -o bin/templ-starter cmd/*.go
	@docker-compose up
