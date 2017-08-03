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
	"github.com/google/uuid"

	"sync"
)

// TxnManager ...
type TxnManager struct {
	mState *MasterState

	sync.Mutex
}

// NewTxnManager ...
func NewTxnManager() *TxnManager {
	return &TxnManager{
	//		txnMap: make(map[uuid.UUID]*Txn),
	}
}

// OpenTxn ...
func (tm *TxnManager) OpenTxn() *Txn {
	tm.Lock()
	defer tm.Unlock()

	txn := &Txn{id: uuid.New()}
	//	tm.txnMap[txn.id] = txn

	return txn
}

// CloseTxn ...
func (tm *TxnManager) CloseTxn(txn *Txn) {
	tm.Lock()
	defer tm.Unlock()

	//	delete(tm.txnMap, txn.id)
}
