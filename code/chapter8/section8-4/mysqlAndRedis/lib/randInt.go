package lib

import (
	"math/rand"
	"time"
)

func GenRandInt(ceiling int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(ceiling)
}
