MAKEFLAGS += -j2

tailwind:
	npx tailwindcss -i ./cmd/server/static/input.css -o ./cmd/server/static/styles.css -c ./cmd/server/tailwind.config.js -m --watch

air:
	ln -s ./cmd/server/static/ ./bin/static
	cd ./cmd/server/ && air

server: build_server
	./bin/server

# TODO: make working dir consistent
build: clean_bin build_server build_cli copy_files

build_server:
	go build -o ./bin/server ./cmd/server/

build_cli:
	go build -o ./bin/letter_generator ./cmd/cli/

copy_files:
	cp -r ./cmd/server/static/ ./bin/static/

clean_bin:
	cd ./bin/ && rm -rf *

