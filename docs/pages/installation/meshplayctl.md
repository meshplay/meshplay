---
layout: default
title: Install meshplayctl
permalink: installation/meshplayctl
type: installation
category: meshplayctl
redirect_from:
- installation/meshplayctl/
- installation/platforms/meshplayctl
display-title: "true"
language: en
list: exclude
suggested-reading: false
# image: /assets/img/platforms/brew.png
abstract: Install Meshplay CLI
---

Meshplay's command line client is `meshplayctl` and is the recommended tool for configuring and deploying one or more Meshplay deployments. To install `meshplayctl` on your system, you may choose from any of the following supported methods.

`meshplayctl` can be installed via [bash]({{site.baseurl}}/installation/linux-mac/bash), [Homebrew]({{site.baseurl}}/installation/linux-mac/brew), [Scoop]({{site.baseurl}}/installation/windows/scoop) or [directly downloaded](https://github.com/meshplay/meshplay/releases/latest).

# Install Meshplay CLI with Brew

{% include meshplayctl/installation-brew.md %}

# Install Meshplay CLI with Bash

{% include meshplayctl/installation-bash.md %}

# Install Meshplay CLI with Scoop

{% include meshplayctl/installation-scoop.md %}

Continue deploying Meshplay onto one of the [Supported Platforms]({{ site.baseurl }}/installation).

# Related Reading

## Meshplay CLI Guides

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

{% include related-discussions.html tag="meshplayctl" %}

{:toc}
