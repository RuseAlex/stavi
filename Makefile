MAIN_PACKAGE_PATH := ./cmd/web
BINARY_NAME := server

## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## tidy: format code and tidy modfile
tidy:
	go fmt ./...
	go mod tidy -v

## test: test all the code and show the coverage
test:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## run: runs the application
run:
	go run ${MAIN_PACKAGE_PATH}

## build: builds the application
build:
	go build -o=/tmp/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

tidy:
	go fmt ./...
	go mod tidy -v

