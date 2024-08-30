---
layout: default
title: meshplayctl-pattern-apply
permalink: reference/meshplayctl/pattern/apply
redirect_from: reference/meshplayctl/pattern/apply/
type: reference
display-title: "false"
language: en
command: pattern
subcommand: apply
---

# meshplayctl pattern apply

Apply pattern file

## Synopsis

Apply pattern file will trigger deploy of the pattern file
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl pattern apply [flags]

</div>
</pre> 

## Examples

apply a pattern file
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl pattern apply -f [file | URL]

</div>
</pre> 

deploy a saved pattern
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl pattern apply [pattern-name]

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -f, --file string   Path to pattern file
  -h, --help          help for apply
      --skip-save     Skip saving a pattern

</div>
</pre>

## Options inherited from parent commands

<pre class='codeblock-pre'>
<div class='codeblock'>
      --config string   path to config file (default "/home/runner/.meshplay/config.yaml")
  -t, --token string    Path to token file default from current context
  -v, --verbose         verbose output

</div>
</pre>

## Screenshots

Usage of meshplayctl pattern apply
![pattern-apply-usage](/assets/img/meshplayctl/patternApply.png)

## See Also

Go back to [command reference index](/reference/meshplayctl/), if you want to add content manually to the CLI documentation, please refer to the [instruction](/project/contributing/contributing-cli#preserving-manually-added-documentation) for guidance.
