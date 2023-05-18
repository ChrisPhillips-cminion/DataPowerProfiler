BINARY_NAME=DataPower-profiler

build:
	mkdir build
	GOARCH=arm64 GOOS=darwin go build -o build/${BINARY_NAME}-darwin-arm64 main.go
	GOARCH=amd64 GOOS=darwin go build -o build/${BINARY_NAME}-darwin-amd64 main.go
	GOARCH=amd64 GOOS=linux go build -o build/${BINARY_NAME}-linux-amd64 main.go
	GOARCH=amd64 GOOS=windows go build -o build/${BINARY_NAME}-windows-amd64.exe main.go

clean:
	go clean
	rm -rf build
