---
layout: default
title: meshplayctl-exp-workspace-list
permalink: reference/meshplayctl/exp/workspace/list
redirect_from: reference/meshplayctl/exp/workspace/list/
type: reference
display-title: "false"
language: en
command: exp
subcommand: workspace
---

# meshplayctl exp workspace list

List registered workspaces

## Synopsis

List name of all registered workspaces
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp workspace list [flags]

</div>
</pre> 

## Examples

List all registered workspace
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp workspace list --orgId [orgId]

</div>
</pre> 

Documentation for workspace can be found at:
<pre class='codeblock-pre'>
<div class='codeblock'>
https://docs.khulnasoft.com/cloud/spaces/workspaces/

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
