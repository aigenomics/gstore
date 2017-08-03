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

// MemDB is a simple in-memory key value store. It is used primarily
// for testing.
//
// This implements some KV interface
type MemDB struct {
	m map[string][]byte
}

// NewMemDB is ..
func NewMemDB() *MemDB {
	return &MemDB{
		m: make(map[string][]byte),
	}
}

// Put ...
func (db *MemDB) Put(key []byte, val []byte) error {
	db.m[string(key)] = val
	return nil
}

// Get ...
func (db *MemDB) Get(key []byte) ([]byte, error) {
	val, _ := db.m[string(key)]
	//	if !ok {
	//	}
	return val, nil
}

// Get and Scan?
