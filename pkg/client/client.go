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

// Client is ..
type Client struct {
}

// Write ...
func (c *Client) Write() {

	// Talk to the master
	// Create a txn (mostly lock)
	// Get a set of
	//
	// Split the data into stripes
	//
	//
	// Encode the stripe.
	//
	// Write maybe directly to each replica and check consistency
	//
	// Close the txn
}

// Read ..
func (c *Client) Read() {
	// Talk to the master and
}

// aas
