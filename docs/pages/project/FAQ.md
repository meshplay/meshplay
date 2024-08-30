---
layout: page
title: Frequently Asked Questions
permalink: project/faq
abstract: General commonly asked questions and answers about Meshplay.
language: en
type: project
category: project
---

## General FAQs

<details>
    <summary>
    <h6>Question: What is Meshplay?</h6>
</summary>

<p><strong>Answer:</strong> Meshplay is a self-service engineering platform that enables collaborative design and operation of cloud and cloud native infrastructure.</p>
</details>

<details>
    <summary>
    <h6>Question: Why was Meshplay created?</h6>
</summary>

<p><strong>Answer:</strong> As an open source, vendor neutral project, Meshplay was created out of the necessity to enable platform engineers, site reliability engineers, DevSecOps teams - all engineers to collaborate in the management of their infrastucture and workloads. Meshplay was created as an extensible platform to serve a broad set of modern application management needs.</p>
</details>

<details>
    <summary>
    <h6>Question: What does Meshplay do?</h6>
</summary>

<p><strong>Answer:</strong> Collaborative infrastructure management. Meshplay enables you to design and operate cloud native infrastructure visually, collaboratively, with confidence, and in partnership with your teammates.</p>
</details>

<!-- - _offers a catalog of operational best practices._
- _offersompare apples-to-apples performance across different infrastructure configurations._
- _Understand behavioral differences between service deployments._
- _Track your application performance from version to version._ -->

<details>
    <summary>
    <h6>Question: Is Meshplay an open source project?</h6>
</summary>
<p><strong>Answer:</strong> Yes, Meshplay is a Cloud Native Computing Foundation (CNCF) project and is licensed under Apache v2. As an internal developer platform, Meshplay is <a href="/extensibility">highly extensible</a>, offering multiple forms of extension points within which users and partners can customize and extend Meshplay's functionality.</p>
</details>

<details>
    <summary>
<h6>Question: Why should I use Meshplay?</h6>
</summary>
<p><strong>Answer:</strong> Meshplay is a powerful tool for managing ​Kubernetes infrastructure. It seamlessly integrates with different hundreds of tools and offers extensibility through many different <a href="{{site.baseurl}}/extensibility#extension-points">extension points</a>. With Meshplay, you can easily discover your environment, collaboratively manage multiple Kubernetes clusters, connect your Git and Helm repos, and analyze app and infra performance.</p>
</details>


## User FAQs

<details>
    <summary>
    <h6>Question: What is meshplayctl?</h6>
</summary>
<strong>Answer:</strong> A command line interface to manage Meshplay. `meshplayctl` can manage any number of Meshplay deployments.
</details>

<details>
<summary>
<h6>Question: How do I install Meshplay?</h6>
</summary>
<p><strong>Answer:</strong> Meshplay runs on a <a href="{{site.baseurl}}/installation">number of platforms</a>. You are encouraged to use <code>meshplayctl</code> to configure and control Meshplay deployments. Install `meshplayctl` using any of these options:</p>
<ul>
<li><a href="/installation/linux-mac/bash">Bash user</a></li>
<li><a href="/installation/linux-mac/brew">Brew user</a></li>
<li><a href="/installation/windows/scoop">Scoop user</a></li>
<li><a href="https://github.com/meshplay/meshplay/releases/latest">Direct download</a></li>
</ul>
</details>

<details>
<summary><h6>Question: What architecture does Meshplay have?</h6></summary>
<p><strong>Answer:</strong> An extensible architecture. There are several components, languages and they have different purposes. See Meshplay's <a href="/concepts/architecture">Architecture</a>.</p>
</details>

<details>
<summary>
<h6>Question: What is the difference between <code>make server</code> and <code>meshplayctl system start</code>? Do they both run Meshplay on my local machine?</h6>
</summary>
<strong>Answer:</strong> Yes, both of them do run Meshplay on your local machine. `make server` builds Meshplay from source and runs it on your local OS, while `meshplayctl system start` runs Meshplay as a set of containers in Docker or in Kubernetes on your local machine.
</details>

