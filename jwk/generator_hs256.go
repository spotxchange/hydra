// Copyright © 2017 Aeneas Rekkas <aeneas+oss@aeneas.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jwk

import (
	"crypto/rand"
	"crypto/x509"
	"io"

	"github.com/pkg/errors"
	jose "gopkg.in/square/go-jose.v1"
)

type HS256Generator struct{}

func (g *HS256Generator) Generate(id string) (*jose.JsonWebKeySet, error) {
	// Taken from NewHMACKey
	key := &[16]byte{}
	_, err := io.ReadFull(rand.Reader, key[:])
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var sliceKey = key[:]

	return &jose.JsonWebKeySet{
		Keys: []jose.JsonWebKey{
			{
				Algorithm:    "HS256",
				Key:          sliceKey,
				KeyID:        id,
				Certificates: []*x509.Certificate{},
			},
		},
	}, nil
}
