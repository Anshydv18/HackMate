package handlers

import (
	requests "Hackmate/Model/Requests"
	response "Hackmate/Model/Response"
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

}
