# Copyright Meshplay Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

include install/Makefile.core.mk
include install/Makefile.show-help.mk

#-----------------------------------------------------------------------------
# Docker-based Builds
#-----------------------------------------------------------------------------
.PHONY: docker-build docker-local-cloud docker-cloud docker-playground-build docker-testing-env-build docker-testing-env 

## Build Meshplay Server and UI container.
docker-build:
	# `make docker-build` builds Meshplay inside of a multi-stage Docker container.
	# This method does NOT require that you have Go, NPM, etc. installed locally.
	DOCKER_BUILDKIT=1 docker build -f install/docker/Dockerfile -t khulnasoft/meshplay --build-arg TOKEN=$(GLOBAL_TOKEN) --build-arg GIT_COMMITSHA=$(GIT_COMMITSHA) --build-arg GIT_VERSION=$(GIT_VERSION) --build-arg RELEASE_CHANNEL=${RELEASE_CHANNEL} .

## Build Meshplay Server and UI container in Playground mode.
docker-playground-build:
	# `make docker-playground-build` builds Meshplay inside of a multi-stage Docker container.
	# This method does NOT require that you have Go, NPM, etc. installed locally.
	DOCKER_BUILDKIT=1 docker build -f install/docker/Dockerfile -t khulnasoft/meshplay --build-arg TOKEN=$(GLOBAL_TOKEN) --build-arg GIT_COMMITSHA=$(GIT_COMMITSHA) --build-arg GIT_VERSION=$(GIT_VERSION) --build-arg RELEASE_CHANNEL=${RELEASE_CHANNEL} --build-arg PROVIDER=$(LOCAL_PROVIDER) --build-arg PROVIDER_BASE_URLS=$(MESHPLAY_CLOUD_PROD) --build-arg PLAYGROUND=true .

## Build Meshplay Server and UI container for e2e testing.
docker-testing-env-build:
	# `make docker-build` builds Meshplay inside of a multi-stage Docker container.
	# This method does NOT require that you have Go, NPM, etc. installed locally.
	DOCKER_BUILDKIT=1 docker build -f install/docker/testing/Dockerfile -t khulnasoft/meshplay-testing-env --build-arg GIT_VERSION=$(GIT_VERSION) .

## Meshplay Cloud for user authentication.
## Runs Meshplay in a container locally and points to locally-running
docker-local-cloud:
	(docker rm -f meshplay) || true
	docker run --name meshplay -d \
	--link meshplay-cloud:meshplay-cloud \
	-e PROVIDER_BASE_URLS=$(REMOTE_PROVIDER_LOCAL) \
	-e DEBUG=true \
	-e ADAPTER_URLS=$(ADAPTER_URLS) \
	-e KEYS_PATH=$(KEYS_PATH) \
	-p 9081:8080 \
	khulnasoft/meshplay ./meshplay

## Runs Meshplay in a container locally and points to remote
## Remote Provider for user authentication.
docker-cloud:
	(docker rm -f meshplay) || true
	docker run --name meshplay -d \
	-e PROVIDER_BASE_URLS=$(MESHPLAY_CLOUD_PROD) \
	-e DEBUG=true \
	-e ADAPTER_URLS=$(ADAPTER_URLS) \
	-e KEYS_PATH=$(KEYS_PATH) \
	-v meshplay-config:/home/appuser/.meshplay/config \
  -v $(HOME)/.kube:/home/appuser/.kube:ro \
	-p 9081:8080 \
	khulnasoft/meshplay ./meshplay

## Runs Meshplay in a container locally and points to remote
## Remote Provider for user authentication.
docker-testing-env:
	docker run --rm --name meshplaytesting  -d \
	-e PROVIDER_BASE_URLS=$(MESHPLAY_CLOUD_PROD) \
	-e DEBUG=true \
	-e ADAPTER_URLS=$(ADAPTER_URLS) \
	-e KEYS_PATH=$(KEYS_PATH) \
	-v meshplay-config:/home/appuser/.meshplay/config \
  -v $(HOME)/.kube:/home/appuser/.kube:ro \
	-p 9081:8080 \
	khulnasoft/meshplay-testing-env ./meshplay

