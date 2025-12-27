# Task 11: Add Transactions to the Block
The transactions are stored inside the block, but so far there is no attribute to store the transactions in a block. You have to update the Block structure and the CreateBlock function to include a new attribute for storing transactions and receiving a list of transactions as input. You’ll also update the Genesis function to create a coinbase transaction that will reward the miner who successfully mines the block.

In this task, perform the following operations in the /usercode/blockChainScratch/blockchain/block.go file:
1. Update the Block structure to include a new attribute, Transactions, which will store a list of transactions associated with the block.
2. Update the CreateBlock function to receive this list of transactions as an additional parameter and initialize the new block to include this list.
3. Update the Genesis function to do the following:
    a. Create a new coinbase transaction with the following properties:
        ii. Sender: "Coinbase"
        iii. Receiver: "Genesis"
        iv. Amount: Reward amount for mining a block
        v. Coinbase: true
    b. Include the coinbase transaction in the list of transactions passed to the CreateBlock function when generating the Genesis block.

If you’re not sure how to proceed, check the Hint tab or the Solution tab.

## HINT: Task 11: Add Transactions to the Block
The Transactions attribute should be of type []*Transaction.

In the CreateBlock function, assign the transactions slice to the Transactions attribute of the block structure.

In the Genesis function, create a new Transaction instance for the coinbase transaction and include it in a slice that is passed to the CreateBlock function.


## SOLUTION: Task 11: Add Transactions to the Block
Update the code in the /usercode/blockChainScratch/blockchain/block.go file with the following:

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
	Transactions []*Transaction
}

func CreateBlock(data string, prevHash string, transactions []*Transaction) *Block {
	rand.Seed(time.Now().UnixNano())

	initialNonce := rand.Intn(10000)
	
	block := &Block{"", data, prevHash, initialNonce, transactions}

	newPow := NewProofOfWork(block)

	nonce, hash := newPow.MineBlock()

	block.Hash = string(hash[:])
	block.Nonce = nonce

	return block
}

func Genesis() *Block {
	coinbaseTransaction := &Transaction{
		Sender:   "Coinbase",
		Receiver: "Genesis",
		Amount:   0.0,
		Coinbase: true,
	}

	return CreateBlock("Genesis", "", []*Transaction{coinbaseTransaction})
}
```