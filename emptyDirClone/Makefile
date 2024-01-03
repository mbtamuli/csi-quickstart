# Just typing `make` shows all targets and their description.
.PHONY: all
all: help

SHELL=/bin/bash

# Binary PATH
BINARY := bin/emptydirclone
# Image URL to use all building/pushing image targets
IMG ?= ghcr.io/mbtamuli/csi-quickstart/emptydirclone
# Image tag
TAG ?= latest
# Version of the application
VERSION ?= $(GIT_INFO)
# Namespace to deploy the CSI plugin
CSI_NAMESPACE ?= emptydirclone
# E2E Test args
ifdef E2E_TEST_ARGS
E2E_ARGS = --args $(E2E_TEST_ARGS)
endif

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Cluster

.PHONY: cluster-up
cluster-up: ## Create a local Kubernetes cluster using `kind`.
	kind create cluster

.PHONY: cluster-down
cluster-down: ## Destroy the local cluster.
	kind delete cluster

##@ Development

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: csc-test
csc-test: fmt vet build ## Run `csc` tests.
	@./tests/csc_tests.sh $(BINARY)

.PHONY: e2e
e2e: fmt vet ## Run e2e tests.
	env KUBECONFIG=${PWD}/tests/e2e/kubeconfig.yaml go test ./tests/e2e -test.v $(E2E_ARGS)


STERN_EXISTS := $(shell stern --version 2>/dev/null)
.PHONY: logs
logs:
ifdef STERN_EXISTS
	stern --namespace emptydirclone \
		--selector 'app=emptydirclone-plugin' \
		--container '(csi-volume-driver|node-driver-registrar)'
else
	kubectl logs --namespace emptydirclone \
		--selector 'app=emptydirclone-plugin' \
		--container csi-volume-driver \
		--container node-driver-registrar \
		--follow --ignore-errors --prefix
endif

##@ Build

.PHONY: build
build: ## Build binary for testing with `csc`
	go build -ldflags "-s -w -X main.version=${VERSION}" -o $(BINARY) main.go

.PHONY: docker-build
docker-build: ## Build docker image with the plugin.
	docker build --build-arg VERSION=${VERSION} --tag ${IMG}:${TAG} .
	docker build --build-arg VERSION=${VERSION} --target debug --tag ${IMG}:debug .

.PHONY: docker-push
docker-push: ## Push docker image with the plugin.
	docker push ${IMG}:${TAG}
	docker push ${IMG}:debug

CLUSTER_CONTEXT=$(shell kubectl config current-context)
KIND_CLUSTER_NAME=$(CLUSTER_CONTEXT:kind-%=%)
.PHONY: kind-load
kind-load: ## Load docker image with the plugin, into Kind cluster targeted by the current context(`kubectl config current-context`).
	kind load docker-image ${IMG}:${TAG} --name $(KIND_CLUSTER_NAME)
	kind load docker-image ${IMG}:debug --name $(KIND_CLUSTER_NAME)

##@ Deploy

.PHONY: deploy
deploy: ## Deploy the plugin to the K8s cluster specified in $KUBECONFIG.
	@kubectl create namespace $(CSI_NAMESPACE) \
		&& kubectl config set-context $(shell kubectl config current-context) --namespace=$(CSI_NAMESPACE) \
		&& kubectl apply --filename deploy

.PHONY: undeploy
undeploy: ## Undeploy the plugin from the K8s cluster specified in $KUBECONFIG.
	@kubectl delete --filename deploy --ignore-not-found=true --force=true \
		&& kubectl delete namespace $(CSI_NAMESPACE) --ignore-not-found=true
## Helpers

# Get information from git to reliably identify
# the commit from a branch/tag to set the version
GIT_HASH=$(shell git describe --always --long)
GIT_TAG=$(shell git name-rev --name-only --tags HEAD | sed 's/^undefined$$//')
GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GIT_INFO=$(if $(GIT_TAG),$(GIT_TAG),$(GIT_BRANCH)-$(GIT_HASH))
