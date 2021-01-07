![alt-text-2](https://miro.medium.com/max/722/1*P0z3ortvUH4gZtIyj4al_A.png "title-2")

# Features

- Server Stream
- Client Stream
- Bi-direcional Stream (default)

# Init

```sh
go mod init github.com/mauriciolucas22/gRPC-examples
go get google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go
```

# Run Server

```sh
go run cmd/server/server.go
```

# Run Client

```sh
go run cmd/client/client.go
```

# To anothers streams edit client.go

```go
// AddUser(client)
// AddUserVerbose(client)
// AddUsers(client)
AddUserStreamBoth(client) // DEFAULT > bi-direcional stream
```

# Reduce users count sended by stream in client.go

```go
// Alter number 10 to your counter
for i := 0; i < 10; i++ {
		reqs = append(reqs, &pb.User{
			Id:    strconv.Itoa(i),
			Name:  "Jesus Loves you " + strconv.Itoa(i),
			Email: "@jesus",
		})
	}
```

<!-- go run cmd/server/server.go

go mod init github.com/mauriciolucas22/gRPC-examples

go get google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go

-- Only Mac
brew install protoc-gen-go
brew install protoc-gen-go-grpc -->

<!-- protoc --proto_path=proto proto/_.proto --go_out=pb -->

<!-- protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb -->
