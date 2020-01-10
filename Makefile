GIT_VERSION=$(git describe --dirty)
CURRENT_TIME=$(date +%Y%m%d%H%M%S)

LD_VERSION_FLAGS=-X github.com/elotl/cloud-instance-provider/main.buildVersion=$(GIT_VERSION) -X github.com/elotl/cloud-instance-provider/main.buildTime=$(CURRENT_TIME)
LDFLAGS=-ldflags "$(LD_VERSION_FLAGS)"

BINARIES=virtual-kubelet

REGISTRY_REPO=elotl/virtual-kubelet

TOP_DIR=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))
PKG_SRC=$(shell find $(TOP_DIR)pkg -type f -name '*.go')
CMD_SRC=$(shell find $(TOP_DIR)cmd/virtual-kubelet -type f -name '*.go')
VENDOR_SRC=$(shell find $(TOP_DIR)vendor -type f -name '*.go')

all: $(BINARIES)

virtual-kubelet: $(PKG_SRC) $(VENDOR_SRC) $(CMD_SRC)
	CGO_ENABLED=0 go build $(LDFLAGS) -o $(TOP_DIR)$@ $(TOP_DIR)cmd/virtual-kubelet

img: $(BINARIES)
	@echo "Checking if BUILD_NUMBER is set" && test -n "$(BUILD_NUMBER)"
	img build -t $(REGISTRY_REPO):$(BUILD_NUMBER) \
		-t $(REGISTRY_REPO):dev .

login-img:
	@echo "Checking if REGISTRY_USER is set" && test -n "$(REGISTRY_USER)"
	@echo "Checking if REGISTRY_PASSWORD is set" && test -n "$(REGISTRY_PASSWORD)"
	@img login -u "$(REGISTRY_USER)" -p "$(REGISTRY_PASSWORD)" "$(REGISTRY_SERVER)"

push-img: img
	@echo "Checking if BUILD_NUMBER is set" && test -n "$(BUILD_NUMBER)"
	img push $(REGISTRY_REPO):$(BUILD_NUMBER)
	img push $(REGISTRY_REPO):dev

clean:
	rm -f $(BINARIES)

.PHONY: all clean
