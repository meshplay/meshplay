---
layout: default
title: meshplayctl-system-context-create
permalink: reference/meshplayctl/system/context/create
redirect_from: reference/meshplayctl/system/context/create/
type: reference
display-title: "false"
language: en
command: system
subcommand: context
---

# meshplayctl system context create

Create a new context (a named Meshplay deployment)

## Synopsis

Add a new context to Meshplay config.yaml file
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system context create context-name [flags]

</div>
</pre> 

## Examples

Create new context
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system context create [context-name]

</div>
</pre> 

Create new context and provide list of components, platform & URL
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system context create context-name --components meshplay-nsm --platform docker --url http://localhost:9081 --set --yes

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -a, --components stringArray   List of components
  -h, --help                     help for create
  -p, --platform string          Platform to deploy Meshplay
      --provider string          Provider to use with the Meshplay server
  -s, --set                      Set as current context
  -u, --url string               Meshplay Server URL with Port

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

## Screenshots

Usage of meshplayctl context create
![context-create-usage](/assets/img/meshplayctl/newcontext.png)

## See Also

Go back to [command reference index](/reference/meshplayctl/), if you want to add content manually to the CLI documentation, please refer to the [instruction](/project/contributing/contributing-cli#preserving-manually-added-documentation) for guidance.
