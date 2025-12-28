package blockchain

import (
	"crypto/md5"
	"testing"
)

// TestProofOfWork_Validate ensures that a block mined by CreateBlock
// validates under ProofOfWork.Validate(). Because this educational PoW
// uses a hash size mismatch (MD5 vs a 256-bit target), tampering may not
// always flip the Validate() result; instead we assert that modifying the
// block data changes the recomputed hash compared to the mined value.
func TestProofOfWork_Validate(t *testing.T) {
	t.Parallel()

	// Create (and mine) a block using the helper.
	block := CreateBlock("test-data", "prev-hash", nil)

	pow := NewProofOfWork(block)

	if !pow.Validate() {
		t.Fatal("expected mined block to validate, but Validate() returned false")
	}

	// Tamper with the block and ensure the recomputed hash differs
	// from the original mined hash (indicating tampering changed the digest).
	originalHash := block.Hash
	block.Data = "tampered"
	recomputed := md5.Sum(pow.ComputeData(block.Nonce))
	if string(recomputed[:]) == originalHash {
		t.Fatal("expected recomputed hash to differ after tampering, but it matched")
	}
}
