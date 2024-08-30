---
layout: default
title: meshplayctl-components-list
permalink: reference/meshplayctl/components/list
redirect_from: reference/meshplayctl/components/list/
type: reference
display-title: "false"
language: en
command: components
subcommand: list
---

# meshplayctl components list

List registered components

## Synopsis

List all components registered in Meshplay Server
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl components list [flags]

</div>
</pre> 

## Examples

View list of components
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl components list

</div>
</pre> 

View list of components with specified page number (25 components per page)
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl components list --page 2

</div>
</pre> 

To view the number of components present in Meshplay
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl components list --count

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help       help for list
  -p, --page int   (optional) List next set of components with --page (default = 1) (default 1)

</div>
</pre>

## Options inherited from parent commands

<pre class='codeblock-pre'>
<div class='codeblock'>
      --config string   path to config file (default "/home/runner/.meshplay/config.yaml")
  -v, --verbose         verbose output

</div>
</pre>

## See Also

Go back to [command reference index](/reference/meshplayctl/), if you want to add content manually to the CLI documentation, please refer to the [instruction](/project/contributing/contributing-cli#preserving-manually-added-documentation) for guidance.
