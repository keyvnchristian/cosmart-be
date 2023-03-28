## build: Compile server binaries.
build:
	go build -o cosmart-be ./main.go

## build-linux: Compile binaries for linux.define
build-linux:
	GOARCH=amd64 GOOS=linux go build ./main.go

## run: Execute server.
run:
	./cosmart-be

## test: run test and show coverage.
test:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

start: build run