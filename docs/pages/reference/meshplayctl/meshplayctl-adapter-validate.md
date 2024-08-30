---
layout: default
title: meshplayctl-adapter-validate
permalink: reference/meshplayctl/adapter/validate
redirect_from: reference/meshplayctl/adapter/validate/
type: reference
display-title: "false"
language: en
command: adapter
subcommand: validate
---

# meshplayctl adapter validate

Validate conformance to predefined standards

## Synopsis

Validate predefined conformance to different standard specifications
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl adapter validate [flags]

</div>
</pre> 

## Examples

Validate conformance to predefined standards
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl adapter validate [mesh name] --adapter [name of the adapter] --tokenPath [path to token for authentication] --spec [specification to be used for conformance test] --namespace [namespace to be used]

</div>
</pre> 

Validate Istio to predefined standards
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl adapter validate istio --adapter meshplay-istio --spec smi

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -a, --adapter string   (Required) Adapter to use for validation (default "meshplay-nsm")
  -h, --help             help for validate
  -s, --spec string      (Required) specification to be used for conformance test (smi/istio-vet) (default "smi")
  -t, --token string     Path to token for authenticating to Meshplay API
  -w, --watch            Watch for events and verify operation (in beta testing)

</div>
</pre>

## Options inherited from parent commands

<pre class='codeblock-pre'>
<div class='codeblock'>
      --config string   path to config file (default "/home/runner/.meshplay/config.yaml")
  -v, --verbose         verbose output

</div>
</pre>

## Screenshots

Usage of meshplayctl adapter validate
![mesh-validate-usage](/assets/img/meshplayctl/mesh-validate.png)

## See Also

Go back to [command reference index](/reference/meshplayctl/), if you want to add content manually to the CLI documentation, please refer to the [instruction](/project/contributing/contributing-cli#preserving-manually-added-documentation) for guidance.
