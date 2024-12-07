# Go Slice Modification During Iteration
This example demonstrates a common error in Go: modifying a slice while iterating over it using a traditional `for` loop.  Incorrectly removing elements can result in unexpected behavior and potential panics due to index out-of-bounds errors.  The solution provides a safer approach using a new slice.

## Bug
The `bug.go` file contains a function that attempts to remove elements with a value of 0 from a slice.  This function produces an incorrect result and will likely panic. 

## Solution
The `bugSolution.go` file showcases a corrected version that creates a new slice to avoid the issues encountered in the buggy implementation.
