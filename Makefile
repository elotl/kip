GIT_VERSION=$(shell git describe --dirty)
CURRENT_TIME=$(shell date +%Y%m%d%H%M%S)
IMAGE_TAG=$(GIT_VERSION)
ifneq ($(findstring -,$(GIT_VERSION)),)
IMAGE_DEV_OR_LATEST=dev
else
IMAGE_DEV_OR_LATEST=latest
endif

LD_VERSION_FLAGS=-X main.buildVersion=$(GIT_VERSION) -X main.buildTime=$(CURRENT_TIME)
LDFLAGS=-ldflags "$(LD_VERSION_FLAGS)"

BINARIES=virtual-kubelet milpactl

REGISTRY_REPO=elotl/virtual-kubelet

TOP_DIR=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))
PKG_SRC=$(shell find $(TOP_DIR)pkg -type f -name '*.go')
CMD_SRC=$(shell find $(TOP_DIR)cmd/virtual-kubelet -type f -name '*.go')
VENDOR_SRC=$(shell find $(TOP_DIR)vendor -type f -name '*.go')
MILPACTL_SRC=$(shell find $(TOP_DIR)cmd/milpactl -type f -name '*.go')
GENERATED_SRC=$(TOP_DIR)pkg/clientapi/clientapi.pb.go \
			  $(TOP_DIR)pkg/api/deepcopy_generated.go

all: $(BINARIES)

virtual-kubelet: $(PKG_SRC) $(VENDOR_SRC) $(CMD_SRC) $(GENERATED_SRC)
	CGO_ENABLED=0 go build $(LDFLAGS) -o $(TOP_DIR)$@ $(TOP_DIR)cmd/virtual-kubelet


$(TOP_DIR)pkg/clientapi/clientapi.pb.go: $(TOP_DIR)pkg/clientapi/clientapi.proto
	protoc -I=$(TOP_DIR) --go_out=plugins=grpc:. $(TOP_DIR)pkg/clientapi/clientapi.proto

$(TOP_DIR)pkg/api/deepcopy_generated.go: $(TOP_DIR)pkg/api/types.go
	deepcopy-gen --input-dirs=github.com/elotl/cloud-instance-provider/pkg/api

# milpactl is compiled staticly so it'll run on pods
milpactl: $(PKG_SRC) $(VENDOR_SRC) $(MILPACTL_SRC)
	cd cmd/milpactl && CGO_ENABLED=0 GOOS=linux go build $(LDFLAGS) -o $(TOP_DIR)milpactl

img: $(BINARIES)
	@echo "Checking if IMAGE_TAG is set" && test -n "$(IMAGE_TAG)"
	img build -t $(REGISTRY_REPO):$(IMAGE_TAG) \
		-t $(REGISTRY_REPO):$(IMAGE_DEV_OR_LATEST) .

login-img:
	@echo "Checking if REGISTRY_USER is set" && test -n "$(REGISTRY_USER)"
	@echo "Checking if REGISTRY_PASSWORD is set" && test -n "$(REGISTRY_PASSWORD)"
	@img login -u "$(REGISTRY_USER)" -p "$(REGISTRY_PASSWORD)" "$(REGISTRY_SERVER)"

push-img: img
	@echo "Checking if IMAGE_TAG is set" && test -n "$(IMAGE_TAG)"
	img push $(REGISTRY_REPO):$(IMAGE_TAG)
	img push $(REGISTRY_REPO):$(IMAGE_DEV_OR_LATEST)

clean:
	rm -f $(BINARIES)

.PHONY: all clean
