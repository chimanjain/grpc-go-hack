
# grpc-go-hack

grpc-go-hack server and client demonstrate how to use grpc go libraries to perform unary, server streaming RPCs.

Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.

See the definition of the project service in `project/project.proto`. The definition is architectured against [Airbrake-Projects](https://airbrake.io/docs/api/#list-projects-v4).

## Run the code

To compile and run the server, assuming you are in the root of the `grpc-go-hack`, simply:

```sh
go run project_server/project_server.go
```

Likewise, to run the client, open another console as the same root i.e. `grpc-go-hack` and run:

```sh
go run project_client/project_client.go
```

## To make changes

The code in this repo is already compiled but if you want to make changes to proto file, you will have
to complile it again.

If you are using a MacOS environment, install brew and follow the steps:

```sh
brew install protobuf
brew install protoc-gen-go
brew install protoc-gen-go-grpc
```

After installing above formulas, simple make **Changes** in `project.proto` file, run the following command in the root directory i.e. `grpc-go-hack`:

```sh
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative project/project.proto
```
