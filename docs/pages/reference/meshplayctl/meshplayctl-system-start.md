---
layout: default
title: meshplayctl-system-start
permalink: reference/meshplayctl/system/start
redirect_from: reference/meshplayctl/system/start/
type: reference
display-title: "false"
language: en
command: system
subcommand: start
---

# meshplayctl system start

Start Meshplay

## Synopsis

Start Meshplay and each of its cloud native components.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system start [flags]

</div>
</pre> 

## Examples

Start meshplay
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system start

</div>
</pre> 

(optional) skip opening of MeshplayUI in browser.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system start --skip-browser

</div>
</pre> 

(optional) skip checking for new updates available in Meshplay.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system start --skip-update

</div>
</pre> 

Reset Meshplay's configuration file to default settings.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system start --reset

</div>
</pre> 

Specify Platform to deploy Meshplay to.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system start -p docker

</div>
</pre> 

Specify Provider to use.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system start --provider Meshplay

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help              help for start
  -p, --platform string   platform to deploy Meshplay to.
      --provider string   (optional) Defaults to the provider specified in the current context
      --reset             (optional) reset Meshplay's configuration file to default settings.
      --skip-browser      (optional) skip opening of MeshplayUI in browser.
      --skip-update       (optional) skip checking for new Meshplay's container images.

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
