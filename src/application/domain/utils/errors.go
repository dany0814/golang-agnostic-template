package utils

import "errors"

type Type string

// http errors
const (
	Unauthorized          Type = "UNAUTHORIZED"
	BadRequest            Type = "BAD_REQUEST"
	Conflict              Type = "CONFLICT"
	InternalServer        Type = "INTERNAL"
	NotFound              Type = "NOT_FOUND"
	RequestEntityTooLarge Type = "PAYLOAD_TOO_LARGE"
	ServiceUnavailable    Type = "SERVICE_UNAVAILABLE"
	UnsupportedMediaType  Type = "UNSUPPORTED_MEDIA_TYPE"
	Forbidden             Type = "FORBIDDEN"
	AppError              Type = "ERROR_APPLICATION"
)

var ErrEmailUser error = errors.New("invalid email")
var ErrInvalidPhone error = errors.New("invalid phone")

// errors database
var ErrMsgDatabaseConnect string = "failed to database connection"
