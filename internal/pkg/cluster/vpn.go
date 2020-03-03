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

package cluster

import (
	"net"

	clusterv1alpha1 "github.com/oneinfra/oneinfra/m/apis/cluster/v1alpha1"
)

// VPNPeer represents a VPN peer
type VPNPeer struct {
	Name       string
	Address    string
	PrivateKey string
	PublicKey  string
}

// VPNPeerList represents a list of VPN peers
type VPNPeerList []VPNPeer

func newVPNPeersFromv1alpha1(peers []clusterv1alpha1.VPNPeer) VPNPeerList {
	res := []VPNPeer{}
	for _, peer := range peers {
		res = append(res, VPNPeer{
			Name:       peer.Name,
			Address:    peer.Address,
			PrivateKey: peer.PrivateKey,
			PublicKey:  peer.PublicKey,
		})
	}
	return res
}

func newVPNCIDRFromv1alpha1(vpnCIDR string) *net.IPNet {
	_, ipNet, err := net.ParseCIDR(vpnCIDR)
	if err != nil {
		return &net.IPNet{}
	}
	return ipNet
}

func newVPNPeer(name, address, privateKey, publicKey string) VPNPeer {
	return VPNPeer{
		Name:       name,
		Address:    address,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
}

// Export exports the list of VPN peers to a versioned list of VPN peers
func (vpnPeerList VPNPeerList) Export() []clusterv1alpha1.VPNPeer {
	res := []clusterv1alpha1.VPNPeer{}
	for _, peer := range vpnPeerList {
		res = append(res, clusterv1alpha1.VPNPeer{
			Name:       peer.Name,
			Address:    peer.Address,
			PrivateKey: peer.PrivateKey,
			PublicKey:  peer.PublicKey,
		})
	}
	return res
}
