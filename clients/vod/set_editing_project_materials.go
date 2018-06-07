package vod

import (
	"github.com/morlay/goaliyun"
)

type SetEditingProjectMaterialsRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MaterialIds          string `position:"Query" name:"MaterialIds"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	ProjectId            string `position:"Query" name:"ProjectId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SetEditingProjectMaterialsRequest) Invoke(client goaliyun.Client) (*SetEditingProjectMaterialsResponse, error) {
	resp := &SetEditingProjectMaterialsResponse{}
	err := client.Request("vod", "SetEditingProjectMaterials", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetEditingProjectMaterialsResponse struct {
	RequestId goaliyun.String
}
