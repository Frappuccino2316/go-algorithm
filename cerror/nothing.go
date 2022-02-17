package cerror

type NothingError struct {
	message string
}

func (e *NothingError) Error() string {
	return e.message
}

func RaiseError() error {
	return &NothingError{message: "存在しませんでした"}
}
