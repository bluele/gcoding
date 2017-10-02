package gcoding

import (
	"io"
)

type GammaEncoder struct {
	w     *BitsWriter
	alpha *AlphaEncoder
	pos   int
}

func NewGammaEncoder(w *BitsWriter) *GammaEncoder {
	return &GammaEncoder{
		w:     w,
		alpha: NewAlphaEncoder(w),
	}
}

func (enc *GammaEncoder) Write(v uint) {
	n := estimateBit(v)
	enc.alpha.Write(uint(n))
	enc.w.Write(v, n-1)
}

func (enc *GammaEncoder) Flush() error {
	return enc.w.Flush()
}

type GammaDecoder struct{}

func NewGammaDecoder() *GammaDecoder {
	return &GammaDecoder{}
}

func (dec *GammaDecoder) Decode(r io.Reader) ([]uint, error) {
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
		size := vs[0]
		if size <= 1 {
			targets = append(targets, 1)
		} else {
			n, err := br.Read(size - 1)
			if err != nil && err != io.EOF {
				return nil, err
			}
			targets = append(targets, (1<<uint(size-1))|n)
		}
		if err == io.EOF {
			return targets, nil
		}
	}
}
