# Task 8: Validate the Proof of Work Algorithm
The PoW algorithm involves finding a valid block hash that meets a specific target. After running the PoW algorithm and obtaining the nonce, you will use it to derive the hash that meets the target. Implement a function that validates this derived hash and essentially validates the work done by the PoW algorithm.

Note: This task assumes that the Block struct includes a Nonce field (set during mining in Task 7), which will be used here for validation.

In this task, perform the following operations in the /usercode/blockChainScratch/blockchain/proof.go file:
1. Create a Validate function that takes no parameters and returns a boolean value indicating the validity of the derived hash.
2. Call the ComputeData function with the block’s nonce value to obtain the data for hashing.
3. Compute the hash of the data using MD5.
4. Compare the computed hash with the target value by converting both to big integers:
    1. In a Proof of Work system, a hash is considered "less than the target" when its numerical representation is smaller than the target threshold.
    2. This comparison determines whether the required computational work has been performed.
    3. Convert the hash to a big integer and use the Cmp method to compare with the target value.
5. Return true if the derived hash (as a big integer) is less than the target (valid hash), and false otherwise.

If you’re not sure how to proceed, check the Hint tab or the Solution tab.

## HINT: Task 8: Validate the Proof of Work Algorithm
Use the md5.Sum() method to compute the hash of the data.

Convert the hash to a big.Int type using the SetBytes() method for comparison.

Compare the hash value with the Target using the Cmp method.


## SOLUTION: Task 8: Validate the Proof of Work Algorithm
Add the following code in the /usercode/blockChainScratch/blockchain/proof.go file:

```go
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	computedData := pow.ComputeData(pow.Block.Nonce)

	computedHash := md5.Sum(computedData)
	intHash.SetBytes(computedHash[:])

	if intHash.Cmp(pow.Target) == -1 {
		return true
	} else {
		return false
	}
}
```