build:
	@npm install
	@npx tailwindcss -i assets/tailwind.css -o assets/dist/styles.css
	@templ fmt .
	@templ generate
	@go mod tidy
	@go fmt ./...	
	@go build -o bin/templ-starter cmd/main.go
run:
	@npx tailwindcss -i assets/tailwind.css -o assets/dist/styles.css
	@templ fmt .
	@templ generate
	@go mod tidy
	@go fmt ./...	
	@go run cmd/main.go