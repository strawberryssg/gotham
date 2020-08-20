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

package hugolib

import (
	"github.com/gothamhq/gotham/common/collections"
	"github.com/gothamhq/gotham/resources/page"
)

var (
	_ collections.Grouper = (*pageState)(nil)
	_ collections.Slicer  = (*pageState)(nil)
)

// collections.Slicer implementations below. We keep these bridge implementations
// here as it makes it easier to get an idea of "type coverage". These
// implementations have no value on their own.

// Slice is not meant to be used externally. It's a bridge function
// for the template functions. See collections.Slice.
func (p *pageState) Slice(items interface{}) (interface{}, error) {
	return page.ToPages(items)
}

// collections.Grouper  implementations below

// Group creates a PageGroup from a key and a Pages object
// This method is not meant for external use. It got its non-typed arguments to satisfy
// a very generic interface in the tpl package.
func (p *pageState) Group(key interface{}, in interface{}) (interface{}, error) {
	pages, err := page.ToPages(in)
	if err != nil {
		return nil, err
	}
	return page.PageGroup{Key: key, Pages: pages}, nil
}