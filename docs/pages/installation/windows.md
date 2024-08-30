---
layout: default
title: Install Meshplay CLI on Windows
permalink: installation/windows
type: installation
category: meshplayctl
redirect_from:
- installation/platforms/windows
display-title: "true"
language: en
list: include
image: /assets/img/platforms/wsl2.png
abstract: Install Meshplay CLI on Windows
---


On Windows systems, `meshplayctl` can be installed via Scoop or can be [downloaded directly](https://github.com/meshplay/meshplay/releases/latest).

{% include meshplayctl/installation-scoop.md %}

## Install `meshplayctl` as a direct download

Follow the [installation steps]({{ site.baseurl }}/installation#windows) to install the meshplayctl CLI. Then, execute:
<pre class="codeblock-pre">
<div class="codeblock"><div class="clipboardjs">./meshplayctl system start</div></div>
</pre>

Optionally, move the meshplayctl binary to a directory in your PATH.


<!-- Meshplay server supports customizing authentication flow callback URL, which can be configured in the following way
  <pre class="codeblock-pre">
  <div class="codeblock"><div class="clipboardjs">MESHPLAY_SERVER_CALLBACK_URL=https://custom-host ./meshplayctl system start</div></div>
  </pre>

Type `yes` when prompted to choose to configure a file. To get started, choose Docker as your platform to deploy Meshplay. -->

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