package entity

var (
	// ErrInvalidCredential represents an invalid credential error.
	// It happens when email, password, access token, user agent is wrong.
	ErrInvalidCredential = NewAuthError("01-001", "Credential is invalid")
	// ErrInvalidLoginForm is returned when the request doesn't
	// follow the proper request body for login request.
	ErrInvalidLoginForm = NewAuthError("01-007", "Input format is invalid")
	// ErrInvalidInput is returned when the input doesn't comply
	// with the validation, e.g: the length of password or email format.
	ErrInvalidInput = NewAuthError("01-008", "Input is invalid")
	// ErrAccountNotFound is returned when the account can't be found.
	ErrAccountNotFound = NewAuthError("01-009", "Account not found")
	// ErrEmptyAccessToken is returned when the request doesn't contain
	// access token.
	ErrEmptyAccessToken = NewAuthError("01-010", "Access Token is empty")
	// ErrInternalServerError is returned whenever a problem occurs
	// in the system itself.
	ErrInternalServerError = NewAuthError("01-011", "Internal server error")
	// ErrInvalidEmailFormat occurs when email format is invalid
	ErrInvalidEmailFormat = NewAuthError("01-012", "Invalid email format")
	// ErrAccessDenied occurs when request using account that don't
	// have access
	ErrAccessDenied = NewAuthError("01-013", "Access Denied")
	// ErrEmptyID is returned when the request doesn't contain
	// ID
	ErrEmptyID = NewAuthError("01-017", "ID is empty")
	// ErrFailFindUser is returned when the system is fail to find user
	ErrFailFindUser = NewAuthError("01-18", "Fail on find user")
	// ErrInvalidUUIDFormat is returned when the request UUID format is wrong
	ErrInvalidUUIDFormat = NewAuthError("01-19", "Invalid UUID Format")
)

// AuthError represents error happens from authentication service.
type AuthError struct {
	PublicMessage string `json:"message"`
	Code          string `json:"code"`
	message       string
}

// NewAuthError creates an instance of AuthError.
func NewAuthError(code, message string) *AuthError {
	return &AuthError{
		Code:          code,
		PublicMessage: message,
	}
}

// NewAuthErrorWithOriginalMessage creates an instance of AuthError
// from another AuthError imbued with the original message.
//
// All predefined errors should exist as global variable (see top of this file).
// Therefore, this constructor aims to create a new AuthError from previously predefined error.
func NewAuthErrorWithOriginalMessage(err *AuthError, message string) *AuthError {
	return &AuthError{
		Code:          err.Code,
		PublicMessage: err.PublicMessage,
		message:       message,
	}
}

// Error returns original error message
// without masking it for human ro read.
// Please, note that the original message
// must be set in constructor.
// Otherwise, it will always be empty.
func (a *AuthError) Error() string {
	return a.message
}
