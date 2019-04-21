# Gin + MySQL + gRPC + protobuf

```
protoc --go_out=plugins=grpc:. *.proto
protoc --go_out=. *.proto
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. ./proto/hello/hello.proto
```

```
consul agent -dev
go run api/api.go --registry=consul
go run ms1/ms.go --registry=consul
go run ms2/ms.go --registry=consul
go run ms3/ms.go --registry=consul
curl http://localhost:34895/greeter/asim
```
