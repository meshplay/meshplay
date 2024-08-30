---
layout: default
title: Policies
permalink: concepts/logical/policies
type: concepts
abstract: "Meshplay Policies enable you with a broad set of controls and governance of the behavior of systems under Meshplay's management."
language: en
list: include
---

Policies offer an evaluation algorithm to ensure desired behavior enforcement. Policies can be applied to components and relationships, defining rules and actions based on predefined conditions.

## Policy Evaluation

The relationships are a powerful way to design your infrastructure and each of them are backed by one or more policies. Policies evaluate the designs for potential relationships and the decide whether to create/delete/update the relationships.


[![Meshplay Models Policy Evaluation]({{ site.baseurl }}/assets/img/concepts/meshplay-models-policy-evaluation.svg
)]({{ site.baseurl }}/assets/img/concepts/meshplay-models-policy-evaluation.svg)

Meshplay Server has a built-in policy engine, based on Open Policy Agent (OPA). Currently, Meshplay Server is the only place where the policy evals occur. Policy evaluation is invoked each time a design is updated, and each time a design is imported. By default, policies evaluate for all registered relationships.

In any given Meshplay deployment, you can reference and search the full set of registered policies (in Meshplay's internal registry) in using either of Meshplay's client interfaces.

{% include alert.html type="info" title="Viewing All Registered Relationships" content='<p>You can view all registered relationships using either Meshplay UI or Meshplay CLI.</p><dl><dt>Using Meshplay UI...</dt><dd>Navigate to <i>Settings</i>, then to <i>Registry</i></dd>.<dt>Using Meshplay CLI...</dt><dd><code>meshplayctl policy list</code></dd></dl>' %}

<!-- There are different points in time in which policy evaluations are invoked

1. Each time the design is updated.
2. A Design/HelmChart/K8s Manifest/Docker Compose app is imported/uploaded.
3. Ad-hoc invocation from the Actions Center (coming soon). 
-->

### How are conflicts resolved?

In the event of a conflict or tie, Meshplay relies on Open Policy Agent's [reconciliation behavior](https://www.openpolicyagent.org/docs/latest/faq/) for conflict resolution.

{% include alert.html type="warning" title="Conflict Resolution" content="It may happen that certain eval decisions contain results such that two different components create a conflicting relationship with same component. While this is semantically correct, the visual representation of the relationship in such cases may be undesirable, and you may see relationships and components being redrawn depending upon how the client / Meshplay UI visualizes the relationships." %}

{% include alert.html type="dark" title="Future Feature" content="Policy evaluation in WASM runtime is on roadmap for Meshplay v0.8.3." %}