#-----------------------------------------------------------------------------
# Meshplay Server Native Builds
#-----------------------------------------------------------------------------
.PHONY: server wrk2-setup nighthawk-setup server-local server-skip-compgen server-no-content golangci proto-build error build-server server-binary
## Setup wrk2 for local development.
wrk2-setup:
	echo "setup-wrk does not work on Mac Catalina at the moment"
	cd server; cd cmd; git clone https://github.com/khulnasoft/wrk2.git; cd wrk2; make; cd ..

## Setup nighthawk for local development.
nighthawk-setup: dep-check
	cd server; cd cmd; git clone https://github.com/khulnasoft/nighthawk-go.git; cd nighthawk-go; make setup; cd ..

run-local: server-local error

## Build and run Meshplay Server on your local machine
## and point to (expect) a locally running Remote Provider
## for user authentication.
server-local: dep-check
	cd server; cd cmd; go clean; go mod tidy; \
	BUILD="$(GIT_VERSION)" \
	PROVIDER_BASE_URLS=$(REMOTE_PROVIDER_LOCAL) \
	PORT=9081 \
	DEBUG=true \
	ADAPTER_URLS=$(ADAPTER_URLS) \
	APP_PATH=$(APPLICATIONCONFIGPATH) \
	KEYS_PATH=$(KEYS_PATH) \
	go run main.go error.go

## Build Meshplay Server on your local machine.
build-server: dep-check
	cd server; cd cmd; go mod tidy; cd "../.."
	BUILD="$(GIT_VERSION)" \
	PROVIDER_BASE_URLS=$(MESHPLAY_CLOUD_PROD) \
	PORT=9081 \
	DEBUG=true \
	ADAPTER_URLS=$(ADAPTER_URLS) \
	APP_PATH=$(APPLICATIONCONFIGPATH) \
	KEYS_PATH=$(KEYS_PATH) \
	GOPROXY=https://proxy.golang.org,direct GO111MODULE=on go build ./server/cmd/main.go ./server/cmd/error.go
	chmod +x ./main

## Running the meshplay server using binary.
server-binary:
	cd server/cmd; BUILD="$(GIT_VERSION)" PROVIDER_BASE_URLS=$(MESHPLAY_CLOUD_STAGING) ../../main; cd ../../

## Build and run Meshplay Server on your local machine
## and point to Remote Provider in staging environment
server-stg: dep-check
	cd server; cd cmd; go mod tidy; \
	BUILD="$(GIT_VERSION)" \
	PROVIDER_BASE_URLS=$(MESHPLAY_CLOUD_STAGING) \
	PORT=9081 \
	DEBUG=true \
	ADAPTER_URLS=$(ADAPTER_URLS) \
	APP_PATH=$(APPLICATIONCONFIGPATH) \
	KEYS_PATH=$(KEYS_PATH) \
	go run main.go error.go;

## Build and run Meshplay Server on your local machine.
server: dep-check
	cd server; cd cmd; go mod tidy; \
	BUILD="$(GIT_VERSION)" \
	PROVIDER_BASE_URLS=$(MESHPLAY_CLOUD_PROD) \
	PORT=9081 \
	DEBUG=true \
	ADAPTER_URLS=$(ADAPTER_URLS) \
	APP_PATH=$(APPLICATIONCONFIGPATH) \
	KEYS_PATH=$(KEYS_PATH) \
	go run main.go error.go;

## Build and run Meshplay Server on your local machine.
## Disable deployment of the Meshplay Operator to your Kubernetes cluster(s).
server-without-operator: dep-check
	cd server; cd cmd; go mod tidy; \
	BUILD="$(GIT_VERSION)" \
	PROVIDER_BASE_URLS=$(MESHPLAY_CLOUD_PROD) \
	PORT=9081 \
	DISABLE_OPERATOR=true \
	DEBUG=true \
	ADAPTER_URLS=$(ADAPTER_URLS) \
	APP_PATH=$(APPLICATIONCONFIGPATH) \
	KEYS_PATH=$(KEYS_PATH) \
	go run main.go error.go;

