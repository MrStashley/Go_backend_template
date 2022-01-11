# i'm a c programmer at heart :p
# i'm not sure how build automation is normally done in go, if you see this and you know better than me, feel free to
# let me know by making an issue or sending me an email

all: server_main

server_main: ./*
	mkdir -p bin/
	go env -w GOBIN=${HOME}/personal/go/simple-go-server/bin
	go install simple-go-server
	go env -u GOBIN

