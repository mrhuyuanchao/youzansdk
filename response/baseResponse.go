package response

type baseResponse struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}
