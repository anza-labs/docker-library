REGISTRY_PORT ?= 5005
REPOSITORY    ?= library
TAG           ?= dev-$(shell git describe --match='' --always --abbrev=6 --dirty)
PLATFORM      ?= linux/$(shell arch)

# Variables defining the container tool to be used for building images.
# This Makefile supports common container tools like Docker and Podman.
CONTAINER_TOOL ?= docker
BUILD_COMMAND ?= buildx build --load
MANIFEST_ARGS ?=
BUILD_ARGS ?=

# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: all
all: registry

.PHONY: clean
clean: registry-down

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

# renovate: datasource=docker depName=docker.io/distribution/distribution
REGISTRY_VERSION ?= 3.0.0

# renovate: datasource=docker depName=docker.io/anchore/grype
GRYPE_VERSION ?= v0.96.0

.PHONY: registry
registry: ## Start the local registry.
	$(CONTAINER_TOOL) volume create registry_data
	$(CONTAINER_TOOL) run -d \
		--name=registry \
		--publish=$(REGISTRY_PORT):5000 \
		--restart=always \
		--volume=registry_data:/var/lib/registry \
		docker.io/distribution/distribution:$(REGISTRY_VERSION)

.PHONY: registry-down
registry-down: ## Stop and remove the local registry.
	-$(CONTAINER_TOOL) stop registry
	-$(CONTAINER_TOOL) rm registry
	-$(CONTAINER_TOOL) volume rm registry_data

c := ,
PLATFORMS_LIST := $(subst $c, ,$(PLATFORM))
REPOSITORY_LIST := $(subst $c, ,$(REPOSITORY))
TAG_LIST := $(subst $c, ,$(TAG))
BUILD_ARG_LIST := $(addprefix --build-arg=,$(subst $c, ,$(BUILD_ARGS)))

build-%: ## Build images per platform.
	for plat in $(PLATFORMS_LIST); do \
		$(CONTAINER_TOOL) $(BUILD_COMMAND) $(BUILD_ARG_LIST) \
			--platform=$$plat \
			--file=./library/$*/Dockerfile \
			$(foreach repo,$(REPOSITORY_LIST),$(foreach tag,$(TAG_LIST),--tag=$(repo)/$*:$(tag)-$${plat//\//_})) \
			./library/$* ; \
	done

push-%: build-% ## Push images per platform.
	for plat in $(PLATFORMS_LIST); do \
	for image in $(foreach repo,$(REPOSITORY_LIST),$(foreach tag,$(TAG_LIST),$(repo)/$*:$(tag)-$${plat//\//_})); do \
		$(CONTAINER_TOOL) push \
			$${image} ; \
	done; \
	done

scan-%: ## Scan image using grype.
	for repo in $(REPOSITORY_LIST); do \
	for tag in $(TAG_LIST); do \
	for plat in $(PLATFORMS_LIST); do \
		$(CONTAINER_TOOL) run --rm \
			docker.io/anchore/grype:${GRYPE_VERSION} \
			--platform=$${platform} \
			$${repo}/$*:$${tag}; \
	done; \
	done; \
	done

manifest-%: ## Create and push a manifest list that includes all built images.
	for repo in $(REPOSITORY_LIST); do \
	for tag in $(TAG_LIST); do \
		$(MAKE) _manifest-create-$* _REPOSITORY=$${repo} _TAG=$${tag}; \
		$(MAKE) _manifest-annotate-$* _REPOSITORY=$${repo} _TAG=$${tag}; \
		$(MAKE) _manifest-push-$* _REPOSITORY=$${repo} _TAG=$${tag}; \
	done; \
	done

_REPOSITORY ?=
_TAG ?=

_manifest-create-%:
	$(CONTAINER_TOOL) manifest create \
		--amend $(MANIFEST_ARGS) \
		$(_REPOSITORY)/$*:$(_TAG) \
		$(foreach plat,$(PLATFORMS_LIST),$(_REPOSITORY)/$*:$(_TAG)-$(subst /,_,$(plat)))

_manifest-annotate-%:
	for plat in $(PLATFORMS_LIST); do \
		$(CONTAINER_TOOL) manifest annotate \
			--arch=$$(echo $$plat | cut -d/ -f2) \
			$(_REPOSITORY)/$*:$(_TAG) \
			$(_REPOSITORY)/$*:$(_TAG)-$${plat//\//_} ; \
	done

_manifest-push-%:
	$(CONTAINER_TOOL) manifest push \
		$(MANIFEST_ARGS) \
		$(_REPOSITORY)/$*:$(_TAG)

