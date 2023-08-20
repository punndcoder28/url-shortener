package helpers

import "hash/fnv"

func HashURL(url string) string {
	fnvHashedString := fnv.New32a().Sum([]byte(url))

	return string(fnvHashedString)
}
