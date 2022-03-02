package utils

import "github.com/gin-gonic/gin"

func GenSuccessResponse(code int, message string, result interface{}) gin.H {
	return gin.H{
		"code":    code,
		"message": message,
		"result":  result,
	}
}
