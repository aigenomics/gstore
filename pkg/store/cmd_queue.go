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
	"sync"
)

// Command is ...
type Command struct {
	// Get, Put, PutIf, Scan, ...
	// Ops for updating ranges
	t int32
}

// CmdQueue is ...
type CmdQueue struct {
	cmds []*Command
	sync.Mutex
}

// NewCmdQueue returns a new
func NewCmdQueue() *CmdQueue {
	return &CmdQueue{}
}

// Add adds ...
func (q *CmdQueue) Add(cmd *Command) {
	q.Lock()
	defer q.Unlock()
	q.cmds = append(q.cmds, cmd)
}
