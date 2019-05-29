# Go parameters
GOARCH=amd64
GOBUILD=go build -o
GOTEST=go test -v -race ./...
GOLINT=golint ./...
GOCLEAN=go clean ./...
GOINSTALL=go install cmd/*

# Build parameters
MAIA_VER=1.0.0
MAIA_LINUX=env GOOS=linux GOARCH=$(GOARCH)
MAIA_MAC=env GOOS=darwin GOARCH=$(GOARCH)
MAIA_WIN=env GOOS=windows GOARCH=$(GOARCH)

# Maia parameters
MAIA_DIR=./cmd/
MAIA_BIN=maia

# Maia build for macOS and Linux
LINUX=$(MAIA_LINUX) $(GOBUILD) $(MAIA_BIN)_v$(MAIA_VER)-linux
MAC=$(MAIA_MAC) $(GOBUILD) $(MAIA_BIN)_v$(MAIA_VER)-mac
WIN=$(MAIA_WIN) $(GOBUILD) $(MAIA_BIN)_v$(MAIA_VER)-win

MAIA_BUILD_LINUX=$(LINUX) $(MAIA_DIR)
MAIA_BUILD_MAC=$(MAC) $(MAIA_DIR)
MAIA_BUILD_WIN=$(WIN) $(MAIA_DIR)

all: clean test build install
build: build-linux build-mac build-win
build-mac:
	$(MAIA_BUILD_MAC)
build-linux:
	$(MAIA_BUILD_LINUX)
build-win:
	$(MAIA_BUILD_WIN)
install:
	$(GOINSTALL)
test:
	$(GOTEST)
lint:
	$(GOLINT)
clean:
	$(GOCLEAN)
	rm -f maia_v*
