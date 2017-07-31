// Copyright 2017 Kenji Kaneda.
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

package store

import (
	"github.com/google/btree"

	"sync"
)

// RangeManager ...
type RangeManager struct {
	ranges *btree.BTree
	sync.Mutex
}

// NewRangeManager is ...
func NewRangeManager() *RangeManager {
	return &RangeManager{
		ranges: btree.New(4 /* degree */),
	}
}

// Split is ...
func (rs *RangeManager) Split(r *Range) {
	rs.Lock()
	defer rs.Unlock()

	// ....
}

// Merge is ...
func (rs *RangeManager) Merge(r1, r2 *Range) {
	rs.Lock()
	defer rs.Unlock()

	// ...
}

// String is ...
func (rs *RangeManager) String() string {
	return ""
}
