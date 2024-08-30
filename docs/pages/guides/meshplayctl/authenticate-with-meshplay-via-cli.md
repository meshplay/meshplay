---
layout: default
title: Authenticating with Meshplay via CLI
permalink: guides/meshplayctl/authenticate-with-meshplay-via-cli
language: en
type: guides
category: meshplayctl
list: include
abstract: Get your authentication token from Meshplay CLI.
---

To authenticate with Meshplay through `meshplayctl` you will use the command `meshplayctl system login`. Upon execution of this command, select your Provider of choice, then authenticate to your chosen Provider.

## Get your Token

You can retrieve your authentication token from either of Meshplay's two clients: the CLI or the UI.

- Get your token through [Meshplay UI](/extensibility/api#how-to-get-your-token), from the `Get Token` option.

  _Downloading the token_

  <a href="{{ site.baseurl }}/assets/img/token/MeshplayTokenUI.png"><img alt="Meshplay Dashboard" src="{{ site.baseurl }}/assets/img/token/MeshplayTokenUI.png" /></a>
  <br/>
  <br/>

- Get your token through **Meshplay CLI**.
  <br/>
  To get the token through `meshplayctl` you would have to use the following command and the path to token for authenticating to Meshplay API (default "auth.json").
  <br/>
  <pre class="codeblock-pre">
  <div class="codeblock"><div class="clipboardjs"> meshplayctl system login</div></div>
  </pre>
  <br />

**_The need for authentication to `Meshplay` [provider](https://docs.meshplay.io/extensibility/providers) is to save your environment setup while also having persistent/steady sessions and to be able to retrieve performance test results._**

<br/>

For an exhaustive list of `meshplayctl` commands and syntax:

- See [`meshplayctl` Command Reference]({{ site.baseurl }}/reference/meshplayctl).

Guides to using Meshplay's various features and components.

{% include related-discussions.html tag="meshplayctl" %}

