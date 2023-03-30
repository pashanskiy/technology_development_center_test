//go:generate mkdir -p ../api
//go:generate find . -type f -name **.proto -exec protoc {} --proto_path=. --go_out=../api --go_opt=paths=source_relative --go-grpc_out=../api --go-grpc_opt=paths=source_relative ;

package proto
