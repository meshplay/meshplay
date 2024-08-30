---
layout: default
title: KinD
permalink: installation/kubernetes/kind
type: installation
category: kubernetes
redirect_from:
- installation/platforms/kind
display-title: "false"
language: en
list: include
image: /assets/img/platforms/kind.png
abstract: Install Meshplay on KinD. Deploy Meshplay in KinD in-cluster or outside of KinD out-of-cluster.
---

<h1>Quick Start with {{ page.title }} <img src="{{ page.image }}" style="width:35px;height:35px;" /></h1>

Manage your KinD clusters with Meshplay. Deploy Meshplay in your [KinD cluster](#in-cluster-installation).

<div class="prereqs"><h4>Prerequisites</h4>
<ol>
<li>Install the Meshplay command line client, <a href="{{ site.baseurl }}/installation/meshplayctl" class="meshplay-light">meshplayctl</a>.</li>
<li>Install <a href="https://kubernetes.io/docs/tasks/tools/">kubectl</a> on your local machine.</li>
<li>Install <a href="https://kind.sigs.k8s.io/docs/user/quick-start/#installation">KinD</a>, on your local machine.</li>
</ol>
</div>

Also see: [Install Meshplay on Kubernetes]({{ site.baseurl }}/installation/kubernetes)

## Available Deployment Methods

- [Available Deployment Methods](#available-deployment-methods)
- [In-cluster Installation](#in-cluster-installation)
  - [Preflight Checks](#preflight-checks)
    - [Preflight: Cluster Connectivity](#preflight-cluster-connectivity)
  - [Installation: Using `meshplayctl`](#installation-using-meshplayctl)
  - [Alternative Installation: Using Helm](#alternative-installation-using-helm)
  - [Post-Installation Steps](#post-installation-steps)

## In-cluster Installation

Follow the steps below to install Meshplay in your KinD cluster.

### Preflight Checks

Read through the following considerations prior to deploying Meshplay on KinD.

#### Preflight: Cluster Connectivity

Start the KinD, if not started using the following command:
{% capture code_content %}kind create cluster{% endcapture %}
{% include code.html code=code_content %}
Check up on your KinD cluster :
{% capture code_content %}kind get clusters{% endcapture %}
{% include code.html code=code_content %}
Verify your kubeconfig's current context.
{% capture code_content %}kubectl config current-context{% endcapture %}
{% include code.html code=code_content %}

### Installation: Using `meshplayctl`

<details>
<summary>Verify your Meshplay context</summary>
<p>
Verify that your current Meshplay context is set for an in-cluster deployment (`platform: kubernetes`) by executing:
</p>

{% capture code_content %}$ meshplayctl system context view{% endcapture %}
{% include code.html code=code_content %}
<p>
If the context is not set to <code>platform: kubernetes</code>, you can create a new context with Kubernetes as the platform using the following command.
</p>

{% capture code_content %}$ meshplayctl system context create context-name --platform kubernetes --url http://localhost:9081 --set --yes{% endcapture %}
{% include code.html code=code_content %}
<br/>
</details>

With your KIND cluster configured your `current-context`, start Meshplay.

{% capture code_content %}$ meshplayctl system start -p kubernetes{% endcapture %}
{% include code.html code=code_content %}

### Alternative Installation: Using Helm

See [Helm Installation](/installation/kubernetes/helm) guide.

### Post-Installation Steps

Meshplay deploys with LoadBalancer service type by default. If you are using KinD, you may need to expose the Meshplay service. A universal option is to use `meshplayctl system dashboard --port-forward`. A KIND-specific option to use use the [Cloud Provider KIND](https://kind.sigs.k8s.io/docs/user/loadbalancer/). Cloud Provider KIND runs as a standalone binary in your host and connects to your KIND cluster and provisions new Load Balancer containers for your Services.

{% include meshplayctl/system-dashboard.md %}

Optionally, you can verify the health of your Meshplay deployment, using <a href='/reference/meshplayctl/system/check'>meshplayctl system check</a>.

You're ready to use Meshplay! Open your browser and navigate to the Meshplay UI.

{% include_cached installation/accessing-meshplay-ui.md display-title="true" %}

{% include related-discussions.html tag="meshplay" %}