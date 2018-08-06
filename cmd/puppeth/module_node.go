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
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// ToDo get minDeposit, recoveryEpoch, withdrawalDelay, rpcPort parameter
var staminaScript =`
const Web3 = require('web3');

const web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:8545'));

const owner = "{{.Owner}}";

const delegator = "{{.Delegator}}";
const delegatee = "{{.Delegatee}}";

const staminaAddr = "0x000000000000000000000000000000000000dead";
const Stamina = web3.eth.contract([{"constant":true,"inputs":[],"name":"WITHDRAWAL_DELAY","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"RECOVER_EPOCH_LENGTH","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"initialized","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"development","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"MIN_DEPOSIT","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"depositor","type":"address"},{"indexed":true,"name":"delegatee","type":"address"},{"indexed":false,"name":"amount","type":"uint256"}],"name":"Deposited","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"delegator","type":"address"},{"indexed":false,"name":"oldDelegatee","type":"address"},{"indexed":false,"name":"newDelegatee","type":"address"}],"name":"DelegateeChanged","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"depositor","type":"address"},{"indexed":true,"name":"delegatee","type":"address"},{"indexed":false,"name":"amount","type":"uint256"},{"indexed":false,"name":"requestBlockNumber","type":"uint256"},{"indexed":false,"name":"withdrawalIndex","type":"uint256"}],"name":"WithdrawalRequested","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"depositor","type":"address"},{"indexed":true,"name":"delegatee","type":"address"},{"indexed":false,"name":"amount","type":"uint256"},{"indexed":false,"name":"withdrawalIndex","type":"uint256"}],"name":"Withdrawn","type":"event"},{"constant":false,"inputs":[{"name":"minDeposit","type":"uint256"},{"name":"recoveryEpochLength","type":"uint256"},{"name":"withdrawalDelay","type":"uint256"}],"name":"init","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"delegator","type":"address"}],"name":"getDelegatee","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"addr","type":"address"}],"name":"getStamina","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"delegatee","type":"address"}],"name":"getTotalDeposit","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"depositor","type":"address"},{"name":"delegatee","type":"address"}],"name":"getDeposit","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"depositor","type":"address"}],"name":"getNumWithdrawals","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"delegatee","type":"address"}],"name":"getLastRecoveryBlock","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"delegatee","type":"address"}],"name":"getNumRecovery","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"depositor","type":"address"},{"name":"withdrawalIndex","type":"uint256"}],"name":"getWithdrawal","outputs":[{"name":"amount","type":"uint128"},{"name":"requestBlockNumber","type":"uint128"},{"name":"delegatee","type":"address"},{"name":"processed","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"delegator","type":"address"}],"name":"setDelegator","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"delegatee","type":"address"}],"name":"deposit","outputs":[{"name":"","type":"bool"}],"payable":true,"stateMutability":"payable","type":"function"},{"constant":false,"inputs":[{"name":"delegatee","type":"address"},{"name":"amount","type":"uint256"}],"name":"requestWithdrawal","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"withdraw","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"delegatee","type":"address"},{"name":"amount","type":"uint256"}],"name":"addStamina","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"delegatee","type":"address"},{"name":"amount","type":"uint256"}],"name":"subtractStamina","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"}])

const stamina = Stamina.at(staminaAddr);

web3.personal.unlockAccount(owner, "{{.OwnerPass}}", 0);
web3.personal.unlockAccount(delegatee, "{{.DelegateePass}}", 0);


web3.eth.sendTransaction({from: owner, to: delegatee, value:1e18})
web3.eth.sendTransaction({from: owner, to: delegator, value:1e18})

const minDeposit = 1e17;
const recoveryEpoch = 10;
const withdrawalDelay = 30;

const main = async() => {
  await stamina.init(minDeposit, recoveryEpoch, withdrawalDelay, {from: owner, gas: 2e6});

  await stamina.setDelegator(delegator, {from: delegatee});
  
  await stamina.deposit(delegatee, {from: owner, value: 1e18, gas: 2e6});

  setTimeout(function() {
	web3.eth.sendTransaction({from:delegator, to:delegatee, value:5e17});
  });
  
  console.log("Set stamina complete.")  
};

setTimeout(function() {
  main().catch(console.error);
}, 10000);

`


