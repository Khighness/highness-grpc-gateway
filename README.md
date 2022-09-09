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
    --grpc-gateway_out=. \
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
20:27:42 grpc-service | Starting grpc-service on port 5000
20:27:42 grpc-gateway | Starting grpc-gateway on port 5100
20:27:42 grpc-service | 2022-09-08T20:27:42.702+0800 [INFO] grpc-service/main.go:42 GRPC service is serving at 0.0.0.0:10010
20:27:42 grpc-gateway | 2022-09-08T20:27:42.727+0800 [INFO] grpc-gateway/main.go:61 GRPC gateway is serving at 0.0.0.0:10020
```

CURL Test:
```shell
$ curl -X GET 'http://127.0.0.1:10020/v1/hello/K/Highness?timestamp=1662551788'
$ curl -X GET 'http://127.0.0.1:10020/v1/bye/K/Highness?timestamp=1662551788'
$ curl -X POST 'http://127.0.0.1:10020/v2/hello' -d '{"first_name":"K", "last_name":"Highness", "timestamp":1662551788}'                                
```