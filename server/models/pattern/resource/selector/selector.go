package selector

import (
	meshmodel "github.com/layer5io/meshkit/models/meshmodel/registry"
)

const (
	CoreResource = "pattern.meshplay.io/core"
	MeshResource = "pattern.meshplay.io/mesh/workload"
	K8sResource  = "pattern.meshplay.io/k8s"
)

type Selector struct {
	registry *meshmodel.RegistryManager
}

func New(reg *meshmodel.RegistryManager) *Selector {
	return &Selector{
		registry: reg,
	}
}
