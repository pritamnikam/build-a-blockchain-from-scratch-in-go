package blockchain

import (
	"bytes"
	"crypto/md5"
	"math/rand"
	"time"
)

// Block represents a single block in the blockchain.
// It contains a data payload, a reference to the previous block's hash
// and its own computed hash.
type Block struct {
	Hash         string
	Data         string
	PrevHash     string
	Nonce        int
	Transactions []*Transaction
}

// ComputeHash computes a simple hash for the block by concatenating
// the block's data and previous hash and computing an MD5 digest.
// The computed bytes are stored in the `Hash` field.
// Note: MD5 is used here for educational simplicity; real chains
// should use cryptographically secure hashes like SHA-256.
func (b *Block) ComputeHash() {
	concatenatedData := bytes.Join([][]byte{[]byte(b.Data), []byte(b.PrevHash)}, []byte{})
	computedHash := md5.Sum(concatenatedData)
	b.Hash = string(computedHash[:])
}

// CreateBlock constructs a new block with the given data and previous hash,
// computes its hash and returns a pointer to the block.

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

// Genesis creates the first block in the chain with the payload "Genesis".
func Genesis() *Block {
	coinbaseTransaction := &Transaction{
		Sender:   "Coinbase",
		Receiver: "Genesis",
		Amount:   0.0,
		Coinbase: true,
	}

	return CreateBlock("Genesis", "", []*Transaction{coinbaseTransaction})
}
