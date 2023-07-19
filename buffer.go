package ringbuffer

import (
	"errors"
	"fmt"
)

// Predefined errors
var (
	ErrIsFull     = errors.New("buffer is full")
	ErrIsEmpty    = errors.New("buffer is empty")
	ErrNoMoreData = errors.New("no more data")
)

// Buffer represents a circular buffer with a fixed capacity.
type Buffer struct {
	head     int
	tail     int
	len      int
	capacity int
	buf      []byte
}

// New creates a new Buffer with the specified size.
func New(size int) *Buffer {
	return &Buffer{
		head:     0,
		tail:     0,
		len:      0,
		capacity: size,
		buf:      make([]byte, size),
	}
}

// From creates a new Buffer using an existing T slice.
func From(b []byte) *Buffer {
	return &Buffer{
		head:     0,
		tail:     len(b),
		len:      len(b),
		capacity: len(b),
		buf:      b,
	}
}

// String returns a string representation of the Buffer.
func (b *Buffer) String() string {
	return fmt.Sprintf("head=%d, tail=%d, len=%d, capacity=%d, buffer=[%v]", b.head, b.tail, b.len, b.capacity, b.buf)
}

// Write adds data to the buffer. It returns ErrIsFull if there is not enough space.
func (b *Buffer) Write(data []byte) (n int, err error) {
	dataLen := len(data)
	if b.len+dataLen > b.capacity {
		err = ErrIsFull
		return
	}

	for {
		if dataLen = len(data); dataLen <= 0 {
			break
		}

		var freeSpace int
		if b.head <= b.tail {
			freeSpace = min(dataLen, b.capacity-b.tail)
		} else {
			freeSpace = min(dataLen, b.head-b.tail-1)
		}

		copy(b.buf[b.tail:], data[:freeSpace])
		b.tail = (b.tail + freeSpace) % b.capacity
		b.len += freeSpace
		data = data[freeSpace:]
	}

	return
}

// Read retrieves data from the buffer and removes it. It returns ErrIsEmpty if the buffer is empty,
// and ErrNoMoreData if there is not enough data.
func (b *Buffer) Read(data []byte, dataLen int) (n int, err error) {
	n, err = b.read(data, dataLen, true)
	return
}

// Peek retrieves data from the buffer without removing it. It returns ErrIsEmpty if the buffer is empty,
// and ErrNoMoreData if there is not enough data.
func (b *Buffer) Peek(data []byte, dataLen int) (n int, err error) {
	n, err = b.read(data, dataLen, false)
	return
}

func (b *Buffer) read(data []byte, dataLen int, shouldRemove bool) (n int, err error) {
	if b.len == 0 {
		return 0, ErrIsEmpty
	}

	if dataLen > b.len {
		return 0, ErrNoMoreData
	}

	data = data[:dataLen]
	for {
		if dataLen = len(data); dataLen <= 0 {
			break
		}

		var availableData int
		if b.head <= b.tail {
			availableData = min(dataLen, b.tail-b.head)
		} else {
			availableData = min(dataLen, b.capacity-b.head)
		}

		copy(data, b.buf[b.head:b.head+availableData])

		if shouldRemove {
			b.head = (b.head + availableData) % b.capacity
			b.len -= availableData
		}

		data = data[availableData:]
	}

	return dataLen, nil
}

// IsEmpty returns whether the buffer is empty or not.
func (b *Buffer) IsEmpty() bool {
	return b.len == 0
}

// Len returns the length of the data in the buffer.
func (b *Buffer) Len() int {
	return b.len
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
