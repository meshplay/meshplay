---
layout: default
title: meshplayctl
permalink: reference/meshplayctl/main
redirect_from: reference/meshplayctl/main/
type: reference
display-title: "false"
language: en
command: meshplayctl
subcommand: nil
---

# meshplayctl

Meshplay Command Line tool

## Synopsis

As a self-service engineering platform, Meshplay enables collaborative design and operation of cloud native infrastructure.

<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl [flags]

</div>
</pre> 

## Examples

Base command:
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl

</div>
</pre> 

Display help about command/subcommand:
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl --help

</div>
</pre> 

<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system start --help

</div>
</pre> 

For viewing verbose output:
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl -v [or] --verbose

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
      --config string   path to config file (default "/home/runner/.meshplay/config.yaml")
  -h, --help            help for meshplayctl
  -v, --verbose         verbose output

</div>
</pre>

## See Also

Go back to [command reference index](/reference/meshplayctl/), if you want to add content manually to the CLI documentation, please refer to the [instruction](/project/contributing/contributing-cli#preserving-manually-added-documentation) for guidance.
