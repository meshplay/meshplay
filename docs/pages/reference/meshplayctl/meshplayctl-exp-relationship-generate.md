---
layout: default
title: meshplayctl-exp-relationship-generate
permalink: reference/meshplayctl/exp/relationship/generate
redirect_from: reference/meshplayctl/exp/relationship/generate/
type: reference
display-title: "false"
language: en
command: exp
subcommand: relationship
---

# meshplayctl exp relationship generate

generate relationships docs

## Synopsis

generate relationships docs from the google spreadsheets
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl exp relationship generate [flags]

</div>
</pre> 

## Examples

<pre class='codeblock-pre'>
<div class='codeblock'>
    // generate relationships docs

</div>
</pre> 

<pre class='codeblock-pre'>
<div class='codeblock'>
    meshplayctl relationships generate $CRED

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help             help for generate
  -s, --sheetId string   Google Sheet ID

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
