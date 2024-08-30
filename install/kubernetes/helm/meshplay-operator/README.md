# meshplay-operator

![Version: 0.6.0](https://img.shields.io/badge/Version-0.6.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square)

Meshplay Operator chart.

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| KhulnaSoft Authors | <community@khulnasoft.com> |  |
| aisuko | <urakiny@gmail.com> |  |
| leecalcote | <leecalcote@gmail.com> |  |

## Requirements

| Repository | Name | Version |
|------------|------|---------|
|  | meshplay-broker | 0.5.0 |
|  | meshplay-meshsync | 0.5.0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| annotations | object | `{}` |  |
| env | object | `{}` |  |
| fullnameOverride | string | `"meshplay-operator"` |  |
| imagePullSecrets | list | `[]` |  |
| ingress.annotations | object | `{}` |  |
| ingress.enabled | bool | `false` |  |
| ingress.hosts[0].host | string | `"chart-example.local"` |  |
| ingress.hosts[0].paths | list | `[]` |  |
| ingress.tls | list | `[]` |  |
| kubeRbac.args[0] | string | `"--secure-listen-address=0.0.0.0:8443"` |  |
| kubeRbac.args[1] | string | `"--upstream=http://127.0.0.1:8080/"` |  |
| kubeRbac.args[2] | string | `"--logtostderr=false"` |  |
| kubeRbac.args[3] | string | `"--v=10"` |  |
| kubeRbac.image.pullPolicy | string | `"Always"` |  |
| kubeRbac.image.repository | string | `"gcr.io/kubebuilder/kube-rbac-proxy:v0.5.0"` |  |
| kubeRbac.name | string | `"kube-rbac-proxy"` |  |
| meshplay-broker.enabled | bool | `true` |  |
| meshplay-broker.fullnameOverride | string | `"meshplay-broker"` |  |
| meshplay-broker.serviceAccountNameOverride | string | `"meshplay-server"` |  |
| meshplay-meshsync.enabled | bool | `true` |  |
| meshplay-meshsync.fullnameOverride | string | `"meshplay-meshsync"` |  |
| meshplay-meshsync.serviceAccountNameOverride | string | `"meshplay-server"` |  |
| meshplayOperator.args[0] | string | `"--metrics-addr=127.0.0.1:8080"` |  |
| meshplayOperator.args[1] | string | `"--enable-leader-election"` |  |
| meshplayOperator.command[0] | string | `"/manager"` |  |
| meshplayOperator.image.pullPolicy | string | `"Always"` |  |
| meshplayOperator.image.repository | string | `"khulnasoft/meshplay-operator:stable-latest"` |  |
| meshplayOperator.name | string | `"manager"` |  |
| nameOverride | string | `""` |  |
| nodeSelector | object | `{}` |  |
| podSecurityContext | object | `{}` |  |
| probe.livenessProbe.enabled | bool | `false` |  |
| probe.readinessProbe.enabled | bool | `false` |  |
| replicaCount | int | `1` |  |
| resources | object | `{}` |  |
| securityContext | object | `{}` |  |
| service.annotations | object | `{}` |  |
| service.port | int | `10000` |  |
| service.type | string | `"ClusterIP"` |  |
| serviceAccount.create | string | `"create"` |  |
| serviceAccount.name | string | `"meshplay-operator"` |  |
| testCase.enabled | bool | `false` |  |
| tolerations | list | `[]` |  |

