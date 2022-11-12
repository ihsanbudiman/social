package constant

import "net/http"

const (
	// HTTPResponseOK is the response code for OK
	HTTPResponseOK = http.StatusOK

	// HTTPResponseCreated is the response code for Created
	HTTPResponseCreated = http.StatusCreated

	// HTTPResponseBadRequest is the response code for Bad Request
	HTTPResponseBadRequest = http.StatusBadRequest

	// HTTPResponseUnauthorized is the response code for Unauthorized
	HTTPResponseUnauthorized = http.StatusUnauthorized

	// HTTPResponseForbidden is the response code for Forbidden
	HTTPResponseForbidden = http.StatusForbidden

	// HTTPResponseNotFound is the response code for Not Found
	HTTPResponseNotFound = http.StatusNotFound

	// HTTPResponseInternalServerError is the response code for Internal Server Error
	HTTPResponseInternalServerError = http.StatusInternalServerError
)
