---
layout: default
title: meshplayctl-exp-workspace
permalink: reference/meshplayctl/exp/workspace
redirect_from: reference/meshplayctl/exp/workspace/
type: reference
display-title: "false"
language: en
command: exp
subcommand: workspace
---

# meshplayctl exp workspace

View list of workspaces and detail of workspaces

## Synopsis

View list of workspaces and detailed information of a specific workspaces
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp workspace [flags]

</div>
</pre> 

## Examples

To view a list workspaces
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp workspace list --orgId [orgId]

</div>
</pre> 

To create a workspace
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp workspace create --orgId [orgId] --name [name] --description [description]

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
  -h, --help   help for workspace

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
