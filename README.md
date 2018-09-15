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

A demo application is provided in `examples/api`. This application can perform most API operations.

## Building the Demo

### Prerequisites

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

### Build demo binaries

In the **edgeconnect/examples/api** directory, run:

`make clean && make` (defaults to amd64)

The binary can be builds for either `amd64` (default) or `arm` architectures. Run `make help` for available commands.

The binaries are located in the **build** folder.

### Running the API demo on a gateway

```
cd examples/api
make arm
scp build/ecapi_arm <gateway_ip>:~
ssh <gateway_ip>
chmod +x ecapi_arm (may not be necessary)
sudo ./ecapi (options)
```