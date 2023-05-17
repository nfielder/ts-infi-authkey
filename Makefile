BINARY_NAME=ts-infi-authkey
MAIN_FILE=cmd/ts-infi-authkey/main.go

build:
	go build -o bin/${BINARY_NAME} ${MAIN_FILE}

compile:
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux ${MAIN_FILE}
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin ${MAIN_FILE}
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows ${MAIN_FILE}

run:
	go run ${MAIN_FILE}

clean:
	go clean
	rm ./bin/${BINARY_NAME}*

fmt:
	go fmt ./...

lint:
	golangci-lint run
