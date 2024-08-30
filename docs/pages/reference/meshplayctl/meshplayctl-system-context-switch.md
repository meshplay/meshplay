---
layout: default
title: meshplayctl-system-context-switch
permalink: reference/meshplayctl/system/context/switch
redirect_from: reference/meshplayctl/system/context/switch/
type: reference
display-title: "false"
language: en
command: system
subcommand: context
---

# meshplayctl system context switch

switch context

## Synopsis

Configure meshplayctl to actively use one one context vs. another context
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system context switch context-name [flags]

</div>
</pre> 

## Examples

Switch to context named "sample"
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system context switch sample

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help   help for switch

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

Usage of meshplayctl context switch
![context-switch-usage](/assets/img/meshplayctl/contextswitch.png)

## See Also

Go back to [command reference index](/reference/meshplayctl/), if you want to add content manually to the CLI documentation, please refer to the [instruction](/project/contributing/contributing-cli#preserving-manually-added-documentation) for guidance.
