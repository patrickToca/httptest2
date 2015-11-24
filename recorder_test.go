package httptest

import (
	"net/http"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

//go:generate go install github.com/smartystreets/gunit/gunit
//go:generate gunit

type Fixture struct {
	*gunit.Fixture

	recorder *ResponseRecorder
	handler  http.Handler
	request  *http.Request
}

func (this *Fixture) Setup() {
	this.request, _ = http.NewRequest("GET", "http://foo.com/", nil)
	this.recorder = NewRecorder()
}

func (this *Fixture) serveHTTP(action func(http.ResponseWriter, *http.Request)) {
	http.HandlerFunc(action).ServeHTTP(this.recorder, this.request)
}

func (this *Fixture) TestHTTP200ByDefault() {
	this.serveHTTP(nothing)
	this.So(this.recorder.Code, should.Equal, 200)
}
func nothing(response http.ResponseWriter, _ *http.Request) {
	/* no-op */
}

func (this *Fixture) TestFirstStatusCodeOnly() {
	this.serveHTTP(status)
	this.So(this.recorder.Code, should.Equal, 201)
}
func status(response http.ResponseWriter, _ *http.Request) {
	response.WriteHeader(201)
	response.WriteHeader(202)
}

func (this *Fixture) TestExcludesHeadersWrittenAfterStatusCode() {
	this.serveHTTP(headers)
	this.So(this.recorder.HeaderMap["Key1"], should.Resemble, []string{"Value1"})
	this.So(this.recorder.HeaderMap, should.NotContainKey, "Key2")
	this.So(this.recorder.Code, should.Equal, 201)
}
func headers(response http.ResponseWriter, _ *http.Request) {
	response.Header().Set("Key1", "Value1")
	response.WriteHeader(201)
	response.Header().Set("Key2", "Value2")
}

func (this *Fixture) TestCallingWriteCallsWriteHeaderWithHTTP200() {
	this.serveHTTP(write)
	this.So(this.recorder.Body.String(), should.Equal, "hi first")
	this.So(this.recorder.Flushed, should.BeFalse)
}
func write(response http.ResponseWriter, _ *http.Request) {
	response.Write([]byte("hi first"))
	response.WriteHeader(201)
	response.WriteHeader(202)
}

func (this *Fixture) TestFlushSendsHTTP200AndMarksRecorderAsFlushed() {
	this.serveHTTP(flush)
	this.So(this.recorder.Code, should.Equal, 200)
	this.So(this.recorder.Flushed, should.BeTrue)
}
func flush(response http.ResponseWriter, _ *http.Request) {
	response.(http.Flusher).Flush()
	response.WriteHeader(201)
	response.WriteHeader(202)
}
