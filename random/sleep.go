package random

import (
	"math/rand"
	"time"
)

func Sleep() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(200-100) + 100
	time.Sleep(time.Duration(n) * time.Millisecond)
}
