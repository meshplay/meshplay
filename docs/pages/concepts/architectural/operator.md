---
layout: default
title: Operator
permalink: concepts/architecture/operator
type: components
redirect_from: architecture/operator
abstract: "Meshplay Operator controls and manages the lifecycle of components deployed inside a kubernetes cluster"
language: en
display-title: "false"
list: include
---

<link rel="stylesheet" type="text/css" href="{{ site.baseurl }}/_sass/operator.css">

# Meshplay Operator <img src="{{ site.baseurl }}/assets/img/architecture/B203EFA85E89491B.png" width="30" height="35" style="display:inline"/>

Meshplay Operator is a Kubernetes Operator that deploys and manages the lifecycle of two Meshplay components critical to Meshplay's operations of Kubernetes clusters. Deploy one Meshplay Operator per Kubernetes cluster under management - whether Meshplay Server is deploy inside or outside of the clusters under management. 

## Deployments

It is recommended to deploy one Meshplay Operator per cluster.

[![Meshplay Operator and MeshSync]({{ site.baseurl }}/assets/img/architecture/meshplay-operator-and-meshsync.svg
)]({{ site.baseurl }}/assets/img/architecture/meshplay-operator-and-meshsync.svg)

### Initialization Sequence

[![Meshplay Operator and MeshSync]({{ site.baseurl }}/assets/img/architecture/meshplay-operator-deployment-sequence.svg
)]({{ site.baseurl }}/assets/img/architecture/meshplay-operator-deployment-sequence.svg)

## Controllers managed by Meshplay Operator

### Broker Controller

Meshplay broker is one of the core components of the meshplay architecture. This controller manages the lifecycle of broker that meshplay uses for data streaming across the cluster and the outside world.

See [Meshplay Broker]({{site.baseurl}}/concepts/architecture/broker) for more information.

### MeshSync Controller

MeshSync Controller manages the lifecycle of MeshSync that is deployed for resource synchronization for the cluster.

See [MeshSync]({{site.baseurl}}/concepts/architecture/meshsync) for more information.

## Operator FAQs

### When is Meshplay Operator deployed and when is it deleted?  
As a Kubernetes custom controller, Meshplay Operator is provisioned and deprovisioned when Meshplay Server is connected to or disconnected from Kubernetes cluster. Meshplay Server connections to Kubernetes clusters are controlled using Meshplay Server clients: `meshplayctl` or Meshplay UI.  This behavior described below is consistent whether your Meshplay deployment is using Docker or Kubernetes as the platform to host the Meshplay deployment.

**Meshplay CLI**
`meshplayctl` initiates connection to Kubernetes cluster when `meshplayctl system start` is executed and disconnects when `meshplayctl system stop` is executed. This behavior is consistent whether your Meshplay deployment is using Docker or Kubernetes as the platform to host the Meshplay deployment.

**Meshplay UI**
Meshplay UI offers more granular control over the deployment of Meshplay Operator in that you can remove Meshplay Operator from a Kubernetes cluster without disconnecting Meshplay Server from the Kubernetes cluster. You can control the deployment of Meshplay Operator using the on/off switch found in the Meshplay Operator section of  Settings.

### Does the Meshplay Operator use an SDK or framework? 
Yes, Meshplay Operator used the Operator SDK.
