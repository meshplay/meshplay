---
layout: default
title: meshplayctl-registry-generate
permalink: reference/meshplayctl/registry/generate
redirect_from: reference/meshplayctl/registry/generate/
type: reference
display-title: "false"
language: en
command: registry
subcommand: generate
---

# meshplayctl registry generate

Generate Models

## Synopsis

Prerequisite: Excecute this command from the root of a meshplay/meshplay repo fork.

Given a Google Sheet with a list of model names and source locations, generate models and components any Registrant (e.g. GitHub, Artifact Hub) repositories.

Generated Model files are written to local filesystem under `/server/models/<model-name>`.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl registry generate [flags]

</div>
</pre> 

## Examples

Generate Meshplay Models from a Google Spreadsheet (i.e. "Meshplay Integrations" spreadsheet).
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl registry generate --spreadsheet-id "1DZHnzxYWOlJ69Oguz4LkRVTFM79kC2tuvdwizOJmeMw" --spreadsheet-cred 

</div>
</pre> 

Directly generate models from one of the supported registrants by using Registrant Connection Definition and (optional) Registrant Credential Definition
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl registry generate --registrant-def [path to connection definition] --registrant-cred [path to credential definition]

</div>
</pre> 

Generate a specific Model from a Google Spreadsheet (i.e. "Meshplay Integrations" spreadsheet).
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl registry generate --spreadsheet-id "1DZHnzxYWOlJ69Oguz4LkRVTFM79kC2tuvdwizOJmeMw" --spreadsheet-cred --model "[model-name]"

</div>
</pre> 

<pre class='codeblock-pre'>
<div class='codeblock'>
    

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help                      help for generate
  -m, --model string              specific model name to be generated
  -o, --output string             location to output generated models, defaults to ../server/meshmodels (default "../server/meshmodel")
      --registrant-cred string    path pointing to the registrant credential definition
      --registrant-def string     path pointing to the registrant connection definition
      --spreadsheet-cred string   base64 encoded credential to download the spreadsheet
      --spreadsheet-id string     spreadsheet ID for the integration spreadsheet

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
