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

package util

// Stopper is ..
type Stopper struct {
	c chan struct{}
}

// NewStopper ..
func NewStopper() *Stopper {
	return &Stopper{
		c: make(chan struct{}),
	}
}

// IsStopped ...
func (s *Stopper) IsStopped() bool {
	// Non-blocking chann
	//	switch s.c {

	//	}
	return false
}

// Stop ...
func (s *Stopper) Stop() {
	close(s.c)
}
