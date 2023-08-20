package helpers

import "hash/fnv"

func HashURL(url string) string {
	fnvHashedString := fnv.New32a().Sum([]byte(url))

	return string(fnvHashedString)
}

func CreateTempURLFromHash(hash string) string {
	return "http://localhost:8080/" + hash
}
