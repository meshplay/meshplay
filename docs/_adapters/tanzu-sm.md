---
layout: default
title: Meshplay Adapter for Tanzu Service Mesh
name: Meshplay Adapter for Tanzu Service Mesh
component: Tanzu Service Mesh
earliest_version: pre-GA
port: 10011/gRPC
project_status: alpha
github_link: https://github.com/meshplay/meshplay-tanzu-sm
image: /assets/img/service-meshes/tanzu.svg
white_image: /assets/img/service-meshes/tanzu.svg
permalink: extensibility/adapters/tanzu-sm
redirect_from: service-meshes/adapters/tanzu-sm
language: en
---

{% assign sorted_tests_group = site.compatibility | group_by: "meshplay-component" %}
{% for group in sorted_tests_group %}
      {% if group.name == "meshplay-tanzu-mesh" %}
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

## Lifecycle management

The {{page.name}} can install **{{page.earliest_version}}** of {{page.component}}. A number of sample applications for {{page.component}} can also be installed using Meshplay.

The {{ page.name }} is currently under construction ({{ page.project_status }} state), which means that the adapter is not functional and cannot be interacted with through the <a href="{{ site.baseurl }}/installation#6-you-will-now-be-directed-to-the-meshplay-ui"> Meshplay UI </a>at the moment. Check back here to see updates.

Want to contribute? Check our [progress]({{page.github_link}}).


### Suggested Reading

- Examine [Meshplay's architecture]({{ site.baseurl }}/architecture) and how adapters fit in as a component.
- Learn more about [Meshplay Adapters]({{ site.baseurl }}/architecture/adapters).
