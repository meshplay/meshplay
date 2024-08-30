---
layout: default
title: Overview
abstract: An overview of Meshplay concepts and its functionality.
permalink: /getting-started/overview
language: en
redirect_to: /project/overview
---
Meshplay is an extensible engineering platform for the collaborative design and operation of cloud and cloud native infrastructure and applications.

Kubernetes-centric. Kubernetes not required.

## Meshplay is for all cloud and cloud native infrastructure

Infrastructure diversity is a reality for any enterprise. Whether you're running a single Kubernetes cluster or multiple Kubernetes clusters, on one cloud or multiple clouds, you'll find that Meshplay supports your infrastructure diversity (or lack thereof).

## Meshplay's Functionality

Meshplay supports all Kubernetes-based infrastructure including most cloud services of AWS and GCP platforms. Meshplay features can be categorized by:

1. Performance Management
   - Workload and performance characterization with both built-in and external load generators
   - Prometheus and Grafana integration
1. Lifecycle Management (Day 0, Day 1)
   - Cloud and cloud native provisioning
   - Discovery and onboarding of existing environments and workloads
1. Configuration Management (Day 2)
   - Cloud native patterns catalog
   - Configuration best practices
   - Policy engine for relationship inference and context-aware design
1. Collaboration
   - Multi-player infrastructure design and operation
1. Data Plane Intelligence
   - Registry and configuration of WebAssembly filters for Envoy
1. Interoperability and Federation
   - Integration with thousands of cloud services and cloud native projects
   - Manage multiple service meshes concurrently
   - Connect to multiple clusters independently

### Meshplay is for Developers, Operators, and Product Owners

Whether making a Day 0 adoption choice or maintaining a Day 2 deployment, Meshplay has useful capabilities in either circumstance. Targeted audience for Meshplay project would be any technology operators that leverage service mesh in their ecosystem; this includes developers, devops engineers, decision makers, architects, and organizations that rely on microservices platform.

### Meshplay is for cloud native patterns

Through [Models]({{site.baseurl}}/concepts/logical/models), Meshplay describes infrastructure under management, enabling you to define cloud native designs and patterns and then to export those designs and share within the <a href="https://meshplay.khulnasofy.com/catalog" target="_self_">Meshplay Catalog</a>.

### Meshplay is for performance management

Meshplay helps users weigh the value of their cloud native deployments against the overhead incurred in running different deployment scenarios and different configruations. Meshplay provides statistical analysis of the request latency and throughput seen across various permutations of your workload, infrastructure and infrastructure configuration. In addition to request latency and throughput, Meshplay also tracks memory and CPU overhead in of the nodes in your cluster. Establish a performance benchmark and track performance against this baseline as your environment changes over time.


