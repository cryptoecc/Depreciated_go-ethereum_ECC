package eccpow

import (
	"reflect"
	"testing"

	"github.com/Onther-Tech/go-ethereum/common/hexutil"
)

func TestRandomSeed(t *testing.T) {
	prev_hash := hexutil.MustDecode("0xc9149cc0386e689d789a1c2f3d5d169a61a6218ed30e74414dc736e442ef3d1f")
	n := 24
	wc := 3
	wr := 6
	m := set_difficulty(n, wc, wr)

	ecc := ECC{
		H:             newIntMatrix(m, n),
		col_in_row:    newIntMatrix(wr, m),
		row_in_col:    newIntMatrix(wc, n),
		hashVector:    make([]int, n),
		tmpHashVector: make([]byte, n),
		outputWord:    make([]int, n),
		LRqtl:         newFloatMatrix(n, m),
		LRrtl:         newFloatMatrix(n, m),
		LRpt:          make([]float64, n),
		LRft:          make([]float64, n),
	}

	generateSeed(prev_hash)
	a := ecc.generateH()
	b := ecc.generateH()

	if !reflect.DeepEqual(a, b) {
		t.Error("Wrong matrix")
	} else {
		t.Log("Pass")
	}
	//t.Log(a)
	//for i := 0; i < len(a); i++ {
	//	for j:=0; j < len(a[i]); j++{
	//		if a[i][j] != b[i][j] {
	//			t.Log(a[i][j])
	//			t.Error("Wrong matrix")
	//		} else {
	//			t.Log(" true")
	//		}
	//	}
	//}
}

func TestLDPC(t *testing.T) {
	// Create a block to verify
	prev_hash := hexutil.MustDecode("0x3e140b0784516af5e5ec6730f2fb20cca22f32be399b9e4ad77d32541f798cd0")
	cur_hash := hexutil.MustDecode("0xc9149cc0386e689d789a1c2f3d5d169a61a6218ed30e74414dc736e442ef3d1f")

	nonce := runLDPC(prev_hash, cur_hash, 48, 3, 6)

	t.Log(nonce)
}

func BenchmarkECCPoW(b *testing.B) {
	prev_hash := hexutil.MustDecode("0x3e140b0784516af5e5ec6730f2fb20cca22f32be399b9e4ad77d32541f798cd0")
	cur_hash := hexutil.MustDecode("0xc9149cc0386e689d789a1c2f3d5d169a61a6218ed30e74414dc736e442ef3d1f")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		runLDPC(prev_hash, cur_hash, 24, 3, 6)
	}

}
