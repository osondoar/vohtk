package errors

type AppError struct {
	Error   error
	Message string
	Code    int
}
