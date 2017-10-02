package gcoding

// This module is porting from https://github.com/tcnksm/go-casper/tree/master/internal/bits

import (
	"encoding/binary"
	"io"
)

// BitsWriter writes bits into underlying io.Writer
type BitsWriter struct {
	n uint // current number of bits
	v uint // current accumulated value

	wr io.Writer
}

// NewBitsWriter returns a new Writer.
func NewBitsWriter(w io.Writer) *BitsWriter {
	return &BitsWriter{
		wr: w,
	}
}

// BitsWrite writes bits with give size n.
func (w *BitsWriter) Write(bits uint, n uint) error {
	w.v <<= uint(n)
	w.v |= bits & mask(n)
	w.n += n
	for w.n >= 8 {
		b := (w.v >> (uint(w.n) - 8)) & mask(8)
		if err := binary.Write(w.wr, binary.BigEndian, uint8(b)); err != nil {
			return err
		}
		w.n -= 8
	}
	w.v &= mask(8)

	return nil
}

// Flush writes any remaining bits to the underlying io.Writer.
// bits will be left-shifted.
func (w *BitsWriter) Flush() error {
	if w.n != 0 {
		b := (w.v << (8 - uint(w.n))) & mask(8)
		if err := binary.Write(w.wr, binary.BigEndian, uint8(b)); err != nil {
			return err
		}
	}
	return nil
}

// BitsReader reads bits from the given io.Reader.
type BitsReader struct {
	n uint // current number of bits
	v uint // current accumulated value

	rd io.Reader
}

// NewBitsReader returns new a new Reader.
func NewBitsReader(rd io.Reader) *BitsReader {
	return &BitsReader{
		rd: rd,
	}
}

func (r *BitsReader) Read(n uint) (uint, error) {
	var err error

	for r.n <= n {
		r.v <<= 8
		var b uint8
		err = binary.Read(r.rd, binary.BigEndian, &b)
		if err != nil && err != io.EOF {
			return 0, err
		}
		r.v |= uint(b)

		r.n += 8
	}
	v := r.v >> uint(r.n-n)

	r.n -= n
	r.v &= mask(r.n)

	return v, err
}
