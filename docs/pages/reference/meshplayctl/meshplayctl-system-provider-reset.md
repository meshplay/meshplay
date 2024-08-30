---
layout: default
title: meshplayctl-system-provider-reset
permalink: reference/meshplayctl/system/provider/reset
redirect_from: reference/meshplayctl/system/provider/reset/
type: reference
display-title: "false"
language: en
command: system
subcommand: provider
---

# meshplayctl system provider reset

reset provider to default

## Synopsis

Reset provider for current context to default (Meshplay)
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system provider reset [flags]

</div>
</pre> 

## Examples

Reset provider to default
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system provider reset

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

## See Also

Go back to [command reference index](/reference/meshplayctl/), if you want to add content manually to the CLI documentation, please refer to the [instruction](/project/contributing/contributing-cli#preserving-manually-added-documentation) for guidance.
