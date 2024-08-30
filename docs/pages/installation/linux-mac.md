---
layout: default
title: Install Meshplay CLI on Linux or Mac
permalink: installation/linux-mac
type: installation
category: meshplayctl
redirect_from:
- installation/platforms/linux-mac
display-title: "false"
language: en
list: include
image: /assets/img/platforms/linux_mac.png 
abstract: Install Meshplay CLI on Linux or Mac
---

# Overview

To set up and run Meshplay on Linux or macOS, you will need to install `meshplayctl`. `meshplayctl` is the command line interface (CLI) for Meshplay. It is used to install, manage, and operate one or more Meshplay deployments. `meshplayctl` can be installed via `bash` is also available [directly](https://github.com/meshplay/meshplay/releases/latest) or through [Homebrew]({{site.baseurl}}/installation/linux-mac/brew) or [Scoop]({{site.baseurl}}/installation/windows/scoop).

# Brew

{% include meshplayctl/installation-brew.md %}

# Bash

{% include meshplayctl/installation-bash.md %}

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

<!-- 
1. You can either use **Bash** or **Brew** to install <a href="/guides/meshplayctl">meshplayctl</a> ( Meshplay command line interface ).
2. To run **Meshplay**, execute the following command.

   <pre class="codeblock-pre"><div class="codeblock">
   <div class="clipboardjs">meshplayctl system start</div></div>
   </pre>

Meshplay server supports customizing authentication flow callback URL, which can be configured in the following way

<pre class="codeblock-pre"><div class="codeblock">
<div class="clipboardjs">
 $ MESHPLAY_SERVER_CALLBACK_URL=https://custom-host meshplayctl system start

</div></div>
</pre>
-->