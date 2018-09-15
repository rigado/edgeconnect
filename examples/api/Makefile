default: amd64

all: amd64 arm

amd64:
	mkdir -p build
	go build -o build/ecapi

arm:
	mkdir -p build
	GOOS=linux GOARCH=arm go build -o build/ecapi_arm

clean:
	rm -rf build

.PHONY: default clean linux arm all