package gcoding

import (
	"bytes"
	"reflect"
	"testing"
)

func TestGammaDecode(t *testing.T) {
	for i := 0; i < 1000; i++ {
		values := makeRandUInts(100)
		buf := new(bytes.Buffer)
		w := NewBitsWriter(buf)
		enc := NewGammaEncoder(w)
		for _, v := range values {
			enc.Write(uint(v))
		}
		enc.Flush()
		dec := NewGammaDecoder()
		vs, err := dec.Decode(buf)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(values, vs) {
			t.Errorf("%v != %v", values, vs)
		}
	}
}
