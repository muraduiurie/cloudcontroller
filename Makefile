# Makefile
SHELL := /bin/bash
CONTROLLER_GEN := $(shell which controller-gen)
ENVTEST_K8S_VERSION = 1.31.0

# Ensure controller-gen is installed
ifndef CONTROLLER_GEN
$(error "controller-gen is not installed. Install it with: go install sigs.k8s.io/controller-tools/cmd/controller-gen@latest")
endif

# Directories
API_DIR := ./cmd/cloudcontroller/api/...
CRD_OUTPUT_DIR := ./cmd/cloudcontroller/crds
CRD_CHART_DIR := ./charts/cloudcontroller/templates/crds

.PHONY: generate
generate: ## Generate CRDs from Go structs
	@echo "Generating CRDs..."
	$(CONTROLLER_GEN) crd paths=$(API_DIR) output:crd:dir=$(CRD_OUTPUT_DIR)
	@echo "Copying CRDs to Helm Chart."
	rm $(CRD_CHART_DIR)/* && cp $(CRD_OUTPUT_DIR)/* $(CRD_CHART_DIR)
	@echo "CRD generation complete."

.PHONY: install-crds
install-crds: ## Apply generated CRDs to the cluster
	@echo "Applying CRDs to cluster..."
	kubectl apply -f $(CRD_OUTPUT_DIR)
	@echo "CRDs applied."

IMG ?= muraduiurie/cloudcontroller
TAG ?= "latest"
.PHONY: build-push
build-push: ## Build the operator
	@echo "Building docker image..."
	docker build -t $(IMG):$(TAG) . && docker push $(IMG):$(TAG)
	@echo "cloucontroller image built and pushed."

.PHONY: build
build: ## Build the operator
	@echo "Building docker image..."
	docker build -t $(IMG):$(TAG) .
	@echo "cloucontroller image built."

.PHONY: build
push: ## Build the operator
	@echo "Building docker image..."
	docker push $(IMG):$(TAG)
	@echo "cloucontroller image pushed."

.PHONY: clean
clean: ## Remove generated CRDs
	@echo "Cleaning up CRDs..."
	rm -rf $(CRD_OUTPUT_DIR)
	@echo "Cleanup complete."

.PHONY: help
help: ## Show available make targets
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: deploy
deploy: ## Deploy the operator
	@echo "Deploying operator..."
	helm upgrade --install cloudcontroller ./charts/cloudcontroller --set image.repository=$(IMG) --set image.tag=$(TAG)
	@echo "Operator deployed."

.PHONY: uninstall
uninstall: ## Deploy the operator
	@echo "Uninstalling operator..."
	helm uninstall cloudcontroller
	@echo "Operator uninstalled."

.PHYONY: mock
mock: ## Generate mocks
	@echo "Generating mocks..."
	cd cmd/cloudcontroller && mockgen -source=pkg/cloudproviders/gcp/client.go -destination=pkg/cloudproviders/gcp/mock.go -package=gcp && cd -
	@echo "Mocks generated."

.PHYONY: envtest
envtest: ## Setup envtest
	@echo "Setting up envtest..."
	setup-envtest cleanup && setup-envtest use $(ENVTEST_K8S_VERSION) --bin-dir envtest
	@echo "Setup complete."