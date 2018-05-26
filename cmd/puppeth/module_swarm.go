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

// nodeDockerfile is the Dockerfile required to run an Ethereum node.
var swarmDockerfile = `
FROM ethereum/client-go:alltools-v1.8.3

ADD genesis.json /genesis.json
ADD bzzkey.json /bzzkey.json
ADD bzzpass /bzzpass

RUN \
	echo 'mkdir -p /root/.ethereum/keystore/ && cp /bzzkey.json /root/.ethereum/keystore/' > geth.sh && \
	echo $'swarm --bzznetworkid {{.NetworkID}} {{if .BzzAccount}}--bzzaccount {{.BzzAccount}} {{end}}--port {{.Port}} --bzzport {{.bzzPort}} --maxpeers {{.Peers}} {{if .SwarmBoot}}--bootnodes {{.SwarmBoot}}{{end}} --password /bzzpass' >> geth.sh

ENTRYPOINT ["/bin/sh", "geth.sh"]
`

// nodeComposefile is the docker-compose.yml file required to deploy and maintain
// an Ethereum node (bootnode or miner for now).
var swarmComposefile = `
version: '2'
services:
  {{.Type}}:
    build: .
    image: {{.Network}}/{{.Type}}
    ports:
      - "{{.Port}}:{{.Port}}"
      - "{{.Port}}:{{.Port}}/udp"
    volumes:
      - {{.Datadir}}:/root/.ethereum
    environment:
      - PORT={{.bzzPort}}/tcp
    logging:
      driver: "json-file"
      options:
        max-size: "1m"
        max-file: "10"
    restart: always
`

