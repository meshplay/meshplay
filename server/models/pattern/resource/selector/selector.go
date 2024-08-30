package selector

import (
	meshmodel "github.com/khulnasoft/meshkit/models/meshmodel/registry"
)

const (
	CoreResource = "pattern.meshplay.khulnasofy.com/core"
	MeshResource = "pattern.meshplay.khulnasofy.com/mesh/workload"
	K8sResource  = "pattern.meshplay.khulnasofy.com/k8s"
)

type Selector struct {
	registry *meshmodel.RegistryManager
}

func New(reg *meshmodel.RegistryManager) *Selector {
	return &Selector{
		registry: reg,
	}
}
