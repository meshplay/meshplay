---
layout: default
title: meshplayctl-system-restart
permalink: reference/meshplayctl/system/restart
redirect_from: reference/meshplayctl/system/restart/
type: reference
display-title: "false"
language: en
command: system
subcommand: restart
---

# meshplayctl system restart

Stop, then start Meshplay

## Synopsis

Restart all Meshplay containers / pods.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system restart [flags]

</div>
</pre> 

## Examples

Restart all Meshplay containers, their instances and their connected volumes
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system restart

</div>
</pre> 

(optional) skip checking for new updates available in Meshplay.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system restart --skip-update

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help              help for restart
      --provider string   Provider to use with the Meshplay server
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
