# Communication With GRPC
A simple sample project to test GRPC types of communication

In `client.go`: 
- The method `AddUser` refers to a simple request usage;
- The method `AddUserStream` refers to a server stream request;
- The method `AddUsers` refers to a client stream request;
- The method `AddUserBiStream` refers to a bidirectional stream request;

### Generate Proto
```shell
protoc --proto_path=./proto proto/*.proto --go_out=pb --go-grpc_out=pb
```
### Running Server And Client
```shell
go run cmd/server/server.go
go run cmd/client/client.go
```

