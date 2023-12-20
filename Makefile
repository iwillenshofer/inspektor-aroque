BINDIR	:= $(CURDIR)/bin
BINNAME	?= inspektor

GOBIN         			= $(shell go env GOBIN)
ifeq ($(GOBIN),)
GOBIN         			= $(shell go env GOPATH)/bin
endif
PROTOC_GEN_GO				= $(GOBIN)/protoc-gen-go
PROTOC_GEN_GO_GRPC	= $(GOBIN)/protoc-gen-go-grpc
ARCH          			= $(shell uname -p)

BUF_VERSION   			:=	v1.28.1
SWAGGER_UI_VERSION	:=	v4.15.5

# Go Options
PKG         := ./...
TAGS        :=
TESTS       := .
TESTFLAGS   :=
LDFLAGS     := -w -s
GOFLAGS     :=
CGO_ENABLED ?= 0

# Rebuild the binary if any of these files change
SRC := $(shell find . -type f -name '*.go' -print) go.mod go.sum

.PHONY: all
all: build

.PHONY: build
build: $(BINDIR)/$(BINNAME)

$(BINDIR)/$(BINNAME): $(SRC)
	GO111MODULE=on CGO_ENABLED=$(CGO_ENABLED) go build $(GOFLAGS) -trimpath -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o '$(BINDIR)'/$(BINNAME) ./cmd/gateway

.PHONY: run
run: build ## Run the binary
	$(BINDIR)/$(BINNAME)

# #########
# # Proto #
# #########

.PHONY: generate
generate: proto/generate

.PHONY: proto/generate
proto/generate: ## Generate proto files
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) generate


.PHONY: proto/lint
proto/lint: ## Lint proto files
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) lint
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) breaking --against '.git#branch=main'

.PHONY: proto/clean
proto/clean: ## Clean proto files
	rm -rf proto/gen

# ------------------------------------------------------------------------------
# dependencies

.PHONY: download
download: ## Download dependencies
	go mod download

.PHONY: tidy
tidy: download ## Update go.mod and go.sum
	go mod tidy

.PHONY: vendor
vendor: tidy ## Update vendor directory
	go mod vendor

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