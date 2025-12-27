# Build a Blockchain from Scratch in Go

A minimal, educational implementation of a blockchain with Proof of Work consensus mechanism in Go.

## Features

- **Blockchain Structure**: Immutable chain of cryptographically linked blocks
- **Proof of Work (PoW)**: Mining mechanism requiring computational effort (difficulty-based hashing)
- **SHA-256 Hashing**: Industry-standard cryptographic hash function for block security
- **Chain Validation**: Verify blockchain integrity and detect tampering

## Project Structure

```
.
├── blockchain/
│   ├── block.go          # Block structure and hashing logic
│   ├── blockchain.go     # Blockchain management and validation
│   ├── proof.go          # Proof of Work mining implementation
│   └── wallet.go         # (Future) Wallet and transaction support
├── main.go               # Demo application
├── go.mod                # Go module definition
└── README.md             # This file
```

## Prerequisites

- Go 1.16 or higher
- Basic understanding of blockchain concepts

## Getting Started

### Clone and Setup

```bash
cd c:\WORKSPACE\01_EDUCATIVE\03_PROJECTS\05-build-a-blockchain-from-scratch-in-go
go mod init github.com/your-username/blockchain
```

### Run the Demo

```bash
go run main.go
```

**Expected Output:**
```
Mining block: First Block - Transaction Data
Found hash: 0000abc...
Nonce: 1234

Mining block: Second Block - More Transactions
Found hash: 0000def...
Nonce: 5678

=== Blockchain ===

Block #0
Hash: 0000123...
Previous Hash: 
Data: Genesis Block
Nonce: 89

Block #1
Hash: 0000abc...
Previous Hash: 0000123...
Data: First Block - Transaction Data
Nonce: 1234
```

## Key Concepts

### Block
A block contains:
- **Hash**: SHA-256 hash of the block's data, previous hash, and nonce
- **Data**: Transaction payload
- **PrevHash**: Hash of the previous block (creates the chain link)
- **Nonce**: Number used in Proof of Work mining

### Proof of Work
Mining involves finding a nonce such that the block's hash has a required number of leading zeros (controlled by `Difficulty`).

- **Difficulty = 4**: Hash must start with 4 zeros (example: `0000abc123...`)
- Higher difficulty = more computational work required
- Increase `Difficulty` in `blockchain/block.go` for harder mining

### Chain Validation
The `IsValid()` method verifies:
1. Each block's hash is correctly computed
2. Each block's `PrevHash` matches the previous block's hash
3. No tampering has occurred

## Extending the Blockchain

### Add Transactions
Modify the `Block` struct to include a `Transactions` slice instead of a single `Data` field.

### Add a Wallet System
Implement cryptographic signing and verification in `blockchain/wallet.go` for transaction validation.

### Add Persistence
Use a database (SQLite, LevelDB) to persist blocks to disk.

### Implement P2P Networking
Add network functionality to sync blocks across multiple nodes.

## Learning Resources

- [SHA-256 Hash Function](https://en.wikipedia.org/wiki/SHA-2)
- [Proof of Work](https://en.wikipedia.org/wiki/Proof_of_work)
- [Bitcoin Whitepaper](https://bitcoin.org/bitcoin.pdf)
- [Go Cryptography Package](https://golang.org/pkg/crypto/)

## License

Educational project - MIT License

## Troubleshooting

**Q: Mining is slow**
A: Lower the `Difficulty` constant in `blockchain/block.go` (default is 4).

**Q: How do I verify blockchain integrity?**
A: Call `chain.IsValid()` to check if all blocks are properly linked and unmodified.

**Q: Can I use this in production?**
A: No. This is an educational implementation. Production blockchains require:
- Cryptographic signatures for transactions
- UTXO or account models for state management
- Consensus protocols (PoW, PoS, etc.)
- Network layer for distributed consensus