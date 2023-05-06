package errors

type DomenError struct {
	Module string
	Err    error
}

func (r *DomenError) Error() string {
	return r.Err.Error()
}
