## highness-grpc-gateway



## About

The gRPC-Gateway is a plugin of the Google protocol buffers compiler protoc. It reads protobuf service definitions and generates a reverse-proxy server which translates a RESTful HTTP API into gRPC. This server is generated according to the google.api.http annotations in your service definitions.

![architecture](assets/architecture.svg)




## Build

Generate model.pb, grpc.pb, grpc-gateway.pb

```shell
$ protoc -I=. \
    --go_out=.  \
    --go-grpc_out=.  \
    --hello-gateway_out=. \
    ./proto/api/hello.proto
```

Build grpc-service, grpc-gateway

```shell
$ cd cmd/grpc-service && go build -o  grpc-service-application
$ cd ../grpc-gateway && go build -o  grpc-gateway-application
$ cd ../../
```



## Start

> First install [goreman](https://github.com/mattn/goreman), which manages Procfile-based applications.

```shell
$ goreman start
15:33:46 grpc-service | Starting grpc-service on port 5000
15:33:46 grpc-gateway | Starting grpc-gateway on port 5100
15:33:46 grpc-service | [GRPC] 2022/09/07 15:33:46.790663 main.go:36: GRPC service is serving at 0.0.0.0:10010
15:33:46 grpc-gateway | [GATEWAY] 2022/09/07 15:33:46.807039 main.go:46: GRPC gateway is serving at 0.0.0.0:10020
```

CURL Test:
```shell
$ curl -X POST http://127.0.0.1:10020/v1/hello -d '{"name":"KHighness"}'
{"reply":"Hi KHighness"}                      
```