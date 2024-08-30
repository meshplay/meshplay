---
layout: default
title: Troubleshooting Errors while running Meshplay
abstract: Troubleshooting Meshplay errors when running make run-fast / meshplay system start
permalink: guides/troubleshooting/meshplay-server
redirect_from: guides/troubleshooting/running
type: guides
category: troubleshooting
language: en
---

## meshplayctl system start

**Error:**

```
meshplayctl system start : : cannot start Meshplay: rendered manifests contain a resource that already exists.
Unable to continue with install: ServiceAccount "meshplay-operator" in namespace "meshplay" exists and cannot
be imported into the current release: invalid ownership metadata; label validation error: missing key
"app.kubernetes.io/managed-by": must be set to "Helm"; annotation validation error: missing key
"meta.helm.sh/release-name": must be set to "meshplay"; annotation validation error: missing key
"meta.helm.sh/release-namespace": must be set to "meshplay"
```

**(Fix) Clean the cluster using :**

 <pre class="codeblock-pre"><div class="codeblock">
 <div class="clipboardjs">
kubectl delete ns meshplay
kubectl delete clusterroles.rbac.authorization.k8s.io meshplay-controller-role meshplay-operator-role meshplay-proxy-role meshplay-metrics-reader
kubectl delete clusterrolebindings.rbac.authorization.k8s.io meshplay-controller-rolebinding meshplay-operator-rolebinding meshplay-proxy-rolebinding
 </div></div>
 </pre>

_Issue Reference : [https://github.com/meshplay/meshplay/issues/4578](https://github.com/meshplay/meshplay/issues/4578)_

### make server

**Error:**

```
FATA[0000] constraints not implemented on sqlite, consider using DisableForeignKeyConstraintWhenMigrating, more details https://github.com/go-gorm/gorm/wiki/GORM-V2-Release-Note-Draft#all-new-migrator
exit status 1
make: *** [Makefile:76: server] Error 1
```

**Fix:**

1. Flush the database by deleting the `.meshplay/config`
2. `make server`

#### See Also

- [Error Code Reference](/reference/error-codes)
