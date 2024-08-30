---
layout: default
title: Scoop
permalink: installation/windows/scoop
type: installation
category: meshplayctl
redirect_from:
- installation/platforms/scoop
display-title: "false"
language: en
list: include
image: /assets/img/platforms/scoop.png
abstract: Install Meshplay CLI on Windows with Scoop
---
# Install Meshplay CLI with Scoop

{% include meshplayctl/installation-scoop.md %}

# Related Reading

## Mesherctl Guides

Guides to using Meshplay's various features and components.

{% assign sorted_guides = site.pages | sort: "name" %}

<ul>
  {% for item in sorted_guides %}
  {% if item.type=="guides" and item.category=="meshplayctl" and item.list!="exclude" and item.language=="en" -%}
    <li><a href="{{ site.baseurl }}{{ item.url }}">{{ item.title }}</a>
    </li>
    {% endif %}
  {% endfor %}
    <li><a href="{{ site.baseurl }}/installation/upgrades#upgrading-meshplay-cli">Upgrading Meshplay CLI</a></li>
</ul>

