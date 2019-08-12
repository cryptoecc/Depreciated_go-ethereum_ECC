package eccpow

import (
	"bytes"
	"testing"

	"github.com/Onther-Tech/go-ethereum/common/hexutil"
)

func TestLDPC(t *testing.T) {
	// Create a block to verify
	hash := hexutil.MustDecode("0xc9149cc0386e689d789a1c2f3d5d169a61a6218ed30e74414dc736e442ef3d1f")

	wantDigest := hexutil.MustDecode("0xe4073cffaef931d37117cefd9afd27ea0f1cad6a981dd2605c4a1ac97c519800")
	wantResult := hexutil.MustDecode("0xd3539235ee2e6f8db665c0a72169f55b7f6c605712330b778ec3944f0eb5a557")

	//digest, result := hashimotoLight(32*1024, cache, hash, nonce)
	if !bytes.Equal(digest, wantDigest) {
		t.Errorf("light hashimoto digest mismatch: have %x, want %x", digest, wantDigest)
	}
	if !bytes.Equal(result, wantResult) {
		t.Errorf("light hashimoto result mismatch: have %x, want %x", result, wantResult)
	}
	nonce = runLDPC(hash, wantDigest)
	if !bytes.Equal(digest, wantDigest) {
		t.Errorf("full hashimoto digest mismatch: have %x, want %x", digest, wantDigest)
	}
	if !bytes.Equal(result, wantResult) {
		t.Errorf("full hashimoto result mismatch: have %x, want %x", result, wantResult)
	}
}
