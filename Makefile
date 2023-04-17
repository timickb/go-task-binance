.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: server
server:
	go build -v -o "./bin/server" ./cmd/server

.PHONY: client
client:
	go build -v -o "./bin/client" ./cmd/client
