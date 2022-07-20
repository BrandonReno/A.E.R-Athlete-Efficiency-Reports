GO=go
BUILD_DIR=bin
BUILD_FLAGS=-ldflags "-w -extldflags '"-static"' -X ${PACKAGE_BUILD_INFO}/util.solutionType=${solution}"
OUTPUT_BIN=${service}
GODIRS_NOVENDOR = $(shell go list ./... | grep -v /vendor/)
GOFILES_NOVENDOR = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

pre-build:
    ifdef output
    	$(info "Output set to ${output}")
    	OUTPUT_BIN=${output}
    endif
    BUILD_FLAGS=-ldflags "-w -extldflags '"-static"' -X ${PACKAGE_BUILD_INFO}/util.solutionType=${solution}"
    ifeq ($(DEBUG),true)
        BUILD_FLAGS=-gcflags="all=-N -l" -ldflags "-extldflags '"-static"' -X ${PACKAGE_BUILD_INFO}/util.solutionType=${solution}"
        $(info "Building ${service} in debug mode")
    endif

build: pre-build
	CGO_ENABLED=0 GODEBUG=tls13=0 GOOS=linux $(GO) build ${BUILD_FLAGS} -o ${BUILD_DIR}/${service}/${OUTPUT_BIN}.bin ./${service}

fmt:
	gofumpt -extra -l -w ${GOFILES_NOVENDOR}

lint:
	golangci-lint run --timeout=10m -v ./...

