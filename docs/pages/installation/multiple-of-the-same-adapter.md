---
layout: default
title: Using Multiple Adapters
permalink: installation/multiple-adapters
redirect_from: guides/multiple-adapters
type: guides
category: installation
language: en
abstract: Meshplay is capable of running zero or more adapters. Meshplay offers many features without the need for adapters. Adapters are optional components that enhance and extend Meshplay's core functionality.
---

## Advanced Configuration

Meshplay is capable of running zero or more adapters. Meshplay offers many features without the need for adapters. Adapters are optional components that enhance and extend Meshplay's core functionality.

### Modifying the default adapter deployment configuration

The number of adapters, type of adapters, where they are deployed, how they are named and what port they are exposed on are all configurable deployment options. To modify the default configuration, find `~/.meshplay/meshplay.yaml` on your system. `~/.meshplay/meshplay.yaml` is a Docker Compose file.

#### Configuration: Running fewer Meshplay adapters

To customize which Meshplay Adapters are used in which deployments, customize your contexts in your meshconfig.

*Recommended:*
Configure your meshconfig, using `meshplayctl system context` to customize which Meshplay Adapters are used in which deployments.

*Alternative:*
Alternatively, directly modify the `~/.meshplay/meshplay.yaml` configuration file, remove the entry(ies) of the adapter(s) you are removing from your deployment.

#### Configuration: Running more than one instance of the same Meshplay adapter

The default configuration of a Meshplay deployment includes one instance of each of the Meshplay adapters (that have reached a stable version status). You may choose to run multiple instances of the same type of Meshplay adapter; e.g. two instances of the Meshplay Adapter for NGINX Service Mesh. To do so, you can use either of Meshplay's clients or to modify your Meshplay deployment:
 - Using `meshplayctl`, modify `~/.meshplay/meshplay.yaml` to include multiple copies of the given adapter.
 - Using Meshplay UI, navigate to the Settings page and enter the host and port of your additional adapter.

#### Configuration: Choosing an adapter while installing Meshplayctl

While installing meshplayctl using bash installation script, we can choose which adapter to be loaded.
This is done by passing ADAPTERS environment variable to meshplay bash script.

*For e.g.* 
`curl -L https://meshplay.io/install | ADAPTERS=consul PLATFORM=kubernetes bash -` installs meshplayctl and starts Meshplay Server in your connected Kubernetes cluster deploying only the Meshplay Adapter for Consul and not the rest of Meshplay's adapters.

<h5>Demo of Meshplay managing deployments across multiple Kubernetes clusters:</h5>

<iframe class="container" width="560" height="315" src="https://www.youtube.com/embed/yWPu3vq4vEs?start=5041" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

See on YouTube: [Cloud Native Austin Virtual Meetup: April 2020](https://youtu.be/yWPu3vq4vEs?t=5041&list=PL3A-A6hPO2IOpTbdH89qR-4AE0ON13Zie)
