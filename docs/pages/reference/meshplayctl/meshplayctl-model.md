---
layout: default
title: meshplayctl-model
permalink: reference/meshplayctl/model
redirect_from: reference/meshplayctl/model/
type: reference
display-title: "false"
language: en
command: model
subcommand: nil
---

# meshplayctl model

View list of models and detail of models

## Synopsis

View list of models and detailed information of a specific model
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl model [flags]

</div>
</pre> 

## Examples

To view total of available models
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl model --count

</div>
</pre> 

To view list of models
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl model list

</div>
</pre> 

To view a specific model
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl model view [model-name]

</div>
</pre> 

To search for a specific model
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl model search [model-name]

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
      --count   (optional) Get the number of models in total
  -h, --help    help for model

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
