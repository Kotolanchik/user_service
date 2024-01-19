package errors

type NotFoundAppError struct {
	AppError
}

func NewNotFoundAppError(msg string) NotFoundAppError {
	return NotFoundAppError{NewAppError(msg)}
}
