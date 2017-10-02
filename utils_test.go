package gcoding

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func makeRandUInts(size int) []uint {
	vs := []uint{}
	for i := 0; i < size; i++ {
		vs = append(vs, uint(rand.Intn(1000)+1))
	}
	return vs
}
