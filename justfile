PROJECT_NAME := "alchemist"
GO_BIN := `go env GOPATH` / "bin"
DIST_PATH := "dist"

# Build the project
build:
    go build -o {{DIST_PATH}}/{{PROJECT_NAME}}

# Clean the project
clean:
    rm -rf vendor {{DIST_PATH}}

# Get test coverage and open it in a browser
coverage: 
    go clean -testcache && go test ./... -coverprofile=cover.out && go tool cover -html=cover.out

# Install globally from source
install:
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b {{GO_BIN}} v2.6.1
    go mod vendor

# Lint the project
lint:
   {{GO_BIN}}/golangci-lint run

# Test the project
test:
    go clean -testcache && go test ./...
