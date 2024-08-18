test: # run tests
	go test ./...

build: # build lib
	go build -o bin/wiserwithbit wiserwithbit/wiserwithbit.go
