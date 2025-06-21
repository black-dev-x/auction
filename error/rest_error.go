package error

type RestError struct {
	Message    string       `json:"message"`
	StatusCode int          `json:"status_code"`
	Err        string       `json:"error"`
	Causes     []ErrorCause `json:"causes"`
}

type ErrorCause struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func (re *RestError) Error() string {
	return re.Message
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: 400,
		Err:        "Bad Request",
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: 500,
		Err:        "Internal Server Error",
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: 404,
		Err:        "Not Found",
	}
}
