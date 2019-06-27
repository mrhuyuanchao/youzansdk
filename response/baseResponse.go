package response

type BaseResponse struct {
	Error ErrorInfo `json:"error_response"`
}

// ErrorInfo ...
type ErrorInfo struct {
	Msg     string `json:"msg"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
