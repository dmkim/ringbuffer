# RingBuffer

RingBuffer is a Go package providing a simple, efficient, data-agnostic and thread-safe implementation of a circular buffer, also known as a ring buffer. This data structure is especially useful in scenarios where there is a need for managing data streams or maintaining a rolling window of the last N items.

## Features
- Adjustable Capacity: Supports dynamic resizing according to your needs.
- Error Handling: Clear indication of Buffer Full and Empty states via distinct error variables.
- Data Agnostic: Can handle any type of data represented in byte slices.

## Getting Started

### Installing

To start using RingBuffer, install Go and run `go get`:

```sh
$ go get -u github.com/dmkim/ringbuffer
```

This will retrieve the library.

### Usage

Import RingBuffer into your Go code:

```go
import "github.com/dmkim/ringbuffer"
```

Then construct a new RingBuffer, and use various methods to interact with it:

```go
// Initialize a new buffer with size 10
buffer := ringbuffer.New(10, false)

// Write data into the buffer
n, err := buffer.Write([]byte("test data"))

// Read data from the buffer
data := make([]byte, n)
n, err = buffer.Read(data, n)

// Check if buffer is empty
if buffer.IsEmpty() {
// Handle empty buffer case
}

// Check if buffer is full
if buffer.IsFull() {
// Handle full buffer case
}

// Other methods like Peek, CanRead, CanWrite, Len, etc are also available
```

## Contributing
Please feel free to submit issues, fork the repository and send pull requests!

## License
This project is licensed under the terms of the MIT license.