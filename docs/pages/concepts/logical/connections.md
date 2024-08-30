---
layout: default
title: Connections
permalink: concepts/logical/connections
type: concepts
abstract: "Meshplay Connections are managed and unmanaged resources that either through discovery or manual entry are managed by a state machine and used within one or more Environments."
language: en
list: include
redirect_from:
- concepts/connections
---
Meshplay Connections are managed and unmanaged resources that either through discovery or manual entry are tracked by Meshplay. Connections can be assigned as resources to an Environment. 

{% include alert.html type="info" title="Connections as resources in Environments" content="Meshplay Environments allow you to logically group related <a href='/concepts/logical/connections'>Connections</a> and their associated <a href='/concepts/logical/credentials'>Credentials</a>. Environments make it easier for you to manage, share, and work with a collection of resources as a group, instead of dealing with all your Connections and Credentials on an individual basis." %}

{% include alert.html type="dark" title="Managed vs Unmanaged Connections" content="Managed Connections are those that are discovered by MeshSync and are managed by Meshplay. Unmanaged Connections are those that are manually added by the user and are not managed by Meshplay." %}

## States and the Lifecycle of Connections

Meshplay tracks the status of each connections throughout the connection's lifecycle. Meshplay is intentional about the currently assigned state and which state a connection may or may not transition to and from. To better understand connection states and their meaning, let's consider an example in which you a `Kubernetes` cluster with `Prometheus` installed.

![]({{site.baseurl}}/assets/img/lifecycle-management/states-for-kubernetes-cluster-connections.svg)

### State: Discovered

All resources discovered by [MeshSync's](meshsync.md) multi-tier discovery or provided as part of config, and if Meshplay can integrate, a connection with state as `Discovered` will be created. Though, the connection/resources are not tested for its reachability/usability i.e. Meshplay has not made an attempt to connect or manage the connection.

When a connection has been discovered, it will be listed in the MeshSync browser / Connections table in Meshplay UI. You can self transition a particular connection to [Register](#state-registered) / [Ignore](#state-ignored) state.

> Example: MeshSync discovers Prometheus components and inform Meshplay Server about available Prometheus connection, but Meshplay is yet to [connect](#state-connected) and start scraping metrics.

### State: Registered

The connection in this state have been verified for its use and reachability but not yet being used. Almost all reachable connections will auto transition to Registered state from [Discovered](#state-discovered) state and it is upto the user what to do with this connection (i.e. User needs to administratively process the connection). It can be transitioned to [Connected](#state-connected), [Maintenance](#state-maintenance) and [Not Found](#state-not-found).

> EExampleg: User manually selects the registered Prometheus connection and transition to the [connected](#state-connected) state (i.e. User administratively processes the connection).

### State: Connected

The connection in this state is administratively processed and being actively managed by Meshplay. User can interface and invoke set of actions with the connection.</br>
From this state the transition can happen to either [Maintenance](#state-maintenance) or [Ignore](#state-ignored) state. </br> Auto transition to [Disconnected](#state-disconnected) state will occur if Meshplay can no longer communicate with the connection, which can occur due to connectivity issue/AuthN-AuthZ/connection was deleted outside Meshplay or any other issue.

> Example: Meshplay is communicating with Prometheus APIs to scrape metrics and present it in the UI.

_Certain connections can auto-transition to connected state._

### State: Ignored

The connection is administratively processed to be ignored from Meshplay's view of management. Meshplay will not re-discover this connection even when current user session gets expired.

> Example: Meshplay server will stop/not scrape metrics from Prometheus. Though, the previous data (if connected previously) will continue to exist and user needs to manually delete.

{% include alert.html type="info" title="Ignored versus Disconnected" content="You might intentionally choose to have Meshplay ignore a given Prometheus connection, explicitly leaving in Meshplay’s field of view, but identifying it as a connection not to manage. This is distinctly different than a Prometheus that Meshplay was managing, but has been turned off/uninstalled and now Meshplay is disconnected from the Prometheus." %}

### State: Maintenance

The connection is administratively processed to be offline for maintenance tasks. This is different from being [Disconnected](#state-disconnected)/[Ignored](#state-ignored).

### State: Disconnected

The connection was previously [discovered](#state-discovered)/[registered](#state-registered)/[connected](#state-connected) but is not available currently. This could happen due to connectivity issue/AuthN-AuthZ/connection was deleted outside meshplay/administratively disconnected.

> Example: Prometheus crashed/API token provided at time of registration is revoked.

{% include alert.html type="info" title="Disconnected vs Deleted" content="The connection was previously connected but is unreachable due to connectivity issue/AuthN-AuthZ/connection was **deleted outside Meshplay** i.e. Connection was deleted beyond the Meshplay's view of management." %}

### State: Deleted

The connection is administratively processed to be deleted and removed from Meshplay's view of management. All the available/collected data will also be deleted.

> Example: Prometheus metrics will no longer be accessible to you from the Meshplay UI.

### State: Not Found

User tried registering the connection **manually** but Meshplay could not connect to it or if the connection is unavailable now. User can delete the connection or try re-registering.

{% include alert.html type="info" title="Not Found vs Disconnected" content="You might attempt to transition to Connected state but the connection is unavaialble now due to being deleted/some other reason. This is distinctly different than a cluster with Prometheuses installed for `application monitoring` which was connected previously but is now unreachable from Meshplay's view of management due to change in API token/similar issue." %}

_Connections like **Registration of Meshplay server with remote provider** (and few other connection types) can self transtion to the valid states._

## Registering Connections with Remote Providers

To register a connection with a remote provider, you need to follow these steps:

1. Obtain the necessary credentials or access tokens from the remote provider.
2. Open the Meshplay UI and navigate to the Connections page.
3. Click on the "Add Connection" button.
4. Fill in the required information, such as the provider type, name, and credentials.
5. Click on the "Register" button to register the connection.

Once the connection is registered, Meshplay will verify its reachability and usability. If successful, the connection will transition to the "Registered" state. From there, you can choose to administratively process the connection and transition it to the "Connected" state.

Note that some connections, such as the registration of Meshplay server with remote providers, can self-transition to valid states.

For more information on the different states and the lifecycle of connections, refer to the documentation above.

![]({{site.baseurl}}/assets/img/architecture/meshplay-server-registration-with-remote-providers.svg)