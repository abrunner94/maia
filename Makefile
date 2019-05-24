# Go parameters
GOARCH=amd64
GOBUILD=go build -o
GOTEST=go test -v -race ./...
GOCLEAN=go clean ./...

# Build parameters
MAIA_LINUX=env GOOS=linux GOARCH=$(GOARCH)
MAIA_MAC=env GOOS=darwin GOARCH=$(GOARCH)
MAIA_WIN=env GOOS=windows GOARCH=$(GOARCH)

# Maia parameters
MAIA_DIR=./cmd/
MAIA_BIN=maia

# Maia build for macOS and Linux
LINUX=$(MAIA_LINUX) $(GOBUILD) $(MAIA_BIN)_linux
MAC=$(MAIA_MAC) $(GOBUILD) $(MAIA_BIN)_mac
WIN=$(MAIA_WIN) $(GOBUILD) $(MAIA_BIN)_win

MAIA_BUILD_LINUX=$(LINUX) $(MAIA_DIR)
MAIA_BUILD_MAC=$(MAC) $(MAIA_DIR)
MAIA_BUILD_WIN=$(WIN) $(MAIA_DIR)

all: clean test build
build: build-linux build-mac build-win
build-mac:
	$(MAIA_BUILD_MAC)
build-linux:
	$(MAIA_BUILD_LINUX)
build-win:
	$(MAIA_BUILD_WIN)
test:
	$(GOTEST)
clean:
	$(GOCLEAN)
	rm -f *_linux
	rm -f *_mac
	rm -f *_win