---
layout: default
title: Meshplayctl system commands
permalink: guides/meshplayctl/system-commands
language: en
type: guides
category: meshplayctl
abstract: Meshplayctl system commands for managing Meshplay deployments.
---

Let's get familiar with meshplayctl system commands. The syntax of the meshplayctl commands goes as follws : `meshplayctl <Main_command> <Argument> <Flags>`

## Main_command : system
### start 
`meshplayctl system start` : This will initiate Meshplay & automatically open it in your default web browser.

<a href="{{ site.baseurl }}/assets/img/syscmd/start.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/start.png" /></a>

`meshplayctl system start --skip-browser` : It skips opening Meshplay in your browser with the provided URL.

<a href="{{ site.baseurl }}/assets/img/syscmd/skipbrowser.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/skipbrowser.png" /></a>

`meshplayctl system start --skip-update` : It is used when you want to skip updating Meshplay if an update is available.

<a href="{{ site.baseurl }}/assets/img/syscmd/system update.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/system update.png" /></a>

`meshplayctl system start --reset` : It resets your Meshplay configuration file to its default configuration.

`meshplayctl system start --platform string` : It allows you specify a platform for deploying Meshplay.

<a href="{{ site.baseurl }}/assets/img/syscmd/platform.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/platform.png" /></a>


### stop 
`meshplayctl system stop` : It stops Meshplay resources & delete its associated namespaces.

<a href="{{ site.baseurl }}/assets/img/syscmd/stop.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/stop.png" /></a>

`meshplayctl system stop --reset` : It stops Meshplay and resets the Meshplay configuration file to its default configuration.

<a href="{{ site.baseurl }}/assets/img/syscmd/stop reset.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/stop reset.png" /></a>

`meshplayctl system stop --keep-namespace` : It stops Meshplay without deleting the associated namespaces.

<a href="{{ site.baseurl }}/assets/img/syscmd/keep namespace.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/keep namespace.png" /></a>

`meshplayctl system stop --force` : Force stops Meshplay instead of gentle way. This is only used in emergency situations when `meshplayctl system stop` can't halt Meshplay.

<a href="{{ site.baseurl }}/assets/img/syscmd/force stop.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/force stop.png" /></a>

### update
`meshplayctl system update` : This updates Meshplay itself, not the meshplayctl. Ensure Meshplay is running when using this.

<a href="{{ site.baseurl }}/assets/img/syscmd/system update.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/system update.png" /></a>

`meshplayctl system update --skip-reset` : Skips the check for a new manifest file.

<a href="{{ site.baseurl }}/assets/img/syscmd/update skip reset.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/update skip reset.png" /></a>

### reset
`meshplayctl system reset` : Resets Meshplay to its default configuration.

<a href="{{ site.baseurl }}/assets/img/syscmd/reset.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/reset.png" /></a>

### restart 
`meshryctl system restart` : Stops Meshplay and then starts it again. Opens the website in your default browser.

<a href="{{ site.baseurl }}/assets/img/syscmd/restart.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/restart.png" /></a>

### status 
`meshplayctl system status` : Displays the status of Meshplay components.

`meshplayctl system status --verbose` : Provides additional data along with Meshplay and its component status.

<a href="{{ site.baseurl }}/assets/img/syscmd/system status.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/system status.png" /></a>


### dashboard
`meshplayctl system dashboard` : Opens the Meshplay dashboard in your default browser.

<a href="{{ site.baseurl }}/assets/img/syscmd/system dahboard.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/system dahboard.png" /></a>

`meshplayctl system dashboard --skip-browser` : Provides the link to the dashboard, allowing you to open it in any browser.

<a href="{{ site.baseurl }}/assets/img/syscmd/dashboard skip.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/dashboard skip.png" /></a>

`meshplayctl system dashboard --port-forward` : If the current port is busy, it opens the dashboard on another port.

<a href="{{ site.baseurl }}/assets/img/syscmd/portforward.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/portforward.png" /></a>


### login 
`meshplayctl system login` : Authenticates you with your selected provider.

<a href="{{ site.baseurl }}/assets/img/syscmd/system login.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/system login.png" /></a>

### check
`meshplayctl system check` : Performs checks for both pre & post mesh deployment scenarios on Meshplay.

<a href="{{ site.baseurl }}/assets/img/syscmd/system check.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/system check.png" /></a>

`meshplayctl system check --preflight` : Runs pre-deployment checks.

`meshplayctl system check --adapter` : Runs checks for a specific Mesh adapter.

`meshplayctl system check --adapters` : Runs checks for Meshplay adapters

`meshplayctl system check --components` : Runs checks for Meshplay components

`meshplayctl system check --operator` : Runs checks for Meshplay Operator

## Main_command : system channel
### channel
`mesheyctl system channel set [stable|stable-version|edge|edge-version]` : Used to set the channel.

`mesheyctl system channel switch [stable|stable-version|edge|edge-version]` : Used to switch between channels.

<a href="{{ site.baseurl }}/assets/img/syscmd/channel set.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/channel set.png" /></a>

`mesheyctl system channel view --all` : Displays all available channels.

`mesheyctl system channel view` : Displays the current channel.

<a href="{{ site.baseurl }}/assets/img/syscmd/channel view.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/channel view.png" /></a>


## Main_command : system context
### create 
`meshplayctl system context create 'context-name'` : Creates a new context with default parameters.

<a href="{{ site.baseurl }}/assets/img/syscmd/context create.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/context create.png" /></a>

`meshplayctl system context create --component stringArray` : Specifies the component to be created in the context.

`meshplayctl system context create --platform string` : Specifies the platform.

`meshplayctl system context create --set` : Sets this  context as default context.

`meshplayctl system context create --url string` : Specifies the target URL.

<a href="{{ site.baseurl }}/assets/img/syscmd/context flag.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/context flag.png" /></a>


###  switch
`meshplayctl system context switch` : Easily switch between different contexts.

###  list
`meshplayctl system context list` : Lists all your available Meshplay contexts.

<a href="{{ site.baseurl }}/assets/img/syscmd/context list.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/context list.png" /></a>

###  delete
`meshplayctl system context delete` : Delete context.

<a href="{{ site.baseurl }}/assets/img/syscmd/context delete.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/context delete.png" /></a>


###  view
`meshplayctl system context view` : Display all your contexts with additional information.

<a href="{{ site.baseurl }}/assets/img/syscmd/context view.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/context view.png" /></a>


## Main_command : system provider
### switch
`meshplayctl system provider switch` : Changes your provider

<a href="{{ site.baseurl }}/assets/img/syscmd/pro switch.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/pro switch.png" /></a>

### list
`meshplayctl system provider list` : Lists all available providers

<a href="{{ site.baseurl }}/assets/img/syscmd/pro list.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/pro list.png" /></a>

### set
`meshplayctl system provider set` : Set your provider

<a href="{{ site.baseurl }}/assets/img/syscmd/pro set.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/pro set.png" /></a>

### view
`meshplayctl system provider view` : Lists your current context and provider

<a href="{{ site.baseurl }}/assets/img/syscmd/pro view.png"><img alt="skip-browser" style="width:500px;height:auto;" src="{{ site.baseurl }}/assets/img/syscmd/pro view.png" /></a>

