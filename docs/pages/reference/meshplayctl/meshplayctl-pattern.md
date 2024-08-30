---
layout: default
title: meshplayctl-pattern
permalink: reference/meshplayctl/pattern
redirect_from: reference/meshplayctl/pattern/
type: reference
display-title: "false"
language: en
command: pattern
subcommand: nil
---

# meshplayctl pattern

Cloud Native Patterns Management

## Synopsis

Manage cloud and cloud native infrastructure using predefined patterns.

<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl pattern [flags]

</div>
</pre> 

## Examples

Apply pattern file:
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl pattern apply --file [path to pattern file | URL of the file]

</div>
</pre> 

Delete pattern file:
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl pattern delete --file [path to pattern file]

</div>
</pre> 

View pattern file:
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl pattern view [pattern name | ID]

</div>
</pre> 

List all patterns:
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl pattern list

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help           help for pattern
  -t, --token string   Path to token file default from current context

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
