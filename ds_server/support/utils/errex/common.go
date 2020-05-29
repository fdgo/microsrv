package errex

import (
	"net/http"
)

const BAD_REQUEST string = "BAD_REQUEST"
const UNAUTHORIZED string = "UNAUTHORIZED"
const FORBIDDEN string = "FORBIDDEN"
const NO_PERMISSION string = "NO_PERMISSION"
const INTERNAL_SERVER_ERROR string = "INTERNAL_SERVER_ERROR"
const INVALID_REQUEST_TYPE string = "INVALID_REQUEST_TYPE"
const SYSTEM_ERROR string = "SYSTEM_ERROR"
const INVALID_PARAMETER string = "INVALID_PARAMETER"
const NOT_FOUND string = "NOT_FOUND"
const ACCESS_TOO_OFTEN string = "ACCESS_TOO_OFTEN"
const CLIENT_TIME_INVALID string = "CLIENT_TIME_INVALID"

var (
	OK                     = NewError(http.StatusOK, "OK", "Success")
	BadRequest             = NewError(http.StatusOK, BAD_REQUEST, "Invalid request")
	Unauthorized           = NewError(http.StatusOK, UNAUTHORIZED, "unauthorized")
	Forbidden              = NewError(http.StatusOK, FORBIDDEN, "forbidden")
	NoPermission           = NewError(http.StatusOK, NO_PERMISSION, "no permission")
	InternalServerError    = NewError(http.StatusOK, INTERNAL_SERVER_ERROR, "server error")
	NotFound               = NewError(http.StatusOK, NOT_FOUND, "not found")
	ParameterError         = NewError(http.StatusOK, INVALID_PARAMETER, "Invalid parameter")
	InvalidReqError        = NewError(http.StatusOK, INVALID_REQUEST_TYPE, "Illegal request")
	SysError               = NewError(http.StatusOK, SYSTEM_ERROR, "System error")
	SysAccess2OftenError   = NewError(http.StatusOK, ACCESS_TOO_OFTEN, "Access too often")
	ClientTimeInvalidError = NewError(http.StatusOK, CLIENT_TIME_INVALID, "Please check your system time.")
)
