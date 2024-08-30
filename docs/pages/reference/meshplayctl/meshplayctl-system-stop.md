---
layout: default
title: meshplayctl-system-stop
permalink: reference/meshplayctl/system/stop
redirect_from: reference/meshplayctl/system/stop/
type: reference
display-title: "false"
language: en
command: system
subcommand: stop
---

# meshplayctl system stop

Stop Meshplay

## Synopsis

Stop all Meshplay containers / remove all Meshplay resources.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system stop [flags]

</div>
</pre> 

## Examples

Stop Meshplay
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system stop

</div>
</pre> 

Reset Meshplay's configuration file to default settings.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system stop --reset

</div>
</pre> 

(optional) keep the Meshplay namespace during uninstallation
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system stop --keep-namespace

</div>
</pre> 

Stop Meshplay forcefully (use it when system stop doesn't work)
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system stop --force

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
      --force            (optional) uninstall Meshplay resources forcefully
  -h, --help             help for stop
      --keep-namespace   (optional) keep the Meshplay namespace during uninstallation
      --reset            (optional) reset Meshplay's configuration file to default settings.

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
