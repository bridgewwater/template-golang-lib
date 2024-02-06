package example_test

import (
	"math/rand"
	"time"
)

var randSeed *rand.Rand

// randomStr
//
//	new random string by cnt
func randomStr(cnt uint) string {
	var letters = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	result := make([]byte, cnt)
	keyL := len(letters)
	if randSeed == nil {
		randSeed = rand.New(rand.NewSource(time.Now().Unix()))
	}
	for i := range result {
		result[i] = letters[randSeed.Intn(keyL)]
	}
	return string(result)
}

// randomInt
//
//	new random int by max
func randomInt(max int) int {
	if randSeed == nil {
		randSeed = rand.New(rand.NewSource(time.Now().Unix()))
	}
	return randSeed.Intn(max)
}
