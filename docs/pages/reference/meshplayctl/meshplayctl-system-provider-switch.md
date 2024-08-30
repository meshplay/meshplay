---
layout: default
title: meshplayctl-system-provider-switch
permalink: reference/meshplayctl/system/provider/switch
redirect_from: reference/meshplayctl/system/provider/switch/
type: reference
display-title: "false"
language: en
command: system
subcommand: provider
---

# meshplayctl system provider switch

switch provider and redeploy

## Synopsis

Switch provider of context in focus and redeploy Meshplay. Run `meshplayctl system provider list` to see the available providers.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system provider switch [provider] [flags]

</div>
</pre> 

## Examples

Switch provider and redeploy Meshplay
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system provider switch [provider]

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
