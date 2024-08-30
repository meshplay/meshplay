---
layout: default
title: meshplayctl-system-reset
permalink: reference/meshplayctl/system/reset
redirect_from: reference/meshplayctl/system/reset/
type: reference
display-title: "false"
language: en
command: system
subcommand: reset
---

# meshplayctl system reset

Reset Meshplay's configuration

## Synopsis

Reset Meshplay to it's default configuration.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system reset [flags]

</div>
</pre> 

## Examples

Resets meshplay.yaml file with a copy from Meshplay repo
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system reset

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help   help for reset

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

Usage of meshplayctl system reset
![reset-usage](/assets/img/meshplayctl/reset.png)

## See Also

Go back to [command reference index](/reference/meshplayctl/), if you want to add content manually to the CLI documentation, please refer to the [instruction](/project/contributing/contributing-cli#preserving-manually-added-documentation) for guidance.
