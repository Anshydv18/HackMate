package response

import (
	dto "Hackmate/Model/Dto"
	"context"
)

type HackathonPost struct {
	BaseResponse
	Posts []*dto.HackathonPost `json:"post"`
}

func (res *HackathonPost) Success(ctx *context.Context, data []*dto.HackathonPost) *HackathonPost {
	res.Status = true
	res.Posts = data
	return res
}
