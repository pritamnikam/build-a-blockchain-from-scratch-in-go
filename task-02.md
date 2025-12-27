# Task 2: Compute the Hash of a Block
After creating the structures of Block and BlockChain, implement a function to compute the hash of a block. Hash functions are vital in a blockchain because they provide data integrity and security. Calculate the hash value based on the block’s data and previous hash using the MD5 hashing algorithm.

In this task, perform the following operations in the /usercode/blockChainScratch/blockchain/block.go file:
1. Create a function called ComputeHash that takes no arguments and has no return value.
2. Concatenate the block’s data and previous hash and calculate the hash of the concatenated data.
3. Assign the calculated hash value to the Hash attribute of the block.

If you’re not sure how to proceed, check the Hint tab or the Solution tab.

## HINT: Task 2: Compute the Hash of a Block
Use the bytes.Join() function to concatenate the block’s data and previous hash.

Use the md5.Sum() function to apply the MD5 hashing algorithm for calculating the hash of the concatenated data.

## SOLUTION: Task 2: Compute the Hash of a Block
In the /usercode/blockChainScratch/blockchain/block.go file, perform the following steps:

Add the following code after the package blockchain line to import the required dependencies:
```go
import (
   "bytes"
   "crypto/md5"
)
```

In the same directory, add the following code to compute the hash of the data:

```go
func (b *Block) ComputeHash() {
   concatenatedData := bytes.Join([][]byte{[]byte(b.Data), []byte(b.PrevHash)}, []byte{})
   
   computedHash := md5.Sum(concatenatedData)
   
   b.Hash = string(computedHash[:])
}
```