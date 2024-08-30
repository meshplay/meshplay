---
layout: default
title: meshplayctl-model-list
permalink: reference/meshplayctl/model/list
redirect_from: reference/meshplayctl/model/list/
type: reference
display-title: "false"
language: en
command: model
subcommand: list
---

# meshplayctl model list

list registered models

## Synopsis

list name of all registered models
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl model list [flags]

</div>
</pre> 

## Examples

View list of models
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl model list

</div>
</pre> 

View list of models with specified page number (25 models per page)
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl model list --page 2

</div>
</pre> 

View number of available models in Meshplay
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl model list --count

</div>
</pre> 

<pre class='codeblock-pre'>
<div class='codeblock'>
    

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help       help for list
  -p, --page int   (optional) List next set of models with --page (default = 1) (default 1)

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
