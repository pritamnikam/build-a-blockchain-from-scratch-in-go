# Blockchain Definitions & Concepts

This comprehensive glossary covers all critical blockchain concepts from the guided tasks (task-00.md → task-15.md) with diagrams, code examples and links to relevant source files.

---

## Quick Diagram: Block Chain Structure

```
[Genesis Block]
    |
    v PrevHash
[Block 1]
    |
    v PrevHash
[Block 2]
    |
    v PrevHash
[Block 3]
```

Each block stores its predecessor's hash. Tampering with any block breaks the chain.

---

## Core Concepts

### Blockchain
A distributed, append-only ledger of blocks linked by cryptographic hashes.

**Code location:** [blockchain/blockchain.go](blockchain/blockchain.go)

**Example:**
```go
chain := blockchain.InitBlockChain()       // Genesis block added
chain.AddBlock("Block 1", "Alice", txs)   // New block appended
```

---

### Block
A container holding transaction data, a previous hash reference, and its own computed hash.

**Fields:**
- `Hash` — cryptographic digest of block contents
- `Data` — payload/identifier
- `PrevHash` — hash of the previous block (chain link)
- `Nonce` — number used in Proof of Work mining
- `Transactions` — slice of transactions in the block

**Code location:** [blockchain/block.go](blockchain/block.go)

**Example:**
```go
type Block struct {
    Hash         string
    Data         string
    PrevHash     string
    Nonce        int
    Transactions []*Transaction
}
```

---

### Genesis Block
The first block in a blockchain; created when initializing the chain.

**Properties:**
- No previous hash (empty string `""`)
- Often contains a special "Genesis" payload
- Includes a coinbase transaction

**Code location:** [blockchain/block.go](blockchain/block.go) — `Genesis()` function

**Example:**
```go
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

---

### Hash
A fixed-size cryptographic digest (fingerprint) of block data.

**Properties:**
- Deterministic: same input → same hash
- Avalanche effect: tiny change → completely different hash
- Used in this project: MD5 (educational; use SHA-256 in production)

**Code location:** [blockchain/proof.go](blockchain/proof.go) — `ComputeData()` and `MineBlock()`

**Hash computation:**
```go
data := bytes.Join([][]byte{
    []byte(pow.Block.PrevHash),
    []byte(pow.Block.Data),
    ...
}, []byte{})
computedHash := md5.Sum(data)
```

---

### Nonce
A number varied during mining to change a block's hash.

**Purpose:** Find a nonce such that the block's hash meets the difficulty target.

**Code location:** [blockchain/block.go](blockchain/block.go#L12) — stored in `Block.Nonce`

**Used in:** [blockchain/proof.go](blockchain/proof.go) — mining algorithm

---

### Proof of Work (PoW)
A consensus mechanism requiring computational work to validate and add blocks.

**Process:**
1. Set a difficulty target
2. Iterate nonce values
3. Compute hash for each nonce
4. Stop when hash < target
5. Broadcast and validate

**Code location:** [blockchain/proof.go](blockchain/proof.go)

**Example:**
```go
pow := blockchain.NewProofOfWork(block)
nonce, hash := pow.MineBlock()    // Mining loop
block.Nonce = nonce
block.Hash = string(hash[:])
```

---

### Difficulty & Target
Parameters controlling mining hardness.

**Difficulty:** Number of leading zero bits required in a valid hash.

**Target:** Numerical threshold; hash must be less than target to be valid.

**Formula:** `Target = 1 << (256 - Difficulty)`

**Code location:** [blockchain/proof.go](blockchain/proof.go#L8)

```go
const Difficulty = 10  // Adjust for harder/easier mining

