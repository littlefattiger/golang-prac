#### run below under features
protoc --proto_path=./proto --go_out=./proto --go_opt=paths=source_relative --go-grpc_out=./proto --go-grpc_opt=paths=source_relative --grpc-gateway_out=./proto --grpc-gateway_opt=paths=source_relative ./proto/gateway/gateway.proto
protoc --proto_path=./proto --go_out=./proto --go_opt=paths=source_relative ./proto/gateway/gateway.proto
protoc --proto_path=./proto --go_out=./proto --go_opt=paths=source_relative --go-grpc_out=./proto --go-grpc_opt=paths=source_relative ./proto/gateway/gateway.proto

Install GRPC GateWay:
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
or
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
