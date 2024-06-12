MAKEFLAGS += -j2

tailwind:
	npx tailwindcss -i ./cmd/server/static/input.css -o ./cmd/server/static/styles.css -c ./cmd/server/tailwind.config.js -m --watch

air:
	cd ./cmd/server/ && air

server: build_server
	./bin/server

build: build_server build_cli

build_server:
	go build -o ./bin/server ./cmd/server/
build_cli:
	go build -o ./bin/letter_generator ./cmd/cli/
