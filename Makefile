tidy:
	go mod tidy

# Coding style static check.
lint: tidy
	# keep this in sync with .github/workflows/golangci-lint.yml
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.6.1
	go mod tidy
	golangci-lint run

test: tidy
	export GOARCH='amd64' && go test ./...
	export GOARCH='386' && go test ./...

# target to run all the possible checks; it's a good habit to run it before
# pushing code
check: lint test
	echo "check done"
