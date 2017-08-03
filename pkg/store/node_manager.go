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

// HealthChecker monitors the status of each node by
// periodically sending heartbeat requests.
// It can consist of multiple gorountines, each of which
// is responsible for a subset of nodes (simple sharding
// by node IDs).
type HealthChecker struct {
}

// NodeManager keeks track of the state of each node.
type NodeManager struct {
	mStore *MasterStore

	healthChecker HealthChecker
}
