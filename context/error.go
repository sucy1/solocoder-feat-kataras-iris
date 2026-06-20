package context

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func NewError(code int, message string, details any) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
		Details: details,
	}
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func (ctx *Context) StopWithErrorResponse(err *ErrorResponse) {
	ctx.StopWithJSON(err.Code, err)
}
