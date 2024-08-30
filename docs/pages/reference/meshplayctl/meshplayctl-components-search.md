---
layout: default
title: meshplayctl-components-search
permalink: reference/meshplayctl/components/search
redirect_from: reference/meshplayctl/components/search/
type: reference
display-title: "false"
language: en
command: components
subcommand: search
---

# meshplayctl components search

search registered components

## Synopsis

search components registered in Meshplay Server based on kind
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl components search [flags]

</div>
</pre> 

## Examples

Search for components using a query
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl components search [query-text]

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help   help for search

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
