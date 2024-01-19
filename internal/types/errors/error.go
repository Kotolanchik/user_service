package errors

type AppError struct {
	Description string `json:"description"`
}

func NewAppError(msg string) AppError {
	return AppError{
		Description: msg,
	}
}

func (e AppError) Error() string {
	return e.Description
}
