build:
	@templ fmt .
	@templ generate
	@go mod tidy
	@go fmt ./...	
	@go build -o bin/templ-starter cmd/*.go

dev:
	@make build
	@ENV=dev ./bin/templ-starter

watch:
	@make build
	@air 

test:
	@docker compose down --remove-orphans
	@make build
	@docker compose up -d
	@cd tests && npx playwright test

docker:
	@docker compose down --remove-orphans
	@make build
	@docker-compose up
prod:
	@make build
	@flyctl deploy --dockerfile Dockerfile.prod

