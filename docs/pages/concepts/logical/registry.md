---
layout: default
title: Registry
permalink: concepts/logical/registry
type: concepts
abstract: Meshplay Registry is a database acting as the central repository for all capabilities known to Meshplay. These capabilities encompass various entities, including models, components, relationships, and policies.
language: en
display-title: "false"
list: include

---
# Meshplay Registry: A Central Hub for Capabilities

The Meshplay Registry is a vital component within Meshplay, serving as a centralized repository for managing a diverse range of cloud and cloud native resources. It stores and organizes crucial information such as models, categories, components, and relationships, enabling efficient interaction and utilization of these resources within the Meshplay ecosystem. You can conveniently [access and manage registry data](#interacting-with-the-meshplay-registry) through Meshplay UI, and through Meshplay CLI ([meshplayctl registry]({{site.baseurl}}/reference/meshplayctl/#meshplay-registry-management)).

As the central repository for all capabilities known to Meshplay, contains various entities.

<details>
  <summary>Contents of the Registry</summary>
  <br /><br />
  <a href="../models">Models</a>: Blueprints defining configurations for interacting with cloud-native infrastructure. They consist of operations, components, relationships, and policies.
  <ul>
    <li><a href="../components">Components</a>: Reusable building blocks for depicting capabilities defined within models.</li>
    <li><a href="../relationships">Relationships</a>: Define the nature of connections between components within a model, describing how they interact and depend on each other.</li>
    <li><a href="../policies">Policies</a>: Enforce specific rules and governance for system behavior under Meshplay's management.</li>
    <li><a href="../connections">Connections</a>: Managed and unmanaged resources that Meshplay can interact with.</li>
    <li><a href="../credentials">Credentials</a>: Optionally, included secrets associated with connections contained in a model.</li>
  </ul>
  <br />
</details>

## Key Concepts and Terminology

- **Registry**: a component within Meshplay that contains a database of known capabilities.
- **Registrar**: The internal Meshplay Server process responsible for managing and maintaining the registry.
- **Registrant** (Entity Source): The source of an entity (e.g., model file, Kubernetes cluster).
Entity (Registree): An individual capability stored in the registry (e.g., model, component).
- **Registrant** *(Host)*: A Meshplay Connection responsible for sourcing and registering entities. A registrant can perform registration for their own entities, or a registrant can act as a proxy on behalf of a third-party entity source.
- **Entity** *(registree)* - an entry in the Meshplay Registry; e.g. a model, component, relationship, or policy. Sometimes referred to as a capability.
<!-- - **Entity Source**: an entity’s original location from which it was sourced; e.g. (source_uri is used as the flag by Meshplay Server to assess whether additional support). The Entity Source should have all the information that Meshplay needs to generate the components.   -->

## Models in the Registry

You will find two types of models in the registry: Static and Dynamic.

- **Static Models:** Pre-defined models included with each Meshplay release. See the full list of static models.
- **Dynamic Models:** Generated at run-time by connecting Meshplay to supported platforms like Kubernetes clusters or cloud providers.

Each Meshplay release comes with a built-in set of models automatically registered at Meshplay Server boot-time. These built-in models offer a core set of entities for Meshplay's supported [integrations](/extensibility/integrations). Once Meshplay Server is running, and as it connects to and discovers your infrastructure, *dynamic models* are automatically generated. A given Meshplay release may not include all possible models found in your environment, so Meshplay automatically generates *and registers* new models and components based on the specific infrastructure Meshplay is connected to. Dyanmic models often lack additional metadata, such as descriptions, tags, and relationships, which are typically included in static models.

## Interacting with the Meshplay Registry

Use either Meshplay UI or CLI to interact with the Registry. Meshplay UI offers a user-friendly visual interface for browsing, searching, and managing registry entries. You can easily explore available models, components, and relationships, gaining insights into their properties and connections. Meshplay CLI offers commands so that you can register, list, retrieve, update, and delete models, components, and relationships directly from the command line.

### Model Generation

The process of generating a Model (and its entities) is a multi-step process and does not require use of Meshplay Server. The process begins with the sourcing of the model information from an authoratitive source: a Registrant. Registrants are responsible for providing all the necessary information to Meshplay to generate the model.

#### Using Meshplay CLI to Generate Models

Meshplay CLI supports the generation of models from a Google Spreadsheet. The Google Spreadsheet should contain a list of model names and source locations from any supported Registrant (e.g. GitHub, Artifact Hub) repositories. The source locations can be a URL to a folder containing Kubernetes CRDs, or to a Helm Chart tar.gz, or an individual Kubernetes Manifest with custom resource definition.

See [`meshplayctl registry generate`](/reference/meshplayctl/registry/generate) for more information.

### Model Registration

Once registered in the Meshplay Registry, Models and their entities are available for use within that specific Meshplay Server.

Meshplay [Adapters]({{ site.baseurl }}/concepts/architecture/adapters) are one example of a Registrant. Registrants are responsible for the registration of entities in the Meshplay Registry. Adapters are responsible for the sourcing and registration of entities and the packaging of these enties into one or more models.

#### Using Meshplay CLI to Register a Model

```bash
meshplayctl model import -f <path-to-model>
```

#### Using Meshplay UI to Register a Model

Visit the Settings --> Registry page and click the "Import" button to import a model.

### Ignoring an Entity

You have control over whether a registered entity (model and all that the model contains) this can be an individual or team-level preference. Use the "Ignore" action to designate whether a given model is allowed to be used within a given Meshplay Server deployment. Models that are ignored remain in the Meshplay Registry but are not available for use within a given Meshplay Server deployment.

