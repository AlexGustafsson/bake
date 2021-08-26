package runtime

type RuntimeError struct {
	Message string
}

func CreateRuntimeError(message string) *RuntimeError {
	return &RuntimeError{
		Message: message,
	}
}

func (err *RuntimeError) Error() string {
	return err.Message
}
