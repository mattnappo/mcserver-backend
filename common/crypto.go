package common

import "golang.org/x/crypto/sha3"

// Sha3 hashes a byte input using Sha3.
func Sha3(b []byte) []byte {
	hash := sha3.New256()
	hash.Write(b)
	return hash.Sum(nil)
}
