package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteLaunchTemplateRequest struct {
	LaunchTemplateName   string `position:"Query" name:"LaunchTemplateName"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LaunchTemplateId     string `position:"Query" name:"LaunchTemplateId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteLaunchTemplateRequest) Invoke(client goaliyun.Client) (*DeleteLaunchTemplateResponse, error) {
	resp := &DeleteLaunchTemplateResponse{}
	err := client.Request("ecs", "DeleteLaunchTemplate", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteLaunchTemplateResponse struct {
	RequestId goaliyun.String
}
