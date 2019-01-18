test:
	go test $(shell go list ./... | grep -v /examples/ ) -covermode=count

test-race:
	go test -race $(shell go list ./... | grep -v /examples/ )

coverage:
	go test $(shell go list ./... | grep -v /examples/ ) -covermode=count -coverprofile=coverage.out && go tool cover -func=coverage.out

coverage-html:
	go test $(shell go list ./... | grep -v /examples/ ) -covermode=count -coverprofile=coverage.out && go tool cover -html=coverage.out

lint: 
	golint -set_exit_status $(shell (go list ./... | grep -v /vendor/))

mocks:
	mockgen -package nfs -source ./nfs/nfs.go FileSystem > ./nfs/fs_mock.go
	mockgen -package nhttp -source ./nhttp/nhttp.go Client > ./nhttp/client_mock.go	
	mockgen -package nio io Reader,Writer,Closer,ReadWriter,ReadCloser,ReadWriteCloser  > ./nio/io_mock.go
	mockgen -package ntime -source ./ntime/ntime.go Timer > ./ntime/timer_mock.go

.PHONY: test test-race coverage coverage-html lint
