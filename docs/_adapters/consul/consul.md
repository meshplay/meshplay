---
layout: default
title: Meshplay Adapter for Consul
name: Meshplay Adapter for Consul
component: Consul
earliest_version: v1.8.4
port: 10002/gRPC
project_status: stable
lab: consul-meshplay-adapter
github_link: https://github.com/meshplay/meshplay-consul
image: /assets/img/service-meshes/consul.svg
white_image: /assets/img/service-meshes/consul-white.svg
language: en
permalink: extensibility/adapters/consul
redirect_from: service-meshes/adapters/consul
---

{% assign sorted_tests_group = site.compatibility | group_by: "meshplay-component" %}
{% for group in sorted_tests_group %}
      {% if group.name == "meshplay-consul" %}
        {% assign items = group.items | sort: "meshplay-component-version" | reverse %}
        {% for item in items %}
          {% if item.meshplay-component-version != "edge" %}
            {% if item.overall-status == "passing" %}
              {% assign adapter_version_dynamic = item.meshplay-component-version %}
              {% break %}
            {% elsif item.overall-status == "failing" %}
              {% continue %}
            {% endif %}
          {% endif %}
        {% endfor %} 
      {% endif %}
{% endfor %}

{% include compatibility/adapter-status.html %}

<!-- {% include adapter-labs.html %} -->

## Lifecycle management

The {{page.name}} can install **{{page.earliest_version}}** of the {{page.component}}.

### Install {{ page.component }}

##### Choose the Meshplay Adapter for {{page.component}}

<a href="{{ site.baseurl }}/assets/img/adapters/consul/consul-adapter.png">
  <img style="width:500px;" src="{{ site.baseurl }}/assets/img/adapters/consul/consul-adapter.png" />
</a>

##### Click on (+) and choose the {{page.earliest_version}} of the {{page.component}}.

<a href="{{ site.baseurl }}/assets/img/adapters/consul/consul-install.png">
  <img style="width:500px;" src="{{ site.baseurl }}/assets/img/adapters/consul/consul-install.png" />
</a>

A number of [sample applications](#sample-applications) for {{page.component}} can also be installed using Meshplay.

### Features

1. Lifecycle management of {{page.component}}
1. Lifecycle management of sample applications
1. Performance management of {{page.component}} and it workloads
   - Prometheus and Grafana integration
1. Configuration management and best practices of {{page.component}}
1. Custom configuration

### Sample Applications

Meshplay supports the deployment of a variety of sample applications on {{ page.name }}. Use Meshplay to deploy any of these sample applications.

- [httpbin]({{site.baseurl}}/guides/sample-apps#httpbin)
  - Httpbin is a simple HTTP request and response service.
- [Bookinfo]({{site.baseurl}}/guides/sample-apps#bookinfo)
  - The sample BookInfo application displays information about a book, similar to a single catalog entry of an online book store.
- [Image Hub]({{site.baseurl}}/guides/sample-apps#imagehub)
  - Image Hub is a sample application written to run on Consul for exploring WebAssembly modules used as Envoy filters.

[![Layer5 Image Hub on HashiCorp Consul]({{ site.baseurl }}/extensibility/adapters/consul/layer5-image-hub-on-hashicorp-consul.png)]({{ site.baseurl }}/extensibility/adapters/consul/layer5-image-hub-on-hashicorp-consul.png)

### Performance management of Consul and it workloads

#### Prometheus and Grafana integration

The {{ page.name }} will connect to {{ page.name }}'s Prometheus and Grafana instances running in the control plane (typically found in a separate namespace) or other instances to which Meshplay has network reachability.

### Architecture

[![Consul Service Mesh Architecture]({{ site.baseurl }}/extensibility/adapters/consul/service-mesh-architecture-consul.png)]({{ site.baseurl }}/extensibility/adapters/consul/service-mesh-architecture-consul.png)

### Suggested Topics

- Examine [Meshplay's architecture]({{ site.baseurl }}/architecture) and how adapters fit in as a component.
- Learn more about [Meshplay Adapters]({{ site.baseurl }}/architecture/adapters).
