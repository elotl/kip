# Override the GO variable to test with a different version of Golang:
#     $ make GO=/usr/lib/go-1.15/bin/go
GO=go
DKR=docker
GIT_VERSION=$(shell git describe --dirty)
CURRENT_TIME=$(shell date +%Y%m%d%H%M%S)
IMAGE_TAG=$(GIT_VERSION)

LD_VERSION_FLAGS=-X main.buildVersion=$(GIT_VERSION) -X main.buildTime=$(CURRENT_TIME) -X github.com/elotl/kip/pkg/util.VERSION=$(GIT_VERSION)
LDFLAGS=-ldflags "$(LD_VERSION_FLAGS)"

BINARIES=kip kipctl

REGISTRY_REPO=elotl/kip

TOP_DIR=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))
PKG_SRC=$(shell find $(TOP_DIR)pkg -type f -name '*.go')
CMD_SRC=$(shell find $(TOP_DIR)cmd/kip -type f -name '*.go')
MODULE_FILES=go.mod go.sum
KIPCTL_SRC=$(shell find $(TOP_DIR)cmd/kipctl -type f -name '*.go')
GENERATED_SRC=$(TOP_DIR)pkg/clientapi/clientapi.pb.go \
			  $(TOP_DIR)pkg/api/deepcopy_generated.go

all: $(BINARIES)

kip: $(PKG_SRC) $(CMD_SRC) $(GENERATED_SRC) $(MODULE_FILES)
	CGO_ENABLED=0 $(GO) build $(LDFLAGS) -o $(TOP_DIR)$@ $(TOP_DIR)cmd/kip


$(TOP_DIR)pkg/clientapi/clientapi.pb.go: $(TOP_DIR)pkg/clientapi/clientapi.proto
	protoc -I=$(TOP_DIR) --go_out=plugins=grpc:. $(TOP_DIR)pkg/clientapi/clientapi.proto

$(TOP_DIR)pkg/api/deepcopy_generated.go: $(TOP_DIR)pkg/api/types.go
	deepcopy-gen \
		--input-dirs=github.com/elotl/kip/pkg/api \
		--go-header-file $(TOP_DIR)scripts/boilerplate.go.txt

# kipctl is compiled staticly so it'll run on pods
kipctl: $(PKG_SRC) $(KIPCTL_SRC) $(MODULE_FILES)
	cd cmd/kipctl && CGO_ENABLED=0 GOOS=linux $(GO) build $(LDFLAGS) -o $(TOP_DIR)kipctl

img: $(BINARIES)
	@echo "Checking if IMAGE_TAG is set" && test -n "$(IMAGE_TAG)"
	$(DKR) build -t $(REGISTRY_REPO):$(IMAGE_TAG) .

login-img:
	@echo "Checking if REGISTRY_USER is set" && test -n "$(REGISTRY_USER)"
	@echo "Checking if REGISTRY_PASSWORD is set" && test -n "$(REGISTRY_PASSWORD)"
	@$(DKR) login -u "$(REGISTRY_USER)" -p "$(REGISTRY_PASSWORD)" "$(REGISTRY_SERVER)"

push-img: img
	@echo "Checking if IMAGE_TAG is set" && test -n "$(IMAGE_TAG)"
	$(DKR) push $(REGISTRY_REPO):$(IMAGE_TAG)

clean:
	rm -f $(BINARIES)

.PHONY: all clean

pricing-updater:
	cd scripts/update_instance_data/pricing-updater/cmd  && $(GO) build $(LDFLAGS) -o $(TOP_DIR)scripts/update_instance_data/pricing-updater/pricing-updater

update-pricing-data: pricing-updater
	echo "scraping data for all providers"
	cd scripts/update_instance_data/pricing-updater/ && ./pricing-updater -scrape-all
	cd scripts/update_instance_data/pricing-updater/scripts && pwd && TOP_KIP_DIR=$(TOP_DIR) $(GO) generate
