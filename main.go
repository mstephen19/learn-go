package main

import (
	"bytes"
	"io"
)

type HashReader interface {
	io.Reader
	hash() string
}

type hashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

func (r *hashReader) hash() string {
	return ""
}

func main() {
	var reader HashReader = &hashReader{
		Reader: bytes.NewReader([]byte("foo")),
		buf:    bytes.NewBuffer([]byte("foo")),
	}

}