## Build and run Meshplay Server with no Kubernetes components on your local machine.
server-skip-compgen:
	cd server; cd cmd; go mod tidy; \
	BUILD="$(GIT_VERSION)" \
	PROVIDER_BASE_URLS=$(MESHPLAY_CLOUD_PROD) \
	PORT=9081 \
	DEBUG=true \
	ADAPTER_URLS=$(ADAPTER_URLS) \
	APP_PATH=$(APPLICATIONCONFIGPATH) \
 	SKIP_COMP_GEN=true \
	KEYS_PATH=$(KEYS_PATH) \
	go run main.go error.go;

## Build and run Meshplay Server on your local machine.
## Do not generate and register Kubernetes models.
server-without-k8s: dep-check
	cd server; cd cmd; go mod tidy; \
	BUILD="$(GIT_VERSION)" \
	REGISTER_STATIC_K8S=false \
	PROVIDER_BASE_URLS=$(MESHPLAY_CLOUD_PROD) \
	PORT=9081 \
	DEBUG=true \
	ADAPTER_URLS=$(ADAPTER_URLS) \
	APP_PATH=$(APPLICATIONCONFIGPATH) \
	KEYS_PATH=$(KEYS_PATH) \
	go run main.go error.go;

server-remote-provider: dep-check
	cd server; cd cmd; go mod tidy; \
	BUILD="$(GIT_VERSION)" \
	PROVIDER=$(REMOTE_PROVIDER) \
	PROVIDER_BASE_URLS=$(MESHPLAY_CLOUD_PROD) \
	PORT=9081 \
	DEBUG=true \
	ADAPTER_URLS=$(ADAPTER_URLS) \
	APP_PATH=$(APPLICATIONCONFIGPATH) \
	KEYS_PATH=$(KEYS_PATH) \
	go run main.go error.go;

server-local-provider: dep-check
	cd server; cd cmd; go mod tidy; \
	BUILD="$(GIT_VERSION)" \
	PROVIDER=$(LOCAL_PROVIDER) \
	PROVIDER_BASE_URLS=$(MESHPLAY_CLOUD_DEV) \
	PORT=9081 \
	DEBUG=true \
	ADAPTER_URLS=$(ADAPTER_URLS) \
	APP_PATH=$(APPLICATIONCONFIGPATH) \
	KEYS_PATH=$(KEYS_PATH) \
	go run main.go error.go;

## Build and run Meshplay Server with no seed content.
server-no-content:
	cd server; cd cmd; go mod tidy; \
	BUILD="$(GIT_VERSION)" \
	PROVIDER_BASE_URLS=$(MESHPLAY_CLOUD_PROD) \
	PORT=9081 \
	DEBUG=true \
	ADAPTER_URLS=$(ADAPTER_URLS) \
	APP_PATH=$(APPLICATIONCONFIGPATH) \
	SKIP_DOWNLOAD_CONTENT=true \
	KEYS_PATH=$(KEYS_PATH) \
	go run main.go error.go;

server-playground: dep-check
	cd server; cd cmd; go mod tidy; \
	BUILD="$(GIT_VERSION)" \
	PROVIDER=$(REMOTE_PROVIDER) \
	PROVIDER_BASE_URLS=$(MESHPLAY_CLOUD_PROD) \
	PORT=9081 \
	DEBUG=true \
	ADAPTER_URLS=$(ADAPTER_URLS) \
	APP_PATH=$(APPLICATIONCONFIGPATH) \
	PLAYGROUND=true \
	KEYS_PATH=$(KEYS_PATH) \
	go run main.go error.go;

## Lint check Meshplay Server.
golangci: error dep-check
	golangci-lint run

