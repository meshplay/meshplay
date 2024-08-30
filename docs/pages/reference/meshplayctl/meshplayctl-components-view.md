---
layout: default
title: meshplayctl-components-view
permalink: reference/meshplayctl/components/view
redirect_from: reference/meshplayctl/components/view/
type: reference
display-title: "false"
language: en
command: components
subcommand: view
---

# meshplayctl components view

view registered components

## Synopsis

view a component registered in Meshplay Server
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl components view [flags]

</div>
</pre> 

## Examples

View details of a specific component
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl components view [component-name]

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help                   help for view
  -o, --output-format string   (optional) format to display in [json|yaml] (default "yaml")
  -s, --save                   (optional) save output as a JSON/YAML file

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
