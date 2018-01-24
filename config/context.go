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

package config

import (
	"github.com/ory/ladon"
	"github.com/spotxchange/fosite"
	"github.com/spotxchange/fosite/handler/oauth2"
	"github.com/spotxchange/hydra/firewall"
	"github.com/spotxchange/hydra/jwk"
	hoa2 "github.com/spotxchange/hydra/oauth2"
	"github.com/spotxchange/hydra/pkg"
	"github.com/spotxchange/hydra/warden/group"
)

type Context struct {
	Connection interface{}

	Hasher         fosite.Hasher
	Warden         firewall.Firewall
	LadonManager   ladon.Manager
	FositeStrategy oauth2.CoreStrategy
	FositeStore    pkg.FositeStorer
	KeyManager     jwk.Manager
	ConsentManager hoa2.ConsentRequestManager
	GroupManager   group.Manager
}
