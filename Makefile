.PHONY: pb
pb:
	protoc -I=. -I=${GOPATH}/src -I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. --js_out=import_style=commonjs,binary:. ./pbtrial/types.proto