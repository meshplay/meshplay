---
layout: default
title: EKS
permalink: installation/kubernetes/eks
type: installation
category: kubernetes
redirect_from:
- installation/platforms/eks
display-title: "false"
language: en
list: include
image: /assets/img/platforms/eks.png
abstract: Install Meshplay on Elastic Kubernetes Service. Deploy Meshplay in EKS in-cluster or outside of EKS out-of-cluster.
---

<h1>Quick Start with {{ page.title }} <img src="{{ page.image }}" style="width:35px;height:35px;" /></h1>

Manage your EKS clusters with Meshplay. Deploy Meshplay in EKS [in-cluster](#in-cluster-installation) or outside of EKS [out-of-cluster](#out-of-cluster-installation). **_Note: It is advisable to [Install Meshplay in your EKS clusters](#install-meshplay-into-your-eks-cluster)_**

<div class="prereqs"><h4>Prerequisites</h4>
  <ol>
    <li>Install the Meshplay command line client, <a href="{{ site.baseurl }}/installation/meshplayctl" class="meshplay-light">meshplayctl</a>.</li>
    <li>Install <a href="https://kubernetes.io/docs/tasks/tools/">kubectl</a> on your local machine.</li>
    <li>Install <a href="https://docs.aws.amazon.com/eks/latest/userguide/getting-started.html">AWS CLI</a>, configured for your environment.</li>
    <li>Access to an active EKS cluster in AWS Account.</li>
  </ol>
</div>

Also see: [Install Meshplay on Kubernetes]({{ site.baseurl }}/installation/kubernetes)

### Available Deployment Methods

- [In-cluster Installation](#in-cluster-installation)
  - [Preflight Checks](#preflight-checks)
    - [Preflight: Cluster Connectivity](#preflight-cluster-connectivity)
  - [Installation: Using `meshplayctl`](#installation-using-meshplayctl)
  - [Installation: Using Helm](#installation-using-helm)
  - [Post-Installation Steps](#post-installation-steps)
- [Out-of-cluster Installation](#out-of-cluster-installation)
  - [Install Meshplay on Docker](#install-meshplay-on-docker)

# In-cluster Installation

Follow the steps below to install Meshplay in your EKS cluster.

**Prerequisites: Cluster Connectivity**

1. Verify your connection to an Elastic Kubernetes Services Cluster using AWS CLI.
1. Login to AWS account using [aws configure](https://docs.aws.amazon.com/cli/latest/userguide/cli-authentication-user.html), if you are using a different method of authentication in AWS, please refer to AWS documentation.
1. After successful login, set the cluster context.
{% capture code_content %}aws eks update-kubeconfig --name [YOUR_CLUSTER_NAME] --region [YOUR_REGION]{% endcapture %}
{% include code.html code=code_content %}
1. _Optional:_ If you are using `eksctl`, follow the [AWS documentation steps](https://docs.aws.amazon.com/eks/latest/userguide/getting-started-eksctl.html).
1. Verify your kubeconfig's current context.
{% capture code_content %}kubectl config current-context{% endcapture %}
{% include code.html code=code_content %}

## Installation: Using `meshplayctl`

Use Meshplay's CLI to streamline your connection to your EKS cluster. Configure Meshplay to connect to your EKS cluster by executing:

{% capture code_content %}$ meshplayctl system config eks{% endcapture %}
{% include code.html code=code_content %}

Once configured, execute the following command to start Meshplay.

{% capture code_content %}$ meshplayctl system start{% endcapture %}
{% include code.html code=code_content %}

## Installation: Using Helm

For detailed instructions on installing Meshplay using Helm V3, please refer to the [Helm Installation](/installation/kubernetes/helm) guide.

## Post-Installation Steps

Optionally, you can verify the health of your Meshplay deployment, using <a href='/reference/meshplayctl/system/check'>meshplayctl system check</a>.

You're ready to use Meshplay! Open your browser and navigate to the Meshplay UI.

{% include_cached installation/accessing-meshplay-ui.md display-title="true" %}

# Out-of-cluster Installation

{% include alert.html title='Out-of-cluster EKS deployments not currently supported' type="warning" alert='Out-of-cluster support for EKS is still beta and on <a href="https://github.com/meshplay/meshplay/blob/master/ROADMAP.md">roadmap</a>.' %}

Install Meshplay on Docker (out-of-cluster) and connect it to your EKS cluster.

## Install Meshplay on Docker

{% capture code_content %}$ meshplayctl system start -p docker{% endcapture %}
{% include code.html code=code_content %}

Configure Meshplay to connect to your cluster by executing:

{% capture code_content %}$ meshplayctl system config eks{% endcapture %}
{% include code.html code=code_content %}

Once you have verified that all the services are up and running, Meshplay UI will be accessible on your local machine on port 9081. Open your browser and access Meshplay at [`http://localhost:9081`](http://localhost:9081).

{% include related-discussions.html tag="meshplay" %}
