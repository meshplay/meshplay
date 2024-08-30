---
layout: default
title: Meshplay Operator, MeshSync, Broker Troubleshooting Guide
permalink: guides/troubleshooting/meshplay-operator-meshsync
language: en
abstract: This documentation provides comprehensive guidance on troubleshooting in Meshplay Operator, MeshSync and Broker, ensuring you can address common issues efficiently.
type: guides
category: troubleshooting
---

{% include alert.html type="dark" title="Meshplay Error Code Reference" content="Have specific error with an error code? See the <a href='/reference/error-codes'>Meshplay Error Code Reference</a> for probable cause and suggested remediations." %}

There are common issues Meshplay users may face while operating the [Meshplay Operator]({{site.baseurl}}/concepts/architecture/operator) and its custom controllers, [MeshSync]({{site.baseurl}}/concepts/architecture/meshsync) and [Broker]({{site.baseurl}}/concepts/architecture/broker), that can be resolved by performing specific actions. This documentation aims to empower users by providing a set of troubleshooting tools and actions.

## Understanding the Status of Meshplay Operator, MeshSync, and Meshplay Broker

The following table describes the various states of MeshSync and Meshplay Broker and their implications.

**MeshSync:**

- **ENABLED:** Custom Resource present. MeshSync Controller is not connected to Broker.
- **DEPLOYED:** Custom Resource present. MeshSync Controller is present but the state is not RUNNING or ERRDISABLE, though
- **RUNNING:** MeshSync pod present and in a running state.
- **CONNECTED:** Deployed and connected to Broker.
- **UNDEPLOYED:** Custom Resource not present.

**Meshplay Broker:**

- **DEPLOYED:** External IP not exposed OR External IP exposed but Meshplay Server is not connected as a client to Broker hence data is not being published.

- **UNDEPLOYED:** Custom Resource not deployed.
- **CONNECTED:** Deployed, sending data to Meshplay Server.

## Meshplay Operator Deployment Scenarios

Because Meshplay is versatile in its deployment models, there are a number of scenarios in which you may need to troubleshoot the health of the operator. The following sections describe the various scenarios and the steps you can take to troubleshoot them.

### In-Cluster Deployment

<!-- Meshplay Operator, MeshSync, and Broker are deployed in the same cluster as Meshplay Server. This is the default deployment scenario when using `meshplayctl system start` or `make run-local`. -->

Whether using [`meshplayctl system start`]({{site.baseurl}}/installation), [`helm install`]({{site.baseurl}}/installation/kubernetes/helm) or `make run-local`, Meshplay Server will automatically connect to any available Kubernetes clusters found in your kubeconfig (under `$HOME/.kube/config`). Once connected, operator, broker(NATS) and meshsync will automatically get deployed in the same clusters.

If everything is fine, by viewing the connection in Meshplay UI, MeshSync should be in **CONNECTED:** state. Otherwise, check the Operator's pod logs:

`kubectl logs <meshplay-operator-pod> -n meshplay`

### Out-of-Cluster Deployment

1. Meshplay Server is deployed on any Docker host (- Meshplay Server is deployed on a Docker host, and Meshplay Operator is deployed on a Kubernetes cluster).
   _or_
2. Meshplay is managing multiple clusters, some of which are not the cluster unto which Meshplay Server is deployed.

## Fault Scenarios

Common failure situations that Meshplay users might face are described below.

1. No deployment of Meshplay Operator, MeshSync, and Broker.
   1. Probable cause: Meshplay Server cannot connect to Kubernetes cluster; cluster unreachable or kubeconfig without proper permissions needed to deploy Meshplay Operator; Kubernetes config initialization issues.
1. Meshplay Operator with MeshSync and Broker deployed, but Meshplay Server is not receiving data from MeshSync or data the [Meshplay Database]({{site.baseurl}}/concepts/architecture/database) is stale.
   1. Probable cause: Meshplay Server lost subscription to Meshplay Broker; Broker server not exposed to external IP; MeshSync not connected to Broker; MeshSync not running; Meshplay Database is stale.
   2. The SQL database in Meshplay serves as a cache for cluster state. A single button allows users to dump/reset the Meshplay Database.
1. Orphaned MeshSync and Broker controllers - Meshplay Operator is not present, but MeshSync and Broker controllers are running.

## Operating Meshplay without Meshplay Operator

Meshplay Operator, MeshSync, and Broker are crucial components in a Meshplay deployment. Meshplay can function without them, but some functions of Meshplay will be disable / unusable. Whether Meshplay Operator is initially deployed via `meshplayctl` command or via Meshplay Server, you can monitor the health of the Meshplay Operator deployment using either the CLI or UI clients.

## Verifying the Status of Meshplay Operator, MeshSync, and Meshplay Broker

## Troubleshooting using Meshplay CLI

The following commands are available to troubleshoot Meshplay Operator, MeshSync, and Broker.

**Meshplay Server and Adapters**

- `meshplayctl system status` - Displays the status of Meshplay Server and Meshplay Adapters.

**Meshplay Operator, MeshSync, and Broker**

- `meshplayctl system check` - Displays the status of Meshplay Operator, MeshSync, and Broker.

## Troubleshooting using Meshplay UI

Based on discussed scenarios, the UI exposes tools to perform the following actions:

- (Re)deploy Operator, MeshSync, Broker.
- Uninstall and Install MeshSync, Broker, Operator.
- Reset Database.
- Ad hoc Connectivity Test for Operator, Meshplay Broker, MeshSync.
- Reconnect Meshplay Server to Meshplay Broker.
- Ad hoc Connectivity Test for Kubernetes context.
- Rediscover kubeconfig, delete, (re)upload kubeconfig.

### Synthetic Test for Ensuring Change in Cluster State

Initiate a synthetic check to verify a fully functional Operator deployment, testing MeshSync/Broker connectivity.

- Empty database shows the main-cluster node.
- Corrupt database triggers an error snackbar with a link to the Settings screen.
- Disconnected Kubernetes displays MeshSync logo pulsating when data is received.

<div class="section">
Future Enhancements for Troubleshooting:

- NATS/MeshSync not running prompts a review of available operations in the Settings panel.

</div>

This documentation provides comprehensive guidance on troubleshooting in Meshplay, ensuring users can address common issues efficiently.

{% include related-discussions.html tag="meshplay" %}

