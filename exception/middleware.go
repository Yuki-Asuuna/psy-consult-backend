package exception

import (
	"github.com/gin-gonic/gin"
	"time"
)

func ErrorHandlingMiddleware(c *gin.Context) {
	c.Next()
	if err := c.Errors.Last(); err != nil {
		var exception *Exception
		if h, ok := err.Err.(*Exception); ok {
			exception = h
		} else {
			exception = UnknownError(err.Error())
		}
		exception.Path = c.Request.Method + " " + c.Request.URL.String()
		exception.Timestamp = time.Now()
		c.JSON(exception.Code, exception)
		return
	}
}
