---
layout: default
title: Upgrading Meshplay and all of its components
abstract: How to Meshplay and all of its components
permalink: installation/upgrades
display-title: "false"
type: guides
category: installation
language: en
abstract: How to upgrade Meshplay and all of its components
---

# Upgrade Guide

## Upgrading Meshplay Server, Adapters, and UI

Various components of Meshplay will need to be upgraded as new releases become available. Meshplay is comprised of a number of components including a server, adapters, UI, and CLI. As an application, Meshplay is a composition of different functional components.

<p style="text-align:center">
<a href="{{site.baseurl}}/assets/img/architecture/upgrading-meshplay.svg">
    <img src="{{site.baseurl}}/assets/img/architecture/upgrading-meshplay.svg" style="margin: 1rem;" />
</a><br /><i><small>Figure: Meshplay components</small></i>
</p>

Some of the components must be upgraded simultaneously, while others may be upgraded independently. The following table depicts components, their versions, and deployment units (deployment groups).

### Versioning of Meshplay components

<table class="meshplaycomponents">
    <tr>
        <th>Components</th>
        <th>Sub-component</th>
        <th>Considering or Updating</th>
    </tr>
    <tr>
        <td class="childcomponent">Meshplay Adapters</td>
        <td>Any and All Adapters</td>
        <td>Docker Deployment: Watchtower updates this component in accordance with the user’s release channel subscription.</td>
    </tr>
    <tr>
        <td rowspan="3" class="childcomponent">Meshplay Server</td>
        <td>Meshplay UI</td>
        <td rowspan="3">Manages lifecycle of Meshplay Operator; Adapters, UI, Load Generators, Database.<br /><br />
Docker Deployment: Watchtower updates this component in accordance with the user’s release channel subscription.</td>
    </tr>
    <tr>
        <td>Load Generators</td>
    </tr>
    <tr>
        <td>Database</td>
    </tr>
    <tr>
        <td rowspan="2" class="childcomponent">Meshplay Operator</td>
        <td>MeshSync</td>
        <td>Meshplay Operator manages the lifecycle of this component and its sub-components.</td>
    </tr>
    <tr>
        <td>Meshplay Broker</td>
        <td>Meshplay Operator manages the lifecycle of this event bus component.</td>
    </tr>
    <tr>
        <td class="childcomponent">`meshplayctl`</td>
        <td></td>
        <td><code>meshplayctl</code> manages the lifecycle of Meshplay Server. <br /><br />
        <ul> 
            <li><code>system start</code> calls system update by default, which updates server and existing adapters, but doesn’t update meshplay.yaml. Unless the <code>skipUpdate</code> flag is used, operators are also updated here.</li>
            <li><code>system reset</code> retrieving docker-compose.yaml from GitHub (use git tag to reset to the right Meshplay version).</li>
            <li><code>system restart</code> also updates operators, unless the <code>skipUpdate</code> flag is used.</li>
            <li><code>system update</code> updates operators in case of both docker and kubernetes deployments.</li>
            <li><code>system context</code> manages config.yaml, which manages meshplay.yaml. </li>
            <li><code>meshplayctl</code> should generally be checking for latest release and informing user.</li>
        </ul>
        </td>
    </tr>
    <tr>
        <td rowspan="2" class="childcomponent"><a style="color:white;" ref="/extensibility/providers">Remote Providers</a></td>
        <td>Meshplay Cloud</td>
        <td>Process Extension: Integrators manage the lifecycle of their Remote Providers. Process is unique per provider.</td>
    </tr>
    <tr>
        <td>Meshplay Cloud</td>
        <td> Static Extension: Integrators manage the lifecycle of their Meshplay Extensions. Process is unique per provider.</td>
    </tr>
</table>

Sub-components deploy as a unit, however, they do not share the same version number.

### Meshplay Docker Deployments

In order to pull the latest images for Meshplay Server, Adapters, and UI, execute the following command:

 <pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">meshplayctl system update</div></div>
 </pre>

If you wish to update a running Meshplay deployment with the images you just pulled, you'll also have to execute:

 <pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">meshplayctl system restart</div></div>
 </pre>

### Meshplay Kubernetes Deployments

Use `kubectl apply` or `helm` to upgrade the Meshplay application manifests in your Kubernetes cluster.

## Upgrading Meshplay CLI

The Meshplay command line client, `meshplayctl`, is available in different package managers. Use the instructions relevant to your environment.

### Upgrading `meshplayctl` using Homebrew

<p>To upgrade `meshplayctl`, execute the following command:</p>

 <pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">brew upgrade meshplayctl</div></div>
 </pre>

### Upgrading `meshplayctl` using Bash

Upgrade `meshplayctl` and run Meshplay on Mac or Linux with this script:

 <pre class="codeblock-pre">
 <div class="codeblock"><div class="clipboardjs">curl -L https://meshplay.io/install | DEPLOY_MESHPLAY=false bash -</div></div>
 </pre>

### Upgrading `meshplayctl` using Scoop

To upgrade `meshplayctl`, execute the following command:

 <pre class="codeblock-pre">
 <div class="codeblock"><div class="clipboardjs">scoop update meshplayctl</div></div>
 </pre>

{% include related-discussions.html tag="meshplay" %}
