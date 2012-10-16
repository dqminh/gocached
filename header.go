package gocached

import (
	"encoding/binary"
)

type BaseHeader struct {
	Magic, Opcode, DataType             uint8
	Opaque                              uint32
	Cas                                 uint64
	KeyLength, ExtrasLength, BodyLength int
}

// Illustration of a request header
// Byte/     0       |       1       |       2       |       3       |
//    /              |               |               |               |
//   |0 1 2 3 4 5 6 7|0 1 2 3 4 5 6 7|0 1 2 3 4 5 6 7|0 1 2 3 4 5 6 7|
//   +---------------+---------------+---------------+---------------+
//  0| Magic         | Opcode        | Key length                    |
//   +---------------+---------------+---------------+---------------+
//  4| Extras length | Data type     | vbucket id                    |
//   +---------------+---------------+---------------+---------------+
//  8| Total body length                                             |
//   +---------------+---------------+---------------+---------------+
// 12| Opaque                                                        |
//   +---------------+---------------+---------------+---------------+
// 16| CAS                                                           |
//   |                                                               |
//   +---------------+---------------+---------------+---------------+
type RequestHeader struct {
	BaseHeader
	VbucketID uint16
}

// Illustration of a response header
// Byte/     0       |       1       |       2       |       3       |
//    /              |               |               |               |
//   |0 1 2 3 4 5 6 7|0 1 2 3 4 5 6 7|0 1 2 3 4 5 6 7|0 1 2 3 4 5 6 7|
//   +---------------+---------------+---------------+---------------+
//  0| Magic         | Opcode        | Key Length                    |
//   +---------------+---------------+---------------+---------------+
//  4| Extras length | Data type     | Status                        |
//   +---------------+---------------+---------------+---------------+
//  8| Total body length                                             |
//   +---------------+---------------+---------------+---------------+
// 12| Opaque                                                        |
//   +---------------+---------------+---------------+---------------+
// 16| CAS                                                           |
//   |                                                               |
//   +---------------+---------------+---------------+---------------+
//   Total 24 bytes
type ResponseHeader struct {
	BaseHeader
	Status uint16
}

// get total length of the header
func (h *BaseHeader) TotalLength() int {
	return h.KeyLength + h.ExtrasLength + h.BodyLength
}

// Export the header content into a bytes array. A header has 24 bytes
func (h *BaseHeader) partialBytes() []byte {
	data := make([]byte, HEADER_LENGTH)
	data[0] = h.Magic                                               // response code
	data[1] = byte(h.Opcode)                                        // opcode
	binary.BigEndian.PutUint16(data[2:4], uint16(h.KeyLength))      // Key Length
	data[4] = byte(h.ExtrasLength)                                  // Extras
	data[5] = 0                                                     // Data Type
	binary.BigEndian.PutUint32(data[8:12], uint32(h.TotalLength())) // Total Body Length
	binary.BigEndian.PutUint32(data[12:16], h.Opaque)               // Opague
	binary.BigEndian.PutUint64(data[16:24], h.Cas)                  // Cas
	return data
}

// return representation of request header as a byte array
func (h *RequestHeader) Bytes() []byte {
	data := h.partialBytes()
	binary.BigEndian.PutUint16(data[6:8], uint16(h.VbucketID))
	return data
}

// return representation fo response header as a byte array
func (h *ResponseHeader) Bytes() []byte {
	data := h.partialBytes()
	binary.BigEndian.PutUint16(data[6:8], uint16(h.Status))
	return data
}
