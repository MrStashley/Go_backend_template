all: server_main

server_main: ./*
	mkdir -p bin/
	go env -w GOBIN=${HOME}/personal/go/simple-go-server/bin
	go install simple-go-server
	go env -u GOBIN

