INSTALL_DIR = $(HOME)/.local/bin
BINARY = compatdata-humanizer
VERSION := $(shell git tag --points-at HEAD | grep . || git rev-parse --short=8 HEAD)

build:
	go build \
		-ldflags "-s -X main.version=$(VERSION)" -o $(BINARY) \
		./cmd/compatdata-humanizer/compatdata-humanizer.go

run:
	go run \
		-ldflags "-X main.version=$(VERSION)" \
		./cmd/compatdata-humanizer/compatdata-humanizer.go

install: build
	cp $(BINARY) "$(INSTALL_DIR)"

uninstall:
	rm "$(INSTALL_DIR)/$(BINARY)"
