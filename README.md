# protobuf-sample

## install protoc

1. install brotobuf-all-x.x.x from [protobuf](https://developers.google.com/protocol-buffers/docs/downloads)
2. unzip zip
3. cd protobuf-x.x.x
4. ./configure
5. make
6. make check
6. sudo make install

## install golang protocol buffer plugin

`go get -u github.com/golang/protobuf/protoc-gen-go`

## comple .proto

`protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto`


for more info :
https://developers.google.com/protocol-buffers/docs/proto3