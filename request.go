package gocached

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
)

type Request struct {
	Header            *RequestHeader
	Key, Extras, Body []byte
}

// given a bytes array, fill it with data from the reader
func fill(s io.Reader, buf []byte) error {
	x, err := io.ReadFull(s, buf)
	if err == nil && x != len(buf) {
		panic("Read full didn't")
	}
	return err
}

func fillRequest(data io.Reader, request *Request) (err error) {
	err = fill(data, request.Extras)
	if err != nil {
		return err
	}
	err = fill(data, request.Key)
	if err != nil {
		return err
	}
	return fill(data, request.Body)
}

// Given a request as a byte array, return a request object
func NewRequest(data []byte) (*Request, error) {
	req := &Request{}
	if uint8(data[0]) != REQUEST {
		return req, errors.New("Invalid Request")
	}

	opcode := uint8(data[1])
	keyLength := binary.BigEndian.Uint16(data[2:])
	extrasLength := data[4]
	vbucket := binary.BigEndian.Uint16(data[6:])
	bodyLength := binary.BigEndian.Uint32(data[8:]) - uint32(keyLength) - uint32(extrasLength)
	opaque := binary.BigEndian.Uint32(data[12:])
	cas := binary.BigEndian.Uint64(data[16:])

	if bodyLength > MAX_BODY_LENGTH {
		return req, errors.New(fmt.Sprintf("%v is too big", bodyLength))
	}

	header := &RequestHeader{
		BaseHeader{
			Magic:        REQUEST,
			Opcode:       opcode,
			Opaque:       opaque,
			Cas:          cas,
			KeyLength:    int(keyLength),
			ExtrasLength: int(extrasLength),
			BodyLength:   int(bodyLength),
		},
		vbucket,
	}

	req.Extras = make([]byte, extrasLength)
	req.Key = make([]byte, keyLength)
	req.Body = make([]byte, bodyLength)
	req.Header = header
	return req, nil
}

// create an instance of Request from incoming data
func ReadRequest(reader io.Reader) (request *Request, err error) {
	header := make([]byte, HEADER_LENGTH)
	bytesRead, err := io.ReadFull(reader, header)
	if err != nil {
		log.Printf("cannot read the request")
		return
	}
	if bytesRead != HEADER_LENGTH {
		log.Printf("request does not have sufficient length")
		return
	}

	request, err = NewRequest(header)
	if err != nil {
		return request, err
	}
	err = fillRequest(reader, request)
	if err != nil {
		return request, err
	}
	return request, nil
}
