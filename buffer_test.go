package ringbuffer

import (
	"testing"
)

func TestBuffer(t *testing.T) {
	bufferSize := 10
	b := New(bufferSize)
	data := []byte("hello")

	// Test IsEmpty
	if !b.IsEmpty() {
		t.Errorf("Expected buffer to be empty, but it was not")
	}

	// Test Len
	if b.Len() != 0 {
		t.Errorf("Expected initial buffer length to be 0, but got %d", b.Len())
	}

	// Test Write
	_, err := b.Write(data)
	if err != nil {
		t.Errorf("Write failed: %v", err)
	}
	if b.len != len(data) {
		t.Errorf("Buffer length expected %d, got %d", len(data), b.len)
	}

	// Test IsEmpty
	if b.IsEmpty() {
		t.Errorf("Expected buffer not to be empty, but it was")
	}

	// Test Len
	if b.Len() != len(data) {
		t.Errorf("Expected buffer length to be %d, but got %d", len(data), b.Len())
	}

	// Test Read
	readData := make([]byte, len(data))
	_, err = b.Read(readData, len(data))
	if err != nil {
		t.Errorf("Read failed: %v", err)
	}
	for i, v := range readData {
		if v != data[i] {
			t.Errorf("Read data expected %v, got %v", data[i], v)
		}
	}
	if b.len != 0 {
		t.Errorf("Buffer length expected 0, got %d", b.len)
	}

	// Test Peek
	_, err = b.Write(data)
	if err != nil {
		t.Errorf("Write failed: %v", err)
	}
	peekData := make([]byte, len(data))
	_, err = b.Peek(peekData, len(data))
	if err != nil {
		t.Errorf("Peek failed: %v", err)
	}
	for i, v := range peekData {
		if v != data[i] {
			t.Errorf("Peek data expected %v, got %v", data[i], v)
		}
	}
	if b.len != len(data) {
		t.Errorf("Buffer length expected %d, got %d", len(data), b.len)
	}
}
