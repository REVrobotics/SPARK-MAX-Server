Requires go [zeromq4](https://github.com/pebbe/zmq4)

To build:

1) Compile zeromq 4 using visual studio (target is DynRelease x64)
2) Install a gcc toolset for windows, in this case mingw-64 - make sure gcc is in PATH
3) Set the following flags on the command line (Recommend a path without spaces for source)

```
set CGO_CFLAGS=-I<path to zeromq> -g -O2
set CGO_LDFLAGS=-L<path to build library> -g -O2
go get github.com/pebbe/zmq4
```

Example: ` & { $env:CGO_CFLAGS='-g -O2 -IC:\Users\Will\Src\zeromq-4.2.3\include'; $env:CGO_LDFLAGS='-g -O2 -LC:\Users\Will\Src\zeromq-4.2.3\bin\x64\Release\v141\dynamic'; go get github.com/pebbe/zmq4}`

Same applies when building

Example: ` & { $env:CGO_CFLAGS='-g -O2 -IC:\Users\Will\Src\zeromq-4.2.3\include'; $env:CGO_LDFLAGS='-g -O2 -LC:\Users\Will\Src\zeromq-4.2.3\bin\x64\Release\v141\dynamic'; go build}`

Alternatively, put the built library somewhere accessible by PATH (go/bin?) and the header in a location accessible by 
