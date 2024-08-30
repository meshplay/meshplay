# envs
GIT_VERSION=$(shell git describe --tags `git rev-list --tags --max-count=1`)
GIT_COMMITSHA=$(shell git rev-list -1 HEAD)
RELEASE_CHANNEL="edge"
LDFLAGS="-w -s -X github.com/khulnasoft/meshplay/meshplayctl/internal/cli/root/constants.version=${GIT_VERSION} -X github.com/khulnasoft/meshplay/meshplayctl/internal/cli/root/constants.commitsha=${GIT_COMMITSHA} -X github.com/khulnasoft/meshplay/meshplayctl/internal/cli/root/constants.releasechannel=${RELEASE_CHANNEL}"

make:
	gofmt -l -s -w .
	go build -ldflags=${LDFLAGS} -o meshplayctl cmd/meshplayctl/main.go
forwin:
	go build -ldflags=${LDFLAGS} -o meshplayctl.exe cmd/meshplayctl/main.go

OUTDIR := bin

ARCH := amd64

BINNAME_DARWIN ?= meshplayctl-darwin-$(ARCH)
BINNAME_LINUX ?= meshplayctl-linux-$(ARCH)
BINNAME_WINDOWS ?= meshplayctl-windows-$(ARCH).exe

# CLI

.PHONY: $(BINNAME_DARWIN)
$(BINNAME_DARWIN):
	@echo
	CGO_ENABLED=0 GOARCH=$(ARCH) GOOS=darwin GO111MODULE=on go build -ldflags=$(LDFLAGS) -gcflags='$(GCFLAGS)' -o $(OUTDIR)/$(BINNAME_DARWIN) ./cmd/meshplayctl/main.go

.PHONY: $(BINNAME_LINUX)
$(BINNAME_LINUX):
	@echo
	CGO_ENABLED=0 GOARCH=$(ARCH) GOOS=linux GO111MODULE=on go build -ldflags=$(LDFLAGS) -gcflags='$(GCFLAGS)' -o $(OUTDIR)/$(BINNAME_LINUX) ./cmd/meshplayctl/main.go

.PHONY: $(BINNAME_WINDOWS)
$(BINNAME_WINDOWS):
	@echo
	CGO_ENABLED=0 GOARCH=$(ARCH) GOOS=windows GO111MODULE=on go build -ldflags=$(LDFLAGS) -gcflags='$(GCFLAGS)' -o $(OUTDIR)/$(BINNAME_WINDOWS) ./cmd/meshplayctl/main.go

.PHONY: clean
clean:
	rm -rf $(OUTDIR)

.PHONY: error
error:
	go run github.com/khulnasoft/meshkit/cmd/errorutil -d . analyze -i ./helpers -o ./helpers

## Runs meshkit error utility to update error codes for meshplayctl.
error-util:
	go run github.com/khulnasoft/meshkit/cmd/errorutil -d . update -i ./helpers/ -o ./helpers

# Check code quality at local by using golangci-lint and return the coverage
.PHONY: lint
lint:
	golangci-lint run --timeout 5m

# output a HTML coverage by using gocov and gocov-html
.PHONY: coverage
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Generate markdown for meshplayctl docs
.PHONY: docs
docs:
	cd doc; go run doc.go
