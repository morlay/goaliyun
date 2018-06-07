package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyDeploymentSetAttributeRequest struct {
	DeploymentSetId      string `position:"Query" name:"DeploymentSetId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	DeploymentSetName    string `position:"Query" name:"DeploymentSetName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyDeploymentSetAttributeRequest) Invoke(client goaliyun.Client) (*ModifyDeploymentSetAttributeResponse, error) {
	resp := &ModifyDeploymentSetAttributeResponse{}
	err := client.Request("ecs", "ModifyDeploymentSetAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDeploymentSetAttributeResponse struct {
	RequestId goaliyun.String
}
