# SPARK Max Interface

All messages details to control the spark remotely, as well as all defined types on the SPARK MAX are defined by protocol buffers.

**SPARK-MAX-Command.proto** - Commands and request/response types

**SPARK-MAX-Types.proto** - Static types and parameter definitions

## GO

Install protoc compiler then the following commands from this folder:

`go get -u github.com/golang/protobuf/protoc-gen-go`

`protoc -I=./ --go_out=./  ./SPARK-MAX*.proto`

**Generating Docs**

Docs for the protocol buffers API was generated with [protoc-gen-doc](https://github.com/pseudomuto/protoc-gen-doc)

Protoc version must be > 3.3.0 if not building via docker.

Command is:

`protoc --doc_out=./ --doc_opt=markdown,api.md ./*.proto`
