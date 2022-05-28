package response

import "net/http"

type ErrorReponse struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e ErrorReponse) AsMessage() *ErrorReponse {
	return &ErrorReponse{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *ErrorReponse {
	return &ErrorReponse{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(message string) *ErrorReponse {
	return &ErrorReponse{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewValidationError(message string) *ErrorReponse {
	return &ErrorReponse{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}
