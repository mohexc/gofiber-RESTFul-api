package utils

type AppError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e AppError) Error() string {
	return e.Message
}

func ErrNotFoundError(message string) error {
	return &AppError{
		Message: message,
		Status:  404,
	}
}
