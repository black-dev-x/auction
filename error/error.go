package errors

type CustomError struct {
	Message    string       `json:"message"`
	StatusCode int          `json:"status_code"`
	Err        string       `json:"error"`
	Causes     []ErrorCause `json:"causes"`
}

type ErrorCause struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func (re *CustomError) Error() string {
	return re.Message
}

func BadRequestError(message string) *CustomError {
	return &CustomError{
		Message:    message,
		StatusCode: 400,
		Err:        "Bad Request",
	}
}

func InternalServerError(message string) *CustomError {
	return &CustomError{
		Message:    message,
		StatusCode: 500,
		Err:        "Internal Server Error",
	}
}

func NotFoundError(message string) *CustomError {
	return &CustomError{
		Message:    message,
		StatusCode: 404,
		Err:        "Not Found",
	}
}
