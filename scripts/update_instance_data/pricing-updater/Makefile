NAME=$(shell basename $(TOP_DIR))

GIT_COMMIT=$(shell git rev-parse --short HEAD)
CURRENT_TIME=$(shell date +%Y%m%d%H%M%S)

TOP_DIR=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))
PKG_SRC=$(shell find $(TOP_DIR)pkg -type f -name "*.go")
CMD_SRC=$(shell find $(TOP_DIR)cmd/main.go -type f -name "*.go")
MODULE_FILES=go.mod go.sum

LD_VERSION_FLAGS=-X main.BuildVersion=$(GIT_COMMIT) -X main.BuildTime=$(CURRENT_TIME)
LDFLAGS=-ldflags "$(LD_VERSION_FLAGS)"

BINARIES=$(NAME)

all: $(BINARIES) $(MODULE_FILES)

$(NAME): $(PKG_SRC) $(CMD_SRC)
	cd cmd/ && go build $(LDFLAGS) -o $(TOP_DIR)$(NAME)

test: $(NAME)
	go test ./...

DKR=docker
AWS_PROFILE ?= "elotl"
AWS_REGION ?= "us-east-1"
REGISTRY_PROTOCOL=https://
REGISTRY_DOMAIN=689494258501.dkr.ecr.us-east-1.amazonaws.com
REGISTRY_REPO_NAME=/elotl/$(NAME)
REGISTRY_SERVER ?= $(REGISTRY_PROTOCOL)$(REGISTRY_DOMAIN)
REGISTRY_REPO ?= $(REGISTRY_DOMAIN)$(REGISTRY_REPO_NAME)
IMAGE_TAG=$(GIT_COMMIT)
ifneq ($(findstring -,$(GIT_COMMIT)),)
IMAGE_DEV_OR_LATEST=dev
else
IMAGE_DEV_OR_LATEST=latest
endif

img: $(BINARIES)
	@echo "Checking if IMAGE_TAG is set" && test -n "$(IMAGE_TAG)"
	$(DKR) build -t $(REGISTRY_REPO):$(IMAGE_TAG) \
		-t $(REGISTRY_REPO):$(IMAGE_DEV_OR_LATEST) .

login-img:
	@echo "Checking if REGISTRY_USER is set" && test -n "$(REGISTRY_USER)"
	@echo "Checking if REGISTRY_PASSWORD is set" && test -n "$(REGISTRY_PASSWORD)"
	@$(DKR) login -u "$(REGISTRY_USER)" -p "$(REGISTRY_PASSWORD)" "$(REGISTRY_SERVER)"

push-img: img
	@echo "Checking if IMAGE_TAG is set" && test -n "$(IMAGE_TAG)"
	$(DKR) push $(REGISTRY_REPO):$(IMAGE_TAG)
	#$(DKR) push $(REGISTRY_REPO):$(IMAGE_DEV_OR_LATEST)

clean:
	rm -f $(BINARIES)

.PHONY: all clean
