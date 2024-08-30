---
layout: default
title: meshplayctl-system-logs
permalink: reference/meshplayctl/system/logs
redirect_from: reference/meshplayctl/system/logs/
type: reference
display-title: "false"
language: en
command: system
subcommand: logs
---

# meshplayctl system logs

Print logs

## Synopsis

Print history of Meshplay's logs and begin tailing them.

It also shows the logs of a specific component.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system logs [flags]

</div>
</pre> 

## Examples

Show logs (without tailing)
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system logs

</div>
</pre> 

Starts tailing Meshplay server debug logs (works with components also)
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system logs --follow

</div>
</pre> 

<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system logs meshplay-istio

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -f, --follow   (Optional) Follow the stream of the Meshplay's logs. Defaults to false.
  -h, --help     help for logs

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
