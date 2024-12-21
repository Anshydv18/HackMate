package handlers

import (
	requests "NotesBuddy/Model/Requests"
	response "NotesBuddy/Model/Response"
	services "NotesBuddy/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserProfile(c *gin.Context) {
	key := "Create_User_Profile"
	request := &requests.UserProfileRequest{}
	response := &response.BaseResponse{}
	ctx, err := request.Initiate(c, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, key, "failed while initation", request))
		return
	}

	if err := request.Validate(ctx); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, key, err.Error(), request))
		return
	}

	if er := services.CreateUserProfile(ctx, request); er != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, key, er.Error(), request))
		return
	}

	c.JSON(http.StatusOK, response.Success(ctx, key, request))
}
