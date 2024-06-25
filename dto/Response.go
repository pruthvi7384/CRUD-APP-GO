package dto

type Response struct {
	ErrorDescription string      `json:"message"`
	Data             interface{} `json:"data"`
	ErrorCode        string      `json:"status"`
}
