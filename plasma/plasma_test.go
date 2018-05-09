// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package plasma

// import (
// 	// "bytes"
// 	// "crypto/ecdsa"
// 	// "time"
// 	// "github.com/ethereum/go-ethereum/common"
//
// 	"testing"
// )
//
// func TestPlasmaBasic(t *testing.T) {
// 	w := New(&DefaultConfig)
// 	p := w.Protocols()
// 	pls := p[0]
// 	if pls.Name != ProtocolName {
// 		t.Fatalf("failed Protocol Name: %v.", pls.Name)
// 	}
// 	if uint64(pls.Version) != ProtocolVersion {
// 		t.Fatalf("failed Protocol Version: %v.", pls.Version)
// 	}
// 	if pls.Length != NumberOfMessageCodes {
// 		t.Fatalf("failed Protocol Length: %v.", pls.Length)
// 	}
// 	if pls.Run == nil {
// 		t.Fatalf("failed pls.Run.")
// 	}
// 	if uint64(w.Version()) != ProtocolVersion {
// 		t.Fatalf("failed plasma Version: %v.", pls.Version)
// 	}
// }
//
// func TestPlasmaBasic(t *testing.T) {
// 	w := New(&DefaultConfig)
// 	p := w.Protocols()
// 	pls := p[0]
// 	if pls.Name != ProtocolName {
// 		t.Fatalf("failed Protocol Name: %v.", pls.Name)
// 	}
// 	if uint64(pls.Version) != ProtocolVersion {
// 		t.Fatalf("failed Protocol Version: %v.", pls.Version)
// 	}
// 	if pls.Length != NumberOfMessageCodes {
// 		t.Fatalf("failed Protocol Length: %v.", pls.Length)
// 	}
// 	if pls.Run == nil {
// 		t.Fatalf("failed pls.Run.")
// 	}
// 	if uint64(w.Version()) != ProtocolVersion {
// 		t.Fatalf("failed plasma Version: %v.", pls.Version)
// 	}
// }
