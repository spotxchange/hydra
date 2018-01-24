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
	"sync"

	"github.com/pkg/errors"
	"github.com/spotxchange/hydra/pkg"
	"gopkg.in/square/go-jose.v1"
)

type MemoryManager struct {
	Keys map[string]*jose.JsonWebKeySet
	sync.RWMutex
}

func (m *MemoryManager) AddKey(set string, key *jose.JsonWebKey) error {
	m.Lock()
	defer m.Unlock()

	m.alloc()
	if m.Keys[set] == nil {
		m.Keys[set] = &jose.JsonWebKeySet{Keys: []jose.JsonWebKey{}}
	}
	m.Keys[set].Keys = append(m.Keys[set].Keys, *key)
	return nil
}

func (m *MemoryManager) AddKeySet(set string, keys *jose.JsonWebKeySet) error {
	for _, key := range keys.Keys {
		m.AddKey(set, &key)
	}
	return nil
}

func (m *MemoryManager) GetKey(set, kid string) (*jose.JsonWebKeySet, error) {
	m.RLock()
	defer m.RUnlock()

	m.alloc()
	keys, found := m.Keys[set]
	if !found {
		return nil, errors.Wrap(pkg.ErrNotFound, "")
	}

	result := keys.Key(kid)
	if len(result) == 0 {
		return nil, errors.Wrap(pkg.ErrNotFound, "")
	}

	return &jose.JsonWebKeySet{
		Keys: result,
	}, nil
}

func (m *MemoryManager) GetKeySet(set string) (*jose.JsonWebKeySet, error) {
	m.RLock()
	defer m.RUnlock()

	m.alloc()
	keys, found := m.Keys[set]
	if !found {
		return nil, errors.Wrap(pkg.ErrNotFound, "")
	}

	return keys, nil
}

func (m *MemoryManager) DeleteKey(set, kid string) error {
	keys, err := m.GetKeySet(set)
	if err != nil {
		return err
	}

	m.Lock()
	var results []jose.JsonWebKey
	for _, key := range keys.Keys {
		if key.KeyID != kid {
			results = append(results)
		}
	}
	m.Keys[set].Keys = results
	defer m.Unlock()

	return nil
}

func (m *MemoryManager) DeleteKeySet(set string) error {
	m.Lock()
	defer m.Unlock()

	delete(m.Keys, set)
	return nil
}

func (m *MemoryManager) alloc() {
	if m.Keys == nil {
		m.Keys = make(map[string]*jose.JsonWebKeySet)
	}
}
