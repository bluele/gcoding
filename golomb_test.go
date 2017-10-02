package gcoding

import (
	"bytes"
	"reflect"
	"testing"
)

func TestGolombEncode(t *testing.T) {
	var bp uint = 128
	total := 0
	for i := 0; i < 100; i++ {
		buf := new(bytes.Buffer)
		w := NewBitsWriter(buf)
		enc := NewGolombEncoder(w, bp)
		vs := makeRandUInts(10)
		for _, v := range vs {
			enc.Write(uint(v))
		}
		enc.Flush()

		dec := NewGolombDeocder(bp)
		v, err := dec.Decode(buf)
		if err != nil {
			t.Fatal(err)
		}
		total += len(buf.Bytes())

		if !reflect.DeepEqual(vs, v) {
			t.Errorf("%v != %v", vs, v)
		}
	}
}
