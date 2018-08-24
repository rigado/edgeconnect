default: amd64

amd64:
	go build -o ecapi

arm:
	GOOS=linux GOARCH=arm go build -o ecapi_arm

clean:
	rm ecapi
	rm ecapi_arm

.PHONY: default clean linux arm