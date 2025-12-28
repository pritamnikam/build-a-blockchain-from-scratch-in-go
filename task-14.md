# Task 14: Sign and Verify the Transactions
To ensure security and integrity in a blockchain system, signing and verifying transactions is essential. A transaction is signed using a private key that provides a way to prove the authenticity of the transaction. The sender associates a transaction with their private key by appending a digital signature to a transaction, allowing other participants in the blockchain to verify that the sender authorized the transaction.

A transaction is verified using the sender’s public key, ensuring it has not been tampered with since it was signed. The public key is available to all participants and can be used by any participant to verify the signature. If the transaction data has been modified in any way, the verification process will fail, indicating that the transaction has been tampered with. You will implement functions to sign and verify transactions using a particular sender’s public and private keys.

In this task, perform the following operations in the /usercode/blockChainScratch/blockchain/wallet.go file:

1. Create a SignTransaction function that takes a transaction pointer as an argument and does the following:
    i. It generates a data string by concatenating the sender, receiver, amount, and transaction’s coinbase fields.
    ii. It hashes the generated data string.
    iii. It signs the hashed data using the current wallet’s private key and obtains the signature.
    iv. It encodes the signature as a base64 string and returns the encoded signature.

2. Create a VerifyTransaction function that takes a transaction, a public key, and a signature string as arguments and does the following:
    i. It generates a data string by concatenating the sender, receiver, amount, and transaction’s coinbase fields.
    ii. It hashes the generated data string.
    iii. It decodes the signature string to obtain the signature bytes.
    iv. It verifies the signature using the provided public key.
    v. It returns nil if the signature is valid and an error if the signature is invalid.

If you’re not sure how to proceed, check the Hint tab or the Solution tab.


## HINT: Task 14: Sign and Verify the Transactions
Use the rsa.SignPKCS1v15 function from the crypto/rsa package to sign the hashed data using the wallet’s private key.

Use base64.StdEncoding.EncodeToString to encode the signature as a base64 string.

Use the rsa.VerifyPKCS1v15 function from the crypto/rsa package to verify the signature with the provided public key.

Use base64.StdEncoding.DecodeString to decode the signature to obtain the signature bytes.

## SOLUTION: Task 14: Sign and Verify the Transactions
Add the following code in the /usercode/blockChainScratch/blockchain/wallet.go file:

```go
func (wallet *Wallet) SignTransaction(transaction *Transaction) (string, error) {
	dataString := fmt.Sprintf("%s%s%f%t", transaction.Sender, transaction.Receiver, transaction.Amount, transaction.Coinbase)

	hashedData := sha256.Sum256([]byte(dataString))

	signature, err := rsa.SignPKCS1v15(rand.Reader, wallet.PrivateKey, crypto.SHA256, hashedData[:])

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

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
```