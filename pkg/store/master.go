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
	"github.com/aigenomics/gstore/pkg/util"
	"golang.org/x/net/context"
)

// APIServer provdes the RPC/HTTP interface (to end clients).
type APIServer struct {
}

// Master manages ranges and nodes.
//
// Master needs to replicate its state.
//
// Master needs leader election.
//
// Master doesn't need to be a single process running on oen machine.
// It can be split into multiple processes distribute d
type Master struct {
	rangeManager *RangeManager

	nodeManager *NodeManager

	// We call this transaction for convenience, but unlike
	// traditional database, it just manages the states of clients
	// writing to the store (maybe read too).
	txnManager *TxnManager

	apiServer *APIServer

	masterStore *MasterStore

	stopper *util.Stopper
}

// NewMaster ...
func NewMaster() *Master {
	return &Master{}
}

// Run ...
func (m *Master) Run(ctx context.Context) {
}

// Stop ...
func (m *Master) Stop() {
	m.stopper.Stop()
}
