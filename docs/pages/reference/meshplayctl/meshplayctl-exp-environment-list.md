---
layout: default
title: meshplayctl-exp-environment-list
permalink: reference/meshplayctl/exp/environment/list
redirect_from: reference/meshplayctl/exp/environment/list/
type: reference
display-title: "false"
language: en
command: exp
subcommand: environment
---

# meshplayctl exp environment list

List registered environments

## Synopsis

List name of all registered environments
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp environment list [flags]

</div>
</pre> 

## Examples

List all registered environment
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp environment list --orgID [orgId]

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
  -h, --help           help for list
  -o, --orgId string   Organization ID

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
