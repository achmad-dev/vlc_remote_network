package utils

import (
	"math/rand"
	"time"
)

func RandomPass() int {
	source := rand.NewSource(time.Now().Unix())
	res := rand.New(source)
	return res.Int()
}
