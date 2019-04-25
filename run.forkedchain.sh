make geth

mkdir chain && mkdir chain/1 && mkdir chain/2

trap 'terminate' SIGINT

terminate() {
  killall -HUP geth
  rm -rf chain
}

build/bin/geth --datadir chain/1 init genesis.json && geth --datadir chain/2 init genesis.json
build/bin/geth --datadir chain/1 --rpc --rpcport 8545 --ws --wsport 8546 --rpccorsdomain "*" --rpcapi "admin,eth,miner,net,personal,web3" --nodiscover --mine --gasprice 1 --minerthreads=1 --targetgaslimit 8000000 --etherbase 0x71562b71999873DB5b286dF957af199Ec94617F7 --fakepow &
build/bin/geth --datadir chain/2 --port 30304 --rpc --rpcport 8645 --ws --wsport 8646 --rpccorsdomain "*" --rpcapi "admin,eth,miner,net,personal,web3" --nodiscover --mine --gasprice 1000000000000000000 --minerthreads=8 --targetgaslimit 8000000 --etherbase 0x3616be06d68dd22886505e9c2caaa9eca84564b8 --fakepow --nodekeyhex 5e048defbb082a83db7aa0c508696b70ecfb25de3973d7ea9cd5eeea2d295c17
