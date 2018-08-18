package sha256

import (
	"crypto/sha256"
	"hash"
	"log"
)

// Sha256init ...
func Sha256init(hash256 interface{}) interface{} {
	if hash256 != nil {
		if v, ok := hash256.(hash.Hash); ok {
			v.Reset()
		}
		return hash256
	}
	return sha256.New()
}

// Sha256update ...
func Sha256update(hash256 interface{}, inputData []byte) {
	if v, ok := hash256.(hash.Hash); ok {
		v.Write(inputData)
	}

}

// Sha256final ...
func Sha256final(hash256 interface{}) (hashData []byte) {
	if v, ok := hash256.(hash.Hash); ok {
		tmpHash := v.Sum(nil)
		log.Println("[sha256_final] sum byte is ", tmpHash)
		hashData = tmpHash
	}
	return
}
