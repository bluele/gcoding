package gcoding

import (
	"io"
)

type AlphaEncoder struct {
	w *BitsWriter
}

func NewAlphaEncoder(w *BitsWriter) *AlphaEncoder {
	return &AlphaEncoder{
		w: w,
	}
}

func (enc *AlphaEncoder) Write(v uint) {
	var i uint
	p, q := (v-1)/8, (v-1)%8
	for i = 0; i < p; i++ {
		enc.w.Write(0, 8)
	}
	if q > 0 {
		enc.w.Write(0, q)
	}
	enc.w.Write(1, 1)
}

func (enc *AlphaEncoder) Flush() error {
	return enc.w.Flush()
}

type AlphaDecoder struct{}

func NewAlphaDecoder() *AlphaDecoder {
	return &AlphaDecoder{}
}

func (dec *AlphaDecoder) Decode(r io.Reader) ([]uint, error) {
	return dec.DecodeN(r, 0)
}

func (dec *AlphaDecoder) decode(r *BitsReader, n int) ([]uint, error) {
	var targets []uint
	var value uint = 1
	count := 0

	for {
		b, err := r.Read(1)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if b > 0 {
			targets = append(targets, value)
			value = 1
			count++
			if n > 0 && count >= n {
				return targets, nil
			}
		} else {
			value++
		}
		if err == io.EOF {
			return targets, nil
		}
	}
}

func (dec *AlphaDecoder) DecodeN(r io.Reader, n int) ([]uint, error) {
	return dec.decode(NewBitsReader(r), n)
}
