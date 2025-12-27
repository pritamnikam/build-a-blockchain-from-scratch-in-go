# Task 10: Create the Transaction Structure
In a blockchain, a transaction represents data transfer from one participant to another within the network. Transactions generally include data such as the sender’s and receiver’s addresses and the amount or quantity being transferred. When a transaction is initiated, it undergoes validation to ensure its authenticity. Once validated, the transaction is included in a block and added to the blockchain through mining.

A coinbase transaction is a special type of transaction that is primarily used to reward miners who successfully mine a new block. It is the first transaction in each block and does not have a specific sender like regular transactions. Instead, it is created by the miner who successfully solves the cryptographic puzzle associated with mining a block. It rewards the miner by specifying the recipient’s address and an amount as a reward for their mining efforts.

You will create a new structure for transactions and update the functionality for adding a block such that a new coinbase transaction is created when adding a block to the blockchain.

In this task, perform the following operations in the /usercode/blockChainScratch/blockchain/blockchain.go file:

1. Create a struct Transaction with the following attributes:
    1. Sender: The sender’s address
    2. Receiver: The receiver’s address
    3. Amount: The amount being transferred
    4. Coinbase: Indicator of whether the transaction is a coinbase transaction or not

2. Update the AddBlock function, which does following:
    1. It takes two additional parameters:
        1. coinbaseRcpt: The recipient of the coinbase transaction
        2. transactions: A slice of transactions to be included in the block
    2. It creates a new coinbase transaction for the block with the following properties:
        1. Sender: "Coinbase"
        2. Receiver: coinbaseRcpt
        3. Amount: Reward amount for mining a block
        4. Coinbase: true
    3. It appends the coinbase transaction as the first element in the transactions slice and invokes the CreateBlock function by passing the slice as a parameter.

Note: Update the CreateBlock function in the /usercode/blockChainScratch/blockchain/block.go file to receive a slice of transactions as a parameter.

If you’re not sure how to proceed, check the Hint tab or the Solution tab.

## HINT: Task 10: Create the Transaction Structure
The sender and receiver addresses should be of type string.

The amount being transferred should be of type float64.

The indicator of whether the transaction is a coinbase transaction or not should be of type bool.


## SOLUTION: Task 10: Create the Transaction Structure
In the /usercode/blockChainScratch/blockchain/blockchain.go file, do the following:

Add the following code to create the Transaction struct:

```go
type Transaction struct {
   Sender   string
   Receiver string
   Amount   float64
   Coinbase bool
}
```

Update the AddBlock function with the following code:
```go
func (chain *BlockChain) AddBlock(data string, coinbaseRcpt string, transactions []*Transaction) {
   prevBlock := chain.Blocks[len(chain.Blocks)-1]

   coinbaseTransaction := &Transaction{
      Sender:   "Coinbase",
      Receiver: coinbaseRcpt,
      Amount:   10.0,
      Coinbase: true,
   }

   newBlock := CreateBlock(data, prevBlock.Hash, append([]*Transaction{coinbaseTransaction}, transactions...))

   chain.Blocks = append(chain.Blocks, newBlock)
}
```