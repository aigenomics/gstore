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

//	"sync"

// RangeScheduler is a scheduler that assigns ssign each range to a set
// of nodes.
type RangeScheduler struct {
	// A list of ranges that need (re)scheduling.
}

// Schedule assigns a given range to a set of nodes satisfying a certain
// condition.
//
// If we don't do any erasure coding, the condition is simple and it
// will be something like
//
// - For each range R, R is instantiated on at lease 3 healthy nodes.
// - Instances of R are distributed well over failure domains.
//
// When we apply erasure coding and additional columar store, we take a
// simple approach where a distinct set of nodes are assigned to
// the columnar range
func (ra *RangeScheduler) Schedule() {
}

// RangeRebalancer triggers reassignemnt if ranges are not distributed
// very well and need split/merge.
type RangeRebalancer struct {
}

// OnEvent implements some listener interface
//
// e.g.) remove a range from a node when a nodeRemoved event is
// received.
func (rs *RangeRebalancer) OnEvent() {
}

// RangeManager ...
type RangeManager struct {
	rangeScheduler  RangeScheduler
	rangeRebalancer RangeRebalancer

	mStore *MasterStore
}

// NewRangeManager is ...
func NewRangeManager() *RangeManager {
	return &RangeManager{}
}

// Assign assigns a range to a given node.
// 1) update the state of the master
// 2) send a command to each node.
//
// As for the second step, we have the following design questions
//
// - Q1: How should the command be sent? (durable) queue in the
//   master? state sync like Borg? Either approach would probably work.
//
// - Q2: How should the node populate the data for a newly assigned range?
//   It needs to copy data from some other nodes. Need to make sure if there is
//   no ongoing write by checking txn. Maybe we need a separate process like
//   range syncer.
//   Just allowing an assignment only when no ongoing txn exists is an easy
//   solution, but that might be too naive.
func (rs *RangeManager) Assign(r *Range) {
}

// Delete deletes an excess range instance we don't need.
func (rs *RangeManager) Delete(r *Range) {
}

// Split is ...
func (rs *RangeManager) Split(r *Range) {
}

// Merge is ...
func (rs *RangeManager) Merge(r1, r2 *Range) {
}
