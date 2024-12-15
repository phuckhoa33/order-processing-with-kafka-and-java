package wrapper_responses

import "github.com/gin-gonic/gin"

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Data struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Response(context *gin.Context, statusCode int, data interface{}) {
	context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	context.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	context.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization")
	context.JSON(statusCode, data)
}

func ErrorResponse(context *gin.Context, statusCode int, message string) {
	Response(context, statusCode, Error{
		Code:  statusCode,
		Error: message,
	})
}
