---
layout: default
title: Architecture
permalink: concepts/architecture
redirect_from: architecture
type: components
abstract: overview of different individual components of Meshplay architecture and how they interact as a system.
language: en
list: include
---

## Components, their Purpose, and Languages

Meshplay and its components are written using the following languages and technologies.

| Components                                                           | Languages and Technologies                                                        |
| :------------------------------------------------------------------- | :-------------------------------------------------------------------------------- |
| Meshplay Server                                                       | Golang, gRPC, GraphQL, [SMP](https://smp-spec.io)                                 |
| [Meshplay Adapters](/concepts/architecture/adapters)                  | Golang, gRPC, [CloudEvents](https://cloudevents.io/), [SMI](https://smi-spec.io), [OAM](https://oam.dev)  |
| [Meshplay WASM Filters](https://github.com/khulnasoft/wasm-filters)     | Rust and C++                                                                      |
| Meshplay UI                                                           | ReactJS, NextJS, BillboardJS                                                      |
| Meshplay Provider UI                                                  | ReactJS, NextJS                                                                   |
| [Meshplay Remote Providers](/extensibility/providers)                 | _any_ - must adhere to Meshplay [Extension Points]({{site.baseurl}}/extensibility) |
| [Meshplay Operator](/concepts/architecture/operator)                  | Golang                                                                            |
| [MeshSync](/concepts/architecture/meshsync)                          | Golang                                                                            |
| [Broker](/concepts/architecture/broker)                              | Golang, NATS                                                                      |
| [Meshplay Database](/concepts/architecture/database)                  | Golang, SQLlite                                                                   | 
| [Meshplay CLI](#meshplay-cli) | Golang                                                                            |

## Deployments

Meshplay deploys as a set of containers. Meshplay's containers can be deployed to either Docker or Kubernetes. Meshplay components connect to one another via gRPC requests. Meshplay Server stores the location of the other components and connects with those components as needed. Typically, a connection from Meshplay Server to Meshplay Adapters is initiated from a client request (usually either `meshplayctl` or Meshplay UI) to gather information from the Adapter or invoke an Adapter's operation.

### Adapters

In Meshplay v0.6.0, Adapters will register with Meshplay Server over HTTP POST. If Meshplay Server is not available, Meshplay Adapters will backoff and retry to connect to Meshplay Server perpetually.

<a href="{{ site.baseurl }}/assets/img/architecture/meshplay-architecture.svg" class="lightbox-image">
<img src="{{ site.baseurl }}/assets/img/architecture/meshplay-architecture.svg" width="50%" /></a>

_Figure: Meshplay deploys inside or outside of a Kubernetes cluster_

#### Adapters and Capabilities Registry

Each Meshplay Adapter delivers its own unique specific functionality. As such, at time of deployment, the Meshplay Adapter will register its cloud native infrastructure-specific capabilities (its operations) with Meshplay Server's capability registry.

<a href="{{ site.baseurl }}/assets/img/adapters/meshplay-adapter-operation-registration.svg" class="lightbox-image">
<img src="{{ site.baseurl }}/assets/img/adapters/meshplay-adapter-operation-registration.svg" width="50%" /></a>

_Figure: Meshplay Adapter Operation Registration_

### Clients

Meshplay's REST API may be consumed by any number of clients. Clients need to present valid JWT token.

<a href="{{ site.baseurl }}/assets/img/architecture/Meshplay-client-architecture.svg" class="lightbox-image">
<img src="{{ site.baseurl }}/assets/img/architecture/Meshplay-client-architecture.svg" width="50%" /></a>


_Figure: Clients use Meshplay's [REST API](extensibility/api#rest), [GraphQL API](extensibility/api#graphql), or a combination of both._

### Providers

As a point of extensibility, Meshplay supports two types of [providers](/extensibility/providers): _Local_ and _Remote_.

<a href="{{ site.baseurl }}/assets/img/architecture/Meshplay-provider-architecture.svg" class="lightbox-image">
<img src="{{ site.baseurl }}/assets/img/architecture/Meshplay-provider-architecture.svg" width="50%" /></a>
<figure>
  <figcaption>Figure: Meshplay Provider architecture</figcaption>
</figure>

## Object Model

This diagram outlines logical constructs within Meshplay and their relationships.

<a href="{{ site.baseurl }}/assets/img/architecture/meshplay_extension_points.svg" class="lightbox-image">
<img src="{{ site.baseurl }}/assets/img/architecture/meshplay_extension_points.svg" width="50%" /></a>
<figure>
  <figcaption>Figure: Meshplay Object Model</figcaption>
</figure>

## Meshplay Operator and MeshSync

Meshplay Operator is the multi-cluster Kubernetes operator that manages MeshSync and Meshplay Broker.

<a href="{{ site.baseurl }}/assets/img/architecture/meshplay-operator-and-meshsync.svg" class="lightbox-image">
<img src="{{ site.baseurl }}/assets/img/architecture/meshplay-operator-and-meshsync.svg" width="50%" /></a>
<figure>
  <figcaption>Figure: Meshplay Operator and MeshSync</figcaption>
</figure>

See the [**Operator**]({{ site.baseurl }}/concepts/architecture/operator) section for more information on the function of an operator and [**MeshSync**]({{ site.baseurl }}/concepts/architecture/meshsync) section for more information on the function of meshsync.

## Database

Meshplay Server's database is responsible for collecting and centralizing the state of all elements under management, including infrastructure, application, and Meshplay's own components. Meshplay's database, while persisted to file, is treated as a cache.

<a href="{{ site.baseurl }}/assets/img/architecture/meshplay-database.svg" class="lightbox-image">
<img src="{{ site.baseurl }}/assets/img/architecture/meshplay-database.svg" width="50%" /></a>
<figure>
  <figcaption>Figure: Meshplay Docker Extension</figcaption>
</figure>

_See the [**Database**]({{ site.baseurl }}/concepts/architecture/database) section for more information on the function of the database._

## Meshplay Docker Extension 

Meshplay's Docker extension provides a simple and flexible way to design and operate cloud native infrastructure on top of Kubernetes using Docker containers. The architecture of this extension is designed to be modular and extensible, with each component serving a specific purpose within the overall deployment process.

<a href="{{ site.baseurl }}/assets/img/architecture/meshplay-docker-extension.svg" class="lightbox-image">
<img src="{{ site.baseurl }}/assets/img/architecture/meshplay-docker-extension.svg" width="50%" /></a>
<figure>
  <figcaption>Figure: Meshplay Docker Extension</figcaption>
</figure>

## Meshplay CLI 

The Command Line Interface ( also known as [meshplayctl](/guides/meshplayctl/working-with-meshplayctl) ) that is used to manage Meshplay. Use `meshplayctl` to both manage the lifecycle of Meshplay itself and to access and invoke any of Meshplay's application and cloud native management functions.


### **Statefulness in Meshplay components**

Some components within Meshplay's architecture are concerned with persisting data while others are only
concerned with a long-lived configuration, while others have no state at all.

| Components        | Persistence  | Description                                                           |
| :---------------- | :----------- | :-------------------------------------------------------------------- |
| [meshplayctl](/guides/meshplayctl/working-with-meshplayctl)        | stateless    | command line interface that has a configuration file                  |
| [Meshplay Adapters](/concepts/architecture/adapters)  | stateless    | interface with cloud native infrastructure on a transactional basis                |
| Meshplay Server    | caches state | application cache is stored in `$HOME/.meshplay/` folder               |
| [Meshplay Providers](/extensibility/providers) | stateful     | location of persistent user preferences, environment, tests and so on |
| [Meshplay Operator](/concepts/architecture/operator)  | stateless    | operator of Meshplay custom controllers, notably MeshSync              |
| [MeshSync](/concepts/architecture/meshsync)          | stateless    | Kubernetes custom controller, continuously running discovery          |

### **Network Ports**

Meshplay uses the following list of network ports to interface with its various components:

{% for adapter in site.adapters -%}
{% if adapter.port -%}
{% capture adapter-ports %}
| <img src="{{ adapter.image }}" style="width:20px" /> [{{ adapter.name }}]({{ site.baseurl }}{{ adapter.url }}) | {{ adapter.port }}/gRPC | Communication with Meshplay Server |
{% endcapture %}
{% endif -%}
{% endfor %}

| Component                |   Port   | Purpose                                         |
| :----------------------- | :------: | :-----------------------------------------------|
| Meshplay Server          | 9081/tcp | UI, REST and GraphQL APIs                           |
| Meshplay Server          | 80/tcp | Websocket                          |
| [Meshplay Broker](/concepts/architecture/broker)           | 4222/tcp | Client communication with Meshplay Server        |
| [Meshplay Broker](/concepts/architecture/broker)            | 8222/tcp | HTTP management port for monitoring Meshplay Broker. Available as of Meshplay v0.5.0 |
| [Meshplay Broker](/concepts/architecture/broker)            | 6222/tcp | Routing port for Broker clustering. Unused as of Meshplay v0.6.0-rc-2             |
| [Meshplay Broker](/concepts/architecture/broker)            | 7422/tcp | Incoming/outgoing leaf node connections. Unused as of Meshplay v0.6.0-rc-2 |
| [Meshplay Broker](/concepts/architecture/broker)            | 7522/tcp | Gateway to gateway communication. Unused as of Meshplay v0.6.0-rc-2 |
| [Meshplay Broker](/concepts/architecture/broker)            | 7777/tcp | used for Prometheus NATS Exporter. Unused as of Meshplay v0.6.0-rc-2 |
| Learn KhulnaSoft Application | 10011/tcp  | SMI conformance testing                        |
| [Meshplay Remote Providers]((/extensibility/providers)) | 443/tcp    | e.g. Meshplay Cloud                             |
{% for adapter in site.adapters -%}
{% if adapter.port -%}
| <img src="{{ adapter.image }}" style="width:20px" data-logo-for-dark="{{ adapter.white_image }}" data-logo-for-light="{{ adapter.image }}" id="logo-dark-light" loading="lazy"/> [{{ adapter.name }}]({{ site.baseurl }}{{ adapter.url }}) | {{ adapter.port }} | Communication with Meshplay Server |
{% endif -%}
{% endfor -%}
| [Meshplay Perf]({{ site.baseurl }}/tasks/performance/managing-performance) | 10013/gRPC    | Performance Management|

See the [**Adapters**]({{ site.baseurl }}/concepts/architecture/adapters) section for more information on the function of an adapter.


### **Meshplay Connections and their Actions**

<table style=" padding-right: 10px;
        margin: 5px 5px 5px 5px;
        display: block;
        max-width: fit-content;
        overflow-x: auto;
        white-space: nowrap;">
  <thead>
    <tr>
      <th>Connection Type</th>
      <th>&nbsp;</th>
      <th>&nbsp;</th>
      <th>Action / Behaviour</th>
      <th>&nbsp;</th>
      <th>&nbsp;</th>
      <th>&nbsp;</th>
      <th>&nbsp;</th>
      <th>&nbsp;</th>
      <th>&nbsp;</th>
      <th>&nbsp;</th>
      <th>&nbsp;</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>&nbsp;</td>
      <td><strong>Connect meshplayctl</strong></td>
      <td><strong>Connect Meshplay UI</strong></td>
      <td><strong>Disconnect</strong></td>
      <td><strong>Ad hoc Connectivity Test</strong></td>
      <td><strong>Ongoing Connectivity Test</strong></td>
      <td><strong>Synthetic Check</strong></td>
      <td><strong>Deploy meshplayctl</strong></td>
      <td><strong>Undeploy meshplayctl</strong></td>
      <td><strong>Deploy Meshplay UI</strong></td>
      <td><strong>Undeploy Meshplay UI</strong></td>
      <td>&nbsp;</td>
    </tr>
    <tr>
      <td>Kubernetes clusters</td>
      <td>`system start`</td>
      <td>Upload kubeconfig</td>
      <td>Click "X" on chip</td>
      <td>On click of connection chip</td>
      <td>Yes, via MeshSync</td>
      <td>No</td>
      <td>No</td>
      <td>No</td>
      <td>No</td>
      <td>No</td>
      <td>&nbsp;</td>
    </tr>
    <tr>
      <td>Grafana Servers</td>
      <td>No</td>
      <td>Enter IP/hostname into Meshplay UI</td>
      <td>Click "X" on chip</td>
      <td>On click of connection chip</td>
      <td>No</td>
      <td>No</td>
      <td>No</td>
      <td>No</td>
      <td>No</td>
      <td>No</td>
      <td>&nbsp;</td>
    </tr>
    <tr>
      <td>Prometheus Servers</td>
      <td>No</td>
      <td>Enter IP/hostname into Meshplay UI</td>
      <td>Click "X" on chip</td>
      <td>On click of connection chip</td>
      <td>Yes, when metrics are configured in a dashboard</td>
      <td>Yes</td>
      <td>No</td>
      <td>No</td>
      <td>No</td>
      <td>No</td>
      <td>&nbsp;</td>
    </tr>
    <tr>
      <td><a href="/concepts/architecture/adapters">Meshplay Adapters</a></td>
      <td>`system check`</td>
      <td>Server to Adapter on every UI refresh</td>
      <td>Click "X on" chip</td>
      <td>Server to Adapter every click on adapter chip in UI</td>
      <td>Server to Adapter every 10 seconds</td>
      <td>-</td>
      <td>Yes, as listed in meshconfig contexts</td>
      <td>Yes, as listed in meshconfig contexts</td>
      <td>Toggle switch needed</td>
      <td>Toggle switch needed</td>
      <td>&nbsp;</td>
    </tr>
    <tr>
      <td><a href="/concepts/architecture/operator">Meshplay Operator</a></td>
      <td>`system check`</td>
      <td>Upon upload of kubeconfig</td>
      <td>No</td>
      <td>On click of connection chip in UI to Server to Kubernetes to Meshplay Operator</td>
      <td>No</td>
      <td>-</td>
      <td>`system start`</td>
      <td>`system stop`</td>
      <td>Upon upload of kubeconfig & Toggle of switch</td>
      <td>Toggle of switch</td>
      <td>&nbsp;</td>
    </tr>
    <tr>
      <td><a href="/concepts/architecture/meshsync">MeshSync</a></td>
      <td>`system check`</td>
      <td>follows the lifecycle of Meshplay Operator</td>
      <td>No</td>
      <td>On click of connection chip in UI to Server to Kubernetes to Meshplay Operator to MeshSync</td>
      <td>Managed by Meshplay Operator</td>
      <td>On click of connection chip</td>
      <td>follows the lifecycle of Meshplay Operator</td>
      <td>follows the lifecycle of Meshplay Operator</td>
      <td>follows the lifecycle of Meshplay Operator</td>
      <td>follows the lifecycle of Meshplay Operator</td>
      <td>&nbsp;</td>
    </tr>
    <tr>
      <td><a href="/concepts/architecture/broker">Broker</a></td>
      <td>`system check`</td>
      <td>follows the lifecycle of Meshplay Operator</td>
      <td>No</td>
      <td>On click of connection chip in UI to Server to Brokers exposed service port</td>
      <td>NATS Topic Subscription</td>
      <td>On click of connection chip</td>
      <td>follows the lifecycle of Meshplay Operator</td>
      <td>follows the lifecycle of Meshplay Operator</td>
      <td>follows the lifecycle of Meshplay Operator</td>
      <td>follows the lifecycle of Meshplay Operator</td>
      <td>&nbsp;</td>
    </tr>
  </tbody>
</table>
<br>

Please also see the [Troubleshooting Toolkit](https://docs.google.com/document/d/1q-aayRqx3QKIk2soTaTTTH-jmHcVXHwNYFsYkFawaME/edit#heading=h.ngupcd4j1pfm) and the [Meshplay v0.7.0: Connection States (Kubnernetes) Design Review](http://discuss.meshplay.khulnasofy.com/t/meshplay-v0-7-0-connection-states-kubnernetes-design-review/958)
