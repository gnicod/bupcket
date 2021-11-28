package error

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var MISSING_FILE = ErrorResponse{
	Code:    100,
	Message: "Missing file in request",
}
var CONFIG_NOT_FOUND = ErrorResponse{
	Code:    100,
	Message: "Config not found",
}
var INVALID_MIME_TYPE = ErrorResponse{
	Code:    100,
	Message: "Invalid MimeTypes",
}
