---
layout: default
title: meshplayctl-system-dashboard
permalink: reference/meshplayctl/system/dashboard
redirect_from: reference/meshplayctl/system/dashboard/
type: reference
display-title: "false"
language: en
command: system
subcommand: dashboard
---

# meshplayctl system dashboard

Open Meshplay UI in browser.

<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system dashboard [flags]

</div>
</pre> 

## Examples

Open Meshplay UI in browser
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system dashboard

</div>
</pre> 

Open Meshplay UI in browser and use port-forwarding (if default port is taken already)
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system dashboard --port-forward

</div>
</pre> 

Open Meshplay UI in browser and use port-forwarding, listen on port 9081 locally, forwarding traffic to meshplay server in the pod
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system dashboard --port-forward -p 9081

</div>
</pre> 

(optional) skip opening of MeshplayUI in browser.
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl system dashboard --skip-browser

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help           help for dashboard
  -p, --port int       (optional) Local port that is not in use from which traffic is to be forwarded to the server running inside the Pod. (default 9081)
      --port-forward   (optional) Use port forwarding to access Meshplay UI
      --skip-browser   (optional) skip opening of MeshplayUI in browser.

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


{% include meshplayctl/system-dashboard.md %}
## See Also

Go back to [command reference index](/reference/meshplayctl/), if you want to add content manually to the CLI documentation, please refer to the [instruction](/project/contributing/contributing-cli#preserving-manually-added-documentation) for guidance.