// nodeDockerfile is the Dockerfile required to run an Ethereum node.
// ToDo. To change installing stamina optionally
var nodeDockerfile = `
FROM onther/ethereum-client:latest


ADD genesis.json /genesis.json
{{if .Unlock}}
	ADD zdelegatee.json /zdelegatee.json
	
	
	ADD signer.json /signer.json
	ADD signer.pass /signer.pass


	RUN mkdir stamina
	RUN apk add --update git nodejs-npm
	ADD app.js /stamina/app.js
	ADD package.json /stamina/package.json
	RUN cd stamina && npm install
	RUN cd ..
{{end}}
RUN \
  echo 'geth --cache 512 init /genesis.json' > geth.sh && \{{if .Unlock}}
	echo 'mkdir -p /root/.ethereum/keystore/ && cp /zdelegatee.json /root/.ethereum/keystore/ && cp /signer.json /root/.ethereum/keystore/' >> geth.sh && \{{end}}
	echo $'geth --networkid {{.NetworkID}} --cache 512 --port {{.Port}} --maxpeers {{.Peers}} {{.LightFlag}} --ethstats \'{{.Ethstats}}\' {{if .Bootnodes}}--bootnodes {{.Bootnodes}}{{end}} {{if .Etherbase}}--etherbase {{.Etherbase}} --mine --minerthreads 1{{end}} {{if .Signer}}--etherbase {{.Signer}} {{end}} {{if .Unlock}}--unlock 0 --password /signer.pass --mine --rpc --rpcapi personal,eth,net,web3,admin,txpool --rpcport 8545 --rpcaddr=0.0.0.0 --rpccorsdomain "*" {{end}} --targetgaslimit {{.GasTarget}} --gasprice {{.GasPrice}}' >> geth.sh

ENTRYPOINT ["/bin/sh", "geth.sh"]
`

var packageJson =`
{
  "name": "stamina",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
  "test": "echo \"Error: no test specified\" && exit 1"
  },
  "author": "",
  "license": "ISC",
  "dependencies": {
    "web3": "0.20.0"
  }
}
`

// nodeComposefile is the docker-compose.yml file required to deploy and maintain
// an Ethereum node (bootnode or miner for now).
var nodeComposefile = `
version: '2'
services:
  {{.Type}}:
    build: .
    image: {{.Network}}/{{.Type}}
    ports:
      - "{{.Port}}:{{.Port}}"
      - "{{.Port}}:{{.Port}}/udp"
    volumes:
      - {{.Datadir}}:/root/.ethereum{{if .Ethashdir}}
      - {{.Ethashdir}}:/root/.ethash{{end}}
    environment:
      - PORT={{.Port}}/tcp
      - TOTAL_PEERS={{.TotalPeers}}
      - LIGHT_PEERS={{.LightPeers}}
      - STATS_NAME={{.Ethstats}}
      - MINER_NAME={{.Etherbase}}
      - GAS_TARGET={{.GasTarget}}
      - GAS_PRICE={{.GasPrice}}
    logging:
      driver: "json-file"
      options:
        max-size: "1m"
        max-file: "10"
    restart: always
`

