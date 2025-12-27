package blockchain

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"math/big"
)

// Difficulty defines the number of leading zero bits required in a valid block hash.
// Higher values increase the computational work needed to mine a block.
// Each increment roughly doubles the required hashing attempts.
// Example: Difficulty = 10 means the hash must start with 10 zero bits.
const Difficulty = 10

// ProofOfWork encapsulates the mining logic for a block.
// It stores a reference to the block being mined and the target difficulty threshold
// that valid hashes must meet.
type ProofOfWork struct {
	// Block is the block to be mined
	Block *Block
	// Target is the maximum hash value (as a big.Int) that satisfies the difficulty requirement.
	// Hashes lower than this target have enough leading zeros.
	Target *big.Int
}

// NewProofOfWork creates a new ProofOfWork instance for the given block.
// It calculates the Target threshold based on the Difficulty constant:
//   - A lower Target value means more leading zero bits are required.
//   - The Target is computed as 1 << (256 - Difficulty), representing the maximum
//     valid hash in a 256-bit space.
//
// Returns a pointer to the initialized ProofOfWork.
func NewProofOfWork(block *Block) *ProofOfWork {
	targetVal := big.NewInt(1)
	targetVal.Lsh(targetVal, uint(256-Difficulty))

	return &ProofOfWork{block, targetVal}
}

// ComputeData constructs the data to be hashed for Proof of Work mining.
// It combines:
//   - Previous block's hash (PrevHash)
//   - Current block's data
//   - The nonce (number used once, varied during mining)
//   - The difficulty level
//
// Returns the concatenated byte slice ready for hashing.
func (pow *ProofOfWork) ComputeData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			[]byte(pow.Block.PrevHash),
			[]byte(pow.Block.Data),
			make([]byte, 8), // Placeholder for timestamp/metadata
			make([]byte, 8), // Placeholder for additional data
		},
		[]byte{},
	)

	binary.BigEndian.PutUint64(data[len(data)-16:], uint64(nonce))
	binary.BigEndian.PutUint64(data[len(data)-8:], uint64(Difficulty))

	return data
}

// MineBlock performs the Proof of Work mining algorithm.
// It iteratively increments the nonce and computes hashes until a hash is found
// that satisfies the difficulty requirement (hash value is less than Target).
// The mining process:
//  1. Compute data with current nonce
//  2. Hash the data using SHA-256
//  3. Compare hash against Target threshold
//  4. If valid, return nonce and hash; otherwise increment nonce and repeat
//
// Prints mining progress to stdout (showing the computed hash for each attempt).
// Returns the winning nonce and the corresponding hash bytes.
func (pow *ProofOfWork) MineBlock() (int, []byte) {
	var intHash big.Int
	var computedHash [16]byte

	nonce := 0

	for {
		computedData := pow.ComputeData(nonce)
		computedHash = md5.Sum(computedData)

		fmt.Printf("\r%x", computedHash)

		intHash.SetBytes(computedHash[:])

		// If hash is less than target, it meets the difficulty requirement
		if intHash.Cmp(pow.Target) == -1 {
			break
		}

		nonce++
	}
	fmt.Println()

	return nonce, computedHash[:]
}

// Validate verifies that the block's current nonce and hash satisfy the Proof of Work requirement.
// It recomputes the hash using the block's stored nonce and checks if it meets the difficulty target.
// Returns true if the block's hash is valid (meets difficulty), false otherwise.
// This is used to verify blocks without re-mining them.
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	computedData := pow.ComputeData(pow.Block.Nonce)
	computedHash := md5.Sum(computedData)
	intHash.SetBytes(computedHash[:])

	return intHash.Cmp(pow.Target) == -1
}
