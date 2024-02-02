package example_test

import (
	"math/rand"
	"time"
)

// randomStr
//
//	new random string by cnt
func randomStr(cnt uint) string {
	var letters = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	result := make([]byte, cnt)
	keyL := len(letters)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range result {
		result[i] = letters[rand.Intn(keyL)]
	}
	return string(result)
}

// randomInt
//
//	new random int by max
func randomInt(max int) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(max)
}
