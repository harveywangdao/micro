# Service

This is an example of creating a micro service.

## Run the example

Run the service

```shell
consul agent -dev
go run server.go --registry=consul
```

Run the client

```shell
go run client.go --run_client --registry=consul
```

And that's all there is to it.
