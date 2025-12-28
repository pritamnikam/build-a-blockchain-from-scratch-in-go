# Build a Blockchain from Scratch in Go

A minimal, educational implementation of a blockchain with Proof of Work consensus mechanism in Go.

## Features

- **Blockchain Structure**: Immutable chain of cryptographically linked blocks
- **Proof of Work (PoW)**: Mining mechanism requiring computational effort (difficulty-based hashing)
- **Transactions & Blocks**: Support for multiple transactions per block with coinbase rewards
- **Wallet System**: RSA-based wallet generation for key management
- **Digital Signatures**: Sign and verify transactions using cryptographic keys
- **Unit Tests**: Comprehensive test coverage for PoW, blocks and wallet functionality
- **MD5 Hashing**: Educational hashing (production should use SHA-256)

## Project Structure

```
.
├── blockchain/
│   ├── block.go              # Block structure with transactions and PoW mining
│   ├── block_test.go         # Unit tests for block creation and PoW validation
│   ├── blockchain.go         # Blockchain chain management with transactions
│   ├── proof.go              # Proof of Work algorithm implementation
│   ├── proof_test.go         # Unit tests for mining and validation
│   ├── wallet.go             # RSA wallet, signing and verification
│   └── wallet_test.go        # Unit tests for transaction signing
├── main.go                   # Full demo with wallets and transactions
├── definitions.md            # Glossary of blockchain concepts
├── go.mod                    # Go module definition (github.com/pritamnikam/...)
├── task-00.md → task-15.md   # Guided exercise tasks
└── README.md                 # This file
```

## Definitions & Glossary

See the comprehensive glossary of blockchain concepts with code examples and task references in [definitions.md](definitions.md). Covers all tasks from task-00.md through task-15.md.


## Prerequisites

- Go 1.20 or higher
- Basic understanding of Go and blockchain concepts

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/pritamnikam/build-a-blockchain-from-scratch-in-go.git
cd build-a-blockchain-from-scratch-in-go
```

### Run the Demo

The demo demonstrates:
1. Creating wallets for Alice and Bob
2. Signing a transaction with Alice's private key
3. Verifying the transaction with Alice's public key
4. Adding a block with the transaction to the blockchain
5. Validating PoW for each block

```powershell
go run main.go
```

**Expected Output:**
```
Alice's wallet created successfully
Bob's wallet created successfully
Alice to Bob Transaction created successfully
Transaction Verified Successfully
<mining output with hash attempts>

Previous hash: 
Data: Genesis
Hash: 64572909aa4583d58dfe87f6b89d667e
IsValidPoW: true

Transactions:
Sender: Coinbase
Receiver: Genesis
Amount: 0.000000
Coinbase: true

Previous hash: 64572909aa4583d58dfe87f6b89d667e
Data: Block 1
Hash: b54f09452958fcf772aeb41c2f7e6cda
IsValidPoW: true

Transactions:
Sender: Coinbase
Receiver: Alice
Amount: 10.000000
Coinbase: true

Sender: <Alice's public key>
Receiver: <Bob's public key>
Amount: 5.000000
Coinbase: false
```

### Run Tests

```powershell
go test ./...
```

This runs unit tests for:
- Block creation and PoW mining
- Transaction signing and verification
- PoW validation

## Key Concepts

See [definitions.md](definitions.md) for a comprehensive glossary with diagrams, code examples and task references.

### Block
A block contains:
- **Hash**: MD5 digest of the block's data, previous hash, and nonce
- **Data**: Block identifier/payload
- **PrevHash**: Hash of the previous block (chain link)
- **Nonce**: Number found during PoW mining
- **Transactions**: List of transactions in the block

### Proof of Work
Mining finds a nonce such that the block's hash meets a difficulty target (specified number of leading zero bits).

- **Difficulty = 10**: Hash must be numerically less than a target threshold
- Higher difficulty = more hashing iterations required
- Adjust `Difficulty` constant in [blockchain/proof.go](blockchain/proof.go) for testing

### Transactions & Wallets
- **Transaction**: Transfer record (Sender, Receiver, Amount, Coinbase flag)
- **Coinbase Transaction**: Miner reward (first transaction in each block)
- **Wallet**: RSA keypair for signing/verifying transactions
- **Digital Signature**: Proof of authorization (sender's private key signs, public key verifies)

## Possible Extensions

### Persistence
Store blocks in a database (SQLite, LevelDB, RocksDB) instead of in-memory storage.

### UTXO Model
Implement Unspent Transaction Output (UTXO) model for state management.

### Account Model
Alternatively, use an account-based model (like Ethereum) for balance tracking.

### P2P Networking
Add network layer for multi-node blockchain synchronization and consensus.

### Merkle Tree
Include Merkle root of transactions for efficient transaction verification.

### Gas & Fees
Implement transaction fees and computational cost (gas) models.

### Smart Contracts
Add simple script execution support for programmable transactions.

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