---
layout: default
title: meshplayctl-pattern-view
permalink: reference/meshplayctl/pattern/view
redirect_from: reference/meshplayctl/pattern/view/
type: reference
display-title: "false"
language: en
command: pattern
subcommand: view
---

# meshplayctl pattern view

Display pattern(s)

## Synopsis

Displays the contents of a specific pattern based on name or id
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl pattern view pattern name [flags]

</div>
</pre> 

## Examples

view a pattern
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl pattern view [pattern-name | ID]

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -a, --all                    (optional) view all patterns available
  -h, --help                   help for view
  -o, --output-format string   (optional) format to display in [json|yaml] (default "yaml")

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

Usage of meshplayctl pattern view
![pattern-view-usage](/assets/img/meshplayctl/patternView.png)

## See Also

Go back to [command reference index](/reference/meshplayctl/), if you want to add content manually to the CLI documentation, please refer to the [instruction](/project/contributing/contributing-cli#preserving-manually-added-documentation) for guidance.