## Build Meshplay's protobufs.
proto-build:
	# see https://developers.google.com/protocol-buffers/docs/reference/go-generated
	# see https://grpc.io/docs/languages/go/quickstart/
	export PATH=$(PATH):$(GOBIN)
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	protoc --proto_path=server/meshes --go_out=server/meshes --go_opt=paths=source_relative --go-grpc_out=server/meshes --go-grpc_opt=paths=source_relative meshops.proto

## Analyze error codes
error: dep-check
	go run github.com/khulnasoft/meshkit/cmd/errorutil -d . analyze -i ./server/helpers -o ./server/helpers --skip-dirs meshplayctl

## Runs meshkit error utility to update error codes for meshplay server only.
error-util:
	go run github.com/khulnasoft/meshkit/cmd/errorutil -d . --skip-dirs meshplayctl update -i ./server/helpers/ -o ./server/helpers

## Build Meshplay UI; Build and run Meshplay Server on your local machine.
ui-server: ui-meshplay-build ui-provider-build server

#-----------------------------------------------------------------------------
# Meshplay UI Native Builds.
#-----------------------------------------------------------------------------
.PHONY: ui-setup ui ui-meshplay-build ui-provider ui-lint ui-provider ui-meshplay ui-build ui-provider-build ui-provider-test

UI_BUILD_SCRIPT = build16
UI_DEV_SCRIPT = dev16

ifeq ($(findstring v19, $(shell node --version)), v19)
	UI_BUILD_SCRIPT = build
	UI_DEV_SCRIPT = dev
else ifeq ($(findstring v18, $(shell node --version)), v18)
	UI_BUILD_SCRIPT = build
	UI_DEV_SCRIPT = dev
else ifeq ($(findstring v17, $(shell node --version)), v17)
	UI_BUILD_SCRIPT = build
	UI_DEV_SCRIPT = dev
endif

## Install dependencies for building Meshplay UI.
ui-setup:
	cd ui; npm i; cd ..
	cd provider-ui; npm i; cd ..

## Run Meshplay UI on your local machine. Listen for changes.
ui:
	cd ui; npm run dev; cd ..;

## Run Meshplay Provider UI  on your local machine. Listen for changes.
ui-provider:
	cd provider-ui; npm run dev; cd ..

## Lint check Meshplay UI and Provider UI on your local machine.
ui-lint:
	cd ui; npm run lint; cd ..

## Lint check Meshplay Provider UI on your local machine.
ui-provider-lint:
	cd provider-ui; npm run lint; cd ..

## Test Meshplay Provider UI on your local machine.
ui-provider-test:
	cd provider-ui; npm run test; cd ..

## Buils all Meshplay UIs  on your local machine.
ui-build: ui-setup
	cd ui; npm run lint:fix && npm run build && npm run export; cd ..
	cd provider-ui; npm run lint:fix && npm run build && npm run export; cd ..

## Build only Meshplay UI on your local machine.
ui-meshplay-build:
	cd ui; npm run build && npm run export; cd ..

## Builds only the provider user interface on your local machine
ui-provider-build:
	cd provider-ui; npm run build && npm run export; cd ..

## Run Meshplay Cypress Integration Tests against your local Meshplay UI (cypress runs in non-interactive mode).
ui-integration-tests: ui-setup
	cd ui; npm run ci-test-integration; cd ..

#-----------------------------------------------------------------------------
# Meshplay Docs
#-----------------------------------------------------------------------------
#Incorporating Make docs commands from the Docs Makefile
.PHONY: docs docs-build site docs-docker docs-meshplayctl
jekyll=bundle exec jekyll

site: docs
site-serve: docs-serve

## Run Meshplay Docs. Listen for changes.
docs:
	cd docs; bundle install; bundle exec jekyll serve --drafts --incremental --config _config_dev.yml

## Run Meshplay Docs. Do not listen for changes.
docs-serve:
	cd docs; bundle install; bundle exec jekyll serve --drafts --config _config_dev.yml

## Build Meshplay Docs on your local machine.
docs-build:
	cd docs; $(jekyll) build --drafts

