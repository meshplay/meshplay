---
layout: default
title: Adapters
permalink: concepts/architecture/adapters
type: components
redirect_from: architecture/adapters
abstract: "Adapters extend Meshplay's management capabilities in any number of ways, including lifecycle, configuration, performance, governance, identity..."
language: en
list: include
---

## What are Meshplay Adapters?

Part of Meshplay's extensibility as a platform, Meshplay Adapters are purpopse-built to address an area in need of management that is either considered optional to the platform and/or is considered an area in which additional depth of control is needed. Adapters extend Meshplay's management capabilities in any number of ways, including lifecycle, configuration, performance, governance, identity and so on. Meshplay Adapters come in different form factors, and depending on their purpose, deliver different sets or capabilities. Each Adapter registers its capabilities with Meshplay Server. Meshplay Server, in-turn, exposes those capabilities for you to control.

## Meshplay Adapters for Lifecycle Management

Adapters that extend Meshplay's lifecycle management capabilities for infrastructure do so, by offering an infrastructure-specific interface to increase the depth of control that Meshplay has over a particular technology. Meshplay uses adapters to offer choice of load generator (for performance management) and for managing different layers of your infrastructure. Lifecycle adapters allow Meshplay to interface with the different cloud native infrastructure, exposing their differentiated value to users.

Meshplay has lifecycle adapters for managing the following cloud native infrastructure.
{% assign sorted = site.adapters | sort: "project_status" | reverse %}

| Adapter Status | Adapter | Port | Earliest Version supported |
| :------------: | :----------: | :--: | :------------------------: |
{% for adapter in sorted -%}
{% if adapter.project_status -%}
| {{ adapter.project_status }} | <img src="{{ adapter.image }}" style="width:20px" data-logo-for-dark="{{ adapter.white_image }}" data-logo-for-light="{{ adapter.image }}" id="logo-dark-light" loading="lazy"/> [{{ adapter.name }}]({{ site.baseurl }}{{ adapter.url }}) | {{ adapter.port }} | {{adapter.earliest_version}} |
{% endif -%}
{% endfor %}

## Meshplay Adapters for Performance Management

_v0.8.0 Roadmap:_ The `meshplay-nighthawk` adapter externalizes Nighthawk as an Meshplay component.

Meshplay Server allows users to generate traffic load tests using Nighthawk, fortio, and wrk2. Using the `meshplay-nigthhawk` adapter, you can schedule, control, and execute performance tests.

Run the `meshplay-nighthawk` adapter as an externalized load generator when you: 

1. Need a smaller sized container image for Meshplay. Nighthawk binaries are dynamically linked (C++) and they need other dependencies to work. This causes bloat in Meshplay Server’s image which doesn’t need them.
1. Need *adaptive load control* of your performance tests, controlling the variability by which the system under test receives load. Use Meshplay Server to run adaptive load tests.
1. Need *distributed load testing* and the ability to horizontally scale Nighthawk, using Nighthawk’s execution forwarding service and results sink.

## Adapter Deployment and Registration

Like every Meshplay component, Meshplay Adapters use MeshKit.

### Adapter FAQs

#### Is each Meshplay adapter made equal?

No, different Meshplay adapters are written to expose the unique value of each cloud native infrastructure. Consequently, they are not equally capable just as each cloud native infrastructure is not equally capable as the other. Each Adapter has a set of operations which are grouped based on predefined operation types. See the [extensibility]({{site.baseurl}}/extensibility) page for more details on adapter operations.

#### How can I create a new adapter?

Yes, see the [extensibility]({{site.baseurl}}/extensibility) documentation for details how to create a new Meshplay Adapter. See the Meshplay Adapter Template repository as boilerplate for your new adapter.

#### Do adapters have to be written in Golang?

No. Adapters much interface with Meshplay Server via gRPC. What language is used in that adapter is the perogative of a given adapter's maintainers.

#### Can I run more than one instance of the same Meshplay adapter?

Yes. The default configuration of a Meshplay deployment includes one instance of each of the Meshplay adapters (that have reached a stable version status). You may choose to run multiple instances of the same type of Meshplay adapter; e.g. two instances of the `meshplay-istio` adapter. To do so, modify ~/.meshplay/meshplay.yaml to include multiple copies of the given adapter.

See the "[Multiple Adapters]({{site.baseurl}}/guides/installation/multiple-adapters)" guide for more information.
