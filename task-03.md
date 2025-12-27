# Task 3: Create a New Block with the Genesis Block
The very first block in a blockchain is known as a Genesis block. It serves as the foundation of the entire blockchain network and is often created during the initialization of a new blockchain system. Since there are no preceding blocks in the chain, the Genesis block is unique because it does not reference previous blocks. You’ll also create a function to generate new blocks in the blockchain, each connected to the previous block.

In this task, perform the following operations in the /usercode/blockChainScratch/blockchain/block.go file:

1. Create a function called CreateBlock, which does the following:
    1. It takes the block data and the previous hash as arguments.
    2. It creates a new block by initializing a Block object with the provided data and previous hash.
    3. It computes the hash of the block and returns the created block.

2. Create a function called Genesis, which does the following:
    1. It takes no arguments and returns a pointer to a Block object.
    2. It creates a block, which contains hardcoded data and an empty previous hash.
    3. It calls the CreateBlock function with the appropriate arguments to create the Genesis block.

If you’re not sure how to proceed, check the Hint tab or the Solution tab.

## HINT: Task 3: Create a New Block with the Genesis Block
The Genesis block will contain hardcoded data (e.g., Genesis) and an empty previous hash (e.g., "").

Use the ComputeHash function to calculate the block’s hash.

## SOLUTION: Task 3: Create a New Block with the Genesis Block
Add the following code in the /usercode/blockChainScratch/blockchain/block.go file:

```go
func CreateBlock(data string, prevHash string) *Block {
	block := &Block{"", data, prevHash}
	block.ComputeHash()
	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", "")
}
```