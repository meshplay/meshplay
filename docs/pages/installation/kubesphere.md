---
layout: default
title: KubeSphere
permalink: installation/kubernetes/kubesphere
type: installation
category: kubernetes
redirect_from:
- installation/platforms/kubershphere
display-title: "false"
language: en
list: include
image: /assets/img/platforms/kubesphere.png
abstract: Install Meshplay on KubeSphere
---

{% include installation/installation_prerequisites.html %}

[Meshplay](https://meshplay.khulnasofy.com/) is the open source, cloud native management plane that enables the adoption, operation, and management of Kubernetes, any service mesh, and their workloads.

This tutorial walks you through an example of deploying Meshplay from the App Store of KubeSphere.


## Prerequisites

- Please make sure you enable the OpenPitrix system.
- You need to create a workspace, a project, and a user account (`project-regular`) for this tutorial. The account needs to be a platform regular user and to be invited as the project operator with the `operator` role. In this tutorial, you log in as `project-regular` and work in the project `demo-project` in the workspace `demo-workspace`. For more information, see Create Workspaces, Projects, Users and Roles.


## Hands-on Lab

Perform the following steps in order:

### 1. <b>Deploy Meshplay from the App Store</b>


1. On the **Overview** page of the project `demo-project`, click **App Store** in the upper-left corner.
2. Search for **Meshplay** in the App Store, and click on the search result to enter the app.

    ![meshplay-app]({{ site.baseurl }}/assets/img/platforms/meshplay-app.png)
3. In the **App Information** page, click **Install** on the upper right corner.

    ![meshplay-install]({{ site.baseurl }}/assets/img/platforms/Meshplay-install.png)

4. In the App Settings page, set the application **Name**, **Location** (as your Namespace), and App Version, and then click Next on the upper right corner.

    ![meshplay-info]({{ site.baseurl }}/assets/img/platforms/Meshplay-info.png)

5. Configure the **values.yaml** file as needed, or click **Install** to use the default configuration.

    ![meshplay-yaml]({{ site.baseurl }}/assets/img/platforms/Meshplay-yaml.png)

6. Wait for the deployment to be finished. Upon completion, **Meshplay** will be shown as **Running** in KubeSphere.

    ![meshplay-app-running]({{ site.baseurl }}/assets/img/platforms/Meshplay-app-running.png)



### 2. <b>Access the Meshplay Dashboard</b>


1. Go to **Services** and click the service name of Meshplay.
2. In the **Resource Status** page, copy the **NodePort** of Meshplay.

    ![meshplay-service]({{ site.baseurl }}/assets/img/platforms/Meshplay-service.png)

3. Access the Meshplay Dashboard by entering **${NodeIP}:${NODEPORT}** in your browser.

    ![meshplay-dashboard]({{ site.baseurl }}/assets/img/platforms/meshplay-dashboard.png)

{% include related-discussions.html tag="meshplay" %}
