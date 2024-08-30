<p style="text-align:center;" align="center"><a href="https://khulnasoft.com/meshplay"><picture align="center">
  <source media="(prefers-color-scheme: dark)" srcset="https://github.com/meshplay/meshplay/blob/master/.github/assets/images/meshplay/meshplay-logo-dark-text-side.svg"  width="70%" align="center" style="margin-bottom:20px;">
  <source media="(prefers-color-scheme: light)" srcset="https://github.com/meshplay/meshplay/blob/master/.github/assets/images/meshplay/meshplay-logo-light-text-side.svg" width="70%" align="center" style="margin-bottom:20px;">
  <img alt="Shows an illustrated light mode meshplay logo in light color mode and a dark mode meshplay logo dark color mode." src="https://raw.githubusercontent.com/meshplay/meshplay/master/.github/assets/images/meshplay/meshplay-logo-tag-light-text-side.png" width="70%" align="center" style="margin-bottom:20px;">
</picture></a><br /><br /></p>

<p align="center">
<a href="https://hub.docker.com/r/meshplay/meshplay" alt="Docker pulls">
  <img src="https://img.shields.io/docker/pulls/meshplay/meshplay.svg" /></a>
<a href="https://github.com/issues?utf8=✓&q=is%3Aopen+is%3Aissue+archived%3Afalse+org%3Akhulnasoft+label%3A%22help+wanted%22+" alt="GitHub issues by-label">
  <img src="https://img.shields.io/github/issues/khulnasoft/meshplay/help%20wanted.svg?color=informational" /></a>
<a href="https://github.com/meshplay/meshplay/blob/master/LICENSE" alt="LICENSE">
  <img src="https://img.shields.io/github/license/meshplay/meshplay?color=brightgreen" /></a>
<a href="https://goreportcard.com/report/github.com/meshplay/meshplay" alt="Go Report Card">
  <img src="https://goreportcard.com/badge/github.com/meshplay/meshplay" /></a>
<a href="https://github.com/meshplay/meshplay/actions" alt="Build Status">
  <img src="https://img.shields.io/github/workflow/status/meshplay/meshplay/Meshplay%20Build%20and%20Releaser%20(edge)" /></a>
<a href="https://bestpractices.coreinfrastructure.org/projects/3564" alt="CLI Best Practices">
  <img src="https://bestpractices.coreinfrastructure.org/projects/3564/badge" /></a>
<a href="http://discuss.meshplay.khulnasoft.com" alt="Discuss Users">
  <img src="https://img.shields.io/discourse/users?label=discuss&logo=discourse&server=https%3A%2F%2Fdiscuss.khulnasoft.com" /></a>
<a href="https://slack.meshplay.khulnasoft.com" alt="Join Slack">
  <img src="https://img.shields.io/badge/Slack-@khulnasoft.svg?logo=slack"></a>
<a href="https://twitter.com/intent/follow?screen_name=meshplayio" alt="Twitter Follow">
  <img src="https://img.shields.io/twitter/follow/meshplayio.svg?label=Follow+Meshplay&style=social" /></a>
</p>

[Meshplay](https://meshplay.khulnasoft.com) is the cloud native management plane offering lifecycle, configuration, and performance management of Kubernetes, service meshes, and your workloads.

<p align="center">
Meshplay is a Cloud Native Computing Foundation project.
</p>

# Docker Extension for Meshplay

The Docker Extension for Meshplay extends Docker Desktop’s position as the cloud native developer’s go-to Kubernetes environment with easy access to the next layer of cloud native infrastructure: service meshes. Service mesh or not, though, Meshplay offers a visual topology for designing Docker-Compose applications, operating Kubernetes, service meshes, and their workloads. Meshplay brings deep support of 10 different service meshes to the fingertips of Docker Desktop developers in connection with Docker Desktop’s ability to deliver Kubernetes locally.

<h2><a name="running"></a>Get Started with the Docker Extension for Meshplay</h2>

<h3>Using Docker Desktop</h3>
<p>Navigate to the Extensions area of Docker Desktop.</p>

<h3>Using <code>docker</code></h3>
<p>Meshplay runs as a set of containers inside your Docker Desktop virtual machine.</p>
<pre>docker extension install meshplay/docker-extension-meshplay</pre>
<p>See the <a href="https://docs.meshplay.khulnasoft.com/installation/quick-start">quick start</a> guide.</p>
<p style="clear:both;">&nbsp;</p>

## Using the Docker Extension for Meshplay

1. Install any service mesh with the check of a box.
1. Import your Docker Compose apps for visual design and deployment to Kubernetes and service meshes.

<p align="center"><a href="https://raw.githubusercontent.com/meshplay/meshplay/master/install/docker-extension/docs/img/docker-desktop-extension-for-meshplay.png"><img src="https://raw.githubusercontent.com/meshplay/meshplay/master/install/docker-extension/docs/img/docker-desktop-extension-for-meshplay.png" width="90%" align="center" /></a></p>

## Docker Extension for Meshplay Architecture

The Docker Extension for Meshplay deploys Meshplay to your local Docker host as a Docker Compose application.

<p align="center"><a href="https://raw.githubusercontent.com/meshplay/meshplay/master/install/docker-extension/docs/img/docker-extension-for-meshplay-architecture.png"><img src="https://raw.githubusercontent.com/meshplay/meshplay/master/install/docker-extension/docs/img/docker-extension-for-meshplay-architecture.png" width="90%" align="center" /></a></p>
Learn more about <a href="https://docs.meshplay.khulnasoft.com/architecture">Meshplay's architecture</a>.

## Docker Extension for Meshplay

From `/install/docker-extension`, familiarize with available `make` targets by executing:

```
make
```
Review the available targets and their purpose. In general, follow this sequence when building and testing changes:

```
make extension
```

Once build is complete:

```
docker extension install meshplay/docker-extension-meshplay:edge-latest
```

Or reinstall with:

```
docker extension update meshplay/docker-extension-meshplay:edge-latest
```

<p style="text-align:center; width:100%;" align="center">
<a href ="https://khulnasoft.com/community"><img alt="MeshMates" src="https://docs.meshplay.khulnasoft.com/assets/img/readme/khulnasoft-community-sign.png" style="margin-right:10px; margin-bottom:7px;" width="28%" align="center" /></a>
</p>
<p style="text-align:center; width:100%;" align="center">
<h3 style="text-align:center;" align="center"><em>Have questions? Need help?</em> <strong>Ask in the <a href="http://discuss.meshplay.khulnasoft.com">Community Forum</a></strong>.</h3></p>

