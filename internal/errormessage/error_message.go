package errormessage

type ErrorMessage struct {
	ErrorMessage    string `json:"message"`
	ErrorStatusCode int    `json:"status_code"`
}

func NewErrorMessage(message string, statusCode int) ErrorMessage {
	return ErrorMessage{message, statusCode}
}

func (e *ErrorMessage) Message() string {
	return e.ErrorMessage
}

func (e *ErrorMessage) StatusCode() int {
	return e.ErrorStatusCode
}
