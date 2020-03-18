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

package node

import (
	"io/ioutil"

	"github.com/oneinfra/oneinfra/internal/pkg/crypto"
	"github.com/oneinfra/oneinfra/internal/pkg/node"
)

// Join joins a node to an existing cluster
func Join(nodename, apiServerEndpoint, caCertFile, token, joinTokenPublicKeyFile, containerRuntimeEndpoint, imageServiceEndpoint string) error {
	caCert, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		return err
	}
	joinTokenPublicKey, err := crypto.NewPublicKeyFromFile(joinTokenPublicKeyFile)
	if err != nil {
		return err
	}
	return node.Join(nodename, apiServerEndpoint, string(caCert), token, joinTokenPublicKey, containerRuntimeEndpoint, imageServiceEndpoint)
}
