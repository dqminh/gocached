package gocached

import (
	"reflect"
	"testing"
)

func TestSize(t *testing.T) {
	res := NewResponse(
		&ResponseHeader{
			BaseHeader{
				Opcode: SET,
				Opaque: 0,
				Cas:    0,
			},
			SUCCESS,
		},
		[]byte{'t', 'e', 's', 't'},
		[]byte{'1', '2'},
		[]byte{'v', 'a', 'l', 'u', 'e'})
	size := res.Size()
	if size != (24 + 4 + 5 + 2) {
		t.Errorf("Invalid Size")
	}
}

func TestBytes(t *testing.T) {
	res := NewResponse(
		&ResponseHeader{
			BaseHeader{
				Opcode: SET,
				Opaque: 0,
				Cas:    0,
			},
			SUCCESS,
		},
		[]byte("test"),
		[]byte("12"),
		[]byte("value"))
	resBytes := res.Bytes()
	expected := []byte{
		RESPONSE, byte(SET),
		0x0, 0x4, // length of key
		0x2,      // extra length
		0x0,      // reserved
		0x0, 0x0, // status
		0x0, 0x0, 0x0, 0xb, // Length of value
		0x0, 0x0, 0x0, 0x0, // opaque
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // CAS
		'1', '2',
		't', 'e', 's', 't',
		'v', 'a', 'l', 'u', 'e'}
	if !reflect.DeepEqual(resBytes, expected) {
		t.Errorf("Expected: %#v, got %#v", expected, resBytes)
	}
}
