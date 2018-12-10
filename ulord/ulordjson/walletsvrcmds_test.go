// Copyright (c) 2014 The ulordsuite developers
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

// TestWalletSvrCmds tests all of the wallet server commands marshal and
// unmarshal into valid results include handling of optional fields being
// omitted in the marshalled command, while optional fields with defaults have
// the default assigned on unmarshalled commands.
func TestWalletSvrCmds(t *testing.T) {
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
			name: "addmultisigaddress",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("addmultisigaddress", 2, []string{"031234", "035678"})
			},
			staticCmd: func() interface{} {
				keys := []string{"031234", "035678"}
				return ulordjson.NewAddMultisigAddressCmd(2, keys, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"addmultisigaddress","params":[2,["031234","035678"]],"id":1}`,
			unmarshalled: &ulordjson.AddMultisigAddressCmd{
				NRequired: 2,
				Keys:      []string{"031234", "035678"},
				Account:   nil,
			},
		},
		{
			name: "addmultisigaddress optional",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("addmultisigaddress", 2, []string{"031234", "035678"}, "test")
			},
			staticCmd: func() interface{} {
				keys := []string{"031234", "035678"}
				return ulordjson.NewAddMultisigAddressCmd(2, keys, ulordjson.String("test"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"addmultisigaddress","params":[2,["031234","035678"],"test"],"id":1}`,
			unmarshalled: &ulordjson.AddMultisigAddressCmd{
				NRequired: 2,
				Keys:      []string{"031234", "035678"},
				Account:   ulordjson.String("test"),
			},
		},
		{
			name: "addwitnessaddress",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("addwitnessaddress", "1address")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewAddWitnessAddressCmd("1address")
			},
			marshalled: `{"jsonrpc":"1.0","method":"addwitnessaddress","params":["1address"],"id":1}`,
			unmarshalled: &ulordjson.AddWitnessAddressCmd{
				Address: "1address",
			},
		},
		{
			name: "createmultisig",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("createmultisig", 2, []string{"031234", "035678"})
			},
			staticCmd: func() interface{} {
				keys := []string{"031234", "035678"}
				return ulordjson.NewCreateMultisigCmd(2, keys)
			},
			marshalled: `{"jsonrpc":"1.0","method":"createmultisig","params":[2,["031234","035678"]],"id":1}`,
			unmarshalled: &ulordjson.CreateMultisigCmd{
				NRequired: 2,
				Keys:      []string{"031234", "035678"},
			},
		},
		{
			name: "dumpprivkey",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("dumpprivkey", "1Address")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewDumpPrivKeyCmd("1Address")
			},
			marshalled: `{"jsonrpc":"1.0","method":"dumpprivkey","params":["1Address"],"id":1}`,
			unmarshalled: &ulordjson.DumpPrivKeyCmd{
				Address: "1Address",
			},
		},
		{
			name: "encryptwallet",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("encryptwallet", "pass")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewEncryptWalletCmd("pass")
			},
			marshalled: `{"jsonrpc":"1.0","method":"encryptwallet","params":["pass"],"id":1}`,
			unmarshalled: &ulordjson.EncryptWalletCmd{
				Passphrase: "pass",
			},
		},
		{
			name: "estimatefee",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("estimatefee", 6)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewEstimateFeeCmd(6)
			},
			marshalled: `{"jsonrpc":"1.0","method":"estimatefee","params":[6],"id":1}`,
			unmarshalled: &ulordjson.EstimateFeeCmd{
				NumBlocks: 6,
			},
		},
		{
			name: "estimatepriority",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("estimatepriority", 6)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewEstimatePriorityCmd(6)
			},
			marshalled: `{"jsonrpc":"1.0","method":"estimatepriority","params":[6],"id":1}`,
			unmarshalled: &ulordjson.EstimatePriorityCmd{
				NumBlocks: 6,
			},
		},
		{
			name: "getaccount",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("getaccount", "1Address")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewGetAccountCmd("1Address")
			},
			marshalled: `{"jsonrpc":"1.0","method":"getaccount","params":["1Address"],"id":1}`,
			unmarshalled: &ulordjson.GetAccountCmd{
				Address: "1Address",
			},
		},
		{
			name: "getaccountaddress",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("getaccountaddress", "acct")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewGetAccountAddressCmd("acct")
			},
			marshalled: `{"jsonrpc":"1.0","method":"getaccountaddress","params":["acct"],"id":1}`,
			unmarshalled: &ulordjson.GetAccountAddressCmd{
				Account: "acct",
			},
		},
		{
			name: "getaddressesbyaccount",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("getaddressesbyaccount", "acct")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewGetAddressesByAccountCmd("acct")
			},
			marshalled: `{"jsonrpc":"1.0","method":"getaddressesbyaccount","params":["acct"],"id":1}`,
			unmarshalled: &ulordjson.GetAddressesByAccountCmd{
				Account: "acct",
			},
		},
		{
			name: "getbalance",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("getbalance")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewGetBalanceCmd(nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getbalance","params":[],"id":1}`,
			unmarshalled: &ulordjson.GetBalanceCmd{
				Account: nil,
				MinConf: ulordjson.Int(1),
			},
		},
		{
			name: "getbalance optional1",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("getbalance", "acct")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewGetBalanceCmd(ulordjson.String("acct"), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getbalance","params":["acct"],"id":1}`,
			unmarshalled: &ulordjson.GetBalanceCmd{
				Account: ulordjson.String("acct"),
				MinConf: ulordjson.Int(1),
			},
		},
		{
			name: "getbalance optional2",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("getbalance", "acct", 6)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewGetBalanceCmd(ulordjson.String("acct"), ulordjson.Int(6))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getbalance","params":["acct",6],"id":1}`,
			unmarshalled: &ulordjson.GetBalanceCmd{
				Account: ulordjson.String("acct"),
				MinConf: ulordjson.Int(6),
			},
		},
		{
			name: "getnewaddress",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("getnewaddress")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewGetNewAddressCmd(nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getnewaddress","params":[],"id":1}`,
			unmarshalled: &ulordjson.GetNewAddressCmd{
				Account: nil,
			},
		},
		{
			name: "getnewaddress optional",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("getnewaddress", "acct")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewGetNewAddressCmd(ulordjson.String("acct"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getnewaddress","params":["acct"],"id":1}`,
			unmarshalled: &ulordjson.GetNewAddressCmd{
				Account: ulordjson.String("acct"),
			},
		},
		{
			name: "getrawchangeaddress",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("getrawchangeaddress")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewGetRawChangeAddressCmd(nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getrawchangeaddress","params":[],"id":1}`,
			unmarshalled: &ulordjson.GetRawChangeAddressCmd{
				Account: nil,
			},
		},
		{
			name: "getrawchangeaddress optional",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("getrawchangeaddress", "acct")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewGetRawChangeAddressCmd(ulordjson.String("acct"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getrawchangeaddress","params":["acct"],"id":1}`,
			unmarshalled: &ulordjson.GetRawChangeAddressCmd{
				Account: ulordjson.String("acct"),
			},
		},
		{
			name: "getreceivedbyaccount",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("getreceivedbyaccount", "acct")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewGetReceivedByAccountCmd("acct", nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getreceivedbyaccount","params":["acct"],"id":1}`,
			unmarshalled: &ulordjson.GetReceivedByAccountCmd{
				Account: "acct",
				MinConf: ulordjson.Int(1),
			},
		},
		{
			name: "getreceivedbyaccount optional",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("getreceivedbyaccount", "acct", 6)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewGetReceivedByAccountCmd("acct", ulordjson.Int(6))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getreceivedbyaccount","params":["acct",6],"id":1}`,
			unmarshalled: &ulordjson.GetReceivedByAccountCmd{
				Account: "acct",
				MinConf: ulordjson.Int(6),
			},
		},
		{
			name: "getreceivedbyaddress",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("getreceivedbyaddress", "1Address")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewGetReceivedByAddressCmd("1Address", nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getreceivedbyaddress","params":["1Address"],"id":1}`,
			unmarshalled: &ulordjson.GetReceivedByAddressCmd{
				Address: "1Address",
				MinConf: ulordjson.Int(1),
			},
		},
		{
			name: "getreceivedbyaddress optional",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("getreceivedbyaddress", "1Address", 6)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewGetReceivedByAddressCmd("1Address", ulordjson.Int(6))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getreceivedbyaddress","params":["1Address",6],"id":1}`,
			unmarshalled: &ulordjson.GetReceivedByAddressCmd{
				Address: "1Address",
				MinConf: ulordjson.Int(6),
			},
		},
		{
			name: "gettransaction",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("gettransaction", "123")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewGetTransactionCmd("123", nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"gettransaction","params":["123"],"id":1}`,
			unmarshalled: &ulordjson.GetTransactionCmd{
				Txid:             "123",
				IncludeWatchOnly: ulordjson.Bool(false),
			},
		},
		{
			name: "gettransaction optional",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("gettransaction", "123", true)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewGetTransactionCmd("123", ulordjson.Bool(true))
			},
			marshalled: `{"jsonrpc":"1.0","method":"gettransaction","params":["123",true],"id":1}`,
			unmarshalled: &ulordjson.GetTransactionCmd{
				Txid:             "123",
				IncludeWatchOnly: ulordjson.Bool(true),
			},
		},
		{
			name: "getwalletinfo",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("getwalletinfo")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewGetWalletInfoCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"getwalletinfo","params":[],"id":1}`,
			unmarshalled: &ulordjson.GetWalletInfoCmd{},
		},
		{
			name: "importprivkey",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("importprivkey", "abc")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewImportPrivKeyCmd("abc", nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"importprivkey","params":["abc"],"id":1}`,
			unmarshalled: &ulordjson.ImportPrivKeyCmd{
				PrivKey: "abc",
				Label:   nil,
				Rescan:  ulordjson.Bool(true),
			},
		},
		{
			name: "importprivkey optional1",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("importprivkey", "abc", "label")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewImportPrivKeyCmd("abc", ulordjson.String("label"), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"importprivkey","params":["abc","label"],"id":1}`,
			unmarshalled: &ulordjson.ImportPrivKeyCmd{
				PrivKey: "abc",
				Label:   ulordjson.String("label"),
				Rescan:  ulordjson.Bool(true),
			},
		},
		{
			name: "importprivkey optional2",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("importprivkey", "abc", "label", false)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewImportPrivKeyCmd("abc", ulordjson.String("label"), ulordjson.Bool(false))
			},
			marshalled: `{"jsonrpc":"1.0","method":"importprivkey","params":["abc","label",false],"id":1}`,
			unmarshalled: &ulordjson.ImportPrivKeyCmd{
				PrivKey: "abc",
				Label:   ulordjson.String("label"),
				Rescan:  ulordjson.Bool(false),
			},
		},
		{
			name: "keypoolrefill",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("keypoolrefill")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewKeyPoolRefillCmd(nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"keypoolrefill","params":[],"id":1}`,
			unmarshalled: &ulordjson.KeyPoolRefillCmd{
				NewSize: ulordjson.Uint(100),
			},
		},
		{
			name: "keypoolrefill optional",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("keypoolrefill", 200)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewKeyPoolRefillCmd(ulordjson.Uint(200))
			},
			marshalled: `{"jsonrpc":"1.0","method":"keypoolrefill","params":[200],"id":1}`,
			unmarshalled: &ulordjson.KeyPoolRefillCmd{
				NewSize: ulordjson.Uint(200),
			},
		},
		{
			name: "listaccounts",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listaccounts")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListAccountsCmd(nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listaccounts","params":[],"id":1}`,
			unmarshalled: &ulordjson.ListAccountsCmd{
				MinConf: ulordjson.Int(1),
			},
		},
		{
			name: "listaccounts optional",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listaccounts", 6)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListAccountsCmd(ulordjson.Int(6))
			},
			marshalled: `{"jsonrpc":"1.0","method":"listaccounts","params":[6],"id":1}`,
			unmarshalled: &ulordjson.ListAccountsCmd{
				MinConf: ulordjson.Int(6),
			},
		},
		{
			name: "listaddressgroupings",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listaddressgroupings")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListAddressGroupingsCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"listaddressgroupings","params":[],"id":1}`,
			unmarshalled: &ulordjson.ListAddressGroupingsCmd{},
		},
		{
			name: "listlockunspent",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listlockunspent")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListLockUnspentCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"listlockunspent","params":[],"id":1}`,
			unmarshalled: &ulordjson.ListLockUnspentCmd{},
		},
		{
			name: "listreceivedbyaccount",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listreceivedbyaccount")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListReceivedByAccountCmd(nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaccount","params":[],"id":1}`,
			unmarshalled: &ulordjson.ListReceivedByAccountCmd{
				MinConf:          ulordjson.Int(1),
				IncludeEmpty:     ulordjson.Bool(false),
				IncludeWatchOnly: ulordjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaccount optional1",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listreceivedbyaccount", 6)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListReceivedByAccountCmd(ulordjson.Int(6), nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaccount","params":[6],"id":1}`,
			unmarshalled: &ulordjson.ListReceivedByAccountCmd{
				MinConf:          ulordjson.Int(6),
				IncludeEmpty:     ulordjson.Bool(false),
				IncludeWatchOnly: ulordjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaccount optional2",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listreceivedbyaccount", 6, true)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListReceivedByAccountCmd(ulordjson.Int(6), ulordjson.Bool(true), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaccount","params":[6,true],"id":1}`,
			unmarshalled: &ulordjson.ListReceivedByAccountCmd{
				MinConf:          ulordjson.Int(6),
				IncludeEmpty:     ulordjson.Bool(true),
				IncludeWatchOnly: ulordjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaccount optional3",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listreceivedbyaccount", 6, true, false)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListReceivedByAccountCmd(ulordjson.Int(6), ulordjson.Bool(true), ulordjson.Bool(false))
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaccount","params":[6,true,false],"id":1}`,
			unmarshalled: &ulordjson.ListReceivedByAccountCmd{
				MinConf:          ulordjson.Int(6),
				IncludeEmpty:     ulordjson.Bool(true),
				IncludeWatchOnly: ulordjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaddress",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listreceivedbyaddress")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListReceivedByAddressCmd(nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaddress","params":[],"id":1}`,
			unmarshalled: &ulordjson.ListReceivedByAddressCmd{
				MinConf:          ulordjson.Int(1),
				IncludeEmpty:     ulordjson.Bool(false),
				IncludeWatchOnly: ulordjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaddress optional1",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listreceivedbyaddress", 6)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListReceivedByAddressCmd(ulordjson.Int(6), nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaddress","params":[6],"id":1}`,
			unmarshalled: &ulordjson.ListReceivedByAddressCmd{
				MinConf:          ulordjson.Int(6),
				IncludeEmpty:     ulordjson.Bool(false),
				IncludeWatchOnly: ulordjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaddress optional2",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listreceivedbyaddress", 6, true)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListReceivedByAddressCmd(ulordjson.Int(6), ulordjson.Bool(true), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaddress","params":[6,true],"id":1}`,
			unmarshalled: &ulordjson.ListReceivedByAddressCmd{
				MinConf:          ulordjson.Int(6),
				IncludeEmpty:     ulordjson.Bool(true),
				IncludeWatchOnly: ulordjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaddress optional3",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listreceivedbyaddress", 6, true, false)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListReceivedByAddressCmd(ulordjson.Int(6), ulordjson.Bool(true), ulordjson.Bool(false))
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaddress","params":[6,true,false],"id":1}`,
			unmarshalled: &ulordjson.ListReceivedByAddressCmd{
				MinConf:          ulordjson.Int(6),
				IncludeEmpty:     ulordjson.Bool(true),
				IncludeWatchOnly: ulordjson.Bool(false),
			},
		},
		{
			name: "listsinceblock",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listsinceblock")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListSinceBlockCmd(nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listsinceblock","params":[],"id":1}`,
			unmarshalled: &ulordjson.ListSinceBlockCmd{
				BlockHash:           nil,
				TargetConfirmations: ulordjson.Int(1),
				IncludeWatchOnly:    ulordjson.Bool(false),
			},
		},
		{
			name: "listsinceblock optional1",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listsinceblock", "123")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListSinceBlockCmd(ulordjson.String("123"), nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listsinceblock","params":["123"],"id":1}`,
			unmarshalled: &ulordjson.ListSinceBlockCmd{
				BlockHash:           ulordjson.String("123"),
				TargetConfirmations: ulordjson.Int(1),
				IncludeWatchOnly:    ulordjson.Bool(false),
			},
		},
		{
			name: "listsinceblock optional2",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listsinceblock", "123", 6)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListSinceBlockCmd(ulordjson.String("123"), ulordjson.Int(6), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listsinceblock","params":["123",6],"id":1}`,
			unmarshalled: &ulordjson.ListSinceBlockCmd{
				BlockHash:           ulordjson.String("123"),
				TargetConfirmations: ulordjson.Int(6),
				IncludeWatchOnly:    ulordjson.Bool(false),
			},
		},
		{
			name: "listsinceblock optional3",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listsinceblock", "123", 6, true)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListSinceBlockCmd(ulordjson.String("123"), ulordjson.Int(6), ulordjson.Bool(true))
			},
			marshalled: `{"jsonrpc":"1.0","method":"listsinceblock","params":["123",6,true],"id":1}`,
			unmarshalled: &ulordjson.ListSinceBlockCmd{
				BlockHash:           ulordjson.String("123"),
				TargetConfirmations: ulordjson.Int(6),
				IncludeWatchOnly:    ulordjson.Bool(true),
			},
		},
		{
			name: "listtransactions",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listtransactions")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListTransactionsCmd(nil, nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listtransactions","params":[],"id":1}`,
			unmarshalled: &ulordjson.ListTransactionsCmd{
				Account:          nil,
				Count:            ulordjson.Int(10),
				From:             ulordjson.Int(0),
				IncludeWatchOnly: ulordjson.Bool(false),
			},
		},
		{
			name: "listtransactions optional1",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listtransactions", "acct")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListTransactionsCmd(ulordjson.String("acct"), nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listtransactions","params":["acct"],"id":1}`,
			unmarshalled: &ulordjson.ListTransactionsCmd{
				Account:          ulordjson.String("acct"),
				Count:            ulordjson.Int(10),
				From:             ulordjson.Int(0),
				IncludeWatchOnly: ulordjson.Bool(false),
			},
		},
		{
			name: "listtransactions optional2",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listtransactions", "acct", 20)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListTransactionsCmd(ulordjson.String("acct"), ulordjson.Int(20), nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listtransactions","params":["acct",20],"id":1}`,
			unmarshalled: &ulordjson.ListTransactionsCmd{
				Account:          ulordjson.String("acct"),
				Count:            ulordjson.Int(20),
				From:             ulordjson.Int(0),
				IncludeWatchOnly: ulordjson.Bool(false),
			},
		},
		{
			name: "listtransactions optional3",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listtransactions", "acct", 20, 1)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListTransactionsCmd(ulordjson.String("acct"), ulordjson.Int(20),
					ulordjson.Int(1), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listtransactions","params":["acct",20,1],"id":1}`,
			unmarshalled: &ulordjson.ListTransactionsCmd{
				Account:          ulordjson.String("acct"),
				Count:            ulordjson.Int(20),
				From:             ulordjson.Int(1),
				IncludeWatchOnly: ulordjson.Bool(false),
			},
		},
		{
			name: "listtransactions optional4",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listtransactions", "acct", 20, 1, true)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListTransactionsCmd(ulordjson.String("acct"), ulordjson.Int(20),
					ulordjson.Int(1), ulordjson.Bool(true))
			},
			marshalled: `{"jsonrpc":"1.0","method":"listtransactions","params":["acct",20,1,true],"id":1}`,
			unmarshalled: &ulordjson.ListTransactionsCmd{
				Account:          ulordjson.String("acct"),
				Count:            ulordjson.Int(20),
				From:             ulordjson.Int(1),
				IncludeWatchOnly: ulordjson.Bool(true),
			},
		},
		{
			name: "listunspent",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listunspent")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListUnspentCmd(nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listunspent","params":[],"id":1}`,
			unmarshalled: &ulordjson.ListUnspentCmd{
				MinConf:   ulordjson.Int(1),
				MaxConf:   ulordjson.Int(9999999),
				Addresses: nil,
			},
		},
		{
			name: "listunspent optional1",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listunspent", 6)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListUnspentCmd(ulordjson.Int(6), nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listunspent","params":[6],"id":1}`,
			unmarshalled: &ulordjson.ListUnspentCmd{
				MinConf:   ulordjson.Int(6),
				MaxConf:   ulordjson.Int(9999999),
				Addresses: nil,
			},
		},
		{
			name: "listunspent optional2",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listunspent", 6, 100)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListUnspentCmd(ulordjson.Int(6), ulordjson.Int(100), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listunspent","params":[6,100],"id":1}`,
			unmarshalled: &ulordjson.ListUnspentCmd{
				MinConf:   ulordjson.Int(6),
				MaxConf:   ulordjson.Int(100),
				Addresses: nil,
			},
		},
		{
			name: "listunspent optional3",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("listunspent", 6, 100, []string{"1Address", "1Address2"})
			},
			staticCmd: func() interface{} {
				return ulordjson.NewListUnspentCmd(ulordjson.Int(6), ulordjson.Int(100),
					&[]string{"1Address", "1Address2"})
			},
			marshalled: `{"jsonrpc":"1.0","method":"listunspent","params":[6,100,["1Address","1Address2"]],"id":1}`,
			unmarshalled: &ulordjson.ListUnspentCmd{
				MinConf:   ulordjson.Int(6),
				MaxConf:   ulordjson.Int(100),
				Addresses: &[]string{"1Address", "1Address2"},
			},
		},
		{
			name: "lockunspent",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("lockunspent", true, `[{"txid":"123","vout":1}]`)
			},
			staticCmd: func() interface{} {
				txInputs := []ulordjson.TransactionInput{
					{Txid: "123", Vout: 1},
				}
				return ulordjson.NewLockUnspentCmd(true, txInputs)
			},
			marshalled: `{"jsonrpc":"1.0","method":"lockunspent","params":[true,[{"txid":"123","vout":1}]],"id":1}`,
			unmarshalled: &ulordjson.LockUnspentCmd{
				Unlock: true,
				Transactions: []ulordjson.TransactionInput{
					{Txid: "123", Vout: 1},
				},
			},
		},
		{
			name: "move",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("move", "from", "to", 0.5)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewMoveCmd("from", "to", 0.5, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"move","params":["from","to",0.5],"id":1}`,
			unmarshalled: &ulordjson.MoveCmd{
				FromAccount: "from",
				ToAccount:   "to",
				Amount:      0.5,
				MinConf:     ulordjson.Int(1),
				Comment:     nil,
			},
		},
		{
			name: "move optional1",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("move", "from", "to", 0.5, 6)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewMoveCmd("from", "to", 0.5, ulordjson.Int(6), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"move","params":["from","to",0.5,6],"id":1}`,
			unmarshalled: &ulordjson.MoveCmd{
				FromAccount: "from",
				ToAccount:   "to",
				Amount:      0.5,
				MinConf:     ulordjson.Int(6),
				Comment:     nil,
			},
		},
		{
			name: "move optional2",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("move", "from", "to", 0.5, 6, "comment")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewMoveCmd("from", "to", 0.5, ulordjson.Int(6), ulordjson.String("comment"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"move","params":["from","to",0.5,6,"comment"],"id":1}`,
			unmarshalled: &ulordjson.MoveCmd{
				FromAccount: "from",
				ToAccount:   "to",
				Amount:      0.5,
				MinConf:     ulordjson.Int(6),
				Comment:     ulordjson.String("comment"),
			},
		},
		{
			name: "sendfrom",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("sendfrom", "from", "1Address", 0.5)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewSendFromCmd("from", "1Address", 0.5, nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendfrom","params":["from","1Address",0.5],"id":1}`,
			unmarshalled: &ulordjson.SendFromCmd{
				FromAccount: "from",
				ToAddress:   "1Address",
				Amount:      0.5,
				MinConf:     ulordjson.Int(1),
				Comment:     nil,
				CommentTo:   nil,
			},
		},
		{
			name: "sendfrom optional1",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("sendfrom", "from", "1Address", 0.5, 6)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewSendFromCmd("from", "1Address", 0.5, ulordjson.Int(6), nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendfrom","params":["from","1Address",0.5,6],"id":1}`,
			unmarshalled: &ulordjson.SendFromCmd{
				FromAccount: "from",
				ToAddress:   "1Address",
				Amount:      0.5,
				MinConf:     ulordjson.Int(6),
				Comment:     nil,
				CommentTo:   nil,
			},
		},
		{
			name: "sendfrom optional2",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("sendfrom", "from", "1Address", 0.5, 6, "comment")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewSendFromCmd("from", "1Address", 0.5, ulordjson.Int(6),
					ulordjson.String("comment"), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendfrom","params":["from","1Address",0.5,6,"comment"],"id":1}`,
			unmarshalled: &ulordjson.SendFromCmd{
				FromAccount: "from",
				ToAddress:   "1Address",
				Amount:      0.5,
				MinConf:     ulordjson.Int(6),
				Comment:     ulordjson.String("comment"),
				CommentTo:   nil,
			},
		},
		{
			name: "sendfrom optional3",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("sendfrom", "from", "1Address", 0.5, 6, "comment", "commentto")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewSendFromCmd("from", "1Address", 0.5, ulordjson.Int(6),
					ulordjson.String("comment"), ulordjson.String("commentto"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendfrom","params":["from","1Address",0.5,6,"comment","commentto"],"id":1}`,
			unmarshalled: &ulordjson.SendFromCmd{
				FromAccount: "from",
				ToAddress:   "1Address",
				Amount:      0.5,
				MinConf:     ulordjson.Int(6),
				Comment:     ulordjson.String("comment"),
				CommentTo:   ulordjson.String("commentto"),
			},
		},
		{
			name: "sendmany",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("sendmany", "from", `{"1Address":0.5}`)
			},
			staticCmd: func() interface{} {
				amounts := map[string]float64{"1Address": 0.5}
				return ulordjson.NewSendManyCmd("from", amounts, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendmany","params":["from",{"1Address":0.5}],"id":1}`,
			unmarshalled: &ulordjson.SendManyCmd{
				FromAccount: "from",
				Amounts:     map[string]float64{"1Address": 0.5},
				MinConf:     ulordjson.Int(1),
				Comment:     nil,
			},
		},
		{
			name: "sendmany optional1",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("sendmany", "from", `{"1Address":0.5}`, 6)
			},
			staticCmd: func() interface{} {
				amounts := map[string]float64{"1Address": 0.5}
				return ulordjson.NewSendManyCmd("from", amounts, ulordjson.Int(6), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendmany","params":["from",{"1Address":0.5},6],"id":1}`,
			unmarshalled: &ulordjson.SendManyCmd{
				FromAccount: "from",
				Amounts:     map[string]float64{"1Address": 0.5},
				MinConf:     ulordjson.Int(6),
				Comment:     nil,
			},
		},
		{
			name: "sendmany optional2",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("sendmany", "from", `{"1Address":0.5}`, 6, "comment")
			},
			staticCmd: func() interface{} {
				amounts := map[string]float64{"1Address": 0.5}
				return ulordjson.NewSendManyCmd("from", amounts, ulordjson.Int(6), ulordjson.String("comment"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendmany","params":["from",{"1Address":0.5},6,"comment"],"id":1}`,
			unmarshalled: &ulordjson.SendManyCmd{
				FromAccount: "from",
				Amounts:     map[string]float64{"1Address": 0.5},
				MinConf:     ulordjson.Int(6),
				Comment:     ulordjson.String("comment"),
			},
		},
		{
			name: "sendtoaddress",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("sendtoaddress", "1Address", 0.5)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewSendToAddressCmd("1Address", 0.5, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendtoaddress","params":["1Address",0.5],"id":1}`,
			unmarshalled: &ulordjson.SendToAddressCmd{
				Address:   "1Address",
				Amount:    0.5,
				Comment:   nil,
				CommentTo: nil,
			},
		},
		{
			name: "sendtoaddress optional1",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("sendtoaddress", "1Address", 0.5, "comment", "commentto")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewSendToAddressCmd("1Address", 0.5, ulordjson.String("comment"),
					ulordjson.String("commentto"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendtoaddress","params":["1Address",0.5,"comment","commentto"],"id":1}`,
			unmarshalled: &ulordjson.SendToAddressCmd{
				Address:   "1Address",
				Amount:    0.5,
				Comment:   ulordjson.String("comment"),
				CommentTo: ulordjson.String("commentto"),
			},
		},
		{
			name: "setaccount",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("setaccount", "1Address", "acct")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewSetAccountCmd("1Address", "acct")
			},
			marshalled: `{"jsonrpc":"1.0","method":"setaccount","params":["1Address","acct"],"id":1}`,
			unmarshalled: &ulordjson.SetAccountCmd{
				Address: "1Address",
				Account: "acct",
			},
		},
		{
			name: "settxfee",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("settxfee", 0.0001)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewSetTxFeeCmd(0.0001)
			},
			marshalled: `{"jsonrpc":"1.0","method":"settxfee","params":[0.0001],"id":1}`,
			unmarshalled: &ulordjson.SetTxFeeCmd{
				Amount: 0.0001,
			},
		},
		{
			name: "signmessage",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("signmessage", "1Address", "message")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewSignMessageCmd("1Address", "message")
			},
			marshalled: `{"jsonrpc":"1.0","method":"signmessage","params":["1Address","message"],"id":1}`,
			unmarshalled: &ulordjson.SignMessageCmd{
				Address: "1Address",
				Message: "message",
			},
		},
		{
			name: "signrawtransaction",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("signrawtransaction", "001122")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewSignRawTransactionCmd("001122", nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"signrawtransaction","params":["001122"],"id":1}`,
			unmarshalled: &ulordjson.SignRawTransactionCmd{
				RawTx:    "001122",
				Inputs:   nil,
				PrivKeys: nil,
				Flags:    ulordjson.String("ALL"),
			},
		},
		{
			name: "signrawtransaction optional1",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("signrawtransaction", "001122", `[{"txid":"123","vout":1,"scriptPubKey":"00","redeemScript":"01"}]`)
			},
			staticCmd: func() interface{} {
				txInputs := []ulordjson.RawTxInput{
					{
						Txid:         "123",
						Vout:         1,
						ScriptPubKey: "00",
						RedeemScript: "01",
					},
				}

				return ulordjson.NewSignRawTransactionCmd("001122", &txInputs, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"signrawtransaction","params":["001122",[{"txid":"123","vout":1,"scriptPubKey":"00","redeemScript":"01"}]],"id":1}`,
			unmarshalled: &ulordjson.SignRawTransactionCmd{
				RawTx: "001122",
				Inputs: &[]ulordjson.RawTxInput{
					{
						Txid:         "123",
						Vout:         1,
						ScriptPubKey: "00",
						RedeemScript: "01",
					},
				},
				PrivKeys: nil,
				Flags:    ulordjson.String("ALL"),
			},
		},
		{
			name: "signrawtransaction optional2",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("signrawtransaction", "001122", `[]`, `["abc"]`)
			},
			staticCmd: func() interface{} {
				txInputs := []ulordjson.RawTxInput{}
				privKeys := []string{"abc"}
				return ulordjson.NewSignRawTransactionCmd("001122", &txInputs, &privKeys, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"signrawtransaction","params":["001122",[],["abc"]],"id":1}`,
			unmarshalled: &ulordjson.SignRawTransactionCmd{
				RawTx:    "001122",
				Inputs:   &[]ulordjson.RawTxInput{},
				PrivKeys: &[]string{"abc"},
				Flags:    ulordjson.String("ALL"),
			},
		},
		{
			name: "signrawtransaction optional3",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("signrawtransaction", "001122", `[]`, `[]`, "ALL")
			},
			staticCmd: func() interface{} {
				txInputs := []ulordjson.RawTxInput{}
				privKeys := []string{}
				return ulordjson.NewSignRawTransactionCmd("001122", &txInputs, &privKeys,
					ulordjson.String("ALL"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"signrawtransaction","params":["001122",[],[],"ALL"],"id":1}`,
			unmarshalled: &ulordjson.SignRawTransactionCmd{
				RawTx:    "001122",
				Inputs:   &[]ulordjson.RawTxInput{},
				PrivKeys: &[]string{},
				Flags:    ulordjson.String("ALL"),
			},
		},
		{
			name: "walletlock",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("walletlock")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewWalletLockCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"walletlock","params":[],"id":1}`,
			unmarshalled: &ulordjson.WalletLockCmd{},
		},
		{
			name: "walletpassphrase",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("walletpassphrase", "pass", 60)
			},
			staticCmd: func() interface{} {
				return ulordjson.NewWalletPassphraseCmd("pass", 60)
			},
			marshalled: `{"jsonrpc":"1.0","method":"walletpassphrase","params":["pass",60],"id":1}`,
			unmarshalled: &ulordjson.WalletPassphraseCmd{
				Passphrase: "pass",
				Timeout:    60,
			},
		},
		{
			name: "walletpassphrasechange",
			newCmd: func() (interface{}, error) {
				return ulordjson.NewCmd("walletpassphrasechange", "old", "new")
			},
			staticCmd: func() interface{} {
				return ulordjson.NewWalletPassphraseChangeCmd("old", "new")
			},
			marshalled: `{"jsonrpc":"1.0","method":"walletpassphrasechange","params":["old","new"],"id":1}`,
			unmarshalled: &ulordjson.WalletPassphraseChangeCmd{
				OldPassphrase: "old",
				NewPassphrase: "new",
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