// deployNode deploys a new Ethereum node container to a remote machine via SSH,
// docker and docker-compose. If an instance with the specified network name
// already exists there, it will be overwritten!
func deployNode(client *sshClient, network string, bootnodes []string, config *nodeInfos, nocache bool) ([]byte, error) {
	kind := "sealnode"
	if config.keyJSON == "" && config.etherbase == "" {
		kind = "bootnode"
		bootnodes = make([]string, 0)
	}

	// Generate the content to upload to the server
	workdir := fmt.Sprintf("%d", rand.Int63())
	files := make(map[string][]byte)

	lightFlag := ""
	if config.peersLight > 0 {
		lightFlag = fmt.Sprintf("--lightpeers=%d --lightserv=50", config.peersLight)
	}

	var signer string
	if config.keyJSON != "" {
		// Clique proof-of-authority signer
		var key struct {
			Address string `json:"address"`
		}
		if err := json.Unmarshal([]byte(config.keyJSON), &key); err == nil {
			signer = common.HexToAddress(key.Address).Hex()
		} else {
			log.Error("Failed to retrieve signer address", "err", err)
		}
	}

	fmt.Sprintf(config.etherbase)

	dockerfile := new(bytes.Buffer)
	template.Must(template.New("").Parse(nodeDockerfile)).Execute(dockerfile, map[string]interface{}{
		"NetworkID": config.network,
		"Port":      config.port,
		"Peers":     config.peersTotal,
		"LightFlag": lightFlag,
		"Bootnodes": strings.Join(bootnodes, ","),
		"Ethstats":  config.ethstats,
		"Etherbase": config.etherbase,
		"GasTarget": uint64(1000000 * config.gasTarget),
		"GasPrice":  uint64(1000000000 * config.gasPrice),
		"Unlock":    config.keyJSON != "",
		"Signer":    signer,
		//"OwnerChek" : ownerCheck != "",
		//"DelegateeCheck": delegateeCheck != "",
	})
	files[filepath.Join(workdir, "Dockerfile")] = dockerfile.Bytes()

	composefile := new(bytes.Buffer)
	template.Must(template.New("").Parse(nodeComposefile)).Execute(composefile, map[string]interface{}{
		"Type":       kind,
		"Datadir":    config.datadir,
		"Ethashdir":  config.ethashdir,
		"Network":    network,
		"Port":       config.port,
		"TotalPeers": config.peersTotal,
		"Light":      config.peersLight > 0,
		"LightPeers": config.peersLight,
		"Ethstats":   config.ethstats[:strings.Index(config.ethstats, ":")],
		"Etherbase":  config.etherbase,
		"GasTarget":  config.gasTarget,
		"GasPrice":   config.gasPrice,
	})
	files[filepath.Join(workdir, "docker-compose.yaml")] = composefile.Bytes()

	var owner string
	if config.owner != "" {
		var key struct {
			Address string `json:"address"`
		}
		if err := json.Unmarshal([]byte(config.owner), &key); err == nil {
			owner = common.HexToAddress(key.Address).Hex()
		} else {
			log.Error("Failed to retrieve delegatee address", "err", err)
		}
	}

	var delegatee string
	if config.delegatee != "" {
		var key struct {
			Address string `json:"address"`
		}
		if err := json.Unmarshal([]byte(config.delegatee), &key); err == nil {
			delegatee = common.HexToAddress(key.Address).Hex()
		} else {
			log.Error("Failed to retrieve delegatee address", "err", err)
		}
	}

	scriptfile := new(bytes.Buffer)
	template.Must(template.New("").Parse(staminaScript)).Execute(scriptfile, map[string]interface{}{
		"Owner": owner,
		"Delegator": config.delegator,
		"Delegatee": delegatee,
		"OwnerPass": config.keyPass,
		"DelegateePass": config.delegateePass,
	})
	files[filepath.Join(workdir, "app.js")] = scriptfile.Bytes()
	files[filepath.Join(workdir, "package.json")] = []byte(packageJson)

	//var ownerCheck string = ""
	//var delegateeCheck string  = ""
	files[filepath.Join(workdir, "genesis.json")] = config.genesis
	if config.keyJSON != "" {
		files[filepath.Join(workdir, "signer.json")] = []byte(config.keyJSON)
		files[filepath.Join(workdir, "signer.pass")] = []byte(config.keyPass)
		files[filepath.Join(workdir, "zdelegatee.json")] = []byte(config.delegatee)

		//if config.keyJSON == config.delegatee { // 1. signer와 delegatee가 같을때, owner key파일 추가
		//	config.delegatee = config.keyJSON
		//	files[filepath.Join(workdir, "zowner.pass")] = []byte(config.owner)
		//	ownerCheck = "checked"
		//} else if config.keyJSON == config.owner { // 2. signer와 owner가 같을때, delegatee key 파일 추가
		//	config.owner = config.keyJSON
		//	files[filepath.Join(workdir, "zdelegatee.json")] = []byte(config.delegatee)
		//	delegateeCheck = "checked"
		//} else if config.keyJSON == config.owner && config.keyJSON == config.delegatee { // 3. 모두 같은 키를 쓸 때
		//	config.delegatee = config.keyJSON
		//	config.owner = config.keyJSON
		//} else {
		//	files[filepath.Join(workdir, "zowner.pass")] = []byte(config.owner)
		//	ownerCheck = "checked"
		//	files[filepath.Join(workdir, "zdelegatee.json")] = []byte(config.delegatee)
		//	delegateeCheck = "checked"
		//}
	}

	// Upload the deployment files to the remote server (and clean up afterwards)
	if out, err := client.Upload(files); err != nil {
		return out, err
	}
	defer client.Run("rm -rf " + workdir)

	// Build and deploy the boot or seal node service
	if nocache {
		return nil, client.Stream(fmt.Sprintf("cd %s && docker-compose -p %s build --pull --no-cache && docker-compose -p %s up -d --force-recreate", workdir, network, network))
	}
	return nil, client.Stream(fmt.Sprintf("cd %s && docker-compose -p %s up -d --build --force-recreate", workdir, network))

}

// nodeInfos is returned from a boot or seal node status check to allow reporting
// various configuration parameters.
type nodeInfos struct {
	genesis    []byte
	network    int64
	datadir    string
	ethashdir  string
	ethstats   string
	port       int
	enode      string
	peersTotal int
	peersLight int
	etherbase  string
	keyJSON    string
	keyPass    string
	gasTarget  float64
	gasPrice   float64

	owner         string
	ownerPass     string

	delegator     string
	delegatee     string
	delegateePass string
}

