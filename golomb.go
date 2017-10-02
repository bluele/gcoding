package gcoding

import (
	"fmt"
	"io"
)

type GolombEncoder struct {
	w *BitsWriter
	b uint
	k uint
}

func NewGolombEncoder(w *BitsWriter, b uint) *GolombEncoder {
	if (b & (b - 1)) != 0 {
		panic(fmt.Sprintf("b is not 2^n"))
	}
	return &GolombEncoder{
		w: w,
		b: b,
		k: estimateBit(b),
	}
}

func (enc *GolombEncoder) Write(n uint) error {
	p, q := n/enc.b+1, n%enc.b
	aenc := NewAlphaEncoder(enc.w)
	aenc.Write(p)
	v := q & mask(enc.k)
	enc.w.Write(v, enc.k)

	return nil
}

func (enc *GolombEncoder) Flush() error {
	return enc.w.Flush()
}

type GolombDecoder struct {
	b uint
	k uint
}

func NewGolombDeocder(b uint) *GolombDecoder {
	if (b & (b - 1)) != 0 {
		panic(fmt.Sprintf("b is not 2^n"))
	}
	return &GolombDecoder{
		b: b,
		k: estimateBit(b),
	}
}

func (dec *GolombDecoder) Decode(r io.Reader) ([]uint, error) {
	var targets []uint
	br := NewBitsReader(r)
	ad := NewAlphaDecoder()
	for {
		vs, err := ad.decode(br, 1)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if len(vs) == 0 {
			return targets, nil
		}
		p := uint(vs[0]) - 1
		q, err := br.Read(dec.k)
		if err != nil && err != io.EOF {
			return nil, err
		}
		targets = append(targets, p*dec.b+q)
		if err == io.EOF {
			return targets, nil
		}
	}
}
