package vod

import (
	"github.com/morlay/goaliyun"
)

type GetEditingProjectRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	ProjectId            string `position:"Query" name:"ProjectId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GetEditingProjectRequest) Invoke(client goaliyun.Client) (*GetEditingProjectResponse, error) {
	resp := &GetEditingProjectResponse{}
	err := client.Request("vod", "GetEditingProject", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetEditingProjectResponse struct {
	RequestId goaliyun.String
	Project   GetEditingProjectProject
}

type GetEditingProjectProject struct {
	ProjectId    goaliyun.String
	CreationTime goaliyun.String
	ModifiedTime goaliyun.String
	Status       goaliyun.String
	Description  goaliyun.String
	Title        goaliyun.String
	Timeline     goaliyun.String
	CoverURL     goaliyun.String
}
