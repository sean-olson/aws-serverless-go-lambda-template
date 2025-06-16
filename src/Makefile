GOOS ?= linux
GOARCH ?= amd64
BUILD_MODE ?= release
BINARY_NAME ?= bootstrap
ZIP_NAME ?= $(BINARY_NAME).zip

# Conditional build flags
ifeq ($(BUILD_MODE),debug)
    GOFLAGS=-gcflags="all=-N -l"
else
    GOFLAGS=
endif

local-build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(GOFLAGS) -o bin/$(BINARY_NAME) handler.go
	cd bin && zip $(ZIP_NAME) $(BINARY_NAME)

docker-build:
	docker run --rm -v "$(PWD)":/go/src/app -w /go/src/app \
	    amazonlinux:2 \
	    /bin/bash -c "\
	        yum install -y gcc golang zip && \
	        GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(GOFLAGS) -o bin/$(BINARY_NAME) handler.go && \
	        cd bin && zip $(ZIP_NAME) $(BINARY_NAME)"

clean:
	rm -f bin/$(BINARY_NAME) bin/$(ZIP_NAME)

deploy: docker-build
	serverless deploy
