---
layout: default
title: Minikube
permalink: installation/kubernetes/minikube
type: installation
category: kubernetes
redirect_from:
- installation/platforms/minikube
display-title: "false"
language: en
list: include
image: /assets/img/platforms/minikube.png
abstract: Install Meshplay on Minikube. Deploy Meshplay in Minikube in-cluster or outside of Minikube out-of-cluster.
---

<h1>Quick Start with {{ page.title }} <img src="{{ page.image }}" style="width:35px;height:35px;" /></h1>

Manage your Minikube clusters with Meshplay. Deploy Meshplay in Minikube [in-cluster](#in-cluster-installation) or outside of Minikube [out-of-cluster](#out-of-cluster-installation). **_Note: It is advisable to [Install Meshplay in your Minikube clusters](#install-meshplay-into-your-minikube-cluster)_**

<div class="prereqs"><p><strong style="font-size: 20px;">Prerequisites</strong> </p> 
  <ol>
    <li>Install the Meshplay command line client, <a href="{{ site.baseurl }}/installation/meshplayctl" class="meshplay-light">meshplayctl</a>.</li>
    <li>Install <a href="https://kubernetes.io/docs/tasks/tools/">kubectl</a> installed on your local machine.</li>
  </ol>
</div>

Also see: [Install Meshplay on Kubernetes]({{ site.baseurl }}/installation/kubernetes)

## Available Deployment Methods

- [In-cluster Installation](#in-cluster-installation)
  - [Preflight Checks](#preflight-checks)
    - [Preflight: Cluster Connectivity](#preflight-cluster-connectivity)
    - [Preflight: Plan your access to Meshplay UI](#preflight-plan-your-access-to-meshplay-ui)
  - [Installation: Using `meshplayctl`](#installation-using-meshplayctl)
  - [Installation: Using Helm](#installation-using-helm)
  - [Installation: Manual Steps](#installation-manual-steps)
  - [Installation: Docker Driver Users](#installation-docker-driver-users)
- [Out-of-cluster Installation](#out-of-cluster-installation)
  - [Installation: Install Meshplay on Docker](#installation-install-meshplay-on-docker)
  - [Installation: Upload Config File in Meshplay Web UI](#installation-upload-config-file-in-meshplay-web-ui)
  - [Post-Installation Steps](#post-installation-steps)

# In-cluster Installation

Follow the steps below to install Meshplay in your Minikube cluster.

## Preflight Checks

Read through the following considerations prior to deploying Meshplay on Minikube.

### Preflight: Cluster Connectivity

Start the minikube, if not started using the following command:
{% capture code_content %}minikube start --cpus 4 --memory 8192 --kubernetes-version=v1.14.1{% endcapture %}
{% include code.html code=code_content %}
Check up on your minikube cluster :
{% capture code_content %}minikube status{% endcapture %}
{% include code.html code=code_content %}
Verify your kubeconfig's current context.
{% capture code_content %}kubectl config current-context{% endcapture %}
{% include code.html code=code_content %}

### Preflight: Plan your access to Meshplay UI

1. If you are using port-forwarding, please refer to the [port-forwarding]({{ site.baseurl }}/reference/meshplayctl/system/dashboard) guide for detailed instructions.
2. Customize your Meshplay Provider Callback URL. Meshplay Server supports customizing authentication flow callback URL, which can be configured in the following way:

{% capture code_content %}$ MESHPLAY_SERVER_CALLBACK_URL=https://custom-host meshplayctl system start{% endcapture %}
{% include code.html code=code_content %}

Meshplay should now be running in your Minikube cluster and Meshplay UI should be accessible at the `INTERNAL IP` of `meshplay` service.

## Installation: Using `meshplayctl`

Use Meshplay's CLI to streamline your connection to your Minikube cluster. Configure Meshplay to connect to your Minikube cluster by executing:

{% capture code_content %}$ meshplayctl system config minikube{% endcapture %}
{% include code.html code=code_content %}

Once configured, execute the following command to start Meshplay.

{% capture code_content %}$ meshplayctl system start{% endcapture %}
{% include code.html code=code_content %}

If you encounter any authentication issues, you can use `meshplayctl system login`. For more information, click [here](/guides/meshplayctl/authenticate-with-meshplay-via-cli) to learn more.

## Installation: Using Helm

For detailed instructions on installing Meshplay using Helm V3, please refer to the [Helm Installation](/installation/kubernetes/helm) guide.

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

## Installation: Docker Driver Users

Follow the [installation steps](/installation/quick-start) to setup the meshplayctl CLI and install Meshplay.

**Users using docker driver**:
After completing the Meshplay installation, execute the following commands to establish connectivity between Meshplay Server and Kubernetes cluster:

 <pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">docker network connect bridge meshplay_meshplay_1</div></div>
 </pre>

<br/>

<pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">docker network connect minikube meshplay_meshplay_1</div></div>
 </pre>

To establish connectivity between a particular Meshplay Adapter and Kubernetes server, use _"docker ps"_ to identify the name of the desired container, and execute the following commands:

<pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">docker network connect bridge &#60; container name of the desired adapter &#62;</div></div>
 </pre>

<br/>

 <pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">docker network connect minikube &#60; container name of the desired adapter &#62;</div></div>
 </pre>

# Out-of-cluster Installation

Install Meshplay on Docker (out-of-cluster) and connect it to your Minikube cluster.

## Installation: Install Meshplay on Docker

{% capture code_content %}$ meshplayctl system start -p docker{% endcapture %}
{% include code.html code=code_content %}

Configure Meshplay to connect to your cluster by executing:

{% capture code_content %}$ meshplayctl system config minikube{% endcapture %}
{% include code.html code=code_content %}

Once you have verified that all the services are up and running, Meshplay UI will be accessible on your local machine on port 9081. Open your browser and access Meshplay at [`http://localhost:9081`](http://localhost:9081).

## Installation: Upload Config File in Meshplay Web UI

- Run the below command to generate the _"config_minikube.yaml"_ file for your cluster:

 <pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">kubectl config view --minify --flatten > config_minikube.yaml</div></div>
 </pre>

- Upload the generated config file by navigating to _Settings > Environment > Out of Cluster Deployment_ in the Web UI and using the _"Upload kubeconfig"_ option.

## Post-Installation Steps

Optionally, you can verify the health of your Meshplay deployment, using <a href='/reference/meshplayctl/system/check'>meshplayctl system check</a>.

You're ready to use Meshplay! Open your browser and navigate to the Meshplay UI.

{% include_cached installation/accessing-meshplay-ui.md %}

{% include related-discussions.html tag="meshplay" %}

