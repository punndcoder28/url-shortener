package helpers

import (
	"encoding/hex"
	"hash/fnv"
)

/*
HashURL - This function returns a hashed string of the URL
*/
func HashURL(url string) string {
	hash := fnv.New32a()
	hash.Write([]byte(url))
	hashValue := hash.Sum(nil)

	hashHex := hex.EncodeToString(hashValue)

	return hashHex
}

/*
CreateTempURLFromHash - This function returns a temporary URL from the hash
*/
func CreateTempURLFromHash(hash string) string {
	return "http://localhost:8080/" + hash
}
