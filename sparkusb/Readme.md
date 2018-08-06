commands.proto contains the message details to control the spark remotely. To build:

**GO**

Install protoc compiler then:

`go get -u github.com/golang/protobuf/protoc-gen-go`

`protoc -I=./ --go_out=./  ./commands.proto`
