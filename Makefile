# Go parameters
GOARCH=amd64
GOBUILD=go build -o
GOTEST=go test ./cmd/
GOCLEAN=go clean ./cmd/

# Build parameters
MAIA_LINUX=env GOOS=linux GOARCH=$(GOARCH)
MAIA_MAC=env GOOS=darwin GOARCH=$(GOARCH)

# Maia parameters
MAIA_DIR=./cmd/
MAIA_BIN=maia

# Maia build for macOS and Linux
LINUX=$(MAIA_LINUX) $(GOBUILD) $(MAIA_BIN)_linux
MAC=$(MAIA_MAC) $(GOBUILD) $(MAIA_BIN)_mac
MAIA_BUILD_LINUX=$(LINUX) $(MAIA_DIR)
MAIA_BUILD_MAC=$(MAC) $(MAIA_DIR)

all: clean test build
build: 
	$(MAIA_BUILD_LINUX)
	$(MAIA_BUILD_MAC)
test:
	$(GOTEST)
clean:
	$(GOCLEAN)
	rm -f *_linux
	rm -f *_mac
	



