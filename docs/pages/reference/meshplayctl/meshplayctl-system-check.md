---
layout: default
title: meshplayctl-system-check
permalink: reference/meshplayctl/system/check
redirect_from: reference/meshplayctl/system/check/
type: reference
display-title: "false"
language: en
command: system
subcommand: check
---

# meshplayctl system check

Meshplay environment check

## Synopsis

Verify environment pre/post-deployment of Meshplay.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system check [flags]

</div>
</pre> 

## Examples

Run system checks for both pre and post mesh deployment scenarios on Meshplay
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system check

</div>
</pre> 

Run Pre-mesh deployment checks (Docker and Kubernetes)
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system check --preflight

</div>
</pre> 

Run checks on specific mesh adapter
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system check --adapter meshplay-istio:10000

</div>
</pre> 

or
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system check --adapter meshplay-istio

</div>
</pre> 

Run checks for all the mesh adapters
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system check --adapters

</div>
</pre> 

Verify the health of Meshplay Operator's deployment with MeshSync and Broker
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system check --operator

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
      --adapter string   Check status of specified meshplay adapter
      --adapters         Check status of meshplay adapters
      --components       Check status of Meshplay components
  -h, --help             help for check
      --operator         Verify the health of Meshplay Operator's deployment with MeshSync and Broker
      --pre              Verify environment readiness to deploy Meshplay
      --preflight        Verify environment readiness to deploy Meshplay

</div>
</pre>

## Options inherited from parent commands

<pre class='codeblock-pre'>
<div class='codeblock'>
      --config string    path to config file (default "/home/runner/.meshplay/config.yaml")
  -c, --context string   (optional) temporarily change the current context.
  -v, --verbose          verbose output
  -y, --yes              (optional) assume yes for user interactive prompts.

</div>
</pre>

## Screenshots

Usage of meshplayctl system check
![check-usage](/assets/img/meshplayctl/check.png)

## See Also

Go back to [command reference index](/reference/meshplayctl/), if you want to add content manually to the CLI documentation, please refer to the [instruction](/project/contributing/contributing-cli#preserving-manually-added-documentation) for guidance.
