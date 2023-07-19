# Ring Buffer in Go

This project provides a circular buffer, or ring buffer, with a fixed capacity, implemented in Go.

## Usage

To use this package, first download it:

```shell
go get github.com/dmkim/ringbuffer
```

## Example
```go
package main

import (
	"github.com/dmkim/ringbuffer"
	"fmt"
)

func main() {
	buf := ringbuffer.New(10)
	n, err := buf.Write([]byte("Hello"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes written.\n", n)

	data := make([]byte, 5)
	n, err = buf.Read(data, len(data))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read %d bytes: %s\n", n, data)
}
```

## Benchmark
```shell
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkRingBuffer/smallnest_ring_buffer-16              151471              8096 ns/op            1088 B/op          2 allocs/op
BenchmarkRingBuffer/old_ring_buffer-16                    118152              9912 ns/op             112 B/op          1 allocs/op
BenchmarkRingBuffer/dmkm_ring_buffer-16                   146571              7750 ns/op             112 B/op          1 allocs/op
```