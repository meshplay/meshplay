---
layout: page
title: Contributing to Meshplay Server
permalink: project/contributing/contributing-server
redirect_from: project/contributing/contributing-server/
abstract: How to contribute to Meshplay Server
language: en
type: project
category: contributing
list: include
---

As a new contributor, you’re going to want to familiarize with the project in order to resolve the issues in the best way. Installing and playing around with Meshplay will give you context for any issues that you might work on.

Once an issue has been addressed, you’ll need to test it as well. Ideally, these tests are run from the user’s perspective (someone running Meshplay in a container), not from a contributor’s perspective (someone running Meshplay as a locally-compiled service).

## Compiling and Running Meshplay server

To build and run Meshplay server from source:

1. Build the static assets for the UI by running

{% capture code_content %}make ui-setup
make ui-build{% endcapture %}
{% include code.html code=code_content %}

2. Build & run the server code by running

{% capture code_content %}make server{% endcapture %}
{% include code.html code=code_content %}

Any time changes are made to the Go code, you will have to stop the server and run the above command again.
Once the Meshplay server is up and running, you should be able to access Meshplay on your `localhost` on port `9081` at `http://localhost:9081`. One thing to note, you might NOT see the [Meshplay UI](#contributing-ui) until the UI code is built as well.
After running Meshplay server, you will need to select your **Cloud Provider** by navigating to `localhost:9081`. Only then you will be able to use the Meshplay UI on port `3000`.

**Please note**: If you get error while starting the server as **"Meshplay Development Incompatible"** then follow the below guideline 👇

<a href="{{ site.baseurl }}/assets/img/meshplay-development-incompatible-error.png">
  <img style= "max-width: 450px;" src="{{ site.baseurl }}/assets/img/meshplay-development-incompatible-error.png" />
</a>

Potential Solution:

- Go to your meshplay folder in your local-system where you’ve cloned it.
  Execute:

- `git remote add upstream https://github.com/meshplay/meshplay`
- `git fetch upstream`
- Restart the meshplay server
- Additionally, before restarting the server, if you like to pull the latest changes, you can do: `git pull upstream master`

#### Building Docker image

To build a Docker image of Meshplay, please ensure you have `Docker` installed to be able to build the image. Now, run the following command to build the Docker image:

{% capture code_content %}make docker{% endcapture %}
{% include code.html code=code_content %}

#### Define and validate errors

Every Golang-based component within the Meshplay ecosystem incorporates a utility to define and manage error messages for every error instance. This is internally done with several make commands, but one can explicitly validate with the help of the following make command. This checks and validates the errors that are present in the particular project.

{% capture code_content %}make error{% endcapture %}
{% include code.html code=code_content %}

For more details, <a href="{{ site.baseurl }}/project/contributing/contributing-error">Error Utility</a>

### Configuring Log levels at Runtime

The server log levels can be configured at runtime by changing the env variable `LOG_LEVEL` defined in file [`server-config.env`](https://github.com/meshplay/meshplay/blob/master/server/cmd/server-config.env). The configuration library (`viper`) watches for the env file, any change in the file content results in the `file_system` event to be emitted and the log level is updated accordingly.

**_Should there be any alterations to the location or name of the environment file, it will result in the inability to configure log levels during runtime. In the event of such modifications, it is essential to update the server to preserve proper functionality._**

```Available Meshplay Server log levels are:
 - Panic - 0
 - Fatal - 1
 - Error - 2
 - Warn  - 3
 - Info  - 4
 - Debug - 5
 - Trace level - 6
```

The default setting for the `LOG_LEVEL` is `4` (Info). However, if the `DEBUG` environmental variable is configured as `TRUE`, it supersedes the value set in the `LOG_LEVEL` environmental variable, and the logging level is then adjusted to `5`(Debug).

### Using custom Meshkit code for Meshplay server development

<ol>
  <li>
    <p>Checkout <strong>meshplay</strong> and <strong>meshkit</strong> repository in adjacent directories.</p>
    {% capture code_content %}
$ git clone https://github.com/meshplay/meshplay.git
$ git clone https://github.com/meshplay/meshkit.git
    {% endcapture %}
    {% include code.html code=code_content %}
  </li>
  <li>
    <p>In your <code>meshplay</code> go.mod, update the meshkit package.</p>
    {% capture code_content %}
github.com/layer5io/meshkit => ../meshkit
    {% endcapture %}
    {% include code.html code=code_content %}
    <p>Remember to remove this go.mod change when creating pull requests.</p>
  </li>
</ol>