## Run Meshplay Docs in a Docker container. Listen for changes.
docs-docker:
	cd docs; docker run --name meshplay-docs --rm -p 4000:4000 -v `pwd`:"/srv/jekyll" jekyll/jekyll:4.0 bash -c "bundle install; jekyll serve --drafts --livereload"

## Build Meshplay CLI docs
docs-meshplayctl:
	cd meshplayctl; make docs;
#-----------------------------------------------------------------------------
# Meshplay Helm Charts
#-----------------------------------------------------------------------------
.PHONY: helm-docs helm-operator-docs helm-meshplay-docs helm-operator-lint helm-lint
## Generate all Meshplay Helm Chart documentation in markdown format.
helm-docs: helm-operator-docs helm-meshplay-docs

## Generate Meshplay Operator Helm Chart documentation in markdown format.
helm-operator-docs: dep-check
	GO111MODULE=on go get github.com/norwoodj/helm-docs/cmd/helm-docs
	$(GOPATH)/bin/helm-docs -c install/kubernetes/helm/meshplay-operator

## Generate Meshplay Server and Adapters Helm Chart documentation in markdown format.
helm-meshplay-docs: dep-check
	GO111MODULE=on go get github.com/norwoodj/helm-docs/cmd/helm-docs
	$(GOPATH)/bin/helm-docs -c install/kubernetes/helm/meshplay

## Lint all of Meshplay's Helm Charts
helm-lint: helm-operator-lint helm-meshplay-lint

## Lint Meshplay Operator Helm Chart
helm-operator-lint:
	helm lint install/kubernetes/helm/meshplay-operator --with-subcharts
## Lint Meshplay Server and Adapter Helm Charts
helm-meshplay-lint:
	helm lint install/kubernetes/helm/meshplay --with-subcharts

#-----------------------------------------------------------------------------
# Meshplay APIs
#-----------------------------------------------------------------------------
.PHONY: swagger-build swagger swagger-docs-build graphql-docs-build graphql-build
## Build Meshplay REST API specifications
swagger-build:
	swagger generate spec -o ./server/helpers/swagger.yaml --scan-models

## Generate and serve Meshplay REST API specifications
swagger: swagger-build
	swagger serve ./server/helpers/swagger.yaml

## Build Meshplay REST API documentation
swagger-docs-build:
	swagger generate spec -o ./docs/_data/swagger.yml --scan-models; \
	swagger flatten ./docs/_data/swagger.yml -o ./docs/_data/swagger.yml --with-expand --format=yaml


## Building Meshplay docs with redocly
redocly-docs-build:
	npx @redocly/cli build-docs ./docs/_data/swagger.yml --config='redocly.yaml' -t custom.hbs

## Build Meshplay GraphQL API documentation
graphql-docs-build:
	cd docs; bundle exec rake graphql:compile_docs

## Build Meshplay GraphQl API specifications
graphql-build: dep-check
	cd server; cd internal/graphql; go run -mod=mod github.com/99designs/gqlgen generate



## testing
test-setup-ui:
	cd ui; npm ci; npx playwright install --with-deps; cd ..

test-ui:
	cd ui; npm run test:e2e; cd ..

#-----------------------------------------------------------------------------
# Dependencies
#-----------------------------------------------------------------------------
.PHONY: dep-check
#.SILENT: dep-check

INSTALLED_GO_VERSION=$(shell go version)

dep-check:

ifeq (,$(findstring $(GOVERSION), $(INSTALLED_GO_VERSION)))
# Only send a warning.
	@echo "Dependency missing: go$(GOVERSION). Ensure 'go$(GOVERSION).x' is installed and available in your 'PATH'"
	@echo "GOVERSION: " $(GOVERSION)
	@echo "INSTALLED_GO_VERSION: " $(INSTALLED_GO_VERSION)
# Force error and stop.
#	$(error Found $(INSTALLED_GO_VERSION). \
#	 Required golang version is: 'go$(GOVERSION).x'. \
#	 Ensure go '$(GOVERSION).x' is installed and available in your 'PATH'.)
endif