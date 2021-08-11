package http

import "github.com/pkg/errors"

var (
	ErrBodyNotAllowed  = errors.New("http: request method or response status code does not allow body")
	ErrHijacked        = errors.New("http: connection has been hijacked")
	ErrContentLength   = errors.New("http: wrote more than the declared Content-Length")
	ErrWriteAfterFlush = errors.New("unused")
)

type Handler interface {
	ServerHTTP(ResponseWriter, *Request)
}

type ResponseWriter interface {
	Header() Header
	Writer([]byte) (int, error)
	WriterHeader(statusCode int)
}

type Flusher interface {
	Flush()
}
