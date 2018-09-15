default: amd64

all: amd64 arm

amd64:
	mkdir -p build
	go build -o build/ecapi

arm:
	mkdir -p build
	GOOS=linux GOARCH=arm go build -o build/ecapi_arm

help:
	@echo Available targets:
	@echo amd64
	@echo arm
	@echo all
	@echo "default (amd64)"
	@echo clean

clean:
	rm -rf build

.PHONY: default clean linux arm all help