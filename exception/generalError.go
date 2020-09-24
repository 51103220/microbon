package exception

type generalError struct {
	content string
}

func (e *generalError) Error() string {
	return e.content
}

func NewgeneralError(text string) error {
	return &generalError{text}
}
