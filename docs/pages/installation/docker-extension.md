---
layout: default
title: Docker Extension
permalink: installation/docker/docker-extension
type: installation
category: docker
redirect_from:
- installation/platforms/docker-extension
display-title: "false"
language: en
list: include
image: /assets/img/platforms/docker.svg
abstract: Install Docker Extension for Meshplay
---

<h1>Quick Start with {{ page.title }} <img src="{{ page.image }}" style="width:35px;height:35px;" /></h1>

The Docker Extension for Meshplay extends Docker Desktop’s position as the developer’s go-to Kubernetes environment with easy access to full the capabilities of Meshplay's collaborative cloud native management features.

## Install the Docker Meshplay Extension

Select one of the following three options to install the Docker Meshplay Extension:

- [Install the Docker Meshplay Extension](#install-the-docker-meshplay-extension)
  - [Using Docker Desktop](#using-docker-desktop)
  - [Using Docker Hub](#using-docker-hub)
  - [Using Docker CLI](#using-docker-cli)
- [Remove Meshplay as a Docker Extension](#remove-meshplay-as-a-docker-extension)

### Using Docker Desktop

Navigate to the Extensions Marketplace of Docker Desktop. From the Dashboard, select Add Extensions in the menu bar or open the Extensions Marketplace from the menu options.

[![Docker Meshplay Extension]({{site.baseurl}}/assets/img/platforms/docker-desktop-meshplay-extension.png)]({{site.baseurl}}/assets/img/platforms/docker-desktop-meshplay-extension.png)

### Using Docker Hub

You can find the [Docker Meshplay Extension in Docker Hub](https://hub.docker.com/extensions/meshplay/docker-extension-meshplay) marketplace to install the Docker Meshplay Extension.

### Using Docker CLI

Meshplay runs as a set of one or more containers inside your Docker Desktop virtual machine.

<!--
{% capture code_content %}docker extension install meshplay/docker-extension-meshplay{% endcapture %} -->
<!-- {% include code.html code=code_content %} -->

<pre class="codeblock-pre" style="padding: 0; font-size: 0px;">
  <div class="codeblock" style="display: block;">
    <!-- Updated style for clipboardjs -->
    <div class="clipboardjs" style="padding: 0; height: 0.5rem; overflow: hidden;">
      <span style="font-size: 0;">docker extension install meshplay/docker-extension-meshplay</span> 
    </div>
    <div class="window-buttons"></div>
    <div id="termynal2" style="width: 100%; height: 200px; max-width: 100%;" data-termynal="">
      <span data-ty="input">docker extension install meshplay/docker-extension-meshplay</span>
      <span data-ty="progress"></span>
      <span data-ty="">Successfully installed Meshplay</span>
      <span data-ty="input">meshplayctl system dashboard</span>
    </div>
  </div>
</pre>



## Remove Meshplay as a Docker Extension

If you want to remove Meshplay as a Docker extension from your system, follow these steps:

**Stop Meshplay Container:**

- First, stop the running Meshplay container (if it's currently running) using the following Docker command:
{% capture code_content %}$ docker stop meshplay-container{% endcapture %}
{% include code.html code=code_content %}
<br />
    
**Remove Meshplay Container:**

- After stopping the container, you can remove it using the following command:
{% capture code_content %}$ docker rm meshplay-container{% endcapture %}
{% include code.html code=code_content %}
<br />

**Remove Meshplay Images:**

- Meshplay might have pulled Docker images for its components. You can remove these images using the `docker rmi` command. Replace the image names with the actual ones you want to remove:
{% capture code_content %}$ docker rmi meshplay/meshplay:latest{% endcapture %}
{% include code.html code=code_content %}
{% capture code_content %}$ docker rmi meshplay/adapters:latest{% endcapture %}
{% include code.html code=code_content %}
...and so on for other Meshplay-related images
<br />
<br />

**Remove Meshplay Volumes (if necessary):**

- Meshplay may have created Docker volumes to persist data. You can list and remove these volumes using the `docker volume ls` and `docker volume rm` commands. For example:
{% capture code_content %}$ docker volume ls{% endcapture %}
{% include code.html code=code_content %}
{% capture code_content %}$ docker volume rm meshplay-data-volume{% endcapture %}
{% include code.html code=code_content %}
...remove other Meshplay-related volumes if present
<br />
<br />

**Remove Docker Network (if necessary):**

- If Meshplay created a custom Docker network, you can remove it using the `docker network rm` command. For example:
{% capture code_content %}$ docker network rm meshplay-network{% endcapture %}
{% include code.html code=code_content %}
<br />

**Clean Up Configuration (optional):**
- If Meshplay created configuration files or directories on your host machine, you can remove them manually if you no longer need them.

<script src="{{ site.baseurl }}/assets/js/terminal.js" data-termynal-container="#termynal2"></script>

{% include related-discussions.html tag="meshplay" %}
