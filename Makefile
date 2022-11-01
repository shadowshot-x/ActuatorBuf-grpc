generate-proto:
	protoc --go_out=. ./protobufs/*.proto

install-protoc-go:
	brew install protobuf
	export GOPATH=$$HOME/go
	export PATH=$$PATH:$$GOPATH/bin
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

delete-protoc:
	brew uninstall protobuf