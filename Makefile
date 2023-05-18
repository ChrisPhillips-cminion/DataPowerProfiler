BINARY_NAME=DataPower-profiler
VERSION=1

build:
	mkdir build
	GOARCH=arm64 GOOS=darwin go build -o build/${BINARY_NAME}-v${VERSION}-darwin-arm64 main.go
	GOARCH=amd64 GOOS=darwin go build -o build/${BINARY_NAME}-v${VERSION}-darwin-amd64 main.go
	GOARCH=amd64 GOOS=linux go build -o build/${BINARY_NAME}-v${VERSION}-linux-amd64 main.go
	GOARCH=amd64 GOOS=windows go build -o build/${BINARY_NAME}-v${VERSION}-windows-amd64.exe main.go

clean:
	go clean
	rm -rf build
