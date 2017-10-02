# gcoding

gcoding is an encoding library for golang.
Currently implements unary, gamma and golomb-rice coding. 

## Example

```go
package main

import (
	"bytes"
	"fmt"

	"github.com/bluele/gcoding"
)

func main() {
	buf := new(bytes.Buffer)
	w := gcoding.NewBitsWriter(buf)
	enc := gcoding.NewGolombEncoder(w, 4)
	for _, v := range []uint{1, 2, 3, 4, 5, 10} {
		enc.Write(v)
	}
	enc.Flush()

	// 10011010 10110100 00100100 10100000
	fmt.Println(gcoding.BitsToString(buf.Bytes()))

	dec := gcoding.NewGolombDeocder(4)

	// [1 2 3 4 5 10] <nil>
	fmt.Println(dec.Decode(buf))
}
```
