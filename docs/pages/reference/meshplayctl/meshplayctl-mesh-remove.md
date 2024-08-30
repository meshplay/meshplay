---
layout: default
title: meshplayctl-mesh-remove
permalink: reference/meshplayctl/mesh/remove
redirect_from: reference/meshplayctl/mesh/remove/
type: reference
display-title: "false"
language: en
command: mesh
subcommand: remove
---

# meshplayctl mesh remove

remove cloud and cloud native infrastructure

## Synopsis

remove cloud and cloud native infrastructure
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl mesh remove [flags]

</div>
</pre> 

## Examples

Remove Linkerd deployment
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl mesh remove linkerd

</div>
</pre> 

Remove a Linkerd control plane found under a specific namespace (linkerd-ns)
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl mesh remove linkerd --namespace linkerd-ns

</div>
</pre> 

<pre class='codeblock-pre'>
<div class='codeblock'>
		

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help               help for remove
  -n, --namespace string   Kubernetes namespace where the mesh is deployed (default "default")

</div>
</pre>

## Options inherited from parent commands

<pre class='codeblock-pre'>
<div class='codeblock'>
      --config string   path to config file (default "/home/runner/.meshplay/config.yaml")
  -t, --token string    Path to token for authenticating to Meshplay API
  -v, --verbose         verbose output

</div>
</pre>

## See Also

Go back to [command reference index](/reference/meshplayctl/), if you want to add content manually to the CLI documentation, please refer to the [instruction](/project/contributing/contributing-cli#preserving-manually-added-documentation) for guidance.
