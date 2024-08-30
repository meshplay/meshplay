---
layout: default
title: meshplayctl-model-view
permalink: reference/meshplayctl/model/view
redirect_from: reference/meshplayctl/model/view/
type: reference
display-title: "false"
language: en
command: model
subcommand: view
---

# meshplayctl model view

view model

## Synopsis

view a model queried by its name
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl model view [flags]

</div>
</pre> 

## Examples

View current provider
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl model view [model-name]

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
