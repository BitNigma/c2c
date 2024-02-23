.PHONE: build 

build : 
		go build -o ./p2p cmd/main.go

.PHONE : test
test :
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build