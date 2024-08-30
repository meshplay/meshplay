---
layout: page
title: Contributing to Meshplay Adapters
permalink: project/contributing/contributing-adapters
abstract: How to contribute to Meshplay Adapters
language: en
type: project
category: contributing
list: include
abstract: How to contribute to Meshplay Adapters
---


Meshplay Adapters are the Extension Points in Meshplay's architecture. Their design, the process of creating a new adapter is documented in [Extensibility: Meshplay Adapters]({{site.baseurl}}/extensibility/adapters).

With the [CONTRIBUTING.md](https://github.com/meshplay/meshplay/blob/master/CONTRIBUTING.md#adapter) in mind, understand that development follows the usual fork-and-pull request workflow described here, see also GitHub Process. On forking GitHub deactivates all workflows. It is safe and good practice to activate them such that the code is validated on each push. This requires that branches filter for “on push” is set to ‘**’ to be triggered also on branches containing ‘/’  in their name.  The actions are parameterized using secrets (see Build & Release Strategy). The Docker image is only built and pushed to Docker Hub if a tag is pushed and the corresponding authentication information is configured. The only secret that should be set in each fork is GO_VERSION, specified in Build & Release Strategy, otherwise, the corresponding action’s default version is used.

Each commit has to be signed off, see [Contributing Overview]({{site.baseurl}}/project/contributing).

### Running an adapter as a container

Testing your local changes running as a container can be accomplished in two ways:

1. Define the adapter’s address in the UI: Unless the running container is named as specified in the docker-run target in the Makefile, the container has to be removed manually first. Then, run `make docker` followed by `make docker-run`. Then, connect to the adapter in the UI in “Settings > Adapters” using `localhost:<port>` if the meshplay server is running as a binary, or <docker IP address>:<port> if it is running as a docker container.
1. Using meshplayctl: In `~/.meshplay/meshplay.yaml`, change the tag specifying the image of the adapter to “latest”. Run make docker, followed by `meshplayctl system start --skip-update`. This assumes meshplayctl system start has been executed at least once before.

### Running an adapter as a process

Another way to test your local changes is to run the adapter as a process. To do this, clone the `meshplay/meshplay` repository, and start Meshplay Server using `make server`. Start the adapter from your IDE, or by executing `make run`. Then, in Meshplay UI, add the adapter using “localhost:<PORT>”.

### Creating a new Meshplay Adapter

Meshplay uses adapters to manage and interact with different cloud native infrastructure. Meshplay adapters are written in Go. Whether you are creating a new adapter or modifying an existing adapter, be sure to read the [Meshplay Adapters](https://docs.google.com/document/d/1b8JAMzr3Rntu7CudRaYv6r6ccACJONAB5t7ISCaPNuA/edit#) design specification. For new adapters, start with the Repository Template(https://github.com/meshplay/meshplay). 

1. Get the proto buf spec file from Meshplay repo:
   `wget https://raw.githubusercontent.com/meshplay/meshplay/master/meshes/meshops.proto`
1. Generate code
   1. Using Go as an example, do the following:
      - adding GOPATH to PATH: `export PATH=$PATH:$GOPATH/bin`
      - install grpc: `go get -u google.golang.org/grpc`
      - install protoc plugin for go: `go get -u github.com/golang/protobuf/protoc-gen-go`
      - Generate Go code: `protoc -I meshes/ meshes/meshops.proto --go_out=plugins=grpc:./meshes/`
   1. For other languages, please refer to gRPC.io for language-specific guides.
1. Implement the service methods and expose the gRPC server on a port of your choice (e.g. 10000).

Tip: The [Meshplay Adapter for Istio](https://github.com/meshplay/meshplay-istio) is a good reference adapter to use as an example of a Meshplay Adapter.
