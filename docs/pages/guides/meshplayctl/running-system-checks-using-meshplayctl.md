---
layout: default
title: Running system checks using Meshplay CLI
permalink: guides/meshplayctl/running-system-checks-using-meshplayctl
language: en
type: guides
category: meshplayctl
list: include
abstract: Run pre-flight and post-deployment system health checks using Meshplay's CLI, meshplayctl.
---

Meshplay's CLI, `meshplayctl`, includes commands for verifying system readiness for a Meshplay deployment and health checks to confirm the health of an existing Meshplay deployment. Whether you have yet to deploy Meshplay or have already deployed Meshplay, `meshplayctl system check` is a useful utility to ensure that your Meshplay deployments are healthy.

<pre class="codeblock-pre">
<div class="codeblock"><div class="clipboardjs">meshplayctl system check</div></div>
</pre>
<br/>
<pre class="codeblock-pre">
<div class="codeblock"><div class="clipboardjs">Verify environment pre/post-deployment of Meshplay.

Usage:
meshplayctl system check [flags]

Flags:
--adapter Check status of Meshplay adapters
-h, --help help for check
--operator Check status of Meshplay operators
--pre Verify environment readiness to deploy Meshplay
--preflight Verify environment readiness to deploy Meshplay

Global Flags:
--config string path to config file (default "/Users/navendu/.meshplay/config.yaml")
-c, --context string (optional) temporarily change the current context.
-v, --verbose verbose output
-y, --yes (optional) assume yes for user interactive prompts.</div></div>
</pre>

## Deployment checks

`meshplayctl system check` command can run two types of system checks. A pre-deployment check which verifies the environment to deploy Meshplay and a post-deployment check which runs validation checks on a running Meshplay deployment.

### Pre-deployment checks

Pre-deployment checks runs checks on the environment and verifies whether it is ready for deploying Meshplay.

The following checks are done here:

- Docker health checks: Checks for the availability of Docker and docker-compose in the user's machine
- Kubernetes health checks: Checks for the availability of a Kubernetes cluster and checks if Meshplay can initialize a Kubernetes client
- Kubernetes version checks: Checks if kubectl and the Kubernetes version are higher than the minimum supported versions

Pre-deployment checks are run with the `--preflight` flag as shown below:

<pre class="codeblock-pre">
<div class="codeblock"><div class="clipboardjs">meshplayctl system check --preflight</div></div>
</pre>

### Post-deployment checks

Post-deployment checks are run after deploying Meshplay in the user's environment. These checks ensure that the running deployment of Meshplay and Meshplay adapters are working as expected.

In addition to the pre-flight checks, the following checks are also run in this check:

- Meshplay version checks: Checks the version of Meshplay server and CLI and shows if a new version is available
- Meshplay Adapter health checks: Checks if all the specified adapters are deployed and reachable

Post-deployment checks are run as shown below:

<pre class="codeblock-pre">
<div class="codeblock"><div class="clipboardjs">meshplayctl system check</div></div>
</pre>

## Additional checks

To check the status of the deployed adapters only, users can leverage the `--adapter` flag as shown below:

<pre class="codeblock-pre">
<div class="codeblock"><div class="clipboardjs">meshplayctl system check --adapter</div></div>
</pre>

Users can also narrow down the tests to just check the status of the Meshplay operator deployed on their Kubernetes cluster:

<pre class="codeblock-pre">
<div class="codeblock"><div class="clipboardjs">meshplayctl system check --operator</div></div>
</pre>

## FAQ

##### Question: While running `meshplayctl system check --preflight` it says I didn't install Kubernetes, but I have Docker installed and the test returned "Meshplay prerequisites met". Is that all good?

**Answer**: _Yes, as long as you've Docker installed, it's fine to run Meshplay. But you will need a Kubernetes cluster to handle tasks such as deploying infrastructure and so on, if you want to do them via Meshplay._

##### Question: I ran a preflight check to see if I satisfy all requirements for Meshplay in my system. It returned postive results but I couldn't start Meshplay. What to do?

**Answer**: _Make sure if you've configured your system to run Meshplay in smooth manner. For configuration, do check out the docs site and [this page]({{ site.baseurl }}/installation) to see instructions related to the platform you use._

##### Question: Do I need a Kubernetes cluster or will a Docker host suffice for Meshplay deployments?

**Answer**: _Meshplay's [performance management](tasks/performance/managing-performance) functionality does not require a Kubernetes cluster. The rest of Meshplay's functionality (e.g. cloud native management) does require a Kubernetes cluster._

##### Question: What are Meshplay's production deployment requirements?

**Answer**: _One or more Kubernetes clusters. A stateful set for Meshplay Server in order to persist performance test results. See [#2451](https://github.com/meshplay/meshplay/issues/2451)._

##### Question: For system checks, do I need any add-ons to pass the check?

**Answer**: _Not necessary. Basic requirements are enough to pass the check._

##### Question: The Adapter check is failing, it returns "Auth token not found".

**Answer**: _You can log in to Meshplay using `meshplayctl system login` which would generate an OAuth token. Once the OAuth token is generated, the check will start to function_

##### Question: I have a Kubernetes cluster enabled but Meshplay couldn't reach the cluster and the checks are failing! What to do?

**Answer**: _To resolve this error, you can upload your kubeconfig file in the Meshplay UI under settings and Meshplay will reconfigure to use your Kubernetes cluster._

##### Question: Under Meshplay Version test, I'm getting an error like "CLI is not up-to-date". Should I update meshplayctl often?

**Answer**: _Yes! You should update the meshplayctl often in order to run Meshplay smoothly. The reason behind it is because not only the CLI is updated, but also the Meshplay app. So it is advisable to update Meshplay often._

##### Question: Is it advisable to keep Meshplay in sleep mode while running system checks?

**Answer**: _Not necessary. It is good to keep Meshplay up and running, else the system checks will fail to detect the Meshplay version._

##### Question: What is the minimum version of k8s cluster and kubectl required to run Meshplay?

**Answer**: _For Kubernetes, version >=1.12.0 is recommended. For kubectl version >=1.12 is recommended._

##### Question: In the "Meshplay Adapter" section of check, I could see only some Meshplay adapters up and running and not all. Is this fine?

**Answer**: _Not a problem, if you feel you need to have all mesh adapters to be up running, you can do so by creating a new context `meshplayctl system context create [context-name] --set` (if you voluntarily deleted mesh adapters in your current context)_

##### Question: I started Meshplay fresh, didn't change any of the details in the context I have. But I see that all adapter checks are failing. What to do?

**Answer**: _Configure Meshplay to use on your Kubernetes cluster, then upload the kubeconfig file via Meshplay UI to notify Meshplay to use that cluster. If that didn't work, feel free to [open up an issue](https://github.com/meshplay/meshplay/issues) in GitHub._

### Suggested Reading

For an exhaustive list of `meshplayctl` commands and syntax:

- See [`meshplayctl` Command Reference]({{ site.baseurl }}/reference/meshplayctl).

Guides to using Meshplay's various features and components.

{% capture tag %}

<li><a href="{{ site.baseurl }}/installation/upgrades#upgrading-meshplay-cli">Upgrading meshplayctl</a></li>

{% endcapture %}

{% include related-discussions.html tag="meshplayctl" %}

