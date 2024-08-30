---
layout: default
title: meshplayctl-system-token-delete
permalink: reference/meshplayctl/system/token/delete
redirect_from: reference/meshplayctl/system/token/delete/
type: reference
display-title: "false"
language: en
command: system
subcommand: token
---

# meshplayctl system token delete

Delete a token from your meshconfig

## Synopsis

Delete the token with provided token name from your meshconfig tokens.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system token delete [flags]

</div>
</pre> 

## Examples

<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system token delete [token-name]

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
      --config string    path to config file (default "/home/runner/.meshplay/config.yaml")
  -c, --context string   (optional) temporarily change the current context.
  -v, --verbose          verbose output
  -y, --yes              (optional) assume yes for user interactive prompts.

</div>
</pre>

## See Also

Go back to [command reference index](/reference/meshplayctl/), if you want to add content manually to the CLI documentation, please refer to the [instruction](/project/contributing/contributing-cli#preserving-manually-added-documentation) for guidance.
