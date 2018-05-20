VERSION := $(shell git describe --tags)

linux:
	go build -o ./build/bot-linux -ldflags="-X main.version=${VERSION}" ./main.go

mac:
	GOOS=darwin GOARCH=amd64 go build -o ./build/bot-mac -ldflags="-X main.version=${VERSION}" ./main.go
	
windows:
	GOOS=windows GOARCH=386 go build -o ./build/bot-windows.exe -ldflags="-X main.version=${VERSION}" ./main.go

clean:
	rm -rf ./build

copyfiles:
	cp config.template.json ./build/config.json

zip:
	zip -r build.zip build

all: linux mac windows copyfiles
