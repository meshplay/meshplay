---
layout: default
title: meshplayctl-system-channel-switch
permalink: reference/meshplayctl/system/channel/switch
redirect_from: reference/meshplayctl/system/channel/switch/
type: reference
display-title: "false"
language: en
command: system
subcommand: channel
---

# meshplayctl system channel switch

switch release channel and version

## Synopsis

Switch release channel and version of context in focus
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system channel switch [stable|stable-version|edge|edge-version] [flags]

</div>
</pre> 

## Examples

Switch between release channels
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system channel switch [stable|stable-version|edge|edge-version]

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

## See Also

Go back to [command reference index](/reference/meshplayctl/), if you want to add content manually to the CLI documentation, please refer to the [instruction](/project/contributing/contributing-cli#preserving-manually-added-documentation) for guidance.
