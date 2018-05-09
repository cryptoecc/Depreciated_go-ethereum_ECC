package plasma

//
//
// import (
// 	"bytes"
// 	"crypto/ecdsa"
// 	"fmt"
// 	"net"
// 	"sync"
// 	"testing"
// 	"time"
//
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/p2p"
// 	"github.com/ethereum/go-ethereum/p2p/discover"
// 	"github.com/ethereum/go-ethereum/p2p/nat"
// )
//
// const NumNodes = 16 // must not exceed the number of keys (32)
//
// type TestNode struct {
// 	pls    *Plasma
// 	server *p2p.Server
//   id      *ecdsa.PrivateKey
// }
//
// var nodes [NumNodes]*TestNode
//
