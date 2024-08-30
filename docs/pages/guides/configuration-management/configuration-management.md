---
layout: default
title: Configuring your Cloud Native Infrastructure and Applications
abstract: Learn how to use Meshplay Designs effectively
permalink: guides/configuration-management/working-with-designs
type: guides
category: configuration
language: en
suggested-reading: false
abstract: Learn how to use Meshplay Designs effectively
redirect_from:
  - /tasks/patterns.md
---

As an cloud native management platform, Meshplay has built-in support infastructure and application configuration management. Use Meshplay to configure your multi-cluster Kubernetes environments on-premises or across clouds.

## What is a Meshplay Design?

Meshplay's core construct of configuration management is that of a design. A design is ultimately a document – and a document in which you capture and describe your desired state of improv the environment.

[![Context Aware Design]({{ site.baseurl }}/assets/img/patterns/context-aware-design.svg
)]({{ site.baseurl }}/assets/img/patterns/context-aware-design.svg)

## Creating a Meshplay Design

You have your choice of how are you would like to create an infrastructure design. You can use the design configurator inside of measure, UI or extensions, like MeshMap, or you can hand code your design, using the mesh model spec, if you like.

You are free to choose how you would like to create an infrastructure design. You can use the design configurator inside of measure, UI or extensions, like MeshMap, or you can hand code your design, using the mesh model spec, if you like.

Use Meshplay UI or meshplayctl to manage [designs]({{ site.baseurl }}/tasks/patterns).

## Pattern Management Through Meshplay UI

Meshplay also comes with seed patterns when users start Meshplay for the first time. These patterns cover common use cases and examples for you as you explore Meshplay.

<img src="{{ site.baseurl }}/assets/img/configuration-management/meshplay-patterns.png" width="50%" />

Users can also import these patterns to their remote provider from this [sample repository](https://github.com/service-mesh-patterns/service-mesh-patterns/tree/master/samples).

<img src="{{ site.baseurl }}/assets/img/configuration-management/pattern-import.png" width="50%" />

Once these patterns are imported, you can then edit these patterns or use the pattern configurator to configure them according to your requirements.

<img src="{{ site.baseurl }}/assets/img/configuration-management/pattern-configure-button.png" width="50%" />

<img src="{{ site.baseurl }}/assets/img/configuration-management/pattern-configure.png" width="50%" />

## Pattern Management Through Meshplay CLI

You can also manage cloud native patterns through Meshplay's CLI, meshplayctl.

The `meshplayctl pattern` subcommand lets you import and apply patterns to your cluster.

For example, if you have your pattern written in a file say, `istio-bookinfo.yaml` which deploys Istio service mesh and onboards the BookInfo app on Istio, you can use meshplayctl to apply this pattern as shown below:

```
meshplayctl pattern apply -f istio-bookinfo.yaml
```

If you already have a pattern imported into Meshplay, you can apply the pattern by name.

```
meshplayctl pattern apply BookInfoApp
```

This will apply the pattern BookInfoApp, which has already been imported into Meshplay.

See [meshplayctl pattern subcommand section]({{ site.baseurl }}/reference/meshplayctl/#cloud-native-pattern-configuration-and-management) for more details on the `pattern` subcommand.

## WASM Filters

Meshplay can be used for managing WebAssembly Filters through the UI or the CLI.

### Filter Management Through Meshplay UI

Like patterns, Meshplay also comes with some sample WebAssembly Filters for you to experiment.

<img src="{{ site.baseurl }}/assets/img/configuration-management/meshplay-filters.png" width="50%" />

You can also import these filters manually to your provider from the [wasm-filters](https://github.com/layer5io/wasm-filters) repo.

Meshplay's sample application [ImageHub]({{ site.baseurl }}/guides/infrastructure-management/deploying-sample-apps) will let you test out configuring these filters out-of-the-box.

You can onboard ImageHub to an installed service mesh as shown below.

<img src="{{ site.baseurl }}/assets/img/configuration-management/image-hub.png" width="50%" />

### Filter Management Through Meshplay CLI

You can also manage WASM filters through Meshplay's CLI, meshplayctl.

The `meshplayctl filter` command lets you import and configure WebAssembly filters.

For example,

```
meshplayctl exp filter apply -f metrics_collector_bg.wasm
```

If you already have a filter imported into Meshplay, you can configure the filter by name.

```
meshplayctl exp filter apply metrics_collector_bg
```