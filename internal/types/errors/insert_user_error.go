package errors

type InsertUserError struct {
	AppError
}

func NewInsertUserError(msg string) InsertUserError {
	return InsertUserError{NewAppError(msg)}
}
