

### Prerequisites

You need to have `Brew` installed on your **Linux** or **macOS** system to perform these actions.

### Install `meshplayctl` using Brew

To install `meshplayctl` using homebrew, execute the following commands.

<pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">
 $ brew install meshplayctl
 </div></div>
</pre>

You're ready to run Meshplay. To do so, execute the following command.

<pre class="codeblock-pre"><div class="codeblock">
<div class="clipboardjs">
 $ meshplayctl system start

</div></div>
</pre>

Meshplay server supports customizing authentication flow callback URL, which can be configured in the following way

<pre class="codeblock-pre"><div class="codeblock">
<div class="clipboardjs">
 $ MESHPLAY_SERVER_CALLBACK_URL=https://custom-host meshplayctl system start

</div></div>
</pre>

`meshplayctl` uses your current Kubernetes context, your KUBECONFIG environment variable (`~/.kube/config` by default). Confirm if this Kubernetes cluster you want Meshplay to interact with by running the following command: `kubectl config get-contexts`.

If there are multiple contexts in your kubeconfig file, specify the one you want to use with the `use-context` subcommand: `kubectl config use-context <context-to-use>`.

### Upgrade `meshplayctl` using Brew

To upgrade `meshplayctl`, execute the following command.

 <pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">
 $ brew upgrade meshplayctl
 </div></div>
 </pre>

<details>
<summary>
Example output of a successful upgrade.
</summary>

<pre><code>
➜  ~ brew upgrade meshplayctl
==> Upgrading 1 outdated package:
meshplay/tap/meshplayctl 0.3.2 -> 0.3.4
==> Upgrading meshplay/tap/meshplayctl
==> Downloading https://github.com/meshplay/meshplay/releases/download/v0.3.4/meshplayctl_0.3.4_Darwin_x86_64.zip
==> Downloading from https://github-production-release-asset-2e65be.s3.amazonaws.com/157554479/17522b00-2af0-11ea-8aef-cbfe8
######################################################################## 100.0%
🍺  /usr/local/Cellar/meshplayctl/0.3.4: 5 files, 10.2MB, built in 4 seconds
Removing: /usr/local/Cellar/meshplayctl/0.3.2... (5 files, 10.2MB)
Removing: /Users/lee/Library/Caches/Homebrew/meshplayctl--0.3.2.zip... (3.9MB)
==> Checking for dependents of upgraded formulae...
==> No dependents found!
</code></pre>
<br />
</details>
