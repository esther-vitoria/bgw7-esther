test:
	- go test ./...

coverage:
	- go test -cover ./... -coverprofile=coverage.out

coverage-html:
	- go tool cover -html=coverage.out

coverage-func:
	- go tool cover -func=coverage.out

lint:
	- staticcheck ./...