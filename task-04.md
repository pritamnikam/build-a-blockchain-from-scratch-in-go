# Task 4: Initialize and Add a Block to the Blockchain
In Task 3, you developed methods to create a Genesis block and add new blocks to the blockchain. In this task, start by creating the Genesis block to initialize a new blockchain. Then, add the functionality to include new blocks in the blockchain.

In this task, perform the following operations in the /usercode/blockChainScratch/blockchain/blockchain.go file:

1. Create a function called InitBlockChain, which does the following:
    i. It takes no arguments and returns a pointer to a BlockChain object.
    ii. It creates the Genesis block using the Genesis function defined in the blockchain package.
    iii. It creates a new BlockChain object and initializes the Blocks attribute with a slice containing only the Genesis block.
    iv. It returns a pointer to the newly created blockchain.

2. Create a function called AddBlock, which does the following:
    i. It takes the data to be added in a block as an argument.
    ii. It gets the previous block in the blockchain.
    iii. It creates a new block using the data and the previous hash.
    iv. It creates a new block using the CreateBlock function defined in the blockchain package. It passes the data and the previous block’s hash as arguments.
    v. It appends the newly created block to the Blocks attribute of the blockchain.

If you’re not sure how to proceed, check the Hint tab or the Solution tab.

## HINT: Task 4: Initialize and Add a Block to the Blockchain
Use blockchain.Blocks[len(blockchain.Blocks)-1] to obtain the previous block from the Blocks attribute of the blockchain.

Use blockchain.Blocks = append(blockchain.Blocks, newBlock) to append the newly created block to the Blocks attribute of the blockchain.

## SOLUTION: Task 4: Initialize and Add a Block to the Blockchain
Add the following code in the /usercode/blockChainScratch/blockchain/blockchain.go file:

```go
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}
```