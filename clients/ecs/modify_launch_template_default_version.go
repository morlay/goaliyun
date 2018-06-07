package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyLaunchTemplateDefaultVersionRequest struct {
	LaunchTemplateName   string `position:"Query" name:"LaunchTemplateName"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LaunchTemplateId     string `position:"Query" name:"LaunchTemplateId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DefaultVersionNumber int64  `position:"Query" name:"DefaultVersionNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyLaunchTemplateDefaultVersionRequest) Invoke(client goaliyun.Client) (*ModifyLaunchTemplateDefaultVersionResponse, error) {
	resp := &ModifyLaunchTemplateDefaultVersionResponse{}
	err := client.Request("ecs", "ModifyLaunchTemplateDefaultVersion", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyLaunchTemplateDefaultVersionResponse struct {
	RequestId goaliyun.String
}
