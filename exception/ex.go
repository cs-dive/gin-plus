package exception

// BusinessException the service level exception, equivalent to resp.BadRequest
type BusinessException struct {
	Code int
	Msg  string
}

func (b *BusinessException) Error() string {
	return b.Msg
}

func NewBusinessErr(msg string) *BusinessException {
	return &BusinessException{40000, msg}
}

func NewBusinessErrWithCode(code int, msg string) *BusinessException {
	return &BusinessException{code, msg}
}
