---
layout: default
title: GKE
permalink: installation/kubernetes/gke
type: installation
category: kubernetes
redirect_from:
- installation/platforms/gke
display-title: "false"
language: en
list: include
image: /assets/img/platforms/gke.png
abstract: Install Meshplay on Google Kubernetes Engine. Deploy Meshplay in GKE in-cluster or outside of GKE out-of-cluster.
---

<h1>Quick Start with {{ page.title }} <img src="{{ page.image }}" style="width:35px;height:35px;" /></h1>

Manage your GKE clusters with Meshplay. Deploy Meshplay in GKE [in-cluster](#in-cluster-installation) or outside of GKE [out-of-cluster](#out-of-cluster-installation). **_Note: It is advisable to [Install Meshplay in your GKE clusters](#in-cluster-installation)_**

<div class="prereqs"><p><strong style="font-size: 20px;">Prerequisites</strong> </p> 
  <ol>
    <li>Install the Meshplay command line client, <a href="{{ site.baseurl }}/installation/meshplayctl" class="meshplay-light">meshplayctl</a>.</li>
    <li>Install <a href="https://kubernetes.io/docs/tasks/tools/">kubectl</a> on your local machine.</li>
    <li>Install <a href="https://cloud.google.com/sdk/docs/install">gCloud CLI</a>, configured for your environment.</li>
    <li>Access to an active GKE cluster in your Google Cloud project.</li>
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
  - [Post-Installation Steps](#post-installation-steps)

# In-cluster Installation

Follow the steps below to install Meshplay in your GKE cluster.

## Preflight Checks

Read through the following considerations prior to deploying Meshplay on GKE.

### Preflight: Cluster Connectivity

1. Verfiy you connection to an Google Kubernetes Engine Cluster using gCloud CLI.
1. Login to GCP account using [gcloud auth login](https://cloud.google.com/sdk/gcloud/reference/auth/login).
1. After a successful login, set the Project Id:
{% capture code_content %}gcloud config set project [PROJECT_ID]
{% endcapture %}
{% include code.html code=code_content %}
1. After setting the Project Id, set the cluster context.
{% capture code_content %}gcloud container clusters get-credentials [CLUSTER_NAME] --zone [CLUSTER_ZONE] {% endcapture %}
{% include code.html code=code_content %}
1. Verify your kubeconfig's current context.
{% capture code_content %}kubectl config current-context{% endcapture %}
{% include code.html code=code_content %}

### Preflight: Plan your access to Meshplay UI

1. If you are using port-forwarding, please refer to the [port-forwarding]({{ site.baseurl }}/reference/meshplayctl/system/dashboard) guide for detailed instructions.
2. If you are using a LoadBalancer, please refer to the [LoadBalancer]({{ site.baseurl }}/installation/kubernetes#exposing-meshplay-serviceloadbalancer) guide for detailed instructions.
3. Customize your Meshplay Provider Callback URL. Meshplay Server supports customizing authentication flow callback URL, which can be configured in the following way:

{% capture code_content %}$ MESHPLAY_SERVER_CALLBACK_URL=https://custom-host meshplayctl system start{% endcapture %}
{% include code.html code=code_content %}

Meshplay should now be running in your GKE cluster and Meshplay UI should be accessible at the `EXTERNAL IP` of `meshplay` service.

## Installation: Using `meshplayctl`

Use Meshplay's CLI to streamline your connection to your GKE cluster. Configure Meshplay to connect to your GKE cluster by executing:

{% capture code_content %}$ meshplayctl system config gke{% endcapture %}
{% include code.html code=code_content %}

Once configured, execute the following command to start Meshplay.

{% capture code_content %}$ meshplayctl system start{% endcapture %}
{% include code.html code=code_content %}

If you encounter any authentication issues, you can use `meshplayctl system login`. For more information, click [here](/guides/meshplayctl/authenticate-with-meshplay-via-cli) to learn more.

## Installation: Using Helm

For detailed instructions on installing Meshplay using Helm V3, please refer to the [Helm Installation](/installation/helm) guide.

## Post-Installation Steps

Optionally, you can verify the health of your Meshplay deployment, using <a href='/reference/meshplayctl/system/check'>meshplayctl system check</a>.

You're ready to use Meshplay! Open your browser and navigate to the Meshplay UI.

{% include_cached installation/accessing-meshplay-ui.md display-title="true" %}

{% include related-discussions.html tag="meshplay" %}