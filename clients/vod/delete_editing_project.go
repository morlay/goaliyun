package vod

import (
	"github.com/morlay/goaliyun"
)

type DeleteEditingProjectRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ProjectIds           string `position:"Query" name:"ProjectIds"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteEditingProjectRequest) Invoke(client goaliyun.Client) (*DeleteEditingProjectResponse, error) {
	resp := &DeleteEditingProjectResponse{}
	err := client.Request("vod", "DeleteEditingProject", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteEditingProjectResponse struct {
	RequestId goaliyun.String
}
