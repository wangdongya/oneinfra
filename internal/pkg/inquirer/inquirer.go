/*
Copyright 2020 Rafael Fernández López <ereslibre@ereslibre.es>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package inquirer

import (
	"github.com/oneinfra/oneinfra/m/internal/pkg/cluster"
	"github.com/oneinfra/oneinfra/m/internal/pkg/component"
	"github.com/oneinfra/oneinfra/m/internal/pkg/infra"
)

// ReconcilerInquirer represents an interface that allows a
// reconciliation cycle to retrieve information about the cluster
type ReconcilerInquirer interface {
	Component() *component.Component
	Hypervisor() *infra.Hypervisor
	Cluster() *cluster.Cluster
	ClusterComponents(component.Role) component.List
	ComponentHypervisor(*component.Component) *infra.Hypervisor
}
