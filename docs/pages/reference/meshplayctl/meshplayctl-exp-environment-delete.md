---
layout: default
title: meshplayctl-exp-environment-delete
permalink: reference/meshplayctl/exp/environment/delete
redirect_from: reference/meshplayctl/exp/environment/delete/
type: reference
display-title: "false"
language: en
command: exp
subcommand: environment
---

# meshplayctl exp environment delete

delete a new environments

## Synopsis

delete a new environments by providing the name and description of the environment
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp environment delete [flags]

</div>
</pre> 

## Examples

delete a new environment
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp environment delete [environmentId]

</div>
</pre> 

Documentation for environment can be found at:
<pre class='codeblock-pre'>
<div class='codeblock'>
https://docs.khulnasoft.com/cloud/spaces/environments/

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help   help for delete

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
