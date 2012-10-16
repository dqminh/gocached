package gocached

import (
	"testing"
)

func TestRequestOpcode(t *testing.T) {
	header := []byte{128, 11, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	request, _ := NewRequest(header)
	if request.Header.Opcode != VERSION {
		t.Errorf("Invalid opcode")
	}
}

func TestInvalidRequestType(t *testing.T) {
	header := []byte{120, 11, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	_, err := NewRequest(header)
	if err == nil {
		t.Errorf("Expected error but did not receive")
	}
}
