.PHONE: build 

build : 
		go build -v ./cmd/apiserver

.PHONE : test
test :
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build