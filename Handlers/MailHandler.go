package handlers

import (
	requests "Hackmate/Model/Requests"
	response "Hackmate/Model/Response"
	services "Hackmate/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context) {
	key := "Sending_Mail"
	request := &requests.MailRequest{}
	response := &response.BaseResponse{}
	ctx, err := request.Initiate(c, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, key, "initiation", request))
		return
	}

	if err := request.Validate(ctx); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, key, "initiation", request))
		return
	}

	if err := services.SendMailService(ctx, request); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, key, err.Error(), request))
		return
	}

	c.JSON(http.StatusOK, response.Success(ctx, key, request))
}

func UploadPhoto(c *gin.Context) {
	key := "upload photo"
	request := &requests.ImageRequest{}
	response := &response.StringResponse{}
	ctx, err := request.Initiate(c, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, key, err.Error(), request))
		return
	}

	if err := request.Validate(ctx); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, key, err.Error(), request))
		return
	}

	url, err := services.UploadPhoto(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, key, err.Error(), request))
		return
	}

	c.JSON(http.StatusOK, response.Success(ctx, key, url))
}
