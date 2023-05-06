package errors

type HttpError struct {
	StatusCode int
	Err        error
}

func (r *HttpError) Error() string {
	return r.Err.Error()
}
