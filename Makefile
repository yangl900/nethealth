all: build-linux build-windows

build-linux:
	GOARCH=amd64 GOOS=linux go build -o nethealth .

build-windows:
	GOARCH=amd64 GOOS=windows go build -o nethealth.exe .
