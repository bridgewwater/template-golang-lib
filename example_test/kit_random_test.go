package example_test

import (
	"math/rand"
	"time"
)

const randomStrLetterCnt = 62

var randomStrLetters = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

var randSeed *rand.Rand

// randomStr
//
//	new random string by cnt
func randomStr(cnt uint) string {
	result := make([]byte, cnt)
	if randSeed == nil {
		randSeed = rand.New(rand.NewSource(time.Now().Unix()))
	}
	for i := range result {
		result[i] = randomStrLetters[randSeed.Intn(randomStrLetterCnt)]
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
