---
layout: default
title: meshplayctl-system-config
permalink: reference/meshplayctl/system/config
redirect_from: reference/meshplayctl/system/config/
type: reference
display-title: "false"
language: en
command: system
subcommand: config
---

# meshplayctl system config

Configure Meshplay

## Synopsis

Configure the Kubernetes cluster used by Meshplay.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system config [flags]

</div>
</pre> 

## Examples

Set configuration according to k8s cluster
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system config [aks|eks|gke|minikube]

</div>
</pre> 

Path to token for authenticating to Meshplay API (optional, can be done alternatively using "login")
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system config --token "~/Downloads/auth.json"

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help   help for config

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

## See Also

Go back to [command reference index](/reference/meshplayctl/), if you want to add content manually to the CLI documentation, please refer to the [instruction](/project/contributing/contributing-cli#preserving-manually-added-documentation) for guidance.
