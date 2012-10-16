package gocached

type Response struct {
	Header            *ResponseHeader
	Key, Extras, Body []byte
}

// Create new Response
func NewResponse(Header *ResponseHeader, Key, Extras, Body []byte) *Response {
	// init some additional information for the header
	Header.Magic = RESPONSE
	Header.KeyLength = len(Key)
	Header.ExtrasLength = len(Extras)
	Header.BodyLength = len(Body)
	// a response contains header, key, extras and body
	return &Response{
		Header: Header,
		Key:    Key,
		Extras: Extras,
		Body:   Body}
}

// Return total size of the response
func (res *Response) Size() int {
	return HEADER_LENGTH + res.Header.TotalLength()
}

// Returns representation of the response as a byte array
func (res *Response) Bytes() []byte {
	data := make([]byte, res.Size())
	start := 24
	next := start

	// copy then header
	copy(data[0:start], res.Header.Bytes())

	// copy extras
	if res.Header.ExtrasLength > 0 {
		next = start + res.Header.ExtrasLength
		copy(data[start:next], res.Extras)
		start = next
	}

	// copy key
	if res.Header.KeyLength > 0 {
		next = start + res.Header.KeyLength
		copy(data[start:next], res.Key)
		start = next
	}

	// copy body
	next = start + res.Header.BodyLength
	copy(data[start:next], res.Body)

	return data
}
