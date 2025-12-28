# Blockchain Definitions

This document collects concise definitions of critical concepts used in this repository and the accompanying tasks.

- **Blockchain**: A distributed ledger composed of an ordered chain of blocks. Each block references the previous block's hash, creating an immutable sequence of records.

- **Block**: A unit in the blockchain that contains data (one or more transactions or arbitrary payload), a `PrevHash` that links to the previous block, a `Hash` computed from the block contents, and optional fields such as `Nonce` and `Transactions`.

- **Genesis Block**: The first block in a blockchain. It has no previous hash (often an empty string) and serves as the chain's anchor.

- **Hash**: A fixed-size digest produced by a cryptographic hash function (the example code uses MD5 for brevity). A block's hash is computed from the block's data and previous hash and provides data integrity: changing the block's contents changes its hash.

- **Nonce**: A number used once that miners vary during the Proof of Work process to change the block's hash output and find a value that meets the difficulty requirement.

- **Proof of Work (PoW)**: A consensus mechanism where miners repeatedly compute hashes (by varying the nonce) until they find a hash that meets a difficulty target (for example, a certain number of leading zero bits). PoW makes block creation computationally expensive and helps secure the network.

- **Difficulty / Target**: Parameters that determine how hard mining is. Difficulty is often expressed as the required number of leading zero bits in the hash; the Target is a numerical threshold that a valid hash must be lower than.

- **Mining**: The process of performing Proof of Work to find a valid nonce and hash for a block. Successful mining produces a block that can be appended to the chain.

- **Transaction**: A simple record representing value transfer between a `Sender` and `Receiver`, often including an `Amount`. Transactions may be packaged into blocks.

- **Coinbase Transaction**: A special transaction included in a mined block that awards the miner (or a specified receiver) with newly created coins. It typically has a `Sender` value such as "Coinbase" and `Coinbase: true`.

- **Wallet**: A holder of cryptographic keys (private/public) used to sign transactions. In the demo, wallets generate RSA keys and sign transaction payloads for verification.

- **Signing & Verification**: Signing uses a private key to create a signature over a transaction digest; verification uses the corresponding public key to confirm the signature is valid and the transaction was not tampered with.

- **Immutable**: Because each block's hash depends on its data and the previous hash, altering any block changes its hash and breaks the chain unless subsequent blocks are recomputed, which provides tamper-evidence.

- **Chain Validation**: The process of verifying that each block's stored hash matches a recomputed hash of its contents and that each block's `PrevHash` equals the previous block's `Hash`. For PoW chains, validation also checks that the block's hash meets the difficulty target.

These definitions are distilled from the guided tasks in this repository (task-00.md through task-05.md) and are intended to support learning and documentation for the codebase.
