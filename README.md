# USB-BLDC-TOOL

**Building**

Steps to build the sparkusb server/cli tool

1) Install and setup git
2) Install and setup a go environemnt (minimum 1.7, recommended 1.10+)
3) Build or install zeromq4
4) 'Go get' all dependancies
5) Clone this repo under `$GOPATH/src/github.com/willtoth/`
6) 'Go build'

#Pre-reqs:

Working Go environment [intalled](https://golang.org/doc/install) 
*Not sure the minimum dependancy, as of development we are using 1.10.3, the ubuntu package in 16.04 usses 1.6 which does not work properly. Ubuntu 18.04 uses 1.10.1 and is ok (Same with Linux Mint 19). Successfully built using go 1.7 on a raspberry pi zero w.*

Pre-build binary for Zeromq:

**Windows**
To build:

1) Compile zeromq 4 using visual studio (target is DynRelease x64)
2) Install a gcc toolset for windows, in this case mingw-64 - make sure gcc is in PATH
3) Set the following flags on the command line (Recommend a path without spaces for source)

```
set CGO_CFLAGS=-I<path to zeromq> -g -O2
set CGO_LDFLAGS=-L<path to build library> -g -O2
go get github.com/pebbe/zmq4
```

Example with powershell: ` & { $env:CGO_CFLAGS='-g -O2 -IC:\Users\Will\Src\zeromq-4.2.3\include'; $env:CGO_LDFLAGS='-g -O2 -LC:\Users\Will\Src\zeromq-4.2.3\bin\x64\Release\v141\dynamic'; go get github.com/pebbe/zmq4}`

Same applies when building

Example with powershell: ` & { $env:CGO_CFLAGS='-g -O2 -IC:\Users\Will\Src\zeromq-4.2.3\include'; $env:CGO_LDFLAGS='-g -O2 -LC:\Users\Will\Src\zeromq-4.2.3\bin\x64\Release\v141\dynamic'; go build}`

Alternatively, put the built library somewhere accessible by PATH (go/bin?) and the header in a location accessible by the compiler

**Linux (Ubuntu)**

`sudo apt-get install libzmq3-dev`

# Go dependencies:

```
go get github.com/pebbe/zmq4
go get -u github.com/spf13/cobra/cobra
go get github.com/willtoth/go-serial
go get -u github.com/golang/protobuf/protoc-gen-go
```

*The serial package is from here*
go.bug.st/serial.v1

With this patch applied: https://patch-diff.githubusercontent.com/raw/bugst/go-serial/pull/33.patch

# sparkusb-client

Pre-reqs:

Working nodejs environment and node package manager (npm)
