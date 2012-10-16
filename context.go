package gocached

import (
	"io"
)

// A request context hold a request and response for that request
type Context struct {
	reader   io.Reader
	writer   io.Writer
	request  *Request
	response *Response
}

// create a new request context
func NewContext(reader io.Reader, writer io.Writer) {
}

// read current requet
func (ctx *Context) readRequest() {
}

// write response out
func (ctx *Context) writeResponse() {
}
