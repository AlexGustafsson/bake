# Disable echoing of commands
MAKEFLAGS += --silent

# Add build-time variables
PREFIX := $(shell go list ./internal/version)
VERSION := v0.1.0
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null)
GO_VERSION := $(shell go version)
COMPILE_TIME := $(shell LC_ALL=en_US date)

BUILD_VARIABLES := -X "$(PREFIX).Version=$(VERSION)" -X "$(PREFIX).Commit=$(COMMIT)" -X "$(PREFIX).GoVersion=$(GO_VERSION)" -X "$(PREFIX).CompileTime=$(COMPILE_TIME)"
BUILD_FLAGS := -ldflags '$(BUILD_VARIABLES)'

server_source := $(shell find . -type f -name '*.go')

# Force macOS to use clang
# https://gcc.gnu.org/bugzilla/show_bug.cgi?id=93082
# https://bugs.llvm.org/show_bug.cgi?id=44406
# https://openradar.appspot.com/radar?id=4952611266494464
ifeq ($(shell uname),Darwin)
	CC=clang
endif

.PHONY: help build release build-specified tools vscode nano prism format lint test install-tools clean

# Produce a short description of available make commands
help:
	pcregrep -Mo '^(#.*\n)+^[^# ]+:' Makefile | sed "s/^\([^# ]\+\):/> \1/g" | sed "s/^#\s\+\(.\+\)/\1/g" | GREP_COLORS='ms=1;34' grep -E --color=always '^>.*|$$' | GREP_COLORS='ms=1;37' grep -E --color=always '^[^>].*|$$'

# Build for the native platform
build: build/bake build/bagels

# Package for release
release: tools
	GOOS=darwin GOARCH=amd64 $(MAKE) build-specified
	GOOS=darwin GOARCH=arm64 $(MAKE) build-specified
	GOOS=linux GOARCH=amd64 $(MAKE) build-specified
	GOOS=linux GOARCH=arm64 $(MAKE) build-specified
	GOOS=windows GOARCH=amd64 $(MAKE) build-specified
	GOOS=windows GOARCH=arm64 $(MAKE) build-specified

# Format Go code
format: $(server_source) Makefile
	gofmt -l -s -w .

# Lint Go code
lint: $(server_source) Makefile
	golint .

# Vet Go code
vet: $(server_source) Makefile
	go vet ./...

# Test Go code
test: $(server_source) Makefile
	go test -v ./...

# Build for the native platform
build/bake: $(server_source) Makefile
	go generate ./...
	go build $(BUILD_FLAGS) -o $@ cmd/bake/*.go

# Build for the native platform
build/bagels: $(server_source) Makefile
	go generate ./...
	go build $(BUILD_FLAGS) -o $@ cmd/bagels/*.go

# Build for the specified platform
build-specified: build/bake-$(GOOS)-$(GOARCH) build/bagels-$(GOOS)-$(GOARCH)
	tar -czf build/bake-$(GOOS)-$(GOARCH).tgz build/bake-$(GOOS)-$(GOARCH)
	tar -czf build/bagels-$(GOOS)-$(GOARCH).tgz build/bagels-$(GOOS)-$(GOARCH)

# Build for the specified platform
build/bake-$(GOOS)-$(GOARCH):
	go generate ./...
	go build $(BUILD_FLAGS) -o $@ cmd/bake/*.go

# Build for the specified platform
build/bagels-$(GOOS)-$(GOARCH):
	go generate ./...
	go build $(BUILD_FLAGS) -o $@ cmd/bagels/*.go

install-tools:
	cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

# Build tools
tools: vscode nano prism

# Build the vscode tool
vscode:
	$(MAKE) -C tools/vscode
	mkdir -p build
	cp tools/vscode/bake-lsp*.vsix build

# Copy the nanorc file to the build output
nano:
	mkdir -p build
	cp tools/nano/bake.nanorc build

# Build the PrismJS grammar
prism:
	$(MAKE) -C tools/prism
	mkdir -p build
	cp tools/prism/build/* build/

# Clean all dynamically created files
clean:
	rm -rf ./build &> /dev/null || true
