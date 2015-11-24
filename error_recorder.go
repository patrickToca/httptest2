package httptest2

// ErrorProneRecorder is an implementation of http.ResponseWriter which allows
// Write errors to be returned in an easily configured way.
type ErrorProneRecorder struct {
	*ResponseRecorder

	Errors   []error
	errIndex int
}

// NewErrorProneRecorder initializes a *ErrorProneRecorder. The errors provided
// are returned in the provided order from the Write method. A nil error allows a call to
// ResponseRecorder.Write.
func NewErrorProneRecorder(errors ...error) *ErrorProneRecorder {
	return &ErrorProneRecorder{
		ResponseRecorder: NewRecorder(),
		Errors:           errors,
	}
}

func (this *ErrorProneRecorder) Write(buffer []byte) (n int, err error) {
	defer this.increment()

	err = this.Errors[this.errIndex]
	if err != nil {
		n = 0
	} else {
		n, _ = this.ResponseRecorder.Write(buffer)
	}
	return n, err
}

func (this *ErrorProneRecorder) increment() {
	this.errIndex++
	if this.errIndex >= len(this.Errors) {
		this.errIndex = 0
	}
}
