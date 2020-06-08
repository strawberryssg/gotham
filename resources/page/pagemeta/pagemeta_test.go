// Copyright 2020 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pagemeta

import (
	"fmt"
	"testing"

	"github.com/gothamhq/gotham/htesting/hqt"

	"github.com/gothamhq/gotham/config"

	qt "github.com/frankban/quicktest"
)

func TestDecodeBuildConfig(t *testing.T) {
	t.Parallel()

	c := qt.New(t)

	configTempl := `
[_build]
render = true
list = %s
publishResources = true`

	for _, test := range []struct {
		list   interface{}
		expect string
	}{
		{"true", Always},
		{"false", Never},
		{`"always"`, Always},
		{`"local"`, ListLocally},
		{`"asdfadf"`, Always},
	} {
		cfg, err := config.FromConfigString(fmt.Sprintf(configTempl, test.list), "toml")
		c.Assert(err, qt.IsNil)
		bcfg, err := DecodeBuildConfig(cfg.Get("_build"))
		c.Assert(err, qt.IsNil)

		eq := qt.CmpEquals(hqt.DeepAllowUnexported(BuildConfig{}))

		c.Assert(bcfg, eq, BuildConfig{
			Render:           true,
			List:             test.expect,
			PublishResources: true,
			set:              true,
		})

	}

}
