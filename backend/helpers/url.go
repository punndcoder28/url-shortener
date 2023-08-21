package helpers

import (
	"encoding/hex"
	"hash/fnv"
)

func HashURL(url string) string {
	hash := fnv.New32a()
	hash.Write([]byte(url))
	hashValue := hash.Sum(nil)

	hashHex := hex.EncodeToString(hashValue)

	return hashHex
}

func CreateTempURLFromHash(hash string) string {
	return "http://localhost:8080/" + hash
}
