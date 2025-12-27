# Task 12: Test the Transactions
Now that the transactions can be added to a block, add blocks to the blockchain with transactions to see if everything works as expected.

In this task, perform the following operations in the /usercode/blockChainScratch/main.go file:

1. Add blocks to the blockchain by passing a dummy recipient of the coinbase transaction for that block.
2. Pass a slice of dummy transactions when adding a block.
3. Print the transactions in each block along with its attributes.

If you’re not sure how to proceed, check the Hint tab or the Solution tab.

## HINT: Task 12: Test the Transactions
The updated call to the AddBlock function will look as follows:

```go
chain.AddBlock("Block", "Coinbase Recipient", []*blockchain.Transaction{
   {Sender: "Coinbase Recipient", Receiver: "PersonB", Amount: DummyAmount}
})   
```


## SOLUTION: Task 12: Test the Transactions
Update the code in the /usercode/blockChainScratch/main.go file with the following:

```go
package main

import (
	"blockChain/blockchain"
	"fmt"
	"strconv"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("Block 1", "Alice", []*blockchain.Transaction{
		{Sender: "Alice", Receiver: "Bob", Amount: 1.5},
		{Sender: "Alice", Receiver: "Charlie", Amount: 19.5},
	})

	chain.AddBlock("Block 2", "Bob", []*blockchain.Transaction{
		{Sender: "Bob", Receiver: "Charlie", Amount: 2.3},
	})

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash of block: %x\n", block.Hash)

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