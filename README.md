# Rigado Edge Connect API Example

This repository contains a Golang based example for interacting with the edge connect APIs.

## API Documentation

[Edge Connect API](https://docs.rigado.com/projects/edge-connect-api/en/latest/)

## Available Functionality

- [X] Broadcasts 
- [X] Get/Set Radio Modes
- [X] Tau board support
- [X] BLE Scanning
- [X] Firmware Programming

## Demo Binary

A demo application is provided in examples/api. This application can perform most API operations.

## Building Demo

Install **go** if you don't have it. [Get Go here](https://golang.org/dl/)

** Make sure to set `$GOPATH` **

Clone the repo into `$GOPATH/src/github.com/rigado`:

```
cd $GOPATH
mkdir -p src/github.com/rigado
cd src/github.com/rigado
git clone https://github.com/rigado/edgeconnect.git
```

Install **dep** if you don't have it. 

On MacOS w/ [Homebrew](https://brew.sh/):
```
brew install dep
brew upgrade dep
```

On other platforms:
```
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

Then install dependencies by running:

`dep ensure`

In the **edgeconnect/examples/api** directory, run:

`make clean && make`

The binary can be builds for either `amd64` or `arm` architectures. See the `Makefile`.

The binaries are located in the **build** folder and can be installed.

### Running the API demo on a gateway

```
make arm
scp build/ecapi_arm <gateway_ip>:~
ssh <gateway_ip>
chmod +x ecapi_arm (may not be necessary)
sudo ./ecapi (options)
```