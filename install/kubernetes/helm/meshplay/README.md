# meshplay

![Version: 0.6.0](https://img.shields.io/badge/Version-0.6.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square)

Meshplay chart for deploying Meshplay

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| Meshplay Authors | <maintainers@meshplay.io> |  |

## Values

| Key                                             | Type | Default                                                                                                                                                                                               | Description |
|-------------------------------------------------|------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|
| affinity                                        | object | `{}`                                                                                                                                                                                                  |  |
| annotations                                     | object | `{}`                                                                                                                                                                                                  |  |
| env.ADAPTER_URLS                                | string | `"meshplay-istio:10000 meshplay-linkerd:10001 meshplay-consul:10002 meshplay-kuma:10007 meshplay-nginx-sm:10010 meshplay-nsm:10004 meshplay-app-mesh:10005 meshplay-traefik-mesh:10006 meshplay-cilium:10012"` | Optionally, pre-configure Meshplay Server with the set of Meshplay Adapters used in the deployment. |
| env.EVENT                                       | string | `"meshplayLocal"`                                                                                                                                                                                      |  |
| env.PROVIDER                                    | string | `"Local"`                                                                                                                                                                                             | Use this security-related setting to enforce selection of one and only one Provider. In this way, your Meshplay deployment will only trust and only allow users to authenticate using the Provider you have configured in this setting. See the [Remote Provider documentation](https://docs.meshplay.io/extensibility/providers) for a description of what a Provider is.  |
| env.MESHPLAY_SERVER_CALLBACK_URL                 | string | `""`                                                                                                                                                                                                  | Configure an OAuth callback URL for Meshplay Server to use when signing into a Remote Provider and your Meshplay Server instance is not directly reachable by that Remote Provider. See the [Remote Provider documentation](https://docs.meshplay.io/extensibility/providers#configurable-oauth-callback-url) for more details. |
| env.PROVIDER_BASE_URLS                          | string | `"https://meshplay.layer5.io"`                                                                                                                                                                         | Configure your Remote Provider of choice. See the [Remote Provider documentation](https://docs.meshplay.io/extensibility/providers) for a description of what a Provider is. |
| fullnameOverride                                | string | `""`                                                                                                                                                                                                  |  |
| image.pullPolicy                                | string | `"Always"`                                                                                                                                                                                            |  |
| image.repository                                | string | `"layer5/meshplay"`                                                                                                                                                                                    |  |
| image.tag                                       | string | `"stable-latest"`                                                                                                                                                                                     |  |
| imagePullSecrets                                | list | `[]`                                                                                                                                                                                                  |  |
| ingress.annotations                             | object | `{}`                                                                                                                                                                                                  |  |
| ingress.enabled                                 | bool | `false`                                                                                                                                                                                               |  |
| ingress.hosts[0].host                           | string | `"chart-example.local"`                                                                                                                                                                               |  |
| ingress.hosts[0].paths                          | list | `[]`                                                                                                                                                                                                  |  |
| ingress.tls                                     | list | `[]`                                                                                                                                                                                                  |  |
| meshplay-app-mesh.enabled                        | bool | `false`                                                                                                                                                                                               | Enable to deploy this Meshplay Adapter upon initial deployment. Meshplay Adapters can be deployed post-installation using either Meshplay CLI or UI. |
| meshplay-app-mesh.fullnameOverride               | string | `"meshplay-app-mesh"`                                                                                                                                                                                  |  |
| meshplay-app-mesh.serviceAccountNameOverride     | string | `"meshplay-server"`                                                                                                                                                                                    |  |
| meshplay-cilium.enabled                          | bool | `false`                                                                                                                                                                                               | Enable to deploy this Meshplay Adapter upon initial deployment. Meshplay Adapters can be deployed post-installation using either Meshplay CLI or UI. |
| meshplay-cilium.fullnameOverride                 | string | `"meshplay-cilium"`                                                                                                                                                                                    |  |
| meshplay-consul.enabled                          | bool | `false`                                                                                                                                                                                               | Enable to deploy this Meshplay Adapter upon initial deployment. Meshplay Adapters can be deployed post-installation using either Meshplay CLI or UI. |
| meshplay-consul.fullnameOverride                 | string | `"meshplay-consul"`                                                                                                                                                                                    |  |
| meshplay-consul.serviceAccountNameOverride       | string | `"meshplay-server"`                                                                                                                                                                                    |  |
| meshplay-istio.enabled                           | bool | `false`                                                                                                                                                                                               | Enable to deploy this Meshplay Adapter upon initial deployment. Meshplay Adapters can be deployed post-installation using either Meshplay CLI or UI. |
| meshplay-istio.fullnameOverride                  | string | `"meshplay-istio"`                                                                                                                                                                                     |  |
| meshplay-istio.serviceAccountNameOverride        | string | `"meshplay-server"`                                                                                                                                                                                    |  |
| meshplay-kuma.enabled                            | bool | `false`                                                                                                                                                                                               | Enable to deploy this Meshplay Adapter upon initial deployment. Meshplay Adapters can be deployed post-installation using either Meshplay CLI or UI. |
| meshplay-kuma.fullnameOverride                   | string | `"meshplay-kuma"`                                                                                                                                                                                      |  |
| meshplay-kuma.serviceAccountNameOverride         | string | `"meshplay-server"`                                                                                                                                                                                    |  |
| meshplay-linkerd.enabled                         | bool | `false`                                                                                                                                                                                               | Enable to deploy this Meshplay Adapter upon initial deployment. Meshplay Adapters can be deployed post-installation using either Meshplay CLI or UI. |
| meshplay-linkerd.fullnameOverride                | string | `"meshplay-linkerd"`                                                                                                                                                                                   |  |
| meshplay-linkerd.serviceAccountNameOverride      | string | `"meshplay-server"`                                                                                                                                                                                    |  |
| meshplay-nginx-sm.enabled                        | bool | `false`                                                                                                                                                                                               | Enable to deploy this Meshplay Adapter upon initial deployment. Meshplay Adapters can be deployed post-installation using either Meshplay CLI or UI. |
| meshplay-nginx-sm.fullnameOverride               | string | `"meshplay-nginx-sm"`                                                                                                                                                                                  |  |
| meshplay-nginx-sm.serviceAccountNameOverride     | string | `"meshplay-server"`                                                                                                                                                                                    |  |
| meshplay-nsm.enabled                             | bool | `false`                                                                                                                                                                                               | Enable to deploy this Meshplay Adapter upon initial deployment. Meshplay Adapters can be deployed post-installation using either Meshplay CLI or UI. |
| meshplay-nsm.fullnameOverride                    | string | `"meshplay-nsm"`                                                                                                                                                                                       |  |
| meshplay-nsm.serviceAccountNameOverride          | string | `"meshplay-server"`                                                                                                                                                                                    |  |
| meshplay-nighthawk.enabled                       | bool | `false`                                                                                                                                                                                               | Enable to deploy this Meshplay Adapter upon initial deployment. Meshplay Adapters can be deployed post-installation using either Meshplay CLI or UI. |
| meshplay-nighthawk.fullnameOverride              | string | `"meshplay-nighthawk"`                                                                                                                                                                                 |  |
| meshplay-nighthawk.serviceAccountNameOverride    | string | `"meshplay-nighthawk"`                                                                                                                                                                                    |  |
| meshplay-operator.enabled                        | bool | `true`                                                                                                                                                                                                | Enable to deploy this Meshplay Operator upon initial deploymeent. Meshplay Operator can be deployed post-installation using Meshplay UI. |
| meshplay-operator.fullnameOverride               | string | `"meshplay-operator"`                                                                                                                                                                                  |  |
| meshplay-osm.enabled                             | bool | `false`                                                                                                                                                                                               | OSM is an archived project. |
| meshplay-osm.fullnameOverride                    | string | `"meshplay-osm"`                                                                                                                                                                                       |  |
| meshplay-osm.serviceAccountNameOverride          | string | `"meshplay-server"`                                                                                                                                                                                    |  |
| meshplay-traefik-mesh.enabled                    | bool | `false`                                                                                                                                                                                               | Enable to deploy this Meshplay Adapter upon initial deployment. Meshplay Adapters can be deployed post-installation using either Meshplay CLI or UI. |
| meshplay-traefik-mesh.fullnameOverride           | string | `"meshplay-traefik-mesh"`                                                                                                                                                                              |  |
| meshplay-traefik-mesh.serviceAccountNameOverride | string | `"meshplay-server"`                                                                                                                                                                                    |  |
| meshplaygateway.enabled                          | bool | `false`                                                                                                                                                                                               |  |
| meshplaygateway.selector.istio                   | string | `"ingressgateway"`                                                                                                                                                                                    |  |
| metadata.name                                   | string | `"meshplay"`                                                                                                                                                                                           |  |
| metadata.namespace                              | string | `"meshplay"`                                                                                                                                                                                           |  |
| nameOverride                                    | string | `""`                                                                                                                                                                                                  |  |
| nodeSelector                                    | object | `{}`                                                                                                                                                                                                  |  |
| podSecurityContext                              | object | `{}`                                                                                                                                                                                                  |  |
| probe.livenessProbe.enabled                     | bool | `false`                                                                                                                                                                                               |  |
| probe.readinessProbe.enabled                    | bool | `false`                                                                                                                                                                                               |  |
| rbac.nodes                                      | bool | `false`                                                                                                                                                                                               |  |
| replicaCount                                    | int | `1`                                                                                                                                                                                                   |  |
| resources                                       | object | `{}`                                                                                                                                                                                                  |  |
| restartPolicy                                   | string | `"Always"`                                                                                                                                                                                            |  |
| securityContext                                 | object | `{}`                                                                                                                                                                                                  |  |
| service.annotations                             | object | `{}`                                                                                                                                                                                                  |  |
| service.port                                    | int | `9081`                                                                                                                                                                                                |  |
| service.target_port                             | int | `8080`                                                                                                                                                                                                |  |
| service.type                                    | string | `"LoadBalancer"`                                                                                                                                                                                      |  |
| serviceAccount.name                             | string | `"meshplay-server"`                                                                                                                                                                                    |  |
| testCase.enabled                                | bool | `false`                                                                                                                                                                                               |  |
| tolerations                                     | list | `[]`                                                                                                                                                                                                  |  |

## Setup Repo Info

```console
helm repo add meshplay meshplay https://meshplay.io/charts/
helm repo update
```

_See [helm repo](https://helm.sh/docs/helm/helm_repo/) for command documentation._

## Installing the Chart

To install the chart with the release name `meshplay`:

```console
kubectl create namespace meshplay
helm install meshplay meshplay/meshplay
```

## Uninstalling the Chart

To uninstall/delete the `meshplay` deployment:

```console
helm delete meshplay
```

## Installing the Chart with a custom namespace

```console
kubectl create namespace meshplay
helm install meshplay meshplay/meshplay --namespace meshplay
```

## Installing the Chart with a custom Meshplay Adapters

Eg: For [Meshplay Adapter for Istio](https://github.com/meshplay/meshplay-istio)
```console
kubectl create namespace meshplay
helm install meshplay meshplay/meshplay --set meshplay-istio.enabled=true
```
