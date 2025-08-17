default: start

start:
	go run ./cmd/main.go

cover:
	go test -coverprofile=coverage.out -cover ./...

cover-html:
	go tool cover -html=coverage.out

cover-filtered:
	./scripts/cover-filtered.sh

cover-filtered-total:
	go tool cover -func=coverage_filtered.out