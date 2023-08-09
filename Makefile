API=cmd/api/main.go

##@ General

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Dependencies

install-all: ## Install all Development dependencies
install-all: install-ginkgo install-swag

GINKGO = $(shell pwd)/bin/ginkgo
install-ginkgo:
	$(call go-install-tool,$(GINKGO),github.com/onsi/ginkgo/v2/ginkgo@v2.11.0)

SWAG = $(shell pwd)/bin/swag
install-swag:
	# version 2 it's almost in there, but keep in 1.8 now.
	$(call go-install-tool,$(SWAG),github.com/swaggo/swag/cmd/swag@latest)

##@ Development

build: ## Build the api project
	go build $(API)

run: ## Run the api project
	go run $(API)

openapi: ## Create the latest openapi specs from c
openapi: install-swag
	$(SWAG) init -g $(API)

test: ## Run tests.
test: install-ginkgo
	$(GINKGO) --race -r ./...

##@ Build

# go-install-tool will 'go install' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-install-tool
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/bin go install $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef
