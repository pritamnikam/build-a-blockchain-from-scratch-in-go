package blockchain

import "testing"

// TestSignVerify ensures that signing a transaction with a wallet's private key
// produces a signature that VerifyTransaction can validate, and that modifying
// the transaction causes verification to fail.
func TestSignVerify(t *testing.T) {
	t.Parallel()

	wallet, err := NewWallet()
	if err != nil {
		t.Fatalf("NewWallet error: %v", err)
	}

	tx := &Transaction{
		Sender:   "Alice",
		Receiver: "Bob",
		Amount:   1.5,
		Coinbase: false,
	}

	sig, err := wallet.SignTransaction(tx)
	if err != nil {
		t.Fatalf("SignTransaction error: %v", err)
	}

	if err := VerifyTransaction(tx, wallet.PublicKey, sig); err != nil {
		t.Fatalf("VerifyTransaction failed: %v", err)
	}

	// Tamper with the transaction and ensure verification fails
	tx.Amount = 2.0
	if err := VerifyTransaction(tx, wallet.PublicKey, sig); err == nil {
		t.Fatal("expected VerifyTransaction to fail for tampered transaction, but it succeeded")
	}
}
