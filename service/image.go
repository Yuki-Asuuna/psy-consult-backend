package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"psy-consult-backend/exception"
	"psy-consult-backend/utils"
	"psy-consult-backend/utils/upload_image"
)

func ImageUpload(c *gin.Context) {
	f, err := c.FormFile("source")
	if err != nil {
		c.Error(exception.ParameterError())
		return
	}
	url, err := upload_image.GetImageUrl(f)
	if err != nil {
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", url))
}
