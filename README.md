# Basic-API-with-gRPC-and-Protobuf
##### this app sends a request(two numbers a and b) from the client, the server get the request, add and multiply the numbers and sends the response back to the client.


## prerequisite
##### install protocol buffer compiler from [here](https://developers.google.com/protocol-buffers/) or use [snap](https://snapcraft.io/protobuf) for linux
##### check if protoc is successfully installed 
    protoc --version
##### add third_party libraries : 
1. get the latest release from [here](https://github.com/protocolbuffers/protobuf/releases).
2. create third_party folder in your project.
3. copy google folder from protoc-*/include to third_party folder.
#### gRPC package for golang 
    go get -u google.golang.org/grpc
#### protobuf packages for golang 
    go get -u github.com/golang/protobuf/protoc-gen-go
    
## Generate Proto files : 
    protoc --proto_path=proto --proto_path=third_party --go_out=plugins=grpc:proto service.proto