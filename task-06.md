# Task 6: Create and Instantiate the Proof of Work Structure
Blockchain networks employ the proof of work (PoW) consensus algorithm to validate transactions and add new blocks. The PoW algorithm requires miners to solve a complex mathematical puzzle, which requires a lot of computational power to solve, ensuring a secure blockchain network.

In the PoW algorithm, you have a target that represents the difficulty level for mining new blocks. It determines the criteria the block’s hash must meet, ensuring that mining requires significant computational work. When a new block is mined, the PoW algorithm repeatedly calculates the hash of the block until it meets the difficulty criteria specified by the target. Miners have to find a hash that, compared to the target, has the required number of leading zeros.

Note: The number of leading zeros required in the block’s hash determines the difficulty level. The more leading zeros required, the higher the difficulty.

You will start by creating a proof of work structure and a function that generates an instance of it for a given block after computing the target value.

In this task, perform the following operations in the /usercode/blockChainScratch/blockchain/proof.go file:
1. Define a variable called Difficulty that represents the desired difficulty level for mining new blocks.
2. Create a struct ProofOfWork with the following attributes:
    i. Block: The block to be validated
    ii. Target: The difficulty level for mining

Create a function NewProofOfWork that takes a block as its parameter and returns a new ProofOfWork instance. Inside the function, do the following:
1. Initialize a target variable with a value of 1.
2. Left shift the target’s value by (256 - Difficulty) bits. This sets the difficulty level.
3. Return a newly created ProofOfWork instance by assigning the block and target values.

Note: The target for mining a new block will remain constant in this example; however, in a real-world blockchain implementation, it is typically adjusted periodically to adapt to changes in the network’s computational power and maintain a consistent block creation rate.

If you’re not sure how to proceed, check the Hint tab or the Solution tab.

## HINT: Task 6: Create and Instantiate the Proof of Work Structure
The Target attribute should be of type big.Int for handling large integers.

Create a new big.Int variable with an initial value of 1 for the target.

Use the Lsh() method to left shift the value 256 - Difficulty bits.

## SOLUTION: Task 6: Create and Instantiate the Proof of Work Structure
In the /usercode/blockChainScratch/blockchain/proof.go file, do the following:

Add the following code after the package blockchain line to import the required dependencies:

```go
import (
   "bytes"
   "crypto/md5"
   "encoding/binary"
   "fmt"
   "math/big"
)
```

Add the following code to create the ProofOfWork struct and its instance:

```go
const Difficulty = 10

type ProofOfWork struct {
   Block  *Block
   Target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
   targetVal := big.NewInt(1)
   targetVal.Lsh(targetVal, uint(256-Difficulty))

   return &ProofOfWork{block, targetVal}
}
```