---
layout: default
title: meshplayctl-model-import
permalink: reference/meshplayctl/model/import
redirect_from: reference/meshplayctl/model/import/
type: reference
display-title: "false"
language: en
command: model
subcommand: import
---

# meshplayctl model import

import models from meshplayctl command

## Synopsis

import model by specifying the directory, file. Use 'import model [filepath]' or 'import model  [directory]'.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl model import [flags]

</div>
</pre> 

## Examples

<pre class='codeblock-pre'>
<div class='codeblock'>
	import model  /path/to/[file.yaml|file.json]

</div>
</pre> 

<pre class='codeblock-pre'>
<div class='codeblock'>
	import model  /path/to/models

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help   help for import

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
