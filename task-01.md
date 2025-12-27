# Task 1: Create Blocks and Blockchain Structures
In this task, you’ll create structures representing a block and blockchain.

To get started, perform the following operations:

In the /usercode/blockChainScratch/blockchain/block.go file, create a struct Block with the following attributes:
Hash: The hash of the current block
Data: The data stored in the block
PrevHash: The hash of the previous block in the chain

In the /usercode/blockChainScratch/blockchain/blockchain.go file, create a struct BlockChain. It should have a Blocks attribute that will be a slice of Block pointers representing the blocks in the blockchain.

If you’re not sure how to proceed, check the Hint tab or the Solution tab.

## HINT: Task 1: Create Blocks and Blockchain Structures
Create a struct using the following syntax:
```go
type StructName struct {
   // attribute 1
   // attribute 2
}   
```
The attributes should be of type string.

The Blocks attribute should be a slice of pointers to the Block objects.

## SOLUTION: Task 1: Create Blocks and Blockchain Structures
Add the following code in the /usercode/blockChainScratch/blockchain/block.go file:

```go
type Block struct {
   Hash         string
   Data         string
   PrevHash     string
}
```

Add the following code in the /usercode/blockChainScratch/blockchain/blockchain.go file:

```go
type BlockChain struct {
   Blocks []*Block
}
```