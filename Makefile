server: build_server
	./bin/server

build: build_server build_cli

build_server:
	go build -o ./bin/server ./cmd/server/main.go
build_cli:
	go build -o ./bin/letter_generator ./cmd/cli/main.go
