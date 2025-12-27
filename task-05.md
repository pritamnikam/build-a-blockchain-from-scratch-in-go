# Task 5: Test the Blockchain
Now that you’ve got the initial structure of the blockchain in place, set up a blockchain, add blocks to it, and examine its structure.

In this task, perform the following operations in the /usercode/blockChainScratch/main.go file:
    1. Use the InitBlockChain and AddBlock functions defined in the blockchain package to initialize the blockchain and add some blocks to it, respectively.
    2. Iterate over the blocks in the blockchain and print the previous hash, data, and each block’s hash.

Note: If you modify the data inside a block, its hash will change. Furthermore, this change in the hash will propagate to the subsequent blocks in the blockchain. To observe this behavior, modify the data "Block 2" in the /usercode/blockChainScratch/main.go file and compare its hash with the previous run when the data was not tampered with. To test this behavior, follow the steps below:
    1. Modify the data "Block 2" to any desired value.
    2. Run the code again with the modified data and compare the hash of "Block 2" with the previous run. You will notice that the hash of "Block 2" has changed; consequently, the hashes of subsequent blocks have also changed.

If you’re not sure how to proceed, check the Hint tab or the Solution tab.

## HINT: Task 5: Test the Blockchain
Iterate over the blocks in the blockchain using a for loop and the range keyword.

Use the Printf() function to format and print the block information, including the %x format specifier for printing byte arrays in hexadecimal format.

## SOLUTION: Task 5: Test the Blockchain
Update the /usercode/blockChainScratch/main.go file with the following code:

```go
package main

import (
	"fmt"
	"blockChain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("Block 1")
	chain.AddBlock("Block 2")
	chain.AddBlock("Block 3")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash of block: %x\n", block.Hash)
		fmt.Println()
	}
}
```

Note: Navigate to the /usercode/blockChainScratch directory in the terminal, and run the following command to test the blockchain:

cd /usercode/blockChainScratch
go run main.go


