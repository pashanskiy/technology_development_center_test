init: 
	go mod tidy

run:
	go run ./cmd/main.go

go-install:
	go get -u google.golang.org/protobuf/proto 2>/dev/null && go install google.golang.org/protobuf/proto
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

go-gen: go-install
	go generate ./...
	go get -u -t -v ./... || :
	go mod tidy
