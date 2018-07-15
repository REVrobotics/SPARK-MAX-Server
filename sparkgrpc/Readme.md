commands.proto contains the server and message details to control the spark remotely. To build:

**GO**
```protoc -I=./ --go_out=plugins=grpc:./ ./commands.proto