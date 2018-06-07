package vod

import (
	"github.com/morlay/goaliyun"
)

type UpdateEditingProjectRequest struct {
	CoverURL             string `position:"Query" name:"CoverURL"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Timeline             string `position:"Query" name:"Timeline"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	Title                string `position:"Query" name:"Title"`
	ProjectId            string `position:"Query" name:"ProjectId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UpdateEditingProjectRequest) Invoke(client goaliyun.Client) (*UpdateEditingProjectResponse, error) {
	resp := &UpdateEditingProjectResponse{}
	err := client.Request("vod", "UpdateEditingProject", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateEditingProjectResponse struct {
	RequestId goaliyun.String
}
