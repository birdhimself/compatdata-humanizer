INSTALL_DIR = $(HOME)/.local/bin
BINARY = compatdata-humanizer

build:
	go build -ldflags "-s" -o $(BINARY) ./cmd/compatdata-humanizer/compatdata-humanizer.go

run:
	go run ./cmd/compatdata-humanizer/compatdata-humanizer.go

install: build
	cp $(BINARY) "$(INSTALL_DIR)"

uninstall:
	rm "$(INSTALL_DIR)/$(BINARY)"
