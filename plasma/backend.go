package plasma

import (
  "github.com/ethereum/go-ethereum/rpc"
)

// Backend interface provides the common API services
type Backend interface {
  // General Plasma APIs
  GetTransaction()
  GetCurrentBlock()
  GetCurrentBlockNum()
  GetBlock()
  ApplyTransaction()

  // Public Root chain APIs
  SubmitDeposit()
  StartExit()
  ChallangeExit()

  // Plasma Operator APIs
  SubmitBlock()
}

// GetAPIs returns available plasma related apis
func GetAPIs(apiBackend Backend) []rpc.API {
  return []rpc.API{
    {
      Namespace: "pls",
      Version:   "1.0",
      Service:   NewPlasmaAPI,
      Public:    true,
    },
    {
      Namespace: "pls",
      Version:   "1.0",
      Service:   NewPlasmaOperatorAPI,
      Public:    true,
    },
  }
}
