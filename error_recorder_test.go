package httptest2

import (
	"errors"
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

type ErrorProneRecorderFixture struct {
	*gunit.Fixture
	recorder *ErrorProneRecorder
}

func (this *ErrorProneRecorderFixture) Setup() {
	this.recorder = NewErrorProneRecorder(errors.New("GOPHERS!"), nil) // first write: err; second write: nil
}

func (this *ErrorProneRecorderFixture) TestWriteReturnsSpecifiedErrors() {
	no := "This first write results in an error"
	n, err := this.recorder.Write([]byte(no))
	this.So(n, should.Equal, 0)
	this.So(err.Error(), should.Equal, "GOPHERS!")
	this.So(this.recorder.Body.String(), should.BeBlank)

	yes := "This write will work just fine"
	n, err = this.recorder.Write([]byte(yes))
	this.So(n, should.Equal, len(yes))
	this.So(err, should.BeNil)
	this.So(this.recorder.Body.String(), should.Equal, yes)
}
