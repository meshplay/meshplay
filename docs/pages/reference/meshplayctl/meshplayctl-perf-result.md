---
layout: default
title: meshplayctl-perf-result
permalink: reference/meshplayctl/perf/result
redirect_from: reference/meshplayctl/perf/result/
type: reference
display-title: "false"
language: en
command: perf
subcommand: result
---

# meshplayctl perf result

List performance test results

## Synopsis

List all the available test results of a performance profile
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl perf result [profile-name] [flags]

</div>
</pre> 

## Examples

List Test results (maximum 25 results)
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl perf result saturday-profile

</div>
</pre> 

View other set of performance results with --page (maximum 25 results)
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl perf result saturday-profile --page 2

</div>
</pre> 

View single performance result with detailed information
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl perf result saturday-profile --view

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help       help for result
  -p, --page int   (optional) List next set of performance results with --page (default = 1) (default 1)
      --view       (optional) View single performance results with more info

</div>
</pre>

## Options inherited from parent commands

<pre class='codeblock-pre'>
<div class='codeblock'>
      --config string          path to config file (default "/home/runner/.meshplay/config.yaml")
  -o, --output-format string   (optional) format to display in [json|yaml]
  -t, --token string           (required) Path to meshplay auth config
  -v, --verbose                verbose output
  -y, --yes                    (optional) assume yes for user interactive prompts.

</div>
</pre>

## Screenshots

Usage of meshplayctl perf result
![perf-result-usage](/assets/img/meshplayctl/perf-result.png)

## See Also

Go back to [command reference index](/reference/meshplayctl/), if you want to add content manually to the CLI documentation, please refer to the [instruction](/project/contributing/contributing-cli#preserving-manually-added-documentation) for guidance.
