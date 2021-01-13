package controller

// Response は返却のためのデータです。
type Response struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

// NewResponse はレスポンスを作成するためのメソッドです。
func NewResponse(message string, data interface{}) *Response {
	Response := new(Response)
	Response.Message = message
	Response.Data = data
	return Response
}
