# This command will run the tests for the project
.PHONY tests:
tests:
	go test -v ./...

# This command will generate a coverage report for the project
.PHONY: coverage
coverage:
	go test -cover -coverprofile=coverage.out ./...

# This command will display the coverage report in an HTML file
.PHONY: coverage-html
coverage-html: coverage
	go tool cover -html="coverage.out"

# This command will display the coverage report in the terminal
.PHONY: coverage-total
coverage-total: coverage
	go tool cover -func="coverage.out"

# This command will run the linter for the project
.PHONY: linter
linter:
	staticcheck ./...