package handlers

import (
	requests "Hackmate/Model/Requests"
	response "Hackmate/Model/Response"
	services "Hackmate/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateHackPost(c *gin.Context) {
	Key := "Create_Hack_Post"
	request := &requests.HackPostRequest{}
	response := &response.BaseResponse{}
	ctx, err := request.Initiate(c, Key)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, Key, err, request))
		return
	}

	if err := request.Validate(ctx); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, Key, err, request))
		return
	}

	if err := services.CreateHackathonPost(ctx, request.HackathonPost); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, Key, err, request))
		return
	}

	c.JSON(http.StatusOK, response.Success(ctx, Key))
}

func FetchUpdatedPost(c *gin.Context) {
	key := "Fetch_All_Post"
	request := &requests.StringRequest{}
	response := &response.HackathonPost{}

	ctx, err := request.Initiate(c, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, key, err))
		return
	}

	if err := request.Validate(ctx); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, key, err))
		return
	}

	data, err := services.FetchAllTrendingPost(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(ctx, key, err))
		return
	}

	c.JSON(http.StatusOK, response.Success(ctx, data))
}
