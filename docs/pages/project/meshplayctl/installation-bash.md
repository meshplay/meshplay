
To install or upgrade `meshplayctl` using `bash`, execute anyone of the following commands.

#### Option 1: Only install `meshplayctl` binary

 <pre class="codeblock-pre">
 <div class="codeblock">
 <div class="clipboardjs">
  $ curl -L https://meshplay.khulnasofy.com/install | DEPLOY_MESHPLAY=false bash -
 </div></div>
 </pre>
<br />
<br />
#### Option 2: Install `meshplayctl` binary and deploy Meshplay on Docker

 <pre class="codeblock-pre">
 <div class="codeblock">
 <div class="clipboardjs">
  $ curl -L https://meshplay.khulnasofy.com/install | PLATFORM=docker bash -
 </div></div>
 </pre>
<br />
<br />
#### Option 3: Install `meshplayctl` binary and deploy Meshplay on Kubernetes

 <pre class="codeblock-pre">
 <div class="codeblock">
 <div class="clipboardjs">
  $ curl -L https://meshplay.khulnasofy.com/install | PLATFORM=kubernetes bash -
 </div></div>
 </pre>
<br />
<br />
#### Option 4: Install `meshplayctl` binary and Meshplay adapter(s)

Install `meshplayctl` binary and include one or more [adapters]({{ site.baseurl }}/concepts/architecture/adapters) to be deployed

 <pre class="codeblock-pre">
 <div class="codeblock">
 <div class="clipboardjs">
  $ curl -L https://meshplay.khulnasofy.com/install | ADAPTERS=consul PLATFORM=kubernetes bash -
 </div></div>
 </pre>
<br />
<br />
### Start Meshplay
You are ready to deploy Meshplay `meshplayctl`. To do so, execute the following command.

 <pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">meshplayctl system start</div></div>
 </pre>
