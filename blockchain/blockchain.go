// Package blockchain contains a minimal, educational blockchain
// implementation (blocks, chain, proof-of-work and wallets).
package blockchain

// Transaction represents a simple transfer of value between two parties.
// Fields are kept simple for educational purposes: Sender and Receiver are
// string identifiers, Amount is the transferred value and Coinbase marks
// miner/coinbase transactions.
type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
	Coinbase bool
}

// BlockChain is a very small in-memory blockchain container that stores
// an ordered slice of block pointers. This implementation is intentionally
// minimal for learning purposes.
type BlockChain struct {
	Blocks []*Block
}

// InitBlockChain initializes a new blockchain with the genesis block.
func InitBlockChain() *BlockChain {
	genesisBlock := Genesis()
	return &BlockChain{[]*Block{genesisBlock}}
}

// AddBlock appends a new block with the provided data to the chain.
// The new block's PrevHash is set to the hash of the current last block.
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
