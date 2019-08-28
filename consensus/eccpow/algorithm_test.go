package eccpow

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/Onther-Tech/go-ethereum/common"
	"github.com/Onther-Tech/go-ethereum/common/hexutil"
)

func TestRandomSeed(t *testing.T) {
	prev_hash := hexutil.MustDecode("0xc9149cc0386e689d789a1c2f3d5d169a61a6218ed30e74414dc736e442ef3d1f")
	parameter := SetDifficultyUsingLevel(0)
	GenerateSeed(prev_hash)

	a := GenerateH(parameter)
	b := GenerateH(parameter)

	if !reflect.DeepEqual(a, b) {
		t.Error("Wrong matrix")
	} else {
		t.Log("Pass")
	}
}

func TestLDPC(t *testing.T) {
	prevHash := hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000")
	curHash := hexutil.MustDecode("0xca2ff06caae7c94dc968be7d76d0fbf60dd2e1989ee9bf0d5931e48564d5143b")
	nonce, mixDigest := RunLDPC(prevHash, curHash)

	wantDigest := hexutil.MustDecode("0x535306ee4b42c92aecd0e71fca98572064f049c2babb2769faa3bbd87d67ec2d")

	if !bytes.Equal(mixDigest, wantDigest) {
		t.Errorf("light hashimoto digest mismatch: have %x, want %x", mixDigest, wantDigest)
	}

	t.Log(nonce)
}

func BenchmarkECCPoW(b *testing.B) {
	prevHash := hexutil.MustDecode("0xd783efa4d392943503f28438ad5830b2d5964696ffc285f338585e9fe0a37a05")
	curHash := hexutil.MustDecode("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RunLDPC(prevHash, curHash)
	}

}

func TestHashRate(t *testing.T) {
	var (
		hashrate = []hexutil.Uint64{100, 200, 300}
		expect   uint64
		ids      = []common.Hash{common.HexToHash("a"), common.HexToHash("b"), common.HexToHash("c")}
	)
	ecc := NewTester(nil, false)
	defer ecc.Close()

	if tot := ecc.Hashrate(); tot != 0 {
		t.Error("expect the result should be zero")
	}

	api := &API{ecc}
	for i := 0; i < len(hashrate); i += 1 {
		if res := api.SubmitHashRate(hashrate[i], ids[i]); !res {
			t.Error("remote miner submit hashrate failed")
		}
		expect += uint64(hashrate[i])
	}
	if tot := ecc.Hashrate(); tot != float64(expect) {
		t.Error("expect total hashrate should be same")
	}
}
