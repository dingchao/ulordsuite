// Copyright (c) 2017 The ulordsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package netsync

import (
	"github.com/ulordsuite/ulord/blockchain"
	"github.com/ulordsuite/ulord/chaincfg"
	"github.com/ulordsuite/ulord/chaincfg/chainhash"
	"github.com/ulordsuite/ulord/mempool"
	"github.com/ulordsuite/ulord/peer"
	"github.com/ulordsuite/ulord/wire"
	"github.com/ulordsuite/ulordutil"
)

// PeerNotifier exposes methods to notify peers of status changes to
// transactions, blocks, etc. Currently server (in the main package) implements
// this interface.
type PeerNotifier interface {
	AnnounceNewTransactions(newTxs []*mempool.TxDesc)

	UpdatePeerHeights(latestBlkHash *chainhash.Hash, latestHeight int32, updateSource *peer.Peer)

	RelayInventory(invVect *wire.InvVect, data interface{})

	TransactionConfirmed(tx *ulordutil.Tx)
}

// Config is a configuration struct used to initialize a new SyncManager.
type Config struct {
	PeerNotifier PeerNotifier
	Chain        *blockchain.BlockChain
	TxMemPool    *mempool.TxPool
	ChainParams  *chaincfg.Params

	DisableCheckpoints bool
	MaxPeers           int

	FeeEstimator *mempool.FeeEstimator
}
