---
layout: default
title: meshplayctl-exp-connections-delete
permalink: reference/meshplayctl/exp/connections/delete
redirect_from: reference/meshplayctl/exp/connections/delete/
type: reference
display-title: "false"
language: en
command: exp
subcommand: connections
---

# meshplayctl exp connections delete

Delete a connection

## Synopsis

Delete
a connection
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp connections delete [flags]

</div>
</pre> 

## Examples

Delete a connection
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp connections delete [connection_id]

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help        help for delete
  -i, --id string   ID of the connection to be deleted

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
