---
layout: default
title: Configuring Autocompletion for `meshplayctl`
permalink: guides/meshplayctl/configuring-autocompletion-for-meshplayctl
language: en
type: guides
category: meshplayctl
list: include
abstract: Bash, Zsh, Oh My Zsh, and fish autocompletion for `meshplayctl` commands.
---

If you would like to have `meshplayctl` commands automatically completed for use as you use `meshplayctl`, then use the following instructions to configure automatic completion within your environment.

## Autocompletion for Bash

### bash <= 3.2

 <pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">source /dev/stdin <<< "$(meshplayctl completion bash)"</div></div>
 </pre>

### bash >= 4.0

 <pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">source <(meshplayctl completion bash)</div></div>
 </pre>

### bash <= 3.2 on MacOS

 <pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">brew install bash-completion # ensure you have bash-completion 1.3+
meshplayctl completion bash > $(brew --prefix)/etc/bash_completion.d/meshplayctl</div></div>
 </pre>

### bash >= 4.0 on MacOS

 <pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">brew install bash-completion@2
meshplayctl completion bash > $(brew --prefix)/etc/bash_completion.d/meshplayctl</div></div>
 </pre>

## Autocompletion for zsh

 <pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">source <(meshplayctl completion zsh)</div></div>
 </pre><br>

If shell completion is not already enabled in your environment you will need to enable it. You can execute the following once:

 <pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">~/.zshrc > echo "autoload -U compinit; compinit"</div></div>
 </pre>

_Note_ : You might need to restart your shell for this setup to take effect.

#### zsh on MacOS and Oh My zsh

 <pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">COMPLETION_DIR=$(echo $fpath | grep -o '[^ ]*completions' | grep -v cache) && mkdir -p $COMPLETION_DIR && meshplayctl completion zsh > "${COMPLETION_DIR}/_meshplayctl"</div></div>
 </pre>

### Autocompletion for fish

 <pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">meshplayctl completion fish | source</div></div>
 </pre><br>

To load fish shell completions for each session, execute once:

 <pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">meshplayctl completion fish > ~/.config/fish/completions/meshplayctl.fish</div></div>
 </pre>

For an exhaustive list of `meshplayctl` commands and syntax:

- See [`meshplayctl` Command Reference]({{ site.baseurl }}/reference/meshplayctl).

Guides to using Meshplay's various features and components.

{% capture tag %}

<li><a href="{{ site.baseurl }}/installation/upgrades#upgrading-meshplay-cli">Upgrading meshplayctl</a></li>

{% endcapture %}


{% include related-discussions.html tag="meshplayctl" %}
