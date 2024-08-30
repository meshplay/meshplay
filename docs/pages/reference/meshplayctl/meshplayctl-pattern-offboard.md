---
layout: default
title: meshplayctl-pattern-offboard
permalink: reference/meshplayctl/pattern/offboard
redirect_from: reference/meshplayctl/pattern/offboard/
type: reference
display-title: "false"
language: en
command: pattern
subcommand: offboard
---

# meshplayctl pattern offboard

Offboard pattern

## Synopsis

Offboard pattern will trigger undeploy of pattern
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl pattern offboard [flags]

</div>
</pre> 

## Examples

Offboard pattern by providing file path
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl pattern offboard -f [filepath]

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -f, --file string   Path to pattern file
  -h, --help          help for offboard

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

## See Also

Go back to [command reference index](/reference/meshplayctl/), if you want to add content manually to the CLI documentation, please refer to the [instruction](/project/contributing/contributing-cli#preserving-manually-added-documentation) for guidance.
