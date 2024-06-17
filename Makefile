MAKEFLAGS += -j2

tailwind:
	npx tailwindcss -i ./cmd/server/static/input.css -o ./cmd/server/static/styles.css -c ./cmd/server/tailwind.config.js -m --watch

air: clean_bin
	ln -s ./cmd/server/static/ ./bin/static
	cd ./cmd/server/ && air

export_server: build_server clean_export
	mkdir output
	mkdir ./output/bin
	mkdir ./output/templates
	cp ./bin/server ./output/bin/server
	cp -r ./cmd/server/static/ ./output/static
	cp -r ./templates/ ./output/templates/

clean_export:
	rm -rf ./output/

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
	rm -rf ./bin/*