targetVal := big.NewInt(1)
targetVal.Lsh(targetVal, uint(256-Difficulty))
```

---

### Mining
The process of performing Proof of Work to find a valid nonce and hash.

**Code location:** [blockchain/proof.go](blockchain/proof.go) — `MineBlock()` function

**Step-by-step:**
```go
func (pow *ProofOfWork) MineBlock() (int, []byte) {
    nonce := 0
    for {
        data := pow.ComputeData(nonce)
        hash := md5.Sum(data)
        if hash < target {
            return nonce, hash[:]
        }
        nonce++
    }
}
```

---

### Transaction
A record of value transfer between participants.

**Fields:**
- `Sender` — address/public key of sender
- `Receiver` — address/public key of receiver
- `Amount` — quantity transferred
- `Coinbase` — whether this is a miner reward

**Code location:** [blockchain/blockchain.go](blockchain/blockchain.go#L3)

**Example:**
```go
type Transaction struct {
    Sender   string
    Receiver string
    Amount   float64
    Coinbase bool
}

tx := &Transaction{
    Sender:   "Alice",
    Receiver: "Bob",
    Amount:   5.0,
    Coinbase: false,
}
```

---

### Coinbase Transaction
A special first transaction in each block that rewards the miner.

**Properties:**
- `Sender` = "Coinbase"
- `Receiver` = miner's address
- `Amount` = mining reward (e.g., 10.0)
- `Coinbase` = true

**Code location:** [blockchain/blockchain.go](blockchain/blockchain.go#L19-L33) — `AddBlock()` function

**Example:**
```go
coinbaseTransaction := &Transaction{
    Sender:   "Coinbase",
    Receiver: coinbaseRcpt,
    Amount:   10.0,
    Coinbase: true,
}
```

---

### Wallet
A container holding a user's cryptographic keypair (private & public keys).

**Fields:**
- `PrivateKey` — secret key for signing transactions (RSA 2048-bit)
- `PublicKey` — public key for verifying signatures

**Code location:** [blockchain/wallet.go](blockchain/wallet.go)

**Example:**
```go
type Wallet struct {
    PrivateKey *rsa.PrivateKey
    PublicKey  *rsa.PublicKey
}

wallet, _ := blockchain.NewWallet()
alicePubKey := wallet.PublicKey.N.String()
```

---

### Digital Signature
A cryptographic proof that a transaction was authorized by the sender and not tampered with.

**Signing:** Private key + transaction data → signature
**Verification:** Public key + signature + transaction data → valid/invalid

**Code location:** [blockchain/wallet.go](blockchain/wallet.go)

**Example:**
```go
// Signing
signature, _ := wallet.SignTransaction(tx)

// Verification
err := blockchain.VerifyTransaction(tx, wallet.PublicKey, signature)
```

---

### Signing & Verification

#### Signing
Creates a digital signature for a transaction using the sender's private key.

**Code:** [blockchain/wallet.go](blockchain/wallet.go) — `SignTransaction()`

```go
func (wallet *Wallet) SignTransaction(tx *Transaction) (string, error) {
    dataString := fmt.Sprintf("%s%s%f%t",
        tx.Sender, tx.Receiver, tx.Amount, tx.Coinbase)
    hashedData := sha256.Sum256([]byte(dataString))
    signature, _ := rsa.SignPKCS1v15(rand.Reader, wallet.PrivateKey, crypto.SHA256, hashedData[:])
    return base64.StdEncoding.EncodeToString(signature), nil
}
```

#### Verification
Validates a signature using the sender's public key, ensuring authenticity.

**Code:** [blockchain/wallet.go](blockchain/wallet.go) — `VerifyTransaction()`

```go
func VerifyTransaction(tx *Transaction, pubKey *rsa.PublicKey, sig string) error {
    dataString := fmt.Sprintf("%s%s%f%t",
        tx.Sender, tx.Receiver, tx.Amount, tx.Coinbase)
    hashedData := sha256.Sum256([]byte(dataString))
    sigBytes, _ := base64.StdEncoding.DecodeString(sig)
    return rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashedData[:], sigBytes)
}
```

---

### Immutability
Property that makes tampering with historical blocks evident.

**How it works:**
1. Change Block 1 data → Block 1 hash changes
2. Block 2's `PrevHash` no longer matches Block 1's hash
3. The chain is "broken" (invalid)

**Consequence:** Any attempt to alter the past invalidates the entire chain.

---

### Chain Validation
Process of verifying blockchain integrity.

**Checks:**
1. Recompute each block's hash; verify it matches stored hash
2. Verify each block's `PrevHash` matches the previous block's hash
3. Verify each block satisfies PoW (hash < target)

**Code location:** [blockchain/proof.go](blockchain/proof.go) — `Validate()` function

**Example:**
```go
pow := blockchain.NewProofOfWork(block)
if pow.Validate() {
    fmt.Println("Block is valid PoW")
}
```

---

## End-to-End Example

```go
package main

