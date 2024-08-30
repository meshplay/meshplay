---
layout: default
title: meshplayctl-completion
permalink: reference/meshplayctl/completion
redirect_from: reference/meshplayctl/completion/
type: reference
display-title: "false"
language: en
command: completion
subcommand: nil
---

# meshplayctl completion

Output shell completion code

## Synopsis

Output shell completion code
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl completion [bash|zsh|fish]

</div>
</pre> 

## Examples

### bash <= 3.2
<pre class='codeblock-pre'>
<div class='codeblock'>
source /dev/stdin <<< "$(meshplayctl completion bash)"

</div>
</pre> 

bash <= 3.2 on osx
ensure you have bash-completion 1.3+
<pre class='codeblock-pre'>
<div class='codeblock'>
brew install bash-completion 

</div>
</pre> 

<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl completion bash > $(brew --prefix)/etc/bash_completion.d/meshplayctl

</div>
</pre> 

### bash >= 4.0
<pre class='codeblock-pre'>
<div class='codeblock'>
source <(meshplayctl completion bash)

</div>
</pre> 

bash >= 4.0 on osx
<pre class='codeblock-pre'>
<div class='codeblock'>
brew install bash-completion@2

</div>
</pre> 

<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl completion bash > $(brew --prefix)/etc/bash_completion.d/meshplayctl

</div>
</pre> 

### zsh
If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:
Might need to start a new shell for this setup to take effect.
<pre class='codeblock-pre'>
<div class='codeblock'>
$ echo "autoload -U compinit; compinit" >> ~/.zshrc

</div>
</pre> 

<pre class='codeblock-pre'>
<div class='codeblock'>
source <(meshplayctl completion zsh)

</div>
</pre> 

zsh on osx / oh-my-zsh
<pre class='codeblock-pre'>
<div class='codeblock'>
COMPLETION_DIR=$(echo $fpath | grep -o '[^ ]*completions' | grep -v cache) && mkdir -p $COMPLETION_DIR && meshplayctl completion zsh > "${COMPLETION_DIR}/_meshplayctl"

</div>
</pre> 

### fish:
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl completion fish | source

</div>
</pre> 

To load fish shell completions for each session, execute once:
<pre class='codeblock-pre'>
<div class='codeblock'>
meshplayctl completion fish > ~/.config/fish/completions/meshplayctl.fish

</div>
</pre> 

## Options

<pre class='codeblock-pre'>
<div class='codeblock'>
  -h, --help   help for completion

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
