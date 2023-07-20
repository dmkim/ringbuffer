package ringbuffer

import (
	"testing"
)

func TestNewBuffer(t *testing.T) {
	b := New(10, false)
	if b.Len() != 0 {
		t.Errorf("Buffer length should be 0 but got %d", b.Len())
	}
	if b.capacity != 10 {
		t.Errorf("Buffer capacity should be 10 but got %d", b.capacity)
	}
}

func TestWriteToBuffer(t *testing.T) {
	b := New(10, false)
	data := []byte("test")
	n, err := b.Write(data)
	if err != nil {
		t.Errorf("Error writing to buffer: %v", err)
	}
	if n != len(data) {
		t.Errorf("Expected to write %d bytes, but wrote %d bytes", len(data), n)
	}
	if b.Len() != len(data) {
		t.Errorf("Buffer length should be %d but got %d", len(data), b.Len())
	}
}

func TestReadFromBuffer(t *testing.T) {
	b := New(10, false)
	data := []byte("test")
	_, _ = b.Write(data)

	readData := make([]byte, len(data))
	n, err := b.Read(readData, len(data))
	if err != nil {
		t.Errorf("Error reading from buffer: %v", err)
	}
	if n != len(data) {
		t.Errorf("Expected to read %d bytes, but read %d bytes", len(data), n)
	}
	if b.Len() != 0 {
		t.Errorf("Buffer length should be 0 but got %d", b.Len())
	}
}

func TestBufferOverflow(t *testing.T) {
	b := New(5, false)
	data := []byte("more than five bytes")
	n, err := b.Write(data)
	if err != nil {
		t.Errorf("Error writing to buffer: %v", err)
	}
	if n != 5 {
		t.Errorf("Expected to write %d bytes, but wrote %d bytes", 5, n)
	}
	if !b.IsFull() {
		t.Error("Buffer should be full")
	}
}
