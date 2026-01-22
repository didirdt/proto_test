# proto_test

A Go project for protobuf code generation and gRPC service definitions. This repository demonstrates how to generate Go types and gRPC service stubs from Protocol Buffer definitions.

# protobuf code generation

Prerequisites:

- Install protoc (https://grpc.io/docs/protoc-installation/)
- Install Go protobuf plugins:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Ensure `$GOPATH/bin` (or `$GOBIN`) is in your `PATH`.

From the repository root run:

```sh

```

This will generate Go types and gRPC stubs in the `protobuff` package.