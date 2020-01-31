package errors

import "errors"

const (
	CodeInvalidBody               = 4001
	CodeInvalidResourceID         = 4002
	CodeInvalidDate               = 4003
	CodeIntersectedPeriod         = 4004
	CodeInvalidPeriodTime         = 4005
	CodeTokenError                = 4006
	CodeNotSupportedAuthorization = 4007
	CodeInvalidQueryTime          = 4008
	CodeInvalidQuery              = 4009
	CodeForbidden                 = 4031
	CodeNotFound                  = 4041
	CodeServerError               = 5001
)

var (
	ErrInvalidParameter = errors.New("invalid parameter")
	ErrTokenExpired     = errors.New("token is expired")
	ErrTokenNotValidYet = errors.New("token not active yet")
	ErrTokenMalformed   = errors.New("that's not even a token")
	ErrTokenInvalid     = errors.New("couldn't handle this token")
)
