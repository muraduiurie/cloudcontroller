# Makefile

CONTROLLER_GEN := $(shell which controller-gen)

# Ensure controller-gen is installed
ifndef CONTROLLER_GEN
$(error "controller-gen is not installed. Install it with: go install sigs.k8s.io/controller-tools/cmd/controller-gen@latest")
endif

# Directories
API_DIR := ./cmd/cloudcontroller/api/...
CRD_OUTPUT_DIR := ./cmd/cloudcontroller/crds

.PHONY: generate
generate: ## Generate CRDs from Go structs
	@echo "Generating CRDs..."
	$(CONTROLLER_GEN) crd paths=$(API_DIR) output:crd:dir=$(CRD_OUTPUT_DIR)
	@echo "CRD generation complete."

.PHONY: install-crds
install-crds: generate ## Apply generated CRDs to the cluster
	@echo "Applying CRDs to cluster..."
	kubectl apply -f $(CRD_OUTPUT_DIR)
	@echo "CRDs applied."

.PHONY: build operator image
build: generate ## Build the operator
	@echo "Building docker image..."
	docker build -t $(IMG) .
	@echo "CRDs applied."

.PHONY: clean
clean: ## Remove generated CRDs
	@echo "Cleaning up CRDs..."
	rm -rf $(CRD_OUTPUT_DIR)
	@echo "Cleanup complete."

.PHONY: help
help: ## Show available make targets
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'