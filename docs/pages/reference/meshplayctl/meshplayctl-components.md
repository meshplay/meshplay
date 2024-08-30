---
layout: default
title: meshplayctl-components
permalink: reference/meshplayctl/components
redirect_from: reference/meshplayctl/components/
type: reference
display-title: "false"
language: en
command: components
subcommand: nil
---

# meshplayctl components

View list of components and detail of components

## Synopsis

View list of components and detailed information of a specific component
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl components [flags]

</div>
</pre> 

## Examples

To view total of available components
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl model --count

</div>
</pre> 

To view list of components
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl components list

</div>
</pre> 

To view a specific component
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl components view [component-name]

</div>
</pre> 

To search for a specific component
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl components search [component-name]

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
      --count   (optional) Get the number of components in total
  -h, --help    help for components

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