// Report converts the typed struct into a plain string->string map, containing
// most - but not all - fields for reporting to the user.
func (info *nodeInfos) Report() map[string]string {
	report := map[string]string{
		"Data directory":           info.datadir,
		"Listener port":            strconv.Itoa(info.port),
		"Peer count (all total)":   strconv.Itoa(info.peersTotal),
		"Peer count (light nodes)": strconv.Itoa(info.peersLight),
		"Ethstats username":        info.ethstats,
	}

	if info.gasTarget > 0 {
		// Miner or signer node
		report["Gas limit (baseline target)"] = fmt.Sprintf("%0.3f MGas", info.gasTarget)
		report["Gas price (minimum accepted)"] = fmt.Sprintf("%0.3f GWei", info.gasPrice)

		if info.etherbase != "" {
			// Ethash proof-of-work miner
			report["Ethash directory"] = info.ethashdir
			report["Miner account"] = info.etherbase
		}
		if info.keyJSON != "" {
			// Clique proof-of-authority signer
			var key struct {
				Address string `json:"address"`
			}
			if err := json.Unmarshal([]byte(info.keyJSON), &key);err == nil {
				report["Signer account"] = common.HexToAddress(key.Address).Hex()
			} else {
				log.Error("Failed to retrieve signer address", "err", err)
			}
		}

		if info.delegatee != "" {
			var key struct {
				Address string `json:"address"`
			}
			if err := json.Unmarshal([]byte(info.delegatee), &key);err == nil {
				report["Delegatee"] = common.HexToAddress(key.Address).Hex()
			} else {
				log.Error("Failed to retrieve delegatee address", "err", err)
			}
		}

		if info.delegator != "" {
			report["Delegator"] = info.delegator
		}
	}
	return report
}

var checkStamina = false // check if state of stamina contract

// checkNode does a health-check against a boot or seal node server to verify
// whether it's running, and if yes, whether it's responsive.
func checkNode(client *sshClient, network string, boot bool) (*nodeInfos, error) {
	kind := "bootnode"
	if !boot {
		kind = "sealnode"
	}
	// Inspect a possible bootnode container on the host
	infos, err := inspectContainer(client, fmt.Sprintf("%s_%s_1", network, kind))
	if err != nil {
		return nil, err
	}
	if !infos.running {
		return nil, ErrServiceOffline
	}
	// Resolve a few types from the environmental variables
	totalPeers, _ := strconv.Atoi(infos.envvars["TOTAL_PEERS"])
	lightPeers, _ := strconv.Atoi(infos.envvars["LIGHT_PEERS"])
	gasTarget, _ := strconv.ParseFloat(infos.envvars["GAS_TARGET"], 64)
	gasPrice, _ := strconv.ParseFloat(infos.envvars["GAS_PRICE"], 64)

	// Container available, retrieve its node ID and its genesis json
	var out []byte
	if out, err = client.Run(fmt.Sprintf("docker exec %s_%s_1 geth --exec admin.nodeInfo.id --cache=16 attach", network, kind)); err != nil {
		return nil, ErrServiceUnreachable
	}
	id := bytes.Trim(bytes.TrimSpace(out), "\"")


	if out, err = client.Run(fmt.Sprintf("docker exec %s_%s_1 cat /genesis.json", network, kind)); err != nil {
		return nil, ErrServiceUnreachable
	}
	genesis := bytes.TrimSpace(out)

	//defer client.Run(fmt.Sprintf("docker exec %s_%s_1 node stamina/app.js", network, kind))

	if kind == "sealnode" && checkStamina != true {
		client.Run(fmt.Sprintf("docker exec %s_%s_1 node stamina/app.js", network, kind))
		checkStamina = true
	}

	keyJSON, keyPass := "", ""
	if out, err = client.Run(fmt.Sprintf("docker exec %s_%s_1 cat /signer.json", network, kind)); err == nil {
		keyJSON = string(bytes.TrimSpace(out))
	}
	if out, err = client.Run(fmt.Sprintf("docker exec %s_%s_1 cat /signer.pass", network, kind)); err == nil {
		keyPass = string(bytes.TrimSpace(out))
	}
	// Run a sanity check to see if the devp2p is reachable
	port := infos.portmap[infos.envvars["PORT"]]
	if err = checkPort(client.server, port); err != nil {
		log.Warn(fmt.Sprintf("%s devp2p port seems unreachable", strings.Title(kind)), "server", client.server, "port", port, "err", err)
	}
	// Assemble and return the useful infos
	stats := &nodeInfos{
		genesis:    genesis,
		datadir:    infos.volumes["/root/.ethereum"],
		ethashdir:  infos.volumes["/root/.ethash"],
		port:       port,
		peersTotal: totalPeers,
		peersLight: lightPeers,
		ethstats:   infos.envvars["STATS_NAME"],
		etherbase:  infos.envvars["MINER_NAME"],
		keyJSON:    keyJSON,
		keyPass:    keyPass,
		gasTarget:  gasTarget,
		gasPrice:   gasPrice,
		delegator:  infos.envvars["DELEGATOR_NAME"],
		delegatee:  infos.envvars["DELEGATEE_NAME"],
	}
	stats.enode = fmt.Sprintf("enode://%s@%s:%d", id, client.address, stats.port)

	return stats, nil
}
