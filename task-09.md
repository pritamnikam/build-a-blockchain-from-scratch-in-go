# Task 9: Test the Proof of Work Algorithm
Once you have computed the nonce, you must store it inside that block. Currently, the Block structure has no attribute to store the nonce. Remember that you are now computing the block’s hash in the PoW algorithm; therefore, there is no need to do so separately for the block. Update the Block struct and the CreateBlock function to incorporate the PoW algorithm.

In this task, perform the following operations:

In the /usercode/blockChainScratch/blockchain/block.go file:
1. Add a Nonce attribute to the Block structure to store the nonce value computed via the PoW algorithm.
2. Remove the ComputeHash function since the hash is computed in the PoW algorithm.
3. Update the CreateBlock function to do the following:
    1. The block should have an initial value for the Nonce attribute.
    2. Create a new ProofOfWork instance for the block.
    3. Obtain the nonce and derived hash values by calling the MineBlock function on the ProofOfWork instance.
    4. Set the derived hash value and nonce values in the respective fields of the block.

In the /usercode/blockChainScratch/main.go file:
1. Create a new instance of the ProofOfWork structure when printing each block.
2. Run the validation to check if the PoW algorithm ran successfully or not.

If you’re not sure how to proceed, check the Hint tab or the Solution tab.


## HINT: Task 9: Test the Proof of Work Algorithm
The Nonce attribute should be of type int.

The initial value for the Nonce attribute in the block should be a random value.

## SOLUTION: Task 9: Test the Proof of Work Algorithm
Update the code in the /usercode/blockChainScratch/blockchain/block.go file

```go
package blockchain
import (
   "math/rand"
   "time"
)

type Block struct {
   Hash     string
   Data     string
   PrevHash string
   Nonce    int
}

func CreateBlock(data string, prevHash string) *Block {
   rand.Seed(time.Now().UnixNano()) // Seed the random number generator
   initialNonce := rand.Intn(10000)

   block := &Block{"", data, prevHash, initialNonce}

   newPow := NewProofOfWork(block)

   nonce, hash := newPow.MineBlock()

   block.Hash = string(hash[:])
   block.Nonce = nonce

   return block
}

func Genesis() *Block {
   return CreateBlock("Genesis", "")
}
```


Update the code in the /usercode/blockChainScratch/main.go file with the following:

```go
package main

import (
   "fmt"
   "strconv"
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

      pow := blockchain.NewProofOfWork(block)
      fmt.Printf("IsValidPoW: %s\n", strconv.FormatBool(pow.Validate()))

      fmt.Println()
   }
}
```


Note: Navigate to the /usercode/blockChainScratch directory in the terminal and run the following command to test the PoW algorithm:

```sh
cd /usercode/blockChainScratch
go run main.go
```