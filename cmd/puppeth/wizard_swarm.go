// Copyright 2017 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// deploySwarm creates a new node configuration based on some user input.
func (w *wizard) deploySwarm(boot bool) {
	// Do some sanity check before the user wastes time on input
	if w.conf.Genesis == nil {
		log.Error("No genesis block configured")
		return
	}

	// Select the server to interact with
	server := w.selectServer()
	if server == "" {
		return
	}
	client := w.servers[server]

	// Retrieve any active node configurations from the server
	infos, err := checkSwarmNode(client, w.network, boot)
	if err != nil {
		if boot {
			infos = &swarmInfos{port: 30399, peersTotal: 50, bzzPort:8500}
		} else {
			infos = &swarmInfos{port: 30399, peersTotal: 50, bzzPort:8500}
		}
	}
	existed := err == nil

	infos.genesis, _ = json.MarshalIndent(w.conf.Genesis, "", "  ")
	infos.network = w.conf.Genesis.Config.ChainId.Int64()

	// Figure out where the user wants to store the persistent data
	fmt.Println()
	if infos.datadir == "" {
		fmt.Printf("Where should data be stored on the remote machine?\n")
		infos.datadir = w.readString()
	} else {
		fmt.Printf("Where should data be stored on the remote machine? (default = %s)\n", infos.datadir)
		infos.datadir = w.readDefaultString(infos.datadir)
	}

	// Figure out which port to listen on
	fmt.Println()
	fmt.Printf("Which TCP/UDP port to listen on? (default = %d)\n", infos.port)
	infos.port = w.readDefaultInt(infos.port)

	// Figure out which bzzport to listen on
	fmt.Println()
	fmt.Printf("Which bzzport to listen on? (default = %d)\n", infos.bzzPort)
	infos.bzzPort = w.readDefaultInt(infos.bzzPort)

	// Figure out how many peers to allow (different based on node type)
	fmt.Println()
	fmt.Printf("How many peers to allow connecting? (default = %d)\n", infos.peersTotal)
	infos.peersTotal = w.readDefaultInt(infos.peersTotal)

	if infos.keyJSON == "" {
		fmt.Println()
		fmt.Println("Please paste the bzzaccount's key JSON:")
		infos.keyJSON = w.readJSON()

		fmt.Println()
		fmt.Println("What's the unlock password for the account? (won't be echoed)")
		infos.keyPass = w.readPassword()

		if _, err := keystore.DecryptKey([]byte(infos.keyJSON), infos.keyPass); err != nil {
			log.Error("Failed to decrypt key with given passphrase")
			return
		}
	}

	fmt.Println()
	if infos.bzzAccount == "" {
		fmt.Println("Please paste the bzzaccount:")
		for {
			if address := w.readAddress(); address != nil {
				infos.bzzAccount = address.Hex()
				break
			}
		}
	} else {
		fmt.Printf("Please paste the bzzaccount (default = %s)\n", infos.bzzAccount)
		infos.bzzAccount = w.readDefaultAddress(common.HexToAddress(infos.bzzAccount)).Hex()
	}

	// Try to deploy the full node on the host
	nocache := false
	if existed {
		fmt.Println()
		fmt.Printf("Should the node be built from scratch (y/n)? (default = no)\n")
		nocache = w.readDefaultString("n") != "n"
	}
	if out, err := deploySwarm(client, w.network, w.conf.bootnodes, infos, nocache); err != nil {
		log.Error("Failed to deploy Swarm node container", "err", err)
		if len(out) > 0 {
			fmt.Printf("%s\n", out)
		}
		return
	}
	// All ok, run a network scan to pick any changes up
	log.Info("Waiting for node to finish booting")
	time.Sleep(3 * time.Second)

	w.networkStats()
}
