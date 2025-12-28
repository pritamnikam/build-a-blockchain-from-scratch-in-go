# Task 13: Add Wallet Functionality
In blockchain, a wallet is a digital container that stores the cryptographic keys used to access and manage a user’s assets. It consists of a private key and a public key. The public key is shared within the blockchain network and is mainly used to receive funds and verify the authenticity of digital signatures. With a public key, anyone can encrypt data that can only be decrypted by the corresponding private key. The private key is a secret cryptographic key that should be kept secure and known only to the owner. It creates digital signatures, signs transactions, and decrypts data encrypted with the corresponding public key.

You will use RSA for key generation, sign transactions, and verify transaction signatures. You’ll implement a wallet structure, generate public and private keys, and create a new wallet based on the generated keys.

In this task, perform the following operations in the /usercode/blockChainScratch/blockchain/wallet.go file:

1. Create a struct Wallet with the following attributes:
    i. PrivateKey: The user’s private key
    ii. PublicKey: The user’s public key
2. Create a GenerateRSAKeys function that does the following:
    i. It generates a pair of RSA keys and the public and private keys.
    ii. It returns the generated private and public keys.
3. Create a NewWallet function that creates a new wallet instance and does the following:
    i. It calls the GenerateRSAKeys function to obtain the private and public keys.
    ii. It creates a new Wallet struct and assigns the generated private and public keys to the respective attributes.
    iii. It returns the created Wallet instance.

If you’re not sure how to proceed, check the Hint tab or the Solution tab.

## HINT: Task 13: Add Wallet Functionality
Explore the crypto/rsa package for information on RSA keys.
The PrivateKey and PublicKey attributes should be of type *rsa.PrivateKey and *rsa.PublicKey, respectively.
Use the rsa.GenerateKey function from the crypto/rsa package to generate a pair of keys.

## SOLUTION: Task 13: Add Wallet Functionality
Add the following code in the /usercode/blockChainScratch/blockchain/wallet.go file:

```go
import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
)

type Wallet struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func GenerateRSAKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		return nil, nil, err
	}

	publicKey := &privateKey.PublicKey

	return privateKey, publicKey, nil
}

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
```