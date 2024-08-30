---
layout: default
title: meshplayctl-exp-environment
permalink: reference/meshplayctl/exp/environment
redirect_from: reference/meshplayctl/exp/environment/
type: reference
display-title: "false"
language: en
command: exp
subcommand: environment
---

# meshplayctl exp environment

View list of environments and detail of environments

## Synopsis

View list of environments and detailed information of a specific environments
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp environment [flags]

</div>
</pre> 

## Examples

To view a list environments
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp environment list --orgID [orgId]

</div>
</pre> 

To create a environment
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp environment create --orgID [orgId] --name [name] --description [description]

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
  -h, --help   help for environment

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