<details>
<summary>
<h6>Question: What systems can I deploy Meshplay onto?</h6>
</summary>
<strong>Answer:</strong> Many. See Meshplay's <a href="{{site.baseurl}}/installation">Compatibility Matrix</a>.
</details>

<details>
<summary><h6>Question: What systems does Meshplay manage?</h6></summary>
<p><strong>Answer:</strong> Many. See Meshplay's <a href="https://meshplay.khulnasofy.com/integrations">Integrations</a></p>
</details>

<details>
<summary><h6>Question: Why is Meshplay Server only receiving MeshSync updates from one of my Kubernetes Clusters?</h6></summary>
<p><strong>Answer:</strong> In order to receive MeshSync updates, Meshplay Server subscribes for updates Meshplay Broker. In other words, Meshplay Server connects to the `meshplay-broker` service port in order to subscribe for streaming MeshSync updates. By default, the Meshplay Broker service is deployed as type Kubernetes Service type <code>LoadBalancer</code>, which requires that your Kubernetes cluster provides an external IP address to the Meshplay Broker service, exposing it external to the Kubernetes cluster.</p>
<p>If you're running Kubernetes in Docker Desktop, an external IP address of <code>localhost</code> is assigned. If you're running Minikube, and execute <code>minikube tunnel</code> to gain access to Meshplay Broker's service, you will find that both Meshplay Broker service endpoints (from two different clusters) are sharing the same <code>localhost:4222</code> address and port number. This port sharing causes conflict and Meshplay Server is only able to connect to one of the Meshplay Brokers.</p>

<p>Few ways to solve this problem:</p>

<ul>
<li>Use an external cloud provider which provides you with the LoadBalancer having an external IP address other than localhost.</li>
<li>Use <a href="https://kind.sigs.k8s.io">Kind</a> cluster with <a href="https://metallb.universe.tf">MetalLB</a> configuration</li>
</ul>
</details>

<details><summary>
<h6>Question: Why does the dashboard not show the infrastructure provisioned or discovered by Meshplay?</h6></summary>
<strong>Answer:</strong> <p>This issue is typically caused by either lack of connectivity between Meshplay Server and Meshplay Broker or by database corruption. Use the following troubleshooting steps to resolve this issue:</p>

<p><strong>Lack of Connectivity</strong></p>

<ol>
<li>Confirm that the Meshplay Broker service is exposed from your cluster using <code>kubectl get svc -n meshplay</code> and that an hostname or IP address is displayed in the External Address column. Meshplay Server should be able to reach this address.</li>
<li>It is possible that MeshSync is not healthy and not sending cluster updates, check for MeshSync status by navigating to Settings in Meshplay UI and clicking on the MeshSync connection.</li>
<li>If MeshSync is healthy, check the status of Meshplay Broker by clicking on the NATS connection.</li>
</ol>

<p>If either is the case, Meshplay Operator will make sure MeshSync and Meshplay Broker deployments are again healthy, wait for some time, otherwise try redeploying Meshplay Operator.</p>

<p><strong>Database Corruption</strong></p>

<p>If MeshSync, Meshplay Broker and Meshplay Operator are healthy, then perhaps, there is corruption in the Meshplay Database. Use the following troubleshooting steps to resolve this issue:</p>
<ul>
<li>Try clearing the database by clicking on the `Flush MeshSync` button associated with the corresponding cluster.</li>
<li>If you don't see the specific entities in Meshplay UI, you may choose to reset Meshplay's database. This option is in the <code>Reset System</code> Tab in <code>Settings</code> page.</li>
</ul>

<p>Note: You can also verify health of your system using <a href="{{site.baseurl}}/reference/meshplayctl/system/check">meshplayctl system check</a></p>

</details>

## Contributing FAQs

<details>
<summary>
<strong>Question: Getting an error while running <code>make server</code> on Windows?</strong>
</summary><strong>Answer:</strong> <p>On Windows, set up the project on Ubuntu WSL2 and you will be able to run the Meshplay UI and the server. For more information please visit <a href="/project/contributing/meshplay-windows">Setting up Meshplay Development Environment on Windows</a>.</p>
</details>

{% include discuss.html %}

<!--Add other questions-->