// deploysSwarm deploys a new Swarm node container to a remote machine via SSH,
// docker and docker-compose. If an instance with the specified network name
// already exists there, it will be overwritten!
func deploySwarm(client *sshClient, network string, swarmboot []string, config *swarmInfos, nocache bool) ([]byte, error) {
	kind := "swarmnode"
	if config.swarmenode == "" {
		kind = "swarmboot"
		swarmboot = make([]string, 0)
	}

	// Generate the content to upload to the server
	workdir := fmt.Sprintf("%d", rand.Int63())
	files := make(map[string][]byte)

	dockerfile := new(bytes.Buffer)
	template.Must(template.New("").Parse(swarmDockerfile)).Execute(dockerfile, map[string]interface{}{
		"NetworkID": config.network,
		"Port":      config.port,
		"Peers":     config.peersTotal,
		"bzzPort":   config.bzzPort,
		"SwarmBoot": strings.Join(swarmboot, ","),
		"Unlock":    config.keyJSON != "",
		"BzzAccount":config.bzzAccount,
	})
	files[filepath.Join(workdir, "Dockerfile")] = dockerfile.Bytes()

	composefile := new(bytes.Buffer)
	template.Must(template.New("").Parse(swarmComposefile)).Execute(composefile, map[string]interface{}{
		"Type":       kind,
		"Datadir":    config.datadir,
		"Network":    network,
		"Port":       config.port,
		"TotalPeers": config.peersTotal,
		"BzzAccount": config.bzzAccount,

	})
	files[filepath.Join(workdir, "docker-compose.yaml")] = composefile.Bytes()

	files[filepath.Join(workdir, "genesis.json")] = config.genesis
	if config.keyJSON != "" {
		files[filepath.Join(workdir, "bzzkey.json")] = []byte(config.keyJSON)
		files[filepath.Join(workdir, "bzzpass")] = []byte(config.keyPass)
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
	return nil, client.Stream(fmt.Sprintf("cd %s && docker-compose -p %s up -d --build --force-recreate", workdir, network)) // 설치 로그
}

// nodeInfos is returned from a boot or seal node status check to allow reporting
// various configuration parameters.
type swarmInfos struct {
	genesis    []byte
	network    int64
	datadir    string
	peersTotal int
	port       int
	bzzPort    int
	enode      string
	swarmenode string
	bzzAccount  string
	keyJSON    string
	keyPass    string
}

// Report converts the typed struct into a plain string->string map, containing
// most - but not all - fields for reporting to the user.
func (info *swarmInfos) Report() map[string]string {
	report := map[string]string{
		"Data directory":           info.datadir,
		"Listener port":            strconv.Itoa(info.port),
		"Peer count (all total)":   strconv.Itoa(info.peersTotal),
	}

	report["Bzz account"] = info.bzzAccount

	if info.keyJSON != "" {
		var key struct {
			Address string `json:"address"`
		}
		if err := json.Unmarshal([]byte(info.keyJSON), &key); err == nil {
			report["Bzz account"] = common.HexToAddress(key.Address).Hex()
		} else {
			log.Error("Failed to retrieve bzz address", "err", err)
		}
	}

	return report
}

// checkNode does a health-check against an boot or seal node server to verify
// whether it's running, and if yes, whether it's responsive.
func checkSwarmNode(client *sshClient, network string, boot bool) (*swarmInfos, error) {
	kind := "bootswarm"
	if !boot {
		kind = "swarmnode"
	}
	// Inspect a possible bootnode container on the host
	infos, err := inspectContainer(client, fmt.Sprintf("%s_%s_1", network, kind))
	if err != nil {
		return nil, err
	}
	if !infos.running {
		return nil, ErrServiceOffline
	}
	totalPeers, _ := strconv.Atoi(infos.envvars["TOTAL_PEERS"])

	// Container available, retrieve its node ID and its genesis json
	var out []byte
	if out, err = client.Run(fmt.Sprintf("docker exec %s_%s_1 geth --exec admin.nodeInfo.id attach /root/.ethereum/bzzd.ipc", network, kind)); err != nil {
		return nil, ErrServiceUnreachable
	}

	id := bytes.Trim(bytes.TrimSpace(out), "\"")

	if out, err = client.Run(fmt.Sprintf("docker exec %s_%s_1 cat /genesis.json", network, kind)); err != nil {
		return nil, ErrServiceUnreachable
	}
	genesis := bytes.TrimSpace(out)

	keyJSON, keyPass := "", ""
	if out, err = client.Run(fmt.Sprintf("docker exec %s_%s_1 cat /bzzkey.json", network, kind)); err == nil {
		keyJSON = string(bytes.TrimSpace(out))
	}
	if out, err = client.Run(fmt.Sprintf("docker exec %s_%s_1 cat /bzzpass", network, kind)); err == nil { // 얘가 문제는 아님
		keyPass = string(bytes.TrimSpace(out))
	}
	// Run a sanity check to see if the devp2p is reachable
	port := infos.portmap[infos.envvars["PORT"]]
	if err = checkPort(client.server, port); err != nil {
		log.Warn(fmt.Sprintf("%s devp2p port seems unreachable", strings.Title(kind)), "server", client.server, "port", port, "err", err)
	}

	//Run a sanity check to see if the devp2p is reachable
	//bzzPort := infos.portmap[infos.envvars["BZZPORT"]]
	//if err = checkBzzPort(client.server, bzzPort); err != nil {
	//	log.Warn(fmt.Sprintf("%s bzzp2p port seems unreachable", strings.Title(kind)), "server", client.server, "bzzport", bzzPort, "err", err)
	//}
	// Assemble and return the useful infos
	stats := &swarmInfos{
		genesis:    genesis,
		datadir:    infos.volumes["/root/.ethereum"],
		port:       port,
		//bzzPort:    bzzPort,
		peersTotal: totalPeers,
		keyJSON:    keyJSON,
		bzzAccount: infos.envvars["BZZ_NAME"],
		keyPass:    keyPass,
	}
	stats.swarmenode = fmt.Sprintf("enode://%s@%s:%d", id, client.address, stats.port)
	//fmt.Println(nil)
	return stats, nil
}