import (
    "blockChain/blockchain"
    "fmt"
)

func main() {
    // 1. Initialize blockchain
    chain := blockchain.InitBlockChain()

    // 2. Create wallets
    aliceWallet, _ := blockchain.NewWallet()
    bobWallet, _ := blockchain.NewWallet()

    // 3. Create and sign transaction
    tx := &blockchain.Transaction{
        Sender:   aliceWallet.PublicKey.N.String(),
        Receiver: bobWallet.PublicKey.N.String(),
        Amount:   5.0,
    }
    signature, _ := aliceWallet.SignTransaction(tx)

    // 4. Verify transaction
    blockchain.VerifyTransaction(tx, aliceWallet.PublicKey, signature)

    // 5. Add block with transaction
    chain.AddBlock("Block 1", "Alice", []*blockchain.Transaction{tx})

    // 6. Display and validate
    for _, block := range chain.Blocks {
        fmt.Printf("Hash: %x\n", block.Hash)
        pow := blockchain.NewProofOfWork(block)
        fmt.Printf("Valid PoW: %v\n", pow.Validate())
    }
}
```

---

## Task-by-Task Reference

| Task | Title | Key Concepts |
|------|-------|--------------|
| [task-00.md](task-00.md) | Get Started | Blockchain, Block, Hash, PoW, Wallet overview |
| [task-01.md](task-01.md) | Create Blocks and Blockchain Structures | `Block` struct, `BlockChain` struct |
| [task-02.md](task-02.md) | Compute the Hash of a Block | Hash function, MD5 hashing |
| [task-03.md](task-03.md) | Create a New Block with the Genesis Block | `CreateBlock()`, `Genesis()` functions |
| [task-04.md](task-04.md) | Initialize and Add a Block to the Blockchain | `InitBlockChain()`, `AddBlock()` |
| [task-05.md](task-05.md) | Test the Blockchain | Basic demo, chain immutability |
| [task-06.md](task-06.md) | Create and Instantiate the Proof of Work Structure | `ProofOfWork` struct, `Difficulty`, `Target` |
| [task-07.md](task-07.md) | Mine a Block | `ComputeData()`, `MineBlock()` mining loop |
| [task-08.md](task-08.md) | Validate the Proof of Work Algorithm | `Validate()` PoW verification |
| [task-09.md](task-09.md) | Test the Proof of Work Algorithm | Integrate PoW into `CreateBlock()` |
| [task-10.md](task-10.md) | Create the Transaction Structure | `Transaction` struct, `Coinbase` |
| [task-11.md](task-11.md) | Add Transactions to the Block | `Block.Transactions` field |
| [task-12.md](task-12.md) | Test the Transactions | Display transactions in blocks |
| [task-13.md](task-13.md) | Add Wallet Functionality | `Wallet` struct, RSA key generation |
| [task-14.md](task-14.md) | Sign and Verify the Transactions | `SignTransaction()`, `VerifyTransaction()` |
| [task-15.md](task-15.md) | Test the Wallets | Full end-to-end demo with signing |

---

**Source Files:**
- [blockchain/block.go](blockchain/block.go)
- [blockchain/blockchain.go](blockchain/blockchain.go)
- [blockchain/proof.go](blockchain/proof.go)
- [blockchain/wallet.go](blockchain/wallet.go)
- [main.go](main.go)

---

*Last Updated: December 28, 2025*
*Covers: tasks 00 through 15*
