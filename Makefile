INSTALL_DIR = $(HOME)/.local/bin
DIST_DIR = dist
BINARY = compatdata-humanizer
BINARY_AMD64 = $(BINARY).amd64
BINARY_ARM64 = $(BINARY).arm64
VERSION := $(shell git tag --points-at HEAD | grep . || git rev-parse --short=8 HEAD)
OWN_ARCH := $(shell go env GOARCH)

build:
	env GOOS=linux GOARCH=amd64 go build \
		-ldflags "-s -X main.version=$(VERSION)" -o "$(DIST_DIR)/$(BINARY_AMD64)" \
		./cmd/compatdata-humanizer/compatdata-humanizer.go
	env GOOS=linux GOARCH=arm64 go build \
		-ldflags "-s -X main.version=$(VERSION)" -o "$(DIST_DIR)/$(BINARY_ARM64)" \
		./cmd/compatdata-humanizer/compatdata-humanizer.go

run:
	go run \
		-ldflags "-X main.version=$(VERSION)" \
		./cmd/compatdata-humanizer/compatdata-humanizer.go

install: build
	cp "$(DIST_DIR)/$(BINARY).$(OWN_ARCH)" "$(INSTALL_DIR)/$(BINARY)"

uninstall:
	rm "$(INSTALL_DIR)/$(BINARY)"
