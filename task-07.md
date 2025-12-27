# Task 7: Mine a Block
In this task, implement the mining process for a block in a blockchain using the PoW algorithm. Mining involves finding a valid hash for a given block that meets the difficulty criteria specified by the target. One important component in the mining process is the use of a nonce. You’ll start by initializing the data to compute the block’s hash and then perform the mining process. Simulate the mining process by repeatedly computing the hash of the block’s data with different nonces until a hash less than the target is found.

In this task, perform the following operations in the /usercode/blockChainScratch/blockchain/proof.go file:

Create a ComputeData function that takes an integer nonce as a parameter and returns the computed data for hashing the block. Inside the function, concatenate the following:
1. The hash of the previous block
2. The data contained within the current block
3. The nonce value
4. The difficulty level

Create a MineBlock function that performs the mining process to find a valid block hash that meets the target. Inside the function, do the following:
1. Initialize a variable nonce with a value of 0.
2. Initialize a loop that runs until a valid block hash is found and do the following:
    i. Call the ComputeData function with the current nonce value to obtain the data for hashing.
    ii. Compute the hash of the data.
    iii. Check if the computed hash value is less than the target defined for the PoW algorithm.
        • If it is, break the loop because a valid block hash has been found.
        • Otherwise, increment the nonce to continue searching for a valid hash.

    iv. Return the nonce and the valid computed hash. This returned nonce from the MineBlock function will be used later for validating the block's hash using the PoW algorithm in Task 8.

If you’re not sure how to proceed, check the Hint tab or the Solution tab.

# HINT: Task 7: Mine a Block
Prepare the data for hashing using the bytes.Join() function to concatenate the previous hash, block’s data, nonce, and Difficulty.
Use the md5.Sum() method to compute the hash of the data.
Convert the hash to a big.Int type using the SetBytes() method for comparison.

## SOLUTION: 

Add the following code in the /usercode/blockChainScratch/blockchain/proof.go file:

```go
func (pow *ProofOfWork) ComputeData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			[]byte(pow.Block.PrevHash),
			[]byte(pow.Block.Data),
			make([]byte, 8),
			make([]byte, 8),
		},
		[]byte{},
	)

	binary.BigEndian.PutUint64(data[len(data)-16:], uint64(nonce))
	binary.BigEndian.PutUint64(data[len(data)-8:], uint64(Difficulty))
	
	return data
}

func (pow *ProofOfWork) MineBlock() (int, []byte) {
	var intHash big.Int
	var computedHash [16]byte

	nonce := 0
	
	for {
		computedData := pow.ComputeData(nonce)
		computedHash = md5.Sum(computedData)

		fmt.Printf("\r%x", computedHash)

		intHash.SetBytes(computedHash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		}

		nonce++
	}
	fmt.Println()

	return nonce, computedHash[:]
}
```
