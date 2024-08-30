---
layout: default
title: meshplayctl-exp-components-view
permalink: reference/meshplayctl/exp/components/view
redirect_from: reference/meshplayctl/exp/components/view/
type: reference
display-title: "false"
language: en
command: exp
subcommand: components
---

# meshplayctl exp components view

view registered components

## Synopsis

view a component registered in Meshplay Server
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp components view [flags]

</div>
</pre> 

## Examples

View details of a specific component
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp components view [component-name]

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help                   help for view
  -o, --output-format string   (optional) format to display in [json|yaml] (default "yaml")

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
