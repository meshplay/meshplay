---
layout: default
title: Meshplay Overview
permalink: project/overview
# redirect_from: project/overview/
language: en
display-title: "false"
type: project
category: project
list: exclude
published: true
abstract: Meshplay is the self-service engineering platform, enabling collaborative design and operation of cloud and cloud native infrastructure.
---

## Meshplay is for all cloud and cloud native infrastructure

Meshplay is an **extensible engineering platform** for the collaborative design and operation of cloud and cloud native infrastructure and applications.

Kubernetes-centric. Kubernetes not required.

Infrastructure diversity is a reality for any enterprise. Whether you're running a single Kubernetes cluster or multiple Kubernetes clusters, on one cloud or multiple clouds, you'll find that Meshplay supports your infrastructure diversity (or lack thereof).

## Meshplay's Functionality

Meshplay supports all Kubernetes-based infrastructure including many cloud services of AWS and GCP platforms. Meshplay features can be categorized by:

### Lifecycle Management (Day 0, Day 1)

- Cloud and cloud native provisioning
- Discovery and onboarding of existing environments and workloads
- Registry and configuration of WebAssembly filters for Envoy

### Configuration Management (Day 2)

- Cloud native patterns catalog
- Configuration best practices
- Policy engine for relationship inference and context-aware design

### Collaboration

- Multi-player infrastructure design and operation

### Performance Management

- Workload and performance characterization with both built-in and external load generators
- Prometheus and Grafana integration

### Interoperability and Federation

- Integration with thousands of cloud services and cloud native projects
- Manage multiple cloud and cloud native environments concurrently
- Connect to multiple clusters independently

### Design patterns and Meshplay Catalog

Through [Models]({{site.baseurl}}/concepts/logical/models), Meshplay describes infrastructure under management, enabling you to define cloud native designs and patterns and then to export those designs and share within the <a href="https://meshplay.io/catalog" target="_self_">Meshplay Catalog</a>.

## Meshplay is for engineering teams

Whether you are a Platform Engineer, Site Reliability Engineer, DevOps Engineer, Developer, or Operator, Meshplay provides a platform for you to collaborate on the design and operation of your cloud native infrastructure.

Whether making a Day 0 adoption choice, a Day 1 configuration and provisioning, or maintaining a Day 2 deployment, Meshplay has useful capabilities in either circumstance. Targeted audience for Meshplay project would be any technology operators that leverage Cloud and cloud native infrastructure.

<!-- 
### Meshplay is for performance management

Meshplay helps users weigh the value of their cloud native deployments against the overhead incurred in running different deployment scenarios and different configruations. Meshplay provides statistical analysis of the request latency and throughput seen across various permutations of your workload, infrastructure and infrastructure configuration. In addition to request latency and throughput, Meshplay also tracks memory and CPU overhead in of the nodes in your cluster. Establish a performance benchmark and track performance against this baseline as your environment changes over time. 
-->

<!-- ### Supported Integrations

#### **Stable**

| Adapter | Status |
| :----------- | -----: |
{% for adapter in site.adapters -%}
{% if adapter.project_status == "stable" -%}
| <img src="{{ adapter.image }}" style="width:20px" /> [{{ adapter.name }}]({{ site.baseurl }}{{ adapter.url }}) | {{ adapter.project_status }} |
{% endif -%}
{% endfor %}

##### **Beta**

| Adapter | Status |
| :----------- | -----: |
{% for adapter in site.adapters -%}
{% if adapter.project_status == "beta" -%}
| <img src="{{ adapter.image }}" style="width:20px" /> [{{ adapter.name }}]({{ site.baseurl }}{{ adapter.url }}) | {{ adapter.project_status }} |
{% endif -%}
{% endfor %}

##### **Alpha** - Meshplay adapters for which we are seeking community-contributed support.

| Adapter | Status |
| :----------- | -----: |
{% for adapter in site.adapters -%}
{% if adapter.project_status == "alpha" -%}
| <img src="{{ adapter.image }}" style="width:20px" /> [{{ adapter.name }}]({{ site.baseurl }}{{ adapter.url }}) | {{ adapter.project_status }} |
{% endif -%}
{% endfor %}
 -->

<!-- ## Meshplay as a project and its community

{% assign sorted_pages = site.pages | sort: "type" | reverse %}

<ul>
    {% for item in sorted_pages %}
    {% if item.type=="project" and item.language=="en" and item.list != "exclude" %}
      <li><a href="{{ site.baseurl }}{{ item.url }}">{{ item.title }}</a>
      {% if item.description != " " %}
        -  {{ item.description }}
      {% endif %}
      </li>
      {% endif %}
    {% endfor %}
</ul> 
-->
