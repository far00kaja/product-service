package httpservice

var Message = "Success"

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// func ResponseData(httpCode int, data interface{}, err error, c *gin.Context) error {
func ResponseData(data interface{}, err error) Response {
	if err != nil {
		Message = err.Error()
	}
	return Response{
		Data:    data,
		Message: Message,
	}
}
