---
layout: default
title: Extensibility
permalink: extensibility
type: Extensibility
abstract: 'Meshplay has an extensible architecture with several different types of extension points.'
# redirect_from:
#   - reference/extensibility
#   - extensibility/
language: en
list: exclude
---

Meshplay has an extensible architecture with several different types of extension points.

 <!-- via [adapters]({{site.baseurl}}/extensibility/adapters), different [load generators]({{site.baseurl}}/extensibility/load-generators) and different [providers]({{site.baseurl}}/extensibility/providers). Meshplay also offers a REST API. -->

## Extension Points

Meshplay is not just an application. It is a set of microservices where the central component is itself called Meshplay. Integrators may extend Meshplay by taking advantage of designated Extension Points. Extension points come in various forms and are available through Meshplay’s architecture.

![Meshplay Extension Points]({{site.baseurl}}/assets/img/architecture/meshplay_extension_points.svg)

_Figure: Extension points available throughout Meshplay_

<!-- 
1. [Adapters]({{site.baseurl}}/extensibility/adapters)
   -  Messaging Framework (CloudEvents and NATS) 
1. [GraphQL API](/extensibility/api#graphql)
1. [Load Generators]({{site.baseurl}}/extensibility/load-generators)
1. [Providers]({{site.baseurl}}/extensibility/providers)
1. [REST API](/extensibility/api#rest)
1. [UI Plugins](extensibility/ui)
1. [Integrations](/extensibility/integrations)
1. [Extensions](/extensibility/extensions) 
-->

## Types of Extension Points

The following points of extension are currently incorporated into Meshplay.

{% assign sorted = site.pages | sort: "extensibility" %}

<ul>
    {% for item in sorted %}
    {% if item.type=="Extensibility" and item.list!="exclude" and item.language !="es"  -%}
      <li><a href="{{ site.baseurl }}{{ item.url }}">{{ item.title }}</a>
      {% if item.abstract != " " %}
              - {{ item.abstract }}
            {% endif %}
            </li>
            {% endif %}
    {% endfor %}
</ul>
