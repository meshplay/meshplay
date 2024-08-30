---
layout: default
title: meshplayctl-system-login
permalink: reference/meshplayctl/system/login
redirect_from: reference/meshplayctl/system/login/
type: reference
display-title: "false"
language: en
command: system
subcommand: login
---

# meshplayctl system login

Authenticate to a Meshplay Server

## Synopsis


Authenticate to the Local or a Remote Provider of a Meshplay Server

The authentication mode is web-based browser flow
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system login [flags]

</div>
</pre> 

## Examples

Login with the Meshplay Provider of your choice: the Local Provider or a Remote Provider.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system login

</div>
</pre> 

Login with the Meshplay Provider by specifying it via -p or --provider flag.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system login -p Meshplay

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help              help for login
  -p, --provider string   login Meshplay with specified provider

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
