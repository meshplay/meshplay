---
layout: default
title: "Extensibility: Providers"
permalink: extensibility/providers
type: Extensibility
#redirect_from: architecture/adapters
abstract: "Meshplay uses providers to enrich the level of depth by which it manages cloud native infrastructure."
language: en
list: include
---

Meshplay offers _Providers_ as a point of extensibility. It has a built-in Local Provider, named **"None"** and supports _Remote Providers_ that are designed to be pluggable. _Remote Providers_ offer a point of extension to users / integrators to deliver enhanced functionality such as authentication and authorization, using Meshplay as a platform.

1. **Extensibility points offer clean separation of Meshplay's core functionality versus plugin functionality.**
   - [MeshMap]({{site.baseurl}}/extensions/meshmap) is an example of a feature to be delivered via Remote Provider.
1. **Remote Providers should be able to offer custom RBAC, custom UI components, and custom backend components**
   - Dynamically loadable frameworks need to be identified or created to serve each of these purposes.

### What functionality do Providers offer?

The functionalities offerred by a _Remote Provider_ may vary broadly between providers. Meshplay offers extension points that Remote Providers are able to use to inject various functionalities, functionalities that are specific to that provider.
Some examples include:

- **Authentication and Authorization**
  - Examples: session management, two factor authentication, LDAP integration.
- **Long-Term Persistence**
  - Examples: Storage and retrieval of performance test results.
  - Examples: Storage and retrieval of user preferences.
- **Enhanced Visualization**
  - Examples: Creation of a visual service mesh topology.
  - Examples: Different charts (metrics), debug (log viewer), distributed trace explorers.
- **Reporting**
  - Examples: Using Meshplay's GraphQL server to compose new dashboards.

<a href="{{ site.baseurl }}/assets/img/providers/provider_screenshot_new.png">
<img src="{{ site.baseurl }}/assets/img/providers/provider_screenshot_new.png" width="50%" /></a>

  <figcaption>Figure: Selecting a provider in Meshplay</figcaption>

## Types of providers

There are two types of providers defined in Meshplay, `local` and `remote`.

- The **Local** provider is built-into Meshplay.
- **Remote providers** can be implemented by anyone or an organization that wishes to integrate with Meshplay. Any number of Remote providers may be available in your Meshplay deployment.

### Remote Providers

