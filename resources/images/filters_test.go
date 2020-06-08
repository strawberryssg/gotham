// Copyright 2019 The Hugo Authors. All rights reserved.
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

package images

import (
	"testing"

	"github.com/gothamhq/gotham/helpers"

	qt "github.com/frankban/quicktest"
)

func TestFilterHash(t *testing.T) {
	c := qt.New(t)

	f := &Filters{}

	c.Assert(helpers.HashString(f.Grayscale()), qt.Equals, helpers.HashString(f.Grayscale()))
	c.Assert(helpers.HashString(f.Grayscale()), qt.Not(qt.Equals), helpers.HashString(f.Invert()))
	c.Assert(helpers.HashString(f.Gamma(32)), qt.Not(qt.Equals), helpers.HashString(f.Gamma(33)))
	c.Assert(helpers.HashString(f.Gamma(32)), qt.Equals, helpers.HashString(f.Gamma(32)))

}
