package vod

import (
	"github.com/morlay/goaliyun"
)

type AddEditingProjectRequest struct {
	CoverURL             string `position:"Query" name:"CoverURL"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	Timeline             string `position:"Query" name:"Timeline"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	Title                string `position:"Query" name:"Title"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddEditingProjectRequest) Invoke(client goaliyun.Client) (*AddEditingProjectResponse, error) {
	resp := &AddEditingProjectResponse{}
	err := client.Request("vod", "AddEditingProject", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddEditingProjectResponse struct {
	RequestId goaliyun.String
	Project   AddEditingProjectProject
}

type AddEditingProjectProject struct {
	ProjectId    goaliyun.String
	CreationTime goaliyun.String
	ModifiedTime goaliyun.String
	Status       goaliyun.String
	Description  goaliyun.String
	Title        goaliyun.String
}
