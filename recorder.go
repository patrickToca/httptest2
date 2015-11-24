// Package httptest provides utilities for HTTP testing.
package httptest2

import (
	"bytes"
	"net/http"
)

// ResponseRecorder is an implementation of http.ResponseWriter that
// records its mutations for later inspection in tests.
type ResponseRecorder struct {
	Code      int           // the HTTP response code from WriteHeader
	HeaderMap http.Header   // the HTTP response headers
	Body      *bytes.Buffer // the bytes.Buffer to append written data to
	Flushed   bool

	wroteStatusCode bool
}

// NewRecorder returns an initialized ResponseRecorder.
func NewRecorder() *ResponseRecorder {
	return &ResponseRecorder{
		HeaderMap: make(http.Header),
		Body:      new(bytes.Buffer),
		Code:      http.StatusOK,
	}
}

// Header returns the response headers. Feel free to modify the headers
// all you like until calling WriteHeader or Write. At that point any
// modifications to the headers will not stick.
func (this *ResponseRecorder) Header() http.Header {
	if this.wroteStatusCode {
		return make(http.Header)
	}
	return this.HeaderMap
}

// WriteHeader sets rw.Code. Only the first call will
// have any effect.
func (this *ResponseRecorder) WriteHeader(status int) {
	if !this.wroteStatusCode {
		this.Code = status
	}
	this.wroteStatusCode = true
}

// Write writes to Body. Will also call WriteHeader(200) if that
// method has not yet been called.
func (this *ResponseRecorder) Write(buffer []byte) (int, error) {
	this.WriteHeader(http.StatusOK)
	return this.Body.Write(buffer)
}

// Flush sets Flushed to true and calls WriteHeader(200) if that
// method has not yet been called.
func (this *ResponseRecorder) Flush() {
	this.WriteHeader(http.StatusOK)
	this.Flushed = true
}
