---
layout: default
title: meshplayctl-system-logout
permalink: reference/meshplayctl/system/logout
redirect_from: reference/meshplayctl/system/logout/
type: reference
display-title: "false"
language: en
command: system
subcommand: logout
---

# meshplayctl system logout

Remove authentication for Meshplay Server

## Synopsis


Remove authentication for Meshplay Server

This command removes the authentication token from the user's filesystem
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system logout [flags]

</div>
</pre> 

## Examples

Logout current session with your Meshplay Provider.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system logout

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help   help for logout

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
