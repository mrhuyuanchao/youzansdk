package response

type baseResponse struct {
	Code    string `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}
