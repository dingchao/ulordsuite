// Copyright (c) 2014-2017 The ulordsuite developers
// Copyright (c) 2015-2017 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package ulordjson_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/ulordsuite/ulord/ulordjson"
)

// TestChainSvrWsCmds tests all of the chain server websocket-specific commands
// marshal and unmarshal into valid results include handling of optional fields
// being omitted in the marshalled command, while optional fields with defaults
// have the default assigned on unmarshalled commands.
func TestChainSvrWsCmds(t *testing.T) {
	t.Parallel()

	testID := int(1)
	tests := []struct {
		name         string
		newCmd       func() (interface{}, error)
		staticCmd    func() interface{}
		marshalled   string
		unmarshalled interface{}
	}{
		{
			name: "authenticate",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("authenticate", "user", "pass")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewAuthenticateCmd("user", "pass")
			},
			marshalled:   `{"jsonrpc":"1.0","method":"authenticate","params":["user","pass"],"id":1}`,
			unmarshalled: &ulordjson.AuthenticateCmd{Username: "user", Passphrase: "pass"},
		},
		{
			name: "notifyblocks",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("notifyblocks")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewNotifyBlocksCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"notifyblocks","params":[],"id":1}`,
			unmarshalled: &ulordjson.NotifyBlocksCmd{},
		},
		{
			name: "stopnotifyblocks",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("stopnotifyblocks")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewStopNotifyBlocksCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"stopnotifyblocks","params":[],"id":1}`,
			unmarshalled: &ulordjson.StopNotifyBlocksCmd{},
		},
		{
			name: "notifynewtransactions",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("notifynewtransactions")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewNotifyNewTransactionsCmd(nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"notifynewtransactions","params":[],"id":1}`,
			unmarshalled: &ulordjson.NotifyNewTransactionsCmd{
				Verbose: ulordjson.Bool(false),
			},
		},
		{
			name: "notifynewtransactions optional",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("notifynewtransactions", true)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewNotifyNewTransactionsCmd(ulordjson.Bool(true))
			},
			marshalled: `{"jsonrpc":"1.0","method":"notifynewtransactions","params":[true],"id":1}`,
			unmarshalled: &ulordjson.NotifyNewTransactionsCmd{
				Verbose: ulordjson.Bool(true),
			},
		},
		{
			name: "stopnotifynewtransactions",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("stopnotifynewtransactions")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewStopNotifyNewTransactionsCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"stopnotifynewtransactions","params":[],"id":1}`,
			unmarshalled: &ulordjson.StopNotifyNewTransactionsCmd{},
		},
		{
			name: "notifyreceived",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("notifyreceived", []string{"1Address"})
			},
			staticCmd: func() interface{} {
				return ulordjson.NewNotifyReceivedCmd([]string{"1Address"})
			},
			marshalled: `{"jsonrpc":"1.0","method":"notifyreceived","params":[["1Address"]],"id":1}`,
			unmarshalled: &ulordjson.NotifyReceivedCmd{
				Addresses: []string{"1Address"},
			},
		},
		{
			name: "stopnotifyreceived",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("stopnotifyreceived", []string{"1Address"})
			},
			staticCmd: func() interface{} {
				return ulordjson.NewStopNotifyReceivedCmd([]string{"1Address"})
			},
			marshalled: `{"jsonrpc":"1.0","method":"stopnotifyreceived","params":[["1Address"]],"id":1}`,
			unmarshalled: &ulordjson.StopNotifyReceivedCmd{
				Addresses: []string{"1Address"},
			},
		},
		{
			name: "notifyspent",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("notifyspent", `[{"hash":"123","index":0}]`)
			},
			staticCmd: func() interface{} {
				ops := []ulordjson.OutPoint{{Hash: "123", Index: 0}}
				return ulordjson.NewNotifySpentCmd(ops)
			},
			marshalled: `{"jsonrpc":"1.0","method":"notifyspent","params":[[{"hash":"123","index":0}]],"id":1}`,
			unmarshalled: &ulordjson.NotifySpentCmd{
				OutPoints: []ulordjson.OutPoint{{Hash: "123", Index: 0}},
			},
		},
		{
			name: "stopnotifyspent",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("stopnotifyspent", `[{"hash":"123","index":0}]`)
			},
			staticCmd: func() interface{} {
				ops := []ulordjson.OutPoint{{Hash: "123", Index: 0}}
				return ulordjson.NewStopNotifySpentCmd(ops)
			},
			marshalled: `{"jsonrpc":"1.0","method":"stopnotifyspent","params":[[{"hash":"123","index":0}]],"id":1}`,
			unmarshalled: &ulordjson.StopNotifySpentCmd{
				OutPoints: []ulordjson.OutPoint{{Hash: "123", Index: 0}},
			},
		},
		{
			name: "rescan",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("rescan", "123", `["1Address"]`, `[{"hash":"0000000000000000000000000000000000000000000000000000000000000123","index":0}]`)
			},
			staticCmd: func() interface{} {
				addrs := []string{"1Address"}
				ops := []ulordjson.OutPoint{{
					Hash:  "0000000000000000000000000000000000000000000000000000000000000123",
					Index: 0,
				}}
				return ulordjson.NewRescanCmd("123", addrs, ops, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"rescan","params":["123",["1Address"],[{"hash":"0000000000000000000000000000000000000000000000000000000000000123","index":0}]],"id":1}`,
			unmarshalled: &ulordjson.RescanCmd{
				BeginBlock: "123",
				Addresses:  []string{"1Address"},
				OutPoints:  []ulordjson.OutPoint{{Hash: "0000000000000000000000000000000000000000000000000000000000000123", Index: 0}},
				EndBlock:   nil,
			},
		},
		{
			name: "rescan optional",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("rescan", "123", `["1Address"]`, `[{"hash":"123","index":0}]`, "456")
			},
			staticCmd: func() interface{} {
				addrs := []string{"1Address"}
				ops := []ulordjson.OutPoint{{Hash: "123", Index: 0}}
				return ulordjson.NewRescanCmd("123", addrs, ops, ulordjson.String("456"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"rescan","params":["123",["1Address"],[{"hash":"123","index":0}],"456"],"id":1}`,
			unmarshalled: &ulordjson.RescanCmd{
				BeginBlock: "123",
				Addresses:  []string{"1Address"},
				OutPoints:  []ulordjson.OutPoint{{Hash: "123", Index: 0}},
				EndBlock:   ulordjson.String("456"),
			},
		},
		{
			name: "loadtxfilter",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("loadtxfilter", false, `["1Address"]`, `[{"hash":"0000000000000000000000000000000000000000000000000000000000000123","index":0}]`)
			},
			staticCmd: func() interface{} {
				addrs := []string{"1Address"}
				ops := []ulordjson.OutPoint{{
					Hash:  "0000000000000000000000000000000000000000000000000000000000000123",
					Index: 0,
				}}
				return ulordjson.NewLoadTxFilterCmd(false, addrs, ops)
			},
			marshalled: `{"jsonrpc":"1.0","method":"loadtxfilter","params":[false,["1Address"],[{"hash":"0000000000000000000000000000000000000000000000000000000000000123","index":0}]],"id":1}`,
			unmarshalled: &ulordjson.LoadTxFilterCmd{
				Reload:    false,
				Addresses: []string{"1Address"},
				OutPoints: []ulordjson.OutPoint{{Hash: "0000000000000000000000000000000000000000000000000000000000000123", Index: 0}},
			},
		},
		{
			name: "rescanblocks",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("rescanblocks", `["0000000000000000000000000000000000000000000000000000000000000123"]`)
			},
			staticCmd: func() interface{} {
				blockhashes := []string{"0000000000000000000000000000000000000000000000000000000000000123"}
				return ulordjson.NewRescanBlocksCmd(blockhashes)
			},
			marshalled: `{"jsonrpc":"1.0","method":"rescanblocks","params":[["0000000000000000000000000000000000000000000000000000000000000123"]],"id":1}`,
			unmarshalled: &ulordjson.RescanBlocksCmd{
				BlockHashes: []string{"0000000000000000000000000000000000000000000000000000000000000123"},
			},
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		// Marshal the command as created by the new static command
		// creation function.
		marshalled, err := ulordjson.MarshalCmd(testID, test.staticCmd())
		if err != nil {
			t.Errorf("MarshalCmd #%d (%s) unexpected error: %v", i,
				test.name, err)
			continue
		}

		if !bytes.Equal(marshalled, []byte(test.marshalled)) {
			t.Errorf("Test #%d (%s) unexpected marshalled data - "+
				"got %s, want %s", i, test.name, marshalled,
				test.marshalled)
			continue
		}

		// Ensure the command is created without error via the generic
		// new command creation function.
		cmd, err := test.newCmd()
		if err != nil {
			t.Errorf("Test #%d (%s) unexpected NewCmd error: %v ",
				i, test.name, err)
		}

		// Marshal the command as created by the generic new command
		// creation function.
		marshalled, err = ulordjson.MarshalCmd(testID, cmd)
		if err != nil {
			t.Errorf("MarshalCmd #%d (%s) unexpected error: %v", i,
				test.name, err)
			continue
		}

		if !bytes.Equal(marshalled, []byte(test.marshalled)) {
			t.Errorf("Test #%d (%s) unexpected marshalled data - "+
				"got %s, want %s", i, test.name, marshalled,
				test.marshalled)
			continue
		}

		var request ulordjson.Request
		if err := json.Unmarshal(marshalled, &request); err != nil {
			t.Errorf("Test #%d (%s) unexpected error while "+
				"unmarshalling JSON-RPC request: %v", i,
				test.name, err)
			continue
		}

		cmd, err = ulordjson.UnmarshalCmd(&request)
		if err != nil {
			t.Errorf("UnmarshalCmd #%d (%s) unexpected error: %v", i,
				test.name, err)
			continue
		}

		if !reflect.DeepEqual(cmd, test.unmarshalled) {
			t.Errorf("Test #%d (%s) unexpected unmarshalled command "+
				"- got %s, want %s", i, test.name,
				fmt.Sprintf("(%T) %+[1]v", cmd),
				fmt.Sprintf("(%T) %+[1]v\n", test.unmarshalled))
			continue
		}
	}
}
