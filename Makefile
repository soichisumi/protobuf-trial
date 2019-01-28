.PHONY: pb
pb:
	protoc -I=. --go_out=. ./pbtrial/types.proto