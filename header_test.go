package gocached

import (
	"reflect"
	"testing"
)

func TestRequestHeader(t *testing.T) {
	h := &RequestHeader{
		BaseHeader{
			Magic:        REQUEST,
			Opcode:       GET,
			Opaque:       0,
			Cas:          0,
			KeyLength:    4,
			ExtrasLength: 0,
			BodyLength:   0,
		},
		5,
	}
	reqBytes := h.Bytes()
	expected := []byte{
		REQUEST,
		byte(GET),
		0x0, 0x4, // length of key
		0x0, 0x0, // extra length
		0x0, 0x5, // vbucket
		0x0, 0x0, 0x0, 0x4, // Length of value
		0x0, 0x0, 0x0, 0x0, // opaque
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // CAS
	}

	if !reflect.DeepEqual(reqBytes, expected) {
		t.Errorf("Expected: %#v, got %#v", expected, reqBytes)
	}
}

func TestResponseHeader(t *testing.T) {
	h := &ResponseHeader{
		BaseHeader{
			Magic:        RESPONSE,
			Opcode:       GET,
			Opaque:       0,
			Cas:          0,
			KeyLength:    4,
			ExtrasLength: 0,
			BodyLength:   5,
		},
		SUCCESS,
	}
	resBytes := h.Bytes()
	expected := []byte{
		RESPONSE,
		byte(GET),
		0x0, 0x4, // length of key
		0x0, 0x0, // extra length
		0x0, 0x0, // status
		0x0, 0x0, 0x0, 0x9, // Length of value
		0x0, 0x0, 0x0, 0x0, // opaque
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // CAS
	}

	if !reflect.DeepEqual(resBytes, expected) {
		t.Errorf("Expected: %#v, got %#v", expected, resBytes)
	}
}
