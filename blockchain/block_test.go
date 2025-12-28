package blockchain

import "testing"

// TestCreateBlock verifies that creating a block sets expected fields
// and that the produced block satisfies its Proof of Work requirement.
func TestCreateBlock(t *testing.T) {
	t.Parallel()

	prevHash := "previous-hash-xyz"
	block := CreateBlock("payload", prevHash, nil)

	if block == nil {
		t.Fatal("CreateBlock returned nil")
	}

	if block.PrevHash != prevHash {
		t.Fatalf("expected PrevHash %q, got %q", prevHash, block.PrevHash)
	}

	if block.Hash == "" {
		t.Fatal("expected non-empty Hash after mining")
	}

	pow := NewProofOfWork(block)
	if !pow.Validate() {
		t.Fatal("mined block did not validate against PoW target")
	}
}
