# httptest2
--
    import "github.com/smartystreets/httptest2"

Package httptest provides utilities for HTTP testing.

## Usage

#### type ResponseRecorder

```go
type ResponseRecorder struct {
	Code      int           // the HTTP response code from WriteHeader
	HeaderMap http.Header   // the HTTP response headers
	Body      *bytes.Buffer // the bytes.Buffer to append written data to
	Flushed   bool
}
```

ResponseRecorder is an implementation of http.ResponseWriter that records its
mutations for later inspection in tests.

#### func  NewRecorder

```go
func NewRecorder() *ResponseRecorder
```
NewRecorder returns an initialized ResponseRecorder.

#### func (*ResponseRecorder) Flush

```go
func (this *ResponseRecorder) Flush()
```
Flush sets Flushed to true and calls WriteHeader(200) if that method has not yet
been called.

#### func (*ResponseRecorder) Header

```go
func (this *ResponseRecorder) Header() http.Header
```
Header returns the response headers. Feel free to modify the headers all you
like until calling WriteHeader or Write. At that point any modifications to the
headers will not stick.

#### func (*ResponseRecorder) Write

```go
func (this *ResponseRecorder) Write(buffer []byte) (int, error)
```
Write writes to Body. Will also call WriteHeader(200) if that method has not yet
been called.

#### func (*ResponseRecorder) WriteHeader

```go
func (this *ResponseRecorder) WriteHeader(status int)
```
WriteHeader sets rw.Code. Only the first call will have any effect.
