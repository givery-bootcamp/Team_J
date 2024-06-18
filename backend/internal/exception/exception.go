package exception

import "net/http"

type Exception struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
}

func (e Exception) Error() string {
	return e.Message
}

func new(status int, message string) *Exception {
	return &Exception{Status: status, Message: message}
}

var (
	// 4xx: Client Error
	ErrDuplicateUser  = new(http.StatusBadRequest, "Duplicate User")
	ErrInvalidRequest = new(http.StatusBadRequest, "Invalid Request")
	ErrSigninFailed   = new(http.StatusBadRequest, "Signin Failed")
	ErrUnauthorized   = new(http.StatusUnauthorized, "Unauthorized")
	ErrNotFound       = new(http.StatusNotFound, "Not Found")
	// 5xx: Server Error
	ErrInternalServerError = new(500, "Internal Server Error")
)