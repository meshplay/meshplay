---
layout: default
title: Meshplay Adapter for Traefik Mesh
name: Meshplay Adapter for Traefik Mesh
component: Traefik Mesh
earliest_version: v1.0
port: 10006/gRPC
project_status: stable
lab: traefik-meshplay-adapter
github_link: https://github.com/meshplay/meshplay-traefik-mesh
image: /assets/img/service-meshes/traefik-mesh.svg
white_image: /assets/img/service-meshes/traefik-mesh.svg
permalink: extensibility/adapters/traefik-mesh
redirect_from: service-meshes/adapters/traefik-mesh
language: en
---

{% assign sorted_tests_group = site.compatibility | group_by: "meshplay-component" %}
{% for group in sorted_tests_group %}
      {% if group.name == "meshplay-traefik-mesh" %}
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

The {{page.name}} can install **{{page.earliest_version}}** of {{page.component}}. A number of sample applications can be installed using the {{page.name}}.

The {{ page.name }} is currently under construction ({{ page.project_status }} state), which means that the adapter is not functional and cannot be interacted with through the <a href="{{ site.baseurl }}installation#6-you-will-now-be-directed-to-the-meshplay-ui"> Meshplay UI </a>at the moment. Check back here to see updates.

Want to contribute? Check our [progress]({{page.github_link}}).

### Sample Applications

The {{ page.name }} includes some sample applications operations. Meshplay can be used to deploy any of these sample applications.

- [Bookinfo]({{site.baseurl}}/guides/sample-apps#bookinfo)
  - Follow this [tutorial workshop](https://github.com/khulnasoft/istio-service-mesh-workshop/blob/master/lab-2/README.md) to set up and deploy the BookInfo sample app on Istio using Meshplay.
- [Httpbin]({{site.baseurl}}/guides/sample-apps#httpbin)
  - Httpbin is a simple HTTP request and response service.

## Suggested Topics

- Examine [Meshplay's architecture]({{ site.baseurl }}/architecture) and how adapters fit in as a component.
- Learn more about [Meshplay Adapters]({{ site.baseurl }}/architecture/adapters).
