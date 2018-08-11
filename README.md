# USB-BLDC-TOOL

Pre-reqs:

Working Go environment [intalled](https://golang.org/doc/install) 
*Not sure the minimum dependancy, as of development we are using 1.10.3, the ubuntu package in 16.04 usses 1.6 which does not work properly. Ubuntu 18.04 uses 1.10.1 and is ok (Same with Linux Mint 19).*

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
