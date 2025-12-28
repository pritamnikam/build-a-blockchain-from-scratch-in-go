# Task 15: Test the Wallets
Now that wallet functionality has been added, test it in the blockchain implementation. You will create wallets for two users, simulate a transaction between them, sign the transaction, verify its authenticity, and display the transaction details.

In this task, perform the following operations in the /usercode/blockChainScratch/main.go file:

1. Create wallets for two dummy users, Alice and Bob. Handle any errors that might occur during wallet creation and display a success message if the wallet is created successfully.
2. Create a new transaction from Alice to Bob with the following details:
    i. Sender: The public key of Alice’s wallet
    ii. Receiver: The public key of Bob’s wallet
    iii. Amount: A dummy transaction amount
3. Sign this newly created transaction using Alice’s wallet and store the generated signature.
4. Verify the transaction’s authenticity by passing the transaction, Alice’s wallet public key, and the signature obtained after signing the transaction.
5. Add the verified transaction to the blockchain by adding a block to the blockchain.
6. Print the transactions in each block along with its attributes.

If you’re not sure how to proceed, check the Hint tab or the Solution tab.


## HINT: Task 15: Test the Wallets
Create new wallets using the NewWallet function provided in the blockchain package.
Obtain the public key to Alice’s wallet using the aliceWallet.PublicKey.N.String method to set the transaction’s sender.
Obtain the public key of Bob’s wallet using the bobWallet.PublicKey.N.String method to set the transaction’s receiver.
Sign and verify the transaction using the SignTransaction and VerifyTransaction functions provided in the blockchain package.

## SOLUTION: Task 15: Test the Wallets
Update the code in the /usercode/blockChainScratch/main.go file with the following:

```go
package main

import (
	"blockChain/blockchain"
	"strconv"
	"fmt"
)

func main() {
	chain := blockchain.InitBlockChain()

	// Create a wallet for Alice.
	aliceWallet, err := blockchain.NewWallet()
	
	if err != nil {
		fmt.Println("Error creating Alice's wallet:", err)
		return
	}
	fmt.Println("Alice's wallet created successfully")

	// Create a wallet for Bob.
	bobWallet, err := blockchain.NewWallet()
	
	if err != nil {
		fmt.Println("Error creating Bob's wallet:", err)
		return
	}
	fmt.Println("Bob's wallet created successfully")

	// Create a transaction from Alice to Bob.
	tx := &blockchain.Transaction{
		Sender:   aliceWallet.PublicKey.N.String(),
		Receiver: bobWallet.PublicKey.N.String(),
		Amount:   5.0,
	}
	fmt.Println("Alice to Bob Transaction created successfully")

	// Sign the transaction with Alice’s wallet.
	signature, err := aliceWallet.SignTransaction(tx)
	if err != nil {
		fmt.Println("Error signing the transaction:", err)
		return
	}

	// Verify the transaction using Alice’s wallet, public key, and the signature.
	err = blockchain.VerifyTransaction(tx, aliceWallet.PublicKey, signature)
	
	if err != nil {
		fmt.Println("Transaction verification failed:", err)
		return
	}

	fmt.Println("Transaction Verified Successfully")

	// Add the verified transaction to the blockchain.
	chain.AddBlock("Block 1", "Alice", []*blockchain.Transaction{tx})
	fmt.Println()

	// Print the blockchain.
	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)


		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("IsValidPoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
		
		fmt.Println("Transactions:")

		for _, tx := range block.Transactions {
			fmt.Printf("Sender: %s\n", tx.Sender)
			fmt.Printf("Receiver: %s\n", tx.Receiver)
			fmt.Printf("Amount: %f\n", tx.Amount)
			fmt.Printf("Coinbase: %t\n", tx.Coinbase)
			fmt.Println()
		}
		fmt.Println()
	}
}
```