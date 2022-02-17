package cerror

type NothingError struct {
	Message string
}

func (e *NothingError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &NothingError{Message: "存在しませんでした"}
}
