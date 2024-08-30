---
layout: default
title: Troubleshooting Meshplay Installations
abstract: Troubleshoot Meshplay installation and deployment
permalink: guides/troubleshooting/installation
type: guides
category: troubleshooting
language: en
abstract: Troubleshoot Meshplay installation and deployment
---

## Meshplay's Preflight Checks

Anytime a `meshplayctl system` command is executed, a series of preflight checks are run. An attempt will be made to connect to the Kubernetes cluster configured in the user's kubeconfig as their current-context .

1. Check whether `meshplayctl` can initialize a Kubernetes client.

   Situation: `meshplayctl` fails to query for pods in the default namespace of the user's current Kubernetes context.

2. Remove `~/.meshplay` to reinitialize Meshplay

   Situation: Unable to start Meshplay Server with `make run-local` due to error of `key/value size is invalid`

## Setting up Meshplay using Kind or Minikube

The difficulty with Minikube and Kind clusters is that they typically don't support LoadBalancer service networking by default. Meshplay UI and Meshplay Broker are configured for LoadBalancer service networking by default. There are a number of solutions this overcoming this challenge. Here are a few methods:

1. Use the MetalLB Minikube add-on that provides load balancing. `minikube addons enable metallb`

   MetalLB setup: [link](https://kubebyexample.com/learning-paths/metallb/install)

2. Use Minikube tunnel to expose services. `minikube tunnel`.

   Docs: [link](https://minikube.sigs.k8s.io/docs/handbook/accessing/#using-minikube-tunnel)

   A simpler way to resolve this issue can be `port-forwarding`. Run the following command in terminal:

   `kubectl port-forward service/meshplay 9081:9081 -n meshplay`

   {% include meshplayctl/system-dashboard.md %}

3. For `kind`, you can prefer installing MetalLB with a custom configmap.

   Docs: [link](https://kind.sigs.k8s.io/docs/user/loadbalancer/)

## Meshplay Operator

By default, Meshplay Operator is installed in all the connected clusters automatically once Meshplay server detects those clusters. The operator can manually be turned off on particular cluster from the settings page.

### Disabling the operator

The env variable DISABLE_OPERATOR=true can be used to signal Meshplay server to not install operator in any of the clusters at any point in time after starting. While using Meshplay server locally, the `make server-without-operator` should be used to start Meshplay in disabled operator mode.

### Meshplay Broker

Example of a healthy Meshplay Broker server with an actively connected (subscribed) Meshplay Server:

```
➜  ~ kubectl logs -n meshplay meshplay-broker-0 nats
[8] 2021/09/08 21:46:03.070952 [INF] Starting nats-server version 2.1.9
[8] 2021/09/08 21:46:03.070982 [INF] Git commit [7c76626]
[8] 2021/09/08 21:46:03.071308 [INF] Starting http monitor on 0.0.0.0:8222
[8] 2021/09/08 21:46:03.071370 [INF] Listening for client connections on 0.0.0.0:4222
[8] 2021/09/08 21:46:03.071512 [INF] Server id is NAAYJNX4LDDNXW5UE7IP7PRQR2W2JP546XSFNUWQQHN7JYY27RG47KSG
[8] 2021/09/08 21:46:03.071516 [INF] Server is ready
```

For details about the state of the Meshplay Server subscription see the http monitor port on Meshplay Broker.

### Meshplay Unable to Connect to Kubernetes

Meshplay is unable to detect the Kubernetes connection running on your local system, even after manually uploading the `.kube config` file

When deploying Meshplay out-of-cluster, verify your kubeconfig's contexts and the ability for Meshplay Server to reach Kubernetes cluster API from whatever host and network that Meshplay Server is being deployed on.

<pre class="codeblock-pre">
<div class="codeblock"><div class="clipboardjs">kubectl config get-contexts</div></div>
</pre>

If you're using Docker Destkop, consider whether you need to change your current Kubernetes context to `docker-desktop`.

<pre class="codeblock-pre">
<div class="codeblock"><div class="clipboardjs">kubectl config use-context
docker-desktop</div></div>
</pre>

## Meshplay Remote Providers

Once Meshplay is installed, the remote provider "Meshplay" can be chosen from UI or by using the command `meshplayctl system login`:

![Providers](/assets/img/providers/provider_screenshot.png)

```bash
➜  ~ meshplayctl system login
Use the arrow keys to navigate: ↓ ↑ → ←
? Select a Provider:
  ▸ Meshplay
    None
```

If you cannot see "Meshplay" Remote Provider and find such error logs in Meshplay Server's logs (`meshplayctl system logs`), please make sure that Meshplay Server is able to reach "https://meshplay.layer5.io" in order to initialize the "Meshplay" Remote Provider.

```bash
time="2021-11-10T11:05:30Z" level=error msg="[Initialize Provider]: Failed to get capabilities Get \"https://meshplay.layer5.io/v0.5.71/capabilities?os=meshplay\": dial tcp 3.140.89.205:443: i/o timeout"
```

For more details about Meshplay Providers:

- [Extensibility: Providers](/extensibility/providers)

## See Also

- [Meshplay Error Code Reference](/reference/error-codes)

