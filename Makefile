.PHONY: init
init:
	mkdir -p dist

.PHONY: build
build:
	go build -o bin/kanna \
		./cmd/kanna

.PHONY: dockerbuild
dockerbuild:
	docker build . -t kanna
