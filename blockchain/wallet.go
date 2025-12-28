// Package blockchain provides a minimal educational blockchain implementation.
// This file contains wallet helper functions for key generation, signing and
// verification of simple transactions used by the demo application.
package blockchain

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
)

// Wallet holds an RSA private/public keypair for signing transactions.
// In a production system wallets use secure key storage and deterministic
// key derivation; here we generate ephemeral RSA keys for simplicity.
type Wallet struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// GenerateRSAKeys generates a new RSA private/public keypair.
// Returns the generated keys or an error if key generation fails.
func GenerateRSAKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		return nil, nil, err
	}

	publicKey := &privateKey.PublicKey

	return privateKey, publicKey, nil
}

// NewWallet creates a new Wallet with a freshly generated RSA keypair.
func NewWallet() (*Wallet, error) {
	privateKey, publicKey, err := GenerateRSAKeys()

	if err != nil {
		return nil, err
	}

	return &Wallet{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

// SignTransaction signs the provided transaction using the wallet's private key.
// The transaction is converted to a reproducible string form, hashed with SHA-256
// and then signed. The signature is returned as a base64 encoded string.
func (wallet *Wallet) SignTransaction(transaction *Transaction) (string, error) {
	dataString := fmt.Sprintf("%s%s%f%t", transaction.Sender, transaction.Receiver, transaction.Amount, transaction.Coinbase)

	hashedData := sha256.Sum256([]byte(dataString))

	signature, err := rsa.SignPKCS1v15(rand.Reader, wallet.PrivateKey, crypto.SHA256, hashedData[:])

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

// VerifyTransaction verifies that the provided signature is valid for the
// given transaction and public key. Returns nil on success, or an error when
// verification fails.
func VerifyTransaction(transaction *Transaction, publicKey *rsa.PublicKey, signature string) error {
	dataString := fmt.Sprintf("%s%s%f%t", transaction.Sender, transaction.Receiver, transaction.Amount, transaction.Coinbase)

	hashedData := sha256.Sum256([]byte(dataString))

	signatureBytes, err := base64.StdEncoding.DecodeString(signature)

	if err != nil {
		return err
	}

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashedData[:], signatureBytes)

	if err != nil {
		return errors.New("Transaction Signature not valid.")
	}
	return nil
}
