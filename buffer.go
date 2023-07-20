package ringbuffer

import (
	"errors"
	"fmt"
)

// Error variables to denote Buffer Full and Empty states.
var (
	ErrIsFull  = errors.New("buffer is full")
	ErrIsEmpty = errors.New("buffer is empty")
)

// Buffer struct contains data and metadata for the buffer.
type Buffer struct {
	head     int
	tail     int
	length   int
	capacity int
	variable bool
	buf      []byte
}

// New creates and initializes a new buffer with a given size.
// The 'variable' flag determines whether the size of the buffer is fixed or variable.
func New(size int, variable bool) *Buffer {
	return &Buffer{
		head:     0,
		tail:     0,
		length:   0,
		capacity: size,
		variable: variable,
		buf:      make([]byte, size),
	}
}

// NewFrom initializes a new buffer using an existing byte slice.
// The capacity of the buffer is set to the length of the byte slice,
// and the variable property is set to the provided value.
func NewFrom(b []byte, variable bool) *Buffer {
	return &Buffer{
		head:     0,
		tail:     0,
		length:   0,
		capacity: len(b),
		variable: variable,
		buf:      b,
	}
}

// String method returns a formatted string containing the state of the Buffer.
func (b *Buffer) String() string {
	return fmt.Sprintf("head=%d, tail=%d, length=%d, capacity=%d, buffer=[%v]", b.head, b.tail, b.length, b.capacity, b.buf)
}

// Write writes data into the buffer.
func (b *Buffer) Write(data []byte) (n int, err error) {
	dataLen := len(data)

	if b.variable {
		if b.length+dataLen > b.capacity {
			b.increaseCapacity(max(b.length+dataLen, b.capacity))
		}
	}

	dataLen = min(dataLen, b.capacity-b.length)
	if dataLen == 0 {
		err = ErrIsFull
		return
	}

	n = dataLen
	data = data[:dataLen]

	for {
		if dataLen = len(data); dataLen <= 0 {
			break
		}

		var freeSpace int
		if b.head <= b.tail {
			freeSpace = min(dataLen, b.capacity-b.tail)
		} else {
			freeSpace = min(dataLen, b.head-b.tail)
		}

		copy(b.buf[b.tail:], data[:freeSpace])
		b.tail = (b.tail + freeSpace) % b.capacity
		b.length += freeSpace
		data = data[freeSpace:]
	}

	return
}

// ReadAll reads all data from the buffer.
func (b *Buffer) ReadAll() []byte {
	data := make([]byte, b.length)
	_, _ = b.Read(data, b.length)
	return data
}

// PeekAll reads all data from the buffer without removing it.
func (b *Buffer) PeekAll() []byte {
	data := make([]byte, b.length)
	_, _ = b.Peek(data, b.length)
	return data
}

// Read reads data from the buffer.
func (b *Buffer) Read(data []byte, dataLen int) (n int, err error) {
	n, err = b.read(data, dataLen, true)
	return
}

// Peek reads data from the buffer without removing it.
func (b *Buffer) Peek(data []byte, dataLen int) (n int, err error) {
	n, err = b.read(data, dataLen, false)
	return
}

// read reads data from the buffer with option to remove after reading.
func (b *Buffer) read(data []byte, dataLen int, shouldRemove bool) (n int, err error) {
	if b.length == 0 {
		return 0, ErrIsEmpty
	}

	dataLen = min(dataLen, b.length)

	n = dataLen
	data = data[:dataLen]

	for {
		if dataLen = len(data); dataLen <= 0 {
			break
		}

		var availableData int
		if b.head < b.tail {
			availableData = min(dataLen, b.tail-b.head)
		} else {
			availableData = min(dataLen, b.capacity-b.head)
		}

		copy(data, b.buf[b.head:b.head+availableData])

		if shouldRemove {
			b.head = (b.head + availableData) % b.capacity
			b.length -= availableData
		}

		data = data[availableData:]
	}

	return
}

// CanWrite checks if buffer has space to write data of given length.
func (b *Buffer) CanWrite(dataLen int) bool {
	return dataLen <= b.length
}

// CanRead checks if buffer has enough data to read of given length.
func (b *Buffer) CanRead(dataLen int) bool {
	return dataLen <= b.length
}

// IsEmpty checks if the buffer is empty.
func (b *Buffer) IsEmpty() bool {
	return b.length == 0
}

// IsFull checks if the buffer is full.
func (b *Buffer) IsFull() bool {
	return b.length == b.capacity
}

// Len returns the current length of the buffer.
func (b *Buffer) Len() int {
	return b.length
}

// increaseCapacity increases the capacity of the buffer.
func (b *Buffer) increaseCapacity(incCap int) {
	newData := make([]byte, b.capacity+incCap)
	n, _ := b.Peek(newData, b.length)

	b.head = 0
	b.tail = n
	b.length = n
	b.capacity = len(newData)
	b.buf = newData
}

// min returns the minimum of two integers.
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}
