BINDIR	:= $(CURDIR)/bin
BINNAME	?= inspektor

GOBIN         			= $(shell go env GOBIN)
ifeq ($(GOBIN),)
GOBIN         			= $(shell go env GOPATH)/bin
endif
ARCH          			= $(shell uname -p)

# Go Options
PKG         := ./...
TAGS        :=
TESTS       := .
TESTFLAGS   :=
LDFLAGS     := -w
GOFLAGS     :=
CGO_ENABLED ?= 0

# Rebuild the binary if any of these files change
SRC := $(shell find . -type f -name '*.go' -print) go.mod go.sum

.PHONY: all
all: build

.PHONY: build
build: $(BINDIR)/$(BINNAME)

$(BINDIR)/$(BINNAME): $(SRC)
	mkdir -p $(BINDIR)
	CGO_ENABLED=$(CGO_ENABLED) go build -o $(BINDIR) -ldflags $(LDFLAGS) $(PKG)

.PHONY: run
run: build ## Run the binary
	$(BINDIR)/$(BINNAME)

# ------------------------------------------------------------------------------
# dependencies

.PHONY: download
download: ## Download dependencies
	go mod download

.PHONY: tidy
tidy: download ## Update go.mod and go.sum
	go mod tidy

# ------------------------------------------------------------------------------
# docker

.PHONY: docker-build
docker-build: ## Build docker image
	docker build -t inspektor .

# ------------------------------------------------------------------------------

.PHONY: clean
clean: ## Remove build artifacts
	rm -rf '$(BINDIR)'

.PHONY: help
help: ## Display this help screen
	@echo "Usage: make [target] ..."
	@echo
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort