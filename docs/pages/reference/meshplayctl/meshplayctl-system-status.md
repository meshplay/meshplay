---
layout: default
title: meshplayctl-system-status
permalink: reference/meshplayctl/system/status
redirect_from: reference/meshplayctl/system/status/
type: reference
display-title: "false"
language: en
command: system
subcommand: status
---

# meshplayctl system status

Check Meshplay status

## Synopsis

Check status of Meshplay and Meshplay components.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system status [flags]

</div>
</pre> 

## Examples

Check status of Meshplay, Meshplay adapters, Meshplay Operator and its controllers.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system status

</div>
</pre> 

(optional) Extra data in status table
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system status --verbose

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help      help for status
  -v, --verbose   (optional) Extra data in status table

</div>
</pre>

## Options inherited from parent commands

<pre class='codeblock-pre'>
<div class='codeblock'>
      --config string    path to config file (default "/home/runner/.meshplay/config.yaml")
  -c, --context string   (optional) temporarily change the current context.
  -y, --yes              (optional) assume yes for user interactive prompts.

</div>
</pre>

## Screenshots

Usage of meshplayctl system status
![status-usage](/assets/img/meshplayctl/status.png)

## See Also

Go back to [command reference index](/reference/meshplayctl/), if you want to add content manually to the CLI documentation, please refer to the [instruction](/project/contributing/contributing-cli#preserving-manually-added-documentation) for guidance.
