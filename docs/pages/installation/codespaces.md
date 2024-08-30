---
layout: default
title: Codespaces
permalink: installation/codespaces
type: installation
category: kubernetes
redirect_from:
- installation/platforms/codespaces
display-title: "false"
language: en
list: include
image: /assets/img/platforms/codespaces.png
abstract: Build and contribute to Meshplay using GitHub Codespaces
---

<h1>Quick Start with {{ page.title }} <img src="{{ page.image }}" style="width:35px;height:35px;" /></h1>

Use Minikube in GitHub Codespace to setup your development environment for Meshplay.

<div class="prereqs"><p><strong style="font-size: 20px;">Prerequisites</strong> </p> 
  <ol>
    <li>Install the Meshplay command line client, <a href="{{ site.baseurl }}/installation/meshplayctl" class="meshplay-light">meshplayctl</a>.</li>
  </ol>
</div>

## Available Deployment Methods

- [In-cluster Installation](#in-cluster-installation)
  - [Preflight Checks](#preflight-checks)
    - [Preflight: Cluster Connectivity](#preflight-cluster-connectivity)
  - [Installation: Using `meshplayctl`](#installation-using-meshplayctl)
  - [Installation: Using Helm](#installation-using-helm)
  - [Installation: Manual Steps](#installation-manual-steps)
- [Post-Installation Steps](#post-installation-steps)

# In-cluster Installation

Follow the steps below to install Meshplay in your Minikube cluster.

## Preflight Checks

Read through the following considerations prior to deploying Meshplay on Minikube.

### Preflight: Cluster Connectivity


You can develop and run Meshplay in a GitHub Codespace using your choice of tool:

- A command shell, via an SSH connection initiated using GitHub CLI.
- One of the JetBrains IDEs, via the JetBrains Gateway.
- The Visual Studio Code desktop application.
- A browser-based version of Visual Studio Code.

{% include alert.html type="dark" title="Choice of Codespace Tool" content="For the best experience, run Codespace in your locally <a href='https://docs.github.com/en/codespaces/developing-in-codespaces/developing-in-a-codespace'>installed IDE</a>. Alternatively, you can <br /><a href='https://github.com/codespaces/new?hide_repo_select=true&ref=master&repo=157554479&machine=premiumLinux'><img alt='Open in GitHub Codespaces' src='https://github.com/codespaces/badge.svg' /></a>" %}

Start the minikube, if not started using the following command:
{% capture code_content %}minikube start --cpus 4 --memory 4096{% endcapture %}
{% include code.html code=code_content %}
Please allocate cpus based on the machine you selected in the Github codespaces and to check up on your minikube cluster :
{% capture code_content %}minikube status{% endcapture %}
{% include code.html code=code_content %}
Verify your kubeconfig's current context.
{% capture code_content %}kubectl cluster-info{% endcapture %}
{% include code.html code=code_content %}

## Installation: Using `meshplayctl`

Use Meshplay's CLI to streamline your connection to your Minikube cluster. Configure Meshplay to connect to your Minikube cluster by executing:

{% capture code_content %}$ meshplayctl system config minikube{% endcapture %}
{% include code.html code=code_content %}

Once configured, execute the following command to start Meshplay.

{% capture code_content %}$ meshplayctl system start{% endcapture %}
{% include code.html code=code_content %}

If you encounter any authentication issues, you can use `meshplayctl system login`. For more information, click [here](/guides/meshplayctl/authenticate-with-meshplay-via-cli) to learn more.

## Installation: Using Helm

For detailed instructions on installing Meshplay using Helm V3, please refer to the [Helm Installation](/installation/helm) guide.

## Installation: Manual Steps

You may also manually generate and load the kubeconfig file for Meshplay to use:

**The following configuration yaml will be used by Meshplay. Copy and paste the following in your config file** :

{% capture code_content %}apiVersion: v1
clusters:

- cluster:
  certificate-authority-data: < cert shortcutted >
  server: https://192.168.99.100:8443
  name: minikube
  contexts:
- context:
  cluster: minikube
  user: minikube
  name: minikube
  current-context: minikube
  kind: Config
  preferences: {}
  users:
- name: minikube
  user:
  client-certificate-data: < cert shortcutted >
  client-key-data: < key shortcutted >{% endcapture %}
  {% include code.html code=code_content %}

_Note_: Make sure _current-context_ is set to _minikube_.

<br />
**To allow Meshplay to auto detect your config file, Run** :
{% capture code_content %}kubectl config view --minify --flatten > config_minikube.yaml{% endcapture %}
{% include code.html code=code_content %}

<br />
Meshplay should now be connected with your managed Kubernetes instance. Take a look at the [Meshplay guides]({{ site.baseurl }}/guides) for advanced usage tips.

# Post-Installation Steps

Otionally, you can verify the health of your Meshplay deployment, using <a href='/reference/meshplayctl/system/check'>meshplayctl system check</a>.

You're ready to use Meshplay! Open your browser and navigate to the Meshplay UI.

{% include_cached installation/accessing-meshplay-ui.md %}

For further information to access meshplay-ui/port-forwarding in Github Codespace, read the [docs](https://docs.github.com/en/codespaces/developing-in-a-codespace/forwarding-ports-in-your-codespace?tool=vscode)

{% include related-discussions.html tag="meshplay" %}