package gcoding

import (
	"fmt"
	"strings"
)

func mask(n uint) uint {
	return (1 << uint(n)) - 1
}

func estimateBit(v uint) uint {
	var i, size uint
	size = 32
	for i = 0; i < size; i++ {
		b := v & (1 << (size - 1 - i))
		if b > 0 {
			return size - i
		}
	}
	return 0
}

// BitsToString returns pretty format for bits
func BitsToString(bs []byte) string {
	var ss []string
	for _, b := range bs {
		ss = append(ss, fmt.Sprintf("%08b", b))
	}
	return strings.Join(ss, " ")
}
