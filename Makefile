MAKEFLAGS += -j2

tailwind:
	npx tailwindcss -i ./cmd/server/static/input.css -o ./cmd/server/static/styles.css -c ./cmd/server/tailwind.config.js -m --watch

air: clean_bin
	ln -s ./cmd/server/static/ ./bin/static
	cd ./cmd/server/ && air

export_server: build_server clean_export
	mkdir letter_generator
	mkdir ./letter_generator/bin
	mkdir ./letter_generator/templates
	cp ./bin/server ./letter_generator/bin/server
	cp -r ./cmd/server/static/ ./letter_generator/static
	cp -r ./templates/ ./letter_generator/templates/


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
	cp -r ./cmd/server/static/ ./static/

clean_bin:
	rm -rf ./bin/*

