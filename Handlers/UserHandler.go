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
	response := &response.UserResponse{}
	ctx, err := request.Initiate(c, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, err, "failed while initation", request))
		return
	}

	if err := request.Validate(ctx); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, err, key, request))
		return
	}

	data, err := services.CreateUserProfile(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, err, key, request))
		return
	}

	c.JSON(http.StatusOK, response.Success(ctx, key, data))
}

func Login(c *gin.Context) {
	key := "Login"
	request := &requests.PhoneRequest{}
	response := &response.UserResponse{}
	ctx, err := request.Initiate(c, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, err, "initiation", request))
		return
	}

	if err := request.Validate(ctx); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, err, "initiation", request))
		return
	}

	data, err := services.Login(ctx, request.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, err, key, request))
		return
	}

	c.JSON(http.StatusOK, response.Success(ctx, key, data))
}
