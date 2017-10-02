package gcoding

import (
	"bytes"
	"reflect"
	"testing"
)

func TestAlphaEncoder(t *testing.T) {
	for i := 0; i < 100; i++ {
		buf := new(bytes.Buffer)
		w := NewBitsWriter(buf)
		enc := NewAlphaEncoder(w)
		dec := NewAlphaDecoder()

		vs := makeRandUInts(100)
		for _, v := range vs {
			enc.Write(uint(v))
		}
		enc.Flush()
		v, err := dec.Decode(buf)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(vs, v) {
			t.Errorf("%v != %v", vs, v)
		}
	}
}
