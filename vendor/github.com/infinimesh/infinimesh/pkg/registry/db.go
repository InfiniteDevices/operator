//--------------------------------------------------------------------------
// Copyright 2018 Infinite Devices GmbH
// www.infinimesh.io
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
//--------------------------------------------------------------------------

package registry

import (
	"github.com/infinimesh/infinimesh/pkg/node/dgraph"
)

//Device data struct with Certificates data strcuture as slice
type Device struct {
	dgraph.Object
	Tags    []string `json:"tags,omitempty"`
	Enabled bool     `json:"enabled"`

	Certificates []*X509Cert `json:"certificates,omitempty"`
}

//X509Cert is Certificate data struct which is refered in Device data strcuture
type X509Cert struct {
	dgraph.Node
	PemData              string `json:"pem_data,omitempty"`
	Algorithm            string `json:"algorithm,omitempty"`
	Fingerprint          []byte `json:"fingerprint,omitempty"`
	FingerprintAlgorithm string `json:"fingerprint.algorithm,omitempty"`
}
