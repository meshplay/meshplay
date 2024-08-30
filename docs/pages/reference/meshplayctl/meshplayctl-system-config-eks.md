---
layout: default
title: meshplayctl-system-config-eks
permalink: reference/meshplayctl/system/config/eks
redirect_from: reference/meshplayctl/system/config/eks/
type: reference
display-title: "false"
language: en
command: system
subcommand: config
---

# meshplayctl system config eks

Configure Meshplay to use EKS cluster

## Synopsis

Configure Meshplay to connect to EKS cluster
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system config eks [flags]

</div>
</pre> 

## Examples

Configure Meshplay to connect to EKS cluster using auth token
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system config eks --token auth.json

</div>
</pre> 

Configure Meshplay to connect to EKS cluster (if session is logged in using login subcommand)
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system config eks

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help           help for eks
  -t, --token string   Path to token for authenticating to Meshplay API

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