The use of a Remote Provider, puts Meshplay into multi-user mode and requires user authentication. This provides security for the public-facing Meshplay UI as the remote provider enforces identity with authentication and authorization. You should also use a remote provider when your use of Meshplay is ongoing or used in a team environment (used by multiple people). This can be seen when using Meshplay Playground, where a user is prompted to login through the _Layer5 Meshplay Cloud_ remote provider. Visit [Meshplay Playground](https://playground.meshplay.io/) to experience this.

A specific remote provider can be enforced in a Meshplay instance by passing the name of the provider with the env variable `PROVIDER`.  

Name: **"Meshplay"** (default)

- Enforces user authentication.
- Long-term term persistence of test results.
- Save environment setup.
- Retrieve performance test results.
- Retrieve conformance test results.
- Free to use.

### Local Provider

The use of the Local Provider, **"None"**, puts Meshplay into a single-user mode and does not require authentication. Use the Local provider when your use of Meshplay is intended to be shortlived.

Name: **“None”**

- No user authentication.
- Container-local storage of test results. Ephemeral.
- Environment setup not saved.
- No performance test result history.
- No conformance test result history.
- Free to use.

### Design Principles: Meshplay Remote Provider Framework

Meshplay's Remote Provider extensibility framework is designed to enable the following functionalities:

1. **Pluggable UI Functionality**
   - Out-of-tree custom UI components with seamless user experience.
   - A system of remote retrieval of extension packages (ReactJS components and Golang binaries).

1. **Pluggable Backend Functionality**
   - Remote Providers have any number of capabilities unbeknownst to Meshplay.

1. **Pluggable AuthZ**
   - Design an extensible role based access control system such that Remote Providers can determine their own set of controls. Remote Providers to return JWTs with custom roles, permission keys and permission keychains.

## Building a Provider

Meshplay interfaces with providers through a Go interface. The Provider implementations have to be placed in the code and compiled together today. A Provider instance will have to be injected into Meshplay when the program starts.

Meshplay keeps the implementation of Remote Providers separate so that they are brought in through a separate process and injected into Meshplay at runtime (OR) change the way the code works to make the Providers invoke Meshplay.

### Remote Provider Extension Points

Interwoven into Meshplay’s web-based user interface are a variety of extension points. Each extension point is carefully carved out to afford a seamless user experience. Each extension point is identified by a name and type. The following Meshplay UI extension points are available:

|**Name**   |**Type**   | **Description**  |
|---|---|---|
|navigator |Menu Items   |This is supposed to be a full page extension which will get a dedicated endpoint in the meshplay UI. And will be listed in the meshplay UI’s navigator/sidebar. Menu items may refer to full page extensions.   |
|user_prefs   |Single Component   |This is supposed to be remote react components which will get placed in a pre-existing page and will not have a dedicated endpoint. As of now, the only place where this extension can be loaded is the “User Preference” section under meshplay settings.   |
|account   |Full Page   |Remote Reactjs components (or other) are placed in a pre-existing page and will have dedicated endpoint: `/extension/account`.   |
|collaborator |Single Component|This is supposed to be remote react components which will get placed in a pre-existing page and will not have a dedicated endpoint. Currently, placed at the Header component of Mehery UI. Its work is to show active Meshplay users under the same remote provider.|
  |/extension/\<your name here>|Full Page|The Provider package is unzipped into Meshplay server filesystem under `/app/provider-pkg/<package-name>`.|

Remote Providers must fulfill the following endpoints:

1. `/login` - return valid token
1. `/logout` - invalidating token
1. `/capabilities` - return capabilities.json

## UI Extension Points

All UI extensions will be hosted under the endpoint `<meshplayserver:port/provider>`.

### UserPrefs

The UserPrefs extension point expects and loads a component to be displayed into `/userpreferences` page.

### Navigator

The Navigator extension point loads a set of menu items to be displayed in the menu bar on the left hand side of the Meshplay UI.

## Capabilities Endpoint Example

Meshplay Server will proxy all requests to remote provider endpoints. Endpoints are dynamically determined and identified in the _**capabilities**_ section of the `/capabilities` endpoint. Providers as an object have the following attributes (this must be returned as a response to `/capabilities` endpoint):
<details>
<summary>Capabilities Endpoint Example</summary>
{% capture code_content %}{
  "provider_type": "remote",
  "package_version": "v0.1.0",
  "package_url": "https://layer5labs.github.io/meshplay-extensions-packages/provider.tar.gz",
  "provider_name": "Meshplay",
  "provider_description": [
    "Persistent sessions",
    "Save environment setup",
    "Retrieve performance test results",
    "Free use"
  ],
  "extensions": {
    "navigator": [
      {
        "title": "MeshMap",
        "href": {
          "uri": "/meshmap",
          "external": false
        },
        "component": "/provider/navigator/meshmap/index.js",
        "icon": "/provider/navigator/img/meshmap-icon.svg",
        "link:": true,
        "show": true,
        "children": [
          {
            "title": "View: Single Mesh",
            "href": {
              "uri": "/meshmap/mesh/all",
              "external": false
            },
            "component": "/provider/navigator/meshmap/index.js",
            "icon": "/provider/navigator/img/singlemesh-icon.svg",
            "link": false,
            "show": true
          }
        ]
      }
    ],
    "account": [
      {
          "title": "Overview",
          "on_click_callback": 1,
          "href": {
              "uri": "/account/overview",
              "external": false
          },
          "component": "/provider/account/profile/overview.js",
          "link": true,
          "show": true,
          "type": "full_page"
      },
      {
          "title": "Profile",
          "on_click_callback": 1,
          "href": {
              "uri": "/account/profile",
              "external": false
          },
          "component": "/provider/account/profile/profile.js",
          "link": true,
          "show": true,
          "type": "full_page"
      },
      {
          "title": "API Tokens",
          "on_click_callback": 1,
          "href": {
              "uri": "/account/tokens",
              "external": false
          },
          "component": "/provider/account/profile/tokens.js",
          "link": true,
          "show": true,
          "type": "full_page"
      }
    ],
    "user_prefs": [
      {
        "component": "/provider/userprefs/meshmap-preferences.js"
      }
    ],
    "collaborator": [
      {
        "component": "/provider/collaborator/avatar.js"
      }
    ]
  },
  "capabilities": [
    {
        "feature": "sync-prefs",
        "endpoint": "/user/preferences"
    },
    {
        "feature": "persist-results",
        "endpoint": "/results"
    },
    {
        "feature": "persist-result",
        "endpoint": "/result"
    },
    {
        "feature": "persist-smi-results",
        "endpoint": "/smi/results"
    },
    {
        "feature": "persist-smi-result",
        "endpoint": "/smi/result"
    },
    {
        "feature": "persist-metrics",
        "endpoint": "/result/metrics"
    },
    {
        "feature": "persist-smp-test-profile",
        "endpoint": "/user/test-config"
    },
    {
        "feature": "persist-performance-profiles",
        "endpoint": "/user/performance/profiles"
    },
    {
        "feature": "persist-schedules",
        "endpoint": "/user/schedules"
    },
    {
        "feature": "persist-meshplay-patterns",
        "endpoint": "/patterns"
    },
    {
        "feature": "persist-meshplay-filters",
        "endpoint": "/filters"
    },
    {
        "feature": "persist-meshplay-applications",
        "endpoint": "/applications"
    },
    {
        "feature": "persist-meshplay-pattern-resources",
        "endpoint": "/patterns/resource"
    },
    {
        "feature": "meshplay-patterns-catalog",
        "endpoint": "/patterns/catalog"
    },
    {
        "feature": "meshplay-filters-catalog",
        "endpoint": "/filters/catalog"
    },
    {
        "feature": "clone-meshplay-patterns",
        "endpoint": "/patterns/clone"
    },
    {
        "feature": "clone-meshplay-filters",
        "endpoint": "/filters/clone"
    },
    {
        "feature": "share-designs",
        "endpoint": "/api/content/design/share"
    },
    {
        "feature": "persist-connection",
        "endpoint": "/api/connection"
    }
  ]
}{% endcapture %}
{% include code.html code=code_content %}
</details>

#### Meshplay Server Registration

Every Meshplay server is capable of registering itself with the remote provider, considering that remote provider supports this feature as a capability.
On successful authentication with the remote provider, Meshplay server registers itself by sending a POST request to the remote provider through the `persist-connection` capability. The body of the request should include information so as to uniquely indentify Meshplay server and its status.

Example of the request body:

{% capture code_content %}
  {
    "server_id": "xxxx-xxxxx-xxxx-xxxx",
    "server_version": "vx.x.x",
    "server_build-sha": "xxxx-xxxxx",
    "server_location": "<protocol>://<hostname>:<port>”
  }
{% endcapture %}
{% include code.html code=code_content %}

## Configurable OAuth Callback URL

As of the release `v0.5.39`, Meshplay Server has extended its capability to configure the callback URL for connected _Remote Providers_. This is helpful when you are deploying Meshplay behind an ingress gateway or reverse proxy. You can specify a custom, redirect endpoint for the connected Remote Provider.
Let's understand this through an example deployment scenario below.

### Example Deployment Scenario: Meshplay Cloud, Istio, and Azure Kubernetes Service

User has deployed the Meshplay behind a Istio Ingress Gateway and the Istio is also behind an Application Gateway (e.g. AKS Application Gateway). Generally, when you use a GitHub Remote Provider for authentication, it redirect the request to the Istio Ingress Gateway FQDN. In this setup, redirection won't be successful because the Ingress Gateway is behind an additional Application Gateway. In this case, you have to define where the request should be redirected once it is authenticated from GitHub.

**_SOLUTION_**

You can define a custom callback URL by setting up the `MESHPLAY_SERVER_CALLBACK_URL` environment variable while installing Meshplay.

**Example using Helm**

First, construct the URL in the format
```https://[CUSTOM_URL]/api/user/token```

- Where `[CUSTOM_URL]` is the URL where **Meshplay server** will be or is installed. For example, `https://k8s-staging.test.io/`.
- And `api/user/token` is the **Auth Endpoint** and is append at the end of your custom URL.

So, the final URL would look similar to
```https://k8s-staging.test.io/api/user/token```

{% include alert.html type="info" title="Note" content="Callback URL is not the same as Provider URL. In scenarios where Meshplay server and Provider is installed on same server, pay attention to paths or subdomains." %}

Next, set the `MESHPLAY_SERVER_CALLBACK_URL` variable when running the `helm install`. Below is an example:

{% capture code_content %}helm install meshplay meshplay/meshplay --namespace meshplay --set env.MESHPLAY_SERVER_CALLBACK_URL=<https://k8s-staging.test.io/api/user/token>
{% endcapture %}
{% include code.html code=code_content %}

**NOTE**

If Meshplay server is accessible in a path of the URL such as `https://k8s-staging.test.io/meshplay`, then the callback URL will be `https://k8s-staging.test.io/meshplay/api/user/token`.

## Managing your Remote Provider Extension Code

Remote Provider extensions are kept out-of-tree from Meshplay (server and UI). You might need to build your extensions under the same environment and set of dependencies as Meshplay. The Meshplay framework of extensibility has been designed such that in-tree extensions can be safely avoided while still providing a robust platform from which to extend Meshplay’s functionality. Often, herein lies the delineation of open vs. closed functionality within Meshplay. Remote Providers can bring (plugin) what functionality that they want behind this extensible interface (more about Meshplay extensibility), at least that is up to the point that Meshplay has provided a way to plug that feature in.

Offering out-of-tree support for Meshplay extensions means that:

1. source code to your Meshplay extensions are not required to be open source,
1. liability to Meshplay’s stability is significantly reduced, avoiding potential bugs in extended components.

Through clearly defined extension points, Meshplay extensions may be offered as closed source capabilities that plug into open source Meshplay code. To facilitate integration of your Meshplay extensions, you might automate the building and releasing of your separate, but interdependent code repositories. You will be responsible for sustaining both your ReactJS and Golang-based extensions.
