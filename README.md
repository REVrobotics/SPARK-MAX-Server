# SPARK MAX Backend Tool

## Building

Steps to build the sparkusb server/cli tool

1) Install and setup git
2) Install and setup a go environemnt (minimum 1.7, recommended 1.10+)
3) On linux instatll zeromq `sudo apt-get install libzmq3-dev` (windows includes pre-built binaries)
4) Clone this repo under `$GOPATH/src/github.com/willtoth/`
5) Run `make deps` or manually get go dependencies
6) run `make` to build the project

## Pre-reqs:

- Working Go environment [intalled](https://golang.org/doc/install) 
*Not sure the minimum dependancy, as of development we are using 1.10.3, the ubuntu package in 16.04 usses 1.6 which does not work properly. Ubuntu 18.04 uses 1.10.1 and is ok (Same with Linux Mint 19). Successfully built using go 1.7 on a raspberry pi zero w.*

- Make build system

## Go dependencies:

To install the below run `make deps`

*Manual package list:*
```
go get github.com/pebbe/zmq4
go get -u github.com/spf13/cobra/cobra
go get go.bug.st/serial.v1
go get github.com/tarm/serial
go get -u github.com/golang/protobuf/protoc-gen-go
```
## Structure

This application can run both as a command line utility or remote command server to interface to the SPARK MAX. It includes the following packages:

- sparkmax - Protocol buffer definitions of all types and messages and hardware interface
- spark0mq - Includes the ZeroMQ server for network communication
- cmd - All defined commands for either cli tool, zmq server, or both depending on the command

## Front-end Client

Moved to its own repo [SPARK-Max-Client](https://github.com/REVrobotics/SPARK-MAX-Client)

## Known Issues

On Linux machines with ModemManager installed, the device may show up as busy for ~15 seconds while modem manager tries to decide if its a modem. To disable this add a custom UDEV rule, [here](https://linux-tips.com/t/prevent-modem-manager-to-capture-usb-serial-devices/284) is a detailed description. **VID** = 0483 **PID** = 5740

For example:

*/etc/udev/rules.d/99-ttyacms.rules* has the line:

`ATTRS{idVendor}=="0483" ATTRS{idProduct}=="5740", ENV{ID_MM_DEVICE_IGNORE}="1"`