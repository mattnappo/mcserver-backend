package common

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

// Hash represents an Sha3 hash.
type Hash struct {
	Hash []byte `json:"hash"`
}

// Sha3 hashes a byte input using Sha3.
func Sha3(b []byte) Hash {
	hash := sha3.New256()
	hash.Write(b)

	newHash := Hash{
		Hash: hash.Sum(nil),
	}
	return newHash
}

// String encodes a Hash to a string.
func (hash *Hash) String() string {
	return hex.EncodeToString(hash.Hash)
}
