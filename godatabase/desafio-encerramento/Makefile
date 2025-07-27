default: dev

start:
	go run ./cmd/main.go
dev:
	air

up-build:
	docker-compose up --build
up:
	docker-compose up -d

down:
	docker-compose down

down-v:
	docker-compose down -v

logs:
	docker-compose logs -f
	
ps:
	docker-compose ps

cover:
	go test -coverprofile=coverage.out -cover ./...

cover-html:
	go tool cover -html=coverage.out
