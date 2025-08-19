### gRPC client & server 

-----

Prerequisites:

```sh
brew install go # latest 2 versions 
brew install protoc # protocol buffer compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest # go plugins for the protocol buffer
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin" # update your PATH so that the protoc compiler can find the plugins
```

-----

Largely copied from: https://grpc.io/docs/languages/go/quickstart/

Run the server: 

```sh
go run . --serve
```

Run the client: 

```sh
go run . --client --name=bob
```

Should see: 

```
2025/08/19 18:16:43 server listening at 127.0.0.1:50051
```

and 

```
2025/08/19 18:14:15 Greeting: Hello bob
```

-----

The `./helloworld/helloworld_grpc.pb.go` & `./helloworld/helloworld.pb.go` is 
generated from the `./helloworld/helloworld.proto` file by running: 

```sh
export PATH="$PATH:$(go env GOPATH)/bin" 
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
helloworld/helloworld.proto
```